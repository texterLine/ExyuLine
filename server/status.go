// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package server

import (
	"github.com/sandertv/gophertunnel/minecraft"
)

// statusProvider handles the way the server shows up in the server list. The
// online players and maximum players are not changeable from outside the
// server, but the server name may be changed at any time.
type statusProvider struct {
	name string
}

// ServerStatus returns the player count, max players and the server's name as
// a minecraft.ServerStatus.
func (s statusProvider) ServerStatus(playerCount, maxPlayers int) minecraft.ServerStatus {
	return minecraft.ServerStatus{
		ServerName:  s.name,
		PlayerCount: playerCount,
		MaxPlayers:  maxPlayers,
	}
}
