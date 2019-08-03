package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	sugar "github.com/jopbrown/gtk-sugar"

	"github.com/jopbrown/gtk-sugar/lib/gtk"
)

var sw = NewStopWatch()

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
	win.SetTitle("Stopwatch")
	win.SetDefaultSize(400, 200)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.ConnectDefault(gtk.MainQuit)

	table := gtk.TableNew(10, 10, true)
	win.Add(table)

	timeLabel := gtk.LabelNew("")
	timeLabel.SetMarkup(makeupOfElapsed(0))
	table.AttachDefaults(timeLabel, 1, 9, 1, 5)

	startBtn := gtk.ButtonNewWithLabel("Start")
	startBtn.ConnectDefault(func() {
		log.Println("click start")
		if sw.IsRunning() {
			return
		}

		go func() {
			log.Println("start")
			sw.Start()
			for sw.IsRunning() {
				gtk.Invoke(func() {
					timeLabel.SetMarkup(makeupOfElapsed(sw.Elapsed()))
				})
				time.Sleep(100 * time.Millisecond)
			}
		}()
	})
	table.AttachDefaults(startBtn, 1, 3, 6, 9)

	stopBtn := gtk.ButtonNewWithLabel("Stop")
	stopBtn.ConnectDefault(func() {
		log.Println("click stop")
		if sw.IsRunning() {
			log.Println("stop")
			sw.Stop()
		}
	})
	table.AttachDefaults(stopBtn, 4, 6, 6, 9)

	resetBtn := gtk.ButtonNewWithLabel("Reset")
	resetBtn.ConnectDefault(func() {
		log.Println("click reset")
		if sw.IsRunning() {
			log.Println("reset")
			sw.Reset()
		}
		timeLabel.SetMarkup(makeupOfElapsed(0))
	})
	table.AttachDefaults(resetBtn, 7, 9, 6, 9)

	win.ShowAll()

	gtk.Main()
}

func makeupOfElapsed(elapsed time.Duration) string {
	return fmt.Sprintf("<span font_desc='40'>%s</span>", elapsed.Round(100*time.Millisecond).String())
}

type StopWatch struct {
	startTime time.Time
	stopTime  time.Time
	isRunning bool
	m         sync.RWMutex
}

func NewStopWatch() *StopWatch {
	return &StopWatch{}
}

func (sw *StopWatch) IsRunning() bool {
	sw.m.RLock()
	defer sw.m.RUnlock()
	return sw.isRunning
}

func (sw *StopWatch) Start() {
	sw.m.Lock()
	defer sw.m.Unlock()
	sw.isRunning = true
	if sw.startTime.IsZero() {
		sw.startTime = time.Now()
		sw.stopTime = time.Time{}
		return
	}

	sw.startTime = sw.startTime.Add(time.Since(sw.stopTime))
}

func (sw *StopWatch) Stop() {
	sw.m.Lock()
	defer sw.m.Unlock()
	sw.isRunning = false
	sw.stopTime = time.Now()
}

func (sw *StopWatch) Reset() {
	sw.m.Lock()
	defer sw.m.Unlock()
	sw.isRunning = false
	t := time.Time{}
	sw.startTime = t
	sw.stopTime = t
}

func (sw *StopWatch) Elapsed() time.Duration {
	sw.m.RLock()
	defer sw.m.RUnlock()

	if sw.isRunning {
		return time.Since(sw.startTime)
	}

	if sw.startTime.IsZero() {
		return 0
	}

	return sw.stopTime.Sub(sw.startTime)
}
