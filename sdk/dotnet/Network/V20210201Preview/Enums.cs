// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Network.V20210201Preview
{
    /// <summary>
    /// Address prefix type.
    /// </summary>
    [EnumType]
    public readonly struct AddressPrefixType : IEquatable<AddressPrefixType>
    {
        private readonly string _value;

        private AddressPrefixType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AddressPrefixType IPPrefix { get; } = new AddressPrefixType("IPPrefix");
        public static AddressPrefixType ServiceTag { get; } = new AddressPrefixType("ServiceTag");

        public static bool operator ==(AddressPrefixType left, AddressPrefixType right) => left.Equals(right);
        public static bool operator !=(AddressPrefixType left, AddressPrefixType right) => !left.Equals(right);

        public static explicit operator string(AddressPrefixType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AddressPrefixType other && Equals(other);
        public bool Equals(AddressPrefixType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Connectivity topology type.
    /// </summary>
    [EnumType]
    public readonly struct ConnectivityTopology : IEquatable<ConnectivityTopology>
    {
        private readonly string _value;

        private ConnectivityTopology(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectivityTopology HubAndSpokeTopology { get; } = new ConnectivityTopology("HubAndSpokeTopology");
        public static ConnectivityTopology MeshTopology { get; } = new ConnectivityTopology("MeshTopology");

        public static bool operator ==(ConnectivityTopology left, ConnectivityTopology right) => left.Equals(right);
        public static bool operator !=(ConnectivityTopology left, ConnectivityTopology right) => !left.Equals(right);

        public static explicit operator string(ConnectivityTopology value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectivityTopology other && Equals(other);
        public bool Equals(ConnectivityTopology other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Configuration Deployment Type.
    /// </summary>
    [EnumType]
    public readonly struct DeploymentType : IEquatable<DeploymentType>
    {
        private readonly string _value;

        private DeploymentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentType AdminPolicy { get; } = new DeploymentType("AdminPolicy");
        public static DeploymentType UserPolicy { get; } = new DeploymentType("UserPolicy");
        public static DeploymentType Routing { get; } = new DeploymentType("Routing");
        public static DeploymentType Connectivity { get; } = new DeploymentType("Connectivity");

        public static bool operator ==(DeploymentType left, DeploymentType right) => left.Equals(right);
        public static bool operator !=(DeploymentType left, DeploymentType right) => !left.Equals(right);

        public static explicit operator string(DeploymentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentType other && Equals(other);
        public bool Equals(DeploymentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Group connectivity type.
    /// </summary>
    [EnumType]
    public readonly struct GroupConnectivity : IEquatable<GroupConnectivity>
    {
        private readonly string _value;

        private GroupConnectivity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GroupConnectivity None { get; } = new GroupConnectivity("None");
        public static GroupConnectivity DirectlyConnected { get; } = new GroupConnectivity("DirectlyConnected");

        public static bool operator ==(GroupConnectivity left, GroupConnectivity right) => left.Equals(right);
        public static bool operator !=(GroupConnectivity left, GroupConnectivity right) => !left.Equals(right);

        public static explicit operator string(GroupConnectivity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GroupConnectivity other && Equals(other);
        public bool Equals(GroupConnectivity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Group member type.
    /// </summary>
    [EnumType]
    public readonly struct MemberType : IEquatable<MemberType>
    {
        private readonly string _value;

        private MemberType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static MemberType VirtualNetwork { get; } = new MemberType("VirtualNetwork");
        public static MemberType Subnet { get; } = new MemberType("Subnet");

        public static bool operator ==(MemberType left, MemberType right) => left.Equals(right);
        public static bool operator !=(MemberType left, MemberType right) => !left.Equals(right);

        public static explicit operator string(MemberType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is MemberType other && Equals(other);
        public bool Equals(MemberType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScopeAccesses : IEquatable<ScopeAccesses>
    {
        private readonly string _value;

        private ScopeAccesses(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScopeAccesses Security { get; } = new ScopeAccesses("Security");
        public static ScopeAccesses Routing { get; } = new ScopeAccesses("Routing");
        public static ScopeAccesses Connectivity { get; } = new ScopeAccesses("Connectivity");

        public static bool operator ==(ScopeAccesses left, ScopeAccesses right) => left.Equals(right);
        public static bool operator !=(ScopeAccesses left, ScopeAccesses right) => !left.Equals(right);

        public static explicit operator string(ScopeAccesses value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScopeAccesses other && Equals(other);
        public bool Equals(ScopeAccesses other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates the access allowed for this particular rule
    /// </summary>
    [EnumType]
    public readonly struct SecurityConfigurationRuleAccess : IEquatable<SecurityConfigurationRuleAccess>
    {
        private readonly string _value;

        private SecurityConfigurationRuleAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityConfigurationRuleAccess Allow { get; } = new SecurityConfigurationRuleAccess("Allow");
        public static SecurityConfigurationRuleAccess Deny { get; } = new SecurityConfigurationRuleAccess("Deny");
        public static SecurityConfigurationRuleAccess AlwaysAllow { get; } = new SecurityConfigurationRuleAccess("AlwaysAllow");

        public static bool operator ==(SecurityConfigurationRuleAccess left, SecurityConfigurationRuleAccess right) => left.Equals(right);
        public static bool operator !=(SecurityConfigurationRuleAccess left, SecurityConfigurationRuleAccess right) => !left.Equals(right);

        public static explicit operator string(SecurityConfigurationRuleAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityConfigurationRuleAccess other && Equals(other);
        public bool Equals(SecurityConfigurationRuleAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates if the traffic matched against the rule in inbound or outbound.
    /// </summary>
    [EnumType]
    public readonly struct SecurityConfigurationRuleDirection : IEquatable<SecurityConfigurationRuleDirection>
    {
        private readonly string _value;

        private SecurityConfigurationRuleDirection(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityConfigurationRuleDirection Inbound { get; } = new SecurityConfigurationRuleDirection("Inbound");
        public static SecurityConfigurationRuleDirection Outbound { get; } = new SecurityConfigurationRuleDirection("Outbound");

        public static bool operator ==(SecurityConfigurationRuleDirection left, SecurityConfigurationRuleDirection right) => left.Equals(right);
        public static bool operator !=(SecurityConfigurationRuleDirection left, SecurityConfigurationRuleDirection right) => !left.Equals(right);

        public static explicit operator string(SecurityConfigurationRuleDirection value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityConfigurationRuleDirection other && Equals(other);
        public bool Equals(SecurityConfigurationRuleDirection other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Network protocol this rule applies to.
    /// </summary>
    [EnumType]
    public readonly struct SecurityConfigurationRuleProtocol : IEquatable<SecurityConfigurationRuleProtocol>
    {
        private readonly string _value;

        private SecurityConfigurationRuleProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityConfigurationRuleProtocol Tcp { get; } = new SecurityConfigurationRuleProtocol("Tcp");
        public static SecurityConfigurationRuleProtocol Udp { get; } = new SecurityConfigurationRuleProtocol("Udp");
        public static SecurityConfigurationRuleProtocol Icmp { get; } = new SecurityConfigurationRuleProtocol("Icmp");
        public static SecurityConfigurationRuleProtocol Esp { get; } = new SecurityConfigurationRuleProtocol("Esp");
        public static SecurityConfigurationRuleProtocol Any { get; } = new SecurityConfigurationRuleProtocol("Any");
        public static SecurityConfigurationRuleProtocol Ah { get; } = new SecurityConfigurationRuleProtocol("Ah");

        public static bool operator ==(SecurityConfigurationRuleProtocol left, SecurityConfigurationRuleProtocol right) => left.Equals(right);
        public static bool operator !=(SecurityConfigurationRuleProtocol left, SecurityConfigurationRuleProtocol right) => !left.Equals(right);

        public static explicit operator string(SecurityConfigurationRuleProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityConfigurationRuleProtocol other && Equals(other);
        public bool Equals(SecurityConfigurationRuleProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Security Type.
    /// </summary>
    [EnumType]
    public readonly struct SecurityType : IEquatable<SecurityType>
    {
        private readonly string _value;

        private SecurityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SecurityType AdminPolicy { get; } = new SecurityType("AdminPolicy");
        public static SecurityType UserPolicy { get; } = new SecurityType("UserPolicy");

        public static bool operator ==(SecurityType left, SecurityType right) => left.Equals(right);
        public static bool operator !=(SecurityType left, SecurityType right) => !left.Equals(right);

        public static explicit operator string(SecurityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SecurityType other && Equals(other);
        public bool Equals(SecurityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
