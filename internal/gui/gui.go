package gui

import (
	"github.com/rivo/tview"
	"github.com/tauraamui/dragoncli/internal/gui/views"
)

type view interface {
	Render() tview.Primitive
}

type Gui struct {
	app     *tview.Application
	pages   *tview.Pages
	uiViews map[string]view
}

func NewGui() *Gui {
	pages := tview.NewPages()
	views := map[string]view{
		"login": &views.Login{},
	}

	for viewName, viewRef := range views {
		pages.AddPage(viewName, viewRef.Render(), true, true)
	}

	app := tview.NewApplication()
	app.SetRoot(pages, true).SetFocus(pages)

	return &Gui{
		app:     app,
		pages:   pages,
		uiViews: views,
	}
}

func (g *Gui) Run() error {
	return g.app.Run()
}

func (g *Gui) ShowLogin() {
	g.pages.SwitchToPage("login")
}

func (g *Gui) Close() {
	g.app.Stop()
}
