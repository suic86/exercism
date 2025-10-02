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
	var wg sync.WaitGroup
	c := make(chan rune)

	reader := func(text string) {
		defer wg.Done()
		for _, r := range text {
			c <- r
		}
	}

	for _, text := range texts {
		wg.Add(1)
		go reader(text)
	}

	go func() {
		for r := range c {
			fm[r]++
		}
	}()

	wg.Wait()
	return fm
}
