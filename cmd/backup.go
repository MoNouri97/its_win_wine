/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var BACKUP_PATH = ".its_win_wine/backup"

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

func BackupLinux(row []string) error {
	return BackupPath(row[2], row[0]+"_linux")
}

func BackupWindows(row []string) error {
	return BackupPath(row[1], row[0]+"_win")
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
				err = BackupWindows(row)
			}
			if linux {
				err = BackupLinux(row)
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

	backupCmd.Flags().BoolP("windows", "w", false, "create backups for the win paths")
	backupCmd.Flags().BoolP("linux", "l", false, "create backups for the linux paths")
}
