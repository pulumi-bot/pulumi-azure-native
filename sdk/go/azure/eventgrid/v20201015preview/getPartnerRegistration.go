// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPartnerRegistration(ctx *pulumi.Context, args *LookupPartnerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerRegistrationResult, error) {
	var rv LookupPartnerRegistrationResult
	err := ctx.Invoke("azure-nextgen:eventgrid/v20201015preview:getPartnerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerRegistrationArgs struct {
	// Name of the partner registration.
	PartnerRegistrationName string `pulumi:"partnerRegistrationName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about a partner registration.
type LookupPartnerRegistrationResult struct {
	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds []string `pulumi:"authorizedAzureSubscriptionIds"`
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri *string `pulumi:"customerServiceUri"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// URI of the logo.
	LogoUri *string `pulumi:"logoUri"`
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription *string `pulumi:"longDescription"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
	PartnerCustomerServiceExtension *string `pulumi:"partnerCustomerServiceExtension"`
	// The customer service number of the publisher. The expected phone format should start with a '+' sign
	// followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
	// length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
	// +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
	PartnerCustomerServiceNumber *string `pulumi:"partnerCustomerServiceNumber"`
	// Official name of the partner name. For example: "Contoso".
	PartnerName *string `pulumi:"partnerName"`
	// Short description of the partner resource type. The length of this description should not exceed 256 characters.
	PartnerResourceTypeDescription *string `pulumi:"partnerResourceTypeDescription"`
	// Display name of the partner resource type.
	PartnerResourceTypeDisplayName *string `pulumi:"partnerResourceTypeDisplayName"`
	// Name of the partner resource type.
	PartnerResourceTypeName *string `pulumi:"partnerResourceTypeName"`
	// Provisioning state of the partner registration.
	ProvisioningState string `pulumi:"provisioningState"`
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri *string `pulumi:"setupUri"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
	// Visibility state of the partner registration.
	VisibilityState *string `pulumi:"visibilityState"`
}
