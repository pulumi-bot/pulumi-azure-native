// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Logger details.
type Logger struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Logger entity contract properties.
	Properties LoggerContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLogger registers a new resource with the given unique name, arguments, and options.
func NewLogger(ctx *pulumi.Context,
	name string, args *LoggerArgs, opts ...pulumi.ResourceOption) (*Logger, error) {
	if args == nil || args.Credentials == nil {
		return nil, errors.New("missing required argument 'Credentials'")
	}
	if args == nil || args.LoggerType == nil {
		return nil, errors.New("missing required argument 'LoggerType'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &LoggerArgs{}
	}
	var resource Logger
	err := ctx.RegisterResource("azurerm:apimanagement/v20170301:Logger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogger gets an existing Logger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggerState, opts ...pulumi.ResourceOption) (*Logger, error) {
	var resource Logger
	err := ctx.ReadResource("azurerm:apimanagement/v20170301:Logger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Logger resources.
type loggerState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Logger entity contract properties.
	Properties *LoggerContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type LoggerState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Logger entity contract properties.
	Properties LoggerContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (LoggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerState)(nil)).Elem()
}

type loggerArgs struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials map[string]string `pulumi:"credentials"`
	// Logger description.
	Description *string `pulumi:"description"`
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered *bool `pulumi:"isBuffered"`
	// Logger type.
	LoggerType string `pulumi:"loggerType"`
	// Logger identifier. Must be unique in the API Management service instance.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling settings for an ApplicationInsights logger.
	Sampling *LoggerSamplingContract `pulumi:"sampling"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Logger resource.
type LoggerArgs struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials pulumi.StringMapInput
	// Logger description.
	Description pulumi.StringPtrInput
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered pulumi.BoolPtrInput
	// Logger type.
	LoggerType pulumi.StringInput
	// Logger identifier. Must be unique in the API Management service instance.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Sampling settings for an ApplicationInsights logger.
	Sampling LoggerSamplingContractPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (LoggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerArgs)(nil)).Elem()
}
