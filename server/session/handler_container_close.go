// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/texterLine/exyuline/server/world"
)

// ContainerCloseHandler handles the ContainerClose packet.
type ContainerCloseHandler struct{}

// Handle ...
func (h *ContainerCloseHandler) Handle(p packet.Packet, s *Session, tx *world.Tx, c Controllable) error {
	pk := p.(*packet.ContainerClose)

	c.MoveItemsToInventory()

	var containerType byte
	switch pk.WindowID {
	case 0:
		// Closing of the normal inventory.
		s.invOpened = false
	case byte(s.openedWindowID.Load()):
		containerType = byte(s.openedContainerID.Load())
		s.closeCurrentContainer(tx, true)
	case 0xff:
		// Sent when an inventory/container is opened at the same time as chat.
		s.invOpened = false
		if s.containerOpened.Load() {
			s.closeCurrentContainer(tx, false)
		}
		return nil
	default:
		containerType = pk.ContainerType
	}
	s.writePacket(&packet.ContainerClose{
		WindowID:      pk.WindowID,
		ContainerType: containerType,
	})
	return nil
}
