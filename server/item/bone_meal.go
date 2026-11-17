// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/particle"
)

// BoneMealResult represents the outcome of a bone meal interaction with a block,
// determining the intensity of the particle effect displayed.
type BoneMealResult int

const (
	// BoneMealResultNone indicates that the bone meal had no effect on the block.
	BoneMealResultNone BoneMealResult = iota
	// BoneMealResultSmall indicates a minor growth effect, produces a small particle burst.
	BoneMealResultSmall
	// BoneMealResultArea indicates a significant growth effect over an area, produces a large particle burst.
	BoneMealResultArea
)

// BoneMeal is an item used to force growth in plants & crops.
type BoneMeal struct{}

// BoneMealAffected represents a block that is affected when bone meal is used on it.
type BoneMealAffected interface {
	// BoneMeal attempts to affect the block using a bone meal item.
	BoneMeal(pos cube.Pos, tx *world.Tx) BoneMealResult
}

// UseOnBlock ...
func (b BoneMeal) UseOnBlock(pos cube.Pos, _ cube.Face, _ mgl64.Vec3, tx *world.Tx, _ User, ctx *UseContext) bool {
	if bm, ok := tx.Block(pos).(BoneMealAffected); ok {
		result := bm.BoneMeal(pos, tx)
		if result == BoneMealResultNone {
			return false
		}

		ctx.SubtractFromCount(1)
		tx.AddParticle(pos.Vec3(), particle.BoneMeal{
			Area: result == BoneMealResultArea,
		})
		return true
	}
	return false
}

// EncodeItem ...
func (b BoneMeal) EncodeItem() (name string, meta int16) {
	return "minecraft:bone_meal", 0
}
