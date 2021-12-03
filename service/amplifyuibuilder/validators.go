// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateComponent struct {
}

func (*validateOpCreateComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTheme struct {
}

func (*validateOpCreateTheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateThemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateThemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteComponent struct {
}

func (*validateOpDeleteComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTheme struct {
}

func (*validateOpDeleteTheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteThemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteThemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExchangeCodeForToken struct {
}

func (*validateOpExchangeCodeForToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExchangeCodeForToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExchangeCodeForTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExchangeCodeForTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportComponents struct {
}

func (*validateOpExportComponents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportComponents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportComponentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportComponentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportThemes struct {
}

func (*validateOpExportThemes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportThemes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportThemesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportThemesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetComponent struct {
}

func (*validateOpGetComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTheme struct {
}

func (*validateOpGetTheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetThemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetThemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListComponents struct {
}

func (*validateOpListComponents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListComponents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListComponentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListComponentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListThemes struct {
}

func (*validateOpListThemes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListThemes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListThemesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListThemesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRefreshToken struct {
}

func (*validateOpRefreshToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRefreshToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RefreshTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRefreshTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateComponent struct {
}

func (*validateOpUpdateComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateTheme struct {
}

func (*validateOpUpdateTheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateTheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateThemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateThemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateComponent{}, middleware.After)
}

func addOpCreateThemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTheme{}, middleware.After)
}

func addOpDeleteComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteComponent{}, middleware.After)
}

func addOpDeleteThemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTheme{}, middleware.After)
}

func addOpExchangeCodeForTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExchangeCodeForToken{}, middleware.After)
}

func addOpExportComponentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportComponents{}, middleware.After)
}

func addOpExportThemesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportThemes{}, middleware.After)
}

func addOpGetComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetComponent{}, middleware.After)
}

func addOpGetThemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTheme{}, middleware.After)
}

func addOpListComponentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListComponents{}, middleware.After)
}

func addOpListThemesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListThemes{}, middleware.After)
}

func addOpRefreshTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRefreshToken{}, middleware.After)
}

func addOpUpdateComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateComponent{}, middleware.After)
}

func addOpUpdateThemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateTheme{}, middleware.After)
}

