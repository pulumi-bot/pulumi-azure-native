// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the runbook property type.
type ContentHash struct {
	// Gets or sets the content hash algorithm used to hash the content.
	Algorithm string `pulumi:"algorithm"`
	// Gets or sets expected hash value of the content.
	Value string `pulumi:"value"`
}

// ContentHashInput is an input type that accepts ContentHashArgs and ContentHashOutput values.
// You can construct a concrete instance of `ContentHashInput` via:
//
//          ContentHashArgs{...}
type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

// Definition of the runbook property type.
type ContentHashArgs struct {
	// Gets or sets the content hash algorithm used to hash the content.
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	// Gets or sets expected hash value of the content.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ContentHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (i ContentHashArgs) ToContentHashOutput() ContentHashOutput {
	return i.ToContentHashOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput)
}

func (i ContentHashArgs) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput).ToContentHashPtrOutputWithContext(ctx)
}

// ContentHashPtrInput is an input type that accepts ContentHashArgs, ContentHashPtr and ContentHashPtrOutput values.
// You can construct a concrete instance of `ContentHashPtrInput` via:
//
//          ContentHashArgs{...}
//
//  or:
//
//          nil
type ContentHashPtrInput interface {
	pulumi.Input

	ToContentHashPtrOutput() ContentHashPtrOutput
	ToContentHashPtrOutputWithContext(context.Context) ContentHashPtrOutput
}

type contentHashPtrType ContentHashArgs

func ContentHashPtr(v *ContentHashArgs) ContentHashPtrInput {
	return (*contentHashPtrType)(v)
}

func (*contentHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (i *contentHashPtrType) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i *contentHashPtrType) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashPtrOutput)
}

// Definition of the runbook property type.
type ContentHashOutput struct{ *pulumi.OutputState }

func (ContentHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (o ContentHashOutput) ToContentHashOutput() ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o.ToContentHashPtrOutputWithContext(context.Background())
}

func (o ContentHashOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o.ApplyT(func(v ContentHash) *ContentHash {
		return &v
	}).(ContentHashPtrOutput)
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashPtrOutput struct{ *pulumi.OutputState }

func (ContentHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (o ContentHashPtrOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) Elem() ContentHashOutput {
	return o.ApplyT(func(v *ContentHash) ContentHash { return *v }).(ContentHashOutput)
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

// Definition of the content source.
type ContentSource struct {
	// Gets or sets the hash.
	Hash *ContentHash `pulumi:"hash"`
	// Gets or sets the content source type.
	Type *string `pulumi:"type"`
	// Gets or sets the value of the content. This is based on the content source type.
	Value *string `pulumi:"value"`
	// Gets or sets the version of the content.
	Version *string `pulumi:"version"`
}

// ContentSourceInput is an input type that accepts ContentSourceArgs and ContentSourceOutput values.
// You can construct a concrete instance of `ContentSourceInput` via:
//
//          ContentSourceArgs{...}
type ContentSourceInput interface {
	pulumi.Input

	ToContentSourceOutput() ContentSourceOutput
	ToContentSourceOutputWithContext(context.Context) ContentSourceOutput
}

// Definition of the content source.
type ContentSourceArgs struct {
	// Gets or sets the hash.
	Hash ContentHashPtrInput `pulumi:"hash"`
	// Gets or sets the content source type.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Gets or sets the value of the content. This is based on the content source type.
	Value pulumi.StringPtrInput `pulumi:"value"`
	// Gets or sets the version of the content.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (i ContentSourceArgs) ToContentSourceOutput() ContentSourceOutput {
	return i.ToContentSourceOutputWithContext(context.Background())
}

func (i ContentSourceArgs) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceOutput)
}

func (i ContentSourceArgs) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return i.ToContentSourcePtrOutputWithContext(context.Background())
}

func (i ContentSourceArgs) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceOutput).ToContentSourcePtrOutputWithContext(ctx)
}

// ContentSourcePtrInput is an input type that accepts ContentSourceArgs, ContentSourcePtr and ContentSourcePtrOutput values.
// You can construct a concrete instance of `ContentSourcePtrInput` via:
//
//          ContentSourceArgs{...}
//
//  or:
//
//          nil
type ContentSourcePtrInput interface {
	pulumi.Input

	ToContentSourcePtrOutput() ContentSourcePtrOutput
	ToContentSourcePtrOutputWithContext(context.Context) ContentSourcePtrOutput
}

