package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func askValue(text string, allowedValues []string) string {
	var value string
	for {
		fmt.Println(text)
		_, err := fmt.Scanln(&value)
		if err != nil {
			log.Fatal(err)
		}
		if valueInSlice(value, allowedValues) {
			break
		}
		fmt.Println("Wrong option")
	}
	return value
}

func valueInSlice(val string, slice []string) bool {
	for _, value := range slice {
		if val == value {
			return true
		}
	}
	return false
}

func printSizes(files map[int64][]string) []int64 {
	var sizes []int64

	orderValues := []string{"1", "2"}

	sorting := askValue("Size sorting options:\n1. Descending\n2. Ascending", orderValues)

	for key := range files {
		sizes = append(sizes, key)
	}

	sort.Slice(sizes, func(i, j int) bool {
		if sorting == "1" {
			return sizes[i] > sizes[j]
		} else {
			return sizes[i] < sizes[j]
		}
	})

	for _, size := range sizes {
		fmt.Println(size, "bytes")
		for _, path := range files[size] {
			fmt.Println(path)
		}
	}

	return sizes
}

type fileInfo struct {
	size int64
	path string
}

func checkDuplicates(files map[int64][]string, sizes []int64) (ret []fileInfo) {
	answerValues := []string{"yes", "no"}
	answer := askValue("Check for duplicates?", answerValues)
	counter := 1
	if answer == "yes" {
		if len(sizes) > 0 {
			for _, size := range sizes {
				if len(files[size]) > 1 {
					var duplicates = make(map[string][]string)
					for _, path := range files[size] {
						file, err := os.Open(path)
						if err != nil {
							log.Fatal(err)
						}
						md5Hash := md5.New()
						if _, err := io.Copy(md5Hash, file); err != nil {
							log.Fatal(err)
						}
						hash := string(md5Hash.Sum(nil))
						duplicates[hash] = append(duplicates[hash], path)
						err = file.Close()
						if err != nil {
							log.Fatal(err)
						}
					}
					printBytes := true
					for hash, paths := range duplicates {
						if len(paths) > 1 {
							if printBytes {
								fmt.Println(size, "bytes")
								printBytes = false
							}
							fmt.Printf("Hash: %x\n", hash)
							for _, path := range paths {
								fmt.Printf("%v. %s\n", counter, path)
								ret = append(ret, fileInfo{size: size, path: path})
								counter += 1
							}
						}
					}
				}
			}
		}
	}
	return
}

func filesToDelete(files []fileInfo) []int {
	var indexes []int
	fmt.Print("Enter file numbers to delete:\n")
	reader := bufio.NewReader(os.Stdin)
	strNumbers, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) // Exit if we have an unexpected error
	}
	numbers := strings.Split(strings.TrimSpace(strNumbers), " ")
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil || number > len(files) {
			fmt.Println("Wrong format")
			return filesToDelete(files)
		}
		indexes = append(indexes, number)
	}
	return indexes
}

func deleteDuplicates(files []fileInfo) {
	answerValues := []string{"yes", "no"}
	answer := askValue("Delete files?", answerValues)
	var freedUp int64
	if answer == "yes" {
		indexes := filesToDelete(files)
		for _, index := range indexes {
			freedUp += files[index-1].size
			err := os.Remove(files[index-1].path)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Total freed up space:", freedUp, "bytes")
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		return
	}

	fmt.Println("Enter file format:")
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err) // Exit if we have an unexpected error
	}
	format := "." + strings.TrimSpace(string(b))

	var files = make(map[int64][]string)

	err = filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() && (format == "." || filepath.Ext(path) == format) {
			files[info.Size()] = append(files[info.Size()], path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	sizes := printSizes(files)
	duplicates := checkDuplicates(files, sizes)
	deleteDuplicates(duplicates)
}
