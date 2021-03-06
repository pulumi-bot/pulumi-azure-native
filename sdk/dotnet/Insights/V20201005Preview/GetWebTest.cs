// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20201005Preview
{
    public static class GetWebTest
    {
        /// <summary>
        /// An Application Insights WebTest definition.
        /// </summary>
        public static Task<GetWebTestResult> InvokeAsync(GetWebTestArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebTestResult>("azure-native:insights/v20201005preview:getWebTest", args ?? new GetWebTestArgs(), options.WithVersion());
    }


    public sealed class GetWebTestArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights WebTest resource.
        /// </summary>
        [Input("webTestName", required: true)]
        public string WebTestName { get; set; } = null!;

        public GetWebTestArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebTestResult
    {
        /// <summary>
        /// An XML configuration specification for a WebTest.
        /// </summary>
        public readonly Outputs.WebTestPropertiesResponseConfiguration? Configuration;
        /// <summary>
        /// The collection of content validation properties
        /// </summary>
        public readonly Outputs.WebTestPropertiesResponseContentValidation? ContentValidation;
        /// <summary>
        /// User defined description for this WebTest.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Is the test actively being monitored.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Validate that the WebTest returns the http status code provided.
        /// </summary>
        public readonly int? ExpectedHttpStatusCode;
        /// <summary>
        /// Interval in seconds between test runs for this WebTest. Default value is 300.
        /// </summary>
        public readonly int? Frequency;
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// When set, validation will ignore the status code.
        /// </summary>
        public readonly bool? IgnoreHttpsStatusCode;
        /// <summary>
        /// The kind of WebTest that this web test watches. Choices are ping and multistep.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// A list of where to physically run the tests from to give global coverage for accessibility of your application.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebTestGeolocationResponse> Locations;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The collection of request properties
        /// </summary>
        public readonly Outputs.WebTestPropertiesResponseRequest? Request;
        /// <summary>
        /// Allow for retries should this WebTest fail.
        /// </summary>
        public readonly bool? RetryEnabled;
        /// <summary>
        /// A number of days to check still remain before the the existing SSL cert expires.
        /// </summary>
        public readonly int? SSLCertRemainingLifetimeCheck;
        /// <summary>
        /// Checks to see if the SSL cert is still valid.
        /// </summary>
        public readonly bool? SSLCheck;
        /// <summary>
        /// Unique ID of this WebTest. This is typically the same value as the Name field.
        /// </summary>
        public readonly string SyntheticMonitorId;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Seconds until this WebTest will timeout and fail. Default value is 30.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The kind of web test this is, valid choices are ping, multistep, basic, and standard.
        /// </summary>
        public readonly string WebTestKind;
        /// <summary>
        /// User defined name if this WebTest.
        /// </summary>
        public readonly string WebTestName;

        [OutputConstructor]
        private GetWebTestResult(
            Outputs.WebTestPropertiesResponseConfiguration? configuration,

            Outputs.WebTestPropertiesResponseContentValidation? contentValidation,

            string? description,

            bool? enabled,

            int? expectedHttpStatusCode,

            int? frequency,

            string id,

            bool? ignoreHttpsStatusCode,

            string? kind,

            string location,

            ImmutableArray<Outputs.WebTestGeolocationResponse> locations,

            string name,

            string provisioningState,

            Outputs.WebTestPropertiesResponseRequest? request,

            bool? retryEnabled,

            int? sSLCertRemainingLifetimeCheck,

            bool? sSLCheck,

            string syntheticMonitorId,

            ImmutableDictionary<string, string>? tags,

            int? timeout,

            string type,

            string webTestKind,

            string webTestName)
        {
            Configuration = configuration;
            ContentValidation = contentValidation;
            Description = description;
            Enabled = enabled;
            ExpectedHttpStatusCode = expectedHttpStatusCode;
            Frequency = frequency;
            Id = id;
            IgnoreHttpsStatusCode = ignoreHttpsStatusCode;
            Kind = kind;
            Location = location;
            Locations = locations;
            Name = name;
            ProvisioningState = provisioningState;
            Request = request;
            RetryEnabled = retryEnabled;
            SSLCertRemainingLifetimeCheck = sSLCertRemainingLifetimeCheck;
            SSLCheck = sSLCheck;
            SyntheticMonitorId = syntheticMonitorId;
            Tags = tags;
            Timeout = timeout;
            Type = type;
            WebTestKind = webTestKind;
            WebTestName = webTestName;
        }
    }
}
