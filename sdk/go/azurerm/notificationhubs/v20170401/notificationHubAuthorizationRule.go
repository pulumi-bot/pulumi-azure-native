// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of a Namespace AuthorizationRules.
//
// ## Example Usage
// ### NotificationHubAuthorizationRuleCreate
//
// ```go
// package main
//
// import (
// 	notificationhubs "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/notificationhubs/v20170401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notificationhubs.NewNotificationHubAuthorizationRule(ctx, "notificationHubAuthorizationRule", &notificationhubs.NotificationHubAuthorizationRuleArgs{
// 			AuthorizationRuleName: pulumi.String("DefaultListenSharedAccessSignature"),
// 			NamespaceName:         pulumi.String("nh-sdk-ns"),
// 			NotificationHubName:   pulumi.String("nh-sdk-hub"),
// 			ResourceGroupName:     pulumi.String("5ktrial"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type NotificationHubAuthorizationRule struct {
	pulumi.CustomResourceState

	// A string that describes the claim type
	ClaimType pulumi.StringOutput `pulumi:"claimType"`
	// A string that describes the claim value
	ClaimValue pulumi.StringOutput `pulumi:"claimValue"`
	// The created time for this rule
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// A string that describes the authorization rule.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The last modified time for this rule
	ModifiedTime pulumi.StringOutput `pulumi:"modifiedTime"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The revision number for the rule
	Revision pulumi.IntOutput `pulumi:"revision"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// The sku of the created namespace
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotificationHubAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *NotificationHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	if args == nil || args.AuthorizationRuleName == nil {
		return nil, errors.New("missing required argument 'AuthorizationRuleName'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.NotificationHubName == nil {
		return nil, errors.New("missing required argument 'NotificationHubName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NotificationHubAuthorizationRuleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:notificationhubs/latest:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azurerm:notificationhubs/v20160301:NotificationHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NotificationHubAuthorizationRule
	err := ctx.RegisterResource("azurerm:notificationhubs/v20170401:NotificationHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationHubAuthorizationRule gets an existing NotificationHubAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	var resource NotificationHubAuthorizationRule
	err := ctx.ReadResource("azurerm:notificationhubs/v20170401:NotificationHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationHubAuthorizationRule resources.
type notificationHubAuthorizationRuleState struct {
	// A string that describes the claim type
	ClaimType *string `pulumi:"claimType"`
	// A string that describes the claim value
	ClaimValue *string `pulumi:"claimValue"`
	// The created time for this rule
	CreatedTime *string `pulumi:"createdTime"`
	// A string that describes the authorization rule.
	KeyName *string `pulumi:"keyName"`
	// Resource location
	Location *string `pulumi:"location"`
	// The last modified time for this rule
	ModifiedTime *string `pulumi:"modifiedTime"`
	// Resource name
	Name *string `pulumi:"name"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The revision number for the rule
	Revision *int `pulumi:"revision"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The sku of the created namespace
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type NotificationHubAuthorizationRuleState struct {
	// A string that describes the claim type
	ClaimType pulumi.StringPtrInput
	// A string that describes the claim value
	ClaimValue pulumi.StringPtrInput
	// The created time for this rule
	CreatedTime pulumi.StringPtrInput
	// A string that describes the authorization rule.
	KeyName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The last modified time for this rule
	ModifiedTime pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey pulumi.StringPtrInput
	// The revision number for the rule
	Revision pulumi.IntPtrInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey pulumi.StringPtrInput
	// The sku of the created namespace
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (NotificationHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleState)(nil)).Elem()
}

type notificationHubAuthorizationRuleArgs struct {
	// Authorization Rule Name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// The notification hub name.
	NotificationHubName string `pulumi:"notificationHubName"`
	// Properties of the Namespace AuthorizationRules.
	Properties SharedAccessAuthorizationRuleProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NotificationHubAuthorizationRule resource.
type NotificationHubAuthorizationRuleArgs struct {
	// Authorization Rule Name.
	AuthorizationRuleName pulumi.StringInput
	// The namespace name.
	NamespaceName pulumi.StringInput
	// The notification hub name.
	NotificationHubName pulumi.StringInput
	// Properties of the Namespace AuthorizationRules.
	Properties SharedAccessAuthorizationRulePropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NotificationHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleArgs)(nil)).Elem()
}
