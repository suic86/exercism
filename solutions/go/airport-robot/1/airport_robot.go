package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(message string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", i.LanguageName(), message)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(message string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", p.LanguageName(), message)
}

func SayHello(message string, greeter Greeter) string {
	return greeter.Greet(message)
}
