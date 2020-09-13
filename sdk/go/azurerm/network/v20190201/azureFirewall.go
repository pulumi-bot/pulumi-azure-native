// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Azure Firewall resource
//
// ## Example Usage
// ### Create Azure Firewall
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20190201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewAzureFirewall(ctx, "azureFirewall", &network.AzureFirewallArgs{
// 			ApplicationRuleCollections: network.AzureFirewallApplicationRuleCollectionArray{
// 				&network.AzureFirewallApplicationRuleCollectionArgs{
// 					Name: pulumi.String("apprulecoll"),
// 				},
// 			},
// 			AzureFirewallName: pulumi.String("azurefirewall"),
// 			IpConfigurations: network.AzureFirewallIPConfigurationArray{
// 				&network.AzureFirewallIPConfigurationArgs{
// 					Name: pulumi.String("azureFirewallIpConfiguration"),
// 				},
// 			},
// 			NatRuleCollections: network.AzureFirewallNatRuleCollectionArray{
// 				&network.AzureFirewallNatRuleCollectionArgs{
// 					Name: pulumi.String("natrulecoll"),
// 				},
// 			},
// 			NetworkRuleCollections: network.AzureFirewallNetworkRuleCollectionArray{
// 				&network.AzureFirewallNetworkRuleCollectionArgs{
// 					Name: pulumi.String("netrulecoll"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			Tags: pulumi.StringMap{
// 				"key1": pulumi.String("value1"),
// 			},
// 			ThreatIntelMode: pulumi.String("Alert"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type AzureFirewall struct {
	pulumi.CustomResourceState

	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayOutput `pulumi:"applicationRuleCollections"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionResponseArrayOutput `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionResponseArrayOutput `pulumi:"networkRuleCollections"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrOutput `pulumi:"threatIntelMode"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAzureFirewall registers a new resource with the given unique name, arguments, and options.
func NewAzureFirewall(ctx *pulumi.Context,
	name string, args *AzureFirewallArgs, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	if args == nil || args.AzureFirewallName == nil {
		return nil, errors.New("missing required argument 'AzureFirewallName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AzureFirewallArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:AzureFirewall"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:AzureFirewall"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureFirewall
	err := ctx.RegisterResource("azurerm:network/v20190201:AzureFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureFirewall gets an existing AzureFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureFirewallState, opts ...pulumi.ResourceOption) (*AzureFirewall, error) {
	var resource AzureFirewall
	err := ctx.ReadResource("azurerm:network/v20190201:AzureFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureFirewall resources.
type azureFirewallState struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollectionResponse `pulumi:"applicationRuleCollections"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollectionResponse `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollectionResponse `pulumi:"networkRuleCollections"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AzureFirewallState struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionResponseArrayInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionResponseArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionResponseArrayInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AzureFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallState)(nil)).Elem()
}

type azureFirewallArgs struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `pulumi:"applicationRuleCollections"`
	// The name of the Azure Firewall.
	AzureFirewallName string `pulumi:"azureFirewallName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollection `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollection `pulumi:"networkRuleCollections"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
}

// The set of arguments for constructing a AzureFirewall resource.
type AzureFirewallArgs struct {
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections AzureFirewallApplicationRuleCollectionArrayInput
	// The name of the Azure Firewall.
	AzureFirewallName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configuration of the Azure Firewall resource.
	IpConfigurations AzureFirewallIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections AzureFirewallNatRuleCollectionArrayInput
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections AzureFirewallNetworkRuleCollectionArrayInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
}

func (AzureFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureFirewallArgs)(nil)).Elem()
}
