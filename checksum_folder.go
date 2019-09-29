// makes a checksum of all files in a folder

package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Details struct {
	Path string
	Size int64
	Hash []byte
}

func md5File(d *Details) {
	f, err := os.Open(d.Path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	d.Hash = h.Sum(nil)
}

func findFiles(root string) []Details {
	var files []Details
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		itemDetails := Details{Path: path, Size: info.Size(), Hash: nil}
		files = append(files, itemDetails)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error walking path %q: %v\n", root, err)
	}
	return files
}

func main() {
	flag.Parse()
	folder := flag.Args()
	if len(folder) != 1 {
		fmt.Printf("must include one filepath argument\n")
		os.Exit(2)
	}
	filelist := findFiles(folder[0])
	for i := range filelist {
		md5File(&filelist[i])
		fmt.Println(filelist[i])
	}
}
