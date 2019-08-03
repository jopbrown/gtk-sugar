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
