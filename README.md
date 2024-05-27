# Blog Aggregator
An extension of a guided project from [boot.dev](https://www.boot.dev/learn/build-blog-aggregator)

## What is this?
This project collects rss feeds that you give it, and allows you to condense all of the RSS feeds for a given user into one list!

## Why is this here?
I'm going to use this as a pet project to practice my fullstack app skeels

## How do I download and use this project?
I don't see why you'd want to, but

first you'd need to clone the repo on to your machine with `git clone`

### Back end
cd into the "server" subdirectory

Add a .env folder with the following keys:

```
PORT="8000"
DB_CONNECTOR="<a url link to your Postgres database>"
```
then all you need to do is run

```bash
$ go build
$ ./blog-aggregator
```

you can then test the health by going to the `localhost:8000/v1/healthz` with an http client of your choice.
If you get a 200 back, the server is up and running!

### Front End
you need to do all your typical Next.js app stuff

cd into "next-blogator"

install dependencies:
```bash
npm install
```
and run the dev frontend server
```bash
npm run dev
```

