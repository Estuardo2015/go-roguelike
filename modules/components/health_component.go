package components

type HealthComponent struct {
	health    int
	maxHealth int
}

func (hc HealthComponent) SetHealth(hp, maxHp int) {
	hc.health = hp
	hc.maxHealth = maxHp
}

func (hc HealthComponent) Health() int {
	return hc.health
}

func (hc HealthComponent) MaxHealth() int {
	return hc.maxHealth
}

func (hc HealthComponent) AddHealth(hp int) {
	hc.health = hc.health + hp
	if hc.health > hc.maxHealth {
		hc.health = hc.maxHealth
	}
}

func (hc HealthComponent) SubtractHealth(hp int) {
	hc.health = hc.health - hp
}
