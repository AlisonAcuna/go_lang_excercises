package main

// which games would have been possible if the bag contained
// only 12 red cubes, 13 green cubes, and 14 blue cubes?
// What is the sum of the IDs of those games?

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ColorCount map[string]int

func main() {
	input := readIntsFromFile("/Users/alison.acuna/Code/advent_of_code/cubeconundrum/day_two_input.txt")

	_, err := parseGames(input)
	// fmt.Print(games[0])
	fmt.Print(err)
	// Write function that counts the number of games which never pull more than 12 red cubes, 13 green cubes, or 14 blue cubes
}

func readIntsFromFile(filename string) []string {
	file, _ := os.Open(filename)

	defer file.Close()
	var games []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		games = append(games, data)
	}
	return games
}

func removeGamePrefix(input string) string {
	colonIndex := strings.Index(input, ":")
	if colonIndex != -1 && strings.HasPrefix(input, "Game ") {
		return strings.TrimSpace(input[colonIndex+1:])
	}
	return input
}

func parseGames(input []string) ([]ColorCount, error) {
	var result []ColorCount
	for _, game := range input {
		rounds := strings.Split(game, ";")
		fmt.Print(rounds[0])
		for _, round := range rounds {
			round = removeGamePrefix(round)
			// START HERE: Remove "Game" to address: strconv.Atoi: parsing "Game": invalid syntax
			colorCount := make(ColorCount)
			pairs := strings.Split(strings.TrimSpace(round), ",")
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				parts := strings.Split(pair, " ")
				if len(parts) < 2 {
					continue
				}
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, err
				}
				color := parts[1]
				colorCount[color] = count
			}
			result = append(result, colorCount)
		}

	}
	return result, nil
}
