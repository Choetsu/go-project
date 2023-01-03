package broadcast

type broadcaster struct {
	input chan interface{}        // Input channel for the broadcaster to receive messages
	reg   chan chan<- interface{} // Channel for registering new output channels
	unreg chan chan<- interface{} // Channel for unregistering output channels

	outputs map[chan<- interface{}]bool // Map of output channels to broadcast messages to (true = active, false = inactive)
}

type Broadcaster interface {
	Register(chan<- interface{})   // Register a new output channel (messages will be sent to it)
	Unregister(chan<- interface{}) // Unregister an output channel (no more messages will be sent to it)
	Close() error                  // Close the broadcaster (no more messages will be sent to any output channel)
	Submit(interface{})            // Submit a message to be broadcasted to all output channels
	TrySubmit(interface{}) bool    // Try to submit a message to be broadcasted to all output channels
}

func (b *broadcaster) broadcast(m interface{}) { //
	for ch := range b.outputs {
		ch <- m
	}
}

func (b *broadcaster) run() {
	for {
		select {
		case m := <-b.input: // Receive a message from the input channel
			b.broadcast(m) // Broadcast the message to all output channels
		case ch, ok := <-b.reg: // Receive a new output channel from the reg channel
			if !ok {
				return
			} else {
				b.outputs[ch] = true // Add the output channel to the map of output channels
			}
		case ch := <-b.unreg: // Receive an output channel from the unreg channel
			delete(b.outputs, ch) // Remove the output channel from the map of output channels
		}
	}
}

func NewBroadcaster(buflen int) Broadcaster {
	b := &broadcaster{ // Create a new broadcaster
		input:   make(chan interface{}, buflen),
		reg:     make(chan chan<- interface{}),
		unreg:   make(chan chan<- interface{}),
		outputs: make(map[chan<- interface{}]bool),
	}
	go b.run() // Start the broadcaster
	return b
}

func (b *broadcaster) Register(newch chan<- interface{}) { // Register a new output channel
	b.reg <- newch // Register the new output channel
}

func (b *broadcaster) Unregister(newch chan<- interface{}) { // Unregister an output channel
	b.unreg <- newch // Unregister the output channel
}

func (b *broadcaster) Close() error { // Close the broadcaster
	close(b.reg)   // Close the reg channel
	close(b.unreg) // Close the unreg channel
	// close(b.input) // Close the input channel
	return nil
}

func (b *broadcaster) Submit(m interface{}) { // Submit a message to be broadcasted to all output channels
	if b.input == nil {
		b.input <- m // Send the message to the input channel
	}
}

func (b *broadcaster) TrySubmit(m interface{}) bool { // Try to submit a message to be broadcasted to all output channels
	if b == nil {
		return false
	}

	select {
	case b.input <- m: // Send the message to the input channel
		return true
	default:
		return false
	}
}
