
package scope

type CurrentLimitBehavior int

// These are the defined values for the Current Limit Behavior.
const (
	Trip CurrentLimitBehavior = iota
	Regulate
)
