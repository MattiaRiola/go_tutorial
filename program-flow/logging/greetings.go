package logging

import (
	"fmt"
)

var name string

func Name(b string) {
	name = b
}
func Greetings() {

	fmt.Printf("Hello %s!\n", name)

}
