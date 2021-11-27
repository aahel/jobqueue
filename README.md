# jobqueue

###
`mq` packge contains :-
-  broker implementation. it provides functionality to subscribe to a topic and publish messages to  topic.
- `client.go`- a client library for the mq. creates a light abstraction around the broker for clients to consume.

`main.go`
tests the queue by publishing messages to a topic every two seconds and subscribing to the topic to get them in fifo order.

### RUN
`go run main.go`



