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

// ConcurrentFrequency counts the frequence of each rune in a given array of strings
// and returns this a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	fm := FreqMap{}

	c := make(chan FreqMap)

	for _, text := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(text)
	}

	for range texts {
		for r, cnt := range <-c {
			fm[r] += cnt
		}
	}

	return fm
}
