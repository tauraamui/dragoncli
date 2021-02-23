package gui

import (
	"github.com/rivo/tview"
)

type view interface {
	render() tview.Primitive
}

type loginView struct {
	rendered tview.Primitive
	username string
	password string
}

func (l *loginView) updateUsername(username string) { l.username = username }
func (l *loginView) updatePassword(password string) { l.password = password }

// tview.NewForm().
// 	AddInputField("Username", "", 20, nil, l.updateUsername).
// 	AddPasswordField("Password", "", 10, '*', l.updatePassword).
// 	AddButton("Login", nil).
// 	AddButton("Cancel", nil),

func (l *loginView) render() tview.Primitive {
	if l.rendered == nil {
		l.rendered = l.getFormPrimative()
		// l.rendered = (tview.NewFlex().
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 2, false).
		// 	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
		// 		AddItem(
		// 			// tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"),
		// 			l.getFormPrimative(),
		// 			30, 1, true,
		// 		).
		// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 0, 1, false), 0, 1, false).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 0, 2, false))

	}
	return l.rendered
}

func (l loginView) getFormPrimative() tview.Primitive {
	return tview.NewForm().
		AddInputField("Username", "", 30, nil, l.updateUsername).
		AddPasswordField("Password", "", 30, '*', l.updatePassword).
		AddButton("Login", nil).
		AddButton("Cancel", nil)
}

type Gui struct {
	app     *tview.Application
	pages   *tview.Pages
	uiViews map[string]view
}

func NewGui() *Gui {
	pages := tview.NewPages()
	views := map[string]view{
		"login": &loginView{},
	}

	for viewName, viewRef := range views {
		pages.AddPage(viewName, viewRef.render(), true, true)
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
