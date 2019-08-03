package sugar

import (
	"io"
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

func (candy *candy) NewWrapper(id string) CandyWrapper {
	return NewCandyWrapper(candy, id)
}
