// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

// Package server implements a high-level implementation of a Minecraft server.
// Creating such a server may be done using the `server.New()` function.
// Additional configuration may be passed by using server.Config{...}.New().
// `Server.Listen()` may be called to start and run the server. It should be
// followed up by a loop as such:
//
//	for p := range srv.Accept() {
//	  // Use p
//	}
//
// `Server.Accept()` blocks until a new player connects to the server and
// spawns in the default world.
package server
