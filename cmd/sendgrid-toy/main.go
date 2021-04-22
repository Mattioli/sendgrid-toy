package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	providers "github.com/Mattioli/sendgrid-toy/providers/email"
	"github.com/spf13/cobra"
)

var cmd *cobra.Command

func init() {
	cmd = &cobra.Command{
		Use:   "send",
		Short: "Send any email via CLI using sendgrid as a provider!",
		RunE:  run,
	}
}

func main() {
	cmd.Execute()
}

func run(cmd *cobra.Command, args []string) error {
	p := providers.NewSendGridProvider()

	jsonFile, err := os.Open("parameters.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	fmt.Println("Successfully Opened parameters.json")

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read content: %s\n", string(byteValue))

	err = p.Send(byteValue)
	if err != nil {
		fmt.Println("Error sending email!")
		return err
	}

	return nil
}
