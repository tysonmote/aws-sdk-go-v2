// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsAwsjson10_serializeOpCreateFHIRDatastore struct {
}

func (*awsAwsjson10_serializeOpCreateFHIRDatastore) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpCreateFHIRDatastore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateFHIRDatastoreInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.CreateFHIRDatastore")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentCreateFHIRDatastoreInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDeleteFHIRDatastore struct {
}

func (*awsAwsjson10_serializeOpDeleteFHIRDatastore) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDeleteFHIRDatastore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteFHIRDatastoreInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.DeleteFHIRDatastore")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDeleteFHIRDatastoreInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeFHIRDatastore struct {
}

func (*awsAwsjson10_serializeOpDescribeFHIRDatastore) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeFHIRDatastore) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeFHIRDatastoreInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.DescribeFHIRDatastore")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeFHIRDatastoreInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeFHIRExportJob struct {
}

func (*awsAwsjson10_serializeOpDescribeFHIRExportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeFHIRExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeFHIRExportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.DescribeFHIRExportJob")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeFHIRExportJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpDescribeFHIRImportJob struct {
}

func (*awsAwsjson10_serializeOpDescribeFHIRImportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpDescribeFHIRImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeFHIRImportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.DescribeFHIRImportJob")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentDescribeFHIRImportJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpListFHIRDatastores struct {
}

func (*awsAwsjson10_serializeOpListFHIRDatastores) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpListFHIRDatastores) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListFHIRDatastoresInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.ListFHIRDatastores")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentListFHIRDatastoresInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpStartFHIRExportJob struct {
}

func (*awsAwsjson10_serializeOpStartFHIRExportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpStartFHIRExportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartFHIRExportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.StartFHIRExportJob")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentStartFHIRExportJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson10_serializeOpStartFHIRImportJob struct {
}

func (*awsAwsjson10_serializeOpStartFHIRImportJob) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson10_serializeOpStartFHIRImportJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartFHIRImportJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.0")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("HealthLake.StartFHIRImportJob")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson10_serializeOpDocumentStartFHIRImportJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson10_serializeDocumentDatastoreFilter(v *types.DatastoreFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CreatedAfter != nil {
		ok := object.Key("CreatedAfter")
		ok.Double(smithytime.FormatEpochSeconds(*v.CreatedAfter))
	}

	if v.CreatedBefore != nil {
		ok := object.Key("CreatedBefore")
		ok.Double(smithytime.FormatEpochSeconds(*v.CreatedBefore))
	}

	if v.DatastoreName != nil {
		ok := object.Key("DatastoreName")
		ok.String(*v.DatastoreName)
	}

	if len(v.DatastoreStatus) > 0 {
		ok := object.Key("DatastoreStatus")
		ok.String(string(v.DatastoreStatus))
	}

	return nil
}

func awsAwsjson10_serializeDocumentInputDataConfig(v types.InputDataConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.InputDataConfigMemberS3Uri:
		av := object.Key("S3Uri")
		av.String(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsAwsjson10_serializeDocumentOutputDataConfig(v types.OutputDataConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.OutputDataConfigMemberS3Uri:
		av := object.Key("S3Uri")
		av.String(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsAwsjson10_serializeDocumentPreloadDataConfig(v *types.PreloadDataConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.PreloadDataType) > 0 {
		ok := object.Key("PreloadDataType")
		ok.String(string(v.PreloadDataType))
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentCreateFHIRDatastoreInput(v *CreateFHIRDatastoreInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.DatastoreName != nil {
		ok := object.Key("DatastoreName")
		ok.String(*v.DatastoreName)
	}

	if len(v.DatastoreTypeVersion) > 0 {
		ok := object.Key("DatastoreTypeVersion")
		ok.String(string(v.DatastoreTypeVersion))
	}

	if v.PreloadDataConfig != nil {
		ok := object.Key("PreloadDataConfig")
		if err := awsAwsjson10_serializeDocumentPreloadDataConfig(v.PreloadDataConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDeleteFHIRDatastoreInput(v *DeleteFHIRDatastoreInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeFHIRDatastoreInput(v *DescribeFHIRDatastoreInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeFHIRExportJobInput(v *DescribeFHIRExportJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	if v.JobId != nil {
		ok := object.Key("JobId")
		ok.String(*v.JobId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentDescribeFHIRImportJobInput(v *DescribeFHIRImportJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	if v.JobId != nil {
		ok := object.Key("JobId")
		ok.String(*v.JobId)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentListFHIRDatastoresInput(v *ListFHIRDatastoresInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filter != nil {
		ok := object.Key("Filter")
		if err := awsAwsjson10_serializeDocumentDatastoreFilter(v.Filter, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentStartFHIRExportJobInput(v *StartFHIRExportJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.DataAccessRoleArn != nil {
		ok := object.Key("DataAccessRoleArn")
		ok.String(*v.DataAccessRoleArn)
	}

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	if v.JobName != nil {
		ok := object.Key("JobName")
		ok.String(*v.JobName)
	}

	if v.OutputDataConfig != nil {
		ok := object.Key("OutputDataConfig")
		if err := awsAwsjson10_serializeDocumentOutputDataConfig(v.OutputDataConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson10_serializeOpDocumentStartFHIRImportJobInput(v *StartFHIRImportJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.DataAccessRoleArn != nil {
		ok := object.Key("DataAccessRoleArn")
		ok.String(*v.DataAccessRoleArn)
	}

	if v.DatastoreId != nil {
		ok := object.Key("DatastoreId")
		ok.String(*v.DatastoreId)
	}

	if v.InputDataConfig != nil {
		ok := object.Key("InputDataConfig")
		if err := awsAwsjson10_serializeDocumentInputDataConfig(v.InputDataConfig, ok); err != nil {
			return err
		}
	}

	if v.JobName != nil {
		ok := object.Key("JobName")
		ok.String(*v.JobName)
	}

	return nil
}