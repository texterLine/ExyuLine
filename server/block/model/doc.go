// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

// Package model has world.BlockModel implementations for each world.Block implementation in the block package. These
// models may be reused by blocks that share the same model.
// Some block models require additional fields and knowledge about the surrounding blocks to calculate the correct
// model.
package model
