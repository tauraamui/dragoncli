package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/tauraamui/dragoncli/internal/gui"
)

type Login struct {
	app                *tview.Application
	loginForm          *tview.Form
	username, password string
}

func NewLogin(app *tview.Application) *Login {
	login := Login{
		app: app,
	}
	form := tview.NewForm().
		AddInputField("Username", "", 20, nil, login.UpdateUsername).
		AddPasswordField("Password", "", 10, '*', login.UpdatePassword).
		AddButton("Login", login.Login).
		AddButton("Cancel", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Authenticate").SetTitleAlign(tview.AlignLeft)

	login.loginForm = form

	return &login
}

func (l *Login) AddToPages(pages *tview.Pages) {
	pages.AddPage("login", l.loginForm, true, false)
}

func (l *Login) UpdateUsername(username string) {
	l.username = username
}

func (l *Login) UpdatePassword(password string) {
	l.password = password
}

func (l *Login) Login() {
	modal := tview.NewModal()
	modal.SetText(fmt.Sprintf("Authenticating with username: %s and password: %s", l.username, l.password)).
		AddButtons([]string{"Ok"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		})
}

func (l *Login) Active() error {
	return l.app.SetRoot(l.loginForm, true).SetFocus(l.loginForm).Run()
}

func main() {
	app := gui.NewGui()
	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	app.Close()
	// }()
	app.Run()
}
