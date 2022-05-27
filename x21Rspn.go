package httpMssg

import (
	"net/http"
	"strconv"
)

type Rspn struct {
	Rspn *http.Response
	mssg []byte
}
	func rspn_estb (rspn *http.Response, mssg []byte) (*Rspn) {
		return &Rspn {Rspn: rspn, mssg: mssg}	
	}
	func (i *Rspn) ExtrCode () (string) {
		return strconv.Itoa (i.Rspn.StatusCode)
	}
	func (i *Rspn) ExtrRspn () ([]byte) {
		return i.mssg
	}
	