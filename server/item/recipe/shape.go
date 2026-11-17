// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package recipe

// Shape make up the shape of a shaped recipe. It consists of a width and a height.
type Shape [2]int

// Width returns the width of the shape.
func (s Shape) Width() int {
	return s[0]
}

// Height returns the height of the shape.
func (s Shape) Height() int {
	return s[1]
}

// NewShape creates a new shape using the provided width and height.
func NewShape(width, height int) Shape {
	return Shape{width, height}
}
