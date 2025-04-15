// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/VapiAI/server-sdk-go/internal"
	time "time"
)

type CreateTestSuiteRunDto struct {
	// This is the name of the test suite run.
	Name *string `json:"name,omitempty" url:"-"`
}

type TestSuiteRunControllerFindAllPaginatedRequest struct {
	// This is the page number to return. Defaults to 1.
	Page *float64 `json:"-" url:"page,omitempty"`
	// This is the sort order for pagination. Defaults to 'DESC'.
	SortOrder *TestSuiteRunControllerFindAllPaginatedRequestSortOrder `json:"-" url:"sortOrder,omitempty"`
	// This is the maximum number of items to return. Defaults to 100.
	Limit *float64 `json:"-" url:"limit,omitempty"`
	// This will return items where the createdAt is greater than the specified value.
	CreatedAtGt *time.Time `json:"-" url:"createdAtGt,omitempty"`
	// This will return items where the createdAt is less than the specified value.
	CreatedAtLt *time.Time `json:"-" url:"createdAtLt,omitempty"`
	// This will return items where the createdAt is greater than or equal to the specified value.
	CreatedAtGe *time.Time `json:"-" url:"createdAtGe,omitempty"`
	// This will return items where the createdAt is less than or equal to the specified value.
	CreatedAtLe *time.Time `json:"-" url:"createdAtLe,omitempty"`
	// This will return items where the updatedAt is greater than the specified value.
	UpdatedAtGt *time.Time `json:"-" url:"updatedAtGt,omitempty"`
	// This will return items where the updatedAt is less than the specified value.
	UpdatedAtLt *time.Time `json:"-" url:"updatedAtLt,omitempty"`
	// This will return items where the updatedAt is greater than or equal to the specified value.
	UpdatedAtGe *time.Time `json:"-" url:"updatedAtGe,omitempty"`
	// This will return items where the updatedAt is less than or equal to the specified value.
	UpdatedAtLe *time.Time `json:"-" url:"updatedAtLe,omitempty"`
}

type UpdateTestSuiteRunDto struct {
	// This is the name of the test suite run.
	Name *string `json:"name,omitempty" url:"-"`
}

