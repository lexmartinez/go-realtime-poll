# go-realtime-poll
:postbox: A Simple Electronic Realtime Poll written in Go 

Built following the original idea of [@christiannwamba](https://github.com/christiannwamba) at `codementor.io` article: [How to Build an Electronic Realtime Poll in 5 Steps](https://www.codementor.io/christiannwamba/how-to-build-an-electronic-realtime-poll-in-5-steps-es41iluc9)

## Build Setup

``` bash

# clone repo
$ git clone https://github.com/lexmartinez/go-realtime-poll

# change directory to cloned app
$ cd go-realtime-poll

# go build
$ go build server.go

# start the server
$ ./server

```

## Environment
For pusher connection following environment variables must be configured into `.env` file located at project root (with your own values)

```bash
APP_ID=545346574
APP_KEY=867564392204459023343
APP_SECRET=b434mnm45n3j28fad
APP_CLUSTER=us2
```````

## Screenshots

![](https://github.com/lexmartinez/go-realtime-poll/raw/master/screenshots/screenshot-1.png)

![](https://github.com/lexmartinez/go-realtime-poll/raw/master/screenshots/screenshot-2.png)

![](https://github.com/lexmartinez/go-realtime-poll/raw/master/screenshots/screenshot-3.png)

## License

This project is licensed under MIT License - see the [LICENSE.md](https://github.com/lexmartinez/go-realtime-poll/blob/master/LICENSE.md) file for details
