// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

// Package block exports implementations of the Block interface found in the server/world package. The blocks
// implemented in this package are automatically registered in the server/world package using the world.RegisterBlock
// and world.RegisterItem functions through an init function, so that methods in server/world that return blocks and
// items are able to return them by the respective types implemented here.
package block
