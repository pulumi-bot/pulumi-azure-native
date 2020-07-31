// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801.Outputs
{

    [OutputType]
    public sealed class BackupRequestResponsePropertiesResult
    {
        /// <summary>
        /// Schedule for the backup if it is executed periodically
        /// </summary>
        public readonly Outputs.BackupScheduleResponseResult? BackupSchedule;
        /// <summary>
        /// Databases included in the backup
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseBackupSettingResponseResult> Databases;
        /// <summary>
        /// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Name of the backup
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// SAS URL to the container
        /// </summary>
        public readonly string? StorageAccountUrl;
        /// <summary>
        /// Type of the backup
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private BackupRequestResponsePropertiesResult(
            Outputs.BackupScheduleResponseResult? backupSchedule,

            ImmutableArray<Outputs.DatabaseBackupSettingResponseResult> databases,

            bool? enabled,

            string? name,

            string? storageAccountUrl,

            string type)
        {
            BackupSchedule = backupSchedule;
            Databases = databases;
            Enabled = enabled;
            Name = name;
            StorageAccountUrl = storageAccountUrl;
            Type = type;
        }
    }
}