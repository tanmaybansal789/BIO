package main

import "fmt"

func numCombinations(ranges []InclRange) int {
	count := 1
	for _, r := range ranges {
		count *= (r.end - r.start + 1)
	}
	return count
}

func nthCombination(ranges []InclRange, n int) []int {
	combination := make([]int, len(ranges))

	for i := 0; i < len(ranges); i++ {
		blockSize := numCombinations(ranges[i+1:])

		indexInRange := (n - 1) / blockSize
		combination[i] = ranges[i].start + indexInRange

		n = (n - 1) % blockSize + 1
	}

	return combination
}


type InclRange struct {
	start, end int
}

func main() {
	var cars string
    var n int
    fmt.Scan(&cars, &n)
	numCars := len(cars)

	carToParking := make([]int, numCars)

	for parking, car := range cars {
		carIdx := int(car - 'a')
		carToParking[carIdx] = parking
	}

	carToParkingPreferences := make([]InclRange, numCars)

	for carIdx := 0; carIdx < numCars; carIdx++ {
		parking := carToParking[carIdx]

		start := parking
		end := parking

		nextPreference := parking - 1

		for nextPreference >= 0 {
			conflict := false

			for alreadyParkedCarIdx := carIdx - 1; alreadyParkedCarIdx >= 0; alreadyParkedCarIdx-- {
				alreadyParkedCarParking := carToParking[alreadyParkedCarIdx]
				if alreadyParkedCarParking == nextPreference {

					start = nextPreference
					conflict = true
					break
				}
			}

			if !conflict {
				break
			}

			nextPreference--
		}
		carToParkingPreferences[carIdx] = InclRange{start, end}
	}

	finalParking := nthCombination(carToParkingPreferences, n)
    for _, parking := range finalParking {
        fmt.Printf("%c", 'A' + parking)
    }
    fmt.Println()
}