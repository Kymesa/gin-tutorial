package code

import "fmt"

func App() string {
	return "KEINER MESA"
}

// type People struct {
// 	Nombre string
// }

// func Loop() {
// 	nums := []People{{Nombre: "KEINER"}, {Nombre: "YESID"}}
// 	for _, num := range nums {
// 		fmt.Println(num)
// 	}
// }

type Multiply struct {
	NumberOne int
	NumberTwo int
}

func (n Multiply) Multiply() int {
	return n.NumberOne * n.NumberTwo
}

func Test() {
	user := map[string]string{
		"NAME": "KEINER",
	}

	for key, value := range user {

		fmt.Printf("MI KEY ES %v Y MI VALOR ES %v ", key, value)
	}

}

type Action interface {
	Ladrar() string
}

type Perro struct{}

func (p Perro) Ladrar() string {
	return "Guauu!!! "
}

func Poly() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[10])
}

func Points() {
	x := 10
	p := &x

	fmt.Println(p)
	fmt.Println(*p)

}

type People struct {
	Name    string
	Balance int
}

func (p *People) AddBalance(balance int) {
	p.Balance += balance
}

func (p *People) OutBalance(balance int) {
	p.Balance -= balance
}

func Test1() {

	people := People{Name: "KEINER", Balance: 100}

	people.AddBalance(50)
	people.OutBalance(100)

	fmt.Println(people)
}
