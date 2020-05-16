package cmd

import (
	"github.com/jroimartin/gocui"
	//"github.com/zorhayashi/notify-server/config"
	"log"
	"fmt"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("help", maxX-23, 0, maxX-1, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "^a: Set mask")
		fmt.Fprintln(v, "^c: Exit")
	}
	if v, err := g.SetView("Menu", maxX-23, 4, maxX-1, 10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "menu"
		fmt.Fprintln(v,"1111")
		fmt.Fprintln(v,"2222")
	}

	if v, err := g.SetView("input", 0, 0, maxX-24, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
		v.Title = "Log's"
		v.Editable = true
		v.Wrap = true
	}

	return nil
}

//Admin 管理界面
func Admin() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Fatalln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}
func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("input", gocui.KeyCtrlA, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			v.Mask ^= '6'
			return nil
		}); err != nil {
		return err
	}
	return nil
}