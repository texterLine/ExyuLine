// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/item/enchantment"
	"github.com/texterLine/exyuline/server/world"
)

type (
	// AttackDamageSource is used for damage caused by other entities, for
	// example when a player attacks another player.
	AttackDamageSource struct {
		// Attacker holds the attacking entity. The entity may be a player or
		// any other entity.
		Attacker world.Entity
	}

	// VoidDamageSource is used for damage caused by an entity being in the
	// void.
	VoidDamageSource struct{}

	// SuffocationDamageSource is used for damage caused by an entity
	// suffocating in a block.
	SuffocationDamageSource struct{}

	// DrowningDamageSource is used for damage caused by an entity drowning in
	// water.
	DrowningDamageSource struct{}

	// FallDamageSource is used for damage caused by falling.
	FallDamageSource struct{}

	// GlideDamageSource is used for damage caused by gliding into a block.
	GlideDamageSource struct{}

	// LightningDamageSource is used for damage caused by being struck by
	// lightning.
	LightningDamageSource struct{}

	// ProjectileDamageSource is used for damage caused by a projectile.
	ProjectileDamageSource struct {
		// Projectile and Owner are the world.Entity that dealt the damage and
		// the one that fired the projectile respectively.
		Projectile, Owner world.Entity
	}

	// ExplosionDamageSource is used for damage caused by an explosion.
	ExplosionDamageSource struct {
		// Source is the source of the explosion that dealt the damage.
		Source world.ExplosionSource
	}
)

func (FallDamageSource) ReducedByArmour() bool     { return false }
func (FallDamageSource) ReducedByResistance() bool { return true }
func (FallDamageSource) Fire() bool                { return false }
func (FallDamageSource) AffectedByEnchantment(e item.EnchantmentType) bool {
	return e == enchantment.FeatherFalling
}
func (FallDamageSource) IgnoreTotem() bool                { return false }
func (GlideDamageSource) ReducedByArmour() bool           { return false }
func (GlideDamageSource) ReducedByResistance() bool       { return true }
func (GlideDamageSource) Fire() bool                      { return false }
func (GlideDamageSource) IgnoreTotem() bool               { return false }
func (LightningDamageSource) ReducedByArmour() bool       { return true }
func (LightningDamageSource) ReducedByResistance() bool   { return true }
func (LightningDamageSource) Fire() bool                  { return false }
func (LightningDamageSource) IgnoreTotem() bool           { return false }
func (AttackDamageSource) ReducedByArmour() bool          { return true }
func (AttackDamageSource) ReducedByResistance() bool      { return true }
func (AttackDamageSource) Fire() bool                     { return false }
func (AttackDamageSource) IgnoreTotem() bool              { return false }
func (VoidDamageSource) ReducedByResistance() bool        { return false }
func (VoidDamageSource) ReducedByArmour() bool            { return false }
func (VoidDamageSource) Fire() bool                       { return false }
func (VoidDamageSource) IgnoreTotem() bool                { return true }
func (SuffocationDamageSource) ReducedByResistance() bool { return false }
func (SuffocationDamageSource) ReducedByArmour() bool     { return false }
func (SuffocationDamageSource) Fire() bool                { return false }
func (SuffocationDamageSource) IgnoreTotem() bool         { return false }
func (DrowningDamageSource) ReducedByResistance() bool    { return false }
func (DrowningDamageSource) ReducedByArmour() bool        { return false }
func (DrowningDamageSource) Fire() bool                   { return false }
func (DrowningDamageSource) IgnoreTotem() bool            { return false }
func (ProjectileDamageSource) ReducedByResistance() bool  { return true }
func (ProjectileDamageSource) ReducedByArmour() bool      { return true }
func (ProjectileDamageSource) Fire() bool                 { return false }
func (ProjectileDamageSource) AffectedByEnchantment(e item.EnchantmentType) bool {
	return e == enchantment.ProjectileProtection
}
func (ProjectileDamageSource) IgnoreTotem() bool        { return false }
func (ExplosionDamageSource) ReducedByResistance() bool { return true }
func (ExplosionDamageSource) ReducedByArmour() bool     { return true }
func (ExplosionDamageSource) Fire() bool                { return false }
func (ExplosionDamageSource) AffectedByEnchantment(e item.EnchantmentType) bool {
	return e == enchantment.BlastProtection
}
func (ExplosionDamageSource) IgnoreTotem() bool { return false }
