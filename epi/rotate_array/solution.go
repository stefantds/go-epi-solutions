package rotate_array

import "github.com/stefantds/go-epi-judge/epi/euclidean_gcd"

func RotateArrayCyclicPermutations(rotateAmount int, a []int) {
	if rotateAmount == 0 {
		return
	}
	n := len(a)
	numCycles := int(euclidean_gcd.EuclideanGCD(rotateAmount, n))
	cycleLength := n / numCycles

	applyCyclicPerm := func(offset int) {
		temp := a[offset]

		for i := 1; i < cycleLength; i++ {
			pos := (offset + i*rotateAmount) % n
			current := a[pos]
			a[pos] = temp
			temp = current
		}
		a[offset] = temp
	}

	for i := 0; i < numCycles; i++ {
		applyCyclicPerm(i)
	}
}

func RotateArrayUsingReverse(rotateAmount int, a []int) {
	if rotateAmount == 0 {
		return
	}
	n := len(a)
	rotateAmount = rotateAmount % n

	reverse := func(startPos, endPos int) {
		for s, e := startPos, endPos; s < e; s, e = s+1, e-1 {
			a[s], a[e] = a[e], a[s]
		}
	}

	reverse(0, n-1)
	reverse(0, rotateAmount-1)
	reverse(rotateAmount, n-1)
}
