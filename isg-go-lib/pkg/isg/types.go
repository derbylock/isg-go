package isg

type InterfaceType string

const (
	HTTP  InterfaceType = "http"
	GRPC  InterfaceType = "grpc"
	TOPIC InterfaceType = "topic"
	JOB   InterfaceType = "job"
	TCP   InterfaceType = "tcp"
	UDP   InterfaceType = "udp"
	DB    InterfaceType = "db"
	OTHER InterfaceType = "other"
)
