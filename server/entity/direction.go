// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/world"
)

// Eyed represents an entity that has eyes.
type Eyed interface {
	// EyeHeight returns the offset from their base position that the eyes of an entity are found at.
	EyeHeight() float64
}

// EyePosition returns the position of the eyes of the entity if the entity implements entity.Eyed, or the
// actual position if it doesn't.
func EyePosition(e world.Entity) mgl64.Vec3 {
	pos := e.Position()
	if eyed, ok := e.(Eyed); ok {
		pos[1] += eyed.EyeHeight()
	}
	return pos
}
