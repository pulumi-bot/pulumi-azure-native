// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Payload of the blockchain member which is exposed in the request/response of the resource provider.
//
// ## Example Usage
// ### BlockchainMembers_Create
//
// ```go
// package main
//
// import (
// 	blockchain "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/blockchain/v20180601preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := blockchain.NewBlockchainMember(ctx, "blockchainMember", &blockchain.BlockchainMemberArgs{
// 			BlockchainMemberName:                pulumi.String("contosemember1"),
// 			Consortium:                          pulumi.String("ContoseConsortium"),
// 			ConsortiumManagementAccountPassword: pulumi.String("1234abcdEFG1"),
// 			Location:                            pulumi.String("southeastasia"),
// 			Password:                            pulumi.String("1234abcdEFG1"),
// 			Protocol:                            pulumi.String("Quorum"),
// 			ResourceGroupName:                   pulumi.String("mygroup"),
// 			ValidatorNodesSku: &blockchain.BlockchainMemberNodesSkuArgs{
// 				Capacity: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type BlockchainMember struct {
	pulumi.CustomResourceState

	// Gets or sets the consortium for the blockchain member.
	Consortium pulumi.StringPtrOutput `pulumi:"consortium"`
	// Gets the managed consortium management account address.
	ConsortiumManagementAccountAddress pulumi.StringOutput `pulumi:"consortiumManagementAccountAddress"`
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword pulumi.StringPtrOutput `pulumi:"consortiumManagementAccountPassword"`
	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName pulumi.StringPtrOutput `pulumi:"consortiumMemberDisplayName"`
	// Gets the role of the member in the consortium.
	ConsortiumRole pulumi.StringPtrOutput `pulumi:"consortiumRole"`
	// Gets the dns endpoint of the blockchain member.
	Dns pulumi.StringOutput `pulumi:"dns"`
	// Gets or sets firewall rules
	FirewallRules FirewallRuleResponseArrayOutput `pulumi:"firewallRules"`
	// The GEO location of the blockchain service.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the basic auth password of the blockchain member.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Gets or sets the blockchain protocol.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Gets or sets the blockchain member provision state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets the public key of the blockchain member (default transaction node).
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// Gets the Ethereum root contract address of the blockchain.
	RootContractAddress pulumi.StringOutput `pulumi:"rootContractAddress"`
	// Gets or sets the blockchain member Sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the service - e.g. "Microsoft.Blockchain"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the auth user name of the blockchain member.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSku BlockchainMemberNodesSkuResponsePtrOutput `pulumi:"validatorNodesSku"`
}

