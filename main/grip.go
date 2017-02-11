package main

import (
	"github.com/tychoish/grip"
	"github.com/tychoish/grip/send"
)

func main() {
	grip.Info("hello world, default")

	opts := send.SMTPOptions{
		Name:   "grip email",
		From:   "garen@tychoish.com",
		Server: "10.8.0.1",
		TruncatedMessageSubjectLength: 5,
	}
	grip.CatchError(opts.AddRecipient("sam", "tycho@10gen.com"))
	sender, err := send.MakeSMTPLogger(&opts)
	if err != nil {
		grip.EmergencyFatal(err)
	}

	grip.SetSender(sender)
	grip.Info("test email tycho")
}
