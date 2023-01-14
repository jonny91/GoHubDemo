package main

import (
	"context"
	"fmt"
	"github.com/jonny91/hubDemo/grades"
	"github.com/jonny91/hubDemo/log"
	"github.com/jonny91/hubDemo/registry"
	"github.com/jonny91/hubDemo/service"
	stlog "log"
)

func main() {
	host, port := "127.0.0.1", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
        ServiceUpdateURL: serviceAddress + "/services",
        HeartbeatURL:     serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
