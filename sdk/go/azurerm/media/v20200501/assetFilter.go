// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Asset Filter.
//
// ## Example Usage
// ### Create an Asset Filter
//
// ```go
// package main
//
// import (
// 	media "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/media/v20200501"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := media.NewAssetFilter(ctx, "assetFilter", &media.AssetFilterArgs{
// 			AccountName: pulumi.String("contosomedia"),
// 			AssetName:   pulumi.String("ClimbingMountRainer"),
// 			FilterName:  pulumi.String("newAssetFilter"),
// 			FirstQuality: &media.FirstQualityArgs{
// 				Bitrate: pulumi.Int(128000),
// 			},
// 			PresentationTimeRange: &media.PresentationTimeRangeArgs{
// 				EndTimestamp:               pulumi.Int(170000000),
// 				ForceEndTimestamp:          pulumi.Bool(false),
// 				LiveBackoffDuration:        pulumi.Int(0),
// 				PresentationWindowDuration: pulumi.Int(9.223372036854776e+18),
// 				StartTimestamp:             pulumi.Int(0),
// 				Timescale:                  pulumi.Int(10000000),
// 			},
// 			ResourceGroupName: pulumi.String("contoso"),
// 			Tracks: media.FilterTrackSelectionArray{
// 				&media.FilterTrackSelectionArgs{
// 					TrackSelections: media.FilterTrackPropertyConditionArray{
// 						&media.FilterTrackPropertyConditionArgs{
// 							Operation: pulumi.String("Equal"),
// 							Property:  pulumi.String("Type"),
// 							Value:     pulumi.String("Audio"),
// 						},
// 						&media.FilterTrackPropertyConditionArgs{
// 							Operation: pulumi.String("NotEqual"),
// 							Property:  pulumi.String("Language"),
// 							Value:     pulumi.String("en"),
// 						},
// 						&media.FilterTrackPropertyConditionArgs{
// 							Operation: pulumi.String("NotEqual"),
// 							Property:  pulumi.String("FourCC"),
// 							Value:     pulumi.String("EC-3"),
// 						},
// 					},
// 				},
// 				&media.FilterTrackSelectionArgs{
// 					TrackSelections: media.FilterTrackPropertyConditionArray{
// 						&media.FilterTrackPropertyConditionArgs{
// 							Operation: pulumi.String("Equal"),
// 							Property:  pulumi.String("Type"),
// 							Value:     pulumi.String("Video"),
// 						},
// 						&media.FilterTrackPropertyConditionArgs{
// 							Operation: pulumi.String("Equal"),
// 							Property:  pulumi.String("Bitrate"),
// 							Value:     pulumi.String("3000000-5000000"),
// 						},
// 					},
// 				},
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
type AssetFilter struct {
	pulumi.CustomResourceState

	// The first quality.
	FirstQuality FirstQualityResponsePtrOutput `pulumi:"firstQuality"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The presentation time range.
	PresentationTimeRange PresentationTimeRangeResponsePtrOutput `pulumi:"presentationTimeRange"`
	// The tracks selection conditions.
	Tracks FilterTrackSelectionResponseArrayOutput `pulumi:"tracks"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssetFilter registers a new resource with the given unique name, arguments, and options.
func NewAssetFilter(ctx *pulumi.Context,
	name string, args *AssetFilterArgs, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.AssetName == nil {
		return nil, errors.New("missing required argument 'AssetName'")
	}
	if args == nil || args.FilterName == nil {
		return nil, errors.New("missing required argument 'FilterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AssetFilterArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:media/latest:AssetFilter"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180701:AssetFilter"),
		},
	})
	opts = append(opts, aliases)
	var resource AssetFilter
	err := ctx.RegisterResource("azurerm:media/v20200501:AssetFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssetFilter gets an existing AssetFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetFilterState, opts ...pulumi.ResourceOption) (*AssetFilter, error) {
	var resource AssetFilter
	err := ctx.ReadResource("azurerm:media/v20200501:AssetFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssetFilter resources.
type assetFilterState struct {
	// The first quality.
	FirstQuality *FirstQualityResponse `pulumi:"firstQuality"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The presentation time range.
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	// The tracks selection conditions.
	Tracks []FilterTrackSelectionResponse `pulumi:"tracks"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type AssetFilterState struct {
	// The first quality.
	FirstQuality FirstQualityResponsePtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The presentation time range.
	PresentationTimeRange PresentationTimeRangeResponsePtrInput
	// The tracks selection conditions.
	Tracks FilterTrackSelectionResponseArrayInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (AssetFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterState)(nil)).Elem()
}

type assetFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The Asset Filter name
	FilterName string `pulumi:"filterName"`
	// The first quality.
	FirstQuality *FirstQuality `pulumi:"firstQuality"`
	// The presentation time range.
	PresentationTimeRange *PresentationTimeRange `pulumi:"presentationTimeRange"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tracks selection conditions.
	Tracks []FilterTrackSelection `pulumi:"tracks"`
}

// The set of arguments for constructing a AssetFilter resource.
type AssetFilterArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Asset name.
	AssetName pulumi.StringInput
	// The Asset Filter name
	FilterName pulumi.StringInput
	// The first quality.
	FirstQuality FirstQualityPtrInput
	// The presentation time range.
	PresentationTimeRange PresentationTimeRangePtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The tracks selection conditions.
	Tracks FilterTrackSelectionArrayInput
}

func (AssetFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetFilterArgs)(nil)).Elem()
}
