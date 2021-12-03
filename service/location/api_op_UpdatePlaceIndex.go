// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties of a given place index resource.
func (c *Client) UpdatePlaceIndex(ctx context.Context, params *UpdatePlaceIndexInput, optFns ...func(*Options)) (*UpdatePlaceIndexOutput, error) {
	if params == nil {
		params = &UpdatePlaceIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePlaceIndex", params, optFns, c.addOperationUpdatePlaceIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePlaceIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePlaceIndexInput struct {

	// The name of the place index resource to update.
	//
	// This member is required.
	IndexName *string

	// Updates the data storage option for the place index resource.
	DataSourceConfiguration *types.DataSourceConfiguration

	// Updates the description for the place index resource.
	Description *string

	// Updates the pricing plan for the place index resource. For more information
	// about each pricing plan option restrictions, see Amazon Location Service pricing
	// (https://aws.amazon.com/location/pricing/).
	PricingPlan types.PricingPlan

	noSmithyDocumentSerde
}

type UpdatePlaceIndexOutput struct {

	// The Amazon Resource Name (ARN) of the upated place index resource. Used to
	// specify a resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:place- index/ExamplePlaceIndex
	//
	// This member is required.
	IndexArn *string

	// The name of the updated place index resource.
	//
	// This member is required.
	IndexName *string

	// The timestamp for when the place index resource was last updated in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePlaceIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePlaceIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePlaceIndex{}, middleware.After)
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
	if err = addOpUpdatePlaceIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePlaceIndex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePlaceIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdatePlaceIndex",
	}
}
