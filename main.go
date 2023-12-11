package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type galaxy struct {
	posX int
	posY int
}

func findMissingNumbers(originalSlice []int, a int) []int {
	numbersMap := make(map[int]bool)
	for _, num := range originalSlice {
		numbersMap[num] = true
	}

	missingNumbers := make([]int, 0)

	for i := 0; i < a; i++ {
		if _, exists := numbersMap[i]; !exists {
			missingNumbers = append(missingNumbers, i)
		}
	}

	return missingNumbers
}

func getExpandedX(missing []int, x int) int {
	toAdd := 0
	for _, v := range missing {
		if v < x {
			toAdd++
		}
	}
	return toAdd + x

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	var galaxies []galaxy
	length := 0
	var xPositions []int

	for scanner.Scan() {
		text := scanner.Text()
		noGalaxies := true
		length = max(length, len(text))
		for i, c := range text {
			if c == '#' {
				noGalaxies = false
				galaxies = append(galaxies, galaxy{posX: i, posY: y})
				xPositions = append(xPositions, i)

			}
		}
		if noGalaxies {
			y++
		}
		y++
	}
	noX := findMissingNumbers(xPositions, length)
	sum := 0
	visited := 0
	for k, v := range galaxies {
		for i, gal := range galaxies {
			if i <= k {
				continue
			}
			x1 := getExpandedX(noX, v.posX)
			x2 := getExpandedX(noX, gal.posX)
			distance := int(math.Abs(float64(x1-x2)) + math.Abs(float64(v.posY-gal.posY)))
			// fmt.Printf("%d %d %d\n", k+1, i+1, distance)
			sum += distance
		}
		visited++
	}
	fmt.Println(sum)
}
