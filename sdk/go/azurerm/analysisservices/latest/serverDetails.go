// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an instance of an Analysis Services resource.
type ServerDetails struct {
	pulumi.CustomResourceState

	// A collection of AS server administrators
	AsAdministrators ServerAdministratorsResponsePtrOutput `pulumi:"asAdministrators"`
	// The SAS container URI to the backup container.
	BackupBlobContainerUri pulumi.StringPtrOutput `pulumi:"backupBlobContainerUri"`
	// The gateway details configured for the AS server.
	GatewayDetails GatewayDetailsResponsePtrOutput `pulumi:"gatewayDetails"`
	// The firewall settings for the AS server.
	IpV4FirewallSettings IPv4FirewallSettingsResponsePtrOutput `pulumi:"ipV4FirewallSettings"`
	// Location of the Analysis Services resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Analysis Services resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode pulumi.StringPtrOutput `pulumi:"querypoolConnectionMode"`
	// The full name of the Analysis Services resource.
	ServerFullName pulumi.StringOutput `pulumi:"serverFullName"`
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuResponseOutput `pulumi:"sku"`
	// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the Analysis Services resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerDetails registers a new resource with the given unique name, arguments, and options.
func NewServerDetails(ctx *pulumi.Context,
	name string, args *ServerDetailsArgs, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ServerDetailsArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:analysisservices/v20160516:ServerDetails"),
		},
		{
			Type: pulumi.String("azurerm:analysisservices/v20170714:ServerDetails"),
		},
		{
			Type: pulumi.String("azurerm:analysisservices/v20170801:ServerDetails"),
		},
		{
			Type: pulumi.String("azurerm:analysisservices/v20170801beta:ServerDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerDetails
	err := ctx.RegisterResource("azurerm:analysisservices/latest:ServerDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerDetails gets an existing ServerDetails resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerDetailsState, opts ...pulumi.ResourceOption) (*ServerDetails, error) {
	var resource ServerDetails
	err := ctx.ReadResource("azurerm:analysisservices/latest:ServerDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerDetails resources.
type serverDetailsState struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministratorsResponse `pulumi:"asAdministrators"`
	// The SAS container URI to the backup container.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// The gateway details configured for the AS server.
	GatewayDetails *GatewayDetailsResponse `pulumi:"gatewayDetails"`
	// The firewall settings for the AS server.
	IpV4FirewallSettings *IPv4FirewallSettingsResponse `pulumi:"ipV4FirewallSettings"`
	// Location of the Analysis Services resource.
	Location *string `pulumi:"location"`
	// The name of the Analysis Services resource.
	Name *string `pulumi:"name"`
	// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState *string `pulumi:"provisioningState"`
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode *string `pulumi:"querypoolConnectionMode"`
	// The full name of the Analysis Services resource.
	ServerFullName *string `pulumi:"serverFullName"`
	// The SKU of the Analysis Services resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State *string `pulumi:"state"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// The type of the Analysis Services resource.
	Type *string `pulumi:"type"`
}

type ServerDetailsState struct {
	// A collection of AS server administrators
	AsAdministrators ServerAdministratorsResponsePtrInput
	// The SAS container URI to the backup container.
	BackupBlobContainerUri pulumi.StringPtrInput
	// The gateway details configured for the AS server.
	GatewayDetails GatewayDetailsResponsePtrInput
	// The firewall settings for the AS server.
	IpV4FirewallSettings IPv4FirewallSettingsResponsePtrInput
	// Location of the Analysis Services resource.
	Location pulumi.StringPtrInput
	// The name of the Analysis Services resource.
	Name pulumi.StringPtrInput
	// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringPtrInput
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode pulumi.StringPtrInput
	// The full name of the Analysis Services resource.
	ServerFullName pulumi.StringPtrInput
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuResponsePtrInput
	// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State pulumi.StringPtrInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
	// The type of the Analysis Services resource.
	Type pulumi.StringPtrInput
}

func (ServerDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsState)(nil)).Elem()
}

type serverDetailsArgs struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministrators `pulumi:"asAdministrators"`
	// The SAS container URI to the backup container.
	BackupBlobContainerUri *string `pulumi:"backupBlobContainerUri"`
	// The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails `pulumi:"gatewayDetails"`
	// The firewall settings for the AS server.
	IpV4FirewallSettings *IPv4FirewallSettings `pulumi:"ipV4FirewallSettings"`
	// Location of the Analysis Services resource.
	Location string `pulumi:"location"`
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode *string `pulumi:"querypoolConnectionMode"`
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName string `pulumi:"serverName"`
	// The SKU of the Analysis Services resource.
	Sku ResourceSku `pulumi:"sku"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServerDetails resource.
type ServerDetailsArgs struct {
	// A collection of AS server administrators
	AsAdministrators ServerAdministratorsPtrInput
	// The SAS container URI to the backup container.
	BackupBlobContainerUri pulumi.StringPtrInput
	// The gateway details configured for the AS server.
	GatewayDetails GatewayDetailsPtrInput
	// The firewall settings for the AS server.
	IpV4FirewallSettings IPv4FirewallSettingsPtrInput
	// Location of the Analysis Services resource.
	Location pulumi.StringInput
	// How the read-write server's participation in the query pool is controlled.<br/>It can have the following values: <ul><li>readOnly - indicates that the read-write server is intended not to participate in query operations</li><li>all - indicates that the read-write server can participate in query operations</li></ul>Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode pulumi.StringPtrInput
	// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput
	// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
	ServerName pulumi.StringInput
	// The SKU of the Analysis Services resource.
	Sku ResourceSkuInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
}

func (ServerDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverDetailsArgs)(nil)).Elem()
}
