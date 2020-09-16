// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The compliance status for the configuration profile assignment.
type ConfigurationProfileAssignmentComplianceResponse struct {
	// The state of compliance, which only appears in the response.
	UpdateStatus string `pulumi:"updateStatus"`
}

// ConfigurationProfileAssignmentComplianceResponseInput is an input type that accepts ConfigurationProfileAssignmentComplianceResponseArgs and ConfigurationProfileAssignmentComplianceResponseOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentComplianceResponseInput` via:
//
//          ConfigurationProfileAssignmentComplianceResponseArgs{...}
type ConfigurationProfileAssignmentComplianceResponseInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentComplianceResponseOutput() ConfigurationProfileAssignmentComplianceResponseOutput
	ToConfigurationProfileAssignmentComplianceResponseOutputWithContext(context.Context) ConfigurationProfileAssignmentComplianceResponseOutput
}

// The compliance status for the configuration profile assignment.
type ConfigurationProfileAssignmentComplianceResponseArgs struct {
	// The state of compliance, which only appears in the response.
	UpdateStatus pulumi.StringInput `pulumi:"updateStatus"`
}

func (ConfigurationProfileAssignmentComplianceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentComplianceResponseArgs) ToConfigurationProfileAssignmentComplianceResponseOutput() ConfigurationProfileAssignmentComplianceResponseOutput {
	return i.ToConfigurationProfileAssignmentComplianceResponseOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentComplianceResponseArgs) ToConfigurationProfileAssignmentComplianceResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentComplianceResponseOutput)
}

func (i ConfigurationProfileAssignmentComplianceResponseArgs) ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentComplianceResponseArgs) ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentComplianceResponseOutput).ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx)
}

// ConfigurationProfileAssignmentComplianceResponsePtrInput is an input type that accepts ConfigurationProfileAssignmentComplianceResponseArgs, ConfigurationProfileAssignmentComplianceResponsePtr and ConfigurationProfileAssignmentComplianceResponsePtrOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentComplianceResponsePtrInput` via:
//
//          ConfigurationProfileAssignmentComplianceResponseArgs{...}
//
//  or:
//
//          nil
type ConfigurationProfileAssignmentComplianceResponsePtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput
	ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput
}

type configurationProfileAssignmentComplianceResponsePtrType ConfigurationProfileAssignmentComplianceResponseArgs

func ConfigurationProfileAssignmentComplianceResponsePtr(v *ConfigurationProfileAssignmentComplianceResponseArgs) ConfigurationProfileAssignmentComplianceResponsePtrInput {
	return (*configurationProfileAssignmentComplianceResponsePtrType)(v)
}

func (*configurationProfileAssignmentComplianceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (i *configurationProfileAssignmentComplianceResponsePtrType) ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentComplianceResponsePtrType) ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentComplianceResponsePtrOutput)
}

// The compliance status for the configuration profile assignment.
type ConfigurationProfileAssignmentComplianceResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentComplianceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponseOutput() ConfigurationProfileAssignmentComplianceResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o.ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentComplianceResponseOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentComplianceResponse) *ConfigurationProfileAssignmentComplianceResponse {
		return &v
	}).(ConfigurationProfileAssignmentComplianceResponsePtrOutput)
}

// The state of compliance, which only appears in the response.
func (o ConfigurationProfileAssignmentComplianceResponseOutput) UpdateStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentComplianceResponse) string { return v.UpdateStatus }).(pulumi.StringOutput)
}

type ConfigurationProfileAssignmentComplianceResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentComplianceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentComplianceResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutput() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) ToConfigurationProfileAssignmentComplianceResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) Elem() ConfigurationProfileAssignmentComplianceResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentComplianceResponse) ConfigurationProfileAssignmentComplianceResponse {
		return *v
	}).(ConfigurationProfileAssignmentComplianceResponseOutput)
}

// The state of compliance, which only appears in the response.
func (o ConfigurationProfileAssignmentComplianceResponsePtrOutput) UpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentComplianceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdateStatus
	}).(pulumi.StringPtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentProperties struct {
	// The Automanage account ARM Resource URI
	AccountId *string `pulumi:"accountId"`
	// A value indicating configuration profile.
	ConfigurationProfile *string `pulumi:"configurationProfile"`
	// The configuration profile custom preferences ARM resource URI
	ConfigurationProfilePreferenceId *string `pulumi:"configurationProfilePreferenceId"`
	// The target VM resource URI
	TargetId *string `pulumi:"targetId"`
}

// ConfigurationProfileAssignmentPropertiesInput is an input type that accepts ConfigurationProfileAssignmentPropertiesArgs and ConfigurationProfileAssignmentPropertiesOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesInput` via:
//
//          ConfigurationProfileAssignmentPropertiesArgs{...}
type ConfigurationProfileAssignmentPropertiesInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput
	ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesOutput
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesArgs struct {
	// The Automanage account ARM Resource URI
	AccountId pulumi.StringPtrInput `pulumi:"accountId"`
	// A value indicating configuration profile.
	ConfigurationProfile pulumi.StringPtrInput `pulumi:"configurationProfile"`
	// The configuration profile custom preferences ARM resource URI
	ConfigurationProfilePreferenceId pulumi.StringPtrInput `pulumi:"configurationProfilePreferenceId"`
	// The target VM resource URI
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ConfigurationProfileAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return i.ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput).ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx)
}

// ConfigurationProfileAssignmentPropertiesPtrInput is an input type that accepts ConfigurationProfileAssignmentPropertiesArgs, ConfigurationProfileAssignmentPropertiesPtr and ConfigurationProfileAssignmentPropertiesPtrOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesPtrInput` via:
//
//          ConfigurationProfileAssignmentPropertiesArgs{...}
//
//  or:
//
//          nil
type ConfigurationProfileAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput
	ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput
}

