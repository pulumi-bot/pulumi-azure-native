// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20190501
{
    /// <summary>
    /// The action of virtual network rule.
    /// </summary>
    [EnumType]
    public readonly struct Action : IEquatable<Action>
    {
        private readonly string _value;

        private Action(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Action Allow { get; } = new Action("Allow");

        public static bool operator ==(Action left, Action right) => left.Equals(right);
        public static bool operator !=(Action left, Action right) => !left.Equals(right);

        public static explicit operator string(Action value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Action other && Equals(other);
        public bool Equals(Action other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The default action of allow or deny when no other rules match.
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
    /// The value that indicates whether the policy is enabled or not.
    /// </summary>
    [EnumType]
    public readonly struct PolicyStatus : IEquatable<PolicyStatus>
    {
        private readonly string _value;

        private PolicyStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PolicyStatus Enabled { get; } = new PolicyStatus("enabled");
        public static PolicyStatus Disabled { get; } = new PolicyStatus("disabled");

        public static bool operator ==(PolicyStatus left, PolicyStatus right) => left.Equals(right);
        public static bool operator !=(PolicyStatus left, PolicyStatus right) => !left.Equals(right);

        public static explicit operator string(PolicyStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PolicyStatus other && Equals(other);
        public bool Equals(PolicyStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The SKU name of the container registry. Required for registry creation.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName Classic { get; } = new SkuName("Classic");
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
    /// The type of trust policy.
    /// </summary>
    [EnumType]
    public readonly struct TrustPolicyType : IEquatable<TrustPolicyType>
    {
        private readonly string _value;

        private TrustPolicyType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TrustPolicyType Notary { get; } = new TrustPolicyType("Notary");

        public static bool operator ==(TrustPolicyType left, TrustPolicyType right) => left.Equals(right);
        public static bool operator !=(TrustPolicyType left, TrustPolicyType right) => !left.Equals(right);

        public static explicit operator string(TrustPolicyType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TrustPolicyType other && Equals(other);
        public bool Equals(TrustPolicyType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct WebhookAction : IEquatable<WebhookAction>
    {
        private readonly string _value;

        private WebhookAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WebhookAction Push { get; } = new WebhookAction("push");
        public static WebhookAction Delete { get; } = new WebhookAction("delete");
        public static WebhookAction Quarantine { get; } = new WebhookAction("quarantine");
        public static WebhookAction Chart_push { get; } = new WebhookAction("chart_push");
        public static WebhookAction Chart_delete { get; } = new WebhookAction("chart_delete");

        public static bool operator ==(WebhookAction left, WebhookAction right) => left.Equals(right);
        public static bool operator !=(WebhookAction left, WebhookAction right) => !left.Equals(right);

        public static explicit operator string(WebhookAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookAction other && Equals(other);
        public bool Equals(WebhookAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the webhook at the time the operation was called.
    /// </summary>
    [EnumType]
    public readonly struct WebhookStatus : IEquatable<WebhookStatus>
    {
        private readonly string _value;

        private WebhookStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WebhookStatus Enabled { get; } = new WebhookStatus("enabled");
        public static WebhookStatus Disabled { get; } = new WebhookStatus("disabled");

        public static bool operator ==(WebhookStatus left, WebhookStatus right) => left.Equals(right);
        public static bool operator !=(WebhookStatus left, WebhookStatus right) => !left.Equals(right);

        public static explicit operator string(WebhookStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WebhookStatus other && Equals(other);
        public bool Equals(WebhookStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
