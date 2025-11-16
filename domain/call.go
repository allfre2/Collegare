package domain

// Protocols
const (
	WebRTC = "webrtc"
)

// Call modes
const (
	Automated = "automated" // IVR ?
	AdHoc     = "adhoc"     // 1 to 1
	Agent     = "agent"
)

type CallConfiguration struct {
	Protocol   string
	Parameters map[string]any
}
