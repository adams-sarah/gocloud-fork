// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package logging

import (
	"fmt"
	"runtime"
	"time"

	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	googleapis_api_monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	googleapis_logging_v2 "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	loggingProjectPathTemplate = gax.MustCompilePathTemplate("projects/{project}")
	loggingLogPathTemplate     = gax.MustCompilePathTemplate("projects/{project}/logs/{log}")
)

// CallOptions contains the retry settings for each method of this client.
type CallOptions struct {
	DeleteLog                        []gax.CallOption
	WriteLogEntries                  []gax.CallOption
	ListLogEntries                   []gax.CallOption
	ListMonitoredResourceDescriptors []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("logging.googleapis.com:443"),
		option.WithScopes(
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/cloud-platform.read-only",
			"https://www.googleapis.com/auth/logging.admin",
			"https://www.googleapis.com/auth/logging.read",
			"https://www.googleapis.com/auth/logging.write",
		),
	}
}

func defaultRetryOptions() []gax.CallOption {
	return []gax.CallOption{
		gax.WithTimeout(45000 * time.Millisecond),
		gax.WithDelayTimeoutSettings(100*time.Millisecond, 1000*time.Millisecond, 1.2),
		gax.WithRPCTimeoutSettings(2000*time.Millisecond, 30000*time.Millisecond, 1.5),
	}
}
func listRetryOptions() []gax.CallOption {
	return []gax.CallOption{
		gax.WithTimeout(45000 * time.Millisecond),
		gax.WithDelayTimeoutSettings(100*time.Millisecond, 1000*time.Millisecond, 1.2),
		gax.WithRPCTimeoutSettings(7000*time.Millisecond, 30000*time.Millisecond, 1.5),
	}
}

func defaultCallOptions() *CallOptions {
	withIdempotentRetryCodes := gax.WithRetryCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
	},
	)
	return &CallOptions{
		DeleteLog:                        append(defaultRetryOptions(), withIdempotentRetryCodes),
		WriteLogEntries:                  defaultRetryOptions(),
		ListLogEntries:                   append(listRetryOptions(), withIdempotentRetryCodes),
		ListMonitoredResourceDescriptors: append(defaultRetryOptions(), withIdempotentRetryCodes),
	}
}

// Client is a client for interacting with LoggingServiceV2.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client googleapis_logging_v2.LoggingServiceV2Client

	// The call options for this service.
	CallOptions *CallOptions

	// The metadata to be sent with each request.
	metadata map[string][]string
}

// NewClient creates a new logging service client.
//
// Service for ingesting and querying logs.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		client:      googleapis_logging_v2.NewLoggingServiceV2Client(conn),
		CallOptions: defaultCallOptions(),
	}
	c.SetGoogleClientInfo("gax", gax.Version)
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) SetGoogleClientInfo(name, version string) {
	c.metadata = map[string][]string{
		"x-goog-api-client": {fmt.Sprintf("%s/%s %s gax/%s go/%s", name, version, gapicNameVersion, gax.Version, runtime.Version())},
	}
}

// Path templates.

// ProjectPath returns the path for the project resource.
func LoggingProjectPath(project string) string {
	path, err := loggingProjectPathTemplate.Render(map[string]string{
		"project": project,
	})
	if err != nil {
		panic(err)
	}
	return path
}

// LogPath returns the path for the log resource.
func LoggingLogPath(project string, log string) string {
	path, err := loggingLogPathTemplate.Render(map[string]string{
		"project": project,
		"log":     log,
	})
	if err != nil {
		panic(err)
	}
	return path
}

// DeleteLog deletes a log and all its log entries.
// The log will reappear if it receives new entries.
func (c *Client) DeleteLog(ctx context.Context, req *googleapis_logging_v2.DeleteLogRequest) error {
	ctx = metadata.NewContext(ctx, c.metadata)
	err := gax.Invoke(ctx, func(ctx context.Context) error {
		var err error
		_, err = c.client.DeleteLog(ctx, req)
		return err
	}, c.CallOptions.DeleteLog...)
	return err
}

