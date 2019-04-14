package main

import (
	"log"

	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gdk"
	"github.com/jopbrown/gtk-sugar/lib/gtk"
)

var (
	gtk *gtk.Gtk
	gdk *gdk.Gdk
)

func init() {
	clt := sugar.NewClient(sugar.ConnStdin(),
		sugar.WithCfgPath(`../../../cfgs/gtk-server.cfg`))
	err := clt.Start()
	if err != nil {
		panic(err)
	}

	gtk = gtk.NewGtk(sugar.NewCandy(clt.Conn()))
	gdk = gdk.NewGdk(sugar.NewCandy(clt.Conn()))
}

func main() {
	defer gtk.ServerExit()

	gtk.Init()

	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Show Image")
	win.SetDefaultSize(300, 300)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	pixbuf, err := gdk.NewPixbufFromFile("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	image := gtk.NewImageFromPixbuf(pixbuf)

	layout := gtk.NewLayout(nil, nil)
	layout.Connect("size-allocate", func() {
		a := win.GetAllocation()
		newbuf := pixbuf.ScaleSimple(a.Width, a.Height, gdk.INTERP_BILINEAR)
		image.SetFromPixbuf(newbuf)
		newbuf.Unref()
	})
	layout.Add(image)
	win.Add(layout)

	win.ShowAll()

	gtk.Main()

	pixbuf.Unref()
}
