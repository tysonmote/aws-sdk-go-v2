// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists geofences stored in a given geofence collection.
func (c *Client) ListGeofences(ctx context.Context, params *ListGeofencesInput, optFns ...func(*Options)) (*ListGeofencesOutput, error) {
	if params == nil {
		params = &ListGeofencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGeofences", params, optFns, c.addOperationListGeofencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGeofencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGeofencesInput struct {

	// The name of the geofence collection storing the list of geofences.
	//
	// This member is required.
	CollectionName *string

	// The pagination token specifying which page of results to return in the response.
	// If no token is provided, the default page is the first page. Default value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListGeofencesOutput struct {

	// Contains a list of geofences stored in the geofence collection.
	//
	// This member is required.
	Entries []types.ListGeofenceResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGeofencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGeofences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGeofences{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListGeofencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGeofences(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListGeofencesAPIClient is a client that implements the ListGeofences operation.
type ListGeofencesAPIClient interface {
	ListGeofences(context.Context, *ListGeofencesInput, ...func(*Options)) (*ListGeofencesOutput, error)
}

var _ ListGeofencesAPIClient = (*Client)(nil)

// ListGeofencesPaginatorOptions is the paginator options for ListGeofences
type ListGeofencesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGeofencesPaginator is a paginator for ListGeofences
type ListGeofencesPaginator struct {
	options   ListGeofencesPaginatorOptions
	client    ListGeofencesAPIClient
	params    *ListGeofencesInput
	nextToken *string
	firstPage bool
}

// NewListGeofencesPaginator returns a new ListGeofencesPaginator
func NewListGeofencesPaginator(client ListGeofencesAPIClient, params *ListGeofencesInput, optFns ...func(*ListGeofencesPaginatorOptions)) *ListGeofencesPaginator {
	if params == nil {
		params = &ListGeofencesInput{}
	}

	options := ListGeofencesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGeofencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGeofencesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListGeofences page.
func (p *ListGeofencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGeofencesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListGeofences(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListGeofences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "ListGeofences",
	}
}
