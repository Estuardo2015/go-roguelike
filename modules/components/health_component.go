package components

type HealthComponent struct {
	health    int
	maxHealth int
}

func (hc HealthComponent) SetHealth(hp, maxHp int) {
	hc.health = hp
	hc.maxHealth = maxHp
}

func (hc HealthComponent) SubtractHealth(hp int) {
	hc.health = hc.health - hp
}

func (hc HealthComponent) AddHealth(hp int) {
	hc.health = hc.health + hp
	if hc.health > hc.maxHealth {
		hc.health = hc.maxHealth
	}
}
