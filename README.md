# Kafka wrapper on top of kafka-go

```shell
go get -u github.com/kazhuravlev/kafka
```

## Consumer

Example with handling messages in json and protobuf format (depends on message header `enc`). You can write own
algorithm to detect which decoder should used or use strict decoder like `new(JSONDecoder)`.

```go
package main

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/kafka/consumer"
)

type Payload struct {
	UserID int64
}

func main() {
	cons, err := consumer.New[Payload](consumer.NewOptions(
		consumer.WithBrokers([]string{"127.0.0.1:9092"}),
		consumer.WithDecoder(&consumer.HeaderDependantDecoder{
			HeaderName: "enc",
			Decoders: map[string]consumer.IDecoder{
				"json":  new(consumer.JSONDecoder),
				"proto": new(consumer.ProtoJSONDecoder),
			},
			DefaultDecoder: new(consumer.AlwaysFailDecoder),
		}),
	))
	if err != nil {
		panic(err)
	}

	if err := cons.Run(context.TODO()); err != nil {
		panic(err)
	}

	for m := range cons.C() {
		fmt.Println(m)
	}
}
```
