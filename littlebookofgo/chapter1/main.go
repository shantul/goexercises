package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
	// 	goku := &Saiyan{
	// 		Name:  "Goku",
	// 		Power: 9000,
	// 	}
	//
	// 	fmt.Println(goku)
	//
	// 	Super(goku)
	//
	// 	fmt.Println(goku)
	//
	// 	goku.Super()
	// 	fmt.Println(goku)
	//
	// 	goku = NewSaiyan("Gohan", 5000)
	// 	fmt.Println(goku)
	//
	// 	gabru := NewerSaiyan("Gabru", 6000)
	// 	fmt.Println(gabru)
	//
	// 	goti := &Saiyan{
	// 		Name:  "Goti",
	// 		Power: 1000,
	// 		Father: &Saiyan{
	// 			Name:   "Goku",
	// 			Power:  9000,
	// 			Father: nil,
	// 		},
	// 	}
	//
	// 	fmt.Println(goti)
	// }
	//
	// func Super(s *Saiyan) {
	// 	s.Power += 10000
	// 	// s = &Saiyan{"Jack", 1000}
	// 	fmt.Println(s)
	// }
	//
	// func (s *Saiyan) Super() {
	// 	s.Power += 10000
	// }
	//
	// func NewSaiyan(name string, power int) *Saiyan {
	// 	return &Saiyan{
	// 		Name:  name,
	// 		Power: power,
	// 	}
	// }
	//
	// func NewerSaiyan(name string, power int) Saiyan {
	// 	return Saiyan{
	// 		Name:  name,
	// 		Power: power,
	// 	}
}

func Super(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}
