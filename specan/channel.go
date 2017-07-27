// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package specan

import "github.com/gotmc/ivi"

// Channel models a generic FGen channel
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

//
func NewChannel(id int, name string, inst ivi.Instrument) Channel {
	return Channel{id, name, inst}
}

//
