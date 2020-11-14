// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about a partner registration.
type PartnerRegistration struct {
	pulumi.CustomResourceState

	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds pulumi.StringArrayOutput `pulumi:"authorizedAzureSubscriptionIds"`
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri pulumi.StringPtrOutput `pulumi:"customerServiceUri"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// URI of the logo.
	LogoUri pulumi.StringPtrOutput `pulumi:"logoUri"`
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription pulumi.StringPtrOutput `pulumi:"longDescription"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
	PartnerCustomerServiceExtension pulumi.StringPtrOutput `pulumi:"partnerCustomerServiceExtension"`
	// The customer service number of the publisher. The expected phone format should start with a '+' sign
	// followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
	// length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
	// +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
	PartnerCustomerServiceNumber pulumi.StringPtrOutput `pulumi:"partnerCustomerServiceNumber"`
	// Official name of the partner name. For example: "Contoso".
	PartnerName pulumi.StringPtrOutput `pulumi:"partnerName"`
	// Short description of the partner resource type. The length of this description should not exceed 256 characters.
	PartnerResourceTypeDescription pulumi.StringPtrOutput `pulumi:"partnerResourceTypeDescription"`
	// Display name of the partner resource type.
	PartnerResourceTypeDisplayName pulumi.StringPtrOutput `pulumi:"partnerResourceTypeDisplayName"`
	// Name of the partner resource type.
	PartnerResourceTypeName pulumi.StringPtrOutput `pulumi:"partnerResourceTypeName"`
	// Provisioning state of the partner registration.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri pulumi.StringPtrOutput `pulumi:"setupUri"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
	// Visibility state of the partner registration.
	VisibilityState pulumi.StringPtrOutput `pulumi:"visibilityState"`
}

