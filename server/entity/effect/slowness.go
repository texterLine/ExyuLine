// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package effect

import (
	"github.com/texterLine/exyuline/server/world"
	"image/color"
)

// Slowness is a lasting effect that decreases the movement speed of a living
// entity by 15% for each level that the effect has.
var Slowness slowness

type slowness struct {
	nopLasting
}

// Start ...
func (slowness) Start(e world.Entity, lvl int) {
	slowness := 1 - float64(lvl)*0.15
	if slowness <= 0 {
		slowness = 0.00001
	}
	if l, ok := e.(living); ok {
		l.SetSpeed(l.Speed() * slowness)
	}
}

// End ...
func (slowness) End(e world.Entity, lvl int) {
	slowness := 1 - float64(lvl)*0.15
	if slowness <= 0 {
		slowness = 0.00001
	}
	if l, ok := e.(living); ok {
		l.SetSpeed(l.Speed() / slowness)
	}
}

// RGBA ...
func (slowness) RGBA() color.RGBA {
	return color.RGBA{R: 0x8b, G: 0xaf, B: 0xe0, A: 0xff}
}
