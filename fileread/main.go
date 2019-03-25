package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	// fileByte, err := ioutil.ReadFile("/Users/ricardo/Desktop/WX20180918-185926.png")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(len(fileByte))
	// err = ioutil.WriteFile("/Users/ricardo/Desktop/WXTmp.png", fileByte, 0666)
	// if err != nil {
	// 	panic(err)
	// }

	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// for i := 0; i < len(sample); i++ {
	// 	fmt.Printf("%x ", sample[i])
	// }
	// fmt.Printf("%+q\n", sample)

	// const placeOfInterest = `⌘`
	// fmt.Printf("plain string: ")
	// fmt.Printf("%s", placeOfInterest)
	// fmt.Printf("\n")
	// fmt.Printf("quoted string: ")
	// fmt.Printf("%+q", placeOfInterest)
	// fmt.Printf("\n")
	// fmt.Printf("hex bytes: ")
	// for i := 0; i < len(placeOfInterest); i++ {
	// 	fmt.Printf("%x ", placeOfInterest[i])
	// }

	// reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	// line, _ := reader.ReadBytes('\n')
	// fmt.Printf("the line:%s\n", line)
	// // 这里可以换上任意的 bufio 的 Read/Write 操作
	// n, _ := reader.ReadSlice('\n')
	// fmt.Printf("the line:%s\n", line)
	// fmt.Println(string(n))

	// reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	// line, err := reader.ReadSlice('\n')
	// if err == bufio.ErrBufferFull {
	// 	fmt.Printf("line:%s\terror:%s\n", line, err)
	// }
	// line, err = reader.ReadSlice('\n')
	// fmt.Printf("line:%s\terror:%s\n", line, err)

	// file, err := os.OpenFile("./l.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// s := "lala"
	// fmt.Println(string(strconv.AppendQuote([]byte(s), "ricardo")))
	// fmt.Println(strconv.Quote(s))

	// bufFile := bufio.NewReader(file)
	// for {
	// 	line, err := bufFile.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else {
	// 		fmt.Println(strings.TrimRight(line, "\n"))
	// 	}
	// }

	// buf := bufio.NewWriter(file)
	// for i := 0; i < 1000000; i++ {
	// 	s := rand.Intn(10000)
	// 	// fmt.Println(s)
	// 	// buf.WriteString(string(s))
	// 	// io.WriteString(file, "测试测试")
	// 	// err := ioutil.WriteFile("./l.txt", []byte("测试测试"), 0666)
	// 	// checkErr(err)
	// 	writeToFile(file, strconv.Itoa(s)+"\n")
	// 	// writeWithIo("./l.txt", "ceshi")
	// }

	s := `Surge port:
	6153问题
	找到gradle.properties里的gradle版本
	解压下载的对应版本的gradle到	/Users/ricardo/.gradle/wrapper/dists/ 下面，
	
	gradle/wrapper/gradle-wrapper.properties 文件的distributionUrl改为
	distributionUrl=file\:/Users/ricardo/.gradle/wrapper/dists/gradle-3.3-all/55gk2rcmfc6p2dg9u9ohc3hw9/gradle-3.3-all.zip
	
	// Presences>Gradle>Offline Work
	// /Users/ricardo/.gradle/wrapper/dists/gradle-3.3-all/55gk2rcmfc6p2dg9u9ohc3hw9
	// 下面的是备用选项，先用上面的
	// Presences>Gradle>Use local cradle distribution
	// Gradle home: /Users/ricardo/.gradle/wrapper/dists/gradle-3.3 (不带all的)`

	n := []byte(s)
	for i, w := 0, 0; i < len(n); i += w {
		runeValue, width := utf8.DecodeRune(n[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

}

func writeWithIo(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.", content)
	}
}

func writeToFile(f *os.File, content string) {
	if _, err := io.WriteString(f, content); err != nil {
		panic(err)
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
