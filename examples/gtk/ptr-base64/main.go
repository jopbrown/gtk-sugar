package main

import (
	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gtks"
)

func main() {
	clt := sugar.NewClient(sugar.ConnStdin(),
		sugar.WithCfgPath(`../../../cfgs/gtk-server.cfg`))
	err := clt.Start()
	if err != nil {
		panic(err)
	}
	defer clt.Stop()

	gtk := gtks.NewGtk(sugar.NewCandy(clt.Conn()))
	gtk.Init()

	win := gtk.NewWindow(gtks.WINDOW_TOPLEVEL)
	win.SetTitle("ptr_base64 demo")
	win.SetDefaultSize(100, 100)
	win.SetPosition(gtks.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	btn := gtk.NewButtonWithLabel("Size Request")
	btn.ConnectDefault(func() {
		r := win.SizeRequest()
		a := win.GetAllocation()
		dialog := gtk.NewMessageDialog(win, gtks.DIALOG_MODAL, gtks.MESSAGE_OTHER, gtks.BUTTONS_CLOSE, "SizeRequest\t= W: %v H: %v\nGetAllocation\t= X: %v Y: %v W: %v H: %v", r.Width, r.Height, a.X, a.Y, a.Width, a.Height)
		dialog.ConnectDefault(dialog.Hide)
		dialog.Run()
	})
	win.Add(btn)

	win.ShowAll()

	gtk.Main()

}
