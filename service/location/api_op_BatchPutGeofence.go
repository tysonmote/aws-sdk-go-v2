// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A batch request for storing geofence geometries into a given geofence
// collection, or updates the geometry of an existing geofence if a geofence ID is
// included in the request.
func (c *Client) BatchPutGeofence(ctx context.Context, params *BatchPutGeofenceInput, optFns ...func(*Options)) (*BatchPutGeofenceOutput, error) {
	if params == nil {
		params = &BatchPutGeofenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutGeofence", params, optFns, c.addOperationBatchPutGeofenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutGeofenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutGeofenceInput struct {

	// The geofence collection storing the geofences.
	//
	// This member is required.
	CollectionName *string

	// The batch of geofences to be stored in a geofence collection.
	//
	// This member is required.
	Entries []types.BatchPutGeofenceRequestEntry

	noSmithyDocumentSerde
}

type BatchPutGeofenceOutput struct {

	// Contains additional error details for each geofence that failed to be stored in
	// a geofence collection.
	//
	// This member is required.
	Errors []types.BatchPutGeofenceError

	// Contains each geofence that was successfully stored in a geofence collection.
	//
	// This member is required.
	Successes []types.BatchPutGeofenceSuccess

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutGeofenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchPutGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchPutGeofence{}, middleware.After)
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
	if err = addOpBatchPutGeofenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutGeofence(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchPutGeofence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchPutGeofence",
	}
}
