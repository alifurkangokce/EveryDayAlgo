package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var maxRel int = releaseTimes[0]
	var maxKey byte = keysPressed[0]
	for i := 1; i <= len(releaseTimes)-1; i++ {
		if releaseTimes[i]-releaseTimes[i-1] >= maxRel {
			if releaseTimes[i]-releaseTimes[i-1] == maxRel {
				if keysPressed[i] > maxKey {
					maxKey = keysPressed[i]
				}
			} else {
				maxKey = keysPressed[i]
			}
			maxRel = releaseTimes[i] - releaseTimes[i-1]
		}
	}
	return maxKey
}
