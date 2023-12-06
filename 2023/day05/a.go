package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type mapping struct {
	start int
	end   int
	add   int
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Could not convert string %s to int", str)
	}
	return num
}

func makeMap(group string) []*mapping {
	mappings := []*mapping{}
	for _, line := range strings.Split(group, "\n")[1:] {
		parts := strings.Split(line, " ")
		dest, src, rnge := stringToInt(parts[0]), stringToInt(parts[1]), stringToInt(parts[2])
		mappings = append(mappings, &mapping{start: src, end: src + rnge, add: dest - src})
	}
	return mappings
}

func readMap(mappings []*mapping, input int) int {
	for _, m := range mappings {
		if m.start <= input && input <= m.end {
			return input + m.add
		}
	}
	return input
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(data), "\n\n")

	seeds := strings.Split(strings.Split(groups[0], ":")[1], " ")[1:]
	seedToSoil := makeMap(groups[1])
	soilToFertilizer := makeMap(groups[2])
	fertilizerToWater := makeMap(groups[3])
	waterToLight := makeMap(groups[4])
	lightToTemperature := makeMap(groups[5])
	temperatureToHumidity := makeMap(groups[6])
	humidityToLocation := makeMap(groups[7])

	locations := []int{}
	for _, seed := range seeds {
		seedNum := stringToInt(seed)
		soil := readMap(seedToSoil, seedNum)
		fertilizer := readMap(soilToFertilizer, soil)
		water := readMap(fertilizerToWater, fertilizer)
		light := readMap(waterToLight, water)
		temperature := readMap(lightToTemperature, light)
		humidity := readMap(temperatureToHumidity, temperature)
		location := readMap(humidityToLocation, humidity)
		locations = append(locations, location)
	}
	fmt.Printf("Lowest location: %d\n", slices.Min(locations))
}
