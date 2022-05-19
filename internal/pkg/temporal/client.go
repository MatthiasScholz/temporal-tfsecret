// https://github.com/temporalio/samples-go/tree/main/encryption
package temporal

import (
	"os"

	encryption "github.com/MatthiasScholz/temporal-plugins-dataconverter/pkg"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

func NewClient(options client.Options) (client.Client, error) {
	if options.HostPort == "" {
		options.HostPort = os.Getenv("TEMPORAL_GRPC_ENDPOINT")
	}

	options.DataConverter = encryption.NewEncryptionDataConverter(
		converter.GetDefaultDataConverter(),
		encryption.DataConverterOptions{KeyID: os.Getenv("DATACONVERTER_ENCRYPTION_KEY_ID")},
	)

	return client.NewClient(options)
}
