package components

import "github.com/Estuardo2015/rogue_wizard/modules/entity"

type MagicComponent struct {
	Spells      map[string]Spell
	LoadedSpell Spell
}

type Spell interface {
	Cast(target entity.Entity)
	Name() string
}
