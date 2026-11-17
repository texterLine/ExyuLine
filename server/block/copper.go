// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
	"math/rand/v2"
)

// Copper is a solid block commonly found in deserts and beaches underneath sand.
type Copper struct {
	solid
	bassDrum

	// Type is the type of copper of the block.
	Type CopperType
	// Oxidation is the level of oxidation of the copper block.
	Oxidation OxidationType
	// Waxed bool is whether the copper block has been waxed with honeycomb.
	Waxed bool
}

func (c Copper) Strip() (world.Block, world.Sound, bool) {
	if c.Waxed {
		c.Waxed = false
		return c, sound.WaxRemoved{}, true
	} else if ot, ok := c.Oxidation.Decrease(); ok {
		c.Oxidation = ot
		return c, sound.CopperScraped{}, true
	}
	return c, nil, false
}

// BreakInfo ...
func (c Copper) BreakInfo() BreakInfo {
	return newBreakInfo(3, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// Wax waxes the copper block to stop it from oxidising further.
func (c Copper) Wax(cube.Pos, mgl64.Vec3) (world.Block, bool) {
	before := c.Waxed
	c.Waxed = true
	return c, !before
}

func (c Copper) CanOxidate() bool {
	return !c.Waxed
}

func (c Copper) OxidationLevel() OxidationType {
	return c.Oxidation
}

func (c Copper) WithOxidationLevel(o OxidationType) Oxidisable {
	c.Oxidation = o
	return c
}

func (c Copper) RandomTick(pos cube.Pos, tx *world.Tx, r *rand.Rand) {
	attemptOxidation(pos, tx, r, c)
}

// EncodeItem ...
func (c Copper) EncodeItem() (name string, meta int16) {
	if c.Type == NormalCopper() && c.Oxidation == UnoxidisedOxidation() && !c.Waxed {
		return "minecraft:copper_block", 0
	}
	name = "copper"
	if c.Type != NormalCopper() {
		name = c.Type.String() + "_" + name
	}
	if c.Oxidation != UnoxidisedOxidation() {
		name = c.Oxidation.String() + "_" + name
	}
	if c.Waxed {
		name = "waxed_" + name
	}
	return "minecraft:" + name, 0
}

// EncodeBlock ...
func (c Copper) EncodeBlock() (string, map[string]any) {
	if c.Type == NormalCopper() && c.Oxidation == UnoxidisedOxidation() && !c.Waxed {
		return "minecraft:copper_block", nil
	}
	name := "copper"
	if c.Type != NormalCopper() {
		name = c.Type.String() + "_" + name
	}
	if c.Oxidation != UnoxidisedOxidation() {
		name = c.Oxidation.String() + "_" + name
	}
	if c.Waxed {
		name = "waxed_" + name
	}
	return "minecraft:" + name, nil
}

// allCopper returns a list of all copper block variants.
func allCopper() (c []world.Block) {
	f := func(waxed bool) {
		for _, t := range CopperTypes() {
			for _, o := range OxidationTypes() {
				c = append(c, Copper{Type: t, Oxidation: o, Waxed: waxed})
			}
		}
	}
	f(true)
	f(false)
	return
}
