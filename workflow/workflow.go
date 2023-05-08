package workflow

import (
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 150,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.Sleep(ctx, 100*time.Second)
	if err != nil {
		log.Fatalln("unable to sleep", err)
		return "", err
	}

	var result string
	err = workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
		return "", err
	}

	return result, err
}
