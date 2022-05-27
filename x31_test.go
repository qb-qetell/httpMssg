package httpMssg

import (
	"fmt"
	"testing"
	"time"
)

func Test_ (t *testing.T) {
	_e1100 := Mssg_Estb ()
	_e1100.SettMssg ([]byte ("hello world!"))
	_e1100.SettRcpn ("https://hello-world.free.beeceptor.com/my/api/path")
	_e1100.SettCmmnMthd ("POST")
	_e1100.SettHdrr ("header1", "12345678")
	_e1100.SettHdrr ("header2", "12345678")
	_e1100.SettHdrr ("header3", "12345678")
	_e1100.SettHdrr ("header4", "12345678")
	_e1200, _e1300, _e1400 := _e1100.Send (time.Second * 4)
	_e1450 := ""
	if _e1300 != nil {
		_e1450 = _e1300.Error ()
	}
	_e1500 := fmt.Sprintf ("Outcome: %t; Description: %s", _e1200, _e1450)
	fmt.Println (_e1500)
	if _e1300 == nil {
		fmt.Println ("Code:", _e1400.ExtrCode ())
		fmt.Println ("Rspn:", string (_e1400.ExtrRspn ()))
	}
}
