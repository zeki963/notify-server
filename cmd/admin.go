package cmd

import (
	"github.com/awesome-gocui/gocui"
	//"github.com/jroimartin/gocui"
	//"github.com/zorhayashi/notify-server/config"
	"log"
	//"io/ioutil"
	"fmt"
)

var (
	msga = "123"
)

//Admin 管理界面
func Admin() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorYellow
	g.SelFrameColor = gocui.ColorGreen
	g.SetManagerFunc(layout)

	if err := initKeybindings(g); err != nil {
		log.Fatalln(err)
	}

	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Fatalln(err)
	}

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("Notify-Server", 0, 0, 25, 4, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Notify-Server"
		fmt.Fprintln(v, `	  __  ___    ___      `)
		fmt.Fprintln(v, `|\ | /  \  |  | |__  \ / `)
		fmt.Fprintln(v, `| \| \__/  |  | |     |  `)
	}

	if v, err := g.SetView("menu", 0, 5, 25, maxY-7, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Menu"
		v.Highlight = true
		// v.SelBgColor = gocui.ColorBlue
		// v.SelFgColor = gocui.ColorBlack

		fmt.Fprintln(v, "All-Server")
		fmt.Fprintln(v, "A-Server")
		fmt.Fprintln(v, "B-Server")
		fmt.Fprintln(v, "C-Server")
	}
	if v, err := g.SetView("help", 0, maxY-6, 25, int(maxY-1), 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Keybindings"
		fmt.Fprintln(v, "Tab: Switch")
		fmt.Fprintln(v, "^s: Setting")
		fmt.Fprintln(v, "^a: Set mask")
		fmt.Fprintln(v, "^c: Exit")
	}

	if v, err := g.SetView("msgbox", 26, 0, int(maxX-1), int(maxY-1), 0); err != nil {
		v.Title = "msgbox"
		if !gocui.IsUnknownView(err) {
			return err
		}
		if _, err := g.SetCurrentView("msgbox"); err != nil {
			return err
		}

		fmt.Fprintf(v, "%s", msga)
		// b, err := ioutil.ReadFile("demolog.txt")
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Fprintf(v, "%s", b)

		v.Editable = true
		v.Wrap = true
	}

	return nil
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone, mask); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, setting); err != nil {
		return err
	}
	if err := g.SetKeybinding("settingpage", gocui.KeyCtrlV, gocui.ModNone, settingsave); err != nil {
		return err
	}
	if err := g.SetKeybinding("settingpage", gocui.KeyEsc, gocui.ModNone, settingcloss); err != nil {
		return err
	}
	//TAB 交換視窗
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchItem); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", gocui.KeyTab, gocui.ModNone, nu); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlL, gocui.ModNone, testlog); err != nil {
		return err
	}

	return nil
}
