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
	"github.com/texterLine/exyuline/server/world"
	"sync/atomic"
)

// ServerBoundLoadingScreenHandler handles loading screen updates from the clients. It is used to ensure that
// the server knows when the client is loading a screen, and when it is done loading it.
type ServerBoundLoadingScreenHandler struct {
	currentID  atomic.Uint32
	expectedID atomic.Uint32
}

// Handle ...
func (h *ServerBoundLoadingScreenHandler) Handle(p packet.Packet, s *Session, _ *world.Tx, _ Controllable) error {
	pk := p.(*packet.ServerBoundLoadingScreen)
	v, ok := pk.LoadingScreenID.Value()
	expected := h.expectedID.Load()

	switch {
	case !ok || expected == 0:
		return nil
	case v != expected:
		return fmt.Errorf("expected loading screen ID %d, got %d", expected, v)
	case pk.Type == packet.LoadingScreenTypeEnd:
		s.changingDimension.Store(false)
		h.expectedID.Store(0)
	}

	return nil
}
