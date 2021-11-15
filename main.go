package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/grpc/proto"
	servers "github.com/wisdommatt/ecommerce-microservice-cart-service/grpc/service-servers"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/internal/cart"
	"github.com/wisdommatt/ecommerce-microservice-cart-service/services"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)

	mustLoadDotenv(log)

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.WithError(err).WithField("dbURL", os.Getenv("DATABASE_URL")).
			Fatal("an error occured while connecting to postgres")
	}
	db.AutoMigrate(&cart.CartItem{})

	tracer := initTracer("cart-service")
	opentracing.SetGlobalTracer(tracer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "2525"
	}

	cartRepo := cart.NewRepository(db, initTracer("postgres"))
	cartService := services.NewCartService(cartRepo, initTracer("cart.ServiceHandlers"))

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.WithError(err).WithField("port", port).Fatal("an error occured while listening on tcp port")
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	proto.RegisterCartServiceServer(grpcServer, servers.NewCartServer(cartService))
	log.WithField("port", port).Info("app running ...")
	log.WithField("port", port).Fatal(grpcServer.Serve(lis))
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