type configurationProfileAssignmentPropertiesPtrType ConfigurationProfileAssignmentPropertiesArgs

func ConfigurationProfileAssignmentPropertiesPtr(v *ConfigurationProfileAssignmentPropertiesArgs) ConfigurationProfileAssignmentPropertiesPtrInput {
	return (*configurationProfileAssignmentPropertiesPtrType)(v)
}

func (*configurationProfileAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *ConfigurationProfileAssignmentProperties {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

// The Automanage account ARM Resource URI
func (o ConfigurationProfileAssignmentPropertiesOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// A value indicating configuration profile.
func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

// The configuration profile custom preferences ARM resource URI
func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfilePreferenceId }).(pulumi.StringPtrOutput)
}

// The target VM resource URI
func (o ConfigurationProfileAssignmentPropertiesOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) Elem() ConfigurationProfileAssignmentPropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) ConfigurationProfileAssignmentProperties { return *v }).(ConfigurationProfileAssignmentPropertiesOutput)
}

// The Automanage account ARM Resource URI
func (o ConfigurationProfileAssignmentPropertiesPtrOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccountId
	}).(pulumi.StringPtrOutput)
}

// A value indicating configuration profile.
func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

// The configuration profile custom preferences ARM resource URI
func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfilePreferenceId
	}).(pulumi.StringPtrOutput)
}

// The target VM resource URI
func (o ConfigurationProfileAssignmentPropertiesPtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesResponse struct {
	// The Automanage account ARM Resource URI
	AccountId *string `pulumi:"accountId"`
	// The configuration setting for the configuration profile.
	Compliance *ConfigurationProfileAssignmentComplianceResponse `pulumi:"compliance"`
	// A value indicating configuration profile.
	ConfigurationProfile *string `pulumi:"configurationProfile"`
	// The configuration profile custom preferences ARM resource URI
	ConfigurationProfilePreferenceId *string `pulumi:"configurationProfilePreferenceId"`
	// The state of onboarding, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The target VM resource URI
	TargetId *string `pulumi:"targetId"`
}

// ConfigurationProfileAssignmentPropertiesResponseInput is an input type that accepts ConfigurationProfileAssignmentPropertiesResponseArgs and ConfigurationProfileAssignmentPropertiesResponseOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesResponseInput` via:
//
//          ConfigurationProfileAssignmentPropertiesResponseArgs{...}
type ConfigurationProfileAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput
	ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesResponseArgs struct {
	// The Automanage account ARM Resource URI
	AccountId pulumi.StringPtrInput `pulumi:"accountId"`
	// The configuration setting for the configuration profile.
	Compliance ConfigurationProfileAssignmentComplianceResponsePtrInput `pulumi:"compliance"`
	// A value indicating configuration profile.
	ConfigurationProfile pulumi.StringPtrInput `pulumi:"configurationProfile"`
	// The configuration profile custom preferences ARM resource URI
	ConfigurationProfilePreferenceId pulumi.StringPtrInput `pulumi:"configurationProfilePreferenceId"`
	// The state of onboarding, which only appears in the response.
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
	// The target VM resource URI
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ConfigurationProfileAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesResponseArgs) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponseOutput).ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}

// ConfigurationProfileAssignmentPropertiesResponsePtrInput is an input type that accepts ConfigurationProfileAssignmentPropertiesResponseArgs, ConfigurationProfileAssignmentPropertiesResponsePtr and ConfigurationProfileAssignmentPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesResponsePtrInput` via:
//
//          ConfigurationProfileAssignmentPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type ConfigurationProfileAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput
	ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput
}

type configurationProfileAssignmentPropertiesResponsePtrType ConfigurationProfileAssignmentPropertiesResponseArgs

func ConfigurationProfileAssignmentPropertiesResponsePtr(v *ConfigurationProfileAssignmentPropertiesResponseArgs) ConfigurationProfileAssignmentPropertiesResponsePtrInput {
	return (*configurationProfileAssignmentPropertiesResponsePtrType)(v)
}

func (*configurationProfileAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesResponsePtrType) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesResponsePtrType) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesResponsePtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *ConfigurationProfileAssignmentPropertiesResponse {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesResponsePtrOutput)
}

// The Automanage account ARM Resource URI
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The configuration setting for the configuration profile.
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) Compliance() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *ConfigurationProfileAssignmentComplianceResponse {
		return v.Compliance
	}).(ConfigurationProfileAssignmentComplianceResponsePtrOutput)
}

// A value indicating configuration profile.
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

// The configuration profile custom preferences ARM resource URI
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string {
		return v.ConfigurationProfilePreferenceId
	}).(pulumi.StringPtrOutput)
}

// The state of onboarding, which only appears in the response.
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The target VM resource URI
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutput() ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ToConfigurationProfileAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) Elem() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) ConfigurationProfileAssignmentPropertiesResponse {
		return *v
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

// The Automanage account ARM Resource URI
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountId
	}).(pulumi.StringPtrOutput)
}

// The configuration setting for the configuration profile.
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) Compliance() ConfigurationProfileAssignmentComplianceResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *ConfigurationProfileAssignmentComplianceResponse {
		if v == nil {
			return nil
		}
		return v.Compliance
	}).(ConfigurationProfileAssignmentComplianceResponsePtrOutput)
}

// A value indicating configuration profile.
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

// The configuration profile custom preferences ARM resource URI
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ConfigurationProfilePreferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfilePreferenceId
	}).(pulumi.StringPtrOutput)
}

// The state of onboarding, which only appears in the response.
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// The target VM resource URI
func (o ConfigurationProfileAssignmentPropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentComplianceResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentComplianceResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponsePtrOutput{})
}
