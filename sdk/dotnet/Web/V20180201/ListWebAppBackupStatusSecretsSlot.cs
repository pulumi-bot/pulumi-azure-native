// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201
{
    public static class ListWebAppBackupStatusSecretsSlot
    {
        public static Task<ListWebAppBackupStatusSecretsSlotResult> InvokeAsync(ListWebAppBackupStatusSecretsSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWebAppBackupStatusSecretsSlotResult>("azurerm:web/v20180201:listWebAppBackupStatusSecretsSlot", args ?? new ListWebAppBackupStatusSecretsSlotArgs(), options.WithVersion());
    }


    public sealed class ListWebAppBackupStatusSecretsSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the backup.
        /// </summary>
        [Input("backupName")]
        public string? BackupName { get; set; }

        /// <summary>
        /// Schedule for the backup if it is executed periodically.
        /// </summary>
        [Input("backupSchedule")]
        public Inputs.BackupScheduleArgs? BackupSchedule { get; set; }

        [Input("databases")]
        private List<Inputs.DatabaseBackupSettingArgs>? _databases;

        /// <summary>
        /// Databases included in the backup.
        /// </summary>
        public List<Inputs.DatabaseBackupSettingArgs> Databases
        {
            get => _databases ?? (_databases = new List<Inputs.DatabaseBackupSettingArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// True if the backup schedule is enabled (must be included in that case), false if the backup schedule should be disabled.
        /// </summary>
        [Input("enabled")]
        public bool? Enabled { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public string? Kind { get; set; }

        /// <summary>
        /// ID of backup.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of web app slot. If not specified then will default to production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        /// <summary>
        /// SAS URL to the container.
        /// </summary>
        [Input("storageAccountUrl", required: true)]
        public string StorageAccountUrl { get; set; } = null!;

        public ListWebAppBackupStatusSecretsSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWebAppBackupStatusSecretsSlotResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// BackupItem resource specific properties
        /// </summary>
        public readonly Outputs.BackupItemResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppBackupStatusSecretsSlotResult(
            string? kind,

            string name,

            Outputs.BackupItemResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
