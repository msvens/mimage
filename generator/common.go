package generator

import (
	"fmt"
	"golang.org/x/net/html"
)

type Matcher = func(node *html.Node) bool

func findNodes(node *html.Node, matcher Matcher) []*html.Node {
	var ret []*html.Node

	var f func(*html.Node)
	f = func(n *html.Node) {
		if matcher(n) {
			ret = append(ret, n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(node)
	return ret
}

func findNode(node *html.Node, matcher Matcher) (*html.Node, error) {
	var ret *html.Node

	var f func(*html.Node)
	f = func(n *html.Node) {
		if ret != nil {
			return
		}
		if matcher(n) {
			ret = n
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(node)
	if ret == nil {
		return ret, fmt.Errorf("Node not found")
	} else {
		return ret, nil
	}
}

func findSibling(node *html.Node, matcher Matcher) (*html.Node, error) {
	var ret *html.Node

	var f func(*html.Node)
	f = func(n *html.Node) {
		if ret != nil {
			return
		}
		if matcher(n) {
			ret = n
		} else {
			for c := n.NextSibling; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(node)
	if ret == nil {
		return ret, fmt.Errorf("Node not found")
	} else {
		return ret, nil
	}
}
