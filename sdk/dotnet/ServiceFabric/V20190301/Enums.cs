// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ServiceFabric.V20190301
{
    /// <summary>
    /// The activation Mode of the service package
    /// </summary>
    [EnumType]
    public readonly struct ArmServicePackageActivationMode : IEquatable<ArmServicePackageActivationMode>
    {
        private readonly string _value;

        private ArmServicePackageActivationMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates the application package activation mode will use shared process.
        /// </summary>
        public static ArmServicePackageActivationMode SharedProcess { get; } = new ArmServicePackageActivationMode("SharedProcess");
        /// <summary>
        /// Indicates the application package activation mode will use exclusive process.
        /// </summary>
        public static ArmServicePackageActivationMode ExclusiveProcess { get; } = new ArmServicePackageActivationMode("ExclusiveProcess");

        public static bool operator ==(ArmServicePackageActivationMode left, ArmServicePackageActivationMode right) => left.Equals(right);
        public static bool operator !=(ArmServicePackageActivationMode left, ArmServicePackageActivationMode right) => !left.Equals(right);

        public static explicit operator string(ArmServicePackageActivationMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ArmServicePackageActivationMode other && Equals(other);
        public bool Equals(ArmServicePackageActivationMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The activation Mode of the service package
    /// </summary>
    [EnumType]
    public readonly struct ArmUpgradeFailureAction : IEquatable<ArmUpgradeFailureAction>
    {
        private readonly string _value;

        private ArmUpgradeFailureAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade fails.
        /// </summary>
        public static ArmUpgradeFailureAction Rollback { get; } = new ArmUpgradeFailureAction("Rollback");
        /// <summary>
        /// Indicates that a manual repair will need to be performed by the administrator if the upgrade fails. Service Fabric will not proceed to the next upgrade domain automatically.
        /// </summary>
        public static ArmUpgradeFailureAction Manual { get; } = new ArmUpgradeFailureAction("Manual");

        public static bool operator ==(ArmUpgradeFailureAction left, ArmUpgradeFailureAction right) => left.Equals(right);
        public static bool operator !=(ArmUpgradeFailureAction left, ArmUpgradeFailureAction right) => !left.Equals(right);

        public static explicit operator string(ArmUpgradeFailureAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ArmUpgradeFailureAction other && Equals(other);
        public bool Equals(ArmUpgradeFailureAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies the move cost for the service.
    /// </summary>
    [EnumType]
    public readonly struct MoveCost : IEquatable<MoveCost>
    {
        private readonly string _value;

        private MoveCost(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Zero move cost. This value is zero.
        /// </summary>
        public static MoveCost Zero { get; } = new MoveCost("Zero");
        /// <summary>
        /// Specifies the move cost of the service as Low. The value is 1.
        /// </summary>
        public static MoveCost Low { get; } = new MoveCost("Low");
        /// <summary>
        /// Specifies the move cost of the service as Medium. The value is 2.
        /// </summary>
        public static MoveCost Medium { get; } = new MoveCost("Medium");
        /// <summary>
        /// Specifies the move cost of the service as High. The value is 3.
        /// </summary>
        public static MoveCost High { get; } = new MoveCost("High");

        public static bool operator ==(MoveCost left, MoveCost right) => left.Equals(right);
        public static bool operator !=(MoveCost left, MoveCost right) => !left.Equals(right);

        public static explicit operator string(MoveCost value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MoveCost other && Equals(other);
        public bool Equals(MoveCost other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies how the service is partitioned.
    /// </summary>
    [EnumType]
    public readonly struct PartitionScheme : IEquatable<PartitionScheme>
    {
        private readonly string _value;

        private PartitionScheme(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates the partition kind is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
        /// </summary>
        public static PartitionScheme Invalid { get; } = new PartitionScheme("Invalid");
        /// <summary>
        /// Indicates that the partition is based on string names, and is a SingletonPartitionSchemeDescription object, The value is 1.
        /// </summary>
        public static PartitionScheme Singleton { get; } = new PartitionScheme("Singleton");
        /// <summary>
        /// Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionSchemeDescription object. The value is 2.
        /// </summary>
        public static PartitionScheme UniformInt64Range { get; } = new PartitionScheme("UniformInt64Range");
        /// <summary>
        /// Indicates that the partition is based on string names, and is a NamedPartitionSchemeDescription object. The value is 3
        /// </summary>
        public static PartitionScheme Named { get; } = new PartitionScheme("Named");

        public static bool operator ==(PartitionScheme left, PartitionScheme right) => left.Equals(right);
        public static bool operator !=(PartitionScheme left, PartitionScheme right) => !left.Equals(right);

        public static explicit operator string(PartitionScheme value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PartitionScheme other && Equals(other);
        public bool Equals(PartitionScheme other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
    /// </summary>
    [EnumType]
    public readonly struct ServiceCorrelationScheme : IEquatable<ServiceCorrelationScheme>
    {
        private readonly string _value;

        private ServiceCorrelationScheme(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An invalid correlation scheme. Cannot be used. The value is zero.
        /// </summary>
        public static ServiceCorrelationScheme Invalid { get; } = new ServiceCorrelationScheme("Invalid");
        /// <summary>
        /// Indicates that this service has an affinity relationship with another service. Provided for backwards compatibility, consider preferring the Aligned or NonAlignedAffinity options. The value is 1.
        /// </summary>
        public static ServiceCorrelationScheme Affinity { get; } = new ServiceCorrelationScheme("Affinity");
        /// <summary>
        /// Aligned affinity ensures that the primaries of the partitions of the affinitized services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value is 2.
        /// </summary>
        public static ServiceCorrelationScheme AlignedAffinity { get; } = new ServiceCorrelationScheme("AlignedAffinity");
        /// <summary>
        /// Non-Aligned affinity guarantees that all replicas of each service will be placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated. The value is 3.
        /// </summary>
        public static ServiceCorrelationScheme NonAlignedAffinity { get; } = new ServiceCorrelationScheme("NonAlignedAffinity");

        public static bool operator ==(ServiceCorrelationScheme left, ServiceCorrelationScheme right) => left.Equals(right);
        public static bool operator !=(ServiceCorrelationScheme left, ServiceCorrelationScheme right) => !left.Equals(right);

        public static explicit operator string(ServiceCorrelationScheme value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceCorrelationScheme other && Equals(other);
        public bool Equals(ServiceCorrelationScheme other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of service (Stateless or Stateful).
    /// </summary>
    [EnumType]
    public readonly struct ServiceKind : IEquatable<ServiceKind>
    {
        private readonly string _value;

        private ServiceKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates the service kind is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
        /// </summary>
        public static ServiceKind Invalid { get; } = new ServiceKind("Invalid");
        /// <summary>
        /// Does not use Service Fabric to make its state highly available or reliable. The value is 1.
        /// </summary>
        public static ServiceKind Stateless { get; } = new ServiceKind("Stateless");
        /// <summary>
        /// Uses Service Fabric to make its state or part of its state highly available and reliable. The value is 2.
        /// </summary>
        public static ServiceKind Stateful { get; } = new ServiceKind("Stateful");

        public static bool operator ==(ServiceKind left, ServiceKind right) => left.Equals(right);
        public static bool operator !=(ServiceKind left, ServiceKind right) => !left.Equals(right);

        public static explicit operator string(ServiceKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceKind other && Equals(other);
        public bool Equals(ServiceKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The service load metric relative weight, compared to other metrics configured for this service, as a number.
    /// </summary>
    [EnumType]
    public readonly struct ServiceLoadMetricWeight : IEquatable<ServiceLoadMetricWeight>
    {
        private readonly string _value;

        private ServiceLoadMetricWeight(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Disables resource balancing for this metric. This value is zero.
        /// </summary>
        public static ServiceLoadMetricWeight Zero { get; } = new ServiceLoadMetricWeight("Zero");
        /// <summary>
        /// Specifies the metric weight of the service load as Low. The value is 1.
        /// </summary>
        public static ServiceLoadMetricWeight Low { get; } = new ServiceLoadMetricWeight("Low");
        /// <summary>
        /// Specifies the metric weight of the service load as Medium. The value is 2.
        /// </summary>
        public static ServiceLoadMetricWeight Medium { get; } = new ServiceLoadMetricWeight("Medium");
        /// <summary>
        /// Specifies the metric weight of the service load as High. The value is 3.
        /// </summary>
        public static ServiceLoadMetricWeight High { get; } = new ServiceLoadMetricWeight("High");

        public static bool operator ==(ServiceLoadMetricWeight left, ServiceLoadMetricWeight right) => left.Equals(right);
        public static bool operator !=(ServiceLoadMetricWeight left, ServiceLoadMetricWeight right) => !left.Equals(right);

        public static explicit operator string(ServiceLoadMetricWeight value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceLoadMetricWeight other && Equals(other);
        public bool Equals(ServiceLoadMetricWeight other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of placement policy for a service fabric service. Following are the possible values.
    /// </summary>
    [EnumType]
    public readonly struct ServicePlacementPolicyType : IEquatable<ServicePlacementPolicyType>
    {
        private readonly string _value;

        private ServicePlacementPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Indicates the type of the placement policy is invalid. All Service Fabric enumerations have the invalid type. The value is zero.
        /// </summary>
        public static ServicePlacementPolicyType Invalid { get; } = new ServicePlacementPolicyType("Invalid");
        /// <summary>
        /// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription, which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 1.
        /// </summary>
        public static ServicePlacementPolicyType InvalidDomain { get; } = new ServicePlacementPolicyType("InvalidDomain");
        /// <summary>
        /// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription indicating that the replicas of the service must be placed in a specific domain. The value is 2.
        /// </summary>
        public static ServicePlacementPolicyType RequiredDomain { get; } = new ServicePlacementPolicyType("RequiredDomain");
        /// <summary>
        /// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription, which indicates that if possible the Primary replica for the partitions of the service should be located in a particular domain as an optimization. The value is 3.
        /// </summary>
        public static ServicePlacementPolicyType PreferredPrimaryDomain { get; } = new ServicePlacementPolicyType("PreferredPrimaryDomain");
        /// <summary>
        /// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two replicas from the same partition in the same domain at any time. The value is 4.
        /// </summary>
        public static ServicePlacementPolicyType RequiredDomainDistribution { get; } = new ServicePlacementPolicyType("RequiredDomainDistribution");
        /// <summary>
        /// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription, which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The value is 5.
        /// </summary>
        public static ServicePlacementPolicyType NonPartiallyPlaceService { get; } = new ServicePlacementPolicyType("NonPartiallyPlaceService");

        public static bool operator ==(ServicePlacementPolicyType left, ServicePlacementPolicyType right) => left.Equals(right);
        public static bool operator !=(ServicePlacementPolicyType left, ServicePlacementPolicyType right) => !left.Equals(right);

        public static explicit operator string(ServicePlacementPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServicePlacementPolicyType other && Equals(other);
        public bool Equals(ServicePlacementPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
