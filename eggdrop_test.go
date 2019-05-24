package eggdrop

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func TestEggDrop(t *testing.T) {
	// t.Skip()
	egg := 2
	floor := 10
	expected := 4
	actual := eggDrop(egg, floor)

	fmt.Printf("eggDrop: Throw Egg %d times\n", actual)
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestEggMemoize(t *testing.T) {
	// t.Skip()
	egg := 2
	floor := 10

	eggIdx := egg + 1
	floorIdx := floor + 1

	// create and initialize memo
	memo := make([][]int, eggIdx)
	for i := 0; i < eggIdx; i++ {
		memo[i] = make([]int, floorIdx)
	}

	for i := 0; i < eggIdx; i++ {
		for j := 0; j < floorIdx; j++ {
			memo[i][j] = MinVal
		}
	}

	expected := 4
	actual := eggMemoize(egg, floor, memo)
	fmt.Printf("eggMemoize: Throw Egg %d times\n", actual)
	fmt.Printf("%# v\n", pretty.Formatter(memo))
	assert.Equal(t, expected, actual, "they should be equal")
}

func TestEggBottomUp(t *testing.T) {
	// t.Skip()
	egg := 2
	floor := 10

	eggIdx := egg + 1
	floorIdx := floor + 1

	// create and initialize memo
	memo := make([][]int, eggIdx)
	for i := 0; i < eggIdx; i++ {
		memo[i] = make([]int, floorIdx)
	}

	for i := 0; i < eggIdx; i++ {
		for j := 0; j < floorIdx; j++ {
			memo[i][j] = MinVal
		}
	}

	expected := 4
	actual := eggBottomUp(egg, floor, memo)
	fmt.Printf("eggBottomUp: Throw Egg %d times\n", actual)
	fmt.Printf("%# v", pretty.Formatter(memo))
	assert.Equal(t, expected, actual, "they should be equal")
}

func BenchmarkEggDrop(b *testing.B) {
	// t.Skip()
	egg := 2
	floor := 10

	for i := 0; i < b.N; i++ {
		eggDrop(egg, floor)
	}
}

func BenchmarkEggMemoize(b *testing.B) {
	// t.Skip()
	egg := 3
	floor := 1000

	eggIdx := egg + 1
	floorIdx := floor + 1

	// create and initialize memo
	memo := make([][]int, eggIdx)
	for i := 0; i < eggIdx; i++ {
		memo[i] = make([]int, floorIdx)
	}

	for i := 0; i < eggIdx; i++ {
		for j := 0; j < floorIdx; j++ {
			memo[i][j] = MinVal
		}
	}

	for i := 0; i < b.N; i++ {
		eggMemoize(egg, floor, memo)
	}
}

func BenchmarkEggBottomUp(b *testing.B) {
	// t.Skip()
	egg := 3
	floor := 1000

	eggIdx := egg + 1
	floorIdx := floor + 1

	// create and initialize memo
	memo := make([][]int, eggIdx)
	for i := 0; i < eggIdx; i++ {
		memo[i] = make([]int, floorIdx)
	}

	for i := 0; i < eggIdx; i++ {
		for j := 0; j < floorIdx; j++ {
			memo[i][j] = MinVal
		}
	}

	for i := 0; i < b.N; i++ {
		eggBottomUp(egg, floor, memo)
	}
}
