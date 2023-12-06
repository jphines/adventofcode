package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type AlmanacEntry struct {
	Destination int
	Source      int
	Range       int
}

func (a *AlmanacEntry) String() string {
	return fmt.Sprintf(
		"`destination: %d, source: %d, range: %d`\n",
		a.Destination,
		a.Source,
		a.Range,
	)
}

func (a *AlmanacEntry) InRange(s int) bool {
	return s >= a.Source && s <= a.Source+a.Range
}

func (a *AlmanacEntry) Convert(source int) int {
	return (a.Destination - a.Source) + source
}

type SeedRange struct {
	Start int
	Range int
}

type Almanac struct {
	Seeds    []int
	AllSeeds []*SeedRange

	SeedToSoil        []*AlmanacEntry
	SoilToFertilizer  []*AlmanacEntry
	FertilizerToWater []*AlmanacEntry
	WaterToLight      []*AlmanacEntry
	LightToTemp       []*AlmanacEntry
	TempToHumidity    []*AlmanacEntry
	HumidityToLoc     []*AlmanacEntry
}

func (a *Almanac) String() string {
	return fmt.Sprintf(
		"seeds: %v\n\n"+
			"seed-to-soil \n%s\n\n"+
			"soil-to-fertilizer \n%s\n\n"+
			"fertilizer-to-water \n%s\n\n"+
			"water-to-light \n%s\n\n"+
			"light-to-temperature \n%s\n\n"+
			"temperature-to-humidity \n%s\n\n"+
			"humidity-to-location \n%s\n",
		a.Seeds,
		a.SeedToSoil,
		a.SoilToFertilizer,
		a.FertilizerToWater,
		a.WaterToLight,
		a.LightToTemp,
		a.TempToHumidity,
		a.HumidityToLoc,
	)
}

func (a *Almanac) Convert(source int, entries []*AlmanacEntry) int {
	for _, entry := range entries {
		if entry.InRange(source) {
			return entry.Convert(source)
		}
	}
	return source
}

func (a *Almanac) Walk(seed int) int {
	soil := a.Convert(seed, a.SeedToSoil)
	fertilizer := a.Convert(soil, a.SoilToFertilizer)
	water := a.Convert(fertilizer, a.FertilizerToWater)
	light := a.Convert(water, a.WaterToLight)
	temp := a.Convert(light, a.LightToTemp)
	humidity := a.Convert(temp, a.TempToHumidity)
	loc := a.Convert(humidity, a.HumidityToLoc)
	return loc
}

func parseSeeds(input string) []int {
	postHead := strings.Split(input, ": ")[1]
	nums := strings.Split(postHead, " ")

	seeds := make([]int, 0)
	for _, num := range nums {
		seed, _ := strconv.Atoi(num)
		seeds = append(seeds, seed)
	}

	return seeds
}

func parseAllSeeds(input string) []*SeedRange {
	postHead := strings.Split(input, ": ")[1]
	nums := strings.Split(postHead, " ")

	seedRange := make([]*SeedRange, 0)

	for i := 0; i < len(nums); i += 2 {
		source, _ := strconv.Atoi(nums[i])
		range_, _ := strconv.Atoi(nums[i+1])
		seedRange = append(seedRange, &SeedRange{source, range_})
	}
	return seedRange
}

func parseEntry(entry string) []*AlmanacEntry {
	lines := strings.Split(entry, "\n")

	entries := make([]*AlmanacEntry, len(lines)-1)

	for i, line := range lines[1:] {
		re := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)

		destination, _ := strconv.Atoi(matches[0][1])
		source, _ := strconv.Atoi(matches[0][2])
		range_, _ := strconv.Atoi(matches[0][3])

		entries[i] = &AlmanacEntry{destination, source, range_}
	}

	return entries
}

func parseInput(input string) *Almanac {
	split := strings.Split(input, "\n\n")

	almanac := &Almanac{}
	almanac.Seeds = parseSeeds(split[0])
	almanac.AllSeeds = parseAllSeeds(split[0])

	almanac.SeedToSoil = parseEntry(split[1])
	almanac.SoilToFertilizer = parseEntry(split[2])
	almanac.FertilizerToWater = parseEntry(split[3])
	almanac.WaterToLight = parseEntry(split[4])
	almanac.LightToTemp = parseEntry(split[5])
	almanac.TempToHumidity = parseEntry(split[6])
	almanac.HumidityToLoc = parseEntry(split[7])

	return almanac
}

func partOne(input string) {
	almanac := parseInput(input)

	minimumLoc := math.MaxInt
	for _, seed := range almanac.Seeds {
		loc := almanac.Walk(seed)
		minimumLoc = min(minimumLoc, loc)
	}

	fmt.Println(minimumLoc)
}

func partTwo(input string) {
	almanac := parseInput(input)

	minimumLoc := math.MaxInt
	for _, sr := range almanac.AllSeeds {
		fmt.Println("trying", sr.Start, sr.Range)
		for seed := sr.Start; seed < sr.Start+sr.Range; seed++ {
			loc := almanac.Walk(seed)
			minimumLoc = min(minimumLoc, loc)
		}
	}

	fmt.Println(minimumLoc)
}

func dayFive(input string) {
	partOne(input)
	partTwo(input)
}
