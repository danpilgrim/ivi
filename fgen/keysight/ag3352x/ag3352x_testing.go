
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
  Amplitude float64 // or string
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

func TestBurstStateOn() {
  fg := FakeInstrument{
    BurstState: "ON",
  }
    // FIXME: Write code to test querying the burst state.
}

// Tests OperationMode() from Continuous mode to Burst mode
func TestOperationMode1(){

}

// Tests OperationMode() from Burst mode To Coninuous mode
func TestOperationMode2(){

}

// Tests OperationMode() Error #1 getting operation mode
func TestOperationMode3(){

}

// Tests OperationMode() Error #2 determining operation mode from fgen
func TestOperationMode4(){

}

// Tests SetOperationMode() from Burst mode to Continuous mode
func TestSetOperationMode1(){

}

// Tests SetOperationMode() from Continuous mode to Burst mode
func TestSetOperationMode2(){

}

// Tests OutputEnabled() sets output to on
func TestOutputEnabled1(){

}

// Tests OutputEnabled() sets output to off
func TestOutputEnabled2(){

}

// Tests DisableOutput() disables output
func TestDisableOutput(){

}

// Tests EnableOutput() enables output
func TestEnableOutput(){

}

// Tests OutputImpedance() returns output channel's impedance in ohms.
func TestOutputImpedance(){

}

// Tests SetOperationMode() sets output channel's impedance in ohms
func TestSetOutputImpedance(){

}

// Tests BurstCount() returns the number of waveform cycles that the function generator produces after it receives a trigger
func TestBurstCount(){

}

// Tests SetBurstCount() sets the number of waveform cycles that the function generator produces after it receives a trigger
func TestSetBurstCount(){

}

// Tests Amplitude() read difference between maximum and minimum waveform values
func TestAmplitude(){

}

// Tests SetAmplitude() specifies the difference between the maximum and minimum waveform values
func TestSetAmplitude(){

}

// Tests DCOffset() reads the difference between the average of the maximum and minimum
// waveform values and the x-axis
func TestDCOffset(){

}

// Tests SetDCOffset() sets the difference between the average of the maximum and
// minimum waveform values and the x-axis (0 volts)
func TestSetDCOffset(){

}

// Tests DutyCycle() reads percentage of time (0-100) during one cycle for which the square wave is at its high value
func TestDutyCycle(){

}

// Tests SetDutyCycle() sets the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value
func TesSetDutyCyclet(){

}

// Tests Frequency() reads the number of waveform cycles generated in one second
func TestFrequency(){

}

// Tests SetFrequency()
func TestSetFrequency(){

}

// Tests StandardWaveform() determines if one of the IVI Standard Waveforms is being output by the function generator
func TestStandardWaveform1(){

}

// Tests StandardWaveform() for fgen.Sine
func TestStandardWaveform2(){

}

// Tests StandardWaveform() for fgen.Square
func TestStandardWaveform3(){

}

// Tests StandardWaveform() for fgen.DC
func TestStandardWaveform4(){

}

//Tests StandardWaveform() for error when unable to get symmetry to determine standard waveform: %s
func TestStandardWaveform5(){

}

//Tests StandardWaveform() for fgen.RampDown
func TestStandardWaveform6(){

}

//Tests StandardWaveform() for fgen.RampUp
func TestStandardWaveform7(){

}

//Tests StandardWaveform() for fgen.Triangle
func TestStandardWaveform8(){

}

// Tests StandardWaveform() for error when unable to determine waveform type RAMP with SYMM %f
func TestStandardWaveform9(){

}

// Tests StandardWaveform() for error when unable to determine standard waveform type: %s
func TestStandardWaveform10(){

}

// Tests SetStandardWaveform() to specify which standard waveform the function generator produces
func TestSetStandardWaveform(){

}

// Tests SetStandardWaveform() fgen.Sine test
func TestSetStandardWaveform1(){

}


// Tests SetStandardWaveform() fgen.Square test
func TestSetStandardWaveform2(){

}

// Tests SetStandardWaveform() fgen.Triangle test
func TestSetStandardWaveform3(){

}
// Tests SetStandardWaveform() fgen.RampUp test
func TestSetStandardWaveform4(){

}

// Tests SetStandardWaveform() fgen.Rampdown test
func TestSetStandardWaveform5(){

}

// Tests SetStandardWaveform() fgen.DC test
func TestSetStandardWaveform6(){

}

// Tests ConfigureStandardWaveform() fgen.Sine test
func TestConfigureStandardWaveform1(){

}
// Tests ConfigureStandardWaveform() fgen.Square test
func TestConfigureStandardWaveform2(){

}
// Tests ConfigureStandardWaveform() fgen.Triangle test
func TestConfigureStandardWaveform3(){

}
// Tests ConfigureStandardWaveform() fgen.RampUp test
func TestConfigureStandardWaveform4(){

}
// Tests ConfigureStandardWaveform() fgen.RampDown test
func TestConfigureStandardWaveform5(){

}
// Tests ConfigureStandardWaveform() fgen.DC test
func TestConfigureStandardWaveform6(){

}

// Tests InternalTriggerRate() for successful triggers per second intput
func TestInternalTriggerRate1(){

}
// Tests InternalTriggerRate() for error on ch.QueryFloat
func TestInternalTriggerRate2(){

}

// Tests SetInternalTriggerRate() to show for triggers per second
func TestSetInternalTriggerRate(){

}

// Tests InternalTriggerRate() to set successful triggers per second
func TestInternalTriggerRate1(){

}


// Tests InternalTriggerPeriod() to show internal trigger period in seconds
func TestInternalTriggerPeriod(){

}
// Tests SetInternalTriggerPeriod() to set intenal trigger period in seconds
func TestSetInternalTriggerPeriod(){

}


// Tests TriggerSource() determines trigger Source, error
func TestTriggerSource1(){

}
// Tests TriggerSource() finds source of InternalTrigger
func TestTriggerSource2(){

}
// Tests TriggerSource() finds source of ExternalTrigger
func TestTriggerSource3(){

}
// Tests TriggerSource() finds source of SoftwareTrigger
func TestTriggerSource4(){

}
// Tests TriggerSource() shows error determining trigger source
func TestTriggerSource5(){

}


// Tests SetTriggerSource() sets trigger source to InternalTrigger
func TestSetTriggerSource1(){

}
// Tests SetTriggerSource() sets trigger source to ExternalTrigger
func TestSetTriggerSource2(){

}
// Tests SetTriggerSource() sets trigger source to SoftwareTrigger
func TestSetTriggerSource3(){

}



  func TestFunction(t *testing.T) {
    expected:="expected answer"
    actual:= callFunction()
    if actual != extpected (
      t.ErrorF("Test failed",expected,actual)
    )

}
