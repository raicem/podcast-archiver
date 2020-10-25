package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	utils "github.com/raicem/podcast-archiver/utils"
)

func main() {
	input := os.Args[1:]

	if len(input) == 0 {
		fmt.Println("Usage: podcast-archiver FEED")
		os.Exit(0)
	}

	feedURL := input[0]

	defaultProtocol := "http"

	// Prefix with the default protocol if one does not exists
	if !strings.HasPrefix(feedURL, "http") {
		feedURL = defaultProtocol + "://" + feedURL
	}

	_, err := url.ParseRequestURI(feedURL)

	if err != nil {
		fmt.Println("Please provide a valid URL.")
		os.Exit(1)
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var podcastsToDownload []utils.Podcast

	for _, item := range feed.Items {
		if len(item.Enclosures) > 0 {
			podcastsToDownload = append(podcastsToDownload, utils.ParseItem(item))
		}
	}

	fmt.Printf("%d episodes from %s will be downloaded\n", len(podcastsToDownload), feed.Title)

	for _, podcast := range podcastsToDownload {
		filename, err := utils.GetFileName(podcast)

		if err != nil {
			fmt.Println("Could not parse the URL. Skipping...")

			continue
		}

		_, err = os.Stat(filename)

		if err == nil {
			fmt.Printf("Episode %s is already downloaded. Skipping...\n", podcast.Title)

			continue
		}

		response, err := http.Get(podcast.URL)

		if err != nil {
			fmt.Printf("Could not download %s. Skipping...\n", podcast.Title)

			continue
		}

		defer response.Body.Close()

		// If the RSS feed does not have the filesize info, try to get it from HTTP header
		if podcast.Filesize == -1 {
			contentLength := response.Header.Get("content-length")
			contentLengthSize, err := strconv.Atoi(contentLength)

			if err != nil {
				contentLengthSize = -1
			}

			podcast.Filesize = contentLengthSize
		}

		out, err := os.Create(filename)

		if err != nil {
			fmt.Println("Could not create the file.")
			os.Exit(1)
		}

		defer out.Close()

		go func() {
			io.Copy(out, response.Body)
		}()

		fmt.Print("\033[s") // Save cursor position
		progress(filename, float32(podcast.Filesize))
	}
}

func progress(filename string, totalSize float32) {
	for {
		fileStats, err := os.Stat(filename)

		if err != nil {
			break
		}

		currentFileSize := float32(fileStats.Size())

		fmt.Print("\033[u\033[K") // Clear the current line
		if currentFileSize >= totalSize {
			printCompleted(filename)

			break
		}

		percentage := (currentFileSize / totalSize) * 100

		printProgress(totalSize, percentage, filename)

		time.Sleep(100 * time.Millisecond)
	}
}

func printProgress(totalSize float32, percentage float32, filename string) {
	totalSizeInMb := totalSize / (1024 * 1024)

	fmt.Printf("\rDownloading %s... %1.0f %% - Total: %1.2f MB", filename, percentage, totalSizeInMb)
}

func printCompleted(filename string) {
	fmt.Printf("\rCompleted downloading %s\n.", filename)
}
