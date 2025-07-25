// Code generated by Fern. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/VapiAI/server-sdk-go/internal"
	time "time"
)

type TestSuiteTestControllerFindAllPaginatedRequest struct {
	// This is the page number to return. Defaults to 1.
	Page *float64 `json:"-" url:"page,omitempty"`
	// This is the sort order for pagination. Defaults to 'DESC'.
	SortOrder *TestSuiteTestControllerFindAllPaginatedRequestSortOrder `json:"-" url:"sortOrder,omitempty"`
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

type CreateTestSuiteTestChatDto struct {
	// These are the scorers used to evaluate the test.
	Scorers []*TestSuiteTestScorerAi `json:"scorers" url:"scorers"`
	// This is the script to be used for the chat test.
	Script string `json:"script" url:"script"`
	// This is the number of attempts allowed for the test.
	NumAttempts *float64 `json:"numAttempts,omitempty" url:"numAttempts,omitempty"`
	// This is the name of the test.
	Name  *string `json:"name,omitempty" url:"name,omitempty"`
	type_ string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateTestSuiteTestChatDto) GetScorers() []*TestSuiteTestScorerAi {
	if c == nil {
		return nil
	}
	return c.Scorers
}

func (c *CreateTestSuiteTestChatDto) GetScript() string {
	if c == nil {
		return ""
	}
	return c.Script
}

func (c *CreateTestSuiteTestChatDto) GetNumAttempts() *float64 {
	if c == nil {
		return nil
	}
	return c.NumAttempts
}

func (c *CreateTestSuiteTestChatDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTestSuiteTestChatDto) Type() string {
	return c.type_
}

func (c *CreateTestSuiteTestChatDto) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateTestSuiteTestChatDto) UnmarshalJSON(data []byte) error {
	type embed CreateTestSuiteTestChatDto
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateTestSuiteTestChatDto(unmarshaler.embed)
	if unmarshaler.Type != "chat" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "chat", unmarshaler.Type)
	}
	c.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateTestSuiteTestChatDto) MarshalJSON() ([]byte, error) {
	type embed CreateTestSuiteTestChatDto
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "chat",
	}
	return json.Marshal(marshaler)
}

func (c *CreateTestSuiteTestChatDto) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateTestSuiteTestVoiceDto struct {
	// These are the scorers used to evaluate the test.
	Scorers []*TestSuiteTestScorerAi `json:"scorers" url:"scorers"`
	// This is the script to be used for the voice test.
	Script string `json:"script" url:"script"`
	// This is the number of attempts allowed for the test.
	NumAttempts *float64 `json:"numAttempts,omitempty" url:"numAttempts,omitempty"`
	// This is the name of the test.
	Name  *string `json:"name,omitempty" url:"name,omitempty"`
	type_ string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateTestSuiteTestVoiceDto) GetScorers() []*TestSuiteTestScorerAi {
	if c == nil {
		return nil
	}
	return c.Scorers
}

func (c *CreateTestSuiteTestVoiceDto) GetScript() string {
	if c == nil {
		return ""
	}
	return c.Script
}

func (c *CreateTestSuiteTestVoiceDto) GetNumAttempts() *float64 {
	if c == nil {
		return nil
	}
	return c.NumAttempts
}

func (c *CreateTestSuiteTestVoiceDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTestSuiteTestVoiceDto) Type() string {
	return c.type_
}

func (c *CreateTestSuiteTestVoiceDto) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateTestSuiteTestVoiceDto) UnmarshalJSON(data []byte) error {
	type embed CreateTestSuiteTestVoiceDto
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateTestSuiteTestVoiceDto(unmarshaler.embed)
	if unmarshaler.Type != "voice" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "voice", unmarshaler.Type)
	}
	c.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateTestSuiteTestVoiceDto) MarshalJSON() ([]byte, error) {
	type embed CreateTestSuiteTestVoiceDto
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "voice",
	}
	return json.Marshal(marshaler)
}

