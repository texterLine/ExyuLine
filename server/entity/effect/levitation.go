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

// Levitation is a lasting effect that causes the affected entity to slowly
// levitate upwards. It is roughly the opposite of the slowFalling effect.
var Levitation levitation

type levitation struct {
	nopLasting
}

// RGBA ...
func (levitation) RGBA() color.RGBA {
	return color.RGBA{R: 0xce, G: 0xff, B: 0xff, A: 0xff}
}
