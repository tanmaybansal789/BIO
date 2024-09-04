package main

import (
	"fmt"
    "unicode/utf8"
	)

const MaxLetters = 26
const MaxScore = 75

var DPTable [MaxScore + 1][MaxLetters]int

func Sum(nums []int) int {
    result := 0

    for _, num := range nums {
        result += num
    }

    return result
}

func CalculateScores(word string) []int {
    scores := make([]int, utf8.RuneCountInString(word))

    for i, char := range word {
        scores[i] = int(char) - 64
    }

    return scores
}

func InitDPTable() {
	for lastLetterIdx := 0; lastLetterIdx < MaxLetters; lastLetterIdx++ {
		DPTable[0][lastLetterIdx] = 1
	}

	for score := 1; score <= MaxScore; score++ {
        
		for lastLetterIdx := 0; lastLetterIdx < MaxLetters; lastLetterIdx++ {
            
			lastLetterScore := lastLetterIdx + 1
            
			nextLetterCap := min(score, MaxLetters)
            
			for nextLetterIdx := 0; nextLetterIdx < nextLetterCap; nextLetterIdx++ {
                
				nextLetterScore := nextLetterIdx + 1
                
				if nextLetterScore != lastLetterScore {
					DPTable[score][lastLetterIdx] += DPTable[max(score - nextLetterScore, 0)][nextLetterIdx]
				}
			}
		}
	}
}

func PositionOfWord(word string) int {
	scores := CalculateScores(word)
    score := Sum(scores)
    
	position := 1

	for i := 0; i < len(word); i++ {
        
		charScore := scores[i]

		for lastCharScore := 1; lastCharScore < charScore; lastCharScore++ {

			if i == 0 || lastCharScore != scores[i-1] {
				position += DPTable[max(score - lastCharScore, 0)][lastCharScore - 1]
			}
		}

		score -= charScore

		if score < 0 {
			break
		}
	}

	return position
}

func main() {
	InitDPTable()

    var word string

    _, err := fmt.Scan(&word)
    if err != nil {
        return
    }

	position := PositionOfWord(word)

	fmt.Println(position)
}