// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/VapiAI/server-sdk-go/internal"
	time "time"
)

type KnowledgeBasesListRequest struct {
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

type CreateTrieveKnowledgeBaseDto struct {
	// This is the name of the knowledge base.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// This is the plan on how to search the vector store while a call is going on.
	VectorStoreSearchPlan *TrieveKnowledgeBaseVectorStoreSearchPlan `json:"vectorStoreSearchPlan,omitempty" url:"vectorStoreSearchPlan,omitempty"`
	// This is the plan if you want us to create a new vector store on your behalf. To use an existing vector store from your account, use `vectoreStoreProviderId`
	VectorStoreCreatePlan *TrieveKnowledgeBaseVectorStoreCreatePlan `json:"vectorStoreCreatePlan,omitempty" url:"vectorStoreCreatePlan,omitempty"`
	// This is an vector store that you already have on your account with the provider. To create a new vector store, use vectorStoreCreatePlan.
	//
	// Usage:
	// - To bring your own vector store from Trieve, go to https://trieve.ai
	// - Create a dataset, and use the datasetId here.
	VectorStoreProviderId *string `json:"vectorStoreProviderId,omitempty" url:"vectorStoreProviderId,omitempty"`
	provider              string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateTrieveKnowledgeBaseDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTrieveKnowledgeBaseDto) GetVectorStoreSearchPlan() *TrieveKnowledgeBaseVectorStoreSearchPlan {
	if c == nil {
		return nil
	}
	return c.VectorStoreSearchPlan
}

func (c *CreateTrieveKnowledgeBaseDto) GetVectorStoreCreatePlan() *TrieveKnowledgeBaseVectorStoreCreatePlan {
	if c == nil {
		return nil
	}
	return c.VectorStoreCreatePlan
}

func (c *CreateTrieveKnowledgeBaseDto) GetVectorStoreProviderId() *string {
	if c == nil {
		return nil
	}
	return c.VectorStoreProviderId
}

func (c *CreateTrieveKnowledgeBaseDto) Provider() string {
	return c.provider
}

func (c *CreateTrieveKnowledgeBaseDto) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateTrieveKnowledgeBaseDto) UnmarshalJSON(data []byte) error {
	type embed CreateTrieveKnowledgeBaseDto
	var unmarshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateTrieveKnowledgeBaseDto(unmarshaler.embed)
	if unmarshaler.Provider != "trieve" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "trieve", unmarshaler.Provider)
	}
	c.provider = unmarshaler.Provider
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "provider")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateTrieveKnowledgeBaseDto) MarshalJSON() ([]byte, error) {
	type embed CreateTrieveKnowledgeBaseDto
	var marshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed:    embed(*c),
		Provider: "trieve",
	}
	return json.Marshal(marshaler)
}

func (c *CreateTrieveKnowledgeBaseDto) String() string {
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

type CustomKnowledgeBase struct {
	// /**
	// This is where the knowledge base request will be sent.
	//
	// Request Example:
	//
	// POST https://{server.url}
	// Content-Type: application/json
	//
	//	{
	//	  "messsage": {
	//	    "type": "knowledge-base-request",
	//	    "messages": [
	//	      {
	//	        "role": "user",
	//	        "content": "Why is ocean blue?"
	//	      }
	//	    ],
	//	    ...other metadata about the call...
	//	  }
	//	}
	//
	// Response Expected:
	// ```
	//
	//	{
	//	  "message": {
	//	     "role": "assistant",
	//	     "content": "The ocean is blue because water absorbs everything but blue.",
	//	  }, // YOU CAN RETURN THE EXACT RESPONSE TO SPEAK
	//	  "documents": [
	//	    {
	//	      "content": "The ocean is blue primarily because water absorbs colors in the red part of the light spectrum and scatters the blue light, making it more visible to our eyes.",
	//	      "similarity": 1
	//	    },
	//	    {
	//	      "content": "Blue light is scattered more by the water molecules than other colors, enhancing the blue appearance of the ocean.",
	//	      "similarity": .5
	//	    }
	//	  ] // OR, YOU CAN RETURN AN ARRAY OF DOCUMENTS THAT WILL BE SENT TO THE MODEL
	//	}
	//
	// ```
	Server *Server `json:"server,omitempty" url:"server,omitempty"`
	// This is the id of the knowledge base.
	Id string `json:"id" url:"id"`
	// This is the org id of the knowledge base.
	OrgId    string `json:"orgId" url:"orgId"`
	provider string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CustomKnowledgeBase) GetServer() *Server {
	if c == nil {
		return nil
	}
	return c.Server
}

func (c *CustomKnowledgeBase) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *CustomKnowledgeBase) GetOrgId() string {
	if c == nil {
		return ""
	}
	return c.OrgId
}

