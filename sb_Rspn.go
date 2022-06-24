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
	func (i *Rspn) ExtrRspn () (string, []byte) {
		return strconv.Itoa (i.Rspn.StatusCode), i.mssg
	}