type contentSourcePtrType ContentSourceArgs

func ContentSourcePtr(v *ContentSourceArgs) ContentSourcePtrInput {
	return (*contentSourcePtrType)(v)
}

func (*contentSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSource)(nil)).Elem()
}

func (i *contentSourcePtrType) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return i.ToContentSourcePtrOutputWithContext(context.Background())
}

func (i *contentSourcePtrType) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourcePtrOutput)
}

// Definition of the content source.
type ContentSourceOutput struct{ *pulumi.OutputState }

func (ContentSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (o ContentSourceOutput) ToContentSourceOutput() ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return o.ToContentSourcePtrOutputWithContext(context.Background())
}

func (o ContentSourceOutput) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return o.ApplyT(func(v ContentSource) *ContentSource {
		return &v
	}).(ContentSourcePtrOutput)
}

// Gets or sets the hash.
func (o ContentSourceOutput) Hash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentSource) *ContentHash { return v.Hash }).(ContentHashPtrOutput)
}

// Gets or sets the content source type.
func (o ContentSourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Gets or sets the value of the content. This is based on the content source type.
func (o ContentSourceOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Value }).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentSourcePtrOutput struct{ *pulumi.OutputState }

func (ContentSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSource)(nil)).Elem()
}

func (o ContentSourcePtrOutput) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return o
}

func (o ContentSourcePtrOutput) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return o
}

func (o ContentSourcePtrOutput) Elem() ContentSourceOutput {
	return o.ApplyT(func(v *ContentSource) ContentSource { return *v }).(ContentSourceOutput)
}

// Gets or sets the hash.
func (o ContentSourcePtrOutput) Hash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentSource) *ContentHash {
		if v == nil {
			return nil
		}
		return v.Hash
	}).(ContentHashPtrOutput)
}

// Gets or sets the content source type.
func (o ContentSourcePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the value of the content. This is based on the content source type.
func (o ContentSourcePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentSourcePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationProperty struct {
	// Gets or sets the name of the Dsc configuration.
	Name *string `pulumi:"name"`
}

// DscConfigurationAssociationPropertyInput is an input type that accepts DscConfigurationAssociationPropertyArgs and DscConfigurationAssociationPropertyOutput values.
// You can construct a concrete instance of `DscConfigurationAssociationPropertyInput` via:
//
//          DscConfigurationAssociationPropertyArgs{...}
type DscConfigurationAssociationPropertyInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput
	ToDscConfigurationAssociationPropertyOutputWithContext(context.Context) DscConfigurationAssociationPropertyOutput
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationPropertyArgs struct {
	// Gets or sets the name of the Dsc configuration.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DscConfigurationAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return i.ToDscConfigurationAssociationPropertyOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyOutput)
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return i.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyOutput).ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx)
}

// DscConfigurationAssociationPropertyPtrInput is an input type that accepts DscConfigurationAssociationPropertyArgs, DscConfigurationAssociationPropertyPtr and DscConfigurationAssociationPropertyPtrOutput values.
// You can construct a concrete instance of `DscConfigurationAssociationPropertyPtrInput` via:
//
//          DscConfigurationAssociationPropertyArgs{...}
//
//  or:
//
//          nil
type DscConfigurationAssociationPropertyPtrInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput
	ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Context) DscConfigurationAssociationPropertyPtrOutput
}

type dscConfigurationAssociationPropertyPtrType DscConfigurationAssociationPropertyArgs

func DscConfigurationAssociationPropertyPtr(v *DscConfigurationAssociationPropertyArgs) DscConfigurationAssociationPropertyPtrInput {
	return (*dscConfigurationAssociationPropertyPtrType)(v)
}

func (*dscConfigurationAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationProperty)(nil)).Elem()
}

func (i *dscConfigurationAssociationPropertyPtrType) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return i.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *dscConfigurationAssociationPropertyPtrType) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyPtrOutput)
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationPropertyOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return o.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationProperty) *DscConfigurationAssociationProperty {
		return &v
	}).(DscConfigurationAssociationPropertyPtrOutput)
}

// Gets or sets the name of the Dsc configuration.
func (o DscConfigurationAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationProperty)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyPtrOutput) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyPtrOutput) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyPtrOutput) Elem() DscConfigurationAssociationPropertyOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationProperty) DscConfigurationAssociationProperty { return *v }).(DscConfigurationAssociationPropertyOutput)
}

