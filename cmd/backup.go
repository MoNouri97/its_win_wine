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

var BACKUP_PATH = ".its_win_wine/backup"

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

// backup a folder
func BackupPath(path, name string) error {
	err := os.MkdirAll(BACKUP_PATH, os.ModeDir)
	if err != nil {
		return err
	}
	newPath := BACKUP_PATH + "/" + name + "_" + time.Now().Format("2006_01_02_15_04")
	os.CopyFS(newPath, os.DirFS(path))
	fmt.Println("Created Backup for " + name + " at")
	fmt.Println(newPath)

	return nil
}

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long:  `Create a backup for the linux and/or the windows path`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(BACKUP_PATH, os.ModeDir)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
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
		data := ReadData()

		for _, row := range data[1:] {
			if win {
				err = BackupPath(row[1], row[0]+"_win")
			}
			if linux {
				err = BackupPath(row[2], row[0]+"_linux")
			}
			//    s, _, err := CompareFolderModTimes(row[1], row[2])
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			// if s == row[1] {
			// 	fmt.Println("the newer file is win")
			// }
			// if s == row[2] {
			// 	fmt.Println("the newer file is linux")
			// }
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().BoolP("windows", "w", false, "the windows path to sync")
	backupCmd.Flags().BoolP("linux", "l", false, "the linux/wine path to sync")
}
