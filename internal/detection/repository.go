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
	wakeUpTime := time.Now().Add(30 * time.Second)
	_, err := temporalClient.ExecuteWorkflow(context.Background(), options, workflow.GreetingWorkflow, wakeUpTime)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
		return err
	}

	fmt.Println("Hittttt")

	return nil
}

func (r *DetectionRepository) EndDetection(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error {
	wakeUpTime := time.Now().Add(3 * time.Second)

	// Send a signal to the running Workflow to be executed right away
	err := temporalClient.SignalWorkflow(context.Background(), alarmScheduleID, "", workflow.SignalType, wakeUpTime)
	if err != nil {
		log.Fatalln("Unable to signale workflow", err)
	}
	log.Println("Signaled workflow to update wake-up time",
		"WorkflowID", alarmScheduleID, "WakeUpTime", wakeUpTime)
	return nil
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
