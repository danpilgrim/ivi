// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project. //
// Changes for Ag3352x reference Keysite's 33500-90901 Manual //
/*
Package Ag3352x implements the IVI driver for the Agilent Ag3352x function
generator.
State Caching: Not implemented
*/

package ag3352x

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 3
	classSpecRevision     = "5.2"
	groupCapabilities     = "IviFgenBase,IviFgenStdfunc,IviFgenTrigger,IviFgenInternalTrigger,IviFgenBurst"
)

// TODO(mdr): Seems like groupCapabilities should be a []string instead of
// string

// 33512B model number
var supportedInstrumentModels = []string{
	"33512B",
}

var channelNames = []string{
	"Output",
}

// Ag3352x provides the IVI driver for an Agilent Ag3352x
// function generator.
type Ag3352x struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new Ag3352x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Ag3352x, error) {

	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)

	for i, ch := range channelNames {
		baseChannel := fgen.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}

	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}

	inherent := ivi.NewInherent(inst, inherentBase)

	fgen := Ag3352x{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}

	if reset {
		err := fgen.Reset()
		return &fgen, err
	}
	return &fgen, nil
}

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	fgen.Channel
}
