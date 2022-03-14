package main

import "fmt"

func main() {
	words := []string{"bella", "label", "roller"}
	// words := []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(words))
}

func commonChars(words []string) []string {
	var stringResults []string
	results := make(map[string]int)
	for i, word := range words {
		temp := make(map[string]int)
		for _, w := range word {
			w := string(w)
			if i == 0 {
				results[w] = results[w] + 1
				continue
			}
			temp[w] = temp[w] + 1
		}
		if i > 0 {
			for k, v := range results {
				if _, isExist := temp[k]; !isExist {
					delete(results, k)
					continue
				}
				if temp[k] < v {
					results[k] = temp[k]
				}
				if i == len(words)-1 {
					if v > 1 {
						for i := 0; i < v; i++ {
							stringResults = append(stringResults, k)
						}
						continue
					}
					stringResults = append(stringResults, k)
				}
			}
			temp = nil
		}
	}
	return stringResults
}