func (c *CreateTestSuiteTestVoiceDto) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type TestSuiteTestChat struct {
	// These are the scorers used to evaluate the test.
	Scorers []*TestSuiteTestScorerAi `json:"scorers" url:"scorers"`
	// This is the unique identifier for the test.
	Id string `json:"id" url:"id"`
	// This is the unique identifier for the test suite this test belongs to.
	TestSuiteId string `json:"testSuiteId" url:"testSuiteId"`
	// This is the unique identifier for the organization this test belongs to.
	OrgId string `json:"orgId" url:"orgId"`
	// This is the ISO 8601 date-time string of when the test was created.
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// This is the ISO 8601 date-time string of when the test was last updated.
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`
	// This is the name of the test.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// This is the script to be used for the chat test.
	Script string `json:"script" url:"script"`
	// This is the number of attempts allowed for the test.
	NumAttempts *float64 `json:"numAttempts,omitempty" url:"numAttempts,omitempty"`
	type_       string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteTestChat) GetScorers() []*TestSuiteTestScorerAi {
	if t == nil {
		return nil
	}
	return t.Scorers
}

func (t *TestSuiteTestChat) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TestSuiteTestChat) GetTestSuiteId() string {
	if t == nil {
		return ""
	}
	return t.TestSuiteId
}

func (t *TestSuiteTestChat) GetOrgId() string {
	if t == nil {
		return ""
	}
	return t.OrgId
}

func (t *TestSuiteTestChat) GetCreatedAt() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.CreatedAt
}

func (t *TestSuiteTestChat) GetUpdatedAt() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.UpdatedAt
}

func (t *TestSuiteTestChat) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TestSuiteTestChat) GetScript() string {
	if t == nil {
		return ""
	}
	return t.Script
}

func (t *TestSuiteTestChat) GetNumAttempts() *float64 {
	if t == nil {
		return nil
	}
	return t.NumAttempts
}

func (t *TestSuiteTestChat) Type() string {
	return t.type_
}

func (t *TestSuiteTestChat) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteTestChat) UnmarshalJSON(data []byte) error {
	type embed TestSuiteTestChat
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
		Type      string             `json:"type"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TestSuiteTestChat(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.Time()
	t.UpdatedAt = unmarshaler.UpdatedAt.Time()
	if unmarshaler.Type != "chat" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "chat", unmarshaler.Type)
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

func (t *TestSuiteTestChat) MarshalJSON() ([]byte, error) {
	type embed TestSuiteTestChat
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
		Type      string             `json:"type"`
	}{
		embed:     embed(*t),
		CreatedAt: internal.NewDateTime(t.CreatedAt),
		UpdatedAt: internal.NewDateTime(t.UpdatedAt),
		Type:      "chat",
	}
	return json.Marshal(marshaler)
}

