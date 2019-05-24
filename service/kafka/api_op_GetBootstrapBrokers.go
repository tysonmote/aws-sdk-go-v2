// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/GetBootstrapBrokersRequest
type GetBootstrapBrokersInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBootstrapBrokersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBootstrapBrokersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBootstrapBrokersInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBootstrapBrokersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Returns a string containing one or more hostname:port pairs.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/GetBootstrapBrokersResponse
type GetBootstrapBrokersOutput struct {
	_ struct{} `type:"structure"`

	// A string containing one or more hostname:port pairs.
	BootstrapBrokerString *string `locationName:"bootstrapBrokerString" type:"string"`
}

// String returns the string representation
func (s GetBootstrapBrokersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBootstrapBrokersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BootstrapBrokerString != nil {
		v := *s.BootstrapBrokerString

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "bootstrapBrokerString", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBootstrapBrokers = "GetBootstrapBrokers"

// GetBootstrapBrokersRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// A list of brokers that a client application can use to bootstrap.
//
//    // Example sending a request using GetBootstrapBrokersRequest.
//    req := client.GetBootstrapBrokersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/GetBootstrapBrokers
func (c *Client) GetBootstrapBrokersRequest(input *GetBootstrapBrokersInput) GetBootstrapBrokersRequest {
	op := &aws.Operation{
		Name:       opGetBootstrapBrokers,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{clusterArn}/bootstrap-brokers",
	}

	if input == nil {
		input = &GetBootstrapBrokersInput{}
	}

	req := c.newRequest(op, input, &GetBootstrapBrokersOutput{})
	return GetBootstrapBrokersRequest{Request: req, Input: input, Copy: c.GetBootstrapBrokersRequest}
}

// GetBootstrapBrokersRequest is the request type for the
// GetBootstrapBrokers API operation.
type GetBootstrapBrokersRequest struct {
	*aws.Request
	Input *GetBootstrapBrokersInput
	Copy  func(*GetBootstrapBrokersInput) GetBootstrapBrokersRequest
}

// Send marshals and sends the GetBootstrapBrokers API request.
func (r GetBootstrapBrokersRequest) Send(ctx context.Context) (*GetBootstrapBrokersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBootstrapBrokersResponse{
		GetBootstrapBrokersOutput: r.Request.Data.(*GetBootstrapBrokersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBootstrapBrokersResponse is the response type for the
// GetBootstrapBrokers API operation.
type GetBootstrapBrokersResponse struct {
	*GetBootstrapBrokersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBootstrapBrokers request.
func (r *GetBootstrapBrokersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}