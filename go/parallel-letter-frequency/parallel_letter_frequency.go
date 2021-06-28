package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given
// array of strings in a seperate GoRoutine and returns
// the data as a FreqMap
// Note that this may not be faster for small texts, ex:
// cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
// BenchmarkSequentialFrequency
// BenchmarkSequentialFrequency-8             69218             17016 ns/op
// BenchmarkConcurrentFrequency
// BenchmarkConcurrentFrequency-8             46052             26011 ns/op
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	for _, text := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(text)
	}
	resultMap := FreqMap{}
	// Receive as many values from the channel as there are input texts
	for range texts {
		m := <-c
		for k, e := range m {
			resultMap[k] += e
		}
	}
	return resultMap
}
