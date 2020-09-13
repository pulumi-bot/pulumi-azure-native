// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170605preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Registered Server resource.
//
// ## Example Usage
// ### RegisteredServers_Create
//
// ```go
// package main
//
// import (
// 	storagesync "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/storagesync/v20170605preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagesync.NewRegisteredServer(ctx, "registeredServer", &storagesync.RegisteredServerArgs{
// 			AgentVersion:               pulumi.String("1.0.277.0"),
// 			LastHeartBeat:              pulumi.String("\"2017-08-08T18:29:06.470652Z\""),
// 			ProvisioningState:          pulumi.String("Created"),
// 			ResourceGroupName:          pulumi.String("SampleResourceGroup_1"),
// 			ServerCertificate:          pulumi.String("\"MIIDFjCCAf6gAwIBAgIQQS+DS8uhc4VNzUkTw7wbRjANBgkqhkiG9w0BAQ0FADAzMTEwLwYDVQQDEyhhbmt1c2hiLXByb2QzLnJlZG1vbmQuY29ycC5taWNyb3NvZnQuY29tMB4XDTE3MDgwMzE3MDQyNFoXDTE4MDgwNDE3MDQyNFowMzExMC8GA1UEAxMoYW5rdXNoYi1wcm9kMy5yZWRtb25kLmNvcnAubWljcm9zb2Z0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALDRvV4gmsIy6jGDPiHsXmvgVP749NNP7DopdlbHaNhjFmYINHl0uWylyaZmgJrROt2mnxN/zEyJtGnqYHlzUr4xvGq/qV5pqgdB9tag/sw9i22gfe9PRZ0FmSOZnXMbLYgLiDFqLtut5gHcOuWMj03YnkfoBEKlFBxWbagvW2yxz/Sxi9OVSJOKCaXra0RpcIHrO/KFl6ho2eE1/7Ykmfa8hZvSdoPd5gHdLiQcMB/pxq+mWp1fI6c8vFZoDu7Atn+NXTzYPKUxKzaisF12TsaKpohUsJpbB3Wocb0F5frn614D2pg14ERB5otjAMWw1m65csQWPI6dP8KIYe0+QPkCAwEAAaMmMCQwIgYDVR0lAQH/BBgwFgYIKwYBBQUHAwIGCisGAQQBgjcKAwwwDQYJKoZIhvcNAQENBQADggEBAA4RhVIBkw34M1RwakJgHvtjsOFxF1tVQA941NtLokx1l2Z8+GFQkcG4xpZSt+UN6wLerdCbnNhtkCErWUDeaT0jxk4g71Ofex7iM04crT4iHJr8mi96/XnhnkTUs+GDk12VgdeeNEczMZz+8Mxw9dJ5NCnYgTwO0SzGlclRsDvjzkLo8rh2ZG6n/jKrEyNXXo+hOqhupij0QbRP2Tvexdfw201kgN1jdZify8XzJ8Oi0bTS0KpJf2pNPOlooK2bjMUei9ANtEdXwwfVZGWvVh6tJjdv6k14wWWJ1L7zhA1IIVb1J+sQUzJji5iX0DrezjTz1Fg+gAzITaA/WsuujlM=\""),
// 			ServerId:                   pulumi.String("\"080d4133-bdb5-40a0-96a0-71a6057bfe9a\""),
// 			ServerManagementtErrorCode: pulumi.Int(0),
// 			ServerOSVersion:            pulumi.String("10.0.14393.0"),
// 			ServerRole:                 pulumi.String("Standalone"),
// 			StorageSyncServiceName:     pulumi.String("SampleStorageSyncService_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RegisteredServer struct {
	pulumi.CustomResourceState

	// Registered Server Agent Version
	AgentVersion pulumi.StringPtrOutput `pulumi:"agentVersion"`
	// Registered Server clusterId
	ClusterId pulumi.StringPtrOutput `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName pulumi.StringPtrOutput `pulumi:"clusterName"`
	// Registered Server last heart beat
	LastHeartBeat pulumi.StringPtrOutput `pulumi:"lastHeartBeat"`
	// Registered Server lastWorkflowId
	LastWorkflowId pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Registered Server Provisioning State
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Registered Server Certificate
	ServerCertificate pulumi.StringPtrOutput `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId pulumi.StringPtrOutput `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementtErrorCode pulumi.IntPtrOutput `pulumi:"serverManagementtErrorCode"`
	// Registered Server OS Version
	ServerOSVersion pulumi.StringPtrOutput `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole pulumi.StringPtrOutput `pulumi:"serverRole"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid pulumi.StringPtrOutput `pulumi:"storageSyncServiceUid"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegisteredServer registers a new resource with the given unique name, arguments, and options.
