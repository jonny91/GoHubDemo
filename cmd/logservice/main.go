package main

import (
	"context"
	"fmt"
	"github.com/jonny91/hubDemo/log"
	"github.com/jonny91/hubDemo/registry"
	"github.com/jonny91/hubDemo/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "0.0.0.0", "12155"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
		return
	}
	<-ctx.Done()

	fmt.Println("Shutting Down - Log Service ...")
}
