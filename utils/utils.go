package utils

import (
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/mmcdole/gofeed"
)

// Podcast is a representation of a Podcast audio file.
type Podcast struct {
	URL      string
	Title    string
	Filesize int
}

// ParseItem returns a Podcast instance from goFeed.Item
func ParseItem(item *gofeed.Item) Podcast {
	firstEnclosure := item.Enclosures[0]

	// Get the filesize from the enclosure
	filesizeInfo := firstEnclosure.Length
	filesize, err := strconv.Atoi(filesizeInfo)

	if err != nil {
		filesize = -1
	}

	return Podcast{
		URL:      firstEnclosure.URL,
		Title:    item.Title,
		Filesize: filesize,
	}
}

// GetFileName Returns a filename for podcast audio file.
func GetFileName(podcast Podcast) (string, error) {
	parsedURL, err := url.Parse(podcast.URL)

	if err != nil {
		return "", err
	}

	filename := path.Base(parsedURL.Path)
	title := strings.Replace(podcast.Title, "/", "", -1)
	filename = title + "-" + filename

	return filename, nil
}

// ReversePodcastsToDownload Reverses the slice of podcast episodes.
func ReversePodcastsToDownload(podcastsToDownload []Podcast) []Podcast {
	var reversed []Podcast

	numberOfPodcastsToDownload := len(podcastsToDownload)

	for i := 0; i < numberOfPodcastsToDownload; i++ {
		reversed = append(reversed, podcastsToDownload[(numberOfPodcastsToDownload-1)-i])
	}

	return reversed
}
