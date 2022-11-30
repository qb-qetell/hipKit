package hipKit

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
	func mssg_Estb (c net.Conn) (   *Mssg) {
		return &Mssg {
			core:     c,
			lsxx: false,
		}
	}
	func (i *Mssg) ExtrCore  () (net.Conn) { return i.core }
	func (i *Mssg) Read (wndw ... time.Duration) (error, []byte, *http.Rqst) {
		_bb00   := time.Second  * 10
		if wndw != nil && len (wndw) > 0 { _bb00 = wndw [0] }
		_bb50 := time.Now ().Add (_bb00).UnixNano ()
		/*--1--*/
		_bc00 := make ([]byte,    0)
		_bc50 := make ([]byte, 1024)
		for {
			_ca00 := time.Now ().Add (time.Millisecond * 10)
			i.core.SetReadDeadline  (_ca00)
			_ca50, _cb00 := i.core.Read (_bc50)
			/*--2--*/
			if _ca50 != 0  { _bc00 = append (_bc00, _bc50 [0:_ca50]...) }
			/*--2--*/
			if _ca50 < len ( _bc50)  {
				break
			} else if _cb00 != nil  &&  _cb00 == io.EOF {
				break
			} else if _cb00 != nil  &&
				errors.Is (_cb00, os.ErrDeadlineExceeded) ==  true {
				break
			} else if time.Now ().UnixNano () >=  _bb50 {
				_da00 := fmt.Sprintf (
					"BA50: Could not read full message. [%s]",
					_cb00.Error (),
				)
				return errors.New (_da00), _bc00, nil
			} else if _cb00 != nil  &&
				errors.Is (_cb00, os.ErrDeadlineExceeded) == false {
				_da00 := fmt.Sprintf (
					"BB00: An error occured while reading this message. " +
					"[%s]", _cb00.Error (),
				)
				return errors.New (_da00), _bc00, nil
			}
		}
		/*--1--*/
		if len (_bc00) == 0 {
			_ca00  := fmt.Sprintf ("BC00: The client sent an empty message.")
			return errors.New (_ca00),   nil, nil
		}
		/*--1--*/
		_bd00,  _be00 := http.RQST_Sdfy (_bc00 [:])
		if      _bd00 != nil {
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
