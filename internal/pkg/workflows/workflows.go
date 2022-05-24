package workflows

import (
	"go.temporal.io/sdk/worker"
)

func Register(w worker.Worker) {
	w.RegisterWorkflow(ProvisionSecretWorkflow)

	//w.RegisterWorkflow(DeprovisionSecret)

	// w.RegisterWorkflow(workflows.BackgroundCheck)
	// w.RegisterWorkflow(workflows.Accept)
	// w.RegisterWorkflow(workflows.EmploymentVerification)
	// //w.RegisterActivity(&activities.Activities{SMTPHost: "mailhog", SMTPPort: 1025})
	// w.RegisterWorkflow(workflows.SSNTrace)
	// w.RegisterWorkflow(workflows.FederalCriminalSearch)
	// w.RegisterWorkflow(workflows.StateCriminalSearch)
	// w.RegisterWorkflow(workflows.MotorVehicleIncidentSearch)
}
