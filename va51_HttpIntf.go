package main

import "errors"
import "net"

type HttpIntf struct {
	intfAdrsP1xx     string
	intfAdrsP2xx        int
	srvxQtxx            int
	mssgScrtEnfrStts   bool
	mssgScrtKeyx     []byte
	mssgScrtCrtf     []byte
	mssgHndl         func (cnnc net.Conn)
}
	func HttpIntf_Estb ()   (*HttpIntf) {
		return &HttpIntf {
			intfAdrsP1xx:            "",
			intfAdrsP2xx:             0,
			srvxQtxx:                 0,
			mssgScrtEnfrStts:     false,
			mssgScrtKeyx:     []byte {},
			mssgScrtCrtf:     []byte {},
		}
	}
	func (i *HttpIntf) SetxIntfAdrs (p1xx string, p2xx int) (error) { return nil }
	func (i *HttpIntf) SetxSrvxQtxx (qtxx int) (error) {
		if qtxx < 1 {
			i.srvxQtxx = 0
			return errors.New (
				"Serving quota provided is less than 0: value must be a " +
				"positive number.",
			)
		}
		return nil
	}
	func (i *HttpIntf) EnfrCnctScrt (keyx, crtf string) (error) { return nil }
	func (i *HttpIntf) RlxxCnctScrt () {
		i.mssgScrtEnfrStts =     false
		i.mssgScrtKeyx     = []byte {}
		i.mssgScrtCrtf     = []byte {}
	}
	func (i *HttpIntf) Strt () (error, chan []string) { return nil, nil}
	func (i *HttpIntf) Halt ()  {}
