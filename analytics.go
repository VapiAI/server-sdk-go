// This file was auto-generated by Fern from our API Definition.

package api

type AnalyticsQueryDto struct {
	// This is the list of metric queries you want to perform.
	Queries []*AnalyticsQuery `json:"queries,omitempty" url:"-"`
}
