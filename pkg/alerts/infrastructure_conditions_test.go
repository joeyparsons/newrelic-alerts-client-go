// +build unit

package alerts

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testInfrastructureConditionPolicyId      = 111111
	testInfrastructureCriticalThresholdValue = 12.3
	testInfrastructureCriticalThreshold      = InfrastructureConditionThreshold{
		Duration: 6,
		Function: "all",
		Value:    &testInfrastructureCriticalThresholdValue,
	}
	testInfrastructureWarningThresholdValue = float64(10)
	testInfrastructureWarningThreshold      = InfrastructureConditionThreshold{
		Duration: 6,
		Function: "all",
		Value:    &testInfrastructureWarningThresholdValue,
	}

	testInfrastructureCondition = InfrastructureCondition{
		Comparison:   "equal",
		CreatedAt:    &testTimestamp,
		Critical:     &testInfrastructureCriticalThreshold,
		Enabled:      true,
		ID:           13890,
		Name:         "Java is running",
		PolicyID:     testInfrastructureConditionPolicyId,
		ProcessWhere: "(commandName = 'java')",
		Type:         "infra_process_running",
		UpdatedAt:    &testTimestamp,
		Warning:      &testInfrastructureWarningThreshold,
		Where:        "(hostname LIKE '%cassandra%')",
	}
	testInfrastructureConditionJson = `
		{
			"type":"infra_process_running",
			"name":"Java is running",
			"enabled":true,
			"where_clause":"(hostname LIKE '%cassandra%')",
			"id":13890,
			"created_at_epoch_millis":` + testTimestampStringMs + `,
			"updated_at_epoch_millis":` + testTimestampStringMs + `,
			"policy_id": 111111,
			"comparison":"equal",
			"critical_threshold":{
				"value":12.3,
				"duration_minutes":6,
				"time_function": "all"
			},
			"warning_threshold": {
				"value": 10,
				"duration_minutes": 6,
				"time_function": "all"
			},
			"process_where_clause":"(commandName = 'java')"
		}`
)

func TestListInfrastructureConditions(t *testing.T) {
	t.Parallel()
	respJSON := fmt.Sprintf(`{ "data":[%s] }`, testInfrastructureConditionJson)
	alerts := newMockResponse(t, respJSON, http.StatusOK)

	expected := []InfrastructureCondition{testInfrastructureCondition}

	actual, err := alerts.ListInfrastructureConditions(testInfrastructureConditionPolicyId)

	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected, actual)
}

func TestGetInfrastructureConditions(t *testing.T) {
	t.Parallel()
	respJSON := fmt.Sprintf(`{ "data":%s }`, testInfrastructureConditionJson)
	alerts := newMockResponse(t, respJSON, http.StatusOK)

	expected := &testInfrastructureCondition

	actual, err := alerts.GetInfrastructureCondition(testInfrastructureCondition.ID)

	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected, actual)
}

func TestCreateInfrastructureConditions(t *testing.T) {
	t.Parallel()
	respJSON := fmt.Sprintf(`{ "data":%s }`, testInfrastructureConditionJson)
	alerts := newMockResponse(t, respJSON, http.StatusOK)

	expected := &testInfrastructureCondition

	actual, err := alerts.CreateInfrastructureCondition(testInfrastructureCondition)

	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected, actual)
}

func TestUpdateInfrastructureConditions(t *testing.T) {
	t.Parallel()
	respJSON := fmt.Sprintf(`{ "data":%s }`, testInfrastructureConditionJson)
	alerts := newMockResponse(t, respJSON, http.StatusOK)

	expected := &testInfrastructureCondition

	actual, err := alerts.UpdateInfrastructureCondition(testInfrastructureCondition)

	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected, actual)
}

func TestDeleteInfrastructureConditions(t *testing.T) {
	t.Parallel()
	respJSON := fmt.Sprintf(`{ "data":%s }`, testInfrastructureConditionJson)
	alerts := newMockResponse(t, respJSON, http.StatusOK)

	err := alerts.DeleteInfrastructureCondition(testInfrastructureCondition.ID)

	require.NoError(t, err)
}
