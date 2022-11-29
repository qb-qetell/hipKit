package main

import "crypto/tls"
import "errors"
import "fmt"
import "net"
import "regexp"

type HttpIntf struct {
	intfAdrsP1xx     string
	intfAdrsP2xx        int
	srvxQtxx            int
	mssgScrtEnfrStts   bool
	mssgScrtRsrc     tls.Config
	mssgHndl         func (cnnc net.Conn)
	shtdSgnl           bool
	lifeStts           bool
}
	func HttpIntf_Estb ()   (*HttpIntf) {
		return &HttpIntf {
			intfAdrsP1xx:            "",
			intfAdrsP2xx:             0,
			srvxQtxx:                 0,
			mssgScrtEnfrStts:     false,
			shtdSgnl:             false,
			shtdSgnl:             false,
		}
	}
	func (i *HttpIntf) SetxIntfAdrs (p1xx string, p2xx int) (error) {
		_ba00 := net.ParseIP ("p1xx")
		if _ba00 == nil {
			_ca00 := fmt.Sprintf ("Interface Address Part 1 is invalid.")
			return errors.New (_ca00)
		}
		/*--1--*/
		if p2xx   <   1 {
			_ca00 := fmt.Sprintf ("Interface Address Part 2 is invalid.")
			return errors.New (_ca00)
		}
		/*--1--*/
		i.intfAdrsP1xx = p1xx
		i.intfAdrsP2xx = p2xx
		/*--1--*/
		return nil
	}
	func (i *HttpIntf) SetxSrvxQtxx (qtxx int) (error) {
		if qtxx < 1 {
			_ca00 := fmt.Sprintf ("Serving quota is invalid.")
			return errors.New (_ca00)
		}
		return nil
	}
	func (i *HttpIntf) EnfrMssgScrt (name, keyx, crtf string) (error) {
		_ba00, _bb00 := tls.LoadX509KeyPair (crtf, keyx)
		/*--1--*/
		if regexp.MustCompile (``).MatchString ("") == false {
			_ca00 := fmt.Sprintf ("Name is invalid.")
			return errors.New (_ca00)
		}
		/*--1--*/
		if _bb00 != nil {
			_ca00 := fmt.Sprintf ("Key and/or certficate are invalid.")
			return errors.New (_ca00)
		}
		/*--1--*/
		i.mssgScrtEnfrStts =  true
		i.mssgScrtRsrc     = tls.Config {
			Certificates: []tls.Certificate {_ba00},
			ServerName:                       name ,
		}
		/*--1--*/
		return nil
	}
	func (i *HttpIntf) RlxxMssgScrt () {
		i.mssgScrtEnfrStts = false
	}
	func (i *HttpIntf) Actv () (error, chan []string) {
		_ba00, _bb00 := net.Listen ("tcp", i.p1xx + ":" + i.p2xx)
		if _bb00 != nil {
			_ca00 := fmt.Sprintf ("%s", _bb00.Error ())
			return errors.New (_ca00), nil
		}
		/*--1--*/
		i.lifeStts = true
		_bc00 := make (chan []string, 1)
		_bc00 <- []string {"ba10"}
		/*--1--*/
		for {
			
		}
		/*--1--*/
		return nil, nil
	} // close conn
	func (i *HttpIntf) Halt ()  {}
