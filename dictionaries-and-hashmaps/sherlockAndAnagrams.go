package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	cCount := make(map[rune]int)
	for _, r := range s1 {
		if _, ok := cCount[r]; !ok {
			cCount[r] = 1
		} else {
			cCount[r] = cCount[r] + 1
		}
	}

	for _, r := range s2 {
		if c, ok := cCount[r]; !ok || c == 0 {
			return false
		}

		cCount[r] = cCount[r] - 1
	}

	return true
}

func main() {
	fmt.Println(sherlockAndAnagrams("aba"))
	fmt.Println(sherlockAndAnagrams("abcdefghijklmnopqrstuvwxyz"))
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	length := len(s)
	if length <= 1 {
		return 0
	}

	var count int32
	for l := 1; l <= length-1; l++ {
		for i := 0; i+l <= length; i++ {
			for j := i + 1; j+l <= length; j++ {
				if isAnagram(s[i:i+l], s[j:j+l]) {
					count++
					fmt.Printf("%s - %s", s[i:i+l], s[j:j+l])
					fmt.Println()
				}
			}
		}
	}

	return count
}

func main1() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
