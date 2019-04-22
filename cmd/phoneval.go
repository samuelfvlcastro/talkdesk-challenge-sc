package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var acFilePath string
var pnFilePath string

var cmd = &cobra.Command{
	Version: "1.0",
	Use:     "phoneval",
	Short:   "Hello, this is a simple phone validation and localization CLI",
	Long: fmt.Sprint(`Welcome stranger!
	
This CLI was created to help phone number validation and localization using two txt files as inputs.
The first file must have one areacode per line and the second one phone number per line.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(acFilePath)
		fmt.Println(pnFilePath)

		areaCodes := retrieveLineData(acFilePath)
		phoneNumbers := retrieveLineData(pnFilePath)

		spew.Dump(areaCodes)
		spew.Dump(phoneNumbers)
	},
}

func retrieveLineData(path string) []string {
	lineData := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

func init() {
	cmd.Flags().StringVarP(&acFilePath, "areacodes", "a", "", "Path to the areacodes input file (required)")
	cmd.MarkFlagRequired("areacodes")

	cmd.Flags().StringVarP(&pnFilePath, "numbers", "n", "", "Path to the phone numbers input file (required)")
	cmd.MarkFlagRequired("numbers")
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