// Gets or sets the name of the Dsc configuration.
func (o DscConfigurationAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationPropertyResponse struct {
	// Gets or sets the name of the Dsc configuration.
	Name *string `pulumi:"name"`
}

// DscConfigurationAssociationPropertyResponseInput is an input type that accepts DscConfigurationAssociationPropertyResponseArgs and DscConfigurationAssociationPropertyResponseOutput values.
// You can construct a concrete instance of `DscConfigurationAssociationPropertyResponseInput` via:
//
//          DscConfigurationAssociationPropertyResponseArgs{...}
type DscConfigurationAssociationPropertyResponseInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput
	ToDscConfigurationAssociationPropertyResponseOutputWithContext(context.Context) DscConfigurationAssociationPropertyResponseOutput
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationPropertyResponseArgs struct {
	// Gets or sets the name of the Dsc configuration.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DscConfigurationAssociationPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput {
	return i.ToDscConfigurationAssociationPropertyResponseOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponseOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponseOutput)
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return i.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponseOutput).ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx)
}

// DscConfigurationAssociationPropertyResponsePtrInput is an input type that accepts DscConfigurationAssociationPropertyResponseArgs, DscConfigurationAssociationPropertyResponsePtr and DscConfigurationAssociationPropertyResponsePtrOutput values.
// You can construct a concrete instance of `DscConfigurationAssociationPropertyResponsePtrInput` via:
//
//          DscConfigurationAssociationPropertyResponseArgs{...}
//
//  or:
//
//          nil
type DscConfigurationAssociationPropertyResponsePtrInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput
	ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Context) DscConfigurationAssociationPropertyResponsePtrOutput
}

type dscConfigurationAssociationPropertyResponsePtrType DscConfigurationAssociationPropertyResponseArgs

func DscConfigurationAssociationPropertyResponsePtr(v *DscConfigurationAssociationPropertyResponseArgs) DscConfigurationAssociationPropertyResponsePtrInput {
	return (*dscConfigurationAssociationPropertyResponsePtrType)(v)
}

func (*dscConfigurationAssociationPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (i *dscConfigurationAssociationPropertyResponsePtrType) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return i.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *dscConfigurationAssociationPropertyResponsePtrType) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// The Dsc configuration property associated with the entity.
type DscConfigurationAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationPropertyResponse) *DscConfigurationAssociationPropertyResponse {
		return &v
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// Gets or sets the name of the Dsc configuration.
func (o DscConfigurationAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) Elem() DscConfigurationAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) DscConfigurationAssociationPropertyResponse {
		return *v
	}).(DscConfigurationAssociationPropertyResponseOutput)
}

// Gets or sets the name of the Dsc configuration.
func (o DscConfigurationAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// Definition of the dsc node configuration.
type DscNodeConfigurationType struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the configuration properties.
	Properties DscNodeConfigurationPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// DscNodeConfigurationTypeInput is an input type that accepts DscNodeConfigurationTypeArgs and DscNodeConfigurationTypeOutput values.
// You can construct a concrete instance of `DscNodeConfigurationTypeInput` via:
//
//          DscNodeConfigurationTypeArgs{...}
type DscNodeConfigurationTypeInput interface {
	pulumi.Input

	ToDscNodeConfigurationTypeOutput() DscNodeConfigurationTypeOutput
	ToDscNodeConfigurationTypeOutputWithContext(context.Context) DscNodeConfigurationTypeOutput
}

// Definition of the dsc node configuration.
type DscNodeConfigurationTypeArgs struct {
	// The name of the resource
	Name pulumi.StringInput `pulumi:"name"`
	// Gets or sets the configuration properties.
	Properties DscNodeConfigurationPropertiesResponseInput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (DscNodeConfigurationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfigurationType)(nil)).Elem()
}

func (i DscNodeConfigurationTypeArgs) ToDscNodeConfigurationTypeOutput() DscNodeConfigurationTypeOutput {
	return i.ToDscNodeConfigurationTypeOutputWithContext(context.Background())
}

func (i DscNodeConfigurationTypeArgs) ToDscNodeConfigurationTypeOutputWithContext(ctx context.Context) DscNodeConfigurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationTypeOutput)
}

