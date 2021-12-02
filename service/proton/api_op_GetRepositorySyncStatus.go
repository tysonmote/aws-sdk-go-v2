// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the repository sync status.
func (c *Client) GetRepositorySyncStatus(ctx context.Context, params *GetRepositorySyncStatusInput, optFns ...func(*Options)) (*GetRepositorySyncStatusOutput, error) {
	if params == nil {
		params = &GetRepositorySyncStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRepositorySyncStatus", params, optFns, c.addOperationGetRepositorySyncStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRepositorySyncStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRepositorySyncStatusInput struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The repository name.
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The repository sync type.
	//
	// This member is required.
	SyncType types.SyncType

	noSmithyDocumentSerde
}

type GetRepositorySyncStatusOutput struct {

	// The repository sync status detail data that's returned by Proton.
	LatestSync *types.RepositorySyncAttempt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRepositorySyncStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRepositorySyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRepositorySyncStatus{}, middleware.After)
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
	if err = addOpGetRepositorySyncStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositorySyncStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRepositorySyncStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetRepositorySyncStatus",
	}
}