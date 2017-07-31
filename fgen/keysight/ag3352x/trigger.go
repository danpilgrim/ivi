// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ag3352x

import (
	"errors"
	"strings"

	"github.com/gotmc/ivi/fgen"
)

// TriggerSource determines the trigger srouce. TriggerSource is the getter for
// the read-write IviFgenTrigger Attribute Trigger Source described in Section
// 9.2.1 of IVI-4.3: IviFgen Class Specification.
// AG3352x - no change (Ref: pg. 432)
func (ch *Channel) TriggerSource() (fgen.TriggerSource, error) {
	var src fgen.TriggerSource
	s, err := ch.QueryString("TRIG:SOUR?\n")
	if err != nil {
		return src, err
	}
	s = strings.TrimSpace(strings.ToUpper(s))
	switch s {
	case "IMM":
		src = fgen.InternalTrigger
	case "EXT":
		src = fgen.ExternalTrigger
	case "BUS":
		src = fgen.SoftwareTrigger
	default:
		return src, errors.New("error determining trigger source")
	}
	return src, nil
}

// SetTriggerSource specifies the trigger srouce. SetTriggerSource is the
// setter for the read-write IviFgenTrigger Attribute Trigger Source described
// in Section 9.2.1 of IVI-4.3: IviFgen Class Specification.
// AG3352x - no change (Ref: pg. 432)
func (ch *Channel) SetTriggerSource(src fgen.TriggerSource) error {
	triggers := map[fgen.TriggerSource]string{
		fgen.InternalTrigger: "IMM",
		fgen.ExternalTrigger: "EXT",
		fgen.SoftwareTrigger: "BUS",
	}
	return ch.Set("TRIGE:SOUR %s\n", triggers[src])
}