// NewBlockchainMember registers a new resource with the given unique name, arguments, and options.
func NewBlockchainMember(ctx *pulumi.Context,
	name string, args *BlockchainMemberArgs, opts ...pulumi.ResourceOption) (*BlockchainMember, error) {
	if args == nil || args.BlockchainMemberName == nil {
		return nil, errors.New("missing required argument 'BlockchainMemberName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BlockchainMemberArgs{}
	}
	var resource BlockchainMember
	err := ctx.RegisterResource("azurerm:blockchain/v20180601preview:BlockchainMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockchainMember gets an existing BlockchainMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockchainMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockchainMemberState, opts ...pulumi.ResourceOption) (*BlockchainMember, error) {
	var resource BlockchainMember
	err := ctx.ReadResource("azurerm:blockchain/v20180601preview:BlockchainMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockchainMember resources.
type blockchainMemberState struct {
	// Gets or sets the consortium for the blockchain member.
	Consortium *string `pulumi:"consortium"`
	// Gets the managed consortium management account address.
	ConsortiumManagementAccountAddress *string `pulumi:"consortiumManagementAccountAddress"`
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string `pulumi:"consortiumManagementAccountPassword"`
	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName *string `pulumi:"consortiumMemberDisplayName"`
	// Gets the role of the member in the consortium.
	ConsortiumRole *string `pulumi:"consortiumRole"`
	// Gets the dns endpoint of the blockchain member.
	Dns *string `pulumi:"dns"`
	// Gets or sets firewall rules
	FirewallRules []FirewallRuleResponse `pulumi:"firewallRules"`
	// The GEO location of the blockchain service.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Sets the basic auth password of the blockchain member.
	Password *string `pulumi:"password"`
	// Gets or sets the blockchain protocol.
	Protocol *string `pulumi:"protocol"`
	// Gets or sets the blockchain member provision state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the public key of the blockchain member (default transaction node).
	PublicKey *string `pulumi:"publicKey"`
	// Gets the Ethereum root contract address of the blockchain.
	RootContractAddress *string `pulumi:"rootContractAddress"`
	// Gets or sets the blockchain member Sku.
	Sku *SkuResponse `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the service - e.g. "Microsoft.Blockchain"
	Type *string `pulumi:"type"`
	// Gets the auth user name of the blockchain member.
	UserName *string `pulumi:"userName"`
	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSku *BlockchainMemberNodesSkuResponse `pulumi:"validatorNodesSku"`
}

type BlockchainMemberState struct {
	// Gets or sets the consortium for the blockchain member.
	Consortium pulumi.StringPtrInput
	// Gets the managed consortium management account address.
	ConsortiumManagementAccountAddress pulumi.StringPtrInput
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword pulumi.StringPtrInput
	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName pulumi.StringPtrInput
	// Gets the role of the member in the consortium.
	ConsortiumRole pulumi.StringPtrInput
	// Gets the dns endpoint of the blockchain member.
	Dns pulumi.StringPtrInput
	// Gets or sets firewall rules
	FirewallRules FirewallRuleResponseArrayInput
	// The GEO location of the blockchain service.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Sets the basic auth password of the blockchain member.
	Password pulumi.StringPtrInput
	// Gets or sets the blockchain protocol.
	Protocol pulumi.StringPtrInput
	// Gets or sets the blockchain member provision state.
	ProvisioningState pulumi.StringPtrInput
	// Gets the public key of the blockchain member (default transaction node).
	PublicKey pulumi.StringPtrInput
	// Gets the Ethereum root contract address of the blockchain.
	RootContractAddress pulumi.StringPtrInput
	// Gets or sets the blockchain member Sku.
	Sku SkuResponsePtrInput
	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags pulumi.StringMapInput
	// The type of the service - e.g. "Microsoft.Blockchain"
	Type pulumi.StringPtrInput
	// Gets the auth user name of the blockchain member.
	UserName pulumi.StringPtrInput
	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSku BlockchainMemberNodesSkuResponsePtrInput
}

func (BlockchainMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockchainMemberState)(nil)).Elem()
}

type blockchainMemberArgs struct {
	// Blockchain member name.
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	// Gets or sets the consortium for the blockchain member.
	Consortium *string `pulumi:"consortium"`
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string `pulumi:"consortiumManagementAccountPassword"`
	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName *string `pulumi:"consortiumMemberDisplayName"`
	// Gets the role of the member in the consortium.
	ConsortiumRole *string `pulumi:"consortiumRole"`
	// Gets or sets firewall rules
	FirewallRules []FirewallRule `pulumi:"firewallRules"`
	// The GEO location of the blockchain service.
	Location *string `pulumi:"location"`
	// Sets the basic auth password of the blockchain member.
	Password *string `pulumi:"password"`
	// Gets or sets the blockchain protocol.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the blockchain member Sku.
	Sku *Sku `pulumi:"sku"`
	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSku *BlockchainMemberNodesSku `pulumi:"validatorNodesSku"`
}

// The set of arguments for constructing a BlockchainMember resource.
type BlockchainMemberArgs struct {
	// Blockchain member name.
	BlockchainMemberName pulumi.StringInput
	// Gets or sets the consortium for the blockchain member.
	Consortium pulumi.StringPtrInput
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword pulumi.StringPtrInput
	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName pulumi.StringPtrInput
	// Gets the role of the member in the consortium.
	ConsortiumRole pulumi.StringPtrInput
	// Gets or sets firewall rules
	FirewallRules FirewallRuleArrayInput
	// The GEO location of the blockchain service.
	Location pulumi.StringPtrInput
	// Sets the basic auth password of the blockchain member.
	Password pulumi.StringPtrInput
	// Gets or sets the blockchain protocol.
	Protocol pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the blockchain member Sku.
	Sku SkuPtrInput
	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags pulumi.StringMapInput
	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSku BlockchainMemberNodesSkuPtrInput
}

func (BlockchainMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockchainMemberArgs)(nil)).Elem()
}
