// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const urlToBeScrapped = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Printf("Starting to read a file from Internet\n")
	resp, err := http.Get(urlToBeScrapped)
	checkError(err)
	fmt.Printf("respons Type received is %T\n", resp)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	content := string(bytes)
	//fmt.Println("Content scrapped from the page is: ", content)

	tours := readToursFromRawData(content)
	for _, val := range tours {
		fmt.Println(val.Name)
	}
}

func readToursFromRawData(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Price, Name string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
