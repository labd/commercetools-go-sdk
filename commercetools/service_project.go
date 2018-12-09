package commercetools

// UpdateInput provides the data required to update a project.
type ProjectUpdateInput struct {
	// The expected version of the project on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the project.
	Actions []ProjectUpdateAction
}

// Service contains client information and bundles all actions.
type ProjectService struct {
	client *Client
}

// Get will return the current project. OAuth2 Scopes:
// view_project_settings:{projectKey}
func (svc *ProjectService) Get() (result *Project, err error) {
	err = svc.client.Get("", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update the current project with the defined UpdateActions. OAuth2
// Scopes: manage_project:{projectKey}
func (svc *ProjectService) Update(input *ProjectUpdateInput) (result *Project, err error) {
	err = svc.client.Update("", nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
