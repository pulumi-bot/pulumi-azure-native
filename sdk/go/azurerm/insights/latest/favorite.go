// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Properties that define a favorite that is associated to an Application Insights component.
//
// ## Example Usage
// ### FavoriteAdd
//
// ```go
// package main
//
// import (
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewFavorite(ctx, "favorite", &insights.FavoriteArgs{
// 			Config:                  pulumi.String("{\"MEDataModelRawJSON\":\"{\\n  \\\"version\\\": \\\"1.4.1\\\",\\n  \\\"isCustomDataModel\\\": true,\\n  \\\"items\\\": [\\n    {\\n      \\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Sum\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"fail\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 2,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"greenHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\n      \\\"chartType\\\": \\\"Bar\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"magentaHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\n      \\\"grouping\\\": {\\n        \\\"kind\\\": \\\"ByDimension\\\",\\n        \\\"dimension\\\": \\\"context.application.version\\\"\\n      },\\n      \\\"chartType\\\": \\\"Grid\\\",\\n      \\\"chartHeight\\\": 1,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"basicException.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"\\n        },\\n        {\\n          \\\"id\\\": \\\"requestFailed.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": true,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"blueHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    }\\n  ],\\n  \\\"currentFilter\\\": {\\n    \\\"eventTypes\\\": [\\n      1,\\n      2\\n    ],\\n    \\\"typeFacets\\\": {},\\n    \\\"isPermissive\\\": false\\n  },\\n  \\\"timeContext\\\": {\\n    \\\"durationMs\\\": 75600000,\\n    \\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\n    \\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\n    \\\"isInitialTime\\\": false,\\n    \\\"grain\\\": 1,\\n    \\\"useDashboardTimeRange\\\": false\\n  },\\n  \\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\n  \\\"timeSource\\\": 0\\n}\"}"),
// 			FavoriteId:              pulumi.String("deadb33f-8bee-4d3b-a059-9be8dac93960"),
// 			FavoriteType:            pulumi.String("shared"),
// 			IsGeneratedFromTemplate: pulumi.Bool(false),
// 			Name:                    pulumi.String("Blah Blah Blah"),
// 			ResourceGroupName:       pulumi.String("my-resource-group"),
// 			ResourceName:            pulumi.String("my-ai-component"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("TagSample01"),
// 				pulumi.String("TagSample02"),
// 			},
// 			Version: pulumi.String("ME"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Favorite struct {
	pulumi.CustomResourceState

	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrOutput `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId pulumi.StringOutput `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrOutput `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrOutput `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The source of the favorite definition.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewFavorite registers a new resource with the given unique name, arguments, and options.
func NewFavorite(ctx *pulumi.Context,
	name string, args *FavoriteArgs, opts ...pulumi.ResourceOption) (*Favorite, error) {
	if args == nil || args.FavoriteId == nil {
		return nil, errors.New("missing required argument 'FavoriteId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &FavoriteArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/v20150501:Favorite"),
		},
	})
	opts = append(opts, aliases)
	var resource Favorite
	err := ctx.RegisterResource("azurerm:insights/latest:Favorite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFavorite gets an existing Favorite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFavorite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FavoriteState, opts ...pulumi.ResourceOption) (*Favorite, error) {
	var resource Favorite
	err := ctx.ReadResource("azurerm:insights/latest:Favorite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Favorite resources.
type favoriteState struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// Internally assigned unique id of the favorite definition.
	FavoriteId *string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *string `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified *string `pulumi:"timeModified"`
	// Unique user id of the specific user that owns this favorite.
	UserId *string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

type FavoriteState struct {
	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrInput
	// Internally assigned unique id of the favorite definition.
	FavoriteId pulumi.StringPtrInput
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrInput
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	// The user-defined name of the favorite.
	Name pulumi.StringPtrInput
	// The source of the favorite definition.
	SourceType pulumi.StringPtrInput
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayInput
	// Date and time in UTC of the last modification that was made to this favorite definition.
	TimeModified pulumi.StringPtrInput
	// Unique user id of the specific user that owns this favorite.
	UserId pulumi.StringPtrInput
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrInput
}

func (FavoriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteState)(nil)).Elem()
}

type favoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config *string `pulumi:"config"`
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId string `pulumi:"favoriteId"`
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType *string `pulumi:"favoriteType"`
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate *bool `pulumi:"isGeneratedFromTemplate"`
	// The user-defined name of the favorite.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// The source of the favorite definition.
	SourceType *string `pulumi:"sourceType"`
	// A list of 0 or more tags that are associated with this favorite definition
	Tags []string `pulumi:"tags"`
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Favorite resource.
type FavoriteArgs struct {
	// Favorite category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
	Config pulumi.StringPtrInput
	// The Id of a specific favorite defined in the Application Insights component
	FavoriteId pulumi.StringInput
	// Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	FavoriteType pulumi.StringPtrInput
	// Flag denoting wether or not this favorite was generated from a template.
	IsGeneratedFromTemplate pulumi.BoolPtrInput
	// The user-defined name of the favorite.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// The source of the favorite definition.
	SourceType pulumi.StringPtrInput
	// A list of 0 or more tags that are associated with this favorite definition
	Tags pulumi.StringArrayInput
	// This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
	Version pulumi.StringPtrInput
}

func (FavoriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*favoriteArgs)(nil)).Elem()
}
