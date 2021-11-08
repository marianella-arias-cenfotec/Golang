# Small web service using gorilla mux

![image](https://user-images.githubusercontent.com/55120163/140699195-105d3e89-723d-4d51-9d4d-311fe1f8565e.png)


This project is about a small web service that with the Mux router will handle the http requests it receives.
Gorilla/mux is the only external library that would be use, is a powerful URL router and dispatcher.

# Previous required instalations

## Install Golang
First of all you need to have golang install in your computer for that go to https://golang.org/, according to your operating system.


## Install Gorilla/mux
$ go get github.com/gorilla/mux


## Install Docker
Install Docker wsl2 https://docs.docker.com/desktop/windows/install/

# Build Docker
$ docker build -t awesomeproject .

# Run Docker
$ docker run -p 8080 :8000 awesomeproject

