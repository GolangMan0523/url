package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Output file
	outputFile := "go_output.txt"
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	// Input URL
	inputURL := "https://antm-pt-prod-dataz-nogbd-nophi-us-east1.s3.amazonaws.com/anthem/2024-04-01_anthem_index.json.gz"

	//Start time
	startTime := time.Now()

	// Regular expressions
	newYorkRegex := regexp.MustCompile(`\b[Nn]ew\s*Y(?:ork)?\b`)
	jsonObjectRegex := regexp.MustCompile(`\{[^{}]*\}(?:\s*(?:,|$))`)
	// HTTP GET request
	response, err := http.Get(inputURL)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer response.Body.Close()

	// Decompress gzip
	decompressor, err := gzip.NewReader(response.Body)
	if err != nil {
		fmt.Println("Error creating decompressor:", err)
		return
	}
	defer decompressor.Close()

	// Read and process data
	reader := bufio.NewReader(decompressor)
	var totalCount int
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading data:", err)
			return
		}
		if err == io.EOF {
			break
		}

		// Extract URLs
		objects := jsonObjectRegex.FindAllString(line, -1)
		for _, stringToParse := range objects {
			// Parsing JSON
			description := getDescription(stringToParse)
			location := getLocation(stringToParse)
			if description != "" && location != "" {
				if newYorkRegex.MatchString(description) && strings.Contains(description, "PPO") {
					// Write to output file
					_, err := fmt.Fprintf(output, "%d:::%s\n\n", totalCount+1, location)
					if err != nil {
						fmt.Println("Error writing to output file:", err)
						return
					}
					totalCount++
					if totalCount%1000 == 0 {
						currentTime := time.Now()
						fmt.Printf("%dk URL abstracted\n", totalCount/1000)
						fmt.Println(currentTime.Sub(startTime), "  seconds")
					}
				}
			}
		}
	}
	endTime := time.Now()
	fmt.Println("URL extraction completed.")
	fmt.Println("Total time taken:", endTime.Sub(startTime))
}

// Structs to represent JSON objects
type JSONObject struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

// Function to parse JSON and extract description
func getDescription(jsonString string) string {
	var obj JSONObject
	var newJsonString string
	if len(jsonString) > 0 && jsonString[len(jsonString)-1] == ',' {
		newJsonString = jsonString[:len(jsonString)-1]
	} else {
		newJsonString = jsonString
	}
	err := json.Unmarshal([]byte(newJsonString), &obj)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}
	return obj.Description
}

// Function to parse JSON and extract location
func getLocation(jsonString string) string {
	var obj JSONObject
	var newJsonString string
	if len(jsonString) > 0 && jsonString[len(jsonString)-1] == ',' {
		newJsonString = jsonString[:len(jsonString)-1]
	} else {
		newJsonString = jsonString
	}
	err := json.Unmarshal([]byte(newJsonString), &obj)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}
	return obj.Location
}
