package config

type Server struct {
	Http         int    `yaml:"http"`
	Websocket    int    `yaml:"ws"`
	Tcp          int    `yaml:"tcp"`
	GrpcHost     string `yaml:"grpc_host"`
	GrpcProtocol string `yaml:"grpc_protocol"`
	GrpcPort     int    `yaml:"grpc_port"`
}
