package main

import (
	g "qq2564874169/goMiniblink"
	f "qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
)

func main() {
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")
	controls.App.SetBgColor(0x00FF)

	var frm = new(controls.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.EvLoad["add_child"] = func(target interface{}) {
		mb := new(g.MiniblinkBrowser).Init()
		mb.SetSize(750, 425)
		mb.SetLocation(15, 15)
		mb.SetAnchor(f.AnchorStyle_Top | f.AnchorStyle_Right | f.AnchorStyle_Bottom | f.AnchorStyle_Left)
		mb.ResourceLoader = append(mb.ResourceLoader, new(g.FileLoader).Init("Res", "local"))
		mb.LoadUri("https://local/control.html")
		frm.AddChild(mb)
	}
	controls.Run(frm)
}
