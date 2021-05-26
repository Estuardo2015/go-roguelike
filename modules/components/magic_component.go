package components

import "github.com/Estuardo2015/rogue_wizard/modules/entity"

type MagicComponent struct {
	spells      map[string]Spell
	loadedSpell Spell
}

func (mc MagicComponent) AddSpell(s Spell) {
	mc.spells[s.Name()] = s
}

func (mc MagicComponent) RemoveSpell(name string) {
	delete(mc.spells, name)
}

func (mc MagicComponent) CastSpell(e entity.Entity) {
	if mc.loadedSpell != nil {
		mc.loadedSpell.Cast(e)
	}
}

type Spell interface {
	Cast(target entity.Entity)
	Name() string
}
