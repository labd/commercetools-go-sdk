package commercetools

import (
	"context"
	"sort"
	"time"
)

// CustomApplicationLabelLocale describes the structure of a localized label.
type CustomApplicationLabelLocale struct {
	Locale string `json:"locale"`
	Value  string `json:"value"`
}

type CustomApplicationNavbarMenu struct {
	Key             string                           `json:"key"`
	URIPath         string                           `json:"uriPath"`
	Icon            string                           `json:"icon"`
	LabelAllLocales []CustomApplicationLabelLocale   `json:"labelAllLocales,omitempty"`
	Permissions     []string                         `json:"permissions"`
	Submenu         []CustomApplicationNavbarSubmenu `json:"submenu"`
}

type CustomApplicationNavbarSubmenu struct {
	Key             string                         `json:"key"`
	URIPath         string                         `json:"uriPath"`
	LabelAllLocales []CustomApplicationLabelLocale `json:"labelAllLocales"`
	Permissions     []string                       `json:"permissions"`
}

type CustomApplicationDraft struct {
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	URL         string                      `json:"url"`
	NavbarMenu  CustomApplicationNavbarMenu `json:"navbarMenu"`
	OAuthScopes []string                    `json:"oAuthScopes,omitempty"`
}

// CustomApplication describes the structure of a Custom Application stored object.
type CustomApplication struct {
	ID          string                      `json:"id"`
	CreatedAt   time.Time                   `json:"createdAt,string"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
	IsActive    bool                        `json:"isActive"`
	Name        string                      `json:"name"`
	Description *string                     `json:"description,omitempty"`
	URL         string                      `json:"url"`
	NavbarMenu  CustomApplicationNavbarMenu `json:"navbarMenu"`
}

type CustomApplicationPagedQueryResponse struct {
	Total   int                 `json:"total,omitempty"`
	Results []CustomApplication `json:"results"`
	Offset  int                 `json:"offset"`
	Limit   int                 `json:"limit"`
	Count   int                 `json:"count"`
}

// ProjectExtension describes the structure of a project extension stored object.
type ProjectExtension struct {
	ID           string              `json:"id"`
	Applications []CustomApplication `json:"applications"`
}

// GraphQLResponseProjectExtension describes the structure of the query result for fetching a Custom Application.
type graphQLResponseProjectExtension struct {
	ProjectExtension *ProjectExtension `json:"projectExtension"`
}

// GraphQLResponseProjectExtensionCreation describes the structure of the query result for creating a Custom Application.
type graphQLResponseProjectExtensionCreation struct {
	CreateProjectExtensionApplication *ProjectExtension `json:"createProjectExtensionApplication"`
}

// GraphQLResponseProjectExtensionUpdate describes the structure of the query result for updating a Custom Application.
type graphQLResponseProjectExtensionUpdate struct {
	UpdateProjectExtensionApplication     *ProjectExtension `json:"updateProjectExtensionApplication"`
	ActivateProjectExtensionApplication   *ProjectExtension `json:"activateProjectExtensionApplication"`
	DeactivateProjectExtensionApplication *ProjectExtension `json:"deactivateProjectExtensionApplication"`
}

// GraphQLResponseProjectExtensionDeletion describes the structure of the query result for deleting a Custom Application.
type graphQLResponseProjectExtensionDeletion struct {
	DeleteProjectExtensionApplication *ProjectExtension `json:"deleteProjectExtensionApplication"`
}

func (client *Client) CustomApplicationCreate(ctx context.Context, draft *CustomApplicationDraft) (*CustomApplication, error) {
	var result graphQLResponseProjectExtensionCreation

	query := client.NewGraphQLQuery(`
		mutation CreateCustomApplicationMutation($draft: ApplicationExtensionDataInput!) {
			createProjectExtensionApplication(data: $draft) {
				id
				applications {
					id
					name
					url
					createdAt
				}
			}
		}`)

	query.Bind("draft", draft)
	err := query.Execute(&result, query.ForMerchantCenter())
	if err != nil {
		return nil, err
	}

	apps := result.CreateProjectExtensionApplication.Applications

	// Sort the apps descending on the created timestamp so we can easily find
	// the last created item
	sort.Slice(apps, func(i, j int) bool {
		return apps[i].CreatedAt.After(apps[j].CreatedAt)
	})

	for _, app := range apps {
		if app.Name == draft.Name && app.URL == draft.URL {
			return client.CustomApplicationGetWithID(ctx, app.ID)
		}
	}

	return nil, nil
}

func (client *Client) CustomApplicationGetWithID(ctx context.Context, ID string) (*CustomApplication, error) {
	var result graphQLResponseProjectExtension

	query := client.NewGraphQLQuery(`
		query FetchCustomApplicationById($applicationId: ID!) {
			projectExtension {
				id
				applications(where: { id: $applicationId }) {
					id
					createdAt
					updatedAt
					isActive
					name
					description
					url
					navbarMenu {
						uriPath
						icon
						labelAllLocales {
							locale
							value
						}
						permissions
						submenu {
							uriPath
							labelAllLocales {
								locale
								value
							}
							permissions
						}
					}
				}
			}
		}
	`)

	query.Bind("applicationId", ID)
	err := query.Execute(&result, query.ForMerchantCenter())
	if err != nil {
		return nil, err
	}
	if len(result.ProjectExtension.Applications) == 1 {
		return &result.ProjectExtension.Applications[0], nil
	}
	return nil, nil
}

func (client *Client) CustomApplicationQuery(ctx context.Context, input *QueryInput) (*CustomApplicationPagedQueryResponse, error) {
	var result graphQLResponseProjectExtension

	query := client.NewGraphQLQuery(`
		query FetchCustomApplicationById {
			projectExtension {
				id
				applications {
					id
					createdAt
					updatedAt
					isActive
					name
					description
					url
					navbarMenu {
						uriPath
						icon
						labelAllLocales {
							locale
							value
						}
						permissions
						submenu {
							uriPath
							labelAllLocales {
								locale
								value
							}
							permissions
						}
					}
				}
			}
		}
	`)

	err := query.Execute(&result, query.ForMerchantCenter())
	if err != nil {
		return nil, err
	}

	response := CustomApplicationPagedQueryResponse{
		Results: result.ProjectExtension.Applications,
		Total:   len(result.ProjectExtension.Applications),
		Count:   len(result.ProjectExtension.Applications),
		Offset:  0,
		Limit:   0,
	}
	return &response, nil
}

func (client *Client) CustomApplicationDeleteWithID(ctx context.Context, ID string) error {
	var result graphQLResponseProjectExtensionDeletion

	query := client.NewGraphQLQuery(`
		mutation DeleteCustomApplication($applicationId: ID!) {
			deleteProjectExtensionApplication(applicationId: $applicationId) {
				id
			}
		}
	`)

	query.Bind("applicationId", ID)

	err := query.Execute(&result, query.ForMerchantCenter())
	return err
}

func (client *Client) CustomApplicationUpdateWithID(ctx context.Context, ID string, draft *CustomApplicationDraft) (*CustomApplication, error) {
	var result graphQLResponseProjectExtensionUpdate

	query := client.NewGraphQLQuery(`
		mutation UpdateCustomApplication(
			$applicationId: ID!
			$draft: ApplicationExtensionDataInput!
		) {
			updateProjectExtensionApplication(
				applicationId: $applicationId
				data: $draft
			) {
				id
				applications(where: { id: $applicationId }) {
					id
				}
			}
		}
	`)
	query.Bind("applicationId", ID)
	query.Bind("draft", draft)

	err := query.Execute(&result, query.ForMerchantCenter())
	if err != nil {
		return nil, err
	}
	return client.CustomApplicationGetWithID(ctx, ID)
}
