package service

import "math/rand"

func GetRandomInt(from, to int) int {
	return rand.Intn(to-from) + from
}
