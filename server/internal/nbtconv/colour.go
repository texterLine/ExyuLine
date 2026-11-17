// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package nbtconv

import (
	"encoding/binary"
	"image/color"
)

// Int32FromRGBA converts a color.RGBA into an int32. These int32s are present in, for example, signs.
func Int32FromRGBA(x color.RGBA) int32 {
	if x.R == 0 && x.G == 0 && x.B == 0 {
		// Default to black colour. The default (0x000000) is a transparent colour. Text with this colour will not show
		// up on the sign.
		return int32(-0x1000000)
	}
	return int32(binary.BigEndian.Uint32([]byte{x.A, x.R, x.G, x.B}))
}

// RGBAFromInt32 converts an int32 into a color.RGBA. These int32s are present in, for example, signs.
func RGBAFromInt32(x int32) color.RGBA {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(x))

	return color.RGBA{A: b[0], R: b[1], G: b[2], B: b[3]}
}
