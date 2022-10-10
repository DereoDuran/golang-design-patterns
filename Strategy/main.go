package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(buffer *strings.Builder)
	End(buffer *strings.Builder)
	AddListItem(buffer *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) End(buffer *strings.Builder) {
}

func (m *MarkdownListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(buffer *strings.Builder) {
	buffer.WriteString("<ul>")
}

func (h *HtmlListStrategy) End(buffer *strings.Builder) {
	buffer.WriteString("</ul>")
}

func (h *HtmlListStrategy) AddListItem(buffer *strings.Builder, item string) {
	buffer.WriteString(" <li>" + item + "</li>")
}

func testStrategy(ls ListStrategy, items []string) {
	var buffer strings.Builder
	ls.Start(&buffer)
	for _, item := range items {
		ls.AddListItem(&buffer, item)
	}
	ls.End(&buffer)
	fmt.Println(buffer.String())
}

func main() {
	text := "alpha bravo charlie delta"
	items := strings.Split(text, " ")

	markdown := &MarkdownListStrategy{}
	html := &HtmlListStrategy{}

	fmt.Println("Testing Markdown Strategy")
	testStrategy(markdown, items)

	fmt.Println("Testing HTML Strategy")
	testStrategy(html, items)

}