func (c *CustomKnowledgeBase) Provider() string {
	return c.provider
}

func (c *CustomKnowledgeBase) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CustomKnowledgeBase) UnmarshalJSON(data []byte) error {
	type embed CustomKnowledgeBase
	var unmarshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CustomKnowledgeBase(unmarshaler.embed)
	if unmarshaler.Provider != "custom-knowledge-base" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "custom-knowledge-base", unmarshaler.Provider)
	}
	c.provider = unmarshaler.Provider
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "provider")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CustomKnowledgeBase) MarshalJSON() ([]byte, error) {
	type embed CustomKnowledgeBase
	var marshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed:    embed(*c),
		Provider: "custom-knowledge-base",
	}
	return json.Marshal(marshaler)
}

func (c *CustomKnowledgeBase) String() string {
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

type TrieveKnowledgeBase struct {
	// This is the name of the knowledge base.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// This is the plan on how to search the vector store while a call is going on.
	VectorStoreSearchPlan *TrieveKnowledgeBaseVectorStoreSearchPlan `json:"vectorStoreSearchPlan,omitempty" url:"vectorStoreSearchPlan,omitempty"`
	// This is the plan if you want us to create a new vector store on your behalf. To use an existing vector store from your account, use `vectoreStoreProviderId`
	VectorStoreCreatePlan *TrieveKnowledgeBaseVectorStoreCreatePlan `json:"vectorStoreCreatePlan,omitempty" url:"vectorStoreCreatePlan,omitempty"`
	// This is an vector store that you already have on your account with the provider. To create a new vector store, use vectorStoreCreatePlan.
	//
	// Usage:
	// - To bring your own vector store from Trieve, go to https://trieve.ai
	// - Create a dataset, and use the datasetId here.
	VectorStoreProviderId *string `json:"vectorStoreProviderId,omitempty" url:"vectorStoreProviderId,omitempty"`
	// This is the id of the knowledge base.
	Id string `json:"id" url:"id"`
	// This is the org id of the knowledge base.
	OrgId    string `json:"orgId" url:"orgId"`
	provider string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TrieveKnowledgeBase) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TrieveKnowledgeBase) GetVectorStoreSearchPlan() *TrieveKnowledgeBaseVectorStoreSearchPlan {
	if t == nil {
		return nil
	}
	return t.VectorStoreSearchPlan
}

func (t *TrieveKnowledgeBase) GetVectorStoreCreatePlan() *TrieveKnowledgeBaseVectorStoreCreatePlan {
	if t == nil {
		return nil
	}
	return t.VectorStoreCreatePlan
}

func (t *TrieveKnowledgeBase) GetVectorStoreProviderId() *string {
	if t == nil {
		return nil
	}
	return t.VectorStoreProviderId
}

func (t *TrieveKnowledgeBase) GetId() string {
	if t == nil {
		return ""
	}
	return t.Id
}

func (t *TrieveKnowledgeBase) GetOrgId() string {
	if t == nil {
		return ""
	}
	return t.OrgId
}

func (t *TrieveKnowledgeBase) Provider() string {
	return t.provider
}

func (t *TrieveKnowledgeBase) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TrieveKnowledgeBase) UnmarshalJSON(data []byte) error {
	type embed TrieveKnowledgeBase
	var unmarshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TrieveKnowledgeBase(unmarshaler.embed)
	if unmarshaler.Provider != "trieve" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "trieve", unmarshaler.Provider)
	}
	t.provider = unmarshaler.Provider
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "provider")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TrieveKnowledgeBase) MarshalJSON() ([]byte, error) {
	type embed TrieveKnowledgeBase
	var marshaler = struct {
		embed
		Provider string `json:"provider"`
	}{
		embed:    embed(*t),
		Provider: "trieve",
	}
	return json.Marshal(marshaler)
}

