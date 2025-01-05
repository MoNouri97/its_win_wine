/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

func ReadData() [][]string {
	f, err := os.OpenFile(DataFile, os.O_RDWR, os.ModeAppend)
	if err != nil {
		println(err.Error())
		os.Exit(1)
		return nil
	}
	defer f.Close()

	prev := csv.NewReader(f)

	data, err := prev.ReadAll()
	if err != nil {
		println(err.Error())
		os.Exit(1)
		return nil
	}
	if len(data) == 0 {
		fmt.Println("No data found in CSV file")
		return nil
	}
	return data
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := ReadData()
		// Get the maximum width for each column
		columnWidths := make([]int, len(data[0]))
		for _, row := range data {
			for i, cell := range row {
				width := utf8.RuneCountInString(cell)
				if width > columnWidths[i] {
					columnWidths[i] = width
				}
			}
		}

		// Print the table
		printSeparator(columnWidths)

		// Print headers
		for i, header := range data[0] {
			fmt.Printf("| %-*s ", columnWidths[i], header)
		}
		fmt.Println("|")

		// printSeparator(columnWidths)

		// Print data rows
		for _, row := range data[1:] {
			for i, cell := range row {
				fmt.Printf("| %-*s ", columnWidths[i], prettyPath(cell))
			}
			fmt.Println("|")
		}

		printSeparator(columnWidths)
	},
}

func printSeparator(widths []int) {
	for _, w := range widths {
		fmt.Print("+", strings.Repeat("-", w+2))
	}
	fmt.Println("+")
}

func prettyPath(arr string) string {
	parts := strings.Split(strings.Trim(arr, "/"), "/")
	if len(parts) > 2 {
		short := parts[len(parts)-3:]
		return ".../" + strings.Join(short, "/")

	}
	return arr
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
