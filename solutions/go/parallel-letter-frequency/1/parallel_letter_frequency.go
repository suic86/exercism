package letter

import (
	"sync"
)

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

// ConcurrentFrequency counts the frequence of each rune in a given array of strings
// and returns this a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	fm := FreqMap{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	counter := func(text string) {
		defer wg.Done()
		for _, r := range text {
			mu.Lock()
			fm[r]++
			mu.Unlock()
		}
	}

	for _, text := range texts {
		wg.Add(1)
		go counter(text)
	}
	wg.Wait()
	return fm
}
