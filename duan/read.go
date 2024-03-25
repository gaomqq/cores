package duan

import (
	"math/rand"
)

func Read() int {

	rand.Seed(11)
	b := rand.Intn(9000) + 1000
	return b
}
