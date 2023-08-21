# RSS Feed Aggregator in Go

## Overview

This project is an RSS feed aggregator built in Go. It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Learning Goals

- Learn how to integrate a Go server with PostgreSQL
- Learn about the basics of database migrations
- Learn about long-running service workers

## Setup

- An editor. I use VS code, you can use whatever you like.
- A command line. I work on Mac OS/Linux, so instructions will be in Bash. I recommend WSL 2 if you're on Windows so you can still use Linux commands.
- The latest Go toolchain.
- If you're in VS Code, I recommend the official Go extension.
- An HTTP client. I use Thunder Client, but you can use whatever you like.