// Definition of the dsc node configuration.
type DscNodeConfigurationTypeOutput struct{ *pulumi.OutputState }

func (DscNodeConfigurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfigurationType)(nil)).Elem()
}

func (o DscNodeConfigurationTypeOutput) ToDscNodeConfigurationTypeOutput() DscNodeConfigurationTypeOutput {
	return o
}

func (o DscNodeConfigurationTypeOutput) ToDscNodeConfigurationTypeOutputWithContext(ctx context.Context) DscNodeConfigurationTypeOutput {
	return o
}

// The name of the resource
func (o DscNodeConfigurationTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DscNodeConfigurationType) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the configuration properties.
func (o DscNodeConfigurationTypeOutput) Properties() DscNodeConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v DscNodeConfigurationType) DscNodeConfigurationPropertiesResponse { return v.Properties }).(DscNodeConfigurationPropertiesResponseOutput)
}

// The type of the resource.
func (o DscNodeConfigurationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DscNodeConfigurationType) string { return v.Type }).(pulumi.StringOutput)
}

// Properties for the DscNodeConfiguration
type DscNodeConfigurationPropertiesResponse struct {
	// Gets or sets the configuration of the node.
	Configuration *DscConfigurationAssociationPropertyResponse `pulumi:"configuration"`
	// Gets or sets creation time.
	CreationTime *string `pulumi:"creationTime"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild *bool `pulumi:"incrementNodeConfigurationBuild"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Number of nodes with this node configuration assigned
	NodeCount *int `pulumi:"nodeCount"`
	// Source of node configuration.
	Source *string `pulumi:"source"`
}

// DscNodeConfigurationPropertiesResponseInput is an input type that accepts DscNodeConfigurationPropertiesResponseArgs and DscNodeConfigurationPropertiesResponseOutput values.
// You can construct a concrete instance of `DscNodeConfigurationPropertiesResponseInput` via:
//
//          DscNodeConfigurationPropertiesResponseArgs{...}
type DscNodeConfigurationPropertiesResponseInput interface {
	pulumi.Input

	ToDscNodeConfigurationPropertiesResponseOutput() DscNodeConfigurationPropertiesResponseOutput
	ToDscNodeConfigurationPropertiesResponseOutputWithContext(context.Context) DscNodeConfigurationPropertiesResponseOutput
}

