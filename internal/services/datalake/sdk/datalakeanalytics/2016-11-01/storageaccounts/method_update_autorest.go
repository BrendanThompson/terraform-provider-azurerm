package storageaccounts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type UpdateResponse struct {
	HttpResponse *http.Response
}

// Update ...
func (c StorageAccountsClient) Update(ctx context.Context, id StorageAccountId, input UpdateStorageAccountParameters) (result UpdateResponse, err error) {
	req, err := c.preparerForUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccounts.StorageAccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccounts.StorageAccountsClient", "Update", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccounts.StorageAccountsClient", "Update", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdate prepares the Update request.
func (c StorageAccountsClient) preparerForUpdate(ctx context.Context, id StorageAccountId, input UpdateStorageAccountParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdate handles the response to the Update request. The method always
// closes the http.Response Body.
func (c StorageAccountsClient) responderForUpdate(resp *http.Response) (result UpdateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
