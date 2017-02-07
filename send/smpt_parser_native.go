// +build go1.5

package send

import "net/mail"

type emailAddressParserImpl struct {
	*mail.AddressParser
}

func (p *emailAddressParserImpl) Init() {
	if p.AddressParser == nil {
		p.AddressParser = &mail.AddressParser{}
	}
}
