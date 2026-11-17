// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// EmoteHandler handles the Emote packet.
type EmoteHandler struct {
	LastEmote time.Time
}

// Handle ...
func (h *EmoteHandler) Handle(p packet.Packet, _ *Session, tx *world.Tx, c Controllable) error {
	pk := p.(*packet.Emote)

	if pk.EntityRuntimeID != selfEntityRuntimeID {
		return errSelfRuntimeID
	}
	if time.Since(h.LastEmote) < time.Second {
		return nil
	}
	h.LastEmote = time.Now()
	emote, err := uuid.Parse(pk.EmoteID)
	if err != nil {
		return err
	}
	for _, viewer := range tx.Viewers(c.Position()) {
		viewer.ViewEmote(c, emote)
	}
	return nil
}
