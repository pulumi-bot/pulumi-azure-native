// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20200601Preview
{
    /// <summary>
    /// Api details.
    /// </summary>
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes the Revision of the Api. If no value is provided, default revision 1 is created
        /// </summary>
        [Output("apiRevision")]
        public Output<string?> ApiRevision { get; private set; } = null!;

        /// <summary>
        /// Description of the Api Revision.
        /// </summary>
        [Output("apiRevisionDescription")]
        public Output<string?> ApiRevisionDescription { get; private set; } = null!;

        /// <summary>
        /// Type of API.
        /// </summary>
        [Output("apiType")]
        public Output<string?> ApiType { get; private set; } = null!;

        /// <summary>
        /// Indicates the Version identifier of the API if the API is versioned
        /// </summary>
        [Output("apiVersion")]
        public Output<string?> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the Api Version.
        /// </summary>
        [Output("apiVersionDescription")]
        public Output<string?> ApiVersionDescription { get; private set; } = null!;

        /// <summary>
        /// Version set details
        /// </summary>
        [Output("apiVersionSet")]
        public Output<Outputs.ApiVersionSetContractDetailsResponse?> ApiVersionSet { get; private set; } = null!;

        /// <summary>
        /// A resource identifier for the related ApiVersionSet.
        /// </summary>
        [Output("apiVersionSetId")]
        public Output<string?> ApiVersionSetId { get; private set; } = null!;

        /// <summary>
        /// Collection of authentication settings included into this API.
        /// </summary>
        [Output("authenticationSettings")]
        public Output<Outputs.AuthenticationSettingsContractResponse?> AuthenticationSettings { get; private set; } = null!;

        /// <summary>
        /// Description of the API. May include HTML formatting tags.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// API name. Must be 1 to 300 characters long.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Indicates if API revision is current api revision.
        /// </summary>
        [Output("isCurrent")]
        public Output<bool?> IsCurrent { get; private set; } = null!;

        /// <summary>
        /// Indicates if API revision is accessible via the gateway.
        /// </summary>
        [Output("isOnline")]
        public Output<bool> IsOnline { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Describes on which protocols the operations in this API can be invoked.
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
        /// </summary>
        [Output("serviceUrl")]
        public Output<string?> ServiceUrl { get; private set; } = null!;

        /// <summary>
        /// API identifier of the source API.
        /// </summary>
        [Output("sourceApiId")]
        public Output<string?> SourceApiId { get; private set; } = null!;

        /// <summary>
        /// Protocols over which API is made available.
        /// </summary>
        [Output("subscriptionKeyParameterNames")]
        public Output<Outputs.SubscriptionKeyParameterNamesContractResponse?> SubscriptionKeyParameterNames { get; private set; } = null!;

        /// <summary>
        /// Specifies whether an API or Product subscription is required for accessing the API.
        /// </summary>
        [Output("subscriptionRequired")]
        public Output<bool?> SubscriptionRequired { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20200601preview:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20200601preview:Api", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/latest:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20160707:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20161010:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20170301:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180101:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180601preview:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20190101:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20191201:Api"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20191201preview:Api"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Api(name, id, options);
        }
    }

    public sealed class ApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Describes the Revision of the Api. If no value is provided, default revision 1 is created
        /// </summary>
        [Input("apiRevision")]
        public Input<string>? ApiRevision { get; set; }

        /// <summary>
        /// Description of the Api Revision.
        /// </summary>
        [Input("apiRevisionDescription")]
        public Input<string>? ApiRevisionDescription { get; set; }

        /// <summary>
        /// Type of API.
        /// </summary>
        [Input("apiType")]
        public InputUnion<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.ApiType>? ApiType { get; set; }

        /// <summary>
        /// Indicates the Version identifier of the API if the API is versioned
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Description of the Api Version.
        /// </summary>
        [Input("apiVersionDescription")]
        public Input<string>? ApiVersionDescription { get; set; }

        /// <summary>
        /// Version set details
        /// </summary>
        [Input("apiVersionSet")]
        public Input<Inputs.ApiVersionSetContractDetailsArgs>? ApiVersionSet { get; set; }

        /// <summary>
        /// A resource identifier for the related ApiVersionSet.
        /// </summary>
        [Input("apiVersionSetId")]
        public Input<string>? ApiVersionSetId { get; set; }

        /// <summary>
        /// Collection of authentication settings included into this API.
        /// </summary>
        [Input("authenticationSettings")]
        public Input<Inputs.AuthenticationSettingsContractArgs>? AuthenticationSettings { get; set; }

        /// <summary>
        /// Description of the API. May include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// API name. Must be 1 to 300 characters long.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Format of the Content in which the API is getting imported.
        /// </summary>
        [Input("format")]
        public InputUnion<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.ContentFormat>? Format { get; set; }

        /// <summary>
        /// Indicates if API revision is current api revision.
        /// </summary>
        [Input("isCurrent")]
        public Input<bool>? IsCurrent { get; set; }

        /// <summary>
        /// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("protocols")]
        private InputList<Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Protocol>? _protocols;

        /// <summary>
        /// Describes on which protocols the operations in this API can be invoked.
        /// </summary>
        public InputList<Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Protocol> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Protocol>());
            set => _protocols = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
        /// </summary>
        [Input("serviceUrl")]
        public Input<string>? ServiceUrl { get; set; }

        /// <summary>
        /// Type of Api to create. 
        ///  * `http` creates a SOAP to REST API 
        ///  * `soap` creates a SOAP pass-through API .
        /// </summary>
        [Input("soapApiType")]
        public InputUnion<string, Pulumi.AzureNextGen.ApiManagement.V20200601Preview.SoapApiType>? SoapApiType { get; set; }

        /// <summary>
        /// API identifier of the source API.
        /// </summary>
        [Input("sourceApiId")]
        public Input<string>? SourceApiId { get; set; }

        /// <summary>
        /// Protocols over which API is made available.
        /// </summary>
        [Input("subscriptionKeyParameterNames")]
        public Input<Inputs.SubscriptionKeyParameterNamesContractArgs>? SubscriptionKeyParameterNames { get; set; }

        /// <summary>
        /// Specifies whether an API or Product subscription is required for accessing the API.
        /// </summary>
        [Input("subscriptionRequired")]
        public Input<bool>? SubscriptionRequired { get; set; }

        /// <summary>
        /// Content value when Importing an API.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Criteria to limit import of WSDL to a subset of the document.
        /// </summary>
        [Input("wsdlSelector")]
        public Input<Inputs.ApiCreateOrUpdatePropertiesWsdlSelectorArgs>? WsdlSelector { get; set; }

        public ApiArgs()
        {
        }
    }
}
