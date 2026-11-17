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
	"github.com/texterLine/exyuline/server/player/dialogue"
	"github.com/texterLine/exyuline/server/world"
)

// NPCRequestHandler handles the NPCRequest packet.
type NPCRequestHandler struct {
	dialogue        dialogue.Dialogue
	entityRuntimeID uint64
}

// Handle ...
func (h *NPCRequestHandler) Handle(p packet.Packet, s *Session, tx *world.Tx, c Controllable) error {
	pk := p.(*packet.NPCRequest)
	if h.entityRuntimeID == 0 {
		// No dialogue is currently open for this session, so there is nothing to submit or close.
		return nil
	}
	switch pk.RequestType {
	case packet.NPCRequestActionExecuteAction:
		if err := h.dialogue.Submit(uint(pk.ActionType), c, tx); err != nil {
			return fmt.Errorf("error submitting dialogue: %w", err)
		}
	case packet.NPCRequestActionExecuteClosingCommands:
		h.dialogue.Close(c, tx)
	}
	return nil
}
