// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a NotificationHub Resource.
type NotificationHub struct {
	pulumi.CustomResourceState

	// Gets or sets datacenter location of the NotificationHub.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets name of the NotificationHub.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets or sets properties of the NotificationHub.
	Properties NotificationHubPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets tags of the NotificationHub.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets resource type of the NotificationHub.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewNotificationHub registers a new resource with the given unique name, arguments, and options.
func NewNotificationHub(ctx *pulumi.Context,
	name string, args *NotificationHubArgs, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NotificationHubArgs{}
	}
	var resource NotificationHub
	err := ctx.RegisterResource("azurerm:notificationhubs/v20140901:NotificationHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationHub gets an existing NotificationHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubState, opts ...pulumi.ResourceOption) (*NotificationHub, error) {
	var resource NotificationHub
	err := ctx.ReadResource("azurerm:notificationhubs/v20140901:NotificationHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationHub resources.
type notificationHubState struct {
	// Gets or sets datacenter location of the NotificationHub.
	Location *string `pulumi:"location"`
	// Gets or sets name of the NotificationHub.
	Name *string `pulumi:"name"`
	// Gets or sets properties of the NotificationHub.
	Properties *NotificationHubPropertiesResponse `pulumi:"properties"`
	// Gets or sets tags of the NotificationHub.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets resource type of the NotificationHub.
	Type *string `pulumi:"type"`
}

type NotificationHubState struct {
	// Gets or sets datacenter location of the NotificationHub.
	Location pulumi.StringPtrInput
	// Gets or sets name of the NotificationHub.
	Name pulumi.StringPtrInput
	// Gets or sets properties of the NotificationHub.
	Properties NotificationHubPropertiesResponsePtrInput
	// Gets or sets tags of the NotificationHub.
	Tags pulumi.StringMapInput
	// Gets or sets resource type of the NotificationHub.
	Type pulumi.StringPtrInput
}

func (NotificationHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubState)(nil)).Elem()
}

type notificationHubArgs struct {
	// Gets or sets NotificationHub data center location.
	Location string `pulumi:"location"`
	// The notification hub name.
	Name string `pulumi:"name"`
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// Gets or sets properties of the NotificationHub.
	Properties NotificationHubProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets NotificationHub tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NotificationHub resource.
type NotificationHubArgs struct {
	// Gets or sets NotificationHub data center location.
	Location pulumi.StringInput
	// The notification hub name.
	Name pulumi.StringInput
	// The namespace name.
	NamespaceName pulumi.StringInput
	// Gets or sets properties of the NotificationHub.
	Properties NotificationHubPropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets NotificationHub tags.
	Tags pulumi.StringMapInput
}

func (NotificationHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubArgs)(nil)).Elem()
}