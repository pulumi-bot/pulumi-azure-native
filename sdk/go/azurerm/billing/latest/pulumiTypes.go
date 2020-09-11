// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Details of the Azure plan.
type AzurePlanResponse struct {
	// The sku description.
	SkuDescription string `pulumi:"skuDescription"`
	// The sku id.
	SkuId *string `pulumi:"skuId"`
}

// Invoice section properties with create subscription permission.
type InvoiceSectionWithCreateSubPermissionResponse struct {
	// The name of the billing profile for the invoice section.
	BillingProfileDisplayName string `pulumi:"billingProfileDisplayName"`
	// The ID of the billing profile for the invoice section.
	BillingProfileId string `pulumi:"billingProfileId"`
	// The billing profile spending limit.
	BillingProfileSpendingLimit string `pulumi:"billingProfileSpendingLimit"`
	// The status of the billing profile.
	BillingProfileStatus string `pulumi:"billingProfileStatus"`
	// Reason for the specified billing profile status.
	BillingProfileStatusReasonCode string `pulumi:"billingProfileStatusReasonCode"`
	// The system generated unique identifier for a billing profile.
	BillingProfileSystemId string `pulumi:"billingProfileSystemId"`
	// Enabled azure plans for the associated billing profile.
	EnabledAzurePlans []AzurePlanResponse `pulumi:"enabledAzurePlans"`
	// The name of the invoice section.
	InvoiceSectionDisplayName string `pulumi:"invoiceSectionDisplayName"`
	// The ID of the invoice section.
	InvoiceSectionId string `pulumi:"invoiceSectionId"`
	// The system generated unique identifier for an invoice section.
	InvoiceSectionSystemId string `pulumi:"invoiceSectionSystemId"`
}

func init() {
	pulumi.RegisterOutputType(AzurePlanResponseOutput{})
	pulumi.RegisterOutputType(AzurePlanResponseArrayOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseArrayOutput{})
}
