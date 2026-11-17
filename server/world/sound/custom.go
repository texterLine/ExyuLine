// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package sound

// Custom is a sound identified by name. It may be used to play sounds defined
// by a resource pack.
type Custom struct {
	// Name is the identifier of the sound.
	Name string
	// Volume is the volume of the sound.
	Volume float64
	// Pitch is the pitch of the sound.
	Pitch float64

	sound
}
