package producer

//go:generate options-gen -from-struct=Options --defaults-from=var
type Options struct {
	addrs []string `option:"mandatory" validate:"required"`

	topic    string
	tls      bool
	clientID string `validate:"required"`
}

var defaultOptions = Options{
	addrs:    nil,
	topic:    "",
	tls:      false,
	clientID: "golang-kafka-client",
}
