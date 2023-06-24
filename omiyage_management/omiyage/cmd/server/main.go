package main

import (
	"database/sql"
	"fmt"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/core/usecase/command"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/infrastructure/grpcserver"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/repository"
	"github.com/Sntree2mi8/samples/omiyage_management/omiyage/internal/service"
	omiyagev1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/omiyage/v1"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := 50051
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}

	conf := mysql.Config{
		User:                 "user",
		Passwd:               "pass",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "omiyage",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	s := grpcserver.New()
	omiyagev1.RegisterOmiyageServiceServer(
		s,
		service.NewOmiyageService(
			command.NewUserSignUpUseCase(
				repository.NewUserRepository(db),
			),
		),
	)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := s.Serve(l); err != nil {
			log.Printf("failed run gRPC server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
