// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20190801Preview
{
    /// <summary>
    /// The top level Linked service resource container.
    /// </summary>
    public partial class LinkedService : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        /// </summary>
        [Output("writeAccessResourceId")]
        public Output<string?> WriteAccessResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedService(string name, LinkedServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20190801preview:LinkedService", name, args ?? new LinkedServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20190801preview:LinkedService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/latest:LinkedService"},
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20151101preview:LinkedService"},
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20200301preview:LinkedService"},
                    new Pulumi.Alias { Type = "azurerm:operationalinsights/v20200801:LinkedService"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LinkedService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkedService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LinkedService(name, id, options);
        }
    }

    public sealed class LinkedServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the linkedServices resource
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public Input<string> LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to get. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Name of the Log Analytics Workspace that will contain the linkedServices resource
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        /// <summary>
        /// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        /// </summary>
        [Input("writeAccessResourceId")]
        public Input<string>? WriteAccessResourceId { get; set; }

        public LinkedServiceArgs()
        {
        }
    }
}
