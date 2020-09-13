// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The essential information related to the peer's ASN.
//
// ## Example Usage
// ### Create a peer ASN
//
// ```go
// package main
//
// import (
// 	peering "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/peering/v20190901preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := peering.NewPeerAsn(ctx, "peerAsn", &peering.PeerAsnArgs{
// 			PeerAsn:     pulumi.Int(65000),
// 			PeerAsnName: pulumi.String("MyPeerAsn"),
// 			PeerContactInfo: &peering.ContactInfoArgs{
// 				Emails: pulumi.StringArray{
// 					pulumi.String("abc@contoso.com"),
// 					pulumi.String("xyz@contoso.com"),
// 				},
// 				Phone: pulumi.StringArray{
// 					pulumi.String("+1 (234) 567-8900"),
// 				},
// 			},
// 			PeerName: pulumi.String("Contoso"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type PeerAsn struct {
	pulumi.CustomResourceState

	// The error message for the validation state
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn pulumi.IntPtrOutput `pulumi:"peerAsn"`
	// The contact information of the peer.
	PeerContactInfo ContactInfoResponsePtrOutput `pulumi:"peerContactInfo"`
	// The name of the peer.
	PeerName pulumi.StringPtrOutput `pulumi:"peerName"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The validation state of the ASN associated with the peer.
	ValidationState pulumi.StringPtrOutput `pulumi:"validationState"`
}

// NewPeerAsn registers a new resource with the given unique name, arguments, and options.
func NewPeerAsn(ctx *pulumi.Context,
	name string, args *PeerAsnArgs, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	if args == nil || args.PeerAsnName == nil {
		return nil, errors.New("missing required argument 'PeerAsnName'")
	}
	if args == nil {
		args = &PeerAsnArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:peering/latest:PeerAsn"),
		},
		{
			Type: pulumi.String("azurerm:peering/v20190801preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azurerm:peering/v20200101preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azurerm:peering/v20200401:PeerAsn"),
		},
	})
	opts = append(opts, aliases)
	var resource PeerAsn
	err := ctx.RegisterResource("azurerm:peering/v20190901preview:PeerAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeerAsn gets an existing PeerAsn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeerAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerAsnState, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	var resource PeerAsn
	err := ctx.ReadResource("azurerm:peering/v20190901preview:PeerAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeerAsn resources.
type peerAsnState struct {
	// The error message for the validation state
	ErrorMessage *string `pulumi:"errorMessage"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn *int `pulumi:"peerAsn"`
	// The contact information of the peer.
	PeerContactInfo *ContactInfoResponse `pulumi:"peerContactInfo"`
	// The name of the peer.
	PeerName *string `pulumi:"peerName"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The validation state of the ASN associated with the peer.
	ValidationState *string `pulumi:"validationState"`
}

type PeerAsnState struct {
	// The error message for the validation state
	ErrorMessage pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn pulumi.IntPtrInput
	// The contact information of the peer.
	PeerContactInfo ContactInfoResponsePtrInput
	// The name of the peer.
	PeerName pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The validation state of the ASN associated with the peer.
	ValidationState pulumi.StringPtrInput
}

func (PeerAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnState)(nil)).Elem()
}

type peerAsnArgs struct {
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn *int `pulumi:"peerAsn"`
	// The peer ASN name.
	PeerAsnName string `pulumi:"peerAsnName"`
	// The contact information of the peer.
	PeerContactInfo *ContactInfo `pulumi:"peerContactInfo"`
	// The name of the peer.
	PeerName *string `pulumi:"peerName"`
	// The validation state of the ASN associated with the peer.
	ValidationState *string `pulumi:"validationState"`
}

// The set of arguments for constructing a PeerAsn resource.
type PeerAsnArgs struct {
	// The Autonomous System Number (ASN) of the peer.
	PeerAsn pulumi.IntPtrInput
	// The peer ASN name.
	PeerAsnName pulumi.StringInput
	// The contact information of the peer.
	PeerContactInfo ContactInfoPtrInput
	// The name of the peer.
	PeerName pulumi.StringPtrInput
	// The validation state of the ASN associated with the peer.
	ValidationState pulumi.StringPtrInput
}

func (PeerAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnArgs)(nil)).Elem()
}
