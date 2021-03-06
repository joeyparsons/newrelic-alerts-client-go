package alerts

import (
	"fmt"

	"github.com/joeyparsons/newrelic-alerts-client-go/pkg/errors"
)

// PluginsCondition represents an alert condition for New Relic Plugins.
type PluginsCondition struct {
	ID                int             `json:"id,omitempty"`
	Name              string          `json:"name,omitempty"`
	Enabled           bool            `json:"enabled"`
	Entities          []string        `json:"entities,omitempty"`
	Metric            string          `json:"metric,omitempty"`
	MetricDescription string          `json:"metric_description,omitempty"`
	RunbookURL        string          `json:"runbook_url,omitempty"`
	Terms             []ConditionTerm `json:"terms,omitempty"`
	ValueFunction     string          `json:"value_function,omitempty"`
	Plugin            AlertPlugin     `json:"plugin,omitempty"`
}

// AlertPlugin represents a plugin to use with a Plugin alert condition.
type AlertPlugin struct {
	ID   string `json:"id,omitempty"`
	GUID string `json:"guid,omitempty"`
}

// ListPluginsConditions returns alert conditions for New Relic plugins for a given alert policy.
func (a *Alerts) ListPluginsConditions(policyID int) ([]*PluginsCondition, error) {
	conditions := []*PluginsCondition{}
	queryParams := listPluginsConditionsParams{
		PolicyID: policyID,
	}

	nextURL := a.config.Region().RestURL("/alerts_plugins_conditions.json")

	for nextURL != "" {
		response := pluginsConditionsResponse{}
		resp, err := a.client.Get(nextURL, &queryParams, &response)

		if err != nil {
			return nil, err
		}

		conditions = append(conditions, response.PluginsConditions...)

		paging := a.pager.Parse(resp)
		nextURL = paging.Next
	}

	return conditions, nil
}

// GetPluginsCondition gets information about an alert condition for a plugin
// given a policy ID and plugin ID.
func (a *Alerts) GetPluginsCondition(policyID int, pluginID int) (*PluginsCondition, error) {
	conditions, err := a.ListPluginsConditions(policyID)

	if err != nil {
		return nil, err
	}

	for _, condition := range conditions {
		if condition.ID == pluginID {
			return condition, nil
		}
	}

	return nil, errors.NewNotFoundf("no condition found for policy %d and condition ID %d", policyID, pluginID)
}

// CreatePluginsCondition creates an alert condition for a plugin.
func (a *Alerts) CreatePluginsCondition(policyID int, condition PluginsCondition) (*PluginsCondition, error) {
	reqBody := pluginConditionRequestBody{
		PluginsCondition: condition,
	}
	resp := pluginConditionResponse{}

	url := fmt.Sprintf("/alerts_plugins_conditions/policies/%d.json", policyID)
	_, err := a.client.Post(a.config.Region().RestURL(url), nil, &reqBody, &resp)

	if err != nil {
		return nil, err
	}

	return &resp.PluginsCondition, nil
}

// UpdatePluginsCondition updates an alert condition for a plugin.
func (a *Alerts) UpdatePluginsCondition(condition PluginsCondition) (*PluginsCondition, error) {
	reqBody := pluginConditionRequestBody{
		PluginsCondition: condition,
	}
	resp := pluginConditionResponse{}

	url := fmt.Sprintf("/alerts_plugins_conditions/%d.json", condition.ID)
	_, err := a.client.Put(a.config.Region().RestURL(url), nil, &reqBody, &resp)

	if err != nil {
		return nil, err
	}

	return &resp.PluginsCondition, nil
}

// DeletePluginsCondition deletes a plugin alert condition.
func (a *Alerts) DeletePluginsCondition(id int) (*PluginsCondition, error) {
	resp := pluginConditionResponse{}
	url := fmt.Sprintf("/alerts_plugins_conditions/%d.json", id)

	_, err := a.client.Delete(a.config.Region().RestURL(url), nil, &resp)

	if err != nil {
		return nil, err
	}

	return &resp.PluginsCondition, nil
}

type listPluginsConditionsParams struct {
	PolicyID int `url:"policy_id,omitempty"`
}

type pluginsConditionsResponse struct {
	PluginsConditions []*PluginsCondition `json:"plugins_conditions,omitempty"`
}

type pluginConditionResponse struct {
	PluginsCondition PluginsCondition `json:"plugins_condition,omitempty"`
}

type pluginConditionRequestBody struct {
	PluginsCondition PluginsCondition `json:"plugins_condition,omitempty"`
}
