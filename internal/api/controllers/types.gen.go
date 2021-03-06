// Package public provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package public

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Account defines model for Account.
type Account string

// CreatedAt defines model for CreatedAt.
type CreatedAt time.Time

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Labels defines model for Labels.
type Labels struct {
	AdditionalProperties map[string]string `json:"-"`
}

// Meta defines model for Meta.
type Meta struct {

	// number of results returned
	Count int `json:"count"`
}

// Run defines model for Run.
type Run struct {

	// Identifier of the tenant
	Account *Account `json:"account,omitempty"`

	// A timestamp when the entry was created
	CreatedAt *CreatedAt `json:"created_at,omitempty"`

	// Unique identifier of a Playbook run
	Id *RunId `json:"id,omitempty"`

	// Additional metadata about the Playbook run. Can be used for filtering purposes.
	Labels *Labels `json:"labels,omitempty"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient *RunRecipient `json:"recipient,omitempty"`

	// Current status of a Playbook run
	Status *RunStatus `json:"status,omitempty"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *RunTimeout `json:"timeout,omitempty"`

	// A timestamp when the entry was last updated
	UpdatedAt *UpdatedAt `json:"updated_at,omitempty"`

	// URL hosting the Playbook
	Url *Url `json:"url,omitempty"`
}

// RunCreated defines model for RunCreated.
type RunCreated struct {

	// status code of the request
	Code int `json:"code"`

	// Unique identifier of a Playbook run
	Id *RunId `json:"id,omitempty"`
}

// RunHost defines model for RunHost.
type RunHost struct {

	// Name used to identify a host within Ansible inventory
	Host *string `json:"host,omitempty"`
	Run  *Run    `json:"run,omitempty"`

	// Current status of a Playbook run
	Status *RunStatus `json:"status,omitempty"`

	// Output produced by running Ansible Playbook on the given host
	Stdout *string `json:"stdout,omitempty"`
}

// RunHosts defines model for RunHosts.
type RunHosts struct {
	Data []RunHost `json:"data"`

	// Information about returned entities
	Meta Meta `json:"meta"`
}

// RunId defines model for RunId.
type RunId string

// RunInput defines model for RunInput.
type RunInput struct {

	// Identifier of the tenant
	Account Account `json:"account"`

	// Additional metadata about the Playbook run. Can be used for filtering purposes.
	Labels *Labels `json:"labels,omitempty"`

	// Identifier of the host to which a given Playbook is addressed
	Recipient RunRecipient `json:"recipient"`

	// Amount of seconds after which the run is considered failed due to timeout
	Timeout *RunTimeout `json:"timeout,omitempty"`

	// URL hosting the Playbook
	Url Url `json:"url"`
}

// RunLabelsNullable defines model for RunLabelsNullable.
type RunLabelsNullable struct {
	AdditionalProperties map[string]string `json:"-"`
}

// RunRecipient defines model for RunRecipient.
type RunRecipient string

// RunStatus defines model for RunStatus.
type RunStatus string

// List of RunStatus
const (
	RunStatus_failure RunStatus = "failure"
	RunStatus_running RunStatus = "running"
	RunStatus_success RunStatus = "success"
	RunStatus_timeout RunStatus = "timeout"
)

// RunTimeout defines model for RunTimeout.
type RunTimeout int

// Runs defines model for Runs.
type Runs struct {
	Data []Run `json:"data"`

	// Information about returned entities
	Meta Meta `json:"meta"`
}

// RunsCreated defines model for RunsCreated.
type RunsCreated []RunCreated

// StatusNullable defines model for StatusNullable.
type StatusNullable string

// List of StatusNullable
const (
	StatusNullable_failure StatusNullable = "failure"
	StatusNullable_running StatusNullable = "running"
	StatusNullable_success StatusNullable = "success"
	StatusNullable_timeout StatusNullable = "timeout"
)

// UpdatedAt defines model for UpdatedAt.
type UpdatedAt time.Time

// Url defines model for Url.
type Url string

// Limit defines model for Limit.
type Limit int

// Offset defines model for Offset.
type Offset int

// RunHostFields defines model for RunHostFields.
type RunHostFields struct {
	Data *[]string `json:"data,omitempty"`
}

// RunHostFilter defines model for RunHostFilter.
type RunHostFilter struct {
	Run *struct {
		Id     *string            `json:"id"`
		Labels *RunLabelsNullable `json:"labels"`
	} `json:"run"`
	Status *StatusNullable `json:"status"`
}

// RunsFields defines model for RunsFields.
type RunsFields struct {
	Data *[]string `json:"data,omitempty"`
}

// RunsFilter defines model for RunsFilter.
type RunsFilter struct {
	Labels    *RunLabelsNullable `json:"labels"`
	Recipient *string            `json:"recipient"`
	Status    *StatusNullable    `json:"status"`
}

// RunsSortBy defines model for RunsSortBy.
type RunsSortBy string

// List of RunsSortBy
const (
	RunsSortBy_created_at      RunsSortBy = "created_at"
	RunsSortBy_created_at_asc  RunsSortBy = "created_at:asc"
	RunsSortBy_created_at_desc RunsSortBy = "created_at:desc"
)

// BadRequest defines model for BadRequest.
type BadRequest Error

// ApiRunHostsListParams defines parameters for ApiRunHostsList.
type ApiRunHostsListParams struct {

	// Allows for filtering based on various criteria
	Filter *RunHostFilter `json:"filter,omitempty"`

	// Defines fields to be returned in the response.
	Fields *RunHostFields `json:"fields,omitempty"`

	// Maximum number of results to return
	Limit *Limit `json:"limit,omitempty"`

	// Indicates the starting position of the query relative to the complete set of items that match the query
	Offset *Offset `json:"offset,omitempty"`
}

// ApiRunsListParams defines parameters for ApiRunsList.
type ApiRunsListParams struct {

	// Allows for filtering based on various criteria
	Filter *RunsFilter `json:"filter,omitempty"`

	// Defines fields to be returned in the response.
	Fields *RunsFields `json:"fields,omitempty"`

	// Sort order
	SortBy *RunsSortBy `json:"sort_by,omitempty"`

	// Maximum number of results to return
	Limit *Limit `json:"limit,omitempty"`

	// Indicates the starting position of the query relative to the complete set of items that match the query
	Offset *Offset `json:"offset,omitempty"`
}

// ApiInternalRunsCreateJSONBody defines parameters for ApiInternalRunsCreate.
type ApiInternalRunsCreateJSONBody []RunInput

// ApiInternalRunsCreateRequestBody defines body for ApiInternalRunsCreate for application/json ContentType.
type ApiInternalRunsCreateJSONRequestBody ApiInternalRunsCreateJSONBody

// Getter for additional properties for Labels. Returns the specified
// element and whether it was found
func (a Labels) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Labels
func (a *Labels) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Labels to handle AdditionalProperties
func (a *Labels) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Labels to handle AdditionalProperties
func (a Labels) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for RunLabelsNullable. Returns the specified
// element and whether it was found
func (a RunLabelsNullable) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for RunLabelsNullable
func (a *RunLabelsNullable) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for RunLabelsNullable to handle AdditionalProperties
func (a *RunLabelsNullable) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for RunLabelsNullable to handle AdditionalProperties
func (a RunLabelsNullable) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
