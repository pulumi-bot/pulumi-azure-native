// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A flow log resource.
type FlowLog struct {
	pulumi.CustomResourceState

	// Flag to enable/disable flow logging.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Parameters that define the configuration of traffic analytics.
	FlowAnalyticsConfiguration TrafficAnalyticsPropertiesResponsePtrOutput `pulumi:"flowAnalyticsConfiguration"`
	// Parameters that define the flow log format.
	Format FlowLogFormatParametersResponsePtrOutput `pulumi:"format"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the flow log.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Parameters that define the retention policy for flow log.
	RetentionPolicy RetentionPolicyParametersResponsePtrOutput `pulumi:"retentionPolicy"`
	// ID of the storage account which is used to store the flow log.
	StorageId pulumi.StringOutput `pulumi:"storageId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Guid of network security group to which flow log will be applied.
	TargetResourceGuid pulumi.StringOutput `pulumi:"targetResourceGuid"`
	// ID of network security group to which flow log will be applied.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	if args == nil || args.FlowLogName == nil {
		return nil, errors.New("missing required argument 'FlowLogName'")
	}
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageId == nil {
		return nil, errors.New("missing required argument 'StorageId'")
	}
	if args == nil || args.TargetResourceId == nil {
		return nil, errors.New("missing required argument 'TargetResourceId'")
	}
	if args == nil {
		args = &FlowLogArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:FlowLog"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:FlowLog"),
		},
	})
	opts = append(opts, aliases)
	var resource FlowLog
	err := ctx.RegisterResource("azure-nextgen:network/v20200301:FlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowLogState, opts ...pulumi.ResourceOption) (*FlowLog, error) {
	var resource FlowLog
	err := ctx.ReadResource("azure-nextgen:network/v20200301:FlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowLog resources.
type flowLogState struct {
	// Flag to enable/disable flow logging.
	Enabled *bool `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Parameters that define the configuration of traffic analytics.
	FlowAnalyticsConfiguration *TrafficAnalyticsPropertiesResponse `pulumi:"flowAnalyticsConfiguration"`
	// Parameters that define the flow log format.
	Format *FlowLogFormatParametersResponse `pulumi:"format"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the flow log.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Parameters that define the retention policy for flow log.
	RetentionPolicy *RetentionPolicyParametersResponse `pulumi:"retentionPolicy"`
	// ID of the storage account which is used to store the flow log.
	StorageId *string `pulumi:"storageId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Guid of network security group to which flow log will be applied.
	TargetResourceGuid *string `pulumi:"targetResourceGuid"`
	// ID of network security group to which flow log will be applied.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type FlowLogState struct {
	// Flag to enable/disable flow logging.
	Enabled pulumi.BoolPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Parameters that define the configuration of traffic analytics.
	FlowAnalyticsConfiguration TrafficAnalyticsPropertiesResponsePtrInput
	// Parameters that define the flow log format.
	Format FlowLogFormatParametersResponsePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the flow log.
	ProvisioningState pulumi.StringPtrInput
	// Parameters that define the retention policy for flow log.
	RetentionPolicy RetentionPolicyParametersResponsePtrInput
	// ID of the storage account which is used to store the flow log.
	StorageId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Guid of network security group to which flow log will be applied.
	TargetResourceGuid pulumi.StringPtrInput
	// ID of network security group to which flow log will be applied.
	TargetResourceId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (FlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogState)(nil)).Elem()
}

type flowLogArgs struct {
	// Flag to enable/disable flow logging.
	Enabled *bool `pulumi:"enabled"`
	// Parameters that define the configuration of traffic analytics.
	FlowAnalyticsConfiguration *TrafficAnalyticsProperties `pulumi:"flowAnalyticsConfiguration"`
	// The name of the flow log.
	FlowLogName string `pulumi:"flowLogName"`
	// Parameters that define the flow log format.
	Format *FlowLogFormatParameters `pulumi:"format"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameters that define the retention policy for flow log.
	RetentionPolicy *RetentionPolicyParameters `pulumi:"retentionPolicy"`
	// ID of the storage account which is used to store the flow log.
	StorageId string `pulumi:"storageId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// ID of network security group to which flow log will be applied.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// Flag to enable/disable flow logging.
	Enabled pulumi.BoolPtrInput
	// Parameters that define the configuration of traffic analytics.
	FlowAnalyticsConfiguration TrafficAnalyticsPropertiesPtrInput
	// The name of the flow log.
	FlowLogName pulumi.StringInput
	// Parameters that define the flow log format.
	Format FlowLogFormatParametersPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Parameters that define the retention policy for flow log.
	RetentionPolicy RetentionPolicyParametersPtrInput
	// ID of the storage account which is used to store the flow log.
	StorageId pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// ID of network security group to which flow log will be applied.
	TargetResourceId pulumi.StringInput
}

func (FlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowLogArgs)(nil)).Elem()
}

type FlowLogInput interface {
	pulumi.Input

	ToFlowLogOutput() FlowLogOutput
	ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput
}

func (FlowLog) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLog)(nil)).Elem()
}

func (i FlowLog) ToFlowLogOutput() FlowLogOutput {
	return i.ToFlowLogOutputWithContext(context.Background())
}

func (i FlowLog) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowLogOutput)
}

type FlowLogOutput struct {
	*pulumi.OutputState
}

func (FlowLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLogOutput)(nil)).Elem()
}

func (o FlowLogOutput) ToFlowLogOutput() FlowLogOutput {
	return o
}

func (o FlowLogOutput) ToFlowLogOutputWithContext(ctx context.Context) FlowLogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FlowLogOutput{})
}
