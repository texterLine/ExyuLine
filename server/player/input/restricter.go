// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package input

// Restricter represents an interface that can manage input locks for a player.
type Restricter interface {
	// LockInput applies an input lock to the player, disabling the specified input.
	LockInput(l Lock)
	// UnlockInput removes an input lock from the player, re-enabling the specified input.
	UnlockInput(l Lock)
	// ClearInputLocks removes all input locks from the player, re-enabling all inputs.
	ClearInputLocks()
	// InputLocked checks if a specific input lock is currently applied to the player.
	InputLocked(l Lock) bool
}
