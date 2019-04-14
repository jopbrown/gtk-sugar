package main

import (
	"fmt"
	"log"
	"sync/atomic"
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
	win.SetTitle("Stopwatch")
	win.SetDefaultSize(400, 200)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	table := gtk.NewTable(10, 10, true)
	win.Add(table)

	startTime := time.Now()
	stop := make(chan bool)
	runFlag := int32(0)

	timeLabel := gtk.NewLabel("")
	timeLabel.SetMarkup(makeupOfElapsed(startTime))
	table.AttachDefaults(timeLabel, 1, 9, 1, 5)

	startBtn := gtk.NewButtonWithLabel("Start")
	startBtn.ConnectDefault(func() {
		log.Println("click start")
		if isSetFlag(&runFlag) {
			return
		}

		go func() {
			log.Println("start")
			setFlag(&runFlag, true)
			for {
				select {
				case <-stop:
					log.Println("got stop")
					setFlag(&runFlag, false)
					return
				case <-time.After(100 * time.Millisecond):
					gtk.Invoke(func() {
						timeLabel.SetMarkup(makeupOfElapsed(startTime))
					})
				}
			}
		}()
	})
	table.AttachDefaults(startBtn, 1, 3, 6, 9)

	stopBtn := gtk.NewButtonWithLabel("Stop")
	stopBtn.ConnectDefault(func() {
		log.Println("click stop")
		if isSetFlag(&runFlag) {
			log.Println("stop")
			stop <- true
		}
	})
	table.AttachDefaults(stopBtn, 4, 6, 6, 9)

	resetBtn := gtk.NewButtonWithLabel("Reset")
	resetBtn.ConnectDefault(func() {
		log.Println("click reset")
		if isSetFlag(&runFlag) {
			log.Println("reset")
			stop <- true
		}
		startTime = time.Now()
		timeLabel.SetMarkup(makeupOfElapsed(startTime))
	})
	table.AttachDefaults(resetBtn, 7, 9, 6, 9)

	win.ShowAll()

	gtk.Main()

	if isSetFlag(&runFlag) {
		log.Println("stop before exit")
		stop <- true
	}
}

func makeupOfElapsed(startTime time.Time) string {
	return fmt.Sprintf("<span font_desc='40'>%s</span>", time.Since(startTime).Round(100*time.Millisecond).String())
}

func setFlag(runFlag *int32, isSet bool) {
	val := int32(0)
	if isSet {
		val = int32(1)
	}
	atomic.StoreInt32(runFlag, val)
}

func isSetFlag(runFlag *int32) bool {
	val := atomic.LoadInt32(runFlag)
	if val == 0 {
		return false
	}

	return true
}
