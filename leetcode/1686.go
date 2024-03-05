package main

func stoneGameVI(aliceValues []int, bobValues []int) int {
	diff := aliceValues
	for i := 0; i < len(aliceValues); i++ {
		diff[i] = aliceValues[i] - bobValues[i]
	}
	visited := make([]int, len(aliceValues))
	_ = visited
	//  turn  sritch
	turn := 0
	for {
		if turn == 0 {
			//alice

		} else {
			// bob
		}
	}
}
