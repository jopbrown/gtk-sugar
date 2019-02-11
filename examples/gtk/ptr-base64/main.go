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
		dialog := gtk.NewMessageDialog(win, gtks.DIALOG_MODAL, gtks.MESSAGE_OTHER, gtks.BUTTONS_CLOSE, "W:\t%v\nH:\t%v", r.Width, r.Height)
		dialog.ConnectDefault(dialog.Destroy)
		dialog.ShowAll()
	})
	win.Add(btn)

	win.ShowAll()

	gtk.Main()

}
