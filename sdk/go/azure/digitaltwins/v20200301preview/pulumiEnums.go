// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The name of the SKU.
type DigitalTwinsSku pulumi.String

const (
	DigitalTwinsSkuF1 = DigitalTwinsSku("F1")
)

func (DigitalTwinsSku) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DigitalTwinsSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DigitalTwinsSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DigitalTwinsSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DigitalTwinsSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of Digital Twins endpoint
type EndpointType pulumi.String

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

func (EndpointType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EndpointType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
