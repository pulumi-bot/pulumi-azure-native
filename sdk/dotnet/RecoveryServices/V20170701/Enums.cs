// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.RecoveryServices.V20170701
{
    /// <summary>
    /// Type of backup management for the backed up item.
    /// </summary>
    [EnumType]
    public readonly struct BackupManagementType : IEquatable<BackupManagementType>
    {
        private readonly string _value;

        private BackupManagementType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BackupManagementType Invalid { get; } = new BackupManagementType("Invalid");
        public static BackupManagementType AzureIaasVM { get; } = new BackupManagementType("AzureIaasVM");
        public static BackupManagementType MAB { get; } = new BackupManagementType("MAB");
        public static BackupManagementType DPM { get; } = new BackupManagementType("DPM");
        public static BackupManagementType AzureBackupServer { get; } = new BackupManagementType("AzureBackupServer");
        public static BackupManagementType AzureSql { get; } = new BackupManagementType("AzureSql");
        public static BackupManagementType AzureStorage { get; } = new BackupManagementType("AzureStorage");
        public static BackupManagementType AzureWorkload { get; } = new BackupManagementType("AzureWorkload");
        public static BackupManagementType DefaultBackup { get; } = new BackupManagementType("DefaultBackup");

        public static bool operator ==(BackupManagementType left, BackupManagementType right) => left.Equals(right);
        public static bool operator !=(BackupManagementType left, BackupManagementType right) => !left.Equals(right);

        public static explicit operator string(BackupManagementType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BackupManagementType other && Equals(other);
        public bool Equals(BackupManagementType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Backup state of this backup item.
    /// </summary>
    [EnumType]
    public readonly struct ProtectionStatus : IEquatable<ProtectionStatus>
    {
        private readonly string _value;

        private ProtectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProtectionStatus Invalid { get; } = new ProtectionStatus("Invalid");
        public static ProtectionStatus NotProtected { get; } = new ProtectionStatus("NotProtected");
        public static ProtectionStatus Protecting { get; } = new ProtectionStatus("Protecting");
        public static ProtectionStatus Protected { get; } = new ProtectionStatus("Protected");
        public static ProtectionStatus ProtectionFailed { get; } = new ProtectionStatus("ProtectionFailed");

        public static bool operator ==(ProtectionStatus left, ProtectionStatus right) => left.Equals(right);
        public static bool operator !=(ProtectionStatus left, ProtectionStatus right) => !left.Equals(right);

        public static explicit operator string(ProtectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProtectionStatus other && Equals(other);
        public bool Equals(ProtectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Workload item type of the item for which intent is to be set
    /// </summary>
    [EnumType]
    public readonly struct WorkloadItemType : IEquatable<WorkloadItemType>
    {
        private readonly string _value;

        private WorkloadItemType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkloadItemType Invalid { get; } = new WorkloadItemType("Invalid");
        public static WorkloadItemType SQLInstance { get; } = new WorkloadItemType("SQLInstance");
        public static WorkloadItemType SQLDataBase { get; } = new WorkloadItemType("SQLDataBase");
        public static WorkloadItemType SAPHanaSystem { get; } = new WorkloadItemType("SAPHanaSystem");
        public static WorkloadItemType SAPHanaDatabase { get; } = new WorkloadItemType("SAPHanaDatabase");
        public static WorkloadItemType SAPAseSystem { get; } = new WorkloadItemType("SAPAseSystem");
        public static WorkloadItemType SAPAseDatabase { get; } = new WorkloadItemType("SAPAseDatabase");

        public static bool operator ==(WorkloadItemType left, WorkloadItemType right) => left.Equals(right);
        public static bool operator !=(WorkloadItemType left, WorkloadItemType right) => !left.Equals(right);

        public static explicit operator string(WorkloadItemType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkloadItemType other && Equals(other);
        public bool Equals(WorkloadItemType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
