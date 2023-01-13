package main

import (
	"context"
	"fmt"
	"github.com/jonny91/hubDemo/log"
	"github.com/jonny91/hubDemo/service"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "0.0.0.0", "12155"
	ctx, err := service.Start(context.Background(), "Log Service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
		return
	}
	<-ctx.Done()

	fmt.Println("Shutting Down - Log Service ...")
}