// NewPartnerRegistration registers a new resource with the given unique name, arguments, and options.
func NewPartnerRegistration(ctx *pulumi.Context,
	name string, args *PartnerRegistrationArgs, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PartnerRegistrationName == nil {
		return nil, errors.New("missing required argument 'PartnerRegistrationName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PartnerRegistrationArgs{}
	}
	var resource PartnerRegistration
	err := ctx.RegisterResource("azure-nextgen:eventgrid/v20200401preview:PartnerRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerRegistration gets an existing PartnerRegistration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerRegistrationState, opts ...pulumi.ResourceOption) (*PartnerRegistration, error) {
	var resource PartnerRegistration
	err := ctx.ReadResource("azure-nextgen:eventgrid/v20200401preview:PartnerRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerRegistration resources.
type partnerRegistrationState struct {
	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds []string `pulumi:"authorizedAzureSubscriptionIds"`
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri *string `pulumi:"customerServiceUri"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// URI of the logo.
	LogoUri *string `pulumi:"logoUri"`
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription *string `pulumi:"longDescription"`
	// Name of the resource
	Name *string `pulumi:"name"`
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
	ProvisioningState *string `pulumi:"provisioningState"`
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri *string `pulumi:"setupUri"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
	// Visibility state of the partner registration.
	VisibilityState *string `pulumi:"visibilityState"`
}

type PartnerRegistrationState struct {
	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds pulumi.StringArrayInput
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// URI of the logo.
	LogoUri pulumi.StringPtrInput
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
	PartnerCustomerServiceExtension pulumi.StringPtrInput
	// The customer service number of the publisher. The expected phone format should start with a '+' sign
	// followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
	// length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
	// +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
	PartnerCustomerServiceNumber pulumi.StringPtrInput
	// Official name of the partner name. For example: "Contoso".
	PartnerName pulumi.StringPtrInput
	// Short description of the partner resource type. The length of this description should not exceed 256 characters.
	PartnerResourceTypeDescription pulumi.StringPtrInput
	// Display name of the partner resource type.
	PartnerResourceTypeDisplayName pulumi.StringPtrInput
	// Name of the partner resource type.
	PartnerResourceTypeName pulumi.StringPtrInput
	// Provisioning state of the partner registration.
	ProvisioningState pulumi.StringPtrInput
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri pulumi.StringPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
	// Visibility state of the partner registration.
	VisibilityState pulumi.StringPtrInput
}

func (PartnerRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationState)(nil)).Elem()
}

type partnerRegistrationArgs struct {
	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds []string `pulumi:"authorizedAzureSubscriptionIds"`
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri *string `pulumi:"customerServiceUri"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// URI of the logo.
	LogoUri *string `pulumi:"logoUri"`
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription *string `pulumi:"longDescription"`
	// The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
	PartnerCustomerServiceExtension *string `pulumi:"partnerCustomerServiceExtension"`
	// The customer service number of the publisher. The expected phone format should start with a '+' sign
	// followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
	// length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
	// +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
	PartnerCustomerServiceNumber *string `pulumi:"partnerCustomerServiceNumber"`
	// Official name of the partner name. For example: "Contoso".
	PartnerName *string `pulumi:"partnerName"`
	// Name of the partner registration.
	PartnerRegistrationName string `pulumi:"partnerRegistrationName"`
	// Short description of the partner resource type. The length of this description should not exceed 256 characters.
	PartnerResourceTypeDescription *string `pulumi:"partnerResourceTypeDescription"`
	// Display name of the partner resource type.
	PartnerResourceTypeDisplayName *string `pulumi:"partnerResourceTypeDisplayName"`
	// Name of the partner resource type.
	PartnerResourceTypeName *string `pulumi:"partnerResourceTypeName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri *string `pulumi:"setupUri"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Visibility state of the partner registration.
	VisibilityState *string `pulumi:"visibilityState"`
}

// The set of arguments for constructing a PartnerRegistration resource.
type PartnerRegistrationArgs struct {
	// List of Azure subscription Ids that are authorized to create a partner namespace
	// associated with this partner registration. This is an optional property. Creating
	// partner namespaces is always permitted under the same Azure subscription as the one used
	// for creating the partner registration.
	AuthorizedAzureSubscriptionIds pulumi.StringArrayInput
	// The extension of the customer service URI of the publisher.
	CustomerServiceUri pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringInput
	// URI of the logo.
	LogoUri pulumi.StringPtrInput
	// Long description for the custom scenarios and integration to be displayed in the portal if needed.
	// Length of this description should not exceed 2048 characters.
	LongDescription pulumi.StringPtrInput
	// The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
	PartnerCustomerServiceExtension pulumi.StringPtrInput
	// The customer service number of the publisher. The expected phone format should start with a '+' sign
	// followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
	// length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
	// +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
	PartnerCustomerServiceNumber pulumi.StringPtrInput
	// Official name of the partner name. For example: "Contoso".
	PartnerName pulumi.StringPtrInput
	// Name of the partner registration.
	PartnerRegistrationName pulumi.StringInput
	// Short description of the partner resource type. The length of this description should not exceed 256 characters.
	PartnerResourceTypeDescription pulumi.StringPtrInput
	// Display name of the partner resource type.
	PartnerResourceTypeDisplayName pulumi.StringPtrInput
	// Name of the partner resource type.
	PartnerResourceTypeName pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// URI of the partner website that can be used by Azure customers to setup Event Grid
	// integration on an event source.
	SetupUri pulumi.StringPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Visibility state of the partner registration.
	VisibilityState pulumi.StringPtrInput
}

func (PartnerRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerRegistrationArgs)(nil)).Elem()
}

type PartnerRegistrationInput interface {
	pulumi.Input

	ToPartnerRegistrationOutput() PartnerRegistrationOutput
	ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput
}

func (PartnerRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistration)(nil)).Elem()
}

func (i PartnerRegistration) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return i.ToPartnerRegistrationOutputWithContext(context.Background())
}

func (i PartnerRegistration) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerRegistrationOutput)
}

type PartnerRegistrationOutput struct {
	*pulumi.OutputState
}

func (PartnerRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistrationOutput)(nil)).Elem()
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutput() PartnerRegistrationOutput {
	return o
}

func (o PartnerRegistrationOutput) ToPartnerRegistrationOutputWithContext(ctx context.Context) PartnerRegistrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PartnerRegistrationOutput{})
}
