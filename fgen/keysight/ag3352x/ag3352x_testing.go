
package ag3352x

import (
  //
  "testing"
  "bites"

  //subsitute with updated git
  "github.com/danpilgrim/ivi.git"
  )


type FakeInstrument struct {
  BurstState string
  OutputChannel string
  Impedance float64
  BurstCount float64
  DCOffset float64
  DutyCycle float64
  Frequency float64
  Function string
  Symmetry float64
  InternalTriggerPeriod float64
  TriggerSource string
}

func (inst *FakeInstrument) Read(p []byte) (int, error) {

  //prints byte amount read to console
  print("READ: ")
  n := bytes.IndexByte(p, 0)
  str := string(p[:n])
  println(str)

  switch s (
    "BURS:STAT ON"
    )

}

func (inst *FakeInstrument) Write(p []byte) (int, error) {

  //prints input to console
  print("INPUT: ")
  n := bytes.IndexByte(p, 0)
  str := string(p[:n])
  println(str)
}

func (inst *FakeInstrument) WriteString(s string) (int, error) {

  switch s {
  case "":
    return
  }
}

func (inst *FakeInstrument) Query(s string) (string, error) {
  switch s {
  case "VOLT?\n":
    return ("2.5000", nil)
  case "BURS:STAT?\n":
    return (inst.BurstState, nil)
  case "OUTP?\n":
    return (inst.OutputChannel, nil)
  case "OUTP:LOAD?\n":
    return (inst.Impedance, nil)
  case "BURS:NCYC?\n":
    return (inst.BurstCount, nil)
  case "APPL:OFFS?\n":
    return (inst.DCOffset, nil)
  case "FUNC:SQU:DCYC?\n":
    return (inst.DutyCycle, nil)
  case "FREQ?\n":
    return (inst.Frequency, nil)
  case "FUNC?\n":
    return (inst.Function, nil)
  case "FUNC:RAMP:SYMM?\n":
    return (inst.Symmetry, nil)
  case "BURS:INT:PER?\n":
    return (inst.InternalTriggerPeriod, nil)
  case "TRIG:SOUR?\n":
    return (inst.TriggerSource, nil)
  default:
    return ("Command not recognized", !nil)
  }

}

/*
func TestBurstStateOn() {
  fg := FakeInstrument{
    BurstState: "ON",
  }
  // FIXME: Write code to test querying the burst state.

}

func TestBurstStateOff() {
  fg := FakeInstrument{
    BurstState: "OFF",
  }
  // FIXME: Write code to test querying the burst state when it's off.

}
*/

  func TestFunction(t *testing.T) {
    expected:="expected answer"
    actual:= callFunction()
    if actual != extpected (
      t.ErrorF("Test failed",expected,actual)
    )

}
