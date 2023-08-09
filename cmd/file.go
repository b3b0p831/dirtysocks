package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file <file>",
	Short: "Open and visualze file data",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		var fileName string = args[0]

		//Open file
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		defer f.Close()
		fmt.Printf("Opening %s\n", fileName)

		// Read data until no data read
		fileData, e := io.ReadAll(f)
		if e != nil {
			fmt.Println(e)
			os.Exit(-1)
		}

		pretyPrintBuf(fileData, printCol)
		fmt.Printf("\nBytes read: %d\n\n", len(fileData))
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	fileCmd.Args = cobra.ExactArgs(1)
}
