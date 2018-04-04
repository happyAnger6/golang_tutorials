package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sem = make(chan struct{}, 100)

func dirents(dir string) []os.FileInfo {
	sem <- struct{}{}
	defer func() { <-sem }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("readdir:%s err:%v\n", dir, err)
		return nil
	}

	return entries
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var n sync.WaitGroup
	fileSizes := make(chan int64)
	start := time.Now()
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nFiles, nBytes int64
	for size := range fileSizes {
		nFiles++
		nBytes += size
	}

	fmt.Printf("files:%d nBytes:%d time:%v\n", nFiles, nBytes, time.Now().Sub(start))
}
