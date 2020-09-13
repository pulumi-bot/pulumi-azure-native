// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A function object, containing all information associated with the named function. All functions are contained under a streaming job.
//
// ## Example Usage
// ### Create a JavaScript function
//
// ```go
// package main
//
// import (
// 	streamanalytics "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/streamanalytics/v20160301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := streamanalytics.NewFunction(ctx, "function", &streamanalytics.FunctionArgs{
// 			FunctionName: pulumi.String("function8197"),
// 			JobName:      pulumi.String("sj8653"),
// 			Properties: &streamanalytics.FunctionPropertiesArgs{
// 				Binding: pulumi.Map{
// 					"properties": pulumi.StringMap{
// 						"script": pulumi.String("function (x, y) { return x + y; }"),
// 					},
// 					"type": pulumi.String("Microsoft.StreamAnalytics/JavascriptUdf"),
// 				},
// 				Inputs: pulumi.StringMapArray{
// 					pulumi.StringMap{
// 						"dataType": pulumi.String("Any"),
// 					},
// 				},
// 				Output: pulumi.StringMap{
// 					"dataType": pulumi.String("Any"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("sjrg1637"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create an Azure ML function
//
// ```go
// package main
//
// import (
// 	streamanalytics "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/streamanalytics/v20160301"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := streamanalytics.NewFunction(ctx, "function", &streamanalytics.FunctionArgs{
// 			FunctionName: pulumi.String("function588"),
// 			JobName:      pulumi.String("sj9093"),
// 			Properties: &streamanalytics.FunctionPropertiesArgs{
// 				Binding: pulumi.Map{
// 					"properties": pulumi.Map{
// 						"apiKey":    pulumi.String("someApiKey=="),
// 						"batchSize": pulumi.Float64(1000),
// 						"endpoint":  pulumi.String("someAzureMLEndpointURL"),
// 						"inputs": pulumi.Map{
// 							"columnNames": pulumi.MapArray{
// 								pulumi.Map{
// 									"dataType": pulumi.String("string"),
// 									"mapTo":    pulumi.Float64(0),
// 									"name":     pulumi.String("tweet"),
// 								},
// 							},
// 							"name": pulumi.String("input1"),
// 						},
// 						"outputs": pulumi.StringMapArray{
// 							pulumi.StringMap{
// 								"dataType": pulumi.String("string"),
// 								"name":     pulumi.String("Sentiment"),
// 							},
// 						},
// 					},
// 					"type": pulumi.String("Microsoft.MachineLearning/WebService"),
// 				},
// 				Inputs: pulumi.StringMapArray{
// 					pulumi.StringMap{
// 						"dataType": pulumi.String("nvarchar(max)"),
// 					},
// 				},
// 				Output: pulumi.StringMap{
// 					"dataType": pulumi.String("nvarchar(max)"),
// 				},
// 			},
// 			ResourceGroupName: pulumi.String("sjrg7"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Function struct {
	pulumi.CustomResourceState

	// Resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties that are associated with a function.
	Properties ScalarFunctionPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFunction registers a new resource with the given unique name, arguments, and options.
func NewFunction(ctx *pulumi.Context,
	name string, args *FunctionArgs, opts ...pulumi.ResourceOption) (*Function, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.JobName == nil {
		return nil, errors.New("missing required argument 'JobName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FunctionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:streamanalytics/latest:Function"),
		},
	})
	opts = append(opts, aliases)
	var resource Function
	err := ctx.RegisterResource("azurerm:streamanalytics/v20160301:Function", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunction gets an existing Function resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionState, opts ...pulumi.ResourceOption) (*Function, error) {
	var resource Function
	err := ctx.ReadResource("azurerm:streamanalytics/v20160301:Function", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Function resources.
type functionState struct {
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties *ScalarFunctionPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type FunctionState struct {
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with a function.
	Properties ScalarFunctionPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (FunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionState)(nil)).Elem()
}

type functionArgs struct {
	// The name of the function.
	FunctionName string `pulumi:"functionName"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with a function.
	Properties *ScalarFunctionProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Function resource.
type FunctionArgs struct {
	// The name of the function.
	FunctionName pulumi.StringInput
	// The name of the streaming job.
	JobName pulumi.StringInput
	// Resource name
	Name pulumi.StringPtrInput
	// The properties that are associated with a function.
	Properties ScalarFunctionPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (FunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionArgs)(nil)).Elem()
}
