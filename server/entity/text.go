// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// NewText creates and returns a new Text entity with the text and position provided.
func NewText(text string, pos mgl64.Vec3) *world.EntityHandle {
	return world.EntitySpawnOpts{Position: pos, NameTag: text}.New(TextType, textConf)
}

var textConf StationaryBehaviourConfig

// TextType is a world.EntityType implementation for Text.
var TextType textType

type textType struct{}

func (t textType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	return &Ent{tx: tx, handle: handle, data: data}
}
func (textType) EncodeEntity() string        { return "exyuline:text" }
func (textType) BBox(world.Entity) cube.BBox { return cube.BBox{} }
func (textType) NetworkEncodeEntity() string { return "minecraft:falling_block" }

func (textType) DecodeNBT(_ map[string]any, data *world.EntityData) { data.Data = textConf.New() }
func (textType) EncodeNBT(_ *world.EntityData) map[string]any       { return nil }
