package domain

type ActionType string

// Action is a result of Command.
// Must contain Type and Attributes according to that Type.
type Action struct {
	Type       ActionType
	Attributes map[string]any
}
