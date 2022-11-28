package main

import "errors"
import "net"

type HIXX struct {
	intfAdrsP1xx     string
	intfAdrsP2xx        int
	srvnQtxx            int
	mssgScrtEnfrStts   bool
	mssgScrtKeyx     []byte
	mssgScrtCrtf     []byte
	mssgHndl         func (cnnc net.Conn)
}
	func HIXX_Estb ()   (*HIXX) {
		return &HIXX  {
		
			srvnQtxx:                 0,
			mssgScrtEnfrStts:     false,
			mssgScrtKeyx:     []byte {},
			mssgScrtCrtf:     []byte {},
		}
	}
	func (i *HIXX) SetxIntfAdrs (p1xx string, p2xx int) (error) { return nil }
	func (i *HIXX) SetxSrvnQtxx (qtxx int) (error) {
		if qtxx < 1 {
			i.srvnQtxx = 0
			return errors.New (
				"Serving quota provided is less than 0: value must be a " +
				"positive number.",
			)
		}
		return nil
	}
	func (i *HIXX) EnfrCnctScrt (keyx, crtf string) (error) { return nil }
	func (i *HIXX) RlxxCnctScrt () {
		i.mssgScrtEnfrStts =     false
		i.mssgScrtKeyx     = []byte {}
		i.mssgScrtCrtf     = []byte {}
	}
	func (i *HIXX) Strt () (error, chan []string) { return nil, nil}
	func (i *HIXX) Halt ()  {}
