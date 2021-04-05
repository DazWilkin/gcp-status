package main

import (
	"strings"

	"golang.org/x/net/html"
)

// contains is a func that confirms whether html.Attribute key|value pairs are in an html.Node's attributes
func contains(aa []html.Attribute, key string, value string) bool {
	for _, a := range aa {
		if a.Key == key {
			if a.Val == value {
				return true
			}
		}
	}
	return false
}

// extractServices extracts service name and status from an HTML document
// The structure is not particularly amenable to parsing
// <table>
//   <tbody>
//     <tr>
//       <td class="service-status">${SERVICE_NAME}</td>
//       <td class="day col1">...</td>
//       <td class="day col2">...</td>
//       ...
//       <td class="day col8">
//         <!-- .... -->
//         <!-- .... -->
//         <span class="end-bubble bubble ${STATUS}"</span>
func extractServices(n *html.Node) []Service {
	services := []Service{}

	var tbody, tr, tds func(*html.Node)

	// Not strictly necessary since `tr` groups `tds`
	tbody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tbody" {
			tr(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tbody(c)
		}
	}

	// Each GCP service is represented by a set of TDs within one TR
	tr = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			tds(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			tr(c)
		}
	}

	// These TDs represent the data need for one GCP service
	// One of the TDs will yield the service name
	// Another TD will yield a SPAN representing the service's status
	tds = func(n *html.Node) {
		s := Service{}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			// Only interested in TD (not e.g. TH)
			if c.Data == "td" {
				if contains(c.Attr, "class", "service-status") {
					s.Name = strings.TrimSpace(c.LastChild.Data)
				}
				if contains(c.Attr, "class", "day col8") {
					// Iterate over comments looking for a SPAN
					for d := c.FirstChild; d != nil; d = d.NextSibling {
						if d.Type == html.ElementNode && d.Data == "span" {
							switch val := d.Attr[0].Val; val {
							case "end-bubble bubble ok":
								s.Up = 1.0
							default:
								s.Up = 0.0 // Anything but one
							}
						}
					}
				}
			}
		}
		// The first row contains headers (TH) that we discard
		// Unable to determine a better way to avoid finding the THs
		if s.Name != "" {
			services = append(services, s)
		}
	}

	// Do it
	tbody(n)

	return services
}
