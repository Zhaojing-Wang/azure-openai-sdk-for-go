//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.6, generator: @autorest/go@4.0.0-preview.50)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package azureopenai

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ChatCompletionsClient contains the methods for the ChatCompletions group.
// Don't use this type directly, use a constructor function instead.
type ChatCompletionsClient struct {
	internal *azcore.Client
}

// Create - Creates a completion for the chat message
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - options - ChatCompletionsClientCreateOptions contains the optional parameters for the ChatCompletionsClient.Create method.
func (client *ChatCompletionsClient) Create(ctx context.Context, deploymentID string, body Paths1L1E8YpDeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchema, options *ChatCompletionsClientCreateOptions) (ChatCompletionsClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, deploymentID, body, options)
	if err != nil {
		return ChatCompletionsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChatCompletionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChatCompletionsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ChatCompletionsClient) createCreateRequest(ctx context.Context, deploymentID string, body Paths1L1E8YpDeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchema, options *ChatCompletionsClientCreateOptions) (*policy.Request, error) {
	host := "https://{endpoint}/openai"
	host = strings.ReplaceAll(host, "{endpoint}", endpoint)
	urlPath := "/deployments/{deployment-id}/chat/completions"
	if deploymentID == "" {
		return nil, errors.New("parameter deploymentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deployment-id}", url.PathEscape(deploymentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ChatCompletionsClient) createHandleResponse(resp *http.Response) (ChatCompletionsClientCreateResponse, error) {
	result := ChatCompletionsClientCreateResponse{}
	if val := resp.Header.Get("apim-request-id"); val != "" {
		result.ApimRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Paths1H0F83DeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchema); err != nil {
		return ChatCompletionsClientCreateResponse{}, err
	}
	return result, nil
}
