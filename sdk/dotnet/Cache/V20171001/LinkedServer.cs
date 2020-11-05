// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cache.V20171001
{
    /// <summary>
    /// Response to put/get linked server (with properties) for Redis cache.
    /// </summary>
    public partial class LinkedServer : Pulumi.CustomResource
    {
        /// <summary>
        /// Fully qualified resourceId of the linked redis cache.
        /// </summary>
        [Output("linkedRedisCacheId")]
        public Output<string> LinkedRedisCacheId { get; private set; } = null!;

        /// <summary>
        /// Location of the linked redis cache.
        /// </summary>
        [Output("linkedRedisCacheLocation")]
        public Output<string> LinkedRedisCacheLocation { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Terminal state of the link between primary and secondary redis cache.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Role of the linked server.
        /// </summary>
        [Output("serverRole")]
        public Output<string> ServerRole { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LinkedServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LinkedServer(string name, LinkedServerArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:cache/v20171001:LinkedServer", name, args ?? new LinkedServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LinkedServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:cache/v20171001:LinkedServer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:cache/latest:LinkedServer"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20170201:LinkedServer"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20180301:LinkedServer"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20190701:LinkedServer"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20200601:LinkedServer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LinkedServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LinkedServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LinkedServer(name, id, options);
        }
    }

    public sealed class LinkedServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified resourceId of the linked redis cache.
        /// </summary>
        [Input("linkedRedisCacheId", required: true)]
        public Input<string> LinkedRedisCacheId { get; set; } = null!;

        /// <summary>
        /// Location of the linked redis cache.
        /// </summary>
        [Input("linkedRedisCacheLocation", required: true)]
        public Input<string> LinkedRedisCacheLocation { get; set; } = null!;

        /// <summary>
        /// The name of the linked server that is being added to the Redis cache.
        /// </summary>
        [Input("linkedServerName", required: true)]
        public Input<string> LinkedServerName { get; set; } = null!;

        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Role of the linked server.
        /// </summary>
        [Input("serverRole", required: true)]
        public Input<string> ServerRole { get; set; } = null!;

        public LinkedServerArgs()
        {
        }
    }
}
