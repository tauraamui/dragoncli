package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/tacusci/logging/v2"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		logging.Fatal("unable to load GUI: %v...", err)
	}
	defer g.Close()

	g.SetManagerFunc(login)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		logging.Fatal(err.Error())
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		logging.Fatal("unrecoverable GUI error occurred: %v...", err)
	}
}

func login(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("login", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "Hello World!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
