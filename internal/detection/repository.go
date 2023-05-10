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
	StartDetection(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error
	EndDetection(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error
}

type DetectionRepository struct{}

func NewDetectionRepository() *DetectionRepository {
	return &DetectionRepository{}
}

func (r *DetectionRepository) StartDetection(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error {
	options := client.StartWorkflowOptions{
		ID:                 alarmScheduleID,
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

func (r *DetectionRepository) EndDetection(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error {

	//err := temporalClient.TerminateWorkflow(ctx, alarmScheduleID, "", "")
	//if err != nil {
	//	log.Fatalln("unable to terminate Workflow", err)
	//}
	//
	//printResults("Workflow terminated", alarmScheduleID, "")
	return nil
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