// WriteLogEntries writes log entries to Stackdriver Logging.  All log entries are
// written by this method.
func (c *Client) WriteLogEntries(ctx context.Context, req *googleapis_logging_v2.WriteLogEntriesRequest) (*googleapis_logging_v2.WriteLogEntriesResponse, error) {
	ctx = metadata.NewContext(ctx, c.metadata)
	var resp *googleapis_logging_v2.WriteLogEntriesResponse
	err := gax.Invoke(ctx, func(ctx context.Context) error {
		var err error
		resp, err = c.client.WriteLogEntries(ctx, req)
		return err
	}, c.CallOptions.WriteLogEntries...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListLogEntries lists log entries.  Use this method to retrieve log entries from Cloud
// Logging.  For ways to export log entries, see
// [Exporting Logs](/logging/docs/export).
func (c *Client) ListLogEntries(ctx context.Context, req *googleapis_logging_v2.ListLogEntriesRequest) *LogEntryIterator {
	ctx = metadata.NewContext(ctx, c.metadata)
	it := &LogEntryIterator{}
	it.apiCall = func() error {
		var resp *googleapis_logging_v2.ListLogEntriesResponse
		err := gax.Invoke(ctx, func(ctx context.Context) error {
			var err error
			req.PageToken = it.nextPageToken
			req.PageSize = it.pageSize
			resp, err = c.client.ListLogEntries(ctx, req)
			return err
		}, c.CallOptions.ListLogEntries...)
		if err != nil {
			return err
		}
		if resp.NextPageToken == "" {
			it.atLastPage = true
		}
		it.nextPageToken = resp.NextPageToken
		it.items = resp.Entries
		return nil
	}
	return it
}

// ListMonitoredResourceDescriptors lists the monitored resource descriptors used by Stackdriver Logging.
func (c *Client) ListMonitoredResourceDescriptors(ctx context.Context, req *googleapis_logging_v2.ListMonitoredResourceDescriptorsRequest) *MonitoredResourceDescriptorIterator {
	ctx = metadata.NewContext(ctx, c.metadata)
	it := &MonitoredResourceDescriptorIterator{}
	it.apiCall = func() error {
		var resp *googleapis_logging_v2.ListMonitoredResourceDescriptorsResponse
		err := gax.Invoke(ctx, func(ctx context.Context) error {
			var err error
			req.PageToken = it.nextPageToken
			req.PageSize = it.pageSize
			resp, err = c.client.ListMonitoredResourceDescriptors(ctx, req)
			return err
		}, c.CallOptions.ListMonitoredResourceDescriptors...)
		if err != nil {
			return err
		}
		if resp.NextPageToken == "" {
			it.atLastPage = true
		}
		it.nextPageToken = resp.NextPageToken
		it.items = resp.ResourceDescriptors
		return nil
	}
	return it
}

// Iterators.
//

// LogEntryIterator manages a stream of *googleapis_logging_v2.LogEntry.
type LogEntryIterator struct {
	// The current page data.
	items         []*googleapis_logging_v2.LogEntry
	atLastPage    bool
	currentIndex  int
	pageSize      int32
	nextPageToken string
	apiCall       func() error
}

// NextPage returns the next page of results.
// It will return at most the number of results specified by the last call to SetPageSize.
// If SetPageSize was never called or was called with a value less than 1,
// the page size is determined by the underlying service.
//
// NextPage may return a second return value of Done along with the last page of results. After
// NextPage returns Done, all subsequent calls to NextPage will return (nil, Done).
//
// Next and NextPage should not be used with the same iterator.
func (it *LogEntryIterator) NextPage() ([]*googleapis_logging_v2.LogEntry, error) {
	if it.atLastPage {
		// We already returned Done with the last page of items. Continue to
		// return Done, but with no items.
		return nil, Done
	}
	if err := it.apiCall(); err != nil {
		return nil, err
	}
	if it.atLastPage {
		return it.items, Done
	}
	return it.items, nil
}

// Next returns the next result. Its second return value is Done if there are no more results.
// Once next returns Done, all subsequent calls will return Done.
//
// Internally, Next retrieves results in bulk. You can call SetPageSize as a performance hint to
// affect how many results are retrieved in a single RPC.
//
// SetPageToken should not be called when using Next.
//
// Next and NextPage should not be used with the same iterator.
func (it *LogEntryIterator) Next() (*googleapis_logging_v2.LogEntry, error) {
	for it.currentIndex >= len(it.items) {
		if it.atLastPage {
			return nil, Done
		}
		if err := it.apiCall(); err != nil {
			return nil, err
		}
		it.currentIndex = 0
	}
	result := it.items[it.currentIndex]
	it.currentIndex++
	return result, nil
}

// PageSize returns the page size for all subsequent calls to NextPage.
func (it *LogEntryIterator) PageSize() int32 {
	return it.pageSize
}

// SetPageSize sets the page size for all subsequent calls to NextPage.
func (it *LogEntryIterator) SetPageSize(pageSize int32) {
	it.pageSize = pageSize
}

// SetPageToken sets the page token for the next call to NextPage, to resume the iteration from
// a previous point.
func (it *LogEntryIterator) SetPageToken(token string) {
	it.nextPageToken = token
}

// NextPageToken returns a page token that can be used with SetPageToken to resume
// iteration from the next page. It returns the empty string if there are no more pages.
func (it *LogEntryIterator) NextPageToken() string {
	return it.nextPageToken
}

// MonitoredResourceDescriptorIterator manages a stream of *googleapis_api_monitoredres.MonitoredResourceDescriptor.
type MonitoredResourceDescriptorIterator struct {
	// The current page data.
	items         []*googleapis_api_monitoredres.MonitoredResourceDescriptor
	atLastPage    bool
	currentIndex  int
	pageSize      int32
	nextPageToken string
	apiCall       func() error
}

// NextPage returns the next page of results.
// It will return at most the number of results specified by the last call to SetPageSize.
// If SetPageSize was never called or was called with a value less than 1,
// the page size is determined by the underlying service.
//
// NextPage may return a second return value of Done along with the last page of results. After
// NextPage returns Done, all subsequent calls to NextPage will return (nil, Done).
//
// Next and NextPage should not be used with the same iterator.
func (it *MonitoredResourceDescriptorIterator) NextPage() ([]*googleapis_api_monitoredres.MonitoredResourceDescriptor, error) {
	if it.atLastPage {
		// We already returned Done with the last page of items. Continue to
		// return Done, but with no items.
		return nil, Done
	}
	if err := it.apiCall(); err != nil {
		return nil, err
	}
	if it.atLastPage {
		return it.items, Done
	}
	return it.items, nil
}

// Next returns the next result. Its second return value is Done if there are no more results.
// Once next returns Done, all subsequent calls will return Done.
//
// Internally, Next retrieves results in bulk. You can call SetPageSize as a performance hint to
// affect how many results are retrieved in a single RPC.
//
// SetPageToken should not be called when using Next.
//
// Next and NextPage should not be used with the same iterator.
func (it *MonitoredResourceDescriptorIterator) Next() (*googleapis_api_monitoredres.MonitoredResourceDescriptor, error) {
	for it.currentIndex >= len(it.items) {
		if it.atLastPage {
			return nil, Done
		}
		if err := it.apiCall(); err != nil {
			return nil, err
		}
		it.currentIndex = 0
	}
	result := it.items[it.currentIndex]
	it.currentIndex++
	return result, nil
}

// PageSize returns the page size for all subsequent calls to NextPage.
func (it *MonitoredResourceDescriptorIterator) PageSize() int32 {
	return it.pageSize
}

// SetPageSize sets the page size for all subsequent calls to NextPage.
func (it *MonitoredResourceDescriptorIterator) SetPageSize(pageSize int32) {
	it.pageSize = pageSize
}

// SetPageToken sets the page token for the next call to NextPage, to resume the iteration from
// a previous point.
func (it *MonitoredResourceDescriptorIterator) SetPageToken(token string) {
	it.nextPageToken = token
}

// NextPageToken returns a page token that can be used with SetPageToken to resume
// iteration from the next page. It returns the empty string if there are no more pages.
func (it *MonitoredResourceDescriptorIterator) NextPageToken() string {
	return it.nextPageToken
}
