package pkg

import "math/rand"

func reservoirSampling(sampleSize int, dataStream []int) {
	reservoir := make([]int, 0, sampleSize)
	reservoir = append(reservoir, dataStream[:sampleSize]...)

	for i := sampleSize; i < len(dataStream); i++ {
		d := rand.Intn(i + 1)
		if d < sampleSize {
			reservoir[d] = dataStream[i]
		}
	}
}
