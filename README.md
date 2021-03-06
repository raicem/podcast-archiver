# podcast-archiver

A CLI tool to download all podcast episodes in a feed.

## Installation

You should have Go installed on your device. To install Go, please refer to [this document](https://golang.org/doc/install).

After installing Go, you can clone this repo and then run this command.

```bash
go install
```

This will put `podcast-archive` executable to in your PATH.

## Usage

```bash
podcast-archiver [OPTIONS...] FEED_URL

OPTIONS
    -limit <number> Number of episodes to download. Default is unlimited.
    -order <direction> Order to start downloading episodes. Possible values are "newest" and "oldest". Default is newest."
```

This will start downloading all episodes in the feed. Files will saved to the current folder.
