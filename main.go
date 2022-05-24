package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"stringsearch/other/progressbar"
)

func iterate(rootpath string) []string {
	returnList := []string{}

	filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if info.IsDir() {
			return nil
		}
		returnList = append(returnList, path)
		// returnDir = "\nFile Name:" + info.Name()
		return nil
	})
	return returnList
}

func main() {
	welcomeMessage := `
:'######::'########:'########::'####:'##::: ##::'######:::::::'######::'########::::'###::::'########:::'######::'##::::'##:
'##... ##:... ##..:: ##.... ##:. ##:: ###:: ##:'##... ##:::::'##... ##: ##.....::::'## ##::: ##.... ##:'##... ##: ##:::: ##:
 ##:::..::::: ##:::: ##:::: ##:: ##:: ####: ##: ##:::..:::::: ##:::..:: ##::::::::'##:. ##:: ##:::: ##: ##:::..:: ##:::: ##:
. ######::::: ##:::: ########::: ##:: ## ## ##: ##::'####::::. ######:: ######:::'##:::. ##: ########:: ##::::::: #########:
:..... ##:::: ##:::: ##.. ##:::: ##:: ##. ####: ##::: ##::::::..... ##: ##...:::: #########: ##.. ##::: ##::::::: ##.... ##:
'##::: ##:::: ##:::: ##::. ##::: ##:: ##:. ###: ##::: ##:::::'##::: ##: ##::::::: ##.... ##: ##::. ##:: ##::: ##: ##:::: ##:
. ######::::: ##:::: ##:::. ##:'####: ##::. ##:. ######::::::. ######:: ########: ##:::: ##: ##:::. ##:. ######:: ##:::: ##:
:......::::::..:::::..:::::..::....::..::::..:::......::::::::......:::........::..:::::..::..:::::..:::......:::..:::::..::
This software is an easy and quick way to search for strings inside files on your (unix based) filesystem.

`
	fmt.Printf("%s", welcomeMessage)

	searchStringReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your search string: ")
	searchString, _ := searchStringReader.ReadString('\n')
	fileNameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Pleas enter your output file name: ")
	fileName, _ := fileNameReader.ReadString('\n')
	searchDirectoryReader := bufio.NewReader(os.Stdin)
	fmt.Print("Pleas enter your search directory root: ")
	searchDir, _ := searchDirectoryReader.ReadString('\n')
	outputFileName := strings.Split(strings.Split(fileName, "\n")[0], ".")[0] + ".txt"
	// currentDirectory := "/Users/martijn/Documents/BreachCompilation"
	fileCount := 0

	var list = iterate(strings.Split(searchDir, "\n")[0])
	totalNumberOfFiles := int64(len(list))
	var bar progressbar.Bar
	bar.NewOption(0, totalNumberOfFiles)

	f, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range list {
		fileCount++
		bar.Play(int64(fileCount))

		b, err := ioutil.ReadFile(dir)
		if err != nil {
			panic(err)
		}
		s := string(b)
		if strings.Contains(s, strings.Split(searchString, "\n")[0]) {
			lines := strings.Split(s, "\n")
			for _, line := range lines {
				if strings.Contains(line, strings.Split(searchString, "\n")[0]) {
					// Write line to file
					_, err = fmt.Fprintln(f, line)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
	bar.Finish()
}
