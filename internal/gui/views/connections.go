package views

import (
	"github.com/rivo/tview"
	"github.com/tacusci/logging/v2"
	"github.com/tauraamui/dragoncli/internal/common"
)

type Connections struct {
	rendered         tview.Primitive
	fetchConnections func() ([]common.ConnectionData, error)
}

func NewConnections(fetchConnections func() ([]common.ConnectionData, error)) *Connections {
	return &Connections{
		fetchConnections: fetchConnections,
	}
}

func (c Connections) Name() string {
	return "connections"
}

func (c *Connections) Render() tview.Primitive {
	_, err := c.fetchConnections()
	if err != nil {
		logging.Error("Unable to fetch active connections: %v", err)
	}
	if c.rendered == nil {
		c.rendered = tview.NewBox().SetBackgroundColor(tview.Styles.ContrastBackgroundColor)
	}
	return c.rendered
}
