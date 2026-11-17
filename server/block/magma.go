// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/item/enchantment"
	"github.com/texterLine/exyuline/server/world"
)

// Magma is a light-emitting Nether block that damages entities standing on it.
type Magma struct {
	solid
	bassDrum
}

// LightEmissionLevel ...
func (Magma) LightEmissionLevel() uint8 {
	return 3
}

// EntityStepOn ...
func (Magma) EntityStepOn(_ cube.Pos, _ *world.Tx, e world.Entity) {
	if fireProof, ok := e.(interface{ FireProof() bool }); ok && fireProof.FireProof() {
		return
	}
	// TODO: Check for Frost Walker once the enchantment is implemented in ExyuLine.
	if sneaking, ok := e.(interface{ Sneaking() bool }); ok && sneaking.Sneaking() {
		return
	}
	if l, ok := e.(livingEntity); ok {
		l.Hurt(1, MagmaDamageSource{})
	}
}

// BreakInfo ...
func (m Magma) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, pickaxeHarvestable, pickaxeEffective, oneOf(m))
}

// EncodeItem ...
func (Magma) EncodeItem() (name string, meta int16) {
	return "minecraft:magma", 0
}

// EncodeBlock ...
func (Magma) EncodeBlock() (string, map[string]any) {
	return "minecraft:magma", nil
}

// MagmaDamageSource is used for damage caused by standing on a magma block.
type MagmaDamageSource struct{}

func (MagmaDamageSource) ReducedByResistance() bool { return true }
func (MagmaDamageSource) ReducedByArmour() bool     { return true }
func (MagmaDamageSource) Fire() bool                { return true }
func (MagmaDamageSource) AffectedByEnchantment(e item.EnchantmentType) bool {
	return e == enchantment.FireProtection
}
func (MagmaDamageSource) IgnoreTotem() bool { return false }
