package puppy

import (
	"fmt"

	"github.com/OskarKwoczka/dog"
)

func Bark() string {
	return "Woof!"
}
func Barks() string {
	return "WOOF! WOOF! WOOF!"
}

func BigBark() string {
	return dog.WhenGrowUp(Barks())
}
func From11() {
	fmt.Println("I'm from v1.1.0")
}
func From13() {
	fmt.Println("I'm from v1.3.0")
}
