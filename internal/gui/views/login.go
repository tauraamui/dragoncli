package views

import "github.com/rivo/tview"

type Login struct {
	rendered tview.Primitive
	username string
	password string
}

func (l Login) Name() string {
	return "login"
}

func (l *Login) Render() tview.Primitive {
	if l.rendered == nil {
		l.rendered = l.getFormPrimative()
	}
	return l.rendered
}

func (l *Login) updateUsername(username string) { l.username = username }
func (l *Login) updatePassword(password string) { l.password = password }

func (l Login) getFormPrimative() tview.Primitive {
	return tview.NewForm().
		AddInputField("Username", "", 30, nil, l.updateUsername).
		AddPasswordField("Password", "", 30, '*', l.updatePassword).
		AddButton("Login", nil).
		AddButton("Cancel", nil)
}
