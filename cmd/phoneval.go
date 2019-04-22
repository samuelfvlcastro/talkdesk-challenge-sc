package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/phoneval"
	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/usecases"
)

var acFilePath string
var pnFilePath string

var cmd = &cobra.Command{
	Version: "1.0",
	Use:     "phoneval",
	Short:   "phoneval validates phone numbers and groups them by area code",
	Long: fmt.Sprint(`Welcome stranger!
	
This CLI was created to help validate phone numbers and grouping them by area code using two txt files as inputs.
The first file must have one area code per line and the second one phone number per line.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		areaCodes := retrieveLineData(acFilePath)
		phoneNumbers := retrieveLineData(pnFilePath)

		validator := phoneval.NewValidator(areaCodes)
		phoneStats := usecases.NewPhoneStats(validator)

		acKeys, acStats := phoneStats.CalculateAreaCodeStatistics(phoneNumbers)
		renderStatistics(acKeys, acStats)
	},
}

func init() {
	cmd.Flags().StringVarP(&acFilePath, "areacodes", "a", "", "Path to the areacodes input file (required)")
	if err := cmd.MarkFlagRequired("areacodes"); err != nil {
		log.Fatal(err)
	}

	cmd.Flags().StringVarP(&pnFilePath, "numbers", "n", "", "Path to the phone numbers input file (required)")
	if err := cmd.MarkFlagRequired("numbers"); err != nil {
		log.Fatal(err)
	}
}

// Execute runs the current command
func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func retrieveLineData(path string) []string {
	lineData := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		lineData = append(lineData, strings.TrimSpace(line))
	}

	return lineData
}

func renderStatistics(acKeys []int, acStats map[int]int) {
	fmt.Println("Area Code Statistics:")
	for _, key := range acKeys {
		fmt.Printf("%d:%d\n", key, acStats[key])
	}
}
