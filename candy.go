package sugar

import (
	"io"
	"time"
)

const (
	MAINLOOP_TIMEGAP = 10 * time.Millisecond
)

// Candy is a sugar with signal handling and main loop control
type Candy interface {
	Sugar
	Connect(widget, signal string, callback func())
	ConnectDefault(widget string, callback func())
	DisConnect(widget, signal string)
	Invoke(callback func())
	Main()
	MainQuit()
	NewWrapper(id string) CandyWrapper
}

type invoke struct {
	callback func()
	done     chan bool
}

type candy struct {
	Sugar
	signalCallback map[string]func()
	signalNameToID map[string]string
	invokeChan     chan *invoke
	quitMain       bool
}

func NewCandy(conn io.ReadWriter) Candy {
	candy := &candy{}
	candy.Sugar = NewSugar(conn)
	candy.signalCallback = make(map[string]func())
	candy.signalNameToID = make(map[string]string)
	candy.invokeChan = make(chan *invoke)

	return candy
}

func signalName(widget, signal string) string {
	return widget + "_" + signal
}

func (candy *candy) Connect(widget, signal string, callback func()) {
	callbackID := candy.ServerConnect(widget, signal)
	candy.signalCallback[callbackID] = callback

	signalName := signalName(widget, signal)
	candy.signalNameToID[signalName] = callbackID
}

func (candy *candy) ConnectDefault(widget string, callback func()) {
	candy.signalCallback[widget] = callback
}

func (candy *candy) DisConnect(widget, signal string) {
	signalName := signalName(widget, signal)
	callbackID, ok := candy.signalNameToID[signalName]

	if ok {
		candy.ServerDisconnect(widget, signal)
		delete(candy.signalCallback, callbackID)
		delete(candy.signalNameToID, signalName)
	}
}

func (candy *candy) Invoke(callback func()) {
	invoke := &invoke{callback: callback, done: make(chan bool)}
	candy.invokeChan <- invoke
	<-invoke.done
}

func (candy *candy) Main() {
	candy.quitMain = false
	for {
		signalID := candy.ServerCallback(SERVER_CALLBACK_UPDATE)
		if callback, ok := candy.signalCallback[signalID]; ok {
			callback()
		}

		select {
		case invoke := <-candy.invokeChan:
			if nil != invoke.callback {
				invoke.callback()
			}
			invoke.done <- true
		default:
			//do nothing
		}

		if candy.quitMain {
			return
		}

		time.Sleep(MAINLOOP_TIMEGAP)
	}
}

func (candy *candy) MainQuit() {
	candy.quitMain = true
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
