// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforPostgreSQL.V20201005PrivatePreview.Inputs
{

    /// <summary>
    /// Represents a server role group.
    /// </summary>
    public sealed class ServerRoleGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If high availability is enabled or not for the server.
        /// </summary>
        [Input("enableHa")]
        public Input<bool>? EnableHa { get; set; }

        /// <summary>
        /// The name of the server role group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role of servers in the server role group.
        /// </summary>
        [Input("role")]
        public InputUnion<string, Pulumi.AzureNextGen.DBforPostgreSQL.V20201005PrivatePreview.ServerRole>? Role { get; set; }

        /// <summary>
        /// The number of servers in the server role group.
        /// </summary>
        [Input("serverCount")]
        public Input<int>? ServerCount { get; set; }

        /// <summary>
        /// The edition of a server (default: GeneralPurpose).
        /// </summary>
        [Input("serverEdition")]
        public InputUnion<string, Pulumi.AzureNextGen.DBforPostgreSQL.V20201005PrivatePreview.ServerEdition>? ServerEdition { get; set; }

        /// <summary>
        /// The storage of a server in MB (max: 2097152 = 2TiB).
        /// </summary>
        [Input("storageQuotaInMb")]
        public Input<int>? StorageQuotaInMb { get; set; }

        /// <summary>
        /// The vCores count of a server (max: 64).
        /// </summary>
        [Input("vCores")]
        public Input<int>? VCores { get; set; }

        public ServerRoleGroupArgs()
        {
        }
    }
}
