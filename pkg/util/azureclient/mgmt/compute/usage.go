package compute

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

// UsageClient is a minimal interface for azure UsageClient
type UsageClient interface {
	UsageClientAddons
}

type usageClient struct {
	compute.UsageClient
}

var _ UsageClient = &usageClient{}

// NewUsageClient creates a new UsageClient
func NewUsageClient(tenantID string, authorizer autorest.Authorizer) UsageClient {
	client := compute.NewUsageClient(tenantID)
	client.Authorizer = authorizer

	return &usageClient{
		UsageClient: client,
	}
}
