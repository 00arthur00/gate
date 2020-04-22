## Gate

concurrent control Gate.
``` golang
// TODO: better with a http server example
func ExampleGate() {
	g := NewGate(10)
	work := func() {
		//do something here
	}
	for {
		go func() {
			defer g.Leave()
			g.Enter()
			work()
		}()
	}
	select {}
}
```