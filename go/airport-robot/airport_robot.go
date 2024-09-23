package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

type Italian struct {
}

type Portuguese struct {
}

func (p Italian) Greet(a string) string {
	return fmt.Sprintf("Ciao %s!", a)
}
func (p Italian) LanguageName() string {
	return "Italian"
}

func (p Portuguese) Greet(a string) string {
	return fmt.Sprintf("Olá %s!", a)
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, p Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", p.LanguageName(), p.Greet(name))
}

// func main() {
// 	italian := Italian{}
// 	portuguese := Portuguese{}

// 	fmt.Println(SayHello("Alice", italian))
// 	fmt.Println(SayHello("Bob", portuguese))
// }