func (t *TestSuiteTestChat) String() string {
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

type TestSuiteTestsPaginatedResponse struct {
	// A list of test suite tests.
	Results []*TestSuiteTestsPaginatedResponseResultsItem `json:"results" url:"results"`
	// Metadata about the pagination.
	Metadata *PaginationMeta `json:"metadata" url:"metadata"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TestSuiteTestsPaginatedResponse) GetResults() []*TestSuiteTestsPaginatedResponseResultsItem {
	if t == nil {
		return nil
	}
	return t.Results
}

func (t *TestSuiteTestsPaginatedResponse) GetMetadata() *PaginationMeta {
	if t == nil {
		return nil
	}
	return t.Metadata
}

func (t *TestSuiteTestsPaginatedResponse) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TestSuiteTestsPaginatedResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler TestSuiteTestsPaginatedResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TestSuiteTestsPaginatedResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TestSuiteTestsPaginatedResponse) String() string {
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

type TestSuiteTestsPaginatedResponseResultsItem struct {
	TestSuiteTestVoice *TestSuiteTestVoice
	TestSuiteTestChat  *TestSuiteTestChat

	typ string
}

func (t *TestSuiteTestsPaginatedResponseResultsItem) GetTestSuiteTestVoice() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestVoice
}

func (t *TestSuiteTestsPaginatedResponseResultsItem) GetTestSuiteTestChat() *TestSuiteTestChat {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestChat
}

func (t *TestSuiteTestsPaginatedResponseResultsItem) UnmarshalJSON(data []byte) error {
	valueTestSuiteTestVoice := new(TestSuiteTestVoice)
	if err := json.Unmarshal(data, &valueTestSuiteTestVoice); err == nil {
		t.typ = "TestSuiteTestVoice"
		t.TestSuiteTestVoice = valueTestSuiteTestVoice
		return nil
	}
	valueTestSuiteTestChat := new(TestSuiteTestChat)
	if err := json.Unmarshal(data, &valueTestSuiteTestChat); err == nil {
		t.typ = "TestSuiteTestChat"
		t.TestSuiteTestChat = valueTestSuiteTestChat
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestsPaginatedResponseResultsItem) MarshalJSON() ([]byte, error) {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return json.Marshal(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return json.Marshal(t.TestSuiteTestChat)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestsPaginatedResponseResultsItemVisitor interface {
	VisitTestSuiteTestVoice(*TestSuiteTestVoice) error
	VisitTestSuiteTestChat(*TestSuiteTestChat) error
}

func (t *TestSuiteTestsPaginatedResponseResultsItem) Accept(visitor TestSuiteTestsPaginatedResponseResultsItemVisitor) error {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return visitor.VisitTestSuiteTestVoice(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return visitor.VisitTestSuiteTestChat(t.TestSuiteTestChat)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type UpdateTestSuiteTestChatDto struct {
	// These are the scorers used to evaluate the test.
	Scorers []*TestSuiteTestScorerAi `json:"scorers,omitempty" url:"scorers,omitempty"`
	// This is the name of the test.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// This is the script to be used for the chat test.
	Script *string `json:"script,omitempty" url:"script,omitempty"`
	// This is the number of attempts allowed for the test.
	NumAttempts *float64 `json:"numAttempts,omitempty" url:"numAttempts,omitempty"`
	type_       string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateTestSuiteTestChatDto) GetScorers() []*TestSuiteTestScorerAi {
	if u == nil {
		return nil
	}
	return u.Scorers
}

func (u *UpdateTestSuiteTestChatDto) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateTestSuiteTestChatDto) GetScript() *string {
	if u == nil {
		return nil
	}
	return u.Script
}

func (u *UpdateTestSuiteTestChatDto) GetNumAttempts() *float64 {
	if u == nil {
		return nil
	}
	return u.NumAttempts
}

func (u *UpdateTestSuiteTestChatDto) Type() string {
	return u.type_
}

func (u *UpdateTestSuiteTestChatDto) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateTestSuiteTestChatDto) UnmarshalJSON(data []byte) error {
	type embed UpdateTestSuiteTestChatDto
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = UpdateTestSuiteTestChatDto(unmarshaler.embed)
	if unmarshaler.Type != "chat" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", u, "chat", unmarshaler.Type)
	}
	u.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateTestSuiteTestChatDto) MarshalJSON() ([]byte, error) {
	type embed UpdateTestSuiteTestChatDto
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
		Type:  "chat",
	}
	return json.Marshal(marshaler)
}

func (u *UpdateTestSuiteTestChatDto) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateTestSuiteTestVoiceDto struct {
	// These are the scorers used to evaluate the test.
	Scorers []*TestSuiteTestScorerAi `json:"scorers,omitempty" url:"scorers,omitempty"`
	// This is the name of the test.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// This is the script to be used for the voice test.
	Script *string `json:"script,omitempty" url:"script,omitempty"`
	// This is the number of attempts allowed for the test.
	NumAttempts *float64 `json:"numAttempts,omitempty" url:"numAttempts,omitempty"`
	type_       string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateTestSuiteTestVoiceDto) GetScorers() []*TestSuiteTestScorerAi {
	if u == nil {
		return nil
	}
	return u.Scorers
}

func (u *UpdateTestSuiteTestVoiceDto) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateTestSuiteTestVoiceDto) GetScript() *string {
	if u == nil {
		return nil
	}
	return u.Script
}

func (u *UpdateTestSuiteTestVoiceDto) GetNumAttempts() *float64 {
	if u == nil {
		return nil
	}
	return u.NumAttempts
}

func (u *UpdateTestSuiteTestVoiceDto) Type() string {
	return u.type_
}

func (u *UpdateTestSuiteTestVoiceDto) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateTestSuiteTestVoiceDto) UnmarshalJSON(data []byte) error {
	type embed UpdateTestSuiteTestVoiceDto
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = UpdateTestSuiteTestVoiceDto(unmarshaler.embed)
	if unmarshaler.Type != "voice" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", u, "voice", unmarshaler.Type)
	}
	u.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateTestSuiteTestVoiceDto) MarshalJSON() ([]byte, error) {
	type embed UpdateTestSuiteTestVoiceDto
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*u),
		Type:  "voice",
	}
	return json.Marshal(marshaler)
}

func (u *UpdateTestSuiteTestVoiceDto) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type TestSuiteTestControllerCreateRequest struct {
	CreateTestSuiteTestVoiceDto *CreateTestSuiteTestVoiceDto
	CreateTestSuiteTestChatDto  *CreateTestSuiteTestChatDto

	typ string
}

func (t *TestSuiteTestControllerCreateRequest) GetCreateTestSuiteTestVoiceDto() *CreateTestSuiteTestVoiceDto {
	if t == nil {
		return nil
	}
	return t.CreateTestSuiteTestVoiceDto
}

func (t *TestSuiteTestControllerCreateRequest) GetCreateTestSuiteTestChatDto() *CreateTestSuiteTestChatDto {
	if t == nil {
		return nil
	}
	return t.CreateTestSuiteTestChatDto
}

func (t *TestSuiteTestControllerCreateRequest) UnmarshalJSON(data []byte) error {
	valueCreateTestSuiteTestVoiceDto := new(CreateTestSuiteTestVoiceDto)
	if err := json.Unmarshal(data, &valueCreateTestSuiteTestVoiceDto); err == nil {
		t.typ = "CreateTestSuiteTestVoiceDto"
		t.CreateTestSuiteTestVoiceDto = valueCreateTestSuiteTestVoiceDto
		return nil
	}
	valueCreateTestSuiteTestChatDto := new(CreateTestSuiteTestChatDto)
	if err := json.Unmarshal(data, &valueCreateTestSuiteTestChatDto); err == nil {
		t.typ = "CreateTestSuiteTestChatDto"
		t.CreateTestSuiteTestChatDto = valueCreateTestSuiteTestChatDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerCreateRequest) MarshalJSON() ([]byte, error) {
	if t.typ == "CreateTestSuiteTestVoiceDto" || t.CreateTestSuiteTestVoiceDto != nil {
		return json.Marshal(t.CreateTestSuiteTestVoiceDto)
	}
	if t.typ == "CreateTestSuiteTestChatDto" || t.CreateTestSuiteTestChatDto != nil {
		return json.Marshal(t.CreateTestSuiteTestChatDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerCreateRequestVisitor interface {
	VisitCreateTestSuiteTestVoiceDto(*CreateTestSuiteTestVoiceDto) error
	VisitCreateTestSuiteTestChatDto(*CreateTestSuiteTestChatDto) error
}

func (t *TestSuiteTestControllerCreateRequest) Accept(visitor TestSuiteTestControllerCreateRequestVisitor) error {
	if t.typ == "CreateTestSuiteTestVoiceDto" || t.CreateTestSuiteTestVoiceDto != nil {
		return visitor.VisitCreateTestSuiteTestVoiceDto(t.CreateTestSuiteTestVoiceDto)
	}
	if t.typ == "CreateTestSuiteTestChatDto" || t.CreateTestSuiteTestChatDto != nil {
		return visitor.VisitCreateTestSuiteTestChatDto(t.CreateTestSuiteTestChatDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerCreateResponse struct {
	TestSuiteTestVoice *TestSuiteTestVoice
	TestSuiteTestChat  *TestSuiteTestChat

	typ string
}

func (t *TestSuiteTestControllerCreateResponse) GetTestSuiteTestVoice() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestVoice
}

func (t *TestSuiteTestControllerCreateResponse) GetTestSuiteTestChat() *TestSuiteTestChat {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestChat
}

func (t *TestSuiteTestControllerCreateResponse) UnmarshalJSON(data []byte) error {
	valueTestSuiteTestVoice := new(TestSuiteTestVoice)
	if err := json.Unmarshal(data, &valueTestSuiteTestVoice); err == nil {
		t.typ = "TestSuiteTestVoice"
		t.TestSuiteTestVoice = valueTestSuiteTestVoice
		return nil
	}
	valueTestSuiteTestChat := new(TestSuiteTestChat)
	if err := json.Unmarshal(data, &valueTestSuiteTestChat); err == nil {
		t.typ = "TestSuiteTestChat"
		t.TestSuiteTestChat = valueTestSuiteTestChat
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerCreateResponse) MarshalJSON() ([]byte, error) {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return json.Marshal(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return json.Marshal(t.TestSuiteTestChat)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerCreateResponseVisitor interface {
	VisitTestSuiteTestVoice(*TestSuiteTestVoice) error
	VisitTestSuiteTestChat(*TestSuiteTestChat) error
}

func (t *TestSuiteTestControllerCreateResponse) Accept(visitor TestSuiteTestControllerCreateResponseVisitor) error {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return visitor.VisitTestSuiteTestVoice(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return visitor.VisitTestSuiteTestChat(t.TestSuiteTestChat)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerFindAllPaginatedRequestSortOrder string

const (
	TestSuiteTestControllerFindAllPaginatedRequestSortOrderAsc  TestSuiteTestControllerFindAllPaginatedRequestSortOrder = "ASC"
	TestSuiteTestControllerFindAllPaginatedRequestSortOrderDesc TestSuiteTestControllerFindAllPaginatedRequestSortOrder = "DESC"
)

func NewTestSuiteTestControllerFindAllPaginatedRequestSortOrderFromString(s string) (TestSuiteTestControllerFindAllPaginatedRequestSortOrder, error) {
	switch s {
	case "ASC":
		return TestSuiteTestControllerFindAllPaginatedRequestSortOrderAsc, nil
	case "DESC":
		return TestSuiteTestControllerFindAllPaginatedRequestSortOrderDesc, nil
	}
	var t TestSuiteTestControllerFindAllPaginatedRequestSortOrder
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TestSuiteTestControllerFindAllPaginatedRequestSortOrder) Ptr() *TestSuiteTestControllerFindAllPaginatedRequestSortOrder {
	return &t
}

type TestSuiteTestControllerFindOneResponse struct {
	TestSuiteTestVoice *TestSuiteTestVoice
	TestSuiteTestChat  *TestSuiteTestChat

	typ string
}

func (t *TestSuiteTestControllerFindOneResponse) GetTestSuiteTestVoice() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestVoice
}

func (t *TestSuiteTestControllerFindOneResponse) GetTestSuiteTestChat() *TestSuiteTestChat {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestChat
}

func (t *TestSuiteTestControllerFindOneResponse) UnmarshalJSON(data []byte) error {
	valueTestSuiteTestVoice := new(TestSuiteTestVoice)
	if err := json.Unmarshal(data, &valueTestSuiteTestVoice); err == nil {
		t.typ = "TestSuiteTestVoice"
		t.TestSuiteTestVoice = valueTestSuiteTestVoice
		return nil
	}
	valueTestSuiteTestChat := new(TestSuiteTestChat)
	if err := json.Unmarshal(data, &valueTestSuiteTestChat); err == nil {
		t.typ = "TestSuiteTestChat"
		t.TestSuiteTestChat = valueTestSuiteTestChat
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerFindOneResponse) MarshalJSON() ([]byte, error) {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return json.Marshal(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return json.Marshal(t.TestSuiteTestChat)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerFindOneResponseVisitor interface {
	VisitTestSuiteTestVoice(*TestSuiteTestVoice) error
	VisitTestSuiteTestChat(*TestSuiteTestChat) error
}

func (t *TestSuiteTestControllerFindOneResponse) Accept(visitor TestSuiteTestControllerFindOneResponseVisitor) error {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return visitor.VisitTestSuiteTestVoice(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return visitor.VisitTestSuiteTestChat(t.TestSuiteTestChat)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerRemoveResponse struct {
	TestSuiteTestVoice *TestSuiteTestVoice
	TestSuiteTestChat  *TestSuiteTestChat

	typ string
}

func (t *TestSuiteTestControllerRemoveResponse) GetTestSuiteTestVoice() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestVoice
}

func (t *TestSuiteTestControllerRemoveResponse) GetTestSuiteTestChat() *TestSuiteTestChat {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestChat
}

func (t *TestSuiteTestControllerRemoveResponse) UnmarshalJSON(data []byte) error {
	valueTestSuiteTestVoice := new(TestSuiteTestVoice)
	if err := json.Unmarshal(data, &valueTestSuiteTestVoice); err == nil {
		t.typ = "TestSuiteTestVoice"
		t.TestSuiteTestVoice = valueTestSuiteTestVoice
		return nil
	}
	valueTestSuiteTestChat := new(TestSuiteTestChat)
	if err := json.Unmarshal(data, &valueTestSuiteTestChat); err == nil {
		t.typ = "TestSuiteTestChat"
		t.TestSuiteTestChat = valueTestSuiteTestChat
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerRemoveResponse) MarshalJSON() ([]byte, error) {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return json.Marshal(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return json.Marshal(t.TestSuiteTestChat)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerRemoveResponseVisitor interface {
	VisitTestSuiteTestVoice(*TestSuiteTestVoice) error
	VisitTestSuiteTestChat(*TestSuiteTestChat) error
}

func (t *TestSuiteTestControllerRemoveResponse) Accept(visitor TestSuiteTestControllerRemoveResponseVisitor) error {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return visitor.VisitTestSuiteTestVoice(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return visitor.VisitTestSuiteTestChat(t.TestSuiteTestChat)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerUpdateRequest struct {
	UpdateTestSuiteTestVoiceDto *UpdateTestSuiteTestVoiceDto
	UpdateTestSuiteTestChatDto  *UpdateTestSuiteTestChatDto

	typ string
}

func (t *TestSuiteTestControllerUpdateRequest) GetUpdateTestSuiteTestVoiceDto() *UpdateTestSuiteTestVoiceDto {
	if t == nil {
		return nil
	}
	return t.UpdateTestSuiteTestVoiceDto
}

func (t *TestSuiteTestControllerUpdateRequest) GetUpdateTestSuiteTestChatDto() *UpdateTestSuiteTestChatDto {
	if t == nil {
		return nil
	}
	return t.UpdateTestSuiteTestChatDto
}

func (t *TestSuiteTestControllerUpdateRequest) UnmarshalJSON(data []byte) error {
	valueUpdateTestSuiteTestVoiceDto := new(UpdateTestSuiteTestVoiceDto)
	if err := json.Unmarshal(data, &valueUpdateTestSuiteTestVoiceDto); err == nil {
		t.typ = "UpdateTestSuiteTestVoiceDto"
		t.UpdateTestSuiteTestVoiceDto = valueUpdateTestSuiteTestVoiceDto
		return nil
	}
	valueUpdateTestSuiteTestChatDto := new(UpdateTestSuiteTestChatDto)
	if err := json.Unmarshal(data, &valueUpdateTestSuiteTestChatDto); err == nil {
		t.typ = "UpdateTestSuiteTestChatDto"
		t.UpdateTestSuiteTestChatDto = valueUpdateTestSuiteTestChatDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerUpdateRequest) MarshalJSON() ([]byte, error) {
	if t.typ == "UpdateTestSuiteTestVoiceDto" || t.UpdateTestSuiteTestVoiceDto != nil {
		return json.Marshal(t.UpdateTestSuiteTestVoiceDto)
	}
	if t.typ == "UpdateTestSuiteTestChatDto" || t.UpdateTestSuiteTestChatDto != nil {
		return json.Marshal(t.UpdateTestSuiteTestChatDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerUpdateRequestVisitor interface {
	VisitUpdateTestSuiteTestVoiceDto(*UpdateTestSuiteTestVoiceDto) error
	VisitUpdateTestSuiteTestChatDto(*UpdateTestSuiteTestChatDto) error
}

func (t *TestSuiteTestControllerUpdateRequest) Accept(visitor TestSuiteTestControllerUpdateRequestVisitor) error {
	if t.typ == "UpdateTestSuiteTestVoiceDto" || t.UpdateTestSuiteTestVoiceDto != nil {
		return visitor.VisitUpdateTestSuiteTestVoiceDto(t.UpdateTestSuiteTestVoiceDto)
	}
	if t.typ == "UpdateTestSuiteTestChatDto" || t.UpdateTestSuiteTestChatDto != nil {
		return visitor.VisitUpdateTestSuiteTestChatDto(t.UpdateTestSuiteTestChatDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerUpdateResponse struct {
	TestSuiteTestVoice *TestSuiteTestVoice
	TestSuiteTestChat  *TestSuiteTestChat

	typ string
}

func (t *TestSuiteTestControllerUpdateResponse) GetTestSuiteTestVoice() *TestSuiteTestVoice {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestVoice
}

func (t *TestSuiteTestControllerUpdateResponse) GetTestSuiteTestChat() *TestSuiteTestChat {
	if t == nil {
		return nil
	}
	return t.TestSuiteTestChat
}

func (t *TestSuiteTestControllerUpdateResponse) UnmarshalJSON(data []byte) error {
	valueTestSuiteTestVoice := new(TestSuiteTestVoice)
	if err := json.Unmarshal(data, &valueTestSuiteTestVoice); err == nil {
		t.typ = "TestSuiteTestVoice"
		t.TestSuiteTestVoice = valueTestSuiteTestVoice
		return nil
	}
	valueTestSuiteTestChat := new(TestSuiteTestChat)
	if err := json.Unmarshal(data, &valueTestSuiteTestChat); err == nil {
		t.typ = "TestSuiteTestChat"
		t.TestSuiteTestChat = valueTestSuiteTestChat
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t TestSuiteTestControllerUpdateResponse) MarshalJSON() ([]byte, error) {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return json.Marshal(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return json.Marshal(t.TestSuiteTestChat)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type TestSuiteTestControllerUpdateResponseVisitor interface {
	VisitTestSuiteTestVoice(*TestSuiteTestVoice) error
	VisitTestSuiteTestChat(*TestSuiteTestChat) error
}

func (t *TestSuiteTestControllerUpdateResponse) Accept(visitor TestSuiteTestControllerUpdateResponseVisitor) error {
	if t.typ == "TestSuiteTestVoice" || t.TestSuiteTestVoice != nil {
		return visitor.VisitTestSuiteTestVoice(t.TestSuiteTestVoice)
	}
	if t.typ == "TestSuiteTestChat" || t.TestSuiteTestChat != nil {
		return visitor.VisitTestSuiteTestChat(t.TestSuiteTestChat)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}
