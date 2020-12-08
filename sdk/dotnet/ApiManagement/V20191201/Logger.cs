// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20191201
{
    /// <summary>
    /// Logger details.
    /// </summary>
    public partial class Logger : Pulumi.CustomResource
    {
        /// <summary>
        /// The name and SendRule connection string of the event hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableDictionary<string, string>> Credentials { get; private set; } = null!;

        /// <summary>
        /// Logger description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether records are buffered in the logger before publishing. Default is assumed to be true.
        /// </summary>
        [Output("isBuffered")]
        public Output<bool?> IsBuffered { get; private set; } = null!;

        /// <summary>
        /// Logger type.
        /// </summary>
        [Output("loggerType")]
        public Output<string> LoggerType { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Logger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Logger(string name, LoggerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20191201:Logger", name, args ?? new LoggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Logger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:apimanagement/v20191201:Logger", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/latest:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20160707:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20161010:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20170301:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180101:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20180601preview:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20190101:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20191201preview:Logger"},
                    new Pulumi.Alias { Type = "azure-nextgen:apimanagement/v20200601preview:Logger"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Logger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Logger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Logger(name, id, options);
        }
    }

    public sealed class LoggerArgs : Pulumi.ResourceArgs
    {
        [Input("credentials", required: true)]
        private InputMap<string>? _credentials;

        /// <summary>
        /// The name and SendRule connection string of the event hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.
        /// </summary>
        public InputMap<string> Credentials
        {
            get => _credentials ?? (_credentials = new InputMap<string>());
            set => _credentials = value;
        }

        /// <summary>
        /// Logger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether records are buffered in the logger before publishing. Default is assumed to be true.
        /// </summary>
        [Input("isBuffered")]
        public Input<bool>? IsBuffered { get; set; }

        /// <summary>
        /// Logger identifier. Must be unique in the API Management service instance.
        /// </summary>
        [Input("loggerId", required: true)]
        public Input<string> LoggerId { get; set; } = null!;

        /// <summary>
        /// Logger type.
        /// </summary>
        [Input("loggerType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.ApiManagement.V20191201.LoggerType> LoggerType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public LoggerArgs()
        {
        }
    }
}
