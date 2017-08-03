
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
  Amplitude float64 //
  Impedance float64 //
  BurstCount float64 //
  DCOffset float64 //
  DutyCycle float64 //
  Frequency float64 //
  Function string
  Symmetry float64 //
  InternalTriggerPeriod float64 //
  TriggerSource string
}

 extractNumber(s string) (float64, error) {

  n = 1
  while (err != nil || n < len(string)){
    str1 := string[0:n]
    strFloat := string[n+1:len(string)-1]
    str2, err := strconv.ParseFloat(strFloat, 64)
    n++
  }
}

// Mock read command
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

// Mock "Write" command
func (inst *FakeInstrument) Write(p []byte) (int, error) {

  //prints input to console
  print("INPUT: ")
  n := bytes.IndexByte(p, 0)
  str := string(p[:n])
  println(str)
}

// Mock "WriteString" command
func (inst *FakeInstrument) WriteString(s string) (int, error) {

string msg = "COMMAND SUCCESS"
string badmsg = "command not recognized"

  //wordArray := strings.Split(s, " ")

  switch s {
  case "BURS:MODE TRIG;STAT ON\n":
    return msg // "Burst mode on, stat one"
  case "BURS:STAT OFF\n":
    return msg // "Burst Mode Off"
  case "OUPT ON\n":
    return msg // Enables output channel
  case "OUTP OFF\n":
    return msg // Disables output channel
  case "OUTP:LOAD %f\n":
    return msg // Sets output channel's impedance in ohms
  case "BURS:NCYC %d\n":
    return msg // Sets number of waveform cycles after trigger
  case "VOLT %f VPP\n":
    return msg // Sets amplitude max/min difference

    //waveformCommands
  case "FUNC SIN\n":
    return msg // fgen.Sine
  case "FUNC SQU\n":
    return msg // fgen.Square
  case "FUNC RAMP; RAMP:SYMM 50\n":
    return msg // fgen.Triangle
  case "FUNC RAMP; RAMP:SYMM 100\n":
    return msg // fgen.RampUp
  case "FUNC RAMP; RAMP:SYMM 0\n":
    return msg // fgen.RampDown
  case "FUNC DC\n":
    return msg // fgen.DC

    //waveformApplyCommands
  case "APPL:SIN %.4f, %.4f, %.4f\n":
    return msg // fgen.Sine
  case "APPL:SQU %.4f, %.4f, %.4f\n":
    return msg // fgen.Square
  case "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 50\n":
    return msg // fgen.Triangle
  case "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 100\n":
    return msg // fgen.RampUp
  case "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 0\n":
    return msg // fgen.RampDown
  case "APPL:DC %.4f, %.4f, %.4f\n":
    return msg // fgen.DC

  case "FUNC:SQU:DCYC %f\n":
    return msg // Sets the percentage of time, specified as 0-100, during one
    // cycle for which the square wave is at its high value
  case "FREQ %f\n":
    return msg // Sets number of waveform cycles per second

  default:
    return nomsg //no command found

  }
}

// Mock Query command
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
  Ag3352x.Channels[0].BurstCount()

  // FIXME: Write code to test querying the burst state when it's off.

}

  func TestFunction(t *testing.T) {
    expected:="expected answer"
    actual:= callFunction()
    if actual != extpected (
      t.ErrorF("Test failed",expected,actual)
    )

}
