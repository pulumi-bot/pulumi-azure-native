// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Log Analytics QueryPack-Query definition.
//
// ## Example Usage
// ### QueryPut
//
// ```go
// package main
//
// import (
// 	insights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/insights/v20190901preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := insights.NewQuery(ctx, "query", &insights.QueryArgs{
// 			Body:          pulumi.String("let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n"),
// 			Description:   pulumi.String("my description"),
// 			DisplayName:   pulumi.String("Exceptions - New in the last 24 hours"),
// 			Id:            pulumi.String("a449f8af-8e64-4b3a-9b16-5a7165ff98c4"),
// 			QueryPackName: pulumi.String("my-querypack"),
// 			Related: &insights.LogAnalyticsQueryPackQueryPropertiesRelatedArgs{
// 				Categories: pulumi.StringArray{
// 					pulumi.String("analytics"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("my-resource-group"),
// 			Tags: pulumi.StringArrayMap{
// 				"my-label": pulumi.StringArray{
// 					pulumi.String("label1"),
// 				},
// 				"my-other-label": pulumi.StringArray{
// 					pulumi.String("label2"),
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
type Query struct {
	pulumi.CustomResourceState

	// Object Id of user creating the query.
	Author pulumi.StringOutput `pulumi:"author"`
	// Body of the query.
	Body pulumi.StringOutput `pulumi:"body"`
	// Description of the query.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique display name for your query within the Query Pack.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional properties that can be set for the query.
	Properties pulumi.MapOutput `pulumi:"properties"`
	// The related metadata items for the function.
	Related LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput `pulumi:"related"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags associated with the query.
	Tags pulumi.StringArrayMapOutput `pulumi:"tags"`
	// Creation Date for the Log Analytics Query, in ISO 8601 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics Query, in ISO 8601 format.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQuery registers a new resource with the given unique name, arguments, and options.
func NewQuery(ctx *pulumi.Context,
	name string, args *QueryArgs, opts ...pulumi.ResourceOption) (*Query, error) {
	if args == nil || args.Body == nil {
		return nil, errors.New("missing required argument 'Body'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Id == nil {
		return nil, errors.New("missing required argument 'Id'")
	}
	if args == nil || args.QueryPackName == nil {
		return nil, errors.New("missing required argument 'QueryPackName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &QueryArgs{}
	}
	var resource Query
	err := ctx.RegisterResource("azurerm:insights/v20190901preview:Query", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuery gets an existing Query resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryState, opts ...pulumi.ResourceOption) (*Query, error) {
	var resource Query
	err := ctx.ReadResource("azurerm:insights/v20190901preview:Query", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Query resources.
type queryState struct {
	// Object Id of user creating the query.
	Author *string `pulumi:"author"`
	// Body of the query.
	Body *string `pulumi:"body"`
	// Description of the query.
	Description *string `pulumi:"description"`
	// Unique display name for your query within the Query Pack.
	DisplayName *string `pulumi:"displayName"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Additional properties that can be set for the query.
	Properties map[string]interface{} `pulumi:"properties"`
	// The related metadata items for the function.
	Related *LogAnalyticsQueryPackQueryPropertiesResponseRelated `pulumi:"related"`
	// Read only system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Tags associated with the query.
	Tags map[string][]string `pulumi:"tags"`
	// Creation Date for the Log Analytics Query, in ISO 8601 format.
	TimeCreated *string `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics Query, in ISO 8601 format.
	TimeModified *string `pulumi:"timeModified"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type QueryState struct {
	// Object Id of user creating the query.
	Author pulumi.StringPtrInput
	// Body of the query.
	Body pulumi.StringPtrInput
	// Description of the query.
	Description pulumi.StringPtrInput
	// Unique display name for your query within the Query Pack.
	DisplayName pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Additional properties that can be set for the query.
	Properties pulumi.MapInput
	// The related metadata items for the function.
	Related LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrInput
	// Read only system data
	SystemData SystemDataResponsePtrInput
	// Tags associated with the query.
	Tags pulumi.StringArrayMapInput
	// Creation Date for the Log Analytics Query, in ISO 8601 format.
	TimeCreated pulumi.StringPtrInput
	// Last modified date of the Log Analytics Query, in ISO 8601 format.
	TimeModified pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (QueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryState)(nil)).Elem()
}

type queryArgs struct {
	// Body of the query.
	Body string `pulumi:"body"`
	// Description of the query.
	Description *string `pulumi:"description"`
	// Unique display name for your query within the Query Pack.
	DisplayName string `pulumi:"displayName"`
	// The id of a specific query defined in the Log Analytics QueryPack
	Id string `pulumi:"id"`
	// Additional properties that can be set for the query.
	Properties map[string]interface{} `pulumi:"properties"`
	// The name of the Log Analytics QueryPack resource.
	QueryPackName string `pulumi:"queryPackName"`
	// The related metadata items for the function.
	Related *LogAnalyticsQueryPackQueryPropertiesRelated `pulumi:"related"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags associated with the query.
	Tags map[string][]string `pulumi:"tags"`
}

// The set of arguments for constructing a Query resource.
type QueryArgs struct {
	// Body of the query.
	Body pulumi.StringInput
	// Description of the query.
	Description pulumi.StringPtrInput
	// Unique display name for your query within the Query Pack.
	DisplayName pulumi.StringInput
	// The id of a specific query defined in the Log Analytics QueryPack
	Id pulumi.StringInput
	// Additional properties that can be set for the query.
	Properties pulumi.MapInput
	// The name of the Log Analytics QueryPack resource.
	QueryPackName pulumi.StringInput
	// The related metadata items for the function.
	Related LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags associated with the query.
	Tags pulumi.StringArrayMapInput
}

func (QueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryArgs)(nil)).Elem()
}
