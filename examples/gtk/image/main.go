package main

import (
	"log"

	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gdk"
	"github.com/jopbrown/gtk-sugar/lib/gtk"
)

func init() {
	clt := sugar.NewClient(sugar.ConnStdio())
	err := clt.Start()
	if err != nil {
		panic(err)
	}

	sugar.GiveCandy(sugar.NewCandy(clt.Conn()))
}

func main() {
	defer gtk.Candy().ServerExit()

	gtk.Init()

	win := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Show Image")
	win.SetDefaultSize(300, 300)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	pixbuf, err := gdk.PixbufNewFromFile("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	image := gtk.ImageNewFromPixbuf(pixbuf)

	layout := gtk.LayoutNew(nil, nil)
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