type TestSuiteRun struct {
	// This is the current status of the test suite run.
	Status TestSuiteRunStatus `json:"status" url:"status"`
	// This is the unique identifier for the test suite run.
	Id string `json:"id" url:"id"`
	// This is the unique identifier for the organization this run belongs to.
	OrgId string `json:"orgId" url:"orgId"`
	// This is the unique identifier for the test suite this run belongs to.
	TestSuiteId string `json:"testSuiteId" url:"testSuiteId"`
	// This is the ISO 8601 date-time string of when the test suite run was created.
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// This is the ISO 8601 date-time string of when the test suite run was last updated.
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`
	// These are the results of the tests in this test suite run.
	TestResults []*TestSuiteRunTestResult `json:"testResults,omitempty" url:"testResults,omitempty"`
	// This is the name of the test suite run.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRun) GetStatus() TestSuiteRunStatus {
	if t == nil {
		return ""
	}
	return t.Status
}

func (t *TestSuiteRun) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TestSuiteRun) GetOrgId() string {
	if t == nil {
		return ""
	}
	return t.OrgId
}

func (t *TestSuiteRun) GetTestSuiteId() string {
	if t == nil {
		return ""
	}
	return t.TestSuiteId
}

func (t *TestSuiteRun) GetCreatedAt() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.CreatedAt
}

func (t *TestSuiteRun) GetUpdatedAt() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.UpdatedAt
}

func (t *TestSuiteRun) GetTestResults() []*TestSuiteRunTestResult {
	if t == nil {
		return nil
	}
	return t.TestResults
}

func (t *TestSuiteRun) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TestSuiteRun) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRun) UnmarshalJSON(data []byte) error {
	type embed TestSuiteRun
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TestSuiteRun(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.Time()
	t.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRun) MarshalJSON() ([]byte, error) {
	type embed TestSuiteRun
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*t),
		CreatedAt: internal.NewDateTime(t.CreatedAt),
		UpdatedAt: internal.NewDateTime(t.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (t *TestSuiteRun) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunScorerAi struct {
	// This is the type of the scorer, which must be AI.
	// This is the result of the test suite.
	Result TestSuiteRunScorerAiResult `json:"result" url:"result"`
	// This is the reasoning provided by the AI scorer.
	Reasoning string `json:"reasoning" url:"reasoning"`
	// This is the rubric used by the AI scorer.
	Rubric string `json:"rubric" url:"rubric"`
	type_  string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunScorerAi) GetResult() TestSuiteRunScorerAiResult {
	if t == nil {
		return ""
	}
	return t.Result
}

func (t *TestSuiteRunScorerAi) GetReasoning() string {
	if t == nil {
		return ""
	}
	return t.Reasoning
}

func (t *TestSuiteRunScorerAi) GetRubric() string {
	if t == nil {
		return ""
	}
	return t.Rubric
}

func (t *TestSuiteRunScorerAi) Type() string {
	return t.type_
}

func (t *TestSuiteRunScorerAi) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunScorerAi) UnmarshalJSON(data []byte) error {
	type embed TestSuiteRunScorerAi
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TestSuiteRunScorerAi(unmarshaler.embed)
	if unmarshaler.Type != "ai" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "ai", unmarshaler.Type)
	}
	t.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunScorerAi) MarshalJSON() ([]byte, error) {
	type embed TestSuiteRunScorerAi
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "ai",
	}
	return json.Marshal(marshaler)
}

func (t *TestSuiteRunScorerAi) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

// This is the result of the test suite.
type TestSuiteRunScorerAiResult string

const (
	TestSuiteRunScorerAiResultPass TestSuiteRunScorerAiResult = "pass"
	TestSuiteRunScorerAiResultFail TestSuiteRunScorerAiResult = "fail"
)

func NewTestSuiteRunScorerAiResultFromString(s string) (TestSuiteRunScorerAiResult, error) {
	switch s {
	case "pass":
		return TestSuiteRunScorerAiResultPass, nil
	case "fail":
		return TestSuiteRunScorerAiResultFail, nil
	}
	var t TestSuiteRunScorerAiResult
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TestSuiteRunScorerAiResult) Ptr() *TestSuiteRunScorerAiResult {
	return &t
}

// This is the current status of the test suite run.
type TestSuiteRunStatus string

const (
	TestSuiteRunStatusQueued     TestSuiteRunStatus = "queued"
	TestSuiteRunStatusInProgress TestSuiteRunStatus = "in-progress"
	TestSuiteRunStatusCompleted  TestSuiteRunStatus = "completed"
	TestSuiteRunStatusFailed     TestSuiteRunStatus = "failed"
)

func NewTestSuiteRunStatusFromString(s string) (TestSuiteRunStatus, error) {
	switch s {
	case "queued":
		return TestSuiteRunStatusQueued, nil
	case "in-progress":
		return TestSuiteRunStatusInProgress, nil
	case "completed":
		return TestSuiteRunStatusCompleted, nil
	case "failed":
		return TestSuiteRunStatusFailed, nil
	}
	var t TestSuiteRunStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TestSuiteRunStatus) Ptr() *TestSuiteRunStatus {
	return &t
}

type TestSuiteRunTestAttempt struct {
	// These are the results of the scorers used to evaluate the test attempt.
	ScorerResults []*TestSuiteRunScorerAi `json:"scorerResults,omitempty" url:"scorerResults,omitempty"`
	// This is the call made during the test attempt.
	Call *TestSuiteRunTestAttemptCall `json:"call,omitempty" url:"call,omitempty"`
	// This is the call ID for the test attempt.
	CallId *string `json:"callId,omitempty" url:"callId,omitempty"`
	// This is the metadata for the test attempt.
	Metadata *TestSuiteRunTestAttemptMetadata `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunTestAttempt) GetScorerResults() []*TestSuiteRunScorerAi {
	if t == nil {
		return nil
	}
	return t.ScorerResults
}

