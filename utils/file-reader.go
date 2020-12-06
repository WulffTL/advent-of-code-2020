package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetArrayFromFile(path string) []int64 {
	root := GetRootPath()
	buf, err := os.Open(root + "/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)

	var returnValue []int64
	for snl.Scan() {
		val := snl.Text()
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		returnValue = append(returnValue, num)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}

	return returnValue
}

// GetArrayOfStringsFromFile returns an array of strings from file
// input file separated by new lines
func GetArrayOfStringsFromFile(path string) []string {
	root := GetRootPath()
	buf, err := os.Open(root + "/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)

	var returnValue []string
	for snl.Scan() {
		val := snl.Text()
		if err != nil {
			log.Fatal(err)
		}
		returnValue = append(returnValue, val)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}

	return returnValue
}
