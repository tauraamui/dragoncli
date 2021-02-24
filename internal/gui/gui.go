package gui

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/tauraamui/dragoncli/internal/gui/views"
)

type view interface {
	Name() string
	Render() tview.Primitive
}

type Pages struct {
	*tview.Pages
}

func NewPages() *Pages {
	return &Pages{
		tview.NewPages(),
	}
}

func (p *Pages) Show(v view) {
	p.SwitchToPage(componentID(v))
}

func (p *Pages) add(v view) {
	p.AddPage(componentID(v), v.Render(), true, true)
}

func (p *Pages) addAndShow(v view) {
	p.add(v)
	p.Show(v)
}

type Gui struct {
	*tview.Application
	main  *Pages
	views map[string]view
}

func NewGui() *Gui {
	gui := Gui{
		Application: tview.NewApplication(),
		main:        NewPages(),
	}

	gui.views = map[string]view{
		"login": views.NewLogin(),
	}

	gui.SetRoot(gui.main, true).SetFocus(gui.main)

	return &gui
}

func (g *Gui) SetFocusToPages() {
	g.SetFocus(g.main)
}

func (g *Gui) Show(v view) {
	g.main.addAndShow(v)
}

// View accessors...
func (g *Gui) Login() *views.Login {
	return g.views["login"].(*views.Login)
}

func componentID(v view) string {
	if v.Name() == "" {
		panic("Component has no name")
	}
	return fmt.Sprintf("%s-%p", v.Name(), v)
}
