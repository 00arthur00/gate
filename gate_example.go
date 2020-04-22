package gate

// TODO: better with a http server example
func ExampleGate() {
	g := NewGate(10)
	work := func() {
		//do something here
	}
	for {
		go func() {
			g.Enter()
			work()
			g.Leave()
		}()
	}
	select {}
}
