package hipKit

import "fmt"
import "time"

func TA00_TA00 (c interface {}, m *Mssg) {
	_ac00 := fmt.Sprintf ("%v",  c)
	fmt.Println (_ac00)
	/*--1--*/
	_ak00 := ":---: " + time.Now ().In (time.FixedZone ("+0000", 0)).Format (
		"2006-01-02 15:04:05 -0700",
	) + ":"
	fmt.Println ("HNDL:", _ak00)
	/*--1--*/
	_ba00, _bb00, _bc00 := m.Read ()
	//_ba00, _bb00, _bc00 := m.Read (time.Second *  4)
	if _ba00 !=  nil {
		_ca00 := fmt.Sprintf  ("Read error. [%s]", _ba00.Error ())
		fmt.Println (_ca00 )
		return
	}
	/*--1--*/
	time.Sleep  (time.Second   *   1)
	fmt.Println (string (_bb00     ))
	fmt.Println (string (_bc00.Core))
	_ = _bb00
	_ = _bc00
	/*--1--*/
	_bd00 :=  m.Rply ([]byte ("Thanks!"))
	_be00 := _bd00.Send (time.Second * 1)
	if _be00 != nil {
		_ca00 := fmt.Sprintf  ("Rply error. [%s]", _be00.Error ())
		fmt.Println (_ca00 )
		return
	}
}

func cnfg_cnfg () (error, interface{}) {
	return nil, []string {"aa", "x"}
}
