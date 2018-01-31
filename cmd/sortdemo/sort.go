package main

import (
	"os"
	"pipeline"
	"bufio"
	"fmt"
	"strconv"
	"time"
)

func main() {
	createNetworkPipeline("small.in", 128, 4)
	time.Sleep(time.Hour)
	/*writeToFile(p, "small.out")
	printFile("small.out")*/

	/*p := createPipeline("large.in", 8000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out")*/
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	pipeline.Init()

	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemorySort(source))
	}

	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	pipeline.Init()

	chunkSize := fileSize / chunkCount

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000 + i)
		pipeline.NetworkSink(addr, pipeline.InMemorySort(source))
		sortAddr = append(sortAddr, addr)
	}

	return nil

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		count ++
		fmt.Println(v)
		if count >= 100 {
			break
		}
	}
}