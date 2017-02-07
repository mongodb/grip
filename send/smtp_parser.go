package send

import "net/mail"

type emailAddressParser interface {
	Init()
	Parse(string) (*mail.Address, error)
	ParseList(string) ([]*mail.Address, error)
}
