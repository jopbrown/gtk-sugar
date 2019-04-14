package main

import (
	"time"

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

	gtk.GiveCandy(sugar.NewCandy(clt.Conn()))
	gtk.Init()

	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Embedded GTK API demo")
	win.SetDefaultSize(400, 200)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	table := gtk.NewTable(10, 10, true)
	win.Add(table)

	helloLabel := gtk.NewLabel("Hello world")
	table.AttachDefaults(helloLabel, 2, 8, 2, 6)

	dialog := gtk.NewMessageDialog(win, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Hello %s", "world")
	dialog.ConnectDefault(dialog.Hide)

	showBtn := gtk.NewButtonWithLabel("Show dialog")
	showBtn.ConnectDefault(func() {
		go func() {
			time.Sleep(time.Second)
			// the goroutine is not run in the main loop.
			// so it's not safe to do any UI operation.
			// need to invoke to main loop.
			gtk.Invoke(func() { dialog.Run() })
		}()
	})
	table.AttachDefaults(showBtn, 1, 4, 6, 9)

	exitBtn := gtk.NewButtonWithLabel("Exit")
	exitBtn.ConnectDefault(gtk.MainQuit)
	table.AttachDefaults(exitBtn, 6, 9, 6, 9)

	win.ShowAll()

	gtk.Main()

}
