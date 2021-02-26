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
	if c.rendered != nil {
		return c.rendered
	}

	connections, err := c.fetchConnections()
	if err != nil {
		logging.Error("Unable to fetch active connections: %v", err)
	}
	rows := len(connections)
	table := tview.NewTable().
		SetBorders(true)

	table.SetCell(0, 0,
		tview.NewTableCell("UUID").
			SetAlign(tview.AlignLeft),
	)

	table.SetCell(0, 1,
		tview.NewTableCell("TITLE").
			SetAlign(tview.AlignLeft),
	)

	table.SetCell(0, 2,
		tview.NewTableCell("SIZE ON DISK").
			SetAlign(tview.AlignLeft),
	)

	for r := 0; r < rows; r++ {
		conn := connections[r]
		table.SetCell(r+1, 0,
			tview.NewTableCell(conn.UUID),
			// SetAlign(tview.AlignRight),
		)

		table.SetCell(r+1, 1,
			tview.NewTableCell(conn.Title),
			// SetAlign(tview.AlignRight),
		)

		table.SetCell(r+1, 2,
			tview.NewTableCell(conn.Size),
			// SetAlign(tview.AlignRight),
		)
	}
	c.rendered = table

	return c.rendered
}
