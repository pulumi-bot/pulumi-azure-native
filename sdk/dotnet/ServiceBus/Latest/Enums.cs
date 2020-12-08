// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ServiceBus.Latest
{
    [EnumType]
    public readonly struct AccessRights : IEquatable<AccessRights>
    {
        private readonly string _value;

        private AccessRights(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessRights Manage { get; } = new AccessRights("Manage");
        public static AccessRights Send { get; } = new AccessRights("Send");
        public static AccessRights Listen { get; } = new AccessRights("Listen");

        public static bool operator ==(AccessRights left, AccessRights right) => left.Equals(right);
        public static bool operator !=(AccessRights left, AccessRights right) => !left.Equals(right);

        public static explicit operator string(AccessRights value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessRights other && Equals(other);
        public bool Equals(AccessRights other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Default Action for Network Rule Set
    /// </summary>
    [EnumType]
    public readonly struct DefaultAction : IEquatable<DefaultAction>
    {
        private readonly string _value;

        private DefaultAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DefaultAction Allow { get; } = new DefaultAction("Allow");
        public static DefaultAction Deny { get; } = new DefaultAction("Deny");

        public static bool operator ==(DefaultAction left, DefaultAction right) => left.Equals(right);
        public static bool operator !=(DefaultAction left, DefaultAction right) => !left.Equals(right);

        public static explicit operator string(DefaultAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DefaultAction other && Equals(other);
        public bool Equals(DefaultAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Enumerates the possible values for the status of a messaging entity.
    /// </summary>
    [EnumType]
    public readonly struct EntityStatus : IEquatable<EntityStatus>
    {
        private readonly string _value;

        private EntityStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EntityStatus Active { get; } = new EntityStatus("Active");
        public static EntityStatus Disabled { get; } = new EntityStatus("Disabled");
        public static EntityStatus Restoring { get; } = new EntityStatus("Restoring");
        public static EntityStatus SendDisabled { get; } = new EntityStatus("SendDisabled");
        public static EntityStatus ReceiveDisabled { get; } = new EntityStatus("ReceiveDisabled");
        public static EntityStatus Creating { get; } = new EntityStatus("Creating");
        public static EntityStatus Deleting { get; } = new EntityStatus("Deleting");
        public static EntityStatus Renaming { get; } = new EntityStatus("Renaming");
        public static EntityStatus Unknown { get; } = new EntityStatus("Unknown");

        public static bool operator ==(EntityStatus left, EntityStatus right) => left.Equals(right);
        public static bool operator !=(EntityStatus left, EntityStatus right) => !left.Equals(right);

        public static explicit operator string(EntityStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityStatus other && Equals(other);
        public bool Equals(EntityStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Filter type that is evaluated against a BrokeredMessage.
    /// </summary>
    [EnumType]
    public readonly struct FilterType : IEquatable<FilterType>
    {
        private readonly string _value;

        private FilterType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FilterType SqlFilter { get; } = new FilterType("SqlFilter");
        public static FilterType CorrelationFilter { get; } = new FilterType("CorrelationFilter");

        public static bool operator ==(FilterType left, FilterType right) => left.Equals(right);
        public static bool operator !=(FilterType left, FilterType right) => !left.Equals(right);

        public static explicit operator string(FilterType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FilterType other && Equals(other);
        public bool Equals(FilterType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The IP Filter Action
    /// </summary>
    [EnumType]
    public readonly struct NetworkRuleIPAction : IEquatable<NetworkRuleIPAction>
    {
        private readonly string _value;

        private NetworkRuleIPAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static NetworkRuleIPAction Allow { get; } = new NetworkRuleIPAction("Allow");

        public static bool operator ==(NetworkRuleIPAction left, NetworkRuleIPAction right) => left.Equals(right);
        public static bool operator !=(NetworkRuleIPAction left, NetworkRuleIPAction right) => !left.Equals(right);

        public static explicit operator string(NetworkRuleIPAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkRuleIPAction other && Equals(other);
        public bool Equals(NetworkRuleIPAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Name of this SKU.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName Basic { get; } = new SkuName("Basic");
        public static SkuName Standard { get; } = new SkuName("Standard");
        public static SkuName Premium { get; } = new SkuName("Premium");

        public static bool operator ==(SkuName left, SkuName right) => left.Equals(right);
        public static bool operator !=(SkuName left, SkuName right) => !left.Equals(right);

        public static explicit operator string(SkuName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuName other && Equals(other);
        public bool Equals(SkuName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The billing tier of this particular SKU.
    /// </summary>
    [EnumType]
    public readonly struct SkuTier : IEquatable<SkuTier>
    {
        private readonly string _value;

        private SkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuTier Basic { get; } = new SkuTier("Basic");
        public static SkuTier Standard { get; } = new SkuTier("Standard");
        public static SkuTier Premium { get; } = new SkuTier("Premium");

        public static bool operator ==(SkuTier left, SkuTier right) => left.Equals(right);
        public static bool operator !=(SkuTier left, SkuTier right) => !left.Equals(right);

        public static explicit operator string(SkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SkuTier other && Equals(other);
        public bool Equals(SkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
