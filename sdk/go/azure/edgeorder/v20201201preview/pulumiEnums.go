// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Type of address.
type AddressType pulumi.String

const (
	// Address type not known.
	AddressTypeNone = AddressType("None")
	// Residential Address.
	AddressTypeResidential = AddressType("Residential")
	// Commercial Address.
	AddressTypeCommercial = AddressType("Commercial")
)

func (AddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AddressType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddressType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Name of the stage.
type NotificationStageName pulumi.String

const (
	// Notification at device prepared stage.
	NotificationStageNameDevicePrepared = NotificationStageName("DevicePrepared")
	// Notification at device dispatched stage.
	NotificationStageNameDispatched = NotificationStageName("Dispatched")
	// Notification at device delivered stage.
	NotificationStageNameDelivered = NotificationStageName("Delivered")
	// Notification at device picked up from user stage.
	NotificationStageNamePickedUp = NotificationStageName("PickedUp")
	// Notification at device received at Azure datacenter stage.
	NotificationStageNameAtAzureDC = NotificationStageName("AtAzureDC")
	// Notification at data copy started stage.
	NotificationStageNameDataCopy = NotificationStageName("DataCopy")
)

func (NotificationStageName) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e NotificationStageName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationStageName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationStageName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotificationStageName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Order type.
type OrderType pulumi.String

const (
	// Purchase Order.
	OrderTypePurchase = OrderType("Purchase")
	// Rental Order.
	OrderTypeRental = OrderType("Rental")
)

func (OrderType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e OrderType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrderType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrderType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrderType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Type of product filter.
type SupportedFilterTypes pulumi.String

const (
	// Ship to country
	SupportedFilterTypesShipToCountries = SupportedFilterTypes("ShipToCountries")
)

func (SupportedFilterTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e SupportedFilterTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedFilterTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Indicates Shipment Logistics type that the customer preferred.
type TransportShipmentTypes pulumi.String

const (
	// Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged = TransportShipmentTypes("CustomerManaged")
	// Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged = TransportShipmentTypes("MicrosoftManaged")
)

func (TransportShipmentTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TransportShipmentTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportShipmentTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportShipmentTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TransportShipmentTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
