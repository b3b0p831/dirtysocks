package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp <IP> <PORT>",
	Short: "Open and visualze socket data",
	Run: func(cmd *cobra.Command, args []string) {
		// Get ip and port
		ip := args[0]
		port := args[1]

		//Connect to ip:port
		con, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
		if err != nil {
			fmt.Println("Failed to connect :(")
			os.Exit(-1)
		}
		fmt.Printf("Successfully connected to %v\n", con.RemoteAddr())
		fmt.Println("Waiting for incoming data...")

		// Read data, BUF_SIZE, until no data left

		data, err := cmd.Flags().GetString("data")

		if err == nil && len(data) > 0 {
			data_bytes := []byte(data)
			data_bytes = append(data_bytes, '\x0d', '\x0a')
			c, e := con.Write(data_bytes)
			if e != nil {
				errorCleanup(e, con)
			}
			fmt.Println()
			pretyPrintBuf(data_bytes)

			fmt.Printf("Bytes written: %d\n\n", c)
		}

		buf := make([]byte, BUF_SIZE)
		total := []byte{}
		c, e := con.Read(buf)
		for c > 0 {
			if e != nil {
				errorCleanup(e, con)
			}

			//
			total = append(total, buf[:c]...)
			if c < BUF_SIZE {
				break
			}
			c, e = con.Read(buf)
		}

		pretyPrintBuf(total)
		fmt.Printf("Bytes read: %d\n\n", len(total))

		con.Close()
	},
}

func errorCleanup(e error, con net.Conn) {
	fmt.Println(e)
	con.Close()
	os.Exit(-1)
}

func init() {
	rootCmd.AddCommand(tcpCmd)
	tcpCmd.Flags().String("data", "", "Data to send before receiving data")

	tcpCmd.Args = cobra.ExactArgs(2)

}
