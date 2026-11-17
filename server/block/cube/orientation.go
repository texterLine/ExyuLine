// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package cube

import "math"

// Orientation represents the rotation of a block divided into 16 options.
type Orientation int

// OrientationFromYaw returns an Orientation value that (roughly) matches the
// yaw passed.
func OrientationFromYaw(yaw float64) Orientation {
	yaw = math.Mod(yaw, 360)
	return Orientation(math.Round(yaw / 360 * 16))
}

// Yaw returns the yaw value that matches the orientation.
func (o Orientation) Yaw() float64 {
	return float64(o) / 16 * 360
}

// Opposite returns the opposite orientation value of the Orientation.
func (o Orientation) Opposite() Orientation {
	return OrientationFromYaw(o.Yaw() + 180)
}

// RotateLeft rotates the orientation left by 90 degrees and returns it.
func (o Orientation) RotateLeft() Orientation {
	return OrientationFromYaw(o.Yaw() - 90)
}

// RotateRight rotates the orientation right by 90 degrees and returns it.
func (o Orientation) RotateRight() Orientation {
	return OrientationFromYaw(o.Yaw() + 90)
}
