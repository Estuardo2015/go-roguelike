package components

type StatComponent struct {
	Stats map[string]Stat
}

func (sc StatComponent) StatValue(stat string) int {
	return sc.Stats[stat].value
}

type Stat struct {
	name  string
	value int
}

/*
    Concentration: How much magic one can perform. Governed by intellect and dexterity
	Encumbrance: How much one can carry. Governed by strength

	Experience
	Level

	Strength: Governs one's carry weight, damage resistance,
	Dexterity: Governs one's chance to successfully cast spells, hit opponents, and concentration
	Intellect: Affects xp gained
	Charisma: Governs bartering ability and persuasion
*/
