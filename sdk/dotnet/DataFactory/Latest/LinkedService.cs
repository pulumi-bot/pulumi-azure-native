// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.Latest
{
    /// <summary>
    /// Linked service resource type.
    /// </summary>
    public partial class LinkedService : Pulumi.CustomResource
    {
        /// <summary>
        /// Etag identifies change in the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of linked service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.LinkedServiceResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedService(string name, LinkedServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/latest:LinkedService", name, args ?? new LinkedServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datafactory/latest:LinkedService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datafactory/v20170901preview:LinkedService"},
                    new Pulumi.Alias { Type = "azurerm:datafactory/v20180601:LinkedService"},
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
        /// The factory name.
        /// </summary>
        [Input("factoryName", required: true)]
        public Input<string> FactoryName { get; set; } = null!;

        /// <summary>
        /// The linked service name.
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public Input<string> LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// Properties of linked service.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.LinkedServiceArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public LinkedServiceArgs()
        {
        }
    }
}
