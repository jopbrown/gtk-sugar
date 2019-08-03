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
	win.SetTitle("ptr_base64 demo")
	win.SetDefaultSize(100, 100)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	btn := gtk.ButtonNewWithLabel("Size Request")
	btn.ConnectDefault(func() {
		r := win.SizeRequest()
		a := win.GetAllocation()
		dialog := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_OTHER, gtk.BUTTONS_CLOSE, "SizeRequest\t= W: %v H: %v\nGetAllocation\t= X: %v Y: %v W: %v H: %v", r.Width, r.Height, a.X, a.Y, a.Width, a.Height)
		dialog.ConnectDefault(dialog.Hide)
		dialog.Run()
	})
	win.Add(btn)

	win.ShowAll()

	gtk.Main()

}
