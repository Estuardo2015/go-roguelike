package modules

type MagicComponent struct {
	Spells      map[string]Spell
	LoadedSpell Spell
}

type Spell interface {
	Cast(target Entity)
	Name() string
}
