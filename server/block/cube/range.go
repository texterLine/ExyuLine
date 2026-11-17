// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package cube

// Range represents the height range of a Dimension in blocks. The first value
// of the Range holds the minimum Y value, the second value holds the maximum Y
// value.
type Range [2]int

// Min returns the minimum Y value of a Range. It is equivalent to Range[0].
func (r Range) Min() int {
	return r[0]
}

// Max returns the maximum Y value of a Range. It is equivalent to Range[1].
func (r Range) Max() int {
	return r[1]
}

// Height returns the total height of the Range, the difference between Max and
// Min.
func (r Range) Height() int {
	return r[1] - r[0]
}
