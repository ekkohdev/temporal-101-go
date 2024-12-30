package finale

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func CertificateGeneratorWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "recipientName", name)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var outFilePath string
	// CreatePdf is the Activity Type defined in the Java Activity code
	err := workflow.ExecuteActivity(ctx, "CreatePdf", name).Get(ctx, &outFilePath)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("CertificateGeneratorWorkflow created file", outFilePath)

	return outFilePath, nil
}
