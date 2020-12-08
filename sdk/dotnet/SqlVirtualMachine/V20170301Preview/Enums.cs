// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.SqlVirtualMachine.V20170301Preview
{
    /// <summary>
    /// Backup schedule type.
    /// </summary>
    [EnumType]
    public readonly struct BackupScheduleType : IEquatable<BackupScheduleType>
    {
        private readonly string _value;

        private BackupScheduleType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BackupScheduleType Manual { get; } = new BackupScheduleType("Manual");
        public static BackupScheduleType Automated { get; } = new BackupScheduleType("Automated");

        public static bool operator ==(BackupScheduleType left, BackupScheduleType right) => left.Equals(right);
        public static bool operator !=(BackupScheduleType left, BackupScheduleType right) => !left.Equals(right);

        public static explicit operator string(BackupScheduleType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BackupScheduleType other && Equals(other);
        public bool Equals(BackupScheduleType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server connectivity option.
    /// </summary>
    [EnumType]
    public readonly struct ConnectivityType : IEquatable<ConnectivityType>
    {
        private readonly string _value;

        private ConnectivityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectivityType LOCAL { get; } = new ConnectivityType("LOCAL");
        public static ConnectivityType PRIVATE { get; } = new ConnectivityType("PRIVATE");
        public static ConnectivityType PUBLIC { get; } = new ConnectivityType("PUBLIC");

        public static bool operator ==(ConnectivityType left, ConnectivityType right) => left.Equals(right);
        public static bool operator !=(ConnectivityType left, ConnectivityType right) => !left.Equals(right);

        public static explicit operator string(ConnectivityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectivityType other && Equals(other);
        public bool Equals(ConnectivityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Day of week to apply the patch on.
    /// </summary>
    [EnumType]
    public readonly struct DayOfWeek : IEquatable<DayOfWeek>
    {
        private readonly string _value;

        private DayOfWeek(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DayOfWeek Monday { get; } = new DayOfWeek("Monday");
        public static DayOfWeek Tuesday { get; } = new DayOfWeek("Tuesday");
        public static DayOfWeek Wednesday { get; } = new DayOfWeek("Wednesday");
        public static DayOfWeek Thursday { get; } = new DayOfWeek("Thursday");
        public static DayOfWeek Friday { get; } = new DayOfWeek("Friday");
        public static DayOfWeek Saturday { get; } = new DayOfWeek("Saturday");
        public static DayOfWeek Sunday { get; } = new DayOfWeek("Sunday");

        public static bool operator ==(DayOfWeek left, DayOfWeek right) => left.Equals(right);
        public static bool operator !=(DayOfWeek left, DayOfWeek right) => !left.Equals(right);

        public static explicit operator string(DayOfWeek value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DayOfWeek other && Equals(other);
        public bool Equals(DayOfWeek other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Disk configuration to apply to SQL Server.
    /// </summary>
    [EnumType]
    public readonly struct DiskConfigurationType : IEquatable<DiskConfigurationType>
    {
        private readonly string _value;

        private DiskConfigurationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DiskConfigurationType NEW { get; } = new DiskConfigurationType("NEW");
        public static DiskConfigurationType EXTEND { get; } = new DiskConfigurationType("EXTEND");
        public static DiskConfigurationType ADD { get; } = new DiskConfigurationType("ADD");

        public static bool operator ==(DiskConfigurationType left, DiskConfigurationType right) => left.Equals(right);
        public static bool operator !=(DiskConfigurationType left, DiskConfigurationType right) => !left.Equals(right);

        public static explicit operator string(DiskConfigurationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiskConfigurationType other && Equals(other);
        public bool Equals(DiskConfigurationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Frequency of full backups. In both cases, full backups begin during the next scheduled time window.
    /// </summary>
    [EnumType]
    public readonly struct FullBackupFrequencyType : IEquatable<FullBackupFrequencyType>
    {
        private readonly string _value;

        private FullBackupFrequencyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FullBackupFrequencyType Daily { get; } = new FullBackupFrequencyType("Daily");
        public static FullBackupFrequencyType Weekly { get; } = new FullBackupFrequencyType("Weekly");

        public static bool operator ==(FullBackupFrequencyType left, FullBackupFrequencyType right) => left.Equals(right);
        public static bool operator !=(FullBackupFrequencyType left, FullBackupFrequencyType right) => !left.Equals(right);

        public static explicit operator string(FullBackupFrequencyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FullBackupFrequencyType other && Equals(other);
        public bool Equals(FullBackupFrequencyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
    /// </summary>
    [EnumType]
    public readonly struct IdentityType : IEquatable<IdentityType>
    {
        private readonly string _value;

        private IdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IdentityType SystemAssigned { get; } = new IdentityType("SystemAssigned");

        public static bool operator ==(IdentityType left, IdentityType right) => left.Equals(right);
        public static bool operator !=(IdentityType left, IdentityType right) => !left.Equals(right);

        public static explicit operator string(IdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IdentityType other && Equals(other);
        public bool Equals(IdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server edition type.
    /// </summary>
    [EnumType]
    public readonly struct SqlImageSku : IEquatable<SqlImageSku>
    {
        private readonly string _value;

        private SqlImageSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlImageSku Developer { get; } = new SqlImageSku("Developer");
        public static SqlImageSku Express { get; } = new SqlImageSku("Express");
        public static SqlImageSku Standard { get; } = new SqlImageSku("Standard");
        public static SqlImageSku Enterprise { get; } = new SqlImageSku("Enterprise");
        public static SqlImageSku Web { get; } = new SqlImageSku("Web");

        public static bool operator ==(SqlImageSku left, SqlImageSku right) => left.Equals(right);
        public static bool operator !=(SqlImageSku left, SqlImageSku right) => !left.Equals(right);

        public static explicit operator string(SqlImageSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlImageSku other && Equals(other);
        public bool Equals(SqlImageSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server Management type.
    /// </summary>
    [EnumType]
    public readonly struct SqlManagementMode : IEquatable<SqlManagementMode>
    {
        private readonly string _value;

        private SqlManagementMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlManagementMode Full { get; } = new SqlManagementMode("Full");
        public static SqlManagementMode LightWeight { get; } = new SqlManagementMode("LightWeight");
        public static SqlManagementMode NoAgent { get; } = new SqlManagementMode("NoAgent");

        public static bool operator ==(SqlManagementMode left, SqlManagementMode right) => left.Equals(right);
        public static bool operator !=(SqlManagementMode left, SqlManagementMode right) => !left.Equals(right);

        public static explicit operator string(SqlManagementMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlManagementMode other && Equals(other);
        public bool Equals(SqlManagementMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server license type.
    /// </summary>
    [EnumType]
    public readonly struct SqlServerLicenseType : IEquatable<SqlServerLicenseType>
    {
        private readonly string _value;

        private SqlServerLicenseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlServerLicenseType PAYG { get; } = new SqlServerLicenseType("PAYG");
        public static SqlServerLicenseType AHUB { get; } = new SqlServerLicenseType("AHUB");
        public static SqlServerLicenseType DR { get; } = new SqlServerLicenseType("DR");

        public static bool operator ==(SqlServerLicenseType left, SqlServerLicenseType right) => left.Equals(right);
        public static bool operator !=(SqlServerLicenseType left, SqlServerLicenseType right) => !left.Equals(right);

        public static explicit operator string(SqlServerLicenseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlServerLicenseType other && Equals(other);
        public bool Equals(SqlServerLicenseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL image sku.
    /// </summary>
    [EnumType]
    public readonly struct SqlVmGroupImageSku : IEquatable<SqlVmGroupImageSku>
    {
        private readonly string _value;

        private SqlVmGroupImageSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlVmGroupImageSku Developer { get; } = new SqlVmGroupImageSku("Developer");
        public static SqlVmGroupImageSku Enterprise { get; } = new SqlVmGroupImageSku("Enterprise");

        public static bool operator ==(SqlVmGroupImageSku left, SqlVmGroupImageSku right) => left.Equals(right);
        public static bool operator !=(SqlVmGroupImageSku left, SqlVmGroupImageSku right) => !left.Equals(right);

        public static explicit operator string(SqlVmGroupImageSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlVmGroupImageSku other && Equals(other);
        public bool Equals(SqlVmGroupImageSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SQL Server workload type.
    /// </summary>
    [EnumType]
    public readonly struct SqlWorkloadType : IEquatable<SqlWorkloadType>
    {
        private readonly string _value;

        private SqlWorkloadType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SqlWorkloadType GENERAL { get; } = new SqlWorkloadType("GENERAL");
        public static SqlWorkloadType OLTP { get; } = new SqlWorkloadType("OLTP");
        public static SqlWorkloadType DW { get; } = new SqlWorkloadType("DW");

        public static bool operator ==(SqlWorkloadType left, SqlWorkloadType right) => left.Equals(right);
        public static bool operator !=(SqlWorkloadType left, SqlWorkloadType right) => !left.Equals(right);

        public static explicit operator string(SqlWorkloadType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SqlWorkloadType other && Equals(other);
        public bool Equals(SqlWorkloadType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Storage workload type.
    /// </summary>
    [EnumType]
    public readonly struct StorageWorkloadType : IEquatable<StorageWorkloadType>
    {
        private readonly string _value;

        private StorageWorkloadType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static StorageWorkloadType GENERAL { get; } = new StorageWorkloadType("GENERAL");
        public static StorageWorkloadType OLTP { get; } = new StorageWorkloadType("OLTP");
        public static StorageWorkloadType DW { get; } = new StorageWorkloadType("DW");

        public static bool operator ==(StorageWorkloadType left, StorageWorkloadType right) => left.Equals(right);
        public static bool operator !=(StorageWorkloadType left, StorageWorkloadType right) => !left.Equals(right);

        public static explicit operator string(StorageWorkloadType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is StorageWorkloadType other && Equals(other);
        public bool Equals(StorageWorkloadType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
