/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Rename a movie file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, _ := regexp.Compile(`^.*\d\d\d\d\.`)

		for i := 0; i < len(args); i++ {
			mName := args[i]
			fmt.Println("I made it here and the movie is:", mName)

			fmt.Println("")
			fmt.Println("Movie Name:", mName)
			idx := r.FindStringIndex(mName)
			ogMatch := strings.Split(mName[idx[0]:idx[1]-1], ".")
			match := strings.Join(ogMatch, " ")

			newTitle := fmt.Sprintf("%s(%s).%s", match[0:len(match)-4], match[len(match)-4:], mName[len(mName)-3:])
			var answer string
			fmt.Println("I made it here 2")

			fmt.Println("The movie will be renamed to:", newTitle)
			fmt.Print("y to accept, n to skip: ")
			fmt.Scan(&answer)
			if answer == "y" {
				os.Rename(mName, newTitle)
				fmt.Println("The file has been renamed!")
			} else {
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
