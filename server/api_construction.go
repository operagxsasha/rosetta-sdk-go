// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// A ConstructionAPIController binds http requests to an api service and writes the service results
// to the http response
type ConstructionAPIController struct {
	service ConstructionAPIServicer
}

// NewConstructionAPIController creates a default api controller
func NewConstructionAPIController(s ConstructionAPIServicer) Router {
	return &ConstructionAPIController{service: s}
}

// Routes returns all of the api route for the ConstructionAPIController
func (c *ConstructionAPIController) Routes() Routes {
	return Routes{
		{
			"ConstructionMetadata",
			strings.ToUpper("Post"),
			"/construction/metadata",
			c.ConstructionMetadata,
		},
		{
			"ConstructionSubmit",
			strings.ToUpper("Post"),
			"/construction/submit",
			c.ConstructionSubmit,
		},
	}
}

// ConstructionMetadata - Get Transaction Construction Metadata
func (c *ConstructionAPIController) ConstructionMetadata(w http.ResponseWriter, r *http.Request) {
	constructionMetadataRequest := &types.ConstructionMetadataRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionMetadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionMetadataRequest is correct
	if err := asserter.ConstructionMetadataRequest(constructionMetadataRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionMetadata(constructionMetadataRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}

// ConstructionSubmit - Submit a Signed Transaction
func (c *ConstructionAPIController) ConstructionSubmit(w http.ResponseWriter, r *http.Request) {
	constructionSubmitRequest := &types.ConstructionSubmitRequest{}
	if err := json.NewDecoder(r.Body).Decode(&constructionSubmitRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	// Assert that ConstructionSubmitRequest is correct
	if err := asserter.ConstructionSubmitRequest(constructionSubmitRequest); err != nil {
		EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)

		return
	}

	result, serviceErr := c.service.ConstructionSubmit(constructionSubmitRequest)
	if serviceErr != nil {
		EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)

		return
	}

	EncodeJSONResponse(result, http.StatusOK, w)
}
