// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401
{
    /// <summary>
    /// Class representing Traffic Manager User Metrics.
    /// 
    /// ## Example Usage
    /// ### TrafficManagerUserMetricsKeys-PUT
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var trafficManagerUserMetricsKey = new AzureRM.Network.V20180401.TrafficManagerUserMetricsKey("trafficManagerUserMetricsKey", new AzureRM.Network.V20180401.TrafficManagerUserMetricsKeyArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class TrafficManagerUserMetricsKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The key returned by the User Metrics operation.
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficManagerUserMetricsKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficManagerUserMetricsKey(string name, TrafficManagerUserMetricsKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180401:TrafficManagerUserMetricsKey", name, args ?? new TrafficManagerUserMetricsKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficManagerUserMetricsKey(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180401:TrafficManagerUserMetricsKey", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:TrafficManagerUserMetricsKey"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrafficManagerUserMetricsKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficManagerUserMetricsKey Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TrafficManagerUserMetricsKey(name, id, options);
        }
    }

    public sealed class TrafficManagerUserMetricsKeyArgs : Pulumi.ResourceArgs
    {
        public TrafficManagerUserMetricsKeyArgs()
        {
        }
    }
}
