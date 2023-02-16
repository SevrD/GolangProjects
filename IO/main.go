package main

// Напишите свои объекты, которые их реализуют и например отдают бесконечный/конечный поток случайных байт. проверять через ioutil.ReadAll() для конечного.

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

type myIO string

type inf struct{}

type endStream struct{}

func (st endStream) Read(p []byte) (n int, err error) {
	count := int(rand.Intn(255))
	for i := 0; i < count; i++ {
		p[i] = byte(int(rand.Intn(255)))
	}
	if count > 100 {
		return count, nil
	}
	return count, io.EOF
}

func (st inf) Read(p []byte) (n int, err error) {
	count := int(rand.Intn(255))
	for i := 0; i < count; i++ {
		p[i] = byte(int(rand.Intn(255)))
	}
	return count, nil
}

func (myStream myIO) Read(p []byte) (n int, err error) {

	for i, value := range myStream {
		p[i] = byte(value)
	}
	return len(myStream), io.EOF
}

func main() {
	test := myIO("1234asdf")
	buffer := make([]byte, len(test))
	test.Read(buffer)
	fmt.Println(buffer)

	var test2 inf
	buffer2 := make([]byte, 255)
	breaker := 0
	for {
		breaker++
		count, err := test2.Read(buffer2)
		fmt.Println(buffer2[:count])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		if breaker >= 1000 {
			break
		}
	}
	fmt.Println(breaker)

	var test3 endStream
	var buffer3 []byte
	buffer3, _ = io.ReadAll(test3)
	fmt.Println(buffer3)
	fmt.Println(len(buffer3))
}
