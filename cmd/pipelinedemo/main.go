package main

import (
	"pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	mainSort()
	//mergeSort()
}

func mainSort() {
	/*const filename = "large.in"
	const number = 10000000*/

	const filename = "small.in"
	const number = 16
	file, error := os.Create(filename)
	if error != nil {
		panic(error)
	}

	defer file.Close()

	p := pipeline.RandomSource(number)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, error = os.Open(filename)
	if error != nil {
		panic(error)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count > 50 {
			break
		}
	}
}

func mergeSort() {
	p := pipeline.Merge(
		pipeline.InMemorySort(
			pipeline.ArraySouce(7273596521315663110, 4, 8249030965139585917, 1, 10)),
		pipeline.InMemorySort(
			pipeline.ArraySouce(9010467728050264449, 0, 3, 2050257992909156333, 9)))
	for v := range p {
		fmt.Println(v);
	}
}