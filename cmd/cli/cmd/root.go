package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "secretify",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("BLA")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Should not be here though

type viewRes struct {
	Data struct {
		Cipher        string    `json:"cipher"`
		ExpiresAt     time.Time `json:"expires_at"`
		HasPassphrase bool      `json:"has_passphrase"`
		RevealOnce    bool      `json:"reveal_once"`
		Deleted       bool      `json:"deleted"`
	} `json:"data"`
}

func get(c *http.Client, url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var t viewRes
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	return t.Data.Cipher, nil
}

type createReq struct {
	Cipher        string `json:"cipher"`
	ExpiresAt     string `json:"expires_at"`
	RevealOnce    bool   `json:"reveal_once"`
	HasPassphrase bool   `json:"has_passphrase"`
	Email         string `json:"email"`
	WebhookAddr   string `json:"webhook_addr"`
}

type createRes struct {
	Data struct {
		CreatedAt  time.Time `json:"created_at"`
		Identifier string    `json:"identifier"`
	} `json:"data"`
}

func create(c *http.Client, url string, input []byte) (string, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(input))
	if err != nil {
		return "", err
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s: %s", resp.Status, string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var t createRes
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	return t.Data.Identifier, nil
}

func client() *http.Client {
	var defaultTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          30,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	c := &http.Client{
		Transport: defaultTransport,
	}
	return c
}
