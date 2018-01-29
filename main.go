package main

import (
	"awesomeProject/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename = "large.in"
	const number = 10000000
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

	p = pipeline.ReaderSource(bufio.NewReader(file))
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if (count > 50) {
			break
		}
	}
}

func mergeSort() {
	p := pipeline.Merge(
		pipeline.InMemorySort(
			pipeline.ArraySouce(2, 4, 5, 1, 10)),
		pipeline.InMemorySort(
			pipeline.ArraySouce(6, 0, 3, 16, 9)))
	for v := range p {
		fmt.Println(v);
	}
}
