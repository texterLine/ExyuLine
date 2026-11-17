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

// Saturation is a lasting effect that causes the affected player to regain
// food and saturation.
var Saturation saturation

type saturation struct {
	nopLasting
}

// Apply ...
func (saturation) Apply(e world.Entity, eff Effect) {
	if i, ok := e.(interface {
		Saturate(food int, saturation float64)
	}); ok {
		i.Saturate(eff.Level(), 2*float64(eff.Level()))
	}
}

// RGBA ...
func (saturation) RGBA() color.RGBA {
	return color.RGBA{R: 0xf8, G: 0x24, B: 0x23, A: 0xff}
}
