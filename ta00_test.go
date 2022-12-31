package hipKit

import "fmt"
import "testing"
import "time"

func TestTA00 (t *testing.T) {
	_ba00 :=  HttpIntf_Estb/**/ (/*-xx-*/)
	_bb00 := _ba00.SetxIntfAdrs ("", 8080)
	if _bb00 != nil {
		_ca00 :=fmt.Sprintf ("SetxIntfAdrs error. [%s]", _bb00.Error ())
		fmt.Println ( _ca00 )
		return
	}
	_bc00 := _ba00.SetxSrvxQtxx (3)
	if _bc00 != nil {
		_ca00 :=fmt.Sprintf ("SetxSrvxQtxx error. [%s]", _bc00.Error ())
		fmt.Println ( _ca00 )
		return
	}/*
	_bd00 := _ba00.EnfrMssgScrt (
		"localhost",
		"/home/octm_qbqt/xb00_Qetell/cf00_hipKit/k.pem",
		"/home/octm_qbqt/xb00_Qetell/cf00_hipKit/c.pem",
	)
	if _bd00 != nil {
		_ca00 := fmt.Sprintf("EnfrMssgScrt error. [%s]", _bd00.Error ())
		fmt.Println (_ca00)
		return
	}*/
	//_ba00.RlxxMssgScrt ()
	_bd50 := _ba00.SetxCnfgPrvd (cnfg_cnfg)
	if _bd50 != nil {
		_ca00 :=fmt.Sprintf ("SetxCnfgPrvd error. [%s]", _bd50.Error ())
		fmt.Println ( _ca00 )
		return
	}
	_be00 := _ba00.SetxMssgHndl (TA00_TA00)
	if _be00 != nil {
		_ca00 :=fmt.Sprintf ("SetxMssgHndl error. [%s]", _be00.Error ())
		fmt.Println ( _ca00 )
		return
	}
	_bf00,  _bg00 := _ba00.Actv()
	if _bf00 != nil {
		_ca00 :=fmt.Sprintf ("Actv----!!!! error. [%s]", _bf00.Error ())
		fmt.Println ( _ca00 )
		return
	}
	go func () {
		 time.Sleep (time.Second * 800000)
		_ba00.Halt  ()
	} ()
	for {
		_ca00 := <-  _bg00
		fmt.Println ("MSSG:", _ca00)
	}
}
