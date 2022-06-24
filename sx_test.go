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
	_ka00.SettHdrr ("hdrr01", "a:123456")
	_ka00.SettHdrr ("hdrr02", "b:123456")
	_ka00.SettHdrr ("hdrr03", "c:123456")
	_ka00.SettHdrr ("hdrr04", "d:123456")
	_kb00, _kc00 := _ka00.Send (time.Second * 4)
	
	if _kb00 != nil {
		_kd00 := fmt.Sprintf ("Errr: %s", _kb00.Error ())
		fmt.Println (_kd00)
	} else {
		_ke00 := _kc00.ExtrRspn ()
		_kg00 := fmt.Sprintf ("Code: %d", _kc00.Rspn.StatusCode)
		_kh00 := fmt.Sprintf ("Rspn: %s", _ke00)
		fmt.Println (_kg00)
		fmt.Println (_kh00)
	}
}
