
package pwrmeter

type OperationState int

const (
	Function OperationState = iota
	Arbitrary
	Sequence
)
