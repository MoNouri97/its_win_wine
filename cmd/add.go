/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new sync entry",
	Long:  `Add a new sync entry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		name := args[0]
		win := args[1]
		linux := args[2]
		f, e := os.Create("./list.csv")
		if e != nil {
			fmt.Println(e)
			return
		}
		writer := csv.NewWriter(f)
		data := [][]string{
			{"Name", "Windows", "Wine_linux"},
			{name, win, linux},
		}

		e = writer.WriteAll(data)
		if e != nil {
			fmt.Println(e)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
