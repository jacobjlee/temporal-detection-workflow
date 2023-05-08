package detection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jacobjlee/temporal-detection-workflow/workflow"
	"go.temporal.io/sdk/client"
)

type Repository interface {
	StartDetection(ctx context.Context, temporalClient client.Client) error
	EndDetection(ctx context.Context, temporalClient client.Client, workflowID string) error
}

type DetectionRepository struct{}

func NewDetectionRepository() *DetectionRepository {
	return &DetectionRepository{}
}

func (r *DetectionRepository) StartDetection(ctx context.Context, temporalClient client.Client) error {
	options := client.StartWorkflowOptions{
		ID:                 "greeting-workflow",
		TaskQueue:          workflow.GreetingTaskQueue,
		WorkflowRunTimeout: time.Second * 100,
	}

	// Start the Workflow
	name := "World"
	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, workflow.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
		return err
	}

	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
		return err
	}

	printResults(greeting, we.GetID(), we.GetRunID())
	return nil
}

func (r *DetectionRepository) EndDetection(ctx context.Context, temporalClient client.Client, workflowID string) error {
	err := temporalClient.TerminateWorkflow(ctx, workflowID, "", "")
	if err != nil {
		log.Fatalln("unable to terminate Workflow", err)
	}

	printResults("Workflow terminated", workflowID, "")
	return nil
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
