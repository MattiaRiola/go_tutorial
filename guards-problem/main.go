package main

import (
	"fmt"
	"math"
)

type Guard struct {
	PrisonerWatched []uint32
}

type Prisoner struct {
	EscapePossibility uint32
	Qi                uint32
}

func main() {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("#######\tWELCOME\t#######")

	const n = 6
	const g = 3
	fmt.Println("input:")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("g: %v\n", g)
	input := [n]uint32{11, 11, 11, 24, 26, 100}
	for i, p := range input {
		fmt.Printf("p%v: %v\n", i, p)
	}

	//F.O. min {max{(Prisoners.EscapePossibility)}}

	maxEscapeProb := 0

	var guards [g]Guard
	for i := 0; i < g; i++ {
		guards[i] = Guard{PrisonerWatched: []uint32{}}
	}

	var prisoners [n]Prisoner
	for i := 0; i < n; i++ {
		prisoners[i] = Prisoner{EscapePossibility: math.MaxInt32, Qi: input[i]}
	}

	//initial situation: guards watch 0 prisoners and prisoners escapePossibility is max)

	fmt.Println("output:")
	fmt.Printf("maxEscapeProb: %v\n", maxEscapeProb)
	fmt.Println("Bye bye")

}

// k can't be 0
func calcEscape(k uint32, Qi uint32) uint32 {
	if k == 0 {
		fmt.Println("ERROR invalid k")
		return math.MaxUint32
	}
	return k * Qi
}
