package utils

import (
	"strings"
	"testing"
)

func TestGetFileNameWithEpisodeName(t *testing.T) {
	filename := "episodename.mp3"
	podcast := Podcast{
		URL:      "https://example.com/podcasts/" + filename,
		Title:    "Episode 03",
		Filesize: 4 * 1024 * 1024,
	}
	expected := podcast.Title + "-" + filename

	result, _ := GetFileName(podcast)

	if expected != result {
		t.Error("File name is not matching!")
	}
}

func TestGetFileNameWithSlashesInEpisodeName(t *testing.T) {
	filename := "episodename.mp3"
	podcast := Podcast{
		URL:      "https://example.com/podcasts/" + filename,
		Title:    "Episode 03 / Topic",
		Filesize: 4 * 1024 * 1024,
	}
	expected := strings.Replace(podcast.Title, "/", "", -1) + "-" + filename

	result, _ := GetFileName(podcast)

	if expected != result {
		t.Error("File name is not matching!")
	}
}

func TestItReversesASlice(t *testing.T) {
	firstPodcast := Podcast{
		URL:      "https://example.com/podcasts1/",
		Title:    "Episode 01",
		Filesize: 4 * 1024 * 1024,
	}

	secondPodcast := Podcast{
		URL:      "https://example.com/podcasts2/",
		Title:    "Episode 02",
		Filesize: 4 * 1024 * 1024,
	}

	thirdPodcast := Podcast{
		URL:      "https://example.com/podcasts3/",
		Title:    "Episode 03",
		Filesize: 4 * 1024 * 1024,
	}
	testSlice := []Podcast{firstPodcast, secondPodcast, thirdPodcast}

	expected := []Podcast{thirdPodcast, secondPodcast, firstPodcast}

	result := ReversePodcastsToDownload(testSlice)

	if expected[0].URL != result[0].URL {
		t.Error("Could not reverse the slice!")
	}
}
