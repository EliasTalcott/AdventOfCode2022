package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	start int
	end   int
	add   int
}

type seedRange struct {
	start int
	end   int
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
		mappings = append(mappings, &mapping{start: dest, end: dest + rnge, add: src - dest})
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

func makeSeedRanges(seeds string) []*seedRange {
	ranges := []*seedRange{}
	parts := strings.Split(seeds, " ")[1:]
	for i := 0; i < len(parts); i += 2 {
		start := stringToInt(parts[i])
		length := stringToInt(parts[i+1])
		ranges = append(ranges, &seedRange{start: start, end: start + length})
	}
	return ranges
}

func seedExists(seeds []*seedRange, seed int) bool {
	for _, r := range seeds {
		if r.start <= seed && seed <= r.end {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(data), "\n\n")

	seeds := makeSeedRanges(strings.Split(groups[0], ":")[1])
	soilToSeed := makeMap(groups[1])
	fertilizerToSoil := makeMap(groups[2])
	waterToFertilizer := makeMap(groups[3])
	lightToWater := makeMap(groups[4])
	temperatureToLight := makeMap(groups[5])
	humidityToTemperature := makeMap(groups[6])
	locationToHumidity := makeMap(groups[7])

	location := 0
	for {
		humidity := readMap(locationToHumidity, location)
		temperature := readMap(humidityToTemperature, humidity)
		light := readMap(temperatureToLight, temperature)
		water := readMap(lightToWater, light)
		fertilizer := readMap(waterToFertilizer, water)
		soil := readMap(fertilizerToSoil, fertilizer)
		seed := readMap(soilToSeed, soil)
		if seedExists(seeds, seed) {
			break
		}
		location++
	}
	fmt.Printf("Lowest location: %d\n", location)
}
