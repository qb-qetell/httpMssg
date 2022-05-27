package httpMssg

type Rspn struct {
	code string
	rspn []byte
}
	func rspn_estb (code string, rspn []byte) (*Rspn) {
		return &Rspn {code: code, rspn: rspn}	
	}
	func (i *Rspn) ExtrCode () (string) {
		t := i.code
		i.code = ""
		return t
	}
	func (i *Rspn) ExtrRspn () ([]byte) {
		t :=  i.rspn
		i.rspn = nil
		return t
	}
	