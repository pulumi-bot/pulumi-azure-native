// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20200201
{
    /// <summary>
    /// The management group details.
    /// </summary>
    public partial class ManagementGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The generic properties of a management group.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ManagementGroupPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.  For example, Microsoft.Management/managementGroups
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagementGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagementGroup(string name, ManagementGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:management/v20200201:ManagementGroup", name, args ?? new ManagementGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagementGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:management/v20200201:ManagementGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ManagementGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagementGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagementGroup(name, id, options);
        }
    }

    public sealed class ManagementGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The details of a management group used during creation.
        /// </summary>
        [Input("details")]
        public Input<Inputs.CreateManagementGroupDetailsArgs>? Details { get; set; }

        /// <summary>
        /// The friendly name of the management group. If no value is passed then this  field will be set to the groupId.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Management Group ID.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ManagementGroupArgs()
        {
        }
    }
}
