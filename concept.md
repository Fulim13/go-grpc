## Grpc

1.  Serialization with protobuffer
2.  Multiplexing - gRPC sends many requests at the same time over one long-lived TCP connection (Good for concurrent scenario)

- With HTTP/1.1, a connection can handle only one request/response at a time.
  If you send request B while request A is still in flight, B waits.
  To get concurrency, clients open several connections, and each new connection pays for: a TCP handshake, a TLS handshake (if using HTTPS)

- With HTTP/2 (one connection, many streams interleaved)

```
conn1: [A1][B1][C1][A2][D1][B2][C2][D2][A3]...
└─ each RPC is a "stream" with its own ID;
its data is split into frames that get mixed together
```
