package cmp

import (
	"github.com/rivo/tview"
)

type Content struct {
	*tview.Pages
}

func NewContent() *Content {
	pages := &tview.Pages{
		Box: tview.NewBox(),
	}
	pages.SetBorder(true).SetTitle("Content")

	return &Content{pages}
}

func (c *Content) SetContent(content tview.Primitive) {
	c.RemovePage("content")
	c.AddPage("content", content, true, true)
}
