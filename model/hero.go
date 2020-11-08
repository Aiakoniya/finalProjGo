package main

import "fmt"

type iBuilder interface {
	setRoles()
	setPath()
	setPower()
	getHero() hero
}

func getBuilder(builderType string) iBuilder {
	if builderType == "warrior" {
		return &warriorBuilder{}
	}

	if builderType == "wizard" {
		return &wizardBuilder{}
	}
	return nil
}

type warriorBuilder struct {
	roles string
	path  string
	power string
}

func newWarriorBuilder() *warriorBuilder {
	return &warriorBuilder{}
}

func (b *warriorBuilder) setRoles() {
	b.roles = "Toxic"
}

func (b *warriorBuilder) setPath() {
	b.path = "Warrior Heart"
}

func (b *warriorBuilder) setPower() {
	b.power = "Strengths and Skills"
}

func (b *warriorBuilder) getHero() hero {
	return hero{
		roles: b.roles,
		path:  b.path,
		power: b.power,
	}
}

type wizardBuilder struct {
	roles string
	path  string
	power string
}

func newWizardBuilder() *wizardBuilder {
	return &wizardBuilder{}
}

func (b *wizardBuilder) setRoles() {
	b.roles = "Authentic"
}

func (b *wizardBuilder) setPath() {
	b.path = "Wizards are wise"
}

func (b *wizardBuilder) setPower() {
	b.power = "magic and mystery"
}

func (b *wizardBuilder) getHero() hero {
	return hero{
		roles: b.roles,
		path:  b.path,
		power: b.power,
	}
}

type hero struct {
	name       string
	equipments []equipment
	roles      string
	path       string
	power      string
}

func (h *hero) getEquip() []equipment {
	return h.equipments
}

type action interface {
	attack()
	def()
	changeEquip() *equipment
}

func (h *hero) attack() {
	j := 0
	for range h.equipments {
		h.equipments[j].attack += h.equipments[j].attack
	}
}

type gamer struct {
	builder iBuilder
}

func newGamer(b iBuilder) *gamer {
	return &gamer{
		builder: b,
	}
}

func (g *gamer) setGamer(b iBuilder) {
	g.builder = b
}

func (g *gamer) setBuilder(b iBuilder) {
	g.builder = b
}

func (g *gamer) buildHero() hero {
	g.builder.setRoles()
	g.builder.setPath()
	g.builder.setPower()
	return g.builder.getHero()
}

func main() {
	warriorBuilder := getBuilder("warrior")
	wizardBuilder := getBuilder("wizard")

	gamer := newGamer(warriorBuilder)
	warriorHero := gamer.buildHero()

	fmt.Printf("Warrior hero's role %s\n", warriorHero.roles)
	fmt.Printf("Warrior hero's path %s\n", warriorHero.path)
	fmt.Printf("Warrior hero's power are %s\n", warriorHero.power)

	gamer.setBuilder(wizardBuilder)
	wizardHero := gamer.buildHero()

	fmt.Printf("\nWizard hero's role %s\n", wizardHero.roles)
	fmt.Printf("Wizard hero's path %s\n", wizardHero.path)
	fmt.Printf("Wizard hero's power are %s\n", wizardHero.power)
}
