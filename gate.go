package gate

// Gate is a concurrent control tool.
// It should initialize with make(chan struct{}, bufferSize).
type Gate chan struct{}

func NewGate(concurrent uint) Gate {
	return make(Gate, concurrent)
}

// Enter to get channel
func (g Gate) Enter() {
	g <- struct{}{}
}

// Leave to release channel
func (g Gate) Leave() {
	<-g
}
