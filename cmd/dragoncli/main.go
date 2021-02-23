package main

import (
	"github.com/tauraamui/dragoncli/internal/gui"
)

func main() {
	app := gui.NewGui()
	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	app.Close()
	// }()
	app.Run()
}
