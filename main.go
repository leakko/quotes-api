package main

import (
	// "encoding/json"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

type Quote struct {
	quoteText string
	quoteAuthor string
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

func openFile(fileName string) (*os.File, error) {
	// Open our jsonFile
	file, err := os.Open(fileName+".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our file till the function finishes
	defer file.Close()
	// else, return the file and set the error to its zero value
	return file, nil
}


func main() {
	
	file, _ := openFile("quotes")
	
	// read our opened jsonFile as a byte array.
	fileBytes, _ := ioutil.ReadAll(file)
	
	// we initialize our Users array
	var quotes Quotes

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(fileBytes, &quotes)

	// we iterate through every quote within our quotes array and
	// print out the quote author and text.
	for i := 0; i < 20; i++ {
		fmt.Println("Author: " + quotes.Quotes[i].quoteAuthor)
		fmt.Println("Quote: " + quotes.Quotes[i].quoteText)
	}
}