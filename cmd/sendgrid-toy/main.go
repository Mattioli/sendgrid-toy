package main

import (
	"fmt"

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
	if len(args) != 7 {
		fmt.Println("You need to pass 6 parameters, in this order: [fromName] [fromEmail] [toName] [toEmail] [subject] [msg]")
		return fmt.Errorf("wrong number of parameters: expected 6, received %d", len(args)-1)
	}

	fromName := args[1]
	fromEmail := args[2]
	toName := args[3]
	toEmail := args[4]
	subject := args[5]
	msg := args[6]

	p := initEmailProviders()

	for _, v := range p {
		err := v.Send(fromName, fromEmail, toName, toEmail, subject, msg)
		if err != nil {
			fmt.Println("Error sending email!")
			return err
		}
	}

	return nil
}

func initEmailProviders() []providers.EmailProvider {
	result := []providers.EmailProvider{providers.NewSendGridProvider()}
	return result
}
