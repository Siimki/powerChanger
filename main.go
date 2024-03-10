package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"math"
)

func main() {
	gpxFile, err := os.Open("zwift.gpx")
	if err != nil {
		fmt.Println("Error opening GPX file:", err)
		return
	}
	defer gpxFile.Close()

	gpxData, err := ioutil.ReadAll(gpxFile)
	if err != nil {
		fmt.Println("Error reading GPX file:", err)
		return
	}

	gpxString := string(gpxData)

	re := regexp.MustCompile(`(?m)<power>(\d+)</power>`)

	// Find all matches for the regular expression
	matches := re.FindAllStringSubmatch(gpxString, -1)

	// Loop through the matches and multiply the power value by 1.2 for 20% power increase
	
	for _, match := range matches {
		power, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error parsing power value:", err)
			continue
		}
		newPower := math.Trunc(float64(power) * 1.2)
		
		gpxString = strings.Replace(gpxString, match[0], fmt.Sprintf("<power>%.1f</power>", math.RoundToEven(newPower)),1)
	}

	re2 := regexp.MustCompile(`(?m)<gpxtpx:hr>(\d+)</gpxtpx:hr>`)
	
	matches2 := re2.FindAllStringSubmatch(gpxString, -1)
	
	for _, match := range matches2 {
		heartRate, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error parsing power value:", err)
			continue
		}
		newHeartRate := float64(heartRate) * 0.9
		gpxString = strings.Replace(gpxString, match[0], fmt.Sprintf("<gpxtpx:hr>%.1f</gpxtpx:hr>", math.RoundToEven(newHeartRate)),1)
	}

	
	err = ioutil.WriteFile("zwift_modified.gpx", []byte(gpxString), 0644)
	if err != nil {
		fmt.Println("Error saving modified GPX file:", err)
		return
	}

	fmt.Println("Done.")
}