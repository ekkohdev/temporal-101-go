package main

import (
	"context"
	"log"
	"os"

	hello "temporal-101-go/hello-workflow"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "my-go-sdk-workflow",
		TaskQueue: "greeting-tasks",
	}

	run, err := c.ExecuteWorkflow(context.Background(), options, hello.GreetSomeone, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", run.GetID(), "RunID", run.GetRunID())

	var result string
	err = run.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
