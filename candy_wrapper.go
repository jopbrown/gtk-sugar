package sugar

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
