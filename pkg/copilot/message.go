package copilot

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
)

// Message represents a role-based chat message that can be
// translated into provider-specific request structures.
type Message struct {
	Role    string
	Content string
}
