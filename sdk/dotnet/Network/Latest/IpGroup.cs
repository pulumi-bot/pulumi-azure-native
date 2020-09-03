// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Latest
{
    /// <summary>
    /// The IpGroups resource information.
    /// 
    /// ## Example Usage
    /// ### CreateOrUpdate_IpGroups
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ipGroup = new AzureRM.Network.Latest.IpGroup("ipGroup", new AzureRM.Network.Latest.IpGroupArgs
    ///         {
    ///             IpAddresses = 
    ///             {
    ///                 "13.64.39.16/32",
    ///                 "40.74.146.80/31",
    ///                 "40.74.147.32/28",
    ///             },
    ///             IpGroupsName = "ipGroups1",
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///             Tags = 
    ///             {
    ///                 { "key1", "value1" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class IpGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// List of references to Azure resources that this IpGroups is associated with.
        /// </summary>
        [Output("firewalls")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> Firewalls { get; private set; } = null!;

        /// <summary>
        /// IpAddresses/IpAddressPrefixes in the IpGroups resource.
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the IpGroups resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IpGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpGroup(string name, IpGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/latest:IpGroup", name, args ?? new IpGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/latest:IpGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:IpGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:IpGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IpGroup(name, id, options);
        }
    }

    public sealed class IpGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// IpAddresses/IpAddressPrefixes in the IpGroups resource.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The name of the ipGroups.
        /// </summary>
        [Input("ipGroupsName", required: true)]
        public Input<string> IpGroupsName { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IpGroupArgs()
        {
        }
    }
}
