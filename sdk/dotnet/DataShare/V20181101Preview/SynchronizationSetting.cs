// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.V20181101Preview
{
    /// <summary>
    /// A Synchronization Setting data transfer object.
    /// </summary>
    public partial class SynchronizationSetting : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of synchronization
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SynchronizationSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SynchronizationSetting(string name, SynchronizationSettingArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datashare/v20181101preview:SynchronizationSetting", name, args ?? new SynchronizationSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SynchronizationSetting(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datashare/v20181101preview:SynchronizationSetting", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datashare/latest:SynchronizationSetting"},
                    new Pulumi.Alias { Type = "azurerm:datashare/v20191101:SynchronizationSetting"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SynchronizationSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SynchronizationSetting Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SynchronizationSetting(name, id, options);
        }
    }

    public sealed class SynchronizationSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Kind of synchronization
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share to add the synchronization setting to.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        /// <summary>
        /// The name of the synchronizationSetting.
        /// </summary>
        [Input("synchronizationSettingName", required: true)]
        public Input<string> SynchronizationSettingName { get; set; } = null!;

        public SynchronizationSettingArgs()
        {
        }
    }
}
