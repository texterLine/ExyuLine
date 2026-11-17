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

// Weakness is a lasting effect that reduces the damage dealt to other entities
// with melee attacks.
var Weakness weakness

type weakness struct {
	nopLasting
}

// Multiplier returns the damage multiplier of the effect.
func (weakness) Multiplier(lvl int) float64 {
	v := 0.2 * float64(lvl)
	if v > 1 {
		v = 1
	}
	return v
}

// RGBA ...
func (weakness) RGBA() color.RGBA {
	return color.RGBA{R: 0x48, G: 0x4d, B: 0x48, A: 0xff}
}