func validateComponentChild(v *types.ComponentChild) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentChild"}
	if v.ComponentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentType"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Properties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Properties"))
	} else if v.Properties != nil {
		if err := validateComponentProperties(v.Properties); err != nil {
			invalidParams.AddNested("Properties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Children != nil {
		if err := validateComponentChildList(v.Children); err != nil {
			invalidParams.AddNested("Children", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentChildList(v []types.ComponentChild) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentChildList"}
	for i := range v {
		if err := validateComponentChild(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentCollectionProperties(v map[string]types.ComponentDataConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentCollectionProperties"}
	for key := range v {
		value := v[key]
		if err := validateComponentDataConfiguration(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentConditionProperty(v *types.ComponentConditionProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentConditionProperty"}
	if v.Then != nil {
		if err := validateComponentProperty(v.Then); err != nil {
			invalidParams.AddNested("Then", err.(smithy.InvalidParamsError))
		}
	}
	if v.Else != nil {
		if err := validateComponentProperty(v.Else); err != nil {
			invalidParams.AddNested("Else", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentDataConfiguration(v *types.ComponentDataConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentDataConfiguration"}
	if v.Model == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Model"))
	}
	if v.Sort != nil {
		if err := validateSortPropertyList(v.Sort); err != nil {
			invalidParams.AddNested("Sort", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentProperties(v map[string]types.ComponentProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentProperties"}
	for key := range v {
		value := v[key]
		if err := validateComponentProperty(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentProperty(v *types.ComponentProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentProperty"}
	if v.BindingProperties != nil {
		if err := validateComponentPropertyBindingProperties(v.BindingProperties); err != nil {
			invalidParams.AddNested("BindingProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.CollectionBindingProperties != nil {
		if err := validateComponentPropertyBindingProperties(v.CollectionBindingProperties); err != nil {
			invalidParams.AddNested("CollectionBindingProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Bindings != nil {
		if err := validateFormBindings(v.Bindings); err != nil {
			invalidParams.AddNested("Bindings", err.(smithy.InvalidParamsError))
		}
	}
	if v.Concat != nil {
		if err := validateComponentPropertyList(v.Concat); err != nil {
			invalidParams.AddNested("Concat", err.(smithy.InvalidParamsError))
		}
	}
	if v.Condition != nil {
		if err := validateComponentConditionProperty(v.Condition); err != nil {
			invalidParams.AddNested("Condition", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentPropertyBindingProperties(v *types.ComponentPropertyBindingProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentPropertyBindingProperties"}
	if v.Property == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Property"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateComponentPropertyList(v []types.ComponentProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ComponentPropertyList"}
	for i := range v {
		if err := validateComponentProperty(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateComponentData(v *types.CreateComponentData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateComponentData"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ComponentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentType"))
	}
	if v.Properties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Properties"))
	} else if v.Properties != nil {
		if err := validateComponentProperties(v.Properties); err != nil {
			invalidParams.AddNested("Properties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Children != nil {
		if err := validateComponentChildList(v.Children); err != nil {
			invalidParams.AddNested("Children", err.(smithy.InvalidParamsError))
		}
	}
	if v.Variants == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Variants"))
	}
	if v.Overrides == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Overrides"))
	}
	if v.BindingProperties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BindingProperties"))
	}
	if v.CollectionProperties != nil {
		if err := validateComponentCollectionProperties(v.CollectionProperties); err != nil {
			invalidParams.AddNested("CollectionProperties", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateThemeData(v *types.CreateThemeData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateThemeData"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExchangeCodeForTokenRequestBody(v *types.ExchangeCodeForTokenRequestBody) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExchangeCodeForTokenRequestBody"}
	if v.Code == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Code"))
	}
	if v.RedirectUri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RedirectUri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFormBindingElement(v *types.FormBindingElement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FormBindingElement"}
	if v.Element == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Element"))
	}
	if v.Property == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Property"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFormBindings(v map[string]types.FormBindingElement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FormBindings"}
	for key := range v {
		value := v[key]
		if err := validateFormBindingElement(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRefreshTokenRequestBody(v *types.RefreshTokenRequestBody) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RefreshTokenRequestBody"}
	if v.Token == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Token"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSortProperty(v *types.SortProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SortProperty"}
	if v.Field == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Field"))
	}
	if len(v.Direction) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Direction"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSortPropertyList(v []types.SortProperty) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SortPropertyList"}
	for i := range v {
		if err := validateSortProperty(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateComponentData(v *types.UpdateComponentData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateComponentData"}
	if v.Properties != nil {
		if err := validateComponentProperties(v.Properties); err != nil {
			invalidParams.AddNested("Properties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Children != nil {
		if err := validateComponentChildList(v.Children); err != nil {
			invalidParams.AddNested("Children", err.(smithy.InvalidParamsError))
		}
	}
	if v.CollectionProperties != nil {
		if err := validateComponentCollectionProperties(v.CollectionProperties); err != nil {
			invalidParams.AddNested("CollectionProperties", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateThemeData(v *types.UpdateThemeData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateThemeData"}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateComponentInput(v *CreateComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateComponentInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.ComponentToCreate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentToCreate"))
	} else if v.ComponentToCreate != nil {
		if err := validateCreateComponentData(v.ComponentToCreate); err != nil {
			invalidParams.AddNested("ComponentToCreate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateThemeInput(v *CreateThemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateThemeInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.ThemeToCreate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThemeToCreate"))
	} else if v.ThemeToCreate != nil {
		if err := validateCreateThemeData(v.ThemeToCreate); err != nil {
			invalidParams.AddNested("ThemeToCreate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteComponentInput(v *DeleteComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteComponentInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteThemeInput(v *DeleteThemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteThemeInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExchangeCodeForTokenInput(v *ExchangeCodeForTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExchangeCodeForTokenInput"}
	if len(v.Provider) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Provider"))
	}
	if v.Request == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Request"))
	} else if v.Request != nil {
		if err := validateExchangeCodeForTokenRequestBody(v.Request); err != nil {
			invalidParams.AddNested("Request", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportComponentsInput(v *ExportComponentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportComponentsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportThemesInput(v *ExportThemesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportThemesInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetComponentInput(v *GetComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetComponentInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetThemeInput(v *GetThemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetThemeInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListComponentsInput(v *ListComponentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListComponentsInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListThemesInput(v *ListThemesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListThemesInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRefreshTokenInput(v *RefreshTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RefreshTokenInput"}
	if len(v.Provider) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Provider"))
	}
	if v.RefreshTokenBody == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RefreshTokenBody"))
	} else if v.RefreshTokenBody != nil {
		if err := validateRefreshTokenRequestBody(v.RefreshTokenBody); err != nil {
			invalidParams.AddNested("RefreshTokenBody", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateComponentInput(v *UpdateComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateComponentInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.UpdatedComponent == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UpdatedComponent"))
	} else if v.UpdatedComponent != nil {
		if err := validateUpdateComponentData(v.UpdatedComponent); err != nil {
			invalidParams.AddNested("UpdatedComponent", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateThemeInput(v *UpdateThemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateThemeInput"}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.UpdatedTheme == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UpdatedTheme"))
	} else if v.UpdatedTheme != nil {
		if err := validateUpdateThemeData(v.UpdatedTheme); err != nil {
			invalidParams.AddNested("UpdatedTheme", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
