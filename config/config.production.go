package config

type configProduction struct {
	HttpPort string `env:"HTTP_PORT" default:"6001"`
	GrpcPort string `env:"GRPC_PORT" default:"6000"`
}

var instance *configProduction

func GetConfigProduction() *configProduction {
	if instance == nil {
		instance = &configProduction{
			HttpPort: "6001",
			GrpcPort: "6000",
		}
	}
	return instance
}
