package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"

	"github.com/Azure/ARO-RP/pkg/api"
)

func (f *frontend) getReady(w http.ResponseWriter, r *http.Request) {
	if f.ready.Load().(bool) && f.env.ArmClientAuthorizer().IsReady() && f.env.AdminClientAuthorizer().IsReady() {
		api.WriteCloudError(w, &api.CloudError{StatusCode: http.StatusOK})
	} else {
		api.WriteError(w, http.StatusInternalServerError, api.CloudErrorCodeInternalServerError, "", "Internal server error.")
	}
}

func (f *frontend) getHealthz(w http.ResponseWriter, r *http.Request) {
	api.WriteCloudError(w, &api.CloudError{StatusCode: http.StatusOK})
}
