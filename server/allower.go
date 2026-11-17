// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package server

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"net"
)

// Allower may be implemented to specifically allow or disallow players from
// joining a Server, by setting the specific Allower implementation through a
// call to Server.Allow.
type Allower interface {
	// Allow filters what connections are allowed to connect to the Server. The
	// address, identity data, and client data of the connection are passed. If
	// Admit returns false, the connection is closed with the string returned as
	// the disconnect message. WARNING: Use the client data at your own risk, it
	// cannot be trusted because it can be freely changed by the player
	// connecting.
	Allow(addr net.Addr, d login.IdentityData, c login.ClientData) (string, bool)
}

// allower is the standard Allower implementation. It accepts all connections.
type allower struct{}

// Allow always returns true.
func (allower) Allow(net.Addr, login.IdentityData, login.ClientData) (string, bool) {
	return "", true
}
