package funcfortests

import (
	"math/rand/v2"
	"runtime"
	"time"
)

func GenerateArr(n, min, max int) []int {
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.IntN(max-min+1)+min)
	}
	return arr
}

func CalculateTime(function func()) float64 {
	start := time.Now()
	function()
	return time.Since(start).Seconds()
}

func CalculateMemory(function func()) float64 {
	runtime.GC()

	var memStart, memEnd runtime.MemStats
	runtime.ReadMemStats(&memStart)

	function()

	runtime.ReadMemStats(&memEnd)

	allocatedBytes := memEnd.TotalAlloc - memStart.TotalAlloc

	return float64(allocatedBytes) / (1024 * 1024)
}
