# Simple HTTP Server

Starts a simple HTTP File Server. Server runs on http://localhost:8888

## Run Server

* Method-1:

```
$ go run main.go
```

* Method-2:

```
$ go build -o simple-http-server main.go
$ ./simple-http-server
```

Once Server starts, launch the browser and navigate to `http://localhost:8888`

## Create Docker Image and Run Server

* Build Simple HTTP Server that can run in the docker container (Linux Version that is)

```
$ GOOS=linux GOARCH=amd64 go build -o simple-http-server main.go
```

* Create Docker Image

```
$ docker build -t simple-http-server:1.0 .
```

* Check Images to confirm its created

```
$ docker images simple-http-server:1.0
```

* Run the Server

```
$ docker run -d -p 80:8888 simple-http-server:1.0
```

The above `docker run` exposes maps host port 80 into container's port 8888 (on which the server is listening on). 
So, with docker container, `http://localhost` is enough and port number is not required.

# Sequence

```
     ,----.                         ,-------.                        ,------.
     |User|                         |Browser|                        |Server|
     `-+--'                         `---+---'                        `--+---'
       | Type in "http://localhost:8888"|                               |    
       | ------------------------------->                               |    
       |                                |                               |    
       |                                |             GET /             |    
       |                                | ------------------------------>    
       |                                |                               |    
       |                                |       Return index.html       |    
       |                                | <------------------------------    
       |                                |                               |    
       |                                | GET /sky-space-dark-galaxy.jpg|    
       |                                | ------------------------------>    
       |                                |                               |    
       |                                |----.                               
       |                                |    | Render index.html, picture    
       |                                |<---'                               
     ,-+--.                         ,---+---.                        ,--+---.
     |User|                         |Browser|                        |Server|
     `----'                         `-------'                        `------'
```