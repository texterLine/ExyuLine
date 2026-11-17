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

// WaterBreathing is a lasting effect that allows the affected entity to breath
// underwater until the effect expires.
var WaterBreathing waterBreathing

type waterBreathing struct {
	nopLasting
}

// RGBA ...
func (waterBreathing) RGBA() color.RGBA {
	return color.RGBA{R: 0x98, G: 0xda, B: 0xc0, A: 0xff}
}
