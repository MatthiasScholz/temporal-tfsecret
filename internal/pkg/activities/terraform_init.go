package activities

import (
	"context"
	//"fmt"

	"go.temporal.io/sdk/activity"
	//"go.temporal.io/server/common/log"

	version "github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"

	"github.com/hashicorp/terraform-exec/tfexec"
)

var TERRAFORM_VERSION = "1.1.2"

type TerraformConfig struct {
	workdir   string
	workspace tfexec.Terraform
}

func installTerraform(ctx context.Context, tf_version string) (path string, err error) {
	logger := activity.GetLogger(ctx)
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(tf_version)),
	}

	logger.Info("Install")
	path, err = installer.Install(context.Background())

	return path, err
}

func TerraformInit(ctx context.Context, config TerraformConfig) (initialized bool, err error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Terraform Init: %s", config.workdir)

	// Always install the terraform binary since worker may change
	execpath, err := installTerraform(ctx, TERRAFORM_VERSION)
	if err != nil {
		logger.Error("Unable to install terraform binary: %s, %e", TERRAFORM_VERSION, err)
		return false, err
	}

	// Initialize terraform working directory
	//workspace, err := tfexec.NewTerraform(config.workdir, execpath)
	workspace, err := tfexec.NewTerraform("../../../examples/root", execpath)
	if err != nil {
		logger.Error("Unable to create terraform context: %s, %e", config.workdir, err)
		return false, err
	}
	logger.Info("Initializing terraform workspace: %s", config.workdir)
	err = workspace.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		logger.Error("Unable to initialize terraform workspace: %s, %e", config.workdir, err)
		return false, err
	}

	config.workspace = *config.workspace
	return true, err
}
