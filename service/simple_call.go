package service

import "fmt"

func A() {
	name := B()("test")
	fmt.Println(name)
}

func B() func(string) string {
	return func(s string) string {
		return "called"
	}
}
