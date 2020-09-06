package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Cat struct {
	Fact string `json:"fact"`
}

type Fact struct {
	Text string `json:"text"`
}

type Program struct {
	En     string `json:"en"`
	Author string `json:"author"`
}

type StarWars struct {
	Quote string `json:"starWarsQuote"`
}

type AnimeStatus struct {
	Data []Anime `json:"data"`
}

type Anime struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Anime     string `json:"anime"`
}

func GetCatFact(fact *Cat) {
	file, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(file.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	_ = json.Unmarshal(data, fact)
}

func GetProgrammingFact(programFact *Program) {
	response, err := http.Get("https://programming-quotes-api.herokuapp.com/quotes/random")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	_ = json.Unmarshal(data, programFact)
}

func GetNumberFact(fact *Fact) {
	response, err := http.Get("http://numbersapi.com/random?json")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	_ = json.Unmarshal(data, fact)
}

func GetRandomFact(fact *Fact) {
	response, err := http.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	_ = json.Unmarshal(data, fact)
}

func GetStarWarsQuote(fact *StarWars) {
	response, err := http.Get("http://swquotesapi.digitaljedi.dk/api/SWQuote/RandomStarWarsQuote")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	_ = json.Unmarshal(data, fact)
}

func GetAnimeQuotes() []byte {
	response, err := http.Get("https://animechanapi.xyz/api/quotes/random")
	if err != nil {
		fmt.Printf("Failed to fetch data from website: %v", err)
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data from website: %v", err)
		os.Exit(1)
	}
	return data
}
