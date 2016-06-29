package links

import (
	"golang.org/x/net/html"
	"os"

	"bytes"
	"fmt"
)

//show a stack of tags
func ShowTags(htmlDoc []byte) {

	buf := bytes.NewBuffer(htmlDoc)
	doc, err := html.Parse(buf)
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

func ShowTagsTree(htmlDoc []byte) {
	buf := bytes.NewBuffer(htmlDoc)
	doc, err := html.Parse(buf)

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
