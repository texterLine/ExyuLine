// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package event

// Context represents the context of an event. Handlers of an event may call methods on the context to change
// the result of the event.
type Context[T any] struct {
	cancel bool
	val    T
}

// C returns a new event context.
func C[T any](v T) *Context[T] {
	return &Context[T]{val: v}
}

// Val returns the subject of the Context.
func (ctx *Context[T]) Val() T {
	return ctx.val
}

// Cancelled returns whether the context has been cancelled.
func (ctx *Context[T]) Cancelled() bool {
	return ctx.cancel
}

// Cancel cancels the context.
func (ctx *Context[T]) Cancel() {
	ctx.cancel = true
}
