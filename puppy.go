package puppy

import (
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
