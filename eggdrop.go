package eggdrop

import "strconv"

const MaxIntVal = int(^uint(0) >> 1)
const MinIntVal = -MaxIntVal - 1
const MinVal = -1

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func eggDrop(egg int, floor int) int {
	if floor <= 1 || egg == 1 { // 0 or 1 floor left, or the last egg
		return floor
	}

	min := MaxIntVal
	for cFloor := 1; cFloor <= floor; cFloor++ {
		result := 1 + MaxInt(eggDrop(egg-1, cFloor-1), eggDrop(egg, floor-cFloor))
		if result < min {
			min = result
		}
	}
	return min
}

func eggMemoize(egg, floor int, memo [][]int) int {
	if memo[egg][floor] != MinVal {
		return memo[egg][floor]
	}

	if floor <= 1 || egg == 1 { // 0 or 1 floor left, or the last egg
		memo[egg][floor] = floor
		return floor
	}

	min := strconv.IntSize
	for cFloor := 1; cFloor <= floor; cFloor++ {
		result := 1 + MaxInt(eggMemoize(egg-1, cFloor-1, memo), eggMemoize(egg, floor-cFloor, memo))
		if result < min {
			min = result
		}
	}
	memo[egg][floor] = min
	return min
}

func eggBottomUp(egg, floor int, memo [][]int) int {

	for cFloor := 1; cFloor <= floor; cFloor++ {
		for e := 1; e <= egg; e++ { // egg left
			if cFloor <= 1 || e == 1 { // 0 or 1 floor left, or the last egg
				memo[e][cFloor] = cFloor
			} else {
				min := MaxIntVal
				for k := 1; k <= cFloor; k++ {
					result := 1 + MaxInt(memo[e-1][k-1], memo[e][cFloor-k])
					if result < min {
						min = result
					}
				}
				memo[e][cFloor] = min
			}
		}
	}
	return memo[egg][floor]
}