func (t *TrieveKnowledgeBase) String() string {
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

type TrieveKnowledgeBaseVectorStoreCreatePlan struct {
	// These are the file ids that will be used to create the vector store. To upload files, use the `POST /files` endpoint.
	FileIds []string `json:"fileIds,omitempty" url:"fileIds,omitempty"`
	// This is an optional field which allows you to specify the number of splits you want per chunk. If not specified, the default 20 is used. However, you may want to use a different number.
	TargetSplitsPerChunk *float64 `json:"targetSplitsPerChunk,omitempty" url:"targetSplitsPerChunk,omitempty"`
	// This is an optional field which allows you to specify the delimiters to use when splitting the file before chunking the text. If not specified, the default [.!?\n] are used to split into sentences. However, you may want to use spaces or other delimiters.
	SplitDelimiters []string `json:"splitDelimiters,omitempty" url:"splitDelimiters,omitempty"`
	// This is an optional field which allows you to specify whether or not to rebalance the chunks created from the file. If not specified, the default true is used. If true, Trieve will evenly distribute remainder splits across chunks such that 66 splits with a target_splits_per_chunk of 20 will result in 3 chunks with 22 splits each.
	RebalanceChunks *bool `json:"rebalanceChunks,omitempty" url:"rebalanceChunks,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) GetFileIds() []string {
	if t == nil {
		return nil
	}
	return t.FileIds
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) GetTargetSplitsPerChunk() *float64 {
	if t == nil {
		return nil
	}
	return t.TargetSplitsPerChunk
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) GetSplitDelimiters() []string {
	if t == nil {
		return nil
	}
	return t.SplitDelimiters
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) GetRebalanceChunks() *bool {
	if t == nil {
		return nil
	}
	return t.RebalanceChunks
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) UnmarshalJSON(data []byte) error {
	type unmarshaler TrieveKnowledgeBaseVectorStoreCreatePlan
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TrieveKnowledgeBaseVectorStoreCreatePlan(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TrieveKnowledgeBaseVectorStoreCreatePlan) String() string {
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

type TrieveKnowledgeBaseVectorStoreSearchPlan struct {
	// If true, stop words (specified in server/src/stop-words.txt in the git repo) will be removed. This will preserve queries that are entirely stop words.
	RemoveStopWords *bool `json:"removeStopWords,omitempty" url:"removeStopWords,omitempty"`
	// This is the score threshold to filter out chunks with a score below the threshold for cosine distance metric. For Manhattan Distance, Euclidean Distance, and Dot Product, it will filter out scores above the threshold distance. This threshold applies before weight and bias modifications. If not specified, this defaults to no threshold. A threshold of 0 will default to no threshold.
	ScoreThreshold *float64 `json:"scoreThreshold,omitempty" url:"scoreThreshold,omitempty"`
	// This is the search method used when searching for relevant chunks from the vector store.
	SearchType TrieveKnowledgeBaseVectorStoreSearchPlanSearchType `json:"searchType" url:"searchType"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) GetRemoveStopWords() *bool {
	if t == nil {
		return nil
	}
	return t.RemoveStopWords
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) GetScoreThreshold() *float64 {
	if t == nil {
		return nil
	}
	return t.ScoreThreshold
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) GetSearchType() TrieveKnowledgeBaseVectorStoreSearchPlanSearchType {
	if t == nil {
		return ""
	}
	return t.SearchType
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) UnmarshalJSON(data []byte) error {
	type unmarshaler TrieveKnowledgeBaseVectorStoreSearchPlan
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TrieveKnowledgeBaseVectorStoreSearchPlan(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TrieveKnowledgeBaseVectorStoreSearchPlan) String() string {
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

// This is the search method used when searching for relevant chunks from the vector store.
type TrieveKnowledgeBaseVectorStoreSearchPlanSearchType string

const (
	TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeFulltext TrieveKnowledgeBaseVectorStoreSearchPlanSearchType = "fulltext"
	TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeSemantic TrieveKnowledgeBaseVectorStoreSearchPlanSearchType = "semantic"
	TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeHybrid   TrieveKnowledgeBaseVectorStoreSearchPlanSearchType = "hybrid"
	TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeBm25     TrieveKnowledgeBaseVectorStoreSearchPlanSearchType = "bm25"
)

func NewTrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeFromString(s string) (TrieveKnowledgeBaseVectorStoreSearchPlanSearchType, error) {
	switch s {
	case "fulltext":
		return TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeFulltext, nil
	case "semantic":
		return TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeSemantic, nil
	case "hybrid":
		return TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeHybrid, nil
	case "bm25":
		return TrieveKnowledgeBaseVectorStoreSearchPlanSearchTypeBm25, nil
	}
	var t TrieveKnowledgeBaseVectorStoreSearchPlanSearchType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TrieveKnowledgeBaseVectorStoreSearchPlanSearchType) Ptr() *TrieveKnowledgeBaseVectorStoreSearchPlanSearchType {
	return &t
}

type KnowledgeBasesCreateRequest struct {
	CreateTrieveKnowledgeBaseDto *CreateTrieveKnowledgeBaseDto
	CreateCustomKnowledgeBaseDto *CreateCustomKnowledgeBaseDto

	typ string
}

func (k *KnowledgeBasesCreateRequest) GetCreateTrieveKnowledgeBaseDto() *CreateTrieveKnowledgeBaseDto {
	if k == nil {
		return nil
	}
	return k.CreateTrieveKnowledgeBaseDto
}

func (k *KnowledgeBasesCreateRequest) GetCreateCustomKnowledgeBaseDto() *CreateCustomKnowledgeBaseDto {
	if k == nil {
		return nil
	}
	return k.CreateCustomKnowledgeBaseDto
}

func (k *KnowledgeBasesCreateRequest) UnmarshalJSON(data []byte) error {
	valueCreateTrieveKnowledgeBaseDto := new(CreateTrieveKnowledgeBaseDto)
	if err := json.Unmarshal(data, &valueCreateTrieveKnowledgeBaseDto); err == nil {
		k.typ = "CreateTrieveKnowledgeBaseDto"
		k.CreateTrieveKnowledgeBaseDto = valueCreateTrieveKnowledgeBaseDto
		return nil
	}
	valueCreateCustomKnowledgeBaseDto := new(CreateCustomKnowledgeBaseDto)
	if err := json.Unmarshal(data, &valueCreateCustomKnowledgeBaseDto); err == nil {
		k.typ = "CreateCustomKnowledgeBaseDto"
		k.CreateCustomKnowledgeBaseDto = valueCreateCustomKnowledgeBaseDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesCreateRequest) MarshalJSON() ([]byte, error) {
	if k.typ == "CreateTrieveKnowledgeBaseDto" || k.CreateTrieveKnowledgeBaseDto != nil {
		return json.Marshal(k.CreateTrieveKnowledgeBaseDto)
	}
	if k.typ == "CreateCustomKnowledgeBaseDto" || k.CreateCustomKnowledgeBaseDto != nil {
		return json.Marshal(k.CreateCustomKnowledgeBaseDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesCreateRequestVisitor interface {
	VisitCreateTrieveKnowledgeBaseDto(*CreateTrieveKnowledgeBaseDto) error
	VisitCreateCustomKnowledgeBaseDto(*CreateCustomKnowledgeBaseDto) error
}

func (k *KnowledgeBasesCreateRequest) Accept(visitor KnowledgeBasesCreateRequestVisitor) error {
	if k.typ == "CreateTrieveKnowledgeBaseDto" || k.CreateTrieveKnowledgeBaseDto != nil {
		return visitor.VisitCreateTrieveKnowledgeBaseDto(k.CreateTrieveKnowledgeBaseDto)
	}
	if k.typ == "CreateCustomKnowledgeBaseDto" || k.CreateCustomKnowledgeBaseDto != nil {
		return visitor.VisitCreateCustomKnowledgeBaseDto(k.CreateCustomKnowledgeBaseDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesCreateResponse struct {
	TrieveKnowledgeBase *TrieveKnowledgeBase
	CustomKnowledgeBase *CustomKnowledgeBase

	typ string
}

func (k *KnowledgeBasesCreateResponse) GetTrieveKnowledgeBase() *TrieveKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.TrieveKnowledgeBase
}

func (k *KnowledgeBasesCreateResponse) GetCustomKnowledgeBase() *CustomKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.CustomKnowledgeBase
}

func (k *KnowledgeBasesCreateResponse) UnmarshalJSON(data []byte) error {
	valueTrieveKnowledgeBase := new(TrieveKnowledgeBase)
	if err := json.Unmarshal(data, &valueTrieveKnowledgeBase); err == nil {
		k.typ = "TrieveKnowledgeBase"
		k.TrieveKnowledgeBase = valueTrieveKnowledgeBase
		return nil
	}
	valueCustomKnowledgeBase := new(CustomKnowledgeBase)
	if err := json.Unmarshal(data, &valueCustomKnowledgeBase); err == nil {
		k.typ = "CustomKnowledgeBase"
		k.CustomKnowledgeBase = valueCustomKnowledgeBase
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesCreateResponse) MarshalJSON() ([]byte, error) {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return json.Marshal(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return json.Marshal(k.CustomKnowledgeBase)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesCreateResponseVisitor interface {
	VisitTrieveKnowledgeBase(*TrieveKnowledgeBase) error
	VisitCustomKnowledgeBase(*CustomKnowledgeBase) error
}

func (k *KnowledgeBasesCreateResponse) Accept(visitor KnowledgeBasesCreateResponseVisitor) error {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return visitor.VisitTrieveKnowledgeBase(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return visitor.VisitCustomKnowledgeBase(k.CustomKnowledgeBase)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesDeleteResponse struct {
	TrieveKnowledgeBase *TrieveKnowledgeBase
	CustomKnowledgeBase *CustomKnowledgeBase

	typ string
}

func (k *KnowledgeBasesDeleteResponse) GetTrieveKnowledgeBase() *TrieveKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.TrieveKnowledgeBase
}

func (k *KnowledgeBasesDeleteResponse) GetCustomKnowledgeBase() *CustomKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.CustomKnowledgeBase
}

func (k *KnowledgeBasesDeleteResponse) UnmarshalJSON(data []byte) error {
	valueTrieveKnowledgeBase := new(TrieveKnowledgeBase)
	if err := json.Unmarshal(data, &valueTrieveKnowledgeBase); err == nil {
		k.typ = "TrieveKnowledgeBase"
		k.TrieveKnowledgeBase = valueTrieveKnowledgeBase
		return nil
	}
	valueCustomKnowledgeBase := new(CustomKnowledgeBase)
	if err := json.Unmarshal(data, &valueCustomKnowledgeBase); err == nil {
		k.typ = "CustomKnowledgeBase"
		k.CustomKnowledgeBase = valueCustomKnowledgeBase
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesDeleteResponse) MarshalJSON() ([]byte, error) {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return json.Marshal(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return json.Marshal(k.CustomKnowledgeBase)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesDeleteResponseVisitor interface {
	VisitTrieveKnowledgeBase(*TrieveKnowledgeBase) error
	VisitCustomKnowledgeBase(*CustomKnowledgeBase) error
}

func (k *KnowledgeBasesDeleteResponse) Accept(visitor KnowledgeBasesDeleteResponseVisitor) error {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return visitor.VisitTrieveKnowledgeBase(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return visitor.VisitCustomKnowledgeBase(k.CustomKnowledgeBase)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesGetResponse struct {
	TrieveKnowledgeBase *TrieveKnowledgeBase
	CustomKnowledgeBase *CustomKnowledgeBase

	typ string
}

func (k *KnowledgeBasesGetResponse) GetTrieveKnowledgeBase() *TrieveKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.TrieveKnowledgeBase
}

func (k *KnowledgeBasesGetResponse) GetCustomKnowledgeBase() *CustomKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.CustomKnowledgeBase
}

func (k *KnowledgeBasesGetResponse) UnmarshalJSON(data []byte) error {
	valueTrieveKnowledgeBase := new(TrieveKnowledgeBase)
	if err := json.Unmarshal(data, &valueTrieveKnowledgeBase); err == nil {
		k.typ = "TrieveKnowledgeBase"
		k.TrieveKnowledgeBase = valueTrieveKnowledgeBase
		return nil
	}
	valueCustomKnowledgeBase := new(CustomKnowledgeBase)
	if err := json.Unmarshal(data, &valueCustomKnowledgeBase); err == nil {
		k.typ = "CustomKnowledgeBase"
		k.CustomKnowledgeBase = valueCustomKnowledgeBase
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesGetResponse) MarshalJSON() ([]byte, error) {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return json.Marshal(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return json.Marshal(k.CustomKnowledgeBase)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesGetResponseVisitor interface {
	VisitTrieveKnowledgeBase(*TrieveKnowledgeBase) error
	VisitCustomKnowledgeBase(*CustomKnowledgeBase) error
}

func (k *KnowledgeBasesGetResponse) Accept(visitor KnowledgeBasesGetResponseVisitor) error {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return visitor.VisitTrieveKnowledgeBase(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return visitor.VisitCustomKnowledgeBase(k.CustomKnowledgeBase)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesListResponseItem struct {
	TrieveKnowledgeBase *TrieveKnowledgeBase
	CustomKnowledgeBase *CustomKnowledgeBase

	typ string
}

func (k *KnowledgeBasesListResponseItem) GetTrieveKnowledgeBase() *TrieveKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.TrieveKnowledgeBase
}

func (k *KnowledgeBasesListResponseItem) GetCustomKnowledgeBase() *CustomKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.CustomKnowledgeBase
}

func (k *KnowledgeBasesListResponseItem) UnmarshalJSON(data []byte) error {
	valueTrieveKnowledgeBase := new(TrieveKnowledgeBase)
	if err := json.Unmarshal(data, &valueTrieveKnowledgeBase); err == nil {
		k.typ = "TrieveKnowledgeBase"
		k.TrieveKnowledgeBase = valueTrieveKnowledgeBase
		return nil
	}
	valueCustomKnowledgeBase := new(CustomKnowledgeBase)
	if err := json.Unmarshal(data, &valueCustomKnowledgeBase); err == nil {
		k.typ = "CustomKnowledgeBase"
		k.CustomKnowledgeBase = valueCustomKnowledgeBase
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesListResponseItem) MarshalJSON() ([]byte, error) {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return json.Marshal(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return json.Marshal(k.CustomKnowledgeBase)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesListResponseItemVisitor interface {
	VisitTrieveKnowledgeBase(*TrieveKnowledgeBase) error
	VisitCustomKnowledgeBase(*CustomKnowledgeBase) error
}

func (k *KnowledgeBasesListResponseItem) Accept(visitor KnowledgeBasesListResponseItemVisitor) error {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return visitor.VisitTrieveKnowledgeBase(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return visitor.VisitCustomKnowledgeBase(k.CustomKnowledgeBase)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesUpdateResponse struct {
	TrieveKnowledgeBase *TrieveKnowledgeBase
	CustomKnowledgeBase *CustomKnowledgeBase

	typ string
}

func (k *KnowledgeBasesUpdateResponse) GetTrieveKnowledgeBase() *TrieveKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.TrieveKnowledgeBase
}

func (k *KnowledgeBasesUpdateResponse) GetCustomKnowledgeBase() *CustomKnowledgeBase {
	if k == nil {
		return nil
	}
	return k.CustomKnowledgeBase
}

func (k *KnowledgeBasesUpdateResponse) UnmarshalJSON(data []byte) error {
	valueTrieveKnowledgeBase := new(TrieveKnowledgeBase)
	if err := json.Unmarshal(data, &valueTrieveKnowledgeBase); err == nil {
		k.typ = "TrieveKnowledgeBase"
		k.TrieveKnowledgeBase = valueTrieveKnowledgeBase
		return nil
	}
	valueCustomKnowledgeBase := new(CustomKnowledgeBase)
	if err := json.Unmarshal(data, &valueCustomKnowledgeBase); err == nil {
		k.typ = "CustomKnowledgeBase"
		k.CustomKnowledgeBase = valueCustomKnowledgeBase
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, k)
}

func (k KnowledgeBasesUpdateResponse) MarshalJSON() ([]byte, error) {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return json.Marshal(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return json.Marshal(k.CustomKnowledgeBase)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", k)
}

type KnowledgeBasesUpdateResponseVisitor interface {
	VisitTrieveKnowledgeBase(*TrieveKnowledgeBase) error
	VisitCustomKnowledgeBase(*CustomKnowledgeBase) error
}

func (k *KnowledgeBasesUpdateResponse) Accept(visitor KnowledgeBasesUpdateResponseVisitor) error {
	if k.typ == "TrieveKnowledgeBase" || k.TrieveKnowledgeBase != nil {
		return visitor.VisitTrieveKnowledgeBase(k.TrieveKnowledgeBase)
	}
	if k.typ == "CustomKnowledgeBase" || k.CustomKnowledgeBase != nil {
		return visitor.VisitCustomKnowledgeBase(k.CustomKnowledgeBase)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", k)
}
