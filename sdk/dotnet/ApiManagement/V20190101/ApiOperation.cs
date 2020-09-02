// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101
{
    /// <summary>
    /// Api Operation details.
    /// </summary>
    public partial class ApiOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the operation. May include HTML formatting tags.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Operation Name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
        /// </summary>
        [Output("method")]
        public Output<string> Method { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Operation Policies
        /// </summary>
        [Output("policies")]
        public Output<string?> Policies { get; private set; } = null!;

        /// <summary>
        /// An entity containing request details.
        /// </summary>
        [Output("request")]
        public Output<Outputs.RequestContractResponseResult?> Request { get; private set; } = null!;

        /// <summary>
        /// Array of Operation responses.
        /// </summary>
        [Output("responses")]
        public Output<ImmutableArray<Outputs.ResponseContractResponseResult>> Responses { get; private set; } = null!;

        /// <summary>
        /// Collection of URL template parameters.
        /// </summary>
        [Output("templateParameters")]
        public Output<ImmutableArray<Outputs.ParameterContractResponseResult>> TemplateParameters { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
        /// </summary>
        [Output("urlTemplate")]
        public Output<string> UrlTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a ApiOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiOperation(string name, ApiOperationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:ApiOperation", name, args ?? new ApiOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiOperation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20190101:ApiOperation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/latest:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20160707:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20161010:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20170301:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180101:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180601preview:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:ApiOperation"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201preview:ApiOperation"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiOperation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiOperation(name, id, options);
        }
    }

    public sealed class ApiOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Description of the operation. May include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Operation Name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A Valid HTTP Operation Method. Typical Http Methods like GET, PUT, POST but not limited by only them.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// Operation identifier within an API. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationId", required: true)]
        public Input<string> OperationId { get; set; } = null!;

        /// <summary>
        /// Operation Policies
        /// </summary>
        [Input("policies")]
        public Input<string>? Policies { get; set; }

        /// <summary>
        /// An entity containing request details.
        /// </summary>
        [Input("request")]
        public Input<Inputs.RequestContractArgs>? Request { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("responses")]
        private InputList<Inputs.ResponseContractArgs>? _responses;

        /// <summary>
        /// Array of Operation responses.
        /// </summary>
        public InputList<Inputs.ResponseContractArgs> Responses
        {
            get => _responses ?? (_responses = new InputList<Inputs.ResponseContractArgs>());
            set => _responses = value;
        }

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("templateParameters")]
        private InputList<Inputs.ParameterContractArgs>? _templateParameters;

        /// <summary>
        /// Collection of URL template parameters.
        /// </summary>
        public InputList<Inputs.ParameterContractArgs> TemplateParameters
        {
            get => _templateParameters ?? (_templateParameters = new InputList<Inputs.ParameterContractArgs>());
            set => _templateParameters = value;
        }

        /// <summary>
        /// Relative URL template identifying the target resource for this operation. May include parameters. Example: /customers/{cid}/orders/{oid}/?date={date}
        /// </summary>
        [Input("urlTemplate", required: true)]
        public Input<string> UrlTemplate { get; set; } = null!;

        public ApiOperationArgs()
        {
        }
    }
}
