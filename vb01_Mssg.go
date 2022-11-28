package main

import "github.com/qb-qetell/http"
import "net"
import "time"

type Mssg struct {
	core net.Conn
}
	func (i *Mssg) GCXX () (net.Conn) { return i.core }
	func (i *Mssg) Read (wndw ... time.Duration) (error, *http.Rqst) { return nil, nil }
	func (i *Mssg) Rply (mssg *http.Rspn, wndw ... time.Duration) (error) { return nil }
