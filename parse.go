package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// parseDashboard extracts service name and status from an HTML document
// The structure is not particularly amenable to parsing
func parse(doc *html.Node) Services {
	services := Services{}
	var body, row func(*html.Node)

	// Dashboard comprises an TABLE/TBODY
	body = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tbody" {
			row(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			body(c)
		}
	}

	// With TRs representing Google Cloud services
	row = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			// Convert the html.Node into a Service or error
			s, err := service(n)
			// If there is no error
			if err == nil {
				// Append it to the list
				services = append(services, s)
			}
		}

		// Repeat
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			row(c)
		}
	}

	// Do it
	body(doc)

	return services
}

// Returns a Google Cloud service (name*regions) or an error
func service(n *html.Node) (Service, error) {
	// Static struct
	// TH: Service name
	// TD: Service Regions (ordered))

	var s Service

	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		th := n.FirstChild.NextSibling
		if th != nil {
			if th.Type == html.ElementNode && th.Data == "th" {
				if th.FirstChild != nil && th.FirstChild.Type == html.TextNode {
					// Only proceed if the Service has a name
					if th.FirstChild.Data != "" && th.FirstChild.Data != "  " {
						s = NewService(th.FirstChild.Data)
						log.Println(s.Name)

						// Won't actually be a TD but sets up the value for the loop
						n := th.NextSibling
						// Iterate over each Region represented by a TD
						s.Regions = regions(n)
					}
				}
			}
		}
	}

	// If the Service has a name
	// And there's at least one Region with a non-null status
	if s.Name != "" && len(s.Regions) > 0 {
		return s, nil
	}

	// Otherwise there's an error
	return s, fmt.Errorf("unable to parse Service from row")
}

func regions(n *html.Node) Regions {
	x := make(Regions)
	for r := Americas; r <= Global; r++ {
		// Move into the next sibling which will be a TD
		n = n.NextSibling
		// Confirm the fact
		if n.Type == html.ElementNode && n.Data == "td" {
			// The TD will always contain an empty (not nil) string ("  ")
			// What's important is whether the FirstChild.NextSibling contains an element
			if n.FirstChild.Type == html.TextNode {
				// If it's null, the status cell for this Region is effectively empty
				// There's no value to report
				// No value should be reported (absense is significant w/ Prometheus)
				if n.FirstChild.NextSibling == nil {
					continue
				}
				// Otherwise
				// The status value is hidden deeper within the element
				if n.FirstChild.NextSibling.Type == html.ElementNode && n.FirstChild.NextSibling.Data == "psd-status-icon" {
					svg := n.FirstChild.NextSibling.FirstChild.NextSibling
					if svg.Type == html.ElementNode && svg.Data == "svg" {
						s, err := status(svg.Attr)
						if err == nil {
							x[r] = s
						}
					}
				}
			}
		}
	}
	return x
}

func status(attrs []html.Attribute) (bool, error) {
	for _, a := range attrs {
		if a.Key != "class" {
			continue
		}

		// key="class"
		if strings.HasSuffix(a.Val, "available") {
			return true, nil
		}
		if strings.HasSuffix(a.Val, "information") {
			// No value
			// Don't report anything
			return false, fmt.Errorf("no status")
		}
		// Else Not available (what is the suffix?)
		return false, nil
	}

	// If here, then the key=class was not found
	return false, fmt.Errorf("unable to find status")
}
