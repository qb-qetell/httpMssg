# Package HTTP Messaging
Package **HTTP Messaging** makes HTTP messaging (API calls) easier.
With a set of just 7 methods, finally saying goodbye to dealing with the complexity of the net/http package.


    // Create (establish) a message
    _e1100 := Mssg_Estb ()
    
    // Set message
    _e1100.SettMssg ([]byte ("hello world!"))
    
    // Set recipient
    _e1100.SettRcpn ("https://hello-world.free.beeceptor.com/my/api/path")
    
    // Set communication method
    _e1100.SettCmmnMthd ("POST")
    
    // Set headers
    _e1100.SettHdrr ("header1", "hello1!")
    _e1100.SettHdrr ("header2", "hello2!")
    _e1100.SettHdrr ("header3", "hello3!")
    _e1100.SettHdrr ("header4", "hello4!")
    
    // Send message
    _e1200, _e1300, _e1400 := _e1100.Send (time.Second * 4)
    _e1450 := ""
    if _e1300 != nil {
    	_e1450 = _e1300.Error ()
    }
    
    // Print response
    _e1500 := fmt.Sprintf ("Outcome: %t; Description: %s", _e1200, _e1450)
    fmt.Println (_e1500)
    if _e1300 == nil {
    	fmt.Println ("Code:", _e1400.ExtrCode ())
    	fmt.Println ("Rspn:", string (_e1400.ExtrRspn ()))
    }

## Author
**Brian Qamardeen**
 - **Email:** brian@qeetell.eon
   **Concert:** [qeetell.one](https://qeetell.one)
   **LinkedIn:** [linkedin.com/in/qeetell](https://www.linkedin.com/in/qeetell/)

## Reference
[https://pkg.go.dev/github.com/qeetell/httpMssg](https://pkg.go.dev/github.com/qeetell/httpMssg)
