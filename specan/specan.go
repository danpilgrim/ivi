
package specan

// The MeasurementFunction defined values are the available measurement functions.
const (
	Function AmplitudeUnits = iota


	ACVolts
	DCCurrent
	ACCurrent
	TwoWireResistance
	FourWireResistance
	ACPlusDCVolts
	ACPlusDCCurrent
	Frequency
	Period
	Temperature
)
