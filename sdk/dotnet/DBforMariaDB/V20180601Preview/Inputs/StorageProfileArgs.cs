// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBforMariaDB.V20180601Preview.Inputs
{

    /// <summary>
    /// Storage Profile properties of a server
    /// </summary>
    public sealed class StorageProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backup retention days for the server.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// Enable Geo-redundant or not for server backup.
        /// </summary>
        [Input("geoRedundantBackup")]
        public Input<string>? GeoRedundantBackup { get; set; }

        /// <summary>
        /// Enable Storage Auto Grow.
        /// </summary>
        [Input("storageAutogrow")]
        public Input<string>? StorageAutogrow { get; set; }

        /// <summary>
        /// Max storage allowed for a server.
        /// </summary>
        [Input("storageMB")]
        public Input<int>? StorageMB { get; set; }

        public StorageProfileArgs()
        {
        }
    }
}
