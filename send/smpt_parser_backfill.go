// +build go.1.1
// +build !go1.5

package send

import "net/mail"

type emailAddressParserImpl struct{}

func (*emailAddressParserImpl) Init() {}

func (*emailAddressParserImpl) Parse(a string) (*mail.Address, error) {
	return mail.ParseAddress(a)
}

func (*emailAddressParserImpl) ParseList(l string) ([]*mail.Address, error) {
	return mail.ParseAddressList(l)
}
