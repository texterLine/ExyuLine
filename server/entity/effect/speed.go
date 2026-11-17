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

// Speed is a lasting effect that increases the speed of an entity by 20% for
// each level that the effect has.
var Speed speed

type speed struct {
	nopLasting
}

// Start ...
func (speed) Start(e world.Entity, lvl int) {
	speed := 1 + float64(lvl)*0.2
	if l, ok := e.(living); ok {
		l.SetSpeed(l.Speed() * speed)
	}
}

// End ...
func (speed) End(e world.Entity, lvl int) {
	speed := 1 + float64(lvl)*0.2
	if l, ok := e.(living); ok {
		l.SetSpeed(l.Speed() / speed)
	}
}

// RGBA ...
func (speed) RGBA() color.RGBA {
	return color.RGBA{R: 0x33, G: 0xeb, B: 0xff, A: 0xff}
}
