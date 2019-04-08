package main

import (
	"log"

	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gdks"
	"github.com/jopbrown/gtk-sugar/lib/gtks"
)

var (
	gtk *gtks.Gtk
	gdk *gdks.Gdk
)

func init() {
	clt := sugar.NewClient(sugar.ConnStdin(),
		sugar.WithCfgPath(`../../../cfgs/gtk-server.cfg`))
	err := clt.Start()
	if err != nil {
		panic(err)
	}

	gtk = gtks.NewGtk(sugar.NewCandy(clt.Conn()))
	gdk = gdks.NewGdk(sugar.NewCandy(clt.Conn()))
}

func main() {
	defer gtk.ServerExit()

	gtk.Init()

	win := gtk.NewWindow(gtks.WINDOW_TOPLEVEL)
	win.SetTitle("Show Image")
	win.SetDefaultSize(300, 300)
	win.SetPosition(gtks.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	pixbuf, err := gdk.NewPixbufFromFile("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	image := gtk.NewImageFromPixbuf(pixbuf)

	layout := gtk.NewLayout(nil, nil)
	layout.Connect("size-allocate", func() {
		a := win.GetAllocation()
		newbuf := pixbuf.ScaleSimple(a.Width, a.Height, gdks.INTERP_BILINEAR)
		image.SetFromPixbuf(newbuf)
		newbuf.Unref()
	})
	layout.Add(image)
	win.Add(layout)

	win.ShowAll()

	gtk.Main()

	pixbuf.Unref()
}
