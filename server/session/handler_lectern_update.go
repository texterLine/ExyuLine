// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/texterLine/exyuline/server/block"
	"github.com/texterLine/exyuline/server/world"
)

// LecternUpdateHandler handles the LecternUpdate packet, sent when a player interacts with a lectern.
type LecternUpdateHandler struct{}

// Handle ...
func (LecternUpdateHandler) Handle(p packet.Packet, _ *Session, tx *world.Tx, c Controllable) error {
	pk := p.(*packet.LecternUpdate)
	pos := blockPosFromProtocol(pk.Position)
	if !canReach(c, pos.Vec3Middle()) {
		return fmt.Errorf("block at %v is not within reach", pos)
	}
	if _, ok := tx.Block(pos).(block.Lectern); !ok {
		return fmt.Errorf("block at %v is not a lectern", pos)
	}
	return c.TurnLecternPage(pos, int(pk.Page))
}
