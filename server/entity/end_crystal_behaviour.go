// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"math/rand/v2"
	"time"

	"github.com/texterLine/exyuline/server/block"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

type endCrystalBehaviour struct {
	showBase      bool
	beamTarget    cube.Pos
	hasBeamTarget bool
	explosionSize float64
}

func (b endCrystalBehaviour) Apply(data *world.EntityData) {
	data.Data = b
}

// Tick continuously generates fire at the End crystal's position while in the
// End, if the block at that position is air and the block below it is not air.
func (endCrystalBehaviour) Tick(e *Ent, tx *world.Tx) *Movement {
	if tx.World().Dimension() == world.End {
		pos := cube.PosFromVec3(e.Position())
		if _, air := tx.Block(pos.Side(cube.FaceDown)).(block.Air); !air {
			if _, air := tx.Block(pos).(block.Air); air {
				fire := block.Fire{}
				tx.SetBlock(pos, fire, nil)
				tx.ScheduleBlockUpdate(pos, fire, time.Duration(30+rand.IntN(10))*time.Second/20)
			}
		}
	}
	return nil
}

// Explode makes the End crystal explode itself when hit by another explosion,
// causing a chain reaction.
func (b endCrystalBehaviour) Explode(e *Ent, _ world.ExplosionSource, impact float64) {
	if impact <= 0 {
		return
	}
	explodeEndCrystal(e, b.explosionSize)
}

// Hurt makes the End crystal explode when damaged by any source, even by
// damage that deals no health. Void damage removes it without an explosion.
func (b endCrystalBehaviour) Hurt(e *Ent, damage float64, src world.DamageSource) (float64, bool) {
	damage = max(damage, 0)
	if _, ok := src.(VoidDamageSource); ok {
		_ = e.Close()
		return damage, true
	}
	explodeEndCrystal(e, b.explosionSize)
	return damage, true
}

func (endCrystalBehaviour) Immobile() bool {
	return true
}

func (b endCrystalBehaviour) ShowBase() bool {
	return b.showBase
}

func (b endCrystalBehaviour) BeamTarget() (cube.Pos, bool) {
	return b.beamTarget, b.hasBeamTarget
}

// explodeEndCrystal closes the End crystal and creates a non-incendiary
// explosion at its base, if the crystal was not closed yet.
func explodeEndCrystal(e *Ent, explosionSize float64) {
	if _, ok := e.H().Entity(e.tx); !ok {
		return
	}
	_ = e.Close()
	block.ExplosionConfig{
		SuppressUnderwaterImpact: true,
	}.Explode(e.tx, world.EntityExplosionSource{
		Entity:        e,
		ExplosionSize: explosionSize,
	})
}
