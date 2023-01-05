package hipKit

import "errors"
import "fmt"
import thxx "github.com/qb-qetell/http/v2"
import "net"
import "time"

type Rspn struct {
	core   net.Conn
	rspn *thxx.Rspn
}
	func rspn_Estb (core net.Conn, rspn []byte) (*Rspn) /*--*/ {
		_ba00 := &Rspn {core: core, rspn: thxx.RSPN_Estb ()}
		_ba00.rspn.Core = rspn
		return _ba00
	}
	func (i *Rspn) SetxCode (code string) { i.rspn.Code = code }
	func (i *Rspn) SetxNote (note string) { i.rspn.Note = note }
	func (i *Rspn) SetxHdrr (name, vlll string) {
		i.rspn.InsrHdrr (name, vlll       )
	}
	func (i *Rspn) Send (wndw ... time.Duration) (error) {
		_ba00    := i.rspn.Lqfy   ()
		/*--1--*/
		_bb00    := time.Second * 10
		if wndw  != nil && len (wndw) > 0 { _bb00 = wndw [0] }
		i.core.SetWriteDeadline (time.Now ().Add (_bb00))
		/*--1--*/
		_, _bc00 := i.core.Write (_ba00)
		if _bc00 != nil {
			_ca00 := fmt.Sprintf (
				"An error occured while trying to send reply. [%s]",
				_bc00.Error (),
			)
			_cb00 := errors.New (_ca00)
			return _cb00
		}
		return nil
	}
