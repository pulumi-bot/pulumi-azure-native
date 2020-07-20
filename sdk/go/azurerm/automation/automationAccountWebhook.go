// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the webhook type.
type AutomationAccountWebhook struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the webhook properties.
	Properties WebhookPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationAccountWebhook registers a new resource with the given unique name, arguments, and options.
func NewAutomationAccountWebhook(ctx *pulumi.Context,
	name string, args *AutomationAccountWebhookArgs, opts ...pulumi.ResourceOption) (*AutomationAccountWebhook, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AutomationAccountWebhookArgs{}
	}
	var resource AutomationAccountWebhook
	err := ctx.RegisterResource("azurerm:automation:AutomationAccountWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationAccountWebhook gets an existing AutomationAccountWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationAccountWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountWebhookState, opts ...pulumi.ResourceOption) (*AutomationAccountWebhook, error) {
	var resource AutomationAccountWebhook
	err := ctx.ReadResource("azurerm:automation:AutomationAccountWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationAccountWebhook resources.
type automationAccountWebhookState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the webhook properties.
	Properties *WebhookPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type AutomationAccountWebhookState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the webhook properties.
	Properties WebhookPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (AutomationAccountWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountWebhookState)(nil)).Elem()
}

type automationAccountWebhookArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The webhook name.
	Name string `pulumi:"name"`
	// Gets or sets the properties of the webhook.
	Properties WebhookCreateOrUpdateProperties `pulumi:"properties"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AutomationAccountWebhook resource.
type AutomationAccountWebhookArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// The webhook name.
	Name pulumi.StringInput
	// Gets or sets the properties of the webhook.
	Properties WebhookCreateOrUpdatePropertiesInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (AutomationAccountWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountWebhookArgs)(nil)).Elem()
}
