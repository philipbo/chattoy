# chattoy
Chat Toy

## Server

    $ go run chat.go
    2016/04/29 10:43:18 chat server on :3000
    
    //指定端口
    $ go run chat.go -p :8888
    




## Client
    $ go run netcat.go
    You are 127.0.0.1:59201

    //指定端口
    $ go run netcat.go -p :8888