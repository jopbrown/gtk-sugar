package sugar

import (
	"sync/atomic"
	"time"
)

func (candy *candy) Invoke(action func()) {
	if atomic.LoadUint32(&candy.isRunning) == 0 {
		panic("The candy loop is not running")
	}

	request := &actionRequest{action: action, done: make(chan struct{})}
	candy.requests <- request
	<-request.done
}

func (candy *candy) RunLoop() {
	if !atomic.CompareAndSwapUint32(&candy.isRunning, 0, 1) {
		panic("The candy loop is already running")
	}
	defer func() {
		atomic.StoreUint32(&candy.isRunning, 0)
	}()

	candy.stop = make(chan struct{})

	// keep sending gtk-server callback request
	go func() {
		for {
			select {
			case candy.requests <- candy.cbRequest:
				<-candy.cbRequest.done
				time.Sleep(_MAINLOOP_TIMEGAP)
			case <-candy.stop:
				return
			}
		}
	}()

mainLoop:
	for {
		select {
		case request := <-candy.requests:
			if nil != request.action {
				request.action()
			}
			request.done <- struct{}{}
		case <-candy.stop:
			break mainLoop
		}
	}
}

func (candy *candy) StopLoop() {
	if atomic.CompareAndSwapUint32(&candy.isRunning, 1, 0) {
		close(candy.stop)
	}
}
