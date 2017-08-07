
package ag3352x

import (
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

  // takes off "\n" ending for strings
  s == strings.TrimSuffix(s, "\n")

  // return messages
  string msg = "COMMAND SUCCESS"
  string badmsg = "command not recognized"

  wordArray := strings.Split(s, " ")

  switch wordArray[0] {
  case "BURS:MODE": // "BURS:MODE TRIG;STAT ON\n" original input
    if wordArray[1] != "TRIG;STAT" {
      return badmsg
    }
    if wordArray[2] != "ON" {
      return badmsg
    }
    return msg // "Burst mode on, stat one"

  case "BURS:STAT": // "BURS:STAT OFF\n" original input
  if wordArray[2] != "ON" {
    return badmsg
  }
  return msg // "Burst Mode Off"

  case "OUPT": // "OUPT ON\n" || "OUTP OFF\n" original input
    if wordArray[1] == "ON\n" {
      return msg // Output channel enabled
    }else if wordArray[1] == "OFF\n" {
      return msg // output channel disabled
    }else{
      return badmsg // Second string not recognized
    }

    // Sets output channel's impedance in ohms
  case "OUTP:LOAD": // "OUTP:LOAD %f\n"  original input
    percent, err = strconv.ParseFloat(wordArray[1], 64)
    if err == nil{
      return msg
    } else {
      return badmsg // Sets amplitude max/min difference
    }
    return badmsg

    // Sets number of waveform cycles after trigger
  case "BURS:NCYC": // "BURS:NCYC %d\n" original input
    percent, err = strconv.ParseFloat(wordArray[1], 64)
    if err == nil{
      return msg
      } else {
        return badmsg // Sets amplitude max/min difference
      }
    return badmsg

  case "VOLT": // "VOLT %f VPP\n" original input

    // test for VPP
    if wordArray[2] != "VPP" {
      return badmsg
    }

    // test for int
    percent, err = strconv.ParseFloat(wordArray[1], 64)
    if err == nil{
      return msg
    } else {
      return badmsg // Sets amplitude max/min difference
    }

    // WaveformCommands fgen.Sine, fgen.Square
  case "FUNC":
     // "FUNC SIN\n" original input, fgen.Sine
    if wordArray[1] == "SIN" {
      return msg
    // "FUNC SQU\n" original input, fgen.Square
    } else if wordArray[1] == "SQU" {
      return msg
    } else {
      return badmsg
    }

  // Waveform Command fgen.Triangle, fgen.RampUp, fgen.RampDown
  case "FUNC RAMP;":
    // Checks second section of command
    if wordArray[1] != "RAMP:SYMM" {
      return badmsg
    }

    // BELOW Checks third section of command
    percent, err = strconv.ParseInt(wordArray[2], 64)

    // "FUNC RAMP; RAMP:SYMM 0\n" original input, fgen.RampDown
    if percent == 0 {
      return msg
    // "FUNC RAMP; RAMP:SYMM 50\n" original input, fgen.Triangle
    } else if percent == 50 {
      return msg
    // "FUNC RAMP; RAMP:SYMM 100\n" original input, fgen.RampUp
    } else if percent == 100 {
      return msg
    // Error returns bad message
    } else {
      return badmsg
    }

  case "FUNC DC": // "FUNC DC\n" original input
    return msg // fgen.DC

    //waveformApplyCommands
  case "APPL:SIN": //"APPL:SIN %.4f, %.4f, %.4f\n": fgen.Sine
    wordArray[1] = strings.TrimSuffix(golang, ",")
    wordArray[2] = strings.TrimSuffix(golang, ",")
    wordArray[3] = wordArray[3]
    return msg

  case "APPL:SQU": //"APPL:SQU %.4f, %.4f, %.4f\n": fgen.Square
    wordArray[1] = strings.TrimSuffix(golang, ",")
    wordArray[2] = strings.TrimSuffix(golang, ",")
    wordArray[3] = wordArray[3]
    return msg

  case "APPL:RAMP": //"APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 50\n":, fgen.Triangle
  //"APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 100\n":// fgen.RampUp
  //"APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 0\n": return msg // fgen.RampDown

    //extracts 1st, 2nd, and 4th float values
    wordArray[1] = strings.TrimSuffix(golang, ",")
    wordArray[2] = strings.TrimSuffix(golang, ",")
    num1, err = strconv.ParseFloat(wordArray[1], 64)
    if err != nil {return badmsg}
    num2, err = strconv.ParseFloat(wordArray[2], 64)
    if err != nil {return badmsg}
    symmetryNum, err = strconv.ParseFloat(wordArray[4], 64)
    if err != nil {return badmsg}

    //takes 3rd number
    wordArray[3] = wordArray[3]
    subString := strings.TrimSuffix(wordArray[3], ";")
    num3, err = strconv.ParseInt(wordArray[2], 64)
    if err != nil {return badmsg}

    //symmetry number
    if symmetryNum == 100.0 {
      return msg;
    } else if symmetryNum == 50.0 {
      return msg;
    } else if symmetryNum == 0.0 {
      return msg;
    } else {
      return badmsg;
    }
    return msg // fgen.Triangle

  case "APPL:DC": // "APPL:DC %.4f, %.4f, %.4f\n", fgen.DC
    wordArray[1] = strings.TrimSuffix(golang, ",")
    wordArray[2] = strings.TrimSuffix(golang, ",")
    num1, err = strconv.ParseFloat(wordArray[1], 64)
    if err != nil {return badmsg}
    num2, err = strconv.ParseFloat(wordArray[2], 64)
    if err != nil {return badmsg}
    num3, err = strconv.ParseFloat(wordArray[3], 64)
    if err != nil {return badmsg}
    return msg

  case "FUNC:SQU:DCYC": // "FUNC:SQU:DCYC %f\n"
    wordArray[1] = wordArray[1]
    return msg // Sets the percentage of time, specified as 0-100, during one
                  // cycle for which the square wave is at its high value

  case "FREQ": // "FREQ %f\n"
    wordArray[1] = wordArray[1]
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
