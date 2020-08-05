// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301
{
    /// <summary>
    /// API details.
    /// </summary>
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Api entity contract properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApiContractPropertiesResponseResult> Properties { get; private set; } = null!;

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
            : base("azurerm:apimanagement/v20170301:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20170301:Api", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
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
        /// Describes the Revision of the Api. If no value is provided, default revision 1 is created
        /// </summary>
        [Input("apiRevision")]
        public Input<string>? ApiRevision { get; set; }

        /// <summary>
        /// Type of API.
        /// </summary>
        [Input("apiType")]
        public Input<string>? ApiType { get; set; }

        /// <summary>
        /// Indicates the Version identifier of the API if the API is versioned
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Api Version Set Contract details.
        /// </summary>
        [Input("apiVersionSet")]
        public Input<Inputs.ApiVersionSetContractArgs>? ApiVersionSet { get; set; }

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
        /// Format of the Content in which the API is getting imported.
        /// </summary>
        [Input("contentFormat")]
        public Input<string>? ContentFormat { get; set; }

        /// <summary>
        /// Content value when Importing an API.
        /// </summary>
        [Input("contentValue")]
        public Input<string>? ContentValue { get; set; }

        /// <summary>
        /// Description of the API. May include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// API name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// Describes on which protocols the operations in this API can be invoked.
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
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
        /// Absolute URL of the backend service implementing this API.
        /// </summary>
        [Input("serviceUrl")]
        public Input<string>? ServiceUrl { get; set; }

        /// <summary>
        /// Protocols over which API is made available.
        /// </summary>
        [Input("subscriptionKeyParameterNames")]
        public Input<Inputs.SubscriptionKeyParameterNamesContractArgs>? SubscriptionKeyParameterNames { get; set; }

        /// <summary>
        /// Criteria to limit import of WSDL to a subset of the document.
        /// </summary>
        [Input("wsdlSelector")]
        public Input<Inputs.ApiCreateOrUpdatePropertiesPropertiesArgs>? WsdlSelector { get; set; }

        public ApiArgs()
        {
        }
    }
}
