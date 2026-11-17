// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package mcrandom

// MixStafford13 implements the Stafford 13 mixing function, a bijective mixing function
// suitable for use in pseudorandom number generators.
func MixStafford13(seed uint64) uint64 {
	seed = (seed ^ (seed >> 30)) * 0xBF58476D1CE4E5B9
	seed = (seed ^ (seed >> 27)) * 0x94D049BB133111EB
	return seed ^ (seed >> 31)
}
