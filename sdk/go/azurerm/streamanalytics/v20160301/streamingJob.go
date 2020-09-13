// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A streaming job object, containing all information associated with the named streaming job.
//
// ## Example Usage
// ### Create a complete streaming job (a streaming job with a transformation, at least 1 input and at least 1 output)
//
// ```go
//
// ```
// ### Create a streaming job shell (a streaming job with no inputs, outputs, transformation, or functions)
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
// 		_, err := streamanalytics.NewStreamingJob(ctx, "streamingJob", &streamanalytics.StreamingJobArgs{
// 			CompatibilityLevel:                 pulumi.String("1.0"),
// 			DataLocale:                         pulumi.String("en-US"),
// 			EventsLateArrivalMaxDelayInSeconds: pulumi.Int(16),
// 			EventsOutOfOrderMaxDelayInSeconds:  pulumi.Int(5),
// 			EventsOutOfOrderPolicy:             pulumi.String("Drop"),
// 			Functions:                          streamanalytics.FunctionArray{},
// 			Inputs:                             streamanalytics.InputArray{},
// 			JobName:                            pulumi.String("sj59"),
// 			Location:                           pulumi.String("West US"),
// 			OutputErrorPolicy:                  pulumi.String("Drop"),
// 			Outputs:                            streamanalytics.OutputArray{},
// 			ResourceGroupName:                  pulumi.String("sjrg6936"),
// 			Sku: &streamanalytics.SkuArgs{
// 				Name: pulumi.String("Standard"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"key1":      pulumi.String("value1"),
// 				"key3":      pulumi.String("value3"),
// 				"randomKey": pulumi.String("randomValue"),
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
type StreamingJob struct {
	pulumi.CustomResourceState

	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrOutput `pulumi:"compatibilityLevel"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrOutput `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrOutput `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionResponseArrayOutput `pulumi:"functions"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputResponseArrayOutput `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState pulumi.StringOutput `pulumi:"jobState"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime pulumi.StringOutput `pulumi:"lastOutputEventTime"`
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrOutput `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrOutput `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrOutput `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputResponseArrayOutput `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationResponsePtrOutput `pulumi:"transformation"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingJob registers a new resource with the given unique name, arguments, and options.
func NewStreamingJob(ctx *pulumi.Context,
	name string, args *StreamingJobArgs, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	if args == nil || args.JobName == nil {
		return nil, errors.New("missing required argument 'JobName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StreamingJobArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:streamanalytics/latest:StreamingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingJob
	err := ctx.RegisterResource("azurerm:streamanalytics/v20160301:StreamingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingJob gets an existing StreamingJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingJobState, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	var resource StreamingJob
	err := ctx.ReadResource("azurerm:streamanalytics/v20160301:StreamingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingJob resources.
type streamingJobState struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate *string `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag *string `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionResponse `pulumi:"functions"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputResponse `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId *string `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState *string `pulumi:"jobState"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime *string `pulumi:"lastOutputEventTime"`
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputResponse `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *TransformationResponse `pulumi:"transformation"`
	// Resource type
	Type *string `pulumi:"type"`
}

type StreamingJobState struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrInput
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate pulumi.StringPtrInput
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrInput
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringPtrInput
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionResponseArrayInput
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputResponseArrayInput
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId pulumi.StringPtrInput
	// Describes the state of the streaming job.
	JobState pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime pulumi.StringPtrInput
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrInput
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrInput
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputResponseArrayInput
	// Describes the provisioning status of the streaming job.
	ProvisioningState pulumi.StringPtrInput
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (StreamingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobState)(nil)).Elem()
}

type streamingJobArgs struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionType `pulumi:"functions"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputType `pulumi:"inputs"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location *string `pulumi:"location"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputType `pulumi:"outputs"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *Transformation `pulumi:"transformation"`
}

// The set of arguments for constructing a StreamingJob resource.
type StreamingJobArgs struct {
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrInput
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrInput
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionTypeArrayInput
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputTypeArrayInput
	// The name of the streaming job.
	JobName pulumi.StringInput
	// Resource location. Required on PUT (CreateOrReplace) requests.
	Location pulumi.StringPtrInput
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrInput
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrInput
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputTypeArrayInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationPtrInput
}

func (StreamingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobArgs)(nil)).Elem()
}
