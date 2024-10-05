// This file was auto-generated by Fern from our API Definition.

package logs

import (
	context "context"
	fmt "fmt"
	vapigosdk "github.com/fern-demo/vapi-go-sdk"
	core "github.com/fern-demo/vapi-go-sdk/core"
	option "github.com/fern-demo/vapi-go-sdk/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) Get(
	ctx context.Context,
	request *vapigosdk.LogsGetRequest,
	opts ...option.RequestOption,
) (*core.Page[*vapigosdk.Log], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.vapi.ai"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/logs"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*float64]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("page", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &core.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	next := 1
	if request.Page != nil {
		next = *request.Page
	}
	readPageResponse := func(response *vapigosdk.LogsPaginatedResponse) *core.PageResponse[*float64, *vapigosdk.Log] {
		next += 1
		results := response.Results
		return &core.PageResponse[*float64, *vapigosdk.Log]{
			Next:    &next,
			Results: results,
		}
	}
	pager := core.NewOffsetPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, &next)
}