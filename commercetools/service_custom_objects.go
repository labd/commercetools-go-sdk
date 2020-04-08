// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CustomObjectsURLPath is the commercetools API path.
const CustomObjectsURLPath = "custom-objects"

// CustomObjectsCreate creates a new instance of type CustomObject
func (client *Client) CustomObjectsCreate(draft *CustomObjectDraft) (result *CustomObject, err error) {
	err = client.Create(CustomObjectsURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectsQuery allows querying for type CustomObject
func (client *Client) CustomObjectsQuery(input *QueryInput) (result *CustomObjectPagedQueryResponse, err error) {
	err = client.Query(CustomObjectsURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CustomObjectsDeleteWithID The version control is optional. If the query contains a version, then it must match the version of the object.
*/
func (client *Client) CustomObjectsDeleteWithID(ID string, version int, dataErasure bool) (result *CustomObject, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("custom-objects/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectsGetWithID for type CustomObject
func (client *Client) CustomObjectsGetWithID(ID string) (result *CustomObject, err error) {
	err = client.Get(strings.Replace("custom-objects/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
