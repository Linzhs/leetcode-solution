package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	var totalTank, curTank, startStation int
	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i] // 预先记录遍历过的消耗
		curTank += gas[i] - cost[i]
		if curTank < 0 {
			startStation = i + 1
			curTank = 0
		}
	}

	if totalTank >= 0 {
		return startStation
	}
	return -1
}
