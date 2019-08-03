package main

import (
	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gtk"
)

func main() {
	clt := sugar.NewClient(sugar.ConnStdin(),
		sugar.WithCfgPath(`../../../cfgs/gtk-server.cfg`))
	err := clt.Start()
	if err != nil {
		panic(err)
	}
	defer clt.Stop()

	sugar.GiveCandy(sugar.NewCandy(clt.Conn()))
	gtk.Init()

	win := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("multi-return demo")
	win.SetDefaultSize(300, 100)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	hbox := gtk.HBoxNew(false, 10)
	win.Add(hbox)

	vbox := gtk.VBoxNew(false, 10)
	hbox.PackStart(vbox, true, true, 50)

	entry := gtk.EntryNew()
	entry.SetText("This is a entry")
	vbox.PackStart(entry, false, false, 10)

	insertBtn := gtk.ButtonNewWithLabel("insert at selected bound")
	insertBtn.ConnectDefault(func() {
		_, start, _ := entry.GetSelectionBounds()
		entry.InsertText("sugar", start)
	})
	vbox.PackStart(insertBtn, false, false, 10)

	win.ShowAll()

	gtk.Main()

}
