package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
	"strings"
	"path"
	"time"
)


type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	time.Sleep(100 * time.Millisecond)
	go func(){
		wc.PrintProgress()
	}()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %v complete", wc.Total)
}


//Test URL
//https://img1.akspic.ru/image/22392-park-oblako-gornyj_hrebet-gornyj_relef-nacionalnyj_park-3840x2160.jpg

func main() {
	var url, fname string
	fmt.Printf("Url: ")
	fmt.Fscan(os.Stdin, &url)
	fname = path.Base(url)
	fmt.Printf("File name : %s\n\n\n",fname)
	getFile(url, fname)
}

func getFile(path string, fname string){
	fmt.Printf("Starting...\n")
	file, err := os.Create(fname)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(path)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	counter := &WriteCounter{}
	rdr := io.TeeReader(resp.Body, counter)
	_, err = io.Copy(file, rdr);
	if err != nil {
		file.Close()
			panic(err)
	}

	fmt.Print("\nDownloaded")
	file.Close()
}
