// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Particle represents a particle that may be added to the world. These particles are then rendered client-
// side, with the server having no control over it after sending.
type Particle interface {
	// Spawn spawns the particle at the position passed. Particles may execute any additional actions here,
	// such as spawning different particles.
	Spawn(w *World, pos mgl64.Vec3)
}
