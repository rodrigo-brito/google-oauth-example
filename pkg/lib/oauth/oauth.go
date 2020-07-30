package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func TokenSource(ctx context.Context, credentialsFile, tokenFile string, scopes ...string) (oauth2.TokenSource, error) {
	content, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %w", err)
	}

	config, err := google.ConfigFromJSON(content, scopes...)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	token, err := tokenFromFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("invalid tokenSource: %w", err)
	}

	return config.TokenSource(ctx, token), nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := new(oauth2.Token)
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}
