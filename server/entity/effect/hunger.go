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

// Hunger is a lasting effect that causes an affected player to gradually lose
// saturation and food.
var Hunger hunger

type hunger struct {
	nopLasting
}

// Apply ...
func (hunger) Apply(e world.Entity, eff Effect) {
	if i, ok := e.(interface {
		Exhaust(points float64)
	}); ok {
		i.Exhaust(float64(eff.Level()) * 0.005)
	}
}

// RGBA ...
func (hunger) RGBA() color.RGBA {
	return color.RGBA{R: 0x58, G: 0x76, B: 0x53, A: 0xff}
}
