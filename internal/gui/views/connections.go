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
	connections, err := c.fetchConnections()
	if err != nil {
		logging.Error("Unable to fetch active connections: %v", err)
	}
	if c.rendered == nil {
		rows := len(connections)
		table := tview.NewTable()

		table.SetCell(0, 0,
			tview.NewTableCell("UUID").
				SetAlign(tview.AlignCenter),
		)

		table.SetCell(0, 1,
			tview.NewTableCell("TITLE").
				SetAlign(tview.AlignCenter),
		)

		table.SetCell(0, 2,
			tview.NewTableCell("DISK SIZE").
				SetAlign(tview.AlignCenter),
		)

		for r := 0; r < rows; r++ {
			conn := connections[r]
			table.SetCell(r+1, 0,
				tview.NewTableCell(conn.UUID).
					SetAlign(tview.AlignCenter),
			)

			table.SetCell(r+1, 1,
				tview.NewTableCell(conn.Title).
					SetAlign(tview.AlignCenter),
			)

			table.SetCell(r+1, 2,
				tview.NewTableCell(conn.Size).
					SetAlign(tview.AlignCenter),
			)
		}
		c.rendered = table
	}
	return c.rendered
}
