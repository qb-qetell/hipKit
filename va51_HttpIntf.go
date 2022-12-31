package hipKit

import "crypto/tls"
import "errors"
import "fmt"
import "net"
import "regexp"
import "sync"
import "time"

type HttpIntf struct {
	intfAdrsP1xx     string
	intfAdrsP2xx        int
	srvxQtxx            int
	mssgScrtEnfrStts   bool
	mssgScrtRsrc     tls.Config
	cnfgPrvd         func () (error, interface{})
	mssgHndl         func (interface{}, *Mssg) ()
	lifeStts           bool
	shtdSgnl           bool
}
	func HttpIntf_Estb ()    (*HttpIntf)   {
		return &HttpIntf {
			intfAdrsP1xx:        "",
			intfAdrsP2xx:         0,
			srvxQtxx:             0,
			mssgScrtEnfrStts: false,
			lifeStts:         false,
			shtdSgnl:         false,
		}
	}
	func (i *HttpIntf) SetxIntfAdrs (p1xx string, p2xx int) (error) {
		_ba00 := net.ParseIP ("p1xx")
		if p1xx != "" && _ba00 == nil {
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
		i.srvxQtxx = qtxx
		return        nil
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
		i.mssgScrtEnfrStts = true
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
	func (i *HttpIntf) SetxCnfgPrvd (m func () (error, interface{})) (error) {
		if m == nil {
			_ca00 := fmt.Sprintf ("Config provider is invalid.")
			return errors.New (_ca00)
		}
		i.cnfgPrvd= m
		return nil
	}
	func (i *HttpIntf) SetxMssgHndl (m func (interface{}, *Mssg) ()) (error) {
		if m == nil {
			_ca00 := fmt.Sprintf ("Message handler is invalid.")
			return errors.New (_ca00)
		}
		i.mssgHndl= m
		return nil
	}
	func (i *HttpIntf) Actv () (error,   chan []string) {
		if i.lifeStts == true || i.shtdSgnl == true {
			_ca00 := fmt.Sprintf ("This interface has alread been used.")
			return errors.New (_ca00), nil
		}
		if i.intfAdrsP2xx ==    0 {
			_ca00 := fmt.Sprintf ("Set the interface address.")
			return errors.New (_ca00), nil
		}
		if i.srvxQtxx     ==    0 {
			_ca00 := fmt.Sprintf ("Set the serving quota.")
			return errors.New (_ca00), nil
		}
		if i.mssgHndl     ==  nil {
			_ca00 := fmt.Sprintf ("Set the message handler.")
			return errors.New (_ca00), nil
		}
		/*--1--*/
		_ba00, _bb00:= net.Listen (
			"tcp", fmt.Sprintf("%s:%d", i.intfAdrsP1xx, i.intfAdrsP2xx),
		)
		if _bb00 != nil {
			_ca00 := fmt.Sprintf ("%s", _bb00.Error ())
			return errors.New (_ca00), nil
		}
		/*--1--*/
		i.lifeStts = true
		_bc00 := make (chan []string, 1)
		/*--1--*/
		actvMssgCntx := 0
		go func (i *HttpIntf, l net.Listener, c chan []string, actvMssgCntx *int) {
			actvMssgCntxMtxx := &sync.Mutex {}
			/*--2--*/
			for {
				time.Sleep (time.Millisecond * 1)
				/*--2--*/
				if i.shtdSgnl == true {
					if *actvMssgCntx  ==  0 {
						i.lifeStts = false
						_db00 := time.Now ().In            (
							time.FixedZone ("+0000", 0),
						).Format (
							"2006-01-02 15:04:05 -0700",
						)
						c <- []string {"ba20", _db00}
						return
					} else  { continue }
				}				
				/*--2--*/
				if *actvMssgCntx == i.srvxQtxx { continue }
				/*--2--*/
				_ca01,  _cb00    := l.Accept ()
				if _cb00 != nil  && i.shtdSgnl == true { continue }
				if _cb00 != nil {
					_da00 := fmt.Sprintf (
						"Could not receive an incoming message. [%s]",
						_cb00.Error (),
					)
					_db00 := time.Now ().In            (
						time.FixedZone ("+0000", 0),
					).Format (
						"2006-01-02 15:04:05 -0700",
					)
					c <- []string {"ba10", _da00, _db00}
					return
				}
				/*--2--*/
				_cb10,  _cb20 := i.cnfgPrvd ()
				if _cb10 != nil {
					_da00 := fmt.Sprintf (
						"Could not collect config. [%s]",
						_cb10.Error (),
					)
					_db00 := time.Now ().In            (
						time.FixedZone ("+0000", 0),
					).Format (
						"2006-01-02 15:04:05 -0700",
					)
					c <- []string {"ba10", _da00, _db00}
					return
				}
				/*--2--*/
				_cc00 := _ca01
				_cc00  =  nil
				if i.mssgScrtEnfrStts == true {
					_cc00 = tls.Server (_ca01, &i.mssgScrtRsrc)
				}
				/*--2--*/
				actvMssgCntxMtxx.Lock   ()
				*actvMssgCntx = *actvMssgCntx + 1
				actvMssgCntxMtxx.Unlock ()
				/*--2--*/
				go func (chnl chan []string, actvMssgCntx *int,
				actvMssgCntxMtxx *sync.Mutex, c1xx, c2xx net.Conn,
				cnfg interface{}) {
					_da00 := mssg_Estb (c1xx)
					if c2xx != nil {_da00 = mssg_Estb (c2xx) }
					/*--3--*/
					defer func (actvMssgCntx     *int,
					actvMssgCntxMtxx *sync.Mutex)    {
						actvMssgCntxMtxx.Lock   ()
						*actvMssgCntx = *actvMssgCntx - 1
						actvMssgCntxMtxx.Unlock ()
					} (actvMssgCntx, actvMssgCntxMtxx)
					/*--3--*/
					defer func (c chan []string) {
						   _ea00 := recover ()
						if _ea00 != nil {
							_fa00 := fmt.Sprintf (
								"The handler paniced. [%v]",
								_ea00,
							)
							_fb00 := time.Now ().In (
								time.FixedZone ("+0000", 0),
							).Format (
								"2006-01-02 15:04:05 -0700",
							)
							c <- []string {"ba10", _fa00,_fb00}
						}
					} (chnl)
					/*--3--*/
					defer c1xx.Close ()
					if c2xx != nil { defer c2xx.Close () }
					/*--3--*/
					i.mssgHndl  (cnfg, _da00)
				} (c, actvMssgCntx, actvMssgCntxMtxx, _ca01, _cc00, _cb20)
			}
		} (i, _ba00, _bc00, &actvMssgCntx)
		/*--1--*/
		go func (intf *HttpIntf,    lstn net.Listener,    actvMssgCntx *int  ) {
			for {
				time.Sleep (time.Millisecond * 1)
				if intf.shtdSgnl == true    &&  *actvMssgCntx >      0 {
					continue
				}
				if intf.shtdSgnl == true  {
					lstn.Close    ()
				}
				if intf.shtdSgnl == true    &&  intf.lifeStts == false {
					break
				}
			}
		} (i, _ba00, &actvMssgCntx)
		/*--1--*/
		return nil, _bc00
	} // close conn
	func (i *HttpIntf) Halt () {
		if i.shtdSgnl == true { return }
		i.shtdSgnl = true		
	}

