/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	util "unzipFiles/Util"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "unzipFiles <zip_file_name/location>",
	DisableFlagsInUseLine: true,
	Short:                 "An application that helps users unzip files",
	Version:               version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Long: `Unzip files is an application that is solely built to help users unzip files `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string
		//take the first argument passed and store it in argument
		argument = args[0]
		//check whether the file exists
		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}
		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("file %v does not exist ", argument)
		}
		//get the current working dir
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		//unzip the file and store it in current working dir
		util.UnzipSource(fileName, wd)
		//remove the extention (.zip) from the file name
		os.Chdir(util.FileNameWithoutExtension(fileName))
		//update working dir
		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.unzipFiles.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
