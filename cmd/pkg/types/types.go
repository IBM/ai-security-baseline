// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package types

import "github.com/gemaraproj/go-gemara"

var ControlFamilies = []string{
	"GA",
	"SC",
	"DI",
	"MR",
	"DE",
	"TE",
}

// Struct for holding the entire YAML structure
type Baseline struct {
	Catalog gemara.ControlCatalog
	Lexicon []LexiconEntry
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
	References []string `yaml:"references"`
}
