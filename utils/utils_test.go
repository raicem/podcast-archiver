package utils

import (
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
