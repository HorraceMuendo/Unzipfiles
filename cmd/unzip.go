/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"os"
	util "unzipFiles/Util"

	"github.com/spf13/cobra"
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "",
	Long:  `.........`,

	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		util.UnzipFile(&zip.File{}, wd)
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unzipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
