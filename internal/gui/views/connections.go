package views

import "github.com/rivo/tview"

type Connections struct {
	rendered tview.Primitive
}

func NewConnections() *Connections {
	return &Connections{}
}

func (c Connections) Name() string {
	return "connections"
}

func (c *Connections) Render() tview.Primitive {
	if c.rendered == nil {
		c.rendered = tview.NewBox().SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	}
	return c.rendered
}
