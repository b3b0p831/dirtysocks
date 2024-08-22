package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var printCol int = 20
var BUF_SIZE int = 2048

var rootCmd = &cobra.Command{
	Use: "dirtysocks",
	Short: "A simple data visualizer for your dirty socks",
	Long:  `[+]------------------Dan's Dirty Socks------------------[+]`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
 		if len(args) == 0 {
        		cmd.Help()  // This will print the help menu
        		os.Exit(0)  // Exit the program after printing help
    		}
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func pretyPrintBuf(sock_data []byte, n_wide int) {
	last_index := 0
	for last_index < len(sock_data) {
		for i := last_index; i < last_index+n_wide && i < len(sock_data); i++ {
			fmt.Printf("%02x ", sock_data[i])
		}

		leftOver := last_index + n_wide - len(sock_data)
		for i := 0; i < leftOver; i++ {
			fmt.Print("   ")
		}

		for i := last_index; i < last_index+n_wide && i < len(sock_data); i++ {
			chr := sock_data[i]
			switch chr {
			case ' ':
				fmt.Print(" _ ")
				continue

			case '\x0a':
				fmt.Printf("\\n ")
				continue

			case '\x09':
				fmt.Printf("\\t ")
				continue

			}

			fmt.Printf(" %s ", string(chr))
		}

		fmt.Println()
		last_index = last_index + n_wide

	}

}
