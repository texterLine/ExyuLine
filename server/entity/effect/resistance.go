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

// Resistance is a lasting effect that reduces the damage taken from any
// sources except for void damage or custom damage.
var Resistance resistance

type resistance struct {
	nopLasting
}

// Multiplier returns a damage multiplier for the damage source passed.
func (resistance) Multiplier(e world.DamageSource, lvl int) float64 {
	if !e.ReducedByResistance() {
		return 1
	}
	if v := 1 - 0.2*float64(lvl); v >= 0 {
		return v
	}
	return 0
}

// RGBA ...
func (resistance) RGBA() color.RGBA {
	return color.RGBA{R: 0x91, G: 0x46, B: 0xf0, A: 0xff}
}
