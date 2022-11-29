package main

import "errors"
import "fmt"
import "github.com/qb-qetell/http"
import "io"
import "net"
import "os"
import "time"

type Mssg struct {
	core net.Conn
	lsxx     bool
}
	func mssg_Estb ()  (*Mssg) {return   &Mssg {  core: nil,   lsxx: false  }}
	func (i *Mssg) ExtrCore () (net.Conn) { return i.core }
	func (i *Mssg) Read (wndw ... time.Duration) (error, []byte, *http.Rqst) {
		_bb00   := time.Second  * 10
		if wndw != nil && len (wndw) > 0 { _bb00 = wndw [0] }
		i.core.SetReadDeadline (time.Now ().Add (_bb00))
		/*--1--*/
		_bc00   := make ([]byte,  0)
		for {
			   _ca00 := make ([]byte, 1024)
			_, _cb00 := i.core.Read (_ca00)
			/*--2--*/
			if        _cb00 != nil && _cb00 == io.EOF {
				_da00 := 0;
				for _db00 := 1; _db00 <= len (_ca00); _db00 ++ {
					if _ca00 [_db00 - 1] == 0 {
						break
					}
					_da00 = _db00
				}
				if _da00 != 0 { _bc00 = append (_bc00, _ca00 [0:_da00 + 1]...) }
				break
			} else if _cb00 != nil &&
				errors.Is (_cb00, os.ErrDeadlineExceeded) == false {
				_da00 := fmt.Sprintf (
					"BA00: Could not read full message on time.",
				)
				return errors.New (_da00), nil, nil
			} else if _cb00 != nil  {
				_da00 := fmt.Sprintf (
					"BB00: Could not read full message. [%s]",
					_cb00.Error (),
				)
				return errors.New (_da00), nil, nil
			}
			/*--2--*/
			_bc00 = append (_bc00, _ca00...)
		}
		/*--1--*/
		if len (_bc00) == 0 {
			_ca00  := fmt.Sprintf ("BC00: The client sent an empty message.")
			return errors.New (_ca00),   nil, nil
		}
		/*--1--*/
		_bd00,  _be00 := http.RQST_Sdfy (_bc00 [:])
		if      _be00 != nil {
			_ca00 := fmt.Sprintf (
				"BD00: Could not duplicate the raw request as a type '%s' " +
				"value. [%s]", "github.com/qb-qetell/http", _bd00.Error (),
			)
			return errors.New (_ca00), _bc00, nil
		}
		/*--1--*/
		return nil, _bc00, _be00
	}
	func (i *Mssg) Rply (mssg []byte) (*Rspn) {
		return rspn_Estb (i.core, mssg)
	}
