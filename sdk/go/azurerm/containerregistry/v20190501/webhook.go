// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An object that represents a webhook for a container registry.
type Webhook struct {
	pulumi.CustomResourceState

	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the webhook.
	Properties WebhookPropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil || args.Actions == nil {
		return nil, errors.New("missing required argument 'Actions'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.RegistryName == nil {
		return nil, errors.New("missing required argument 'RegistryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceUri == nil {
		return nil, errors.New("missing required argument 'ServiceUri'")
	}
	if args == nil {
		args = &WebhookArgs{}
	}
	var resource Webhook
	err := ctx.RegisterResource("azurerm:containerregistry/v20190501:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azurerm:containerregistry/v20190501:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the webhook.
	Properties *WebhookPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type WebhookState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the webhook.
	Properties WebhookPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions []string `pulumi:"actions"`
	// Custom headers that will be added to the webhook notifications.
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	// The location of the webhook. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the webhook.
	Name string `pulumi:"name"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope *string `pulumi:"scope"`
	// The service URI for the webhook to post notifications.
	ServiceUri string `pulumi:"serviceUri"`
	// The status of the webhook at the time the operation was called.
	Status *string `pulumi:"status"`
	// The tags for the webhook.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The list of actions that trigger the webhook to post notifications.
	Actions pulumi.StringArrayInput
	// Custom headers that will be added to the webhook notifications.
	CustomHeaders pulumi.StringMapInput
	// The location of the webhook. This cannot be changed after the resource is created.
	Location pulumi.StringInput
	// The name of the webhook.
	Name pulumi.StringInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
	Scope pulumi.StringPtrInput
	// The service URI for the webhook to post notifications.
	ServiceUri pulumi.StringInput
	// The status of the webhook at the time the operation was called.
	Status pulumi.StringPtrInput
	// The tags for the webhook.
	Tags pulumi.StringMapInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}
