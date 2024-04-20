# Echo Server
Echo server returns whatever has been been send to it over TCP connection.
It is written in GoLang.


## Start server
*Note*: By default server will start on port `7878`

```shell
go run main.go
```

To run in different port eg: `5200`
```shell
go run main.go -port 5200
```

## Start client and interact with server

Using netcat to simulate client interaction

### Installation
Linux:
```shell
sudo apt-get install netcat
```
Mac:
```shell
brew install netcat
```

### Connect to server
```shell
nc localhost 7878
```
