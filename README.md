# podcast-archiver

A CLI tool to download all podcast episodes in a feed.

## Installation

You should have Go installed on your device. To install Go, please refer to [this document](https://golang.org/doc/install).

After installing Go, you can clone this repo and then run this command.

```bash
go install github.com/raicem/podcast-archive
```

This will put `podcast-archive` executable to in your PATH.

## Usage

```bash
podcast-archiver FEED_URL
```

This will start downloading all episodes in the feed. Files will saved to the current folder.
