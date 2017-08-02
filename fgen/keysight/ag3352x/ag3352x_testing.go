
package ag3352x

import (
  //
  "testing"
  "bites"
  "strings"
  //subsitute with updated git
  "github.com/danpilgrim/ivi.git"
  )


type FakeInstrument struct {
  BurstState string
  OutputChannel string
  Amplitude float64
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
func extractLastNumber(s string, int n) (float64, error) {
  while
  str1, err := string[0:n]
  strFloat, err := string[n+1:len(string)-1]
  str2, err := strings.parseFloat(strFloat)

}

func (inst *FakeInstrument) Write(p []byte) (int, error) {

  //prints input to console
  print("INPUT: ")
  n := bytes.IndexByte(p, 0)
  str := string(p[:n])
  println(str)
}

func (inst *FakeInstrument) WriteString(s string) (int, error) {

string msg = "COMMAND SUCCESS"
string badmsg = "command not recognized"
  switch s {
  case "BURS:MODE TRIG;STAT ON\n":
    return msg //"Burst mode on, stat one"
  case "BURS:STAT OFF\n":
    return msg //"Burst Mode Off"
  case "OUPT ON\n":
    return msg //Output Channel enabled
  case "OUTP OFF\n":
    return msg //Output Channel disabled
  case "BURS:MODE TRIG;STAT ON\n":
    return msg //
  case "BURS:MODE TRIG;STAT ON\n":
    return msg
  case "BURS:MODE TRIG;STAT ON\n":
    return msg
  case "BURS:MODE TRIG;STAT ON\n":
    return msg


  }
}

func (inst *FakeInstrument) Query(s string) (string, error) {
  switch s {
  case "VOLT?\n":
    return (inst.Amplitude, nil)
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
