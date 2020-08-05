// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about the connection monitor.
type ConnectionMonitor struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Connection monitor location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the connection monitor.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the connection monitor result.
	Properties ConnectionMonitorResultPropertiesResponseOutput `pulumi:"properties"`
	// Connection monitor tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Connection monitor type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionMonitor registers a new resource with the given unique name, arguments, and options.
func NewConnectionMonitor(ctx *pulumi.Context,
	name string, args *ConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
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
		args = &ConnectionMonitorArgs{}
	}
	var resource ConnectionMonitor
	err := ctx.RegisterResource("azurerm:network/v20191101:ConnectionMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionMonitor gets an existing ConnectionMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionMonitorState, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
	var resource ConnectionMonitor
	err := ctx.ReadResource("azurerm:network/v20191101:ConnectionMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionMonitor resources.
type connectionMonitorState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Connection monitor location.
	Location *string `pulumi:"location"`
	// Name of the connection monitor.
	Name *string `pulumi:"name"`
	// Properties of the connection monitor result.
	Properties *ConnectionMonitorResultPropertiesResponse `pulumi:"properties"`
	// Connection monitor tags.
	Tags map[string]string `pulumi:"tags"`
	// Connection monitor type.
	Type *string `pulumi:"type"`
}

type ConnectionMonitorState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Connection monitor location.
	Location pulumi.StringPtrInput
	// Name of the connection monitor.
	Name pulumi.StringPtrInput
	// Properties of the connection monitor result.
	Properties ConnectionMonitorResultPropertiesResponsePtrInput
	// Connection monitor tags.
	Tags pulumi.StringMapInput
	// Connection monitor type.
	Type pulumi.StringPtrInput
}

func (ConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorState)(nil)).Elem()
}

type connectionMonitorArgs struct {
	// Determines if the connection monitor will start automatically once created.
	AutoStart *bool `pulumi:"autoStart"`
	// Describes the destination of connection monitor.
	Destination *ConnectionMonitorDestination `pulumi:"destination"`
	// List of connection monitor endpoints.
	Endpoints []ConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Connection monitor location.
	Location *string `pulumi:"location"`
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds *int `pulumi:"monitoringIntervalInSeconds"`
	// The name of the connection monitor.
	Name string `pulumi:"name"`
	// The name of the Network Watcher resource.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// Optional notes to be associated with the connection monitor.
	Notes *string `pulumi:"notes"`
	// List of connection monitor outputs.
	Outputs []ConnectionMonitorOutput `pulumi:"outputs"`
	// The name of the resource group containing Network Watcher.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the source of connection monitor.
	Source *ConnectionMonitorSource `pulumi:"source"`
	// Connection monitor tags.
	Tags map[string]string `pulumi:"tags"`
	// List of connection monitor test configurations.
	TestConfigurations []ConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// List of connection monitor test groups.
	TestGroups []ConnectionMonitorTestGroup `pulumi:"testGroups"`
}

// The set of arguments for constructing a ConnectionMonitor resource.
type ConnectionMonitorArgs struct {
	// Determines if the connection monitor will start automatically once created.
	AutoStart pulumi.BoolPtrInput
	// Describes the destination of connection monitor.
	Destination ConnectionMonitorDestinationPtrInput
	// List of connection monitor endpoints.
	Endpoints ConnectionMonitorEndpointArrayInput
	// Connection monitor location.
	Location pulumi.StringPtrInput
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds pulumi.IntPtrInput
	// The name of the connection monitor.
	Name pulumi.StringInput
	// The name of the Network Watcher resource.
	NetworkWatcherName pulumi.StringInput
	// Optional notes to be associated with the connection monitor.
	Notes pulumi.StringPtrInput
	// List of connection monitor outputs.
	Outputs ConnectionMonitorOutputArrayInput
	// The name of the resource group containing Network Watcher.
	ResourceGroupName pulumi.StringInput
	// Describes the source of connection monitor.
	Source ConnectionMonitorSourcePtrInput
	// Connection monitor tags.
	Tags pulumi.StringMapInput
	// List of connection monitor test configurations.
	TestConfigurations ConnectionMonitorTestConfigurationArrayInput
	// List of connection monitor test groups.
	TestGroups ConnectionMonitorTestGroupArrayInput
}

func (ConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorArgs)(nil)).Elem()
}
