package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name    string
	HP      int
	MaxHP   int
	Damage  int
	IsAlive bool
}

func NewHero(name string, damage int) *Hero {
	return &Hero{Name: name, HP: 100, MaxHP: 100, Damage: damage, IsAlive: true}
}

func (h *Hero) Attack(target *Hero) {
	if !h.IsAlive || !target.IsAlive {
		return
	}

	maxPossibleDamage := int(float64(target.HP) * 0.4)
	if maxPossibleDamage < 5 {
		maxPossibleDamage = 5
	}

	actualDamage := rand.Intn(h.Damage) + 1
	if actualDamage > maxPossibleDamage {
		actualDamage = maxPossibleDamage
	}

	fmt.Printf(" %s нанес %d урона %s\n", h.Name, actualDamage, target.Name)
	target.TakeDamage(actualDamage)
}

func (h *Hero) TakeDamage(damage int) {
	h.HP -= damage
	if h.HP <= 0 {
		h.HP = 0
		h.IsAlive = false
	}
}

func (h *Hero) Heal() {
	if !h.IsAlive || h.HP >= h.MaxHP {
		return
	}

	missingHP := h.MaxHP - h.HP
	maxHeal := int(float64(missingHP) * 0.5)
	if maxHeal < 1 {
		maxHeal = 1
	}

	actualHeal := rand.Intn(maxHeal) + 5
	h.HP += actualHeal
	if h.HP > h.MaxHP {
		h.HP = h.MaxHP
	}
	fmt.Printf(" %s восстановил %d HP. Итого: %d\n", h.Name, actualHeal, h.HP)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p1 := NewHero("Гарри Поттер", 50)
	p2 := NewHero("Питер Паркер", 50)

	p1.Attack(p2)
	p2.Heal()
}
