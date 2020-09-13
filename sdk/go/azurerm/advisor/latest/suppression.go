// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
//
// ## Example Usage
// ### CreateSuppression
//
// ```go
// package main
//
// import (
// 	advisor "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/advisor/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := advisor.NewSuppression(ctx, "suppression", &advisor.SuppressionArgs{
// 			Name:             pulumi.String("suppressionName1"),
// 			RecommendationId: pulumi.String("recommendationId"),
// 			ResourceUri:      pulumi.String("resourceUri"),
// 			Ttl:              pulumi.String("07:00:00:00"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Suppression struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The GUID of the suppression.
	SuppressionId pulumi.StringPtrOutput `pulumi:"suppressionId"`
	// The duration for which the suppression is valid.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSuppression registers a new resource with the given unique name, arguments, and options.
func NewSuppression(ctx *pulumi.Context,
	name string, args *SuppressionArgs, opts ...pulumi.ResourceOption) (*Suppression, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.RecommendationId == nil {
		return nil, errors.New("missing required argument 'RecommendationId'")
	}
	if args == nil || args.ResourceUri == nil {
		return nil, errors.New("missing required argument 'ResourceUri'")
	}
	if args == nil {
		args = &SuppressionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:advisor/v20160712preview:Suppression"),
		},
		{
			Type: pulumi.String("azurerm:advisor/v20170331:Suppression"),
		},
		{
			Type: pulumi.String("azurerm:advisor/v20170419:Suppression"),
		},
		{
			Type: pulumi.String("azurerm:advisor/v20200101:Suppression"),
		},
	})
	opts = append(opts, aliases)
	var resource Suppression
	err := ctx.RegisterResource("azurerm:advisor/latest:Suppression", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuppression gets an existing Suppression resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuppression(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuppressionState, opts ...pulumi.ResourceOption) (*Suppression, error) {
	var resource Suppression
	err := ctx.ReadResource("azurerm:advisor/latest:Suppression", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Suppression resources.
type suppressionState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The GUID of the suppression.
	SuppressionId *string `pulumi:"suppressionId"`
	// The duration for which the suppression is valid.
	Ttl *string `pulumi:"ttl"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type SuppressionState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The GUID of the suppression.
	SuppressionId pulumi.StringPtrInput
	// The duration for which the suppression is valid.
	Ttl pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (SuppressionState) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionState)(nil)).Elem()
}

type suppressionArgs struct {
	// The name of the suppression.
	Name string `pulumi:"name"`
	// The recommendation ID.
	RecommendationId string `pulumi:"recommendationId"`
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri string `pulumi:"resourceUri"`
	// The GUID of the suppression.
	SuppressionId *string `pulumi:"suppressionId"`
	// The duration for which the suppression is valid.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Suppression resource.
type SuppressionArgs struct {
	// The name of the suppression.
	Name pulumi.StringInput
	// The recommendation ID.
	RecommendationId pulumi.StringInput
	// The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
	ResourceUri pulumi.StringInput
	// The GUID of the suppression.
	SuppressionId pulumi.StringPtrInput
	// The duration for which the suppression is valid.
	Ttl pulumi.StringPtrInput
}

func (SuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionArgs)(nil)).Elem()
}
