// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"math/rand/v2"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
)

// CopperGrate is a solid block commonly found in deserts and beaches underneath sand.
type CopperGrate struct {
	sourceWaterDisplacer
	solid
	transparent
	bassDrum

	// Oxidation is the level of oxidation of the copper grate.
	Oxidation OxidationType
	// Waxed bool is whether the copper grate has been waxed with honeycomb.
	Waxed bool
}

// BreakInfo ...
func (c CopperGrate) BreakInfo() BreakInfo {
	return newBreakInfo(3, func(t item.Tool) bool {
		return t.ToolType() == item.TypePickaxe && t.HarvestLevel() >= item.ToolTierStone.HarvestLevel
	}, pickaxeEffective, oneOf(c)).withBlastResistance(6)
}

// Wax waxes the copper grate to stop it from oxidising further.
func (c CopperGrate) Wax(cube.Pos, mgl64.Vec3) (world.Block, bool) {
	if c.Waxed {
		return c, false
	}
	c.Waxed = true
	return c, true
}

func (c CopperGrate) Strip() (world.Block, world.Sound, bool) {
	if c.Waxed {
		c.Waxed = false
		return c, sound.WaxRemoved{}, true
	} else if ot, ok := c.Oxidation.Decrease(); ok {
		c.Oxidation = ot
		return c, sound.CopperScraped{}, true
	}
	return c, nil, false
}

func (c CopperGrate) CanOxidate() bool {
	return !c.Waxed
}

func (c CopperGrate) OxidationLevel() OxidationType {
	return c.Oxidation
}

func (c CopperGrate) WithOxidationLevel(o OxidationType) Oxidisable {
	c.Oxidation = o
	return c
}

func (c CopperGrate) RandomTick(pos cube.Pos, tx *world.Tx, r *rand.Rand) {
	attemptOxidation(pos, tx, r, c)
}

// EncodeItem ...
func (c CopperGrate) EncodeItem() (name string, meta int16) {
	return copperBlockName("copper_grate", c.Oxidation, c.Waxed), 0
}

// EncodeBlock ...
func (c CopperGrate) EncodeBlock() (string, map[string]any) {
	return copperBlockName("copper_grate", c.Oxidation, c.Waxed), nil
}

// allCopperGrates returns a list of all copper grate variants.
func allCopperGrates() (c []world.Block) {
	f := func(waxed bool) {
		for _, o := range OxidationTypes() {
			c = append(c, CopperGrate{Oxidation: o, Waxed: waxed})
		}
	}
	f(true)
	f(false)
	return
}
