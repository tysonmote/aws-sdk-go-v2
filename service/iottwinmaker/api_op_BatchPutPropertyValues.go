// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets values for multiple time series properties.
func (c *Client) BatchPutPropertyValues(ctx context.Context, params *BatchPutPropertyValuesInput, optFns ...func(*Options)) (*BatchPutPropertyValuesOutput, error) {
	if params == nil {
		params = &BatchPutPropertyValuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutPropertyValues", params, optFns, c.addOperationBatchPutPropertyValuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutPropertyValuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutPropertyValuesInput struct {

	// An object that maps strings to the property value entries to set. Each string in
	// the mapping must be unique to this object.
	//
	// This member is required.
	Entries []types.PropertyValueEntry

	// The ID of the workspace that contains the properties to set.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type BatchPutPropertyValuesOutput struct {

	// Entries that caused errors in the batch put operation.
	//
	// This member is required.
	ErrorEntries []types.BatchPutPropertyErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutPropertyValuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutPropertyValues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutPropertyValues{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchPutPropertyValuesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchPutPropertyValuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutPropertyValues(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchPutPropertyValuesMiddleware struct {
}

func (*endpointPrefix_opBatchPutPropertyValuesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchPutPropertyValuesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchPutPropertyValuesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchPutPropertyValuesMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchPutPropertyValues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "BatchPutPropertyValues",
	}
}