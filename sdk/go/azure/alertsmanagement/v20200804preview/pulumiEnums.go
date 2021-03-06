// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200804preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// specifies the type of the alert criterion.
type HealthAlertsNamespace pulumi.String

const (
	HealthAlertsNamespaceGuestVmHealth = HealthAlertsNamespace("GuestVmHealth")
)

func (HealthAlertsNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e HealthAlertsNamespace) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthAlertsNamespace) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthAlertsNamespace) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HealthAlertsNamespace) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Health state name
type HealthStateName pulumi.String

const (
	HealthStateNameWarning  = HealthStateName("Warning")
	HealthStateNameCritical = HealthStateName("Critical")
)

func (HealthStateName) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e HealthStateName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthStateName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthStateName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HealthStateName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
