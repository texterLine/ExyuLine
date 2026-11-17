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

// RequestAbilityHandler handles the RequestAbility packet.
type RequestAbilityHandler struct{}

// Handle ...
func (a RequestAbilityHandler) Handle(p packet.Packet, s *Session, _ *world.Tx, c Controllable) error {
	pk := p.(*packet.RequestAbility)
	if pk.Ability == packet.AbilityFlying {
		if !c.GameMode().AllowsFlying() {
			s.conf.Log.Debug("process packet: RequestAbility: flying flag enabled while unable to fly")
			s.SendAbilities(c)
			return nil
		}
		c.StartFlying()
	}
	return nil
}
