// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package cmd

import "sync"

// commands holds a list of registered commands indexed by their name.
var commands sync.Map

// Register registers a command with its name and all aliases that it has. Any command with the same name or
// aliases will be overwritten.
func Register(command Command) {
	commands.Store(command.name, command)
	for _, alias := range command.aliases {
		commands.Store(alias, command)
	}
}

// ByAlias looks up a command by an alias. If found, the command and true are returned. If not, the returned
// command is nil and the bool is false.
func ByAlias(alias string) (Command, bool) {
	command, ok := commands.Load(alias)
	if !ok {
		return Command{}, false
	}
	return command.(Command), ok
}

// Commands returns a map of all registered commands indexed by the alias they were registered with.
func Commands() map[string]Command {
	cmd := make(map[string]Command)
	commands.Range(func(key, value any) bool {
		cmd[key.(string)] = value.(Command)
		return true
	})
	return cmd
}
