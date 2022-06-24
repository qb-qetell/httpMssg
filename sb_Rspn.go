package httpMssg

import (
	"net/http"
)

type Rspn struct {
	Rspn *http.Response
	mssg []byte
}
	func rspn_estb (rspn *http.Response, mssg []byte) (*Rspn) {
		return &Rspn {Rspn: rspn, mssg: mssg}	
	}
	func (i *Rspn) ExtrRspn () ([]byte) {
		return i.mssg
	}
