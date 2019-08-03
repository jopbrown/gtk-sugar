package sugar

import (
	"io"
	"sync/atomic"
	"time"
)

const (
	_MAINLOOP_TIMEGAP = 10 * time.Millisecond
)

// Candy is a sugar with signal handling and main loop control
type Candy interface {
	Sugar
	Connect(widget, signal string, callback func())
	ConnectDefault(widget string, callback func())
	DisConnect(widget, signal string)
	Invoke(action func())
	RunLoop()
	StopLoop()
	NewWrapper(id string) CandyWrapper
}

type actionRequest struct {
	action func()
	done   chan struct{}
}

type candy struct {
	Sugar
	signalCallback map[string]func()
	requests       chan *actionRequest
	cbRequest      *actionRequest
	isRunning      uint32
	stop           chan struct{}
}

func NewCandy(conn io.ReadWriter) Candy {
	candy := &candy{}
	candy.Sugar = NewSugar(conn)
	candy.signalCallback = make(map[string]func())
	candy.requests = make(chan *actionRequest)

	candy.cbRequest = &actionRequest{done: make(chan struct{}), action: func() {
		signalName := candy.ServerCallback(SERVER_CALLBACK_UPDATE)
		if callback, ok := candy.signalCallback[signalName]; ok {
			callback()
		}
	}}

	return candy
}

func signalName(widget, signal string) string {
	return widget + "_" + signal
}

func (candy *candy) Connect(widget, signal string, callback func()) {
	signalName := signalName(widget, signal)
	candy.ServerConnect(widget, signal, signalName)
	candy.signalCallback[signalName] = callback
}

func (candy *candy) ConnectDefault(widget string, callback func()) {
	candy.signalCallback[widget] = callback
}

func (candy *candy) DisConnect(widget, signal string) {
	signalName := signalName(widget, signal)

	_, ok := candy.signalCallback[signalName]

	if ok {
		candy.ServerDisconnect(widget, signalName)
		delete(candy.signalCallback, signalName)
	}
}

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

func (candy *candy) NewWrapper(id string) CandyWrapper {
	return NewCandyWrapper(candy, id)
}

// CandyWrapper is a base for gtk-server widget
type CandyWrapper interface {
	Candy() Candy
	Connect(signal string, callback func())
	ConnectDefault(callback func())
	DisConnect(signal string)
	ID() string
	String() string
}

type candyWrapper struct {
	id    string
	candy Candy
}

func NewCandyWrapper(candy Candy, id string) CandyWrapper {
	return &candyWrapper{id: id, candy: candy}
}

func (w *candyWrapper) Candy() Candy {
	return w.candy
}

func (w *candyWrapper) Connect(signal string, callback func()) {
	w.candy.Connect(w.String(), signal, callback)
}

func (w *candyWrapper) DisConnect(signal string) {
	w.candy.DisConnect(w.String(), signal)
}

func (w *candyWrapper) ConnectDefault(callback func()) {
	w.candy.ConnectDefault(w.String(), callback)
}

func (w *candyWrapper) ID() string {
	return w.id
}

func (w *candyWrapper) String() string {
	return w.id
}
