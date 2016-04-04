package grip

func (self *Journaler) UseNativeLogger() error {
	// name, threshold, default
	sender, err := send.NewNativeLogger(self.Name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	self.sender = sender
	return err
}
func UseNativeLogger() error {
	return std.UseNativeLogger()
}

func (self *Journaler) UseSystemdLogger() error {
	// name, threshold, default
	sender, err := send.NewJournaldLogger(self.Name, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	if err != nil {
		if self.Sender().Name() == "bootstrap" {
			self.sender = sender
		}
		return err
	}
	self.sender = sender
	return nil
}
func UseSystemdLogger() error {
	return std.UseSystemdLogger()
}

func (self *Journaler) UserFileLogger(filename string) error {
	s, err := send.NewFileLogger(self.Name, filename, self.sender.GetThresholdLevel(), self.sender.GetDefaultLevel())
	if err != nil {
		if self.Sender().Name() == "bootstrap" {
			self.sender = s
		}
		return err
	}
	self.sender = s
	return nil
}

func UseFileLogger(filename string) error {
	return std.UserFileLogger(filename)
}
