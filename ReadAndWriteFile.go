// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("Starting to read a file with charachters\n")
	readGivenFile("./adityaOutputFile.txt")
}

func readGivenFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Data being read from the file is usually BYTES: ", data)
	fmt.Println("Converted data in String format is: ", string(data))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Writenewfile() {
	content := "Hello from GO !!"
	file, err := os.Create("./adityaOutputFile.txt")
	checkError(err)
	// length here indicates the Number of charachters being written to the FILE.
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v charachters\n", length)
	defer file.Close()
}
