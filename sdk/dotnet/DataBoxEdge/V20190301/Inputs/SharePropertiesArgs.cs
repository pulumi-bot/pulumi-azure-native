// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190301.Inputs
{

    /// <summary>
    /// The share properties.
    /// </summary>
    public sealed class SharePropertiesArgs : Pulumi.ResourceArgs
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
        /// Current monitoring status of the share.
        /// </summary>
        [Input("monitoringStatus", required: true)]
        public Input<string> MonitoringStatus { get; set; } = null!;

        /// <summary>
        /// Details of the refresh job on this share.
        /// </summary>
        [Input("refreshDetails")]
        public Input<Inputs.RefreshDetailsArgs>? RefreshDetails { get; set; }

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

        public SharePropertiesArgs()
        {
        }
    }
}