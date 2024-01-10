package activities

// import (
// 	"fmt"
//
// 	"go.temporal.io/sdk/workflow"
//
// 	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/provision/awsconfig"
// 	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/provision/tfactivity"
// 	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/provision/tfworkspace"
// 	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/workflows"
//
// 	terraform "github.com/MatthiasScholz/temporal-tfsecret/examples"
//
// 	// TODO Refactor - use HashiCorp
// 	"github.com/dynajoe/temporal-terraform-demo/tfexec"
// )
//
// func CreateSecretActivity(ctx workflow.Context, input workflows.SecretDetailsInput) (workflows.SecretDetailsOutput, error) {
// 	awsConfig := awsconfig.LoadConfig()
//
// 	attemptImport := make(map[string]string)
//
// 	// Lookup vpc by name for import
// 	// foundVpc, err := findVpcByName(ctx, awsConfig, input.Name)
// 	// if err != nil {
// 	// 	return CreateVPCOutput{}, err
// 	// }
// 	// if foundVpc.VpcId != nil {
// 	// 	attemptImport["aws_vpc.vpc"] = *foundVpc.VpcId
// 	// }
//
// 	// Temporal activity aware Terraform workspace wrapper
// 	tfa := tfactivity.New(tfworkspace.Config{
// 		TerraformPath: "aws/vpc",
// 		TerraformFS:   terraform.FS,
// 		S3Backend: tfexec.S3BackendConfig{
// 			Credentials: awsConfig.Credentials,
// 			Region:      "us-west-2",
// 			// TODO Somehow the state bucket must be provisioned
// 			Bucket:      "temporal-terraform-demo-state",
// 			Key:         fmt.Sprintf("vpc-%s.tfstate", input.Name),
// 		},
// 	})
//
// 	// Apply Terraform
// 	applyOutput, err := tfa.Apply(ctx, tfworkspace.ApplyInput{
// 		AttemptImport:  attemptImport,
// 		AwsCredentials: awsConfig.Credentials,
// 		Env: map[string]string{
// 			"AWS_REGION": input.Region,
// 		},
// 		// NOTE: This maps to terraform variables
// 		// TODO This is flacky and creates tight coupling
// 		Vars: map[string]interface{}{
// 			"name":       input.Name,
// 		},
// 	})
// 	if err != nil {
// 		return workflows.SecretDetailsOutput{}, err
// 	}
//
// 	// Extract output from Terraform
// 	// NOTE: This maps to terraform output
// 	// TODO This is flacky and creates tight coupling
// 	secret_name, err := applyOutput.String("secret")
// 	if err != nil {
// 		return workflows.SecretDetailsOutput{}, err
// 	}
// 	secret_arn, err := applyOutput.String("secret_arn")
// 	if err != nil {
// 		return workflows.SecretDetailsOutput{}, err
// 	}
//
// 	return workflows.SecretDetailsOutput{
// 		Name: secret_name,
// 		ARN: secret_arn,
// 	}, nil
// }
//
