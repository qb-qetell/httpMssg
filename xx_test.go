package httpMssg

import (
	"fmt"
	"testing"
	"time"
)

func Test_ (t *testing.T) {
	_ka00 := Mssg_Estb ()
	_ka00.SettMssg ([]byte ("hello world!"))
	_ka00.SettRcpn ("https://hello-world.free.beeceptor.com/my/api/path")
	_ka00.SettCmmnMthd ("POST")
	_ka00.SettHdrr ("header1", "12345678")
	_ka00.SettHdrr ("header2", "12345678")
	_ka00.SettHdrr ("header3", "12345678")
	_ka00.SettHdrr ("header4", "12345678")
	_kb00, _kc00 := _ka00.Send (time.Second * 4)
	
	if _kb00 != nil {
		_kd00 := fmt.Sprintf ("Errr: %s", _kb00.Error ())
		fmt.Println (_kd00)
	} else {
		_ke00, _kf00 := _kc00.ExtrRspn ()
		_kg00 := fmt.Sprintf ("Code: %s", _ke00)
		_kh00 := fmt.Sprintf ("Rspn: %s", _kf00)
		fmt.Println (_kg00)
		fmt.Println (_kh00)
	}
}
