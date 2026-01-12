package domain

// Protocols
const (
	WebRTC = "webrtc"
)

// Call modes
const (
	AdHoc = "adhoc" // 1 to 1
)

type CallConfiguration struct {
	Protocol   string
	Parameters map[string]any
}
