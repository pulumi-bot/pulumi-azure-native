// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190801
{
    /// <summary>
    /// Represents a share on the  Data Box Edge/Gateway device.
    /// </summary>
    public partial class Share : Pulumi.CustomResource
    {
        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The share properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SharePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Share resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Share(string name, ShareArgs args, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190801:Share", name, args ?? new ShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Share(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190801:Share", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Share resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Share Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Share(name, id, options);
        }
    }

    public sealed class ShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access protocol to be used by the share.
        /// </summary>
        [Input("accessProtocol", required: true)]
        public Input<string> AccessProtocol { get; set; } = null!;

        /// <summary>
        /// Azure container mapping for the share.
        /// </summary>
        [Input("azureContainerInfo")]
        public Input<Inputs.AzureContainerInfoArgs>? AzureContainerInfo { get; set; }

        [Input("clientAccessRights")]
        private InputList<Inputs.ClientAccessRightArgs>? _clientAccessRights;

        /// <summary>
        /// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        /// </summary>
        public InputList<Inputs.ClientAccessRightArgs> ClientAccessRights
        {
            get => _clientAccessRights ?? (_clientAccessRights = new InputList<Inputs.ClientAccessRightArgs>());
            set => _clientAccessRights = value;
        }

        /// <summary>
        /// Data policy of the share.
        /// </summary>
        [Input("dataPolicy")]
        public Input<string>? DataPolicy { get; set; }

        /// <summary>
        /// Description for the share.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// Current monitoring status of the share.
        /// </summary>
        [Input("monitoringStatus", required: true)]
        public Input<string> MonitoringStatus { get; set; } = null!;

        /// <summary>
        /// The share name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Details of the refresh job on this share.
        /// </summary>
        [Input("refreshDetails")]
        public Input<Inputs.RefreshDetailsArgs>? RefreshDetails { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Current status of the share.
        /// </summary>
        [Input("shareStatus", required: true)]
        public Input<string> ShareStatus { get; set; } = null!;

        [Input("userAccessRights")]
        private InputList<Inputs.UserAccessRightArgs>? _userAccessRights;

        /// <summary>
        /// Mapping of users and corresponding access rights on the share (required for SMB protocol).
        /// </summary>
        public InputList<Inputs.UserAccessRightArgs> UserAccessRights
        {
            get => _userAccessRights ?? (_userAccessRights = new InputList<Inputs.UserAccessRightArgs>());
            set => _userAccessRights = value;
        }

        public ShareArgs()
        {
        }
    }
}
