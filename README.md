This is my GRPC STUDY with Go
In This repo, there are 3 main folders,
1. proto : this use for implementing the protocol buffer for using GRPC
2. client: this use to save all GRPC implementation of the client side
3. server: this use to save all GRPC implementation of the server side

**There Are 4 Types of communication between client and server in GRPC**

1. Unary API : Regular request and response between client and server (->client request, <-server response, synchhronous)
2. Server Streaming: Client sends a request to the server, but the server gets a stream back (like get sequence of messages) (->client 1 request, <-server response a stream of messages)
3. Client Streaming: Client sends a stream of messages to the server and the server just send a response  (-> client request stream of messages, <- server give 1 response)
4. Bidirectional Streaming: Client and server both communicate  using streams even though it's a "stream" not a "queue", otherwise the order of the messages is preserved (careful, this is asynchronous,
   so the messages or the data that each sends are not waiting for the others, like forgot the classic Request and Response API). (-> client send the stream of data and asynchronously <-server send the stream also)


**Notes: If You already include GRPC in your project doesnt means that you cannot implement the regular request and response**
**Youtube link that I Learn this: https://youtu.be/a6G5-LUlFO4?si=OA-Ty2HHyGEuyeiu**
