// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of the WCF relay resource.
//
// ## Example Usage
// ### RelayCreate
//
// ```go
// package main
//
// import (
// 	relay "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/relay/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := relay.NewWCFRelay(ctx, "wcfRelay", &relay.WCFRelayArgs{
// 			NamespaceName:               pulumi.String("sdk-RelayNamespace-9953"),
// 			RelayName:                   pulumi.String("sdk-Relay-Wcf-1194"),
// 			RelayType:                   pulumi.String("NetTcp"),
// 			RequiresClientAuthorization: pulumi.Bool(true),
// 			RequiresTransportSecurity:   pulumi.Bool(true),
// 			ResourceGroupName:           pulumi.String("RG-eg"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type WCFRelay struct {
	pulumi.CustomResourceState

	// The time the WCF relay was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Returns true if the relay is dynamic; otherwise, false.
	IsDynamic pulumi.BoolOutput `pulumi:"isDynamic"`
	// The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount pulumi.IntOutput `pulumi:"listenerCount"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// WCF relay type.
	RelayType pulumi.StringPtrOutput `pulumi:"relayType"`
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrOutput `pulumi:"requiresClientAuthorization"`
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity pulumi.BoolPtrOutput `pulumi:"requiresTransportSecurity"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewWCFRelay registers a new resource with the given unique name, arguments, and options.
func NewWCFRelay(ctx *pulumi.Context,
	name string, args *WCFRelayArgs, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.RelayName == nil {
		return nil, errors.New("missing required argument 'RelayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WCFRelayArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:relay/v20160701:WCFRelay"),
		},
		{
			Type: pulumi.String("azurerm:relay/v20170401:WCFRelay"),
		},
	})
	opts = append(opts, aliases)
	var resource WCFRelay
	err := ctx.RegisterResource("azurerm:relay/latest:WCFRelay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWCFRelay gets an existing WCFRelay resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWCFRelay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WCFRelayState, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	var resource WCFRelay
	err := ctx.ReadResource("azurerm:relay/latest:WCFRelay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WCFRelay resources.
type wcfrelayState struct {
	// The time the WCF relay was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Returns true if the relay is dynamic; otherwise, false.
	IsDynamic *bool `pulumi:"isDynamic"`
	// The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount *int `pulumi:"listenerCount"`
	// Resource name.
	Name *string `pulumi:"name"`
	// WCF relay type.
	RelayType *string `pulumi:"relayType"`
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity *bool `pulumi:"requiresTransportSecurity"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

type WCFRelayState struct {
	// The time the WCF relay was created.
	CreatedAt pulumi.StringPtrInput
	// Returns true if the relay is dynamic; otherwise, false.
	IsDynamic pulumi.BoolPtrInput
	// The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount pulumi.IntPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// WCF relay type.
	RelayType pulumi.StringPtrInput
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity pulumi.BoolPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The time the namespace was updated.
	UpdatedAt pulumi.StringPtrInput
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (WCFRelayState) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayState)(nil)).Elem()
}

type wcfrelayArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The relay name.
	RelayName string `pulumi:"relayName"`
	// WCF relay type.
	RelayType *string `pulumi:"relayType"`
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity *bool `pulumi:"requiresTransportSecurity"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a WCFRelay resource.
type WCFRelayArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput
	// The relay name.
	RelayName pulumi.StringInput
	// WCF relay type.
	RelayType pulumi.StringPtrInput
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (WCFRelayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayArgs)(nil)).Elem()
}
