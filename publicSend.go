package grip

// generic base method for sending messages.

func (self *Journaler) Send(priority int, message string) {
	if priority >= 7 || priority < 0 {
		m := "'%d' is not a valid journal priority. Using default %d."
		self.SendDefaultf(m, priority, self.defaultLevel)
		self.SendDefault(message)
	} else {
		self.send(convertPriorityInt(priority, self.defaultLevel), message)
	}
}
func Send(priority int, message string) {
	std.Send(priority, message)
}

// special methods for formating and line printing.

func (self *Journaler) Sendf(priority int, message string, a ...interface{}) {
	self.Sendf(priority, message, a...)
}
func Sendf(priority int, message string, a ...interface{}) {
	std.Sendf(priority, message, a...)
}

func (self *Journaler) Sendln(priority int, message string, a ...interface{}) {
	self.Sendln(priority, message, a...)
}
func Sendln(priority int, message string, a ...interface{}) {
	std.Sendln(priority, message, a...)
}

// default methods for sending messages at the default level, whatever it is.

func (self *Journaler) SendDefault(message string) {
	self.send(self.defaultLevel, message)
}
func SendDefault(message string) {
	std.SendDefault(message)
}
func (self *Journaler) SendDefaultf(message string, a ...interface{}) {
	self.sendf(self.defaultLevel, message, a)
}
func SendDefaultf(message string, a ...interface{}) {
	std.SendDefaultf(message, a...)
}
func (self *Journaler) SendDefaultln(a ...interface{}) {
	self.sendln(self.defaultLevel, a)
}
func SendDefaultln(a ...interface{}) {
	std.SendDefaultln(a...)
}
