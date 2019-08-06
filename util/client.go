package util

import (
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
)

func GetBasicClient() (keyvault.BaseClient, error) {
	authorizer, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		deviceConfig := kvauth.NewDeviceFlowConfig(applicationID, tenantID)
		authorizer, derr := deviceConfig.Authorizer()
		if derr != nil {
			return keyvault.New(), derr
		}
	}

	basicClient := keyvault.New()
	basicClient.Authorizer = authorizer

	return basicClient, nil
}
