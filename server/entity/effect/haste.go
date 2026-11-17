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

// Haste is a lasting effect that increases the mining speed of a player by 20%
// for each level of the effect.
var Haste haste

type haste struct {
	nopLasting
}

// Multiplier returns the mining speed multiplier from this effect.
func (haste) Multiplier(lvl int) float64 {
	v := 1 - float64(lvl)*0.1
	if v < 0 {
		v = 0
	}
	return v
}

// RGBA ...
func (haste) RGBA() color.RGBA {
	return color.RGBA{R: 0xd9, G: 0xc0, B: 0x43, A: 0xff}
}
