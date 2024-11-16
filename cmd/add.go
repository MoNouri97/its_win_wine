/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var DataFile = "./list.csv"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new sync entry",
	Example: "its_win_wine add NAME [flags]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires NAME arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		win, err := cmd.Flags().GetString("windows")
		linux, err := cmd.Flags().GetString("linux")
		if err != nil {
			println(err.Error())
			return
		}
		newFile := false
		f, err := os.OpenFile(DataFile, os.O_RDWR, os.ModeAppend)

		if err != nil && os.IsNotExist(err) {
			f, err = os.Create("./list.csv")
			newFile = true
		}

		if err != nil {
			println(err.Error())
			return
		}
		defer f.Close()

		prev := csv.NewReader(f)

		prevData, err := prev.ReadAll()
		if err != nil {
			println(err.Error())
			return
		}
		newFile = len(prevData) == 0
		writer := csv.NewWriter(f)
		var data [][]string
		if newFile {
			data = append(data,
				[]string{"Name", "Windows", "Wine_linux"},
			)
		}

		data = append(data,
			[]string{name, win, linux},
		)

		err = writer.WriteAll(data)
		if err != nil {
			println(err.Error())
			return
		}
		println("DONE ! " + "added " + name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("windows", "w", "", "the windows path to sync")
	addCmd.Flags().StringP("linux", "l", "", "the linux/wine path to sync")
	addCmd.MarkFlagRequired("windows")
	addCmd.MarkFlagRequired("linux")
}
