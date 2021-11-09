package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/go-azure-helpers/sender"
)

func buildServicePrincipalObjectIDFunc(c *Config) func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		objectId, err := objectIdFromMsGraph(ctx, c)
		if err == nil {
			return objectId, nil
		}
		if !strings.HasPrefix(err.Error(), "access denied") {
			return "", err
		}

		log.Printf("Unable to retrieve service principal object ID from Microsoft Graph: %+v", err)
		return objectIdFromAadGraph(ctx, c)
	}
}

func objectIdFromAadGraph(ctx context.Context, c *Config) (string, error) {
	env, err := AzureEnvironmentByNameFromEndpoint(ctx, c.MetadataHost, c.Environment)
	if err != nil {
		return "", err
	}

	s := sender.BuildSender("GoAzureHelpers")

	oauthConfig, err := c.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return "", err
	}

	graphAuth, err := c.GetAuthorizationToken(s, oauthConfig, env.GraphEndpoint)
	if err != nil {
		return "", err
	}

	client := graphrbac.NewServicePrincipalsClientWithBaseURI(env.GraphEndpoint, c.TenantID)
	client.Authorizer = graphAuth
	client.Sender = s

	filter := fmt.Sprintf("appId eq '%s'", c.ClientID)
	listResult, listErr := client.List(ctx, filter)

	if listErr != nil {
		return "", fmt.Errorf("listing Service Principals: %#v", listErr)
	}

	if listResult.Values() == nil || len(listResult.Values()) != 1 || listResult.Values()[0].ObjectID == nil {
		return "", fmt.Errorf("unexpected Service Principal query result: %#v", listResult.Values())
	}

	return *listResult.Values()[0].ObjectID, nil
}

func objectIdFromMsGraph(ctx context.Context, c *Config) (string, error) {
	env, err := environments.EnvironmentFromString(c.Environment)
	if err != nil {
		return "", err
	}

	oauthConfig, err := c.BuildOAuthConfig(string(env.AzureADEndpoint))
	if err != nil {
		return "", err
	}

	msGraphAuth, err := c.GetAuthorizationTokenV2(sender.BuildSender("GoAzureHelpers"), oauthConfig, string(env.MsGraph.Endpoint))
	if err != nil {
		return "", err
	}

	authorizerWrapper, err := auth.NewAutorestAuthorizerWrapper(msGraphAuth)
	if err != nil {
		return "", err
	}

	client := msgraph.NewServicePrincipalsClient(c.TenantID)
	client.BaseClient.ApiVersion = msgraph.Version10
	client.BaseClient.Authorizer = authorizerWrapper
	client.BaseClient.DisableRetries = true
	client.BaseClient.Endpoint = env.MsGraph.Endpoint
	client.BaseClient.RequestMiddlewares = &[]msgraph.RequestMiddleware{hamiltonRequestLogger}
	client.BaseClient.ResponseMiddlewares = &[]msgraph.ResponseMiddleware{hamiltonResponseLogger}

	result, status, err := client.List(ctx, odata.Query{Filter: fmt.Sprintf("appId eq '%s'", c.ClientID)})
	if err != nil {
		if status == http.StatusUnauthorized || status == http.StatusForbidden {
			return "", fmt.Errorf("access denied when listing Service Principals: %+v", err)
		}
		return "", fmt.Errorf("listing Service Principals: %+v", err)
	}

	if result == nil {
		return "", fmt.Errorf("unexpected Service Principal query result, was nil")
	}

	if len(*result) != 1 || (*result)[0].ID == nil {
		return "", fmt.Errorf("unexpected Service Principal query result: %+v", *result)
	}

	return *(*result)[0].ID, nil
}

func hamiltonRequestLogger(req *http.Request) (*http.Request, error) {
	if req == nil {
		return nil, nil
	}

	if dump, err := httputil.DumpRequestOut(req, true); err == nil {
		log.Printf("[DEBUG] GoAzureHelpers Request: \n%s\n", dump)
	} else {
		log.Printf("[DEBUG] GoAzureHelpers Request: %s to %s\n", req.Method, req.URL)
	}

	return req, nil
}

func hamiltonResponseLogger(req *http.Request, resp *http.Response) (*http.Response, error) {
	if resp == nil {
		log.Printf("[DEBUG] GoAzureHelpers Request for %s %s completed with no response", req.Method, req.URL)
		return nil, nil
	}

	if dump, err := httputil.DumpResponse(resp, true); err == nil {
		log.Printf("[DEBUG] GoAzureHelpers Response: \n%s\n", dump)
	} else {
		log.Printf("[DEBUG] GoAzureHelpers Response: %s for %s %s\n", resp.Status, req.Method, req.URL)
	}

	return resp, nil
}
