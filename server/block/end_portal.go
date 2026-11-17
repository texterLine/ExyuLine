// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/block/model"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/portal"
)

// EndPortal is the translucent block that teleports the player to and from the End. It is created by inserting an Eye
// of Ender into all twelve frame blocks of an End portal frame ring.
type EndPortal struct {
	transparent
}

// Model ...
func (EndPortal) Model() world.BlockModel {
	return model.Empty{}
}

// LightEmissionLevel returns 15.
func (EndPortal) LightEmissionLevel() uint8 {
	return 15
}

// HasLiquidDrops ...
func (EndPortal) HasLiquidDrops() bool {
	return false
}

// Portal returns the End dimension. The same block leads back to the Overworld when entered from the End.
func (EndPortal) Portal() world.Dimension {
	return world.End
}

// EncodeNBT encodes the End portal block actor.
func (EndPortal) EncodeNBT() map[string]any {
	return map[string]any{"id": "EndPortal"}
}

// DecodeNBT decodes the End portal block actor.
func (e EndPortal) DecodeNBT(map[string]any) any {
	return e
}

// NeighbourUpdateTick removes the connected portal blocks if the surrounding frame ring is no longer complete,
// like breaking the frame of a nether portal.
func (EndPortal) NeighbourUpdateTick(pos, _ cube.Pos, tx *world.Tx) {
	if portal.EndPortalRingIntact(tx, pos) {
		return
	}
	portal.DeactivateEndPortal(tx, pos)
}

// EntityInside ...
func (EndPortal) EntityInside(_ cube.Pos, tx *world.Tx, e world.Entity) {
	if t, ok := e.(portalTraveller); ok {
		t.TravelThroughPortal(tx, world.End)
	}
}

// EncodeBlock ...
func (EndPortal) EncodeBlock() (string, map[string]any) {
	return "minecraft:end_portal", nil
}
