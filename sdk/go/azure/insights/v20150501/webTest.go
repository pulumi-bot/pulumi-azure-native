// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Application Insights web test definition.
type WebTest struct {
	pulumi.CustomResourceState

	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesResponseConfigurationPtrOutput `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationResponseArrayOutput `pulumi:"locations"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrOutput `pulumi:"retryEnabled"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringOutput `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The kind of web test this is, valid choices are ping and multistep.
	WebTestKind pulumi.StringOutput `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName pulumi.StringOutput `pulumi:"webTestName"`
}

// NewWebTest registers a new resource with the given unique name, arguments, and options.
func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SyntheticMonitorId == nil {
		return nil, errors.New("invalid value for required argument 'SyntheticMonitorId'")
	}
	if args.WebTestName == nil {
		return nil, errors.New("invalid value for required argument 'WebTestName'")
	}
	if args.Frequency == nil {
		args.Frequency = pulumi.IntPtr(300)
	}
	if args.Kind == nil {
		e := WebTestKind("ping")
		args.Kind = &e
	}
	if args.Timeout == nil {
		args.Timeout = pulumi.IntPtr(30)
	}
	if args.WebTestKind == "" {
		args.WebTestKind = WebTestKind("ping")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/latest:WebTest"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20201005preview:WebTest"),
		},
	})
	opts = append(opts, aliases)
	var resource WebTest
	err := ctx.RegisterResource("azure-nextgen:insights/v20150501:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTest gets an existing WebTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azure-nextgen:insights/v20150501:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTest resources.
type webTestState struct {
	// An XML configuration specification for a WebTest.
	Configuration *WebTestPropertiesResponseConfiguration `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int `pulumi:"frequency"`
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations []WebTestGeolocationResponse `pulumi:"locations"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId *string `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int `pulumi:"timeout"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// The kind of web test this is, valid choices are ping and multistep.
	WebTestKind *string `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName *string `pulumi:"webTestName"`
}

type WebTestState struct {
	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesResponseConfigurationPtrInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrInput
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationResponseArrayInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringPtrInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// The kind of web test this is, valid choices are ping and multistep.
	WebTestKind pulumi.StringPtrInput
	// User defined name if this WebTest.
	WebTestName pulumi.StringPtrInput
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	// An XML configuration specification for a WebTest.
	Configuration *WebTestPropertiesConfiguration `pulumi:"configuration"`
	// Purpose/user defined descriptive test for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int `pulumi:"frequency"`
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations []WebTestGeolocation `pulumi:"locations"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId string `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int `pulumi:"timeout"`
	// The kind of web test this is, valid choices are ping and multistep.
	WebTestKind string `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName string `pulumi:"webTestName"`
}

// The set of arguments for constructing a WebTest resource.
type WebTestArgs struct {
	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesConfigurationPtrInput
	// Purpose/user defined descriptive test for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrInput
	// The kind of web test that this web test watches. Choices are ping and multistep.
	Kind *WebTestKind
	// Resource location
	Location pulumi.StringPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrInput
	// The kind of web test this is, valid choices are ping and multistep.
	WebTestKind WebTestKind
	// User defined name if this WebTest.
	WebTestName pulumi.StringInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}

type WebTestInput interface {
	pulumi.Input

	ToWebTestOutput() WebTestOutput
	ToWebTestOutputWithContext(ctx context.Context) WebTestOutput
}

func (*WebTest) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTest)(nil))
}

func (i *WebTest) ToWebTestOutput() WebTestOutput {
	return i.ToWebTestOutputWithContext(context.Background())
}

func (i *WebTest) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestOutput)
}

type WebTestOutput struct {
	*pulumi.OutputState
}

func (WebTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTest)(nil))
}

func (o WebTestOutput) ToWebTestOutput() WebTestOutput {
	return o
}

func (o WebTestOutput) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebTestOutput{})
}
