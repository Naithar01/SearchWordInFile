package no_r_t

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileLine struct {
	text string
	line int
}

type FileInfo struct {
	filename string
	lines    []FileLine
}

// 입력된 파일 경로에 실제 파일이 존재하는지 확인해주는 함수
func GetFilePath(path string) ([]string, error) {
	return filepath.Glob(path)
}

func CheckWord(word string) {
	fmt.Println("-----------")
	fmt.Println("단어", word)
	fmt.Println("-----------")
}

func CheckOsArgsLen() bool {
	if len(os.Args) < 3 {
		fmt.Println("3개 이상의 인수가 필요 * 실행파일, 단어, 파일")
		return false
	}

	return true
}

func FindWordInFile(word, filename string) FileInfo {
	fileinfo := FileInfo{filename: filename, lines: []FileLine{}}

	file, err := os.Open(filename)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	Line := 0

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, word) {
			fileinfo.lines = append(fileinfo.lines, FileLine{text, Line + 1})
		}

		Line++
	}

	return fileinfo
}

func FindWordInFiles(word, path string) []FileInfo {
	file, err := GetFilePath(path) // 파일 경로가 제대로 있는지 확인

	if err != nil {
		panic(err.Error())
	}

	fileInfo := []FileInfo{}

	// 파일이 존재하면
	if len(file) > 0 {
		for _, filename := range file {
			fileInfo = append(fileInfo, FindWordInFile(word, filename))

			return fileInfo
		}

	}

	return fileInfo

}
func Fo() {
	if CheckOsArgsLen() {
		word := os.Args[1] // 단어
		CheckWord(word)

		files := os.Args[2:] // 파일 리스트

		fileInfos := []FileInfo{}

		for _, path := range files {
			fileInfos = append(fileInfos, FindWordInFiles(word, path)...)
		}

		for _, file := range fileInfos {
			fmt.Println("file: ", file.filename)
			for _, line := range file.lines {
				fmt.Println(line.line, ":", line.text)
			}

			fmt.Println("------------")
		}
	}
}
