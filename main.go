package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	mustLoadDotenv(log)

	tracer := initTracer("product-service")
	opentracing.SetGlobalTracer(tracer)
}

func mustLoadDotenv(log *logrus.Logger) {
	err := godotenv.Load(".env", ".env-defaults")
	if err != nil {
		log.WithError(err).Fatal("Unable to load env files")
	}
}

func initTracer(serviceName string) opentracing.Tracer {
	return initJaegerTracer(serviceName)
}

func initJaegerTracer(serviceName string) opentracing.Tracer {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}
	tracer, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatal("ERROR: cannot init Jaeger", err)
	}
	return tracer
}