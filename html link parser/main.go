package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

	func linkNodes(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
			fmt.Println(printText(n))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			linkNodes(c)
		}
	}

func printText(n *html.Node) string {
	if n.Type == html.TextNode {
		str, _ := strings.CutPrefix(n.Data, "\r")
		str, _ = strings.CutSuffix(str, "\r")
		str = strings.TrimSpace(str)
		return str
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += printText(c) + " "
	}
	return ret
}

func main() {
	file, err := os.Open("./example/ex1.html")
	if err != nil {
		log.Fatalf("Unable to open HTML file: %v\n", err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalf("Unable to parse HTML file: %v\n", err)
	}

	linkNodes(doc)
}
