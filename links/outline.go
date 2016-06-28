package links

import (
	"golang.org/x/net/html"
	"os"

	"fmt"
)

//show a stack of tags
func ShowTags(htmDoc []byte) {

	doc, err := html.Parse(htmDoc)
	if err != nil {
		fmt.Errorf("Error showing tags %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)

}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}

var depth int

func ShowTagsTree(htmDoc []byte) {

	doc, err := html.Parse(htmDoc)

	if err != nil {
		fmt.Errorf("Error showing tags %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)

	}

}
