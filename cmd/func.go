package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/awesome-gocui/gocui"
)

func switchItem(g *gocui.Gui, cv *gocui.View) error {
	var err error
	var nv *gocui.View
	switch cv.Name() {
	case "meun":
		if nv, err = g.SetCurrentView("logbox"); err != nil {
			return err
		}
	case "logbox":
		if nv, err = g.SetCurrentView("menu"); err != nil {
			return err
		}
	}
	cv.Highlight = false
	nv.Highlight = true
	return nil
}

func mask(g *gocui.Gui, v *gocui.View) error {
	v.Mask ^= '*'
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	os.Exit(0)
	return nil
	//return gocui.ErrQuit
}
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
func getLine(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	v.Line(cy)

	g.SetCurrentView("logbox")

	return nil
}
func setting(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("settingpage", maxX/2-30, maxY/2-10, maxX/2+30, maxY/2+10, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Setting Page ( ESC/Close @ Ctrl+S/Save ) "
		b, err := ioutil.ReadFile("config/config.toml")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(v, "\033[3%d;%dm%s \033[0m\n", int(6), int(1), b)
		v.Editable = true
		v.Wrap = true
		if _, err := g.SetCurrentView("settingpage"); err != nil {
			return err
		}
	}
	return nil
}

func settingsave(g *gocui.Gui, v *gocui.View) error {
	var text string
	var err error

	_, cy := v.Cursor()
	if text, err = v.Line(cy); err != nil {
		text = ""
	}
	s := text
	bs := []byte(s)
	err = ioutil.WriteFile("config/config2.toml", bs, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := g.SetCurrentView("settingpage"); err != nil {
		return err
	}
	if err := g.DeleteView("settingpage"); err != nil {
		return err
	}
	g.SetCurrentView("menu")
	return nil
}

func settingcloss(g *gocui.Gui, v *gocui.View) error {

	if err := g.DeleteView("settingpage"); err != nil {
		return err
	}
	g.SetCurrentView("menu")
	return nil
}

func testlog(g *gocui.Gui, v *gocui.View) error {
	out, _ := g.View("logbox")
	fmt.Fprintln(out, "2020/05/00 00:00:00 [INFO]  TEST ENGmsg/中文訊息")
	return nil
}
