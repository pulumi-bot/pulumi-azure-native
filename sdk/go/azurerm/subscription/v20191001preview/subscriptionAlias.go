// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Subscription Information with the alias.
//
// ## Example Usage
// ### CreateAlias
//
// ```go
// package main
//
// import (
// 	subscription "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/subscription/v20191001preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := subscription.NewSubscriptionAlias(ctx, "subscriptionAlias", &subscription.SubscriptionAliasArgs{
// 			AliasName: pulumi.String("aliasForNewSub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type SubscriptionAlias struct {
	pulumi.CustomResourceState

	// Alias ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// Put Alias response properties.
	Properties PutAliasResponsePropertiesResponseOutput `pulumi:"properties"`
	// Resource type, Microsoft.Subscription/aliases.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubscriptionAlias registers a new resource with the given unique name, arguments, and options.
func NewSubscriptionAlias(ctx *pulumi.Context,
	name string, args *SubscriptionAliasArgs, opts ...pulumi.ResourceOption) (*SubscriptionAlias, error) {
	if args == nil || args.AliasName == nil {
		return nil, errors.New("missing required argument 'AliasName'")
	}
	if args == nil {
		args = &SubscriptionAliasArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:subscription/latest:SubscriptionAlias"),
		},
		{
			Type: pulumi.String("azurerm:subscription/v20200901:SubscriptionAlias"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionAlias
	err := ctx.RegisterResource("azurerm:subscription/v20191001preview:SubscriptionAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscriptionAlias gets an existing SubscriptionAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscriptionAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionAliasState, opts ...pulumi.ResourceOption) (*SubscriptionAlias, error) {
	var resource SubscriptionAlias
	err := ctx.ReadResource("azurerm:subscription/v20191001preview:SubscriptionAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscriptionAlias resources.
type subscriptionAliasState struct {
	// Alias ID.
	Name *string `pulumi:"name"`
	// Put Alias response properties.
	Properties *PutAliasResponsePropertiesResponse `pulumi:"properties"`
	// Resource type, Microsoft.Subscription/aliases.
	Type *string `pulumi:"type"`
}

type SubscriptionAliasState struct {
	// Alias ID.
	Name pulumi.StringPtrInput
	// Put Alias response properties.
	Properties PutAliasResponsePropertiesResponsePtrInput
	// Resource type, Microsoft.Subscription/aliases.
	Type pulumi.StringPtrInput
}

func (SubscriptionAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAliasState)(nil)).Elem()
}

type subscriptionAliasArgs struct {
	// Alias Name
	AliasName string `pulumi:"aliasName"`
	// Put alias request properties.
	Properties *PutAliasRequestProperties `pulumi:"properties"`
}

// The set of arguments for constructing a SubscriptionAlias resource.
type SubscriptionAliasArgs struct {
	// Alias Name
	AliasName pulumi.StringInput
	// Put alias request properties.
	Properties PutAliasRequestPropertiesPtrInput
}

func (SubscriptionAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAliasArgs)(nil)).Elem()
}
