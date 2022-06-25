package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/DarioCalovic/secretify/pkg/crypto"
	"github.com/DarioCalovic/secretify/pkg/util/nanoid"

	"github.com/spf13/cobra"
)

func init() {
	createCmd.Flags().StringVarP(&url, "url", "u", "https://www.secretify.io/api/v1", "URL to a secretify api endpoint")
	createCmd.Flags().StringVarP(&secretText, "secret", "s", "", "Secret text to be encrypted")
	createCmd.MarkFlagRequired("url")
	createCmd.MarkFlagRequired("secret")
	rootCmd.AddCommand(createCmd)
}

var (
	url        string
	secretText string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a secret",
	Long:  `Creates a secret link waiting to be shared`,
	Run: func(cmd *cobra.Command, args []string) {
		// Generate key
		key, err := nanoid.GenerateIdentifier(26)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Encrypt
		cipher, err := crypto.Encrypt([]byte(secretText), key)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create secret link
		b, err := json.Marshal(createReq{
			ExpiresAt: "1h",
			Cipher:    cipher,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		secretLink, err := create(client(), fmt.Sprintf("%s/secret", url), b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s#%s", secretLink, key)
	},
}
