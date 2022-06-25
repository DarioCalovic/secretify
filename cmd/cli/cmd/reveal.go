package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/DarioCalovic/secretify/pkg/crypto"

	"github.com/spf13/cobra"
)

func init() {
	revealCmd.Flags().StringVarP(&url2, "url", "u", "https://www.secretify.io/api/v1", "URL to a secretify api endpoint")
	revealCmd.Flags().StringVarP(&identifier, "identifier", "i", "", "Secretify identifier")
	revealCmd.MarkFlagRequired("url")
	revealCmd.MarkFlagRequired("identifier")
	rootCmd.AddCommand(revealCmd)
}

var (
	url2       string
	identifier string
)

var revealCmd = &cobra.Command{
	Use:   "reveal",
	Short: "Reveals a secret",
	Long:  `Reveals a secret link`,
	Run: func(cmd *cobra.Command, args []string) {

		// Retrieve id and key
		s := strings.Split(identifier, "#")
		if len(s) != 2 {
			fmt.Println(errors.New("invalid identifier"))
			return
		}

		// Get cipher
		cipher, err := get(client(), fmt.Sprintf("%s/secret/%s/_cipher", url2, s[0]))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Encrypt
		plainttext, err := crypto.Decrypt(cipher, s[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(plainttext)
	},
}
