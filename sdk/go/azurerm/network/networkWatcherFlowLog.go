// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A flow log resource.
type NetworkWatcherFlowLog struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the flow log.
	Properties FlowLogPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkWatcherFlowLog registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, args *NetworkWatcherFlowLogArgs, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkWatcherFlowLogArgs{}
	}
	var resource NetworkWatcherFlowLog
	err := ctx.RegisterResource("azurerm:network:NetworkWatcherFlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkWatcherFlowLog gets an existing NetworkWatcherFlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherFlowLogState, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	var resource NetworkWatcherFlowLog
	err := ctx.ReadResource("azurerm:network:NetworkWatcherFlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkWatcherFlowLog resources.
type networkWatcherFlowLogState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the flow log.
	Properties *FlowLogPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NetworkWatcherFlowLogState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the flow log.
	Properties FlowLogPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkWatcherFlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogState)(nil)).Elem()
}

type networkWatcherFlowLogArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the flow log.
	Name string `pulumi:"name"`
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// Properties of the flow log.
	Properties *FlowLogPropertiesFormat `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkWatcherFlowLog resource.
type NetworkWatcherFlowLogArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the flow log.
	Name pulumi.StringInput
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput
	// Properties of the flow log.
	Properties FlowLogPropertiesFormatPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkWatcherFlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogArgs)(nil)).Elem()
}