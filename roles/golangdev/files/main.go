package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"

	// [START imports]
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	// [END imports]
	"flag"
	"os"
)

const api_key = "TRANSLATE_API_KEY"
const model = "nmt"

var src = flag.String("s", "en", "src language")
var target = flag.String("t", "ja", "")
var query = flag.String("q", "hello", "query")

func apikey() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	apikey := filepath.Join(user.HomeDir, ".honyaku", api_key)
	data, err := ioutil.ReadFile(apikey)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	flag.Parse()
	ctx := context.Background()
	apiKey, err := apikey()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	ops := translate.Options{Source: language.Make(*src), Format: translate.Text, Model: model}
	resp, err := client.Translate(ctx, []string{*query}, language.Make(*target), &ops)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	fmt.Printf("%s\n", resp[0].Text)
	os.Exit(1)
}
