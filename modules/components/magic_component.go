package components

import "github.com/Estuardo2015/rogue_wizard/modules"

type MagicComponent struct {
	Spells      map[string]Spell
	LoadedSpell Spell
}

type Spell interface {
	Cast(target modules.Entity)
	Name() string
}
