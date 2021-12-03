// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources in a resource share that is shared with you but for which
// the invitation is still PENDING. That means that you haven't accepted or
// rejected the invitation and the invitation hasn't expired.
func (c *Client) ListPendingInvitationResources(ctx context.Context, params *ListPendingInvitationResourcesInput, optFns ...func(*Options)) (*ListPendingInvitationResourcesOutput, error) {
	if params == nil {
		params = &ListPendingInvitationResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPendingInvitationResources", params, optFns, c.addOperationListPendingInvitationResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPendingInvitationResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPendingInvitationResourcesInput struct {

	// Specifies the Amazon Resoure Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the invitation. You can use GetResourceShareInvitations to find the ARN of the
	// invitation.
	//
	// This member is required.
	ResourceShareInvitationArn *string

	// Specifies the total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies that you want the results to include only resources that have the
	// specified scope.
	//
	// * ALL – the results include both global and regional resources
	// or resource types.
	//
	// * GLOBAL – the results include only global resources or
	// resource types.
	//
	// * REGIONAL – the results include only regional resources or
	// resource types.
	//
	// The default value is ALL.
	ResourceRegionScope types.ResourceRegionScopeFilter

	noSmithyDocumentSerde
}

type ListPendingInvitationResourcesOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null. This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of objects that contain the information about the resources included
	// the specified resource share.
	Resources []types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPendingInvitationResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPendingInvitationResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPendingInvitationResources{}, middleware.After)
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
	if err = addOpListPendingInvitationResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPendingInvitationResources(options.Region), middleware.Before); err != nil {
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

// ListPendingInvitationResourcesAPIClient is a client that implements the
// ListPendingInvitationResources operation.
type ListPendingInvitationResourcesAPIClient interface {
	ListPendingInvitationResources(context.Context, *ListPendingInvitationResourcesInput, ...func(*Options)) (*ListPendingInvitationResourcesOutput, error)
}

var _ ListPendingInvitationResourcesAPIClient = (*Client)(nil)

// ListPendingInvitationResourcesPaginatorOptions is the paginator options for
// ListPendingInvitationResources
type ListPendingInvitationResourcesPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPendingInvitationResourcesPaginator is a paginator for
// ListPendingInvitationResources
type ListPendingInvitationResourcesPaginator struct {
	options   ListPendingInvitationResourcesPaginatorOptions
	client    ListPendingInvitationResourcesAPIClient
	params    *ListPendingInvitationResourcesInput
	nextToken *string
	firstPage bool
}

// NewListPendingInvitationResourcesPaginator returns a new
// ListPendingInvitationResourcesPaginator
func NewListPendingInvitationResourcesPaginator(client ListPendingInvitationResourcesAPIClient, params *ListPendingInvitationResourcesInput, optFns ...func(*ListPendingInvitationResourcesPaginatorOptions)) *ListPendingInvitationResourcesPaginator {
	if params == nil {
		params = &ListPendingInvitationResourcesInput{}
	}

	options := ListPendingInvitationResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPendingInvitationResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPendingInvitationResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPendingInvitationResources page.
func (p *ListPendingInvitationResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPendingInvitationResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPendingInvitationResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPendingInvitationResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListPendingInvitationResources",
	}
}