func NewRegisteredServer(ctx *pulumi.Context,
	name string, args *RegisteredServerArgs, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil || args.StorageSyncServiceName == nil {
		return nil, errors.New("missing required argument 'StorageSyncServiceName'")
	}
	if args == nil {
		args = &RegisteredServerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storagesync/latest:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180402:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20180701:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20181001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190201:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190301:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20190601:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20191001:RegisteredServer"),
		},
		{
			Type: pulumi.String("azurerm:storagesync/v20200301:RegisteredServer"),
		},
	})
	opts = append(opts, aliases)
	var resource RegisteredServer
	err := ctx.RegisterResource("azurerm:storagesync/v20170605preview:RegisteredServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegisteredServer gets an existing RegisteredServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegisteredServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredServerState, opts ...pulumi.ResourceOption) (*RegisteredServer, error) {
	var resource RegisteredServer
	err := ctx.ReadResource("azurerm:storagesync/v20170605preview:RegisteredServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegisteredServer resources.
type registeredServerState struct {
	// Registered Server Agent Version
	AgentVersion *string `pulumi:"agentVersion"`
	// Registered Server clusterId
	ClusterId *string `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName *string `pulumi:"clusterName"`
	// Registered Server last heart beat
	LastHeartBeat *string `pulumi:"lastHeartBeat"`
	// Registered Server lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// Resource name
	Name *string `pulumi:"name"`
	// Registered Server Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// Registered Server Certificate
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId *string `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementtErrorCode *int `pulumi:"serverManagementtErrorCode"`
	// Registered Server OS Version
	ServerOSVersion *string `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole *string `pulumi:"serverRole"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid *string `pulumi:"storageSyncServiceUid"`
	// Resource type
	Type *string `pulumi:"type"`
}

type RegisteredServerState struct {
	// Registered Server Agent Version
	AgentVersion pulumi.StringPtrInput
	// Registered Server clusterId
	ClusterId pulumi.StringPtrInput
	// Registered Server clusterName
	ClusterName pulumi.StringPtrInput
	// Registered Server last heart beat
	LastHeartBeat pulumi.StringPtrInput
	// Registered Server lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Registered Server Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// Registered Server Certificate
	ServerCertificate pulumi.StringPtrInput
	// Registered Server serverId
	ServerId pulumi.StringPtrInput
	// Registered Server Management Error Code
	ServerManagementtErrorCode pulumi.IntPtrInput
	// Registered Server OS Version
	ServerOSVersion pulumi.StringPtrInput
	// Registered Server serverRole
	ServerRole pulumi.StringPtrInput
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (RegisteredServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerState)(nil)).Elem()
}

type registeredServerArgs struct {
	// Registered Server Agent Version
	AgentVersion *string `pulumi:"agentVersion"`
	// Registered Server clusterId
	ClusterId *string `pulumi:"clusterId"`
	// Registered Server clusterName
	ClusterName *string `pulumi:"clusterName"`
	// Registered Server last heart beat
	LastHeartBeat *string `pulumi:"lastHeartBeat"`
	// Registered Server lastWorkflowId
	LastWorkflowId *string `pulumi:"lastWorkflowId"`
	// Registered Server Provisioning State
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Registered Server Certificate
	ServerCertificate *string `pulumi:"serverCertificate"`
	// Registered Server serverId
	ServerId string `pulumi:"serverId"`
	// Registered Server Management Error Code
	ServerManagementtErrorCode *int `pulumi:"serverManagementtErrorCode"`
	// Registered Server OS Version
	ServerOSVersion *string `pulumi:"serverOSVersion"`
	// Registered Server serverRole
	ServerRole *string `pulumi:"serverRole"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid *string `pulumi:"storageSyncServiceUid"`
}

// The set of arguments for constructing a RegisteredServer resource.
type RegisteredServerArgs struct {
	// Registered Server Agent Version
	AgentVersion pulumi.StringPtrInput
	// Registered Server clusterId
	ClusterId pulumi.StringPtrInput
	// Registered Server clusterName
	ClusterName pulumi.StringPtrInput
	// Registered Server last heart beat
	LastHeartBeat pulumi.StringPtrInput
	// Registered Server lastWorkflowId
	LastWorkflowId pulumi.StringPtrInput
	// Registered Server Provisioning State
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Registered Server Certificate
	ServerCertificate pulumi.StringPtrInput
	// Registered Server serverId
	ServerId pulumi.StringInput
	// Registered Server Management Error Code
	ServerManagementtErrorCode pulumi.IntPtrInput
	// Registered Server OS Version
	ServerOSVersion pulumi.StringPtrInput
	// Registered Server serverRole
	ServerRole pulumi.StringPtrInput
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput
	// Registered Server storageSyncServiceUid
	StorageSyncServiceUid pulumi.StringPtrInput
}

func (RegisteredServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredServerArgs)(nil)).Elem()
}
