References:
  - [https://golang.org/doc/effective_go.html#chan_of_chan]
  - The Go Programming Book section 8.4 channels

Channels are a safe communication (protocol/data structure/mechanism?) for communicating between goroutines. A channel is a first-class value that can be a type, a parameter thats passed arond, allocated. A channel is a conduit for a particular type called an element type. It is the value that is being passed over the channel. This affects the type of channel it is, eg chan int, int is the element being passed over the channel and the type is "chan int".

To create a channel:
  Use the built-in function make. Make is used to create maps, slices, and channels which have an underlying structure and acts as a reference to that.

  ch := make(chan int) 

  ZERO value of channel is nil. Two channels with the same element type can be compared with ==

  A channel has 2 principal operations, send and receive, collectively know as 'communications'. A send statement transmits a value from one goroutine through the channel to another goroutine executing a correspoinding receive expression. Both use the <- operator. In a send statement the <- separates the channel and value operands. In a receive the <- precedes the channel operand. A receive expression whose result is not used is a valid statement.
      ch <- x // a send statement
      x =<- // a receive expression in an assignment statement
      <- ch // a receive statement, result is discarded

Channels support a third operation 'close' which sets a flag indicating taht no more values will ever be sent on this channel, subsequent attempts to send will panic.