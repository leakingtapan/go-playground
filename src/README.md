# Comparing Linux Inter Process Latency
There are mutiple ways to achieve inter process communication on linux machine.
1. Application Protocal
   1. HTTP
   2. HTTP2
2. Transport Protocal
   1. TCP 
   2. unix domain socket
3. Transport Encoding
   1. Streaming
   2. Regular 

# Experiements
## Test scenario
Client send a request to server. The client request contains only the client timestamp before the request is sent. The server reads the server timestamp after client request is recieved and prints both times into stdout.

## Test cases
In the following experiments, I tested the latecy of using differnt combinations of application protocal, transport protocal and transport encoding.
1. Using regular HTTP. This is using TCP as underlying transport.
2. Using HTTP with unix domain socket.
3. Using HTTP with streaming on TCP. 
4. Using HTTP with streaming on unix socket
5. Using gRPC which is a RPC framework

## Results
For each test case, it is ran 100 times and the latency is calculated by the average time differene between client and server. 

### EC2 Instance (t2.small)
| HTTP + TCP | HTTP + TCP + Streaming | HTTP + socket | HTTP + socket + stream | HTTP2 + TCP | gRPC |
| ---------- | ---------------------- | ------------- | ---------------------- | ----------- | ---- |
| 70.4 μs    | 21.80 μs               | 55.10 μs      | 10.70 μs               | 192.00 μs   | 100.60 μs|

### Pi (TBD)
| HTTP + TCP | HTTP + TCP + Streaming | HTTP + socket | HTTP + socket + stream | HTTP2 + TCP | gRPC |
| ---------- | ---------------------- | ------------- | ---------------------- | ----------- | ---- |
|  μs    |  μs               |  μs      |  μs               |  μs   |  μs|

# Conclusions
From the result, I get follow conclusions:
1. Unix domain socket has less latency
2. HTTP streaming has less latency than non-streaming
3. HTTP2 is slower than traditional HTTP. This might be caused by the extra overhead from TLS.
