/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// CompareFolderModTimes compares two paths (files or folders) and returns the path of the most recently modified one,
// along with its modification time and any error encountered.
// If both paths have the same modification time, the first path is returned.
func CompareFolderModTimes(path1, path2 string) (string, time.Time, error) {
	// Clean the paths to handle any . or .. components
	path1 = filepath.Clean(path1)
	path2 = filepath.Clean(path2)

	// Get file info for first path
	info1, err := os.Stat(path1)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error accessing path1 '%s': %w", path1, err)
	}

	// Get file info for second path
	info2, err := os.Stat(path2)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error accessing path2 '%s': %w", path2, err)
	}

	// Compare modification times
	modTime1 := info1.ModTime()
	modTime2 := info2.ModTime()

	if modTime2.After(modTime1) {
		return path2, modTime2, nil
	}
	return path1, modTime1, nil
}

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backup called")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
