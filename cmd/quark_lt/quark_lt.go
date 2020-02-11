package quark_lt

type QuarkLT struct {
	ApiServerStatus  chan bool
	ControllerStatus chan bool
	WebServerStatus  chan bool
}
