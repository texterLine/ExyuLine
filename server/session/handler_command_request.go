// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/texterLine/exyuline/server/world"
)

// CommandRequestHandler handles the CommandRequest packet.
type CommandRequestHandler struct {
	origin protocol.CommandOrigin
}

// Handle ...
func (h *CommandRequestHandler) Handle(p packet.Packet, _ *Session, _ *world.Tx, c Controllable) error {
	pk := p.(*packet.CommandRequest)
	if pk.Internal {
		return fmt.Errorf("command request packet must never have the internal field set to true")
	}

	h.origin = pk.CommandOrigin
	c.ExecuteCommand(pk.CommandLine)
	return nil
}
