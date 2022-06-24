package httpMssg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Mssg struct {
	Mssg *http.Request
	MssgDlvr *http.Client
}
	func Mssg_Estb () (*Mssg) {
		mssg := Mssg { Mssg: &http.Request {} }
		return &mssg
	}
	func (i *Mssg) SettMssg (d []byte) {
		i.Mssg.Body = ioutil.NopCloser (bytes.NewBuffer (d))
	}
	func (i *Mssg) SettRcpn (d string) {
		_e1100, _e1200 := url.ParseRequestURI (d)
		if _e1200 != nil {
			return
		}
		i.Mssg.URL = _e1100
	}
	func (i *Mssg) SettCmmnMthd (d string) {
		i.Mssg.Method = d
	}
	func (i *Mssg) SettHdrr (name, vlll string) {
		if i.Mssg.Header == nil {
			i.Mssg.Header = make (map[string][]string)
		}
		i.Mssg.Header.Add (name, vlll)
	}
	func (i *Mssg) Send (waitTime time.Duration) (error, *Rspn) {
		_e1100 := &http.Client {
			Transport: &http.Transport {
				ResponseHeaderTimeout: waitTime,
				DisableCompression: true,
			},
		}
		if i.MssgDlvr != nil {
			_e1100 = i.MssgDlvr
		}
		_e1200, _e1300 := _e1100.Do (i.Mssg)
		if _e1300 != nil {
			_e1400 := fmt.Sprintf ("Could not send message. [%s]", _e1300.Error ())
			return errors.New (_e1400), nil
		}
		_e1500, _e1600 := io.ReadAll (_e1200.Body)
		if _e1600 != nil {
			_e1700 := fmt.Sprintf ("Could not read possible response. [%s]",
				_e1600.Error ())
			return errors.New (_e1700), nil
		}
		return nil, rspn_estb (_e1200, _e1500)
	}
