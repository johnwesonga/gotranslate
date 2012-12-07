package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	api     = "https://www.googleapis.com/language/translate/v2"
	API_KEY = "PU_API_KEY_HERE"
)

type InputText struct {
	PlainText      string
	TargetLanguage string
	Values         url.Values
}

type Translation struct {
	Data         string
	Translations []struct {
		TranslatedText string
		SourceLanguage string
	}
}

func (i *InputText) TranslateString() (*Translation, error) {
	if len(i.PlainText) == 0 {
		log.Fatal("No text specified")
	}
	if len(i.TargetLanguage) == 0 {
		log.Fatal("No target language specified")
	}

	i.Values = make(url.Values)
	var v = i.Values
	v.Set("target", i.TargetLanguage)
	v.Set("key", API_KEY)
	v.Set("q", i.PlainText)

	u := fmt.Sprintf("%s?%s", api, v.Encode())
	getResp, err := http.Get(u)
	if err != nil {
		log.Fatal("error", err)
		return nil, err
	}
	defer getResp.Body.Close()
	body, _ := ioutil.ReadAll(getResp.Body)
	t := new(Translation)
	json.Unmarshal(body, &t)

	return t, nil

}

func main() {
	input := &InputText{"My name is John, I was born in Nairobi and I am 31 years old", "ES", nil}
	translation, _ := input.TranslateString()
	fmt.Println(translation)
}
