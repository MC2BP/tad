package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const LibraryFolder = "Music"

type Library struct {
	Genres map[string]map[string]Song `json:"genres"`
}

type Song struct {
	Artists []string `json:"artists"`
	Links   []string `json:"links"`
}

func main() {

	//Read file
	file, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	var library Library
	err = json.Unmarshal(file, &library)
	if err != nil {
		panic(err)
	}
	fmt.Println(library)

	//Check directory music
	err = createFolderIfNotExist(LibraryFolder)
	if err != nil {
		panic(err)
	}

	for genreName, genre := range library.Genres {
		createFolderIfNotExist(LibraryFolder + "/" + genreName)
		if err != nil {
			panic(err)
		}
		for songName, song := range genre {
			artists := strings.Join(song.Artists, ", ")
			songName = artists + " - " + songName + ""
			path := LibraryFolder + "/" + genreName + "/" + songName + ".%(ext)s"
			// check if song already present
			_, err := os.Stat(path)
			if err == nil {
				fmt.Println("file already present", err)
				continue
			}

			fmt.Println("Downloading", path)
			for _, url := range song.Links {
				err = exec.Command(
					"youtube-dl",
					"-x",
					"--audio-format",
					"m4a",
					"--prefer-ffmpeg",
					url,
					"-o",
					path,
				).Run()
				if err != nil {
					continue
				}

			}
		}
	}

}

func createFolderIfNotExist(path string) error {
	folderInfo, err := os.Stat(path)
	if err != nil || !folderInfo.IsDir() {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// Command to download:
// yoltube-dl -x --audio-format m4a --prefer-ffmpeg <URL> -o "<ARTIST> - <TITLE>.%(ext)s"
