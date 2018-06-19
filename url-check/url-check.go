// Description: url-check is a minimal tool to inspect remote urls
// Author: Giovan Battista Salinetti (gbsalinetti@extraordy.com)

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Error: missing remote url.")
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<REMOTE_URL>")
		os.Exit(1)
	}

	// If more flags need to be added consider using the flag or pflag packages
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<REMOTE_URL>")
		os.Exit(0)
	}

	// Parse the url
	testURL := os.Args[1]
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("URL Parsing section:\nScheme: %s\nHost: %s\nPath: %s\n\n", u.Scheme, u.Host, u.Path)

	// Configure insecure transport and proxies
	// Since http.ProxyFromEnvironment is used the HTTP_PROXY, HTTPS_PROXY and NO_PROXY env variables
	// are expected to work and provide proxy urls and exclusions to the client.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
	}
	client := &http.Client{Transport: tr}

	// Connect to the remote host
	resp, err := client.Get(testURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print Status and Protocol along with the response body
	fmt.Printf("Response section:\nResponse: %s\nProtocol: %s\n", resp.Status, resp.Proto)
	fmt.Printf("%s", string(body))
}
