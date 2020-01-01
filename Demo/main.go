package main

import (
	"GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows"
	"GoMiniblink/Forms"
	"time"
)

func main() {
	Forms.Provider = new(Windows.Provider).Init()
	Forms.Provider.SetIcon("app.ico")

	var frm = new(Forms.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(GoMiniblink.Rect{Wdith: 300, Height: 500})

	frm.EvLoad["load"] = func(target interface{}) {
		println("onload")
		go func(f *Forms.Form) {
			time.Sleep(3 * time.Second)
			f.Invoke(func(state interface{}) {
				f.ShowInTaskbar(false)
				//f.SetIconVisable(false)
			}, nil)
		}(target.(*Forms.Form))
	}

	Forms.Run(frm)
}
