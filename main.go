package main

import "fmt"

type Hero struct {
	Name    string
	HP      int
	MaxHP   int
	Damage  int
	IsAlive bool
}

func NewHero(name string, damage int) *Hero {
	return &Hero{
		Name:    name,
		HP:      100,
		MaxHP:   100,
		Damage:  damage,
		IsAlive: true,
	}
}

func (h *Hero) Attack(target *Hero) {
	if !h.IsAlive || !target.IsAlive {
		return
	}
	target.TakeDamage(h.Damage)
}

func (h *Hero) TakeDamage(damage int) {
	h.HP -= damage
	if h.HP <= 0 {
		h.HP = 0
		h.IsAlive = false
	}
}

func (h *Hero) Heal(amount int) {
	if !h.IsAlive {
		return
	}
	h.HP += amount
	if h.HP > h.MaxHP {
		h.HP = h.MaxHP
	}
}

func (h *Hero) Status() {
	if h.IsAlive {
		fmt.Printf("[%s] HP: %d\n", h.Name, h.HP)
	} else {
		fmt.Printf("[%s] Повержен\n", h.Name)
	}
}

func main() {
	p1 := NewHero("Игрок 1", 25)
	p2 := NewHero("Игрок 2", 20)

	p1.Attack(p2)
	p2.Attack(p1)
	p1.Heal(10)

	p1.Status()
	p2.Status()
}
