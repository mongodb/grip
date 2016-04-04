package message

import "fmt"

// message.Composer defines an interface with a "Resolve()" method that
// returns the message in string format. Objects that implement this
// interface, in combination to the Compose[*] operations, the
// Resolve() method is only caled if the priority of the method is
// greater than the threshold priority. This makes it possible to
// defer building log messages (that may be somewhat expensive to
// generate) until it's certain that we're going to be outputting the
// message.
type Composer interface {
	Resolve() string
	Loggable() bool

	// A "raw" format of the logging output for use by some Sender
	// implementations that write logged items to interfaces that
	// accept JSON or another structured.

	//	Raw() interface{}
}

func ConvertToComposer(message interface{}) Composer {
	switch message := message.(type) {
	case Composer:
		return message
	case string:
		// we make some weird assumptions here to a level in
		// this conversion, might be messy
		return NewLinesMessage(message)
	case error:
		return NewErrorMessage(message)
	default:
		return NewDefaultMessage(fmt.Sprintf("%v", message))
	}
}
