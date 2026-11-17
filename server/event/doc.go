// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

// Package event exposes a single exported `Context` type that may be used to influence the execution flow of events
// that occur on a server.
// Generally, the caller of `event.C()` calls `Context.Cancelled()` to check if the `Context` was cancelled (using
// `Context.Cancel()`) by whatever code it was passed to.
// who is then able to cancel it by calling `Context.Cancel()`.
package event