func (t *TestSuiteRunTestAttempt) GetCall() *TestSuiteRunTestAttemptCall {
	if t == nil {
		return nil
	}
	return t.Call
}

func (t *TestSuiteRunTestAttempt) GetCallId() *string {
	if t == nil {
		return nil
	}
	return t.CallId
}

func (t *TestSuiteRunTestAttempt) GetMetadata() *TestSuiteRunTestAttemptMetadata {
	if t == nil {
		return nil
	}
	return t.Metadata
}

func (t *TestSuiteRunTestAttempt) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunTestAttempt) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteRunTestAttempt
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteRunTestAttempt(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunTestAttempt) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunTestAttemptCall struct {
	// This is the artifact of the call.
	Artifact *Artifact `json:"artifact,omitempty" url:"artifact,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunTestAttemptCall) GetArtifact() *Artifact {
	if t == nil {
		return nil
	}
	return t.Artifact
}

func (t *TestSuiteRunTestAttemptCall) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunTestAttemptCall) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteRunTestAttemptCall
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteRunTestAttemptCall(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunTestAttemptCall) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunTestAttemptMetadata struct {
	// This is the session ID for the test attempt.
	SessionId string `json:"sessionId" url:"sessionId"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunTestAttemptMetadata) GetSessionId() string {
	if t == nil {
		return ""
	}
	return t.SessionId
}

func (t *TestSuiteRunTestAttemptMetadata) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunTestAttemptMetadata) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteRunTestAttemptMetadata
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteRunTestAttemptMetadata(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunTestAttemptMetadata) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunTestResult struct {
	// This is the test that was run.
	Test *TestSuiteTestVoice `json:"test,omitempty" url:"test,omitempty"`
	// These are the attempts made for this test.
	Attempts []*TestSuiteRunTestAttempt `json:"attempts,omitempty" url:"attempts,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunTestResult) GetTest() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.Test
}

func (t *TestSuiteRunTestResult) GetAttempts() []*TestSuiteRunTestAttempt {
	if t == nil {
		return nil
	}
	return t.Attempts
}

func (t *TestSuiteRunTestResult) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunTestResult) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteRunTestResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteRunTestResult(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunTestResult) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunsPaginatedResponse struct {
	Results  []*TestSuiteRun `json:"results,omitempty" url:"results,omitempty"`
	Metadata *PaginationMeta `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteRunsPaginatedResponse) GetResults() []*TestSuiteRun {
	if t == nil {
		return nil
	}
	return t.Results
}

func (t *TestSuiteRunsPaginatedResponse) GetMetadata() *PaginationMeta {
	if t == nil {
		return nil
	}
	return t.Metadata
}

func (t *TestSuiteRunsPaginatedResponse) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteRunsPaginatedResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteRunsPaginatedResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteRunsPaginatedResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteRunsPaginatedResponse) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TestSuiteRunControllerFindAllPaginatedRequestSortOrder string

const (
	TestSuiteRunControllerFindAllPaginatedRequestSortOrderAsc  TestSuiteRunControllerFindAllPaginatedRequestSortOrder = "ASC"
	TestSuiteRunControllerFindAllPaginatedRequestSortOrderDesc TestSuiteRunControllerFindAllPaginatedRequestSortOrder = "DESC"
)

func NewTestSuiteRunControllerFindAllPaginatedRequestSortOrderFromString(s string) (TestSuiteRunControllerFindAllPaginatedRequestSortOrder, error) {
	switch s {
	case "ASC":
		return TestSuiteRunControllerFindAllPaginatedRequestSortOrderAsc, nil
	case "DESC":
		return TestSuiteRunControllerFindAllPaginatedRequestSortOrderDesc, nil
	}
	var t TestSuiteRunControllerFindAllPaginatedRequestSortOrder
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TestSuiteRunControllerFindAllPaginatedRequestSortOrder) Ptr() *TestSuiteRunControllerFindAllPaginatedRequestSortOrder {
	return &t
}
