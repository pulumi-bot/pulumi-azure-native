// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
type ApplicationInsightsComponentAnalyticsItemProperties struct {
	// A function alias, used when the type of the item is Function
	FunctionAlias *string `pulumi:"functionAlias"`
}

// ApplicationInsightsComponentAnalyticsItemPropertiesInput is an input type that accepts ApplicationInsightsComponentAnalyticsItemPropertiesArgs and ApplicationInsightsComponentAnalyticsItemPropertiesOutput values.
// You can construct a concrete instance of `ApplicationInsightsComponentAnalyticsItemPropertiesInput` via:
//
//          ApplicationInsightsComponentAnalyticsItemPropertiesArgs{...}
type ApplicationInsightsComponentAnalyticsItemPropertiesInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput
}

// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
type ApplicationInsightsComponentAnalyticsItemPropertiesArgs struct {
	// A function alias, used when the type of the item is Function
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
}

func (ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput)
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput).ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx)
}

// ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput is an input type that accepts ApplicationInsightsComponentAnalyticsItemPropertiesArgs, ApplicationInsightsComponentAnalyticsItemPropertiesPtr and ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput values.
// You can construct a concrete instance of `ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput` via:
//
//          ApplicationInsightsComponentAnalyticsItemPropertiesArgs{...}
//
//  or:
//
//          nil
type ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
}

type applicationInsightsComponentAnalyticsItemPropertiesPtrType ApplicationInsightsComponentAnalyticsItemPropertiesArgs

func ApplicationInsightsComponentAnalyticsItemPropertiesPtr(v *ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput {
	return (*applicationInsightsComponentAnalyticsItemPropertiesPtrType)(v)
}

func (*applicationInsightsComponentAnalyticsItemPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput)
}

// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
type ApplicationInsightsComponentAnalyticsItemPropertiesResponse struct {
	// A function alias, used when the type of the item is Function
	FunctionAlias *string `pulumi:"functionAlias"`
}

// A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
type ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		return &v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput)
}

// A function alias, used when the type of the item is Function
func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) Elem() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		return *v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

// A function alias, used when the type of the item is Function
func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionAlias
	}).(pulumi.StringPtrOutput)
}

// The private link scope resource reference.
type PrivateLinkScopedResourceResponse struct {
	// The full resource Id of the private link scope resource.
	ResourceId *string `pulumi:"resourceId"`
	// The private link scope unique Identifier.
	ScopeId *string `pulumi:"scopeId"`
}

// The private link scope resource reference.
type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

// The full resource Id of the private link scope resource.
func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The private link scope unique Identifier.
func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
}

// Geo-physical location to run a web test from. You must specify one or more locations for the test to run from.
type WebTestGeolocation struct {
	// Location ID for the webtest to run from.
	Location *string `pulumi:"location"`
}

// WebTestGeolocationInput is an input type that accepts WebTestGeolocationArgs and WebTestGeolocationOutput values.
// You can construct a concrete instance of `WebTestGeolocationInput` via:
//
//          WebTestGeolocationArgs{...}
type WebTestGeolocationInput interface {
	pulumi.Input

	ToWebTestGeolocationOutput() WebTestGeolocationOutput
	ToWebTestGeolocationOutputWithContext(context.Context) WebTestGeolocationOutput
}

