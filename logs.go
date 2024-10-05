// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	time "time"
)

type LogsGetRequest struct {
	// This is the unique identifier for the org that this log belongs to.
	OrgId *string `json:"-" url:"orgId,omitempty"`
	// This is the type of the log.
	Type *LogsGetRequestType `json:"-" url:"type,omitempty"`
	// This is the ID of the assistant.
	AssistantId *string `json:"-" url:"assistantId,omitempty"`
	// This is the ID of the phone number.
	PhoneNumberId *string `json:"-" url:"phoneNumberId,omitempty"`
	// This is the ID of the customer.
	CustomerId *string `json:"-" url:"customerId,omitempty"`
	// This is the ID of the squad.
	SquadId *string `json:"-" url:"squadId,omitempty"`
	// This is the ID of the call.
	CallId *string `json:"-" url:"callId,omitempty"`
	// This is the page number to return. Defaults to 1.
	Page *float64 `json:"-" url:"page,omitempty"`
	// This is the sort order for pagination. Defaults to 'ASC'.
	SortOrder *LogsGetRequestSortOrder `json:"-" url:"sortOrder,omitempty"`
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

type LogsGetRequestSortOrder string

const (
	LogsGetRequestSortOrderAsc  LogsGetRequestSortOrder = "ASC"
	LogsGetRequestSortOrderDesc LogsGetRequestSortOrder = "DESC"
)

func NewLogsGetRequestSortOrderFromString(s string) (LogsGetRequestSortOrder, error) {
	switch s {
	case "ASC":
		return LogsGetRequestSortOrderAsc, nil
	case "DESC":
		return LogsGetRequestSortOrderDesc, nil
	}
	var t LogsGetRequestSortOrder
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LogsGetRequestSortOrder) Ptr() *LogsGetRequestSortOrder {
	return &l
}

type LogsGetRequestType string

const (
	LogsGetRequestTypeApi      LogsGetRequestType = "API"
	LogsGetRequestTypeWebhook  LogsGetRequestType = "Webhook"
	LogsGetRequestTypeCall     LogsGetRequestType = "Call"
	LogsGetRequestTypeProvider LogsGetRequestType = "Provider"
)

func NewLogsGetRequestTypeFromString(s string) (LogsGetRequestType, error) {
	switch s {
	case "API":
		return LogsGetRequestTypeApi, nil
	case "Webhook":
		return LogsGetRequestTypeWebhook, nil
	case "Call":
		return LogsGetRequestTypeCall, nil
	case "Provider":
		return LogsGetRequestTypeProvider, nil
	}
	var t LogsGetRequestType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LogsGetRequestType) Ptr() *LogsGetRequestType {
	return &l
}