/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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

func clearDirectory(dir string) error {
	// Read directory contents
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// Remove each item in the directory
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		// If it's a directory, we can use RemoveAll to recursively delete it
		// If it's a file, RemoveAll still works (it's equivalent to Remove for files)
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("failed to remove %s: %w", path, err)
		}
	}

	return nil
}

func OverrideXbyY(x, y string) error {
	fmt.Printf("overriding %s \nby %s \n", x, y)
	err := clearDirectory(x)
	if err != nil {
		return err
	}
	// err = os.Mkdir(x, os.ModeDir)
	// if err != nil {
	// 	return err
	// }
	dir := os.DirFS(y)
	return os.CopyFS(x, dir)
}

func handleRow(linux, win bool, row []string) error {
	var err error
	if win {
		err = BackupWindows(row)
		if err != nil {
			return err
		}
		err = OverrideXbyY(row[1], row[2])
		if err != nil {
			return err
		}
	}
	if linux {
		err = BackupLinux(row)
		if err != nil {
			return err
		}
		err = OverrideXbyY(row[2], row[1])
		if err != nil {
			return err
		}
	}
	return err
}

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use: "sync",

	Short: "sync one platfrom to match the other one",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		win, err := cmd.Flags().GetBool("windows")
		if err != nil {
			println(err.Error())
			return
		}
		linux, err := cmd.Flags().GetBool("linux")
		if err != nil {
			println(err.Error())
			return
		}
		game, err := cmd.Flags().GetString("name")
		if err != nil {
			println(err.Error())
			return
		}
		data := ReadData()
		for _, row := range data[1:] {
			if game == "" || row[0] != game {
				err = handleRow(linux, win, row)
				if err != nil {
					println(err.Error())
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().BoolP("windows", "w", false, "sync windows, pull the changes from linux")
	syncCmd.Flags().BoolP("linux", "l", false, "sync linux, pull the changes from windows")
	syncCmd.Flags().StringP("name", "n", "", "Choose the name of the entry to sync")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