// Geo-physical location to run a web test from. You must specify one or more locations for the test to run from.
type WebTestGeolocationArgs struct {
	// Location ID for the webtest to run from.
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (WebTestGeolocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutput() WebTestGeolocationOutput {
	return i.ToWebTestGeolocationOutputWithContext(context.Background())
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutputWithContext(ctx context.Context) WebTestGeolocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationOutput)
}

// WebTestGeolocationArrayInput is an input type that accepts WebTestGeolocationArray and WebTestGeolocationArrayOutput values.
// You can construct a concrete instance of `WebTestGeolocationArrayInput` via:
//
//          WebTestGeolocationArray{ WebTestGeolocationArgs{...} }
type WebTestGeolocationArrayInput interface {
	pulumi.Input

	ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput
	ToWebTestGeolocationArrayOutputWithContext(context.Context) WebTestGeolocationArrayOutput
}

type WebTestGeolocationArray []WebTestGeolocationInput

func (WebTestGeolocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput {
	return i.ToWebTestGeolocationArrayOutputWithContext(context.Background())
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutputWithContext(ctx context.Context) WebTestGeolocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationArrayOutput)
}

// Geo-physical location to run a web test from. You must specify one or more locations for the test to run from.
type WebTestGeolocationResponse struct {
	// Location ID for the webtest to run from.
	Location *string `pulumi:"location"`
}

// Geo-physical location to run a web test from. You must specify one or more locations for the test to run from.
type WebTestGeolocationResponseOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput {
	return o
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutputWithContext(ctx context.Context) WebTestGeolocationResponseOutput {
	return o
}

// Location ID for the webtest to run from.
func (o WebTestGeolocationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestGeolocationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type WebTestGeolocationResponseArrayOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutputWithContext(ctx context.Context) WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) Index(i pulumi.IntInput) WebTestGeolocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTestGeolocationResponse {
		return vs[0].([]WebTestGeolocationResponse)[vs[1].(int)]
	}).(WebTestGeolocationResponseOutput)
}

// An XML configuration specification for a WebTest.
type WebTestPropertiesConfiguration struct {
	// The XML specification of a WebTest to run against an application.
	WebTest *string `pulumi:"webTest"`
}

// WebTestPropertiesConfigurationInput is an input type that accepts WebTestPropertiesConfigurationArgs and WebTestPropertiesConfigurationOutput values.
// You can construct a concrete instance of `WebTestPropertiesConfigurationInput` via:
//
//          WebTestPropertiesConfigurationArgs{...}
type WebTestPropertiesConfigurationInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput
	ToWebTestPropertiesConfigurationOutputWithContext(context.Context) WebTestPropertiesConfigurationOutput
}

// An XML configuration specification for a WebTest.
type WebTestPropertiesConfigurationArgs struct {
	// The XML specification of a WebTest to run against an application.
	WebTest pulumi.StringPtrInput `pulumi:"webTest"`
}

func (WebTestPropertiesConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput {
	return i.ToWebTestPropertiesConfigurationOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput)
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput).ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx)
}

// WebTestPropertiesConfigurationPtrInput is an input type that accepts WebTestPropertiesConfigurationArgs, WebTestPropertiesConfigurationPtr and WebTestPropertiesConfigurationPtrOutput values.
// You can construct a concrete instance of `WebTestPropertiesConfigurationPtrInput` via:
//
//          WebTestPropertiesConfigurationArgs{...}
//
//  or:
//
//          nil
type WebTestPropertiesConfigurationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput
	ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Context) WebTestPropertiesConfigurationPtrOutput
}

type webTestPropertiesConfigurationPtrType WebTestPropertiesConfigurationArgs

func WebTestPropertiesConfigurationPtr(v *WebTestPropertiesConfigurationArgs) WebTestPropertiesConfigurationPtrInput {
	return (*webTestPropertiesConfigurationPtrType)(v)
}

func (*webTestPropertiesConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationPtrOutput)
}

// An XML configuration specification for a WebTest.
type WebTestPropertiesResponseConfiguration struct {
	// The XML specification of a WebTest to run against an application.
	WebTest *string `pulumi:"webTest"`
}

// An XML configuration specification for a WebTest.
type WebTestPropertiesResponseConfigurationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseConfiguration) *WebTestPropertiesResponseConfiguration {
		return &v
	}).(WebTestPropertiesResponseConfigurationPtrOutput)
}

// The XML specification of a WebTest to run against an application.
func (o WebTestPropertiesResponseConfigurationOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseConfiguration) *string { return v.WebTest }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) Elem() WebTestPropertiesResponseConfigurationOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) WebTestPropertiesResponseConfiguration { return *v }).(WebTestPropertiesResponseConfigurationOutput)
}

// The XML specification of a WebTest to run against an application.
func (o WebTestPropertiesResponseConfigurationPtrOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.WebTest
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationPtrOutput{})
}
