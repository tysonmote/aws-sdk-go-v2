// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new AMP workspace.
func (c *Client) CreateWorkspace(ctx context.Context, params *CreateWorkspaceInput, optFns ...func(*Options)) (*CreateWorkspaceOutput, error) {
	if params == nil {
		params = &CreateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspace", params, optFns, c.addOperationCreateWorkspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateWorkspace operation.
type CreateWorkspaceInput struct {

	// An optional user-assigned alias for this workspace. This alias is for user
	// reference and does not need to be unique.
	Alias *string

	// Optional, unique, case-sensitive, user-provided identifier to ensure the
	// idempotency of the request.
	ClientToken *string

	// Optional, user-provided tags for this workspace.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents the output of a CreateWorkspace operation.
type CreateWorkspaceOutput struct {

	// The ARN of the workspace that was just created.
	//
	// This member is required.
	Arn *string

	// The status of the workspace that was just created (usually CREATING).
	//
	// This member is required.
	Status *types.WorkspaceStatus

	// The generated ID of the workspace that was just created.
	//
	// This member is required.
	WorkspaceId *string

	// The tags of this workspace.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkspace{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aps",
		OperationName: "CreateWorkspace",
	}
}
