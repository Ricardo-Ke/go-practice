package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

// func getCurrentDirectory() string {
// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return strings.Replace(dir, "\\", "/", -1)
// }

func main() {

	AppFs := afero.NewMemMapFs()

	AppFs.Mkdir("/Users/ricardo/go/src/github.com/Ricardo-Ke/go-practice/dirScan/dependency", os.ModeDir)

	if ok, _ := afero.DirExists(AppFs, "/Users/ricardo/go/src/github.com/Ricardo-Ke/go-practice/dirScan/dependency"); ok == true {
		fmt.Println("deps exists")
	}

	// str1 := getCurrentDirectory()
	// path := "./"
	// pattern := ".pdb"

	// re, _ := regexp.Compile(pattern)

	// dirList, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	fmt.Println("files not found")
	// }

	// outputFile, outputError := os.OpenFile("sum.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// if outputError != nil {
	// 	fmt.Println("Error with file opening")
	// 	return
	// }

	// for _, v := range dirList {
	// 	if v.Name() == "main" {
	// 		continue
	// 	}
	// 	fmt.Fprintf(outputFile, re.ReplaceAllString(v.Name(), "")+"\r\n")
	// }

}