// Properties for the DscNodeConfiguration
type DscNodeConfigurationPropertiesResponseArgs struct {
	// Gets or sets the configuration of the node.
	Configuration DscConfigurationAssociationPropertyResponsePtrInput `pulumi:"configuration"`
	// Gets or sets creation time.
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild pulumi.BoolPtrInput `pulumi:"incrementNodeConfigurationBuild"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrInput `pulumi:"lastModifiedTime"`
	// Number of nodes with this node configuration assigned
	NodeCount pulumi.IntPtrInput `pulumi:"nodeCount"`
	// Source of node configuration.
	Source pulumi.StringPtrInput `pulumi:"source"`
}

func (DscNodeConfigurationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfigurationPropertiesResponse)(nil)).Elem()
}

func (i DscNodeConfigurationPropertiesResponseArgs) ToDscNodeConfigurationPropertiesResponseOutput() DscNodeConfigurationPropertiesResponseOutput {
	return i.ToDscNodeConfigurationPropertiesResponseOutputWithContext(context.Background())
}

func (i DscNodeConfigurationPropertiesResponseArgs) ToDscNodeConfigurationPropertiesResponseOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationPropertiesResponseOutput)
}

func (i DscNodeConfigurationPropertiesResponseArgs) ToDscNodeConfigurationPropertiesResponsePtrOutput() DscNodeConfigurationPropertiesResponsePtrOutput {
	return i.ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DscNodeConfigurationPropertiesResponseArgs) ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationPropertiesResponseOutput).ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(ctx)
}

// DscNodeConfigurationPropertiesResponsePtrInput is an input type that accepts DscNodeConfigurationPropertiesResponseArgs, DscNodeConfigurationPropertiesResponsePtr and DscNodeConfigurationPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `DscNodeConfigurationPropertiesResponsePtrInput` via:
//
//          DscNodeConfigurationPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type DscNodeConfigurationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDscNodeConfigurationPropertiesResponsePtrOutput() DscNodeConfigurationPropertiesResponsePtrOutput
	ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(context.Context) DscNodeConfigurationPropertiesResponsePtrOutput
}

type dscNodeConfigurationPropertiesResponsePtrType DscNodeConfigurationPropertiesResponseArgs

func DscNodeConfigurationPropertiesResponsePtr(v *DscNodeConfigurationPropertiesResponseArgs) DscNodeConfigurationPropertiesResponsePtrInput {
	return (*dscNodeConfigurationPropertiesResponsePtrType)(v)
}

func (*dscNodeConfigurationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DscNodeConfigurationPropertiesResponse)(nil)).Elem()
}

func (i *dscNodeConfigurationPropertiesResponsePtrType) ToDscNodeConfigurationPropertiesResponsePtrOutput() DscNodeConfigurationPropertiesResponsePtrOutput {
	return i.ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *dscNodeConfigurationPropertiesResponsePtrType) ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationPropertiesResponsePtrOutput)
}

// Properties for the DscNodeConfiguration
type DscNodeConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DscNodeConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfigurationPropertiesResponse)(nil)).Elem()
}

func (o DscNodeConfigurationPropertiesResponseOutput) ToDscNodeConfigurationPropertiesResponseOutput() DscNodeConfigurationPropertiesResponseOutput {
	return o
}

func (o DscNodeConfigurationPropertiesResponseOutput) ToDscNodeConfigurationPropertiesResponseOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponseOutput {
	return o
}

func (o DscNodeConfigurationPropertiesResponseOutput) ToDscNodeConfigurationPropertiesResponsePtrOutput() DscNodeConfigurationPropertiesResponsePtrOutput {
	return o.ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DscNodeConfigurationPropertiesResponseOutput) ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *DscNodeConfigurationPropertiesResponse {
		return &v
	}).(DscNodeConfigurationPropertiesResponsePtrOutput)
}

// Gets or sets the configuration of the node.
func (o DscNodeConfigurationPropertiesResponseOutput) Configuration() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *DscConfigurationAssociationPropertyResponse {
		return v.Configuration
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// Gets or sets creation time.
func (o DscNodeConfigurationPropertiesResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// If a new build version of NodeConfiguration is required.
func (o DscNodeConfigurationPropertiesResponseOutput) IncrementNodeConfigurationBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *bool { return v.IncrementNodeConfigurationBuild }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o DscNodeConfigurationPropertiesResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Number of nodes with this node configuration assigned
func (o DscNodeConfigurationPropertiesResponseOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// Source of node configuration.
func (o DscNodeConfigurationPropertiesResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscNodeConfigurationPropertiesResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type DscNodeConfigurationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DscNodeConfigurationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscNodeConfigurationPropertiesResponse)(nil)).Elem()
}

func (o DscNodeConfigurationPropertiesResponsePtrOutput) ToDscNodeConfigurationPropertiesResponsePtrOutput() DscNodeConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o DscNodeConfigurationPropertiesResponsePtrOutput) ToDscNodeConfigurationPropertiesResponsePtrOutputWithContext(ctx context.Context) DscNodeConfigurationPropertiesResponsePtrOutput {
	return o
}

func (o DscNodeConfigurationPropertiesResponsePtrOutput) Elem() DscNodeConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) DscNodeConfigurationPropertiesResponse { return *v }).(DscNodeConfigurationPropertiesResponseOutput)
}

// Gets or sets the configuration of the node.
func (o DscNodeConfigurationPropertiesResponsePtrOutput) Configuration() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *DscConfigurationAssociationPropertyResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// Gets or sets creation time.
func (o DscNodeConfigurationPropertiesResponsePtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

// If a new build version of NodeConfiguration is required.
func (o DscNodeConfigurationPropertiesResponsePtrOutput) IncrementNodeConfigurationBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IncrementNodeConfigurationBuild
	}).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o DscNodeConfigurationPropertiesResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

// Number of nodes with this node configuration assigned
func (o DscNodeConfigurationPropertiesResponsePtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

// Source of node configuration.
func (o DscNodeConfigurationPropertiesResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfigurationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentSourceOutput{})
	pulumi.RegisterOutputType(ContentSourcePtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(DscNodeConfigurationTypeOutput{})
	pulumi.RegisterOutputType(DscNodeConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DscNodeConfigurationPropertiesResponsePtrOutput{})
}
