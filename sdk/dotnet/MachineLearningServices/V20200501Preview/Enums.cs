// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.MachineLearningServices.V20200501Preview
{
    /// <summary>
    /// The compute environment type for the service.
    /// </summary>
    [EnumType]
    public readonly struct ComputeEnvironmentType : IEquatable<ComputeEnvironmentType>
    {
        private readonly string _value;

        private ComputeEnvironmentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComputeEnvironmentType ACI { get; } = new ComputeEnvironmentType("ACI");
        public static ComputeEnvironmentType AKS { get; } = new ComputeEnvironmentType("AKS");

        public static bool operator ==(ComputeEnvironmentType left, ComputeEnvironmentType right) => left.Equals(right);
        public static bool operator !=(ComputeEnvironmentType left, ComputeEnvironmentType right) => !left.Equals(right);

        public static explicit operator string(ComputeEnvironmentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeEnvironmentType other && Equals(other);
        public bool Equals(ComputeEnvironmentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of compute
    /// </summary>
    [EnumType]
    public readonly struct ComputeType : IEquatable<ComputeType>
    {
        private readonly string _value;

        private ComputeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComputeType AKS { get; } = new ComputeType("AKS");
        public static ComputeType AmlCompute { get; } = new ComputeType("AmlCompute");
        public static ComputeType DataFactory { get; } = new ComputeType("DataFactory");
        public static ComputeType VirtualMachine { get; } = new ComputeType("VirtualMachine");
        public static ComputeType HDInsight { get; } = new ComputeType("HDInsight");
        public static ComputeType Databricks { get; } = new ComputeType("Databricks");
        public static ComputeType DataLakeAnalytics { get; } = new ComputeType("DataLakeAnalytics");

        public static bool operator ==(ComputeType left, ComputeType right) => left.Equals(right);
        public static bool operator !=(ComputeType left, ComputeType right) => !left.Equals(right);

        public static explicit operator string(ComputeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeType other && Equals(other);
        public bool Equals(ComputeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies dataset type.
    /// </summary>
    [EnumType]
    public readonly struct DatasetType : IEquatable<DatasetType>
    {
        private readonly string _value;

        private DatasetType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatasetType Tabular { get; } = new DatasetType("tabular");
        public static DatasetType File { get; } = new DatasetType("file");

        public static bool operator ==(DatasetType left, DatasetType right) => left.Equals(right);
        public static bool operator !=(DatasetType left, DatasetType right) => !left.Equals(right);

        public static explicit operator string(DatasetType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatasetType other && Equals(other);
        public bool Equals(DatasetType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies datastore type.
    /// </summary>
    [EnumType]
    public readonly struct DatastoreTypeArm : IEquatable<DatastoreTypeArm>
    {
        private readonly string _value;

        private DatastoreTypeArm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatastoreTypeArm Blob { get; } = new DatastoreTypeArm("blob");
        public static DatastoreTypeArm Adls { get; } = new DatastoreTypeArm("adls");
        public static DatastoreTypeArm Adls_gen2 { get; } = new DatastoreTypeArm("adls-gen2");
        public static DatastoreTypeArm Dbfs { get; } = new DatastoreTypeArm("dbfs");
        public static DatastoreTypeArm File { get; } = new DatastoreTypeArm("file");
        public static DatastoreTypeArm Mysqldb { get; } = new DatastoreTypeArm("mysqldb");
        public static DatastoreTypeArm Sqldb { get; } = new DatastoreTypeArm("sqldb");
        public static DatastoreTypeArm Psqldb { get; } = new DatastoreTypeArm("psqldb");

        public static bool operator ==(DatastoreTypeArm left, DatastoreTypeArm right) => left.Equals(right);
        public static bool operator !=(DatastoreTypeArm left, DatastoreTypeArm right) => !left.Equals(right);

        public static explicit operator string(DatastoreTypeArm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatastoreTypeArm other && Equals(other);
        public bool Equals(DatastoreTypeArm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether or not the encryption is enabled for the workspace.
    /// </summary>
    [EnumType]
    public readonly struct EncryptionStatus : IEquatable<EncryptionStatus>
    {
        private readonly string _value;

        private EncryptionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EncryptionStatus Enabled { get; } = new EncryptionStatus("Enabled");
        public static EncryptionStatus Disabled { get; } = new EncryptionStatus("Disabled");

        public static bool operator ==(EncryptionStatus left, EncryptionStatus right) => left.Equals(right);
        public static bool operator !=(EncryptionStatus left, EncryptionStatus right) => !left.Equals(right);

        public static explicit operator string(EncryptionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EncryptionStatus other && Equals(other);
        public bool Equals(EncryptionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Header type.
    /// </summary>
    [EnumType]
    public readonly struct Header : IEquatable<Header>
    {
        private readonly string _value;

        private Header(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Header All_files_have_same_headers { get; } = new Header("all_files_have_same_headers");
        public static Header Only_first_file_has_headers { get; } = new Header("only_first_file_has_headers");
        public static Header No_headers { get; } = new Header("no_headers");
        public static Header Combine_all_files_headers { get; } = new Header("combine_all_files_headers");

        public static bool operator ==(Header left, Header right) => left.Equals(right);
        public static bool operator !=(Header left, Header right) => !left.Equals(right);

        public static explicit operator string(Header value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Header other && Equals(other);
        public bool Equals(Header other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    /// </summary>
    [EnumType]
    public readonly struct PrivateEndpointServiceConnectionStatus : IEquatable<PrivateEndpointServiceConnectionStatus>
    {
        private readonly string _value;

        private PrivateEndpointServiceConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateEndpointServiceConnectionStatus Pending { get; } = new PrivateEndpointServiceConnectionStatus("Pending");
        public static PrivateEndpointServiceConnectionStatus Approved { get; } = new PrivateEndpointServiceConnectionStatus("Approved");
        public static PrivateEndpointServiceConnectionStatus Rejected { get; } = new PrivateEndpointServiceConnectionStatus("Rejected");
        public static PrivateEndpointServiceConnectionStatus Disconnected { get; } = new PrivateEndpointServiceConnectionStatus("Disconnected");
        public static PrivateEndpointServiceConnectionStatus Timeout { get; } = new PrivateEndpointServiceConnectionStatus("Timeout");

        public static bool operator ==(PrivateEndpointServiceConnectionStatus left, PrivateEndpointServiceConnectionStatus right) => left.Equals(right);
        public static bool operator !=(PrivateEndpointServiceConnectionStatus left, PrivateEndpointServiceConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateEndpointServiceConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateEndpointServiceConnectionStatus other && Equals(other);
        public bool Equals(PrivateEndpointServiceConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// State of the public SSH port. Possible values are: Disabled - Indicates that the public ssh port is closed on all nodes of the cluster. Enabled - Indicates that the public ssh port is open on all nodes of the cluster. NotSpecified - Indicates that the public ssh port is closed on all nodes of the cluster if VNet is defined, else is open all public nodes. It can be default only during cluster creation time, after creation it will be either enabled or disabled.
    /// </summary>
    [EnumType]
    public readonly struct RemoteLoginPortPublicAccess : IEquatable<RemoteLoginPortPublicAccess>
    {
        private readonly string _value;

        private RemoteLoginPortPublicAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RemoteLoginPortPublicAccess Enabled { get; } = new RemoteLoginPortPublicAccess("Enabled");
        public static RemoteLoginPortPublicAccess Disabled { get; } = new RemoteLoginPortPublicAccess("Disabled");
        public static RemoteLoginPortPublicAccess NotSpecified { get; } = new RemoteLoginPortPublicAccess("NotSpecified");

        public static bool operator ==(RemoteLoginPortPublicAccess left, RemoteLoginPortPublicAccess right) => left.Equals(right);
        public static bool operator !=(RemoteLoginPortPublicAccess left, RemoteLoginPortPublicAccess right) => !left.Equals(right);

        public static explicit operator string(RemoteLoginPortPublicAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RemoteLoginPortPublicAccess other && Equals(other);
        public bool Equals(RemoteLoginPortPublicAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The identity type.
    /// </summary>
    [EnumType]
    public readonly struct ResourceIdentityType : IEquatable<ResourceIdentityType>
    {
        private readonly string _value;

        private ResourceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ResourceIdentityType SystemAssigned { get; } = new ResourceIdentityType("SystemAssigned");
        public static ResourceIdentityType SystemAssigned_UserAssigned { get; } = new ResourceIdentityType("SystemAssigned,UserAssigned");
        public static ResourceIdentityType UserAssigned { get; } = new ResourceIdentityType("UserAssigned");
        public static ResourceIdentityType None { get; } = new ResourceIdentityType("None");

        public static bool operator ==(ResourceIdentityType left, ResourceIdentityType right) => left.Equals(right);
        public static bool operator !=(ResourceIdentityType left, ResourceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ResourceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ResourceIdentityType other && Equals(other);
        public bool Equals(ResourceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Data source type.
    /// </summary>
    [EnumType]
    public readonly struct SourceType : IEquatable<SourceType>
    {
        private readonly string _value;

        private SourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SourceType Delimited_files { get; } = new SourceType("delimited_files");
        public static SourceType Json_lines_files { get; } = new SourceType("json_lines_files");
        public static SourceType Parquet_files { get; } = new SourceType("parquet_files");

        public static bool operator ==(SourceType left, SourceType right) => left.Equals(right);
        public static bool operator !=(SourceType left, SourceType right) => !left.Equals(right);

        public static explicit operator string(SourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SourceType other && Equals(other);
        public bool Equals(SourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Virtual Machine priority
    /// </summary>
    [EnumType]
    public readonly struct VmPriority : IEquatable<VmPriority>
    {
        private readonly string _value;

        private VmPriority(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static VmPriority Dedicated { get; } = new VmPriority("Dedicated");
        public static VmPriority LowPriority { get; } = new VmPriority("LowPriority");

        public static bool operator ==(VmPriority left, VmPriority right) => left.Equals(right);
        public static bool operator !=(VmPriority left, VmPriority right) => !left.Equals(right);

        public static explicit operator string(VmPriority value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VmPriority other && Equals(other);
        public bool Equals(VmPriority other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
