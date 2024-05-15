package main

import "fmt"

type IHero interface {
	setName(n string)
	setPower(p int)
	setHealth(h int)
	getName() string
	getPower() int
	getHealth() int
}

// concrete hero
type Hero struct {
	Name   string
	Power  int
	Health int
}

func (h *Hero) setName(n string) {
	h.Name = n
}
func (h *Hero) setPower(p int) {
	h.Power = p
}
func (h *Hero) setHealth(health int) {
	h.Health = health
}
func (h *Hero) getName() string {
	return h.Name
}
func (h *Hero) getPower() int {
	return h.Power
}
func (h *Hero) getHealth() int {
	return h.Health
}

// waterhero
type WaterHero struct {
	Hero
}

func newWaterHero() IHero {
	return &WaterHero{
		Hero: Hero{
			Name:   "Aquaman",
			Power:  7,
			Health: 8,
		},
	}
}

// airHero
type AirHero struct {
	Hero
}

func newAirHero() IHero {
	return &AirHero{
		Hero: Hero{
			Name:   "Superman",
			Power:  10,
			Health: 10,
		},
	}
}

// factory using switch and type struct
func getHero(heroType IHero) (IHero, error) {
	switch heroType.(type) {
	case *WaterHero:
		return newWaterHero(), nil
	case *AirHero:
		return newAirHero(), nil
	default:
		return nil, fmt.Errorf("wrong hero type")
	}
}

// factory
// func getHero(heroType string) (IHero, error) {
// 	if heroType == "WaterHero" {
// 		return newWaterHero(), nil
// 	}
// 	if heroType == "AirHero" {
// 		return newAirHero(), nil
// 	}
// 	return nil, fmt.Errorf("wrong hero type")
// }

// client
func main() {
	aquaman, _ := getHero(&WaterHero{})
	superman, _ := getHero(&AirHero{})

	printDetails(aquaman)
	printDetails(superman)
}

func printDetails(h IHero) {
	fmt.Printf("Hero: %s\n", h.getName())
	fmt.Printf("Power: %d\n", h.getPower())
	fmt.Printf("Health: %d\n", h.getHealth())
}
