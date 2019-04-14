# GTK-sugar

[![Documentation](https://godoc.org/github.com/jopbrown/gtk-sugar?status.svg)](https://godoc.org/github.com/jopbrown/gtk-sugar?)

**GTK-sugar** is a client package of [GTK-server](http://www.gtk-server.org/intro.html) for Go.

The package aims to develop cross-platform GUI app without cgo binding.

## Features:
* No cgo
* Cross-platform
* Supports GTK 1.x/2.x/3.x and so on...
* Provide both interpreted and [embedded GTK](http://www.gtk-server.org/faq.html#embedded) API
* Runtime depend on GTK-server

## Status
Interpreted API is stable enough to run simple applications.

Embedded API is work in process, breaking change will be made.

## Prerequisites

* Install [Go 1.10 or later](http://golang.org/doc/install.html)
* Install GTK+ and GTK-server
	* Windows
		* Install [GTK+ for Windows Runtime Environment Installer](https://github.com/tschoonj/GTK-for-Windows-Runtime-Environment-Installer/releases)
		* Download [prebuilded GTK-server](https://github.com/jopbrown/gtk-server/releases)
	* Others
		* Install [GTK+](https://www.gtk.org/download/)
		* Install [libffi](http://sourceware.org/libffi/)
		* Build GTK-server from [source](http://www.gtk-server.org/download.html)
		```bash
		tar zxvf gtk-server-2.x.x.tar.gz
		cd gtk-server-2.x.x/src
		./configure --with-gtk2 --with-ffi
		make
		make install
		```
* Install GTK-sugar
	```bash
	go get -u github.com/jopbrown/gtk-sugar
	```
* Make `gtk-server.cfg`
	```bash
	cd cfgs
	makecfg
	```

## Examples

### Interpreted API

```go
package main

import (
	"log"

	sugar "github.com/jopbrown/gtk-sugar"
)

func main() {
	log.Println("start gtk-server...")
	clt := sugar.NewClient(sugar.ConnStdin(), sugar.LookupCfg(), sugar.WithDebug(true))
	err := clt.Start()
	if err != nil {
		panic(err)
	}
	defer clt.Stop()

	gui := sugar.NewGuifyer(clt.Conn())

	gui.Guify("gtk_init", nil, nil)
	win := gui.Guify("gtk_window_new", 0)
	gui.Guify("gtk_window_set_title", win, "Interpreted API demo")
	gui.Guify("gtk_window_set_default_size", win, 400, 200)
	gui.Guify("gtk_window_set_position", win, 1)
	table := gui.Guify("gtk_table_new", 10, 10, true)
	gui.Guify("gtk_container_add", win, table)
	helloLabel := gui.Guify("gtk_label_new", "Hello world")
	gui.Guify("gtk_table_attach_defaults", table, helloLabel, 2, 8, 2, 6)
	showBtn := gui.Guify("gtk_button_new_with_label", "Show dialog")
	gui.Guify("gtk_table_attach_defaults", table, showBtn, 1, 4, 6, 9)
	exitBtn := gui.Guify("gtk_button_new_with_label", "Exit")
	gui.Guify("gtk_table_attach_defaults", table, exitBtn, 6, 9, 6, 9)
	gui.Guify("gtk_widget_show_all", win)
	dialog := gui.Guify("gtk_message_dialog_new", win, "GTK_DIALOG_MODAL", "GTK_MESSAGE_INFO", "GTK_BUTTONS_OK", "Hello world", "''")

	log.Println("start main loop...")
	var event sugar.Response
	for event != win && event != exitBtn {
		event = gui.Guify("gtk_server_callback", "WAIT")
		switch event {
		case dialog:
			gui.Guify("gtk_widget_hide", dialog)
		case showBtn:
			gui.Guify("gtk_dialog_run", dialog)
		}
	}

	log.Println("exit gtk-server...")
	gui.Guify("gtk_server_exit")
}
```

### Embedded GTK API

```go
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

	sugar.GiveCandyToEveryone(sugar.NewCandy(clt.Conn()))
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
```

## TODO
* Core
	* [ ] Client - launch and communication to GTK-server
		* [x] stdin(2-way pipes)
		* [x] fifo(named pipe)
		* [ ] ipc(message queue)
		* [ ] tcp/udp
		* [ ] socket
	* [x] Guifyer - interpreted API
	* [x] Sugar - Guifyer with GTK-server internal functions
	* [X] Candy - Sugar with signal handling and main loop control

* Libarys
	* [ ] GTK - WIP
	* [ ] Xforms
	* [ ] Motif

* Other
	* [x] tool to generate `gtk-server.cfg` from code
	* [ ] more comments
	* [ ] more examples
	* [ ] unit tests

## Why?

I’ve been looking for the librarys for native GUI app in Go.
There are already many librarys, like [gotk3](https://github.com/gotk3/gotk3), [go-gtk](https://github.com/mattn/go-gtk), that binding GTK with Go. However, all of them link to libgtk via cgo which is heavy dependency in compile-time. In my usecase, it's hard to build all targets platform because of cgo.

So, I start to look for pure Go librarys for GUI programing. [shiny](https://github.com/golang/exp/tree/master/shiny) is one of librarys do not need cgo, but it already not maintained and not support X-Server with remote display.

And then [duit](https://github.com/mjl-/duit) emerged, which is pure Go, cross-platform, but runtime depend on [devdraw](https://9fans.github.io/plan9port/man/man1/devdraw.html), a standslone tool talking to X window via stdin/stdout. It is nice, but in early development and not friendly to windows. Nevertheless, duit show a way to separate compile-time dependency and move it to runtime.

Finally, I found GTK-server, a standalone tool talking to libgtk via stdin/stdout. This is the beginning of everything.

## What's in a name?
Take from pronouncing sound like `gtk-server` and meaning of syntactic sugar.

##  Contributing
PRs welcome.

## Licence
MIT
