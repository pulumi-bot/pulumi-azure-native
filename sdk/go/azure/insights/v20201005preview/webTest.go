// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201005preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Application Insights WebTest definition.
type WebTest struct {
	pulumi.CustomResourceState

	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesResponseConfigurationPtrOutput `pulumi:"configuration"`
	// The collection of content validation properties
	ContentValidation WebTestPropertiesResponseContentValidationPtrOutput `pulumi:"contentValidation"`
	// User defined description for this WebTest.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode pulumi.IntPtrOutput `pulumi:"expectedHttpStatusCode"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// When set, validation will ignore the status code.
	IgnoreHttpsStatusCode pulumi.BoolPtrOutput `pulumi:"ignoreHttpsStatusCode"`
	// The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationResponseArrayOutput `pulumi:"locations"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The collection of request properties
	Request WebTestPropertiesResponseRequestPtrOutput `pulumi:"request"`
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrOutput `pulumi:"retryEnabled"`
	// A number of days to check still remain before the the existing SSL cert expires.
	SSLCertRemainingLifetimeCheck pulumi.IntPtrOutput `pulumi:"sSLCertRemainingLifetimeCheck"`
	// Checks to see if the SSL cert is still valid.
	SSLCheck pulumi.BoolPtrOutput `pulumi:"sSLCheck"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringOutput `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
	WebTestKind pulumi.StringOutput `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName pulumi.StringOutput `pulumi:"webTestName"`
}

// NewWebTest registers a new resource with the given unique name, arguments, and options.
func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Locations == nil {
		return nil, errors.New("missing required argument 'Locations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SyntheticMonitorId == nil {
		return nil, errors.New("missing required argument 'SyntheticMonitorId'")
	}
	if args == nil || args.WebTestKind == nil {
		return nil, errors.New("missing required argument 'WebTestKind'")
	}
	if args == nil || args.WebTestName == nil {
		return nil, errors.New("missing required argument 'WebTestName'")
	}
	if args == nil {
		args = &WebTestArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:insights/latest:WebTest"),
		},
		{
			Type: pulumi.String("azure-nextgen:insights/v20150501:WebTest"),
		},
	})
	opts = append(opts, aliases)
	var resource WebTest
	err := ctx.RegisterResource("azure-nextgen:insights/v20201005preview:WebTest", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:insights/v20201005preview:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTest resources.
type webTestState struct {
	// An XML configuration specification for a WebTest.
	Configuration *WebTestPropertiesResponseConfiguration `pulumi:"configuration"`
	// The collection of content validation properties
	ContentValidation *WebTestPropertiesResponseContentValidation `pulumi:"contentValidation"`
	// User defined description for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode *int `pulumi:"expectedHttpStatusCode"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int `pulumi:"frequency"`
	// When set, validation will ignore the status code.
	IgnoreHttpsStatusCode *bool `pulumi:"ignoreHttpsStatusCode"`
	// The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations []WebTestGeolocationResponse `pulumi:"locations"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The collection of request properties
	Request *WebTestPropertiesResponseRequest `pulumi:"request"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// A number of days to check still remain before the the existing SSL cert expires.
	SSLCertRemainingLifetimeCheck *int `pulumi:"sSLCertRemainingLifetimeCheck"`
	// Checks to see if the SSL cert is still valid.
	SSLCheck *bool `pulumi:"sSLCheck"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId *string `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int `pulumi:"timeout"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
	WebTestKind *string `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName *string `pulumi:"webTestName"`
}

type WebTestState struct {
	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesResponseConfigurationPtrInput
	// The collection of content validation properties
	ContentValidation WebTestPropertiesResponseContentValidationPtrInput
	// User defined description for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode pulumi.IntPtrInput
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrInput
	// When set, validation will ignore the status code.
	IgnoreHttpsStatusCode pulumi.BoolPtrInput
	// The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationResponseArrayInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringPtrInput
	// The collection of request properties
	Request WebTestPropertiesResponseRequestPtrInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// A number of days to check still remain before the the existing SSL cert expires.
	SSLCertRemainingLifetimeCheck pulumi.IntPtrInput
	// Checks to see if the SSL cert is still valid.
	SSLCheck pulumi.BoolPtrInput
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
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
	// The collection of content validation properties
	ContentValidation *WebTestPropertiesContentValidation `pulumi:"contentValidation"`
	// User defined description for this WebTest.
	Description *string `pulumi:"description"`
	// Is the test actively being monitored.
	Enabled *bool `pulumi:"enabled"`
	// Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode *int `pulumi:"expectedHttpStatusCode"`
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int `pulumi:"frequency"`
	// When set, validation will ignore the status code.
	IgnoreHttpsStatusCode *bool `pulumi:"ignoreHttpsStatusCode"`
	// The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location string `pulumi:"location"`
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations []WebTestGeolocation `pulumi:"locations"`
	// The collection of request properties
	Request *WebTestPropertiesRequest `pulumi:"request"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Allow for retries should this WebTest fail.
	RetryEnabled *bool `pulumi:"retryEnabled"`
	// A number of days to check still remain before the the existing SSL cert expires.
	SSLCertRemainingLifetimeCheck *int `pulumi:"sSLCertRemainingLifetimeCheck"`
	// Checks to see if the SSL cert is still valid.
	SSLCheck *bool `pulumi:"sSLCheck"`
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId string `pulumi:"syntheticMonitorId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int `pulumi:"timeout"`
	// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
	WebTestKind string `pulumi:"webTestKind"`
	// User defined name if this WebTest.
	WebTestName string `pulumi:"webTestName"`
}

// The set of arguments for constructing a WebTest resource.
type WebTestArgs struct {
	// An XML configuration specification for a WebTest.
	Configuration WebTestPropertiesConfigurationPtrInput
	// The collection of content validation properties
	ContentValidation WebTestPropertiesContentValidationPtrInput
	// User defined description for this WebTest.
	Description pulumi.StringPtrInput
	// Is the test actively being monitored.
	Enabled pulumi.BoolPtrInput
	// Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode pulumi.IntPtrInput
	// Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency pulumi.IntPtrInput
	// When set, validation will ignore the status code.
	IgnoreHttpsStatusCode pulumi.BoolPtrInput
	// The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations WebTestGeolocationArrayInput
	// The collection of request properties
	Request WebTestPropertiesRequestPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Allow for retries should this WebTest fail.
	RetryEnabled pulumi.BoolPtrInput
	// A number of days to check still remain before the the existing SSL cert expires.
	SSLCertRemainingLifetimeCheck pulumi.IntPtrInput
	// Checks to see if the SSL cert is still valid.
	SSLCheck pulumi.BoolPtrInput
	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout pulumi.IntPtrInput
	// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
	WebTestKind pulumi.StringInput
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

func (WebTest) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTest)(nil)).Elem()
}

func (i WebTest) ToWebTestOutput() WebTestOutput {
	return i.ToWebTestOutputWithContext(context.Background())
}

func (i WebTest) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestOutput)
}

type WebTestOutput struct {
	*pulumi.OutputState
}

func (WebTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestOutput)(nil)).Elem()
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
