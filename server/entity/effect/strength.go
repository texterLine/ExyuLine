// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package effect

import (
	"image/color"
)

// Strength is a lasting effect that increases the damage dealt with melee
// attacks when applied to an entity.
var Strength strength

type strength struct {
	nopLasting
}

// Multiplier returns the damage multiplier of the effect.
func (strength) Multiplier(lvl int) float64 {
	return 0.3 * float64(lvl)
}

// RGBA ...
func (strength) RGBA() color.RGBA {
	return color.RGBA{R: 0xff, G: 0xc7, B: 0x00, A: 0xff}
}
