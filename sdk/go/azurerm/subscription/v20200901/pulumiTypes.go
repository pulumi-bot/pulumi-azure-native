// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Put subscription properties.
type PutAliasRequestProperties struct {
	// Determines whether subscription is fieldLed, partnerLed or LegacyEA
	BillingScope *string `pulumi:"billingScope"`
	// The friendly name of the subscription.
	DisplayName *string `pulumi:"displayName"`
	// This parameter can be used to create alias for existing subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// The workload type of the subscription. It can be either Production or DevTest.
	Workload *string `pulumi:"workload"`
}

// PutAliasRequestPropertiesInput is an input type that accepts PutAliasRequestPropertiesArgs and PutAliasRequestPropertiesOutput values.
// You can construct a concrete instance of `PutAliasRequestPropertiesInput` via:
//
//          PutAliasRequestPropertiesArgs{...}
type PutAliasRequestPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput
	ToPutAliasRequestPropertiesOutputWithContext(context.Context) PutAliasRequestPropertiesOutput
}

// Put subscription properties.
type PutAliasRequestPropertiesArgs struct {
	// Determines whether subscription is fieldLed, partnerLed or LegacyEA
	BillingScope pulumi.StringPtrInput `pulumi:"billingScope"`
	// The friendly name of the subscription.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// This parameter can be used to create alias for existing subscription Id
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
	// The workload type of the subscription. It can be either Production or DevTest.
	Workload pulumi.StringPtrInput `pulumi:"workload"`
}

func (PutAliasRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return i.ToPutAliasRequestPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput)
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput).ToPutAliasRequestPropertiesPtrOutputWithContext(ctx)
}

// PutAliasRequestPropertiesPtrInput is an input type that accepts PutAliasRequestPropertiesArgs, PutAliasRequestPropertiesPtr and PutAliasRequestPropertiesPtrOutput values.
// You can construct a concrete instance of `PutAliasRequestPropertiesPtrInput` via:
//
//          PutAliasRequestPropertiesArgs{...}
//
//  or:
//
//          nil
type PutAliasRequestPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput
	ToPutAliasRequestPropertiesPtrOutputWithContext(context.Context) PutAliasRequestPropertiesPtrOutput
}

type putAliasRequestPropertiesPtrType PutAliasRequestPropertiesArgs

func PutAliasRequestPropertiesPtr(v *PutAliasRequestPropertiesArgs) PutAliasRequestPropertiesPtrInput {
	return (*putAliasRequestPropertiesPtrType)(v)
}

func (*putAliasRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesPtrOutput)
}

// Put subscription creation result properties.
type PutAliasResponsePropertiesResponse struct {
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Newly created subscription Id.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// Put subscription creation result properties.
type PutAliasResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PutAliasResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponseOutput() PutAliasResponsePropertiesResponseOutput {
	return o
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponseOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponseOutput {
	return o
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return o.ToPutAliasResponsePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PutAliasResponsePropertiesResponseOutput) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) *PutAliasResponsePropertiesResponse {
		return &v
	}).(PutAliasResponsePropertiesResponsePtrOutput)
}

// The provisioning state of the resource.
func (o PutAliasResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Newly created subscription Id.
func (o PutAliasResponsePropertiesResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v PutAliasResponsePropertiesResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

type PutAliasResponsePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PutAliasResponsePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o PutAliasResponsePropertiesResponsePtrOutput) ToPutAliasResponsePropertiesResponsePtrOutput() PutAliasResponsePropertiesResponsePtrOutput {
	return o
}

func (o PutAliasResponsePropertiesResponsePtrOutput) ToPutAliasResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) PutAliasResponsePropertiesResponsePtrOutput {
	return o
}

func (o PutAliasResponsePropertiesResponsePtrOutput) Elem() PutAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) PutAliasResponsePropertiesResponse { return *v }).(PutAliasResponsePropertiesResponseOutput)
}

// The provisioning state of the resource.
func (o PutAliasResponsePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// Newly created subscription Id.
func (o PutAliasResponsePropertiesResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PutAliasRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PutAliasResponsePropertiesResponsePtrOutput{})
}
