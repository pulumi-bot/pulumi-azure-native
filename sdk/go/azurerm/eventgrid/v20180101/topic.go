// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// EventGrid Topic
type Topic struct {
	pulumi.CustomResourceState

	// Location of the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the topic
	Properties TopicPropertiesResponseOutput `pulumi:"properties"`
	// Tags of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TopicArgs{}
	}
	var resource Topic
	err := ctx.RegisterResource("azurerm:eventgrid/v20180101:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azurerm:eventgrid/v20180101:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Location of the resource
	Location *string `pulumi:"location"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Properties of the topic
	Properties *TopicPropertiesResponse `pulumi:"properties"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type TopicState struct {
	// Location of the resource
	Location pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Properties of the topic
	Properties TopicPropertiesResponsePtrInput
	// Tags of the resource
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Location of the resource
	Location string `pulumi:"location"`
	// Name of the topic
	Name string `pulumi:"name"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Location of the resource
	Location pulumi.StringInput
	// Name of the topic
	Name pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource
	Tags pulumi.StringMapInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}
