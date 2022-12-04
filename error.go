package polybase

type Error struct {
	Reason  string `json:"reason"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
