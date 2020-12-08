// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20191201Preview
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
    /// A message indicating if changes on the service provider require any updates on the consumer.
    /// </summary>
    [EnumType]
    public readonly struct ActionsRequired : IEquatable<ActionsRequired>
    {
        private readonly string _value;

        private ActionsRequired(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ActionsRequired None { get; } = new ActionsRequired("None");
        public static ActionsRequired Recreate { get; } = new ActionsRequired("Recreate");

        public static bool operator ==(ActionsRequired left, ActionsRequired right) => left.Equals(right);
        public static bool operator !=(ActionsRequired left, ActionsRequired right) => !left.Equals(right);

        public static explicit operator string(ActionsRequired value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ActionsRequired other && Equals(other);
        public bool Equals(ActionsRequired other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The private link service connection status.
    /// </summary>
    [EnumType]
    public readonly struct ConnectionStatus : IEquatable<ConnectionStatus>
    {
        private readonly string _value;

        private ConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ConnectionStatus Approved { get; } = new ConnectionStatus("Approved");
        public static ConnectionStatus Pending { get; } = new ConnectionStatus("Pending");
        public static ConnectionStatus Rejected { get; } = new ConnectionStatus("Rejected");
        public static ConnectionStatus Disconnected { get; } = new ConnectionStatus("Disconnected");

        public static bool operator ==(ConnectionStatus left, ConnectionStatus right) => left.Equals(right);
        public static bool operator !=(ConnectionStatus left, ConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(ConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ConnectionStatus other && Equals(other);
        public bool Equals(ConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
    /// Indicates whether or not the encryption is enabled for container registry.
    /// </summary>
    [EnumType]
    public readonly struct EncryptionStatus : IEquatable<EncryptionStatus>
    {
        private readonly string _value;

        private EncryptionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static EncryptionStatus Enabled { get; } = new EncryptionStatus("enabled");
        public static EncryptionStatus Disabled { get; } = new EncryptionStatus("disabled");

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

    [EnumType]
    public readonly struct PipelineOptions : IEquatable<PipelineOptions>
    {
        private readonly string _value;

        private PipelineOptions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PipelineOptions OverwriteTags { get; } = new PipelineOptions("OverwriteTags");
        public static PipelineOptions OverwriteBlobs { get; } = new PipelineOptions("OverwriteBlobs");
        public static PipelineOptions DeleteSourceBlobOnSuccess { get; } = new PipelineOptions("DeleteSourceBlobOnSuccess");
        public static PipelineOptions ContinueOnErrors { get; } = new PipelineOptions("ContinueOnErrors");

        public static bool operator ==(PipelineOptions left, PipelineOptions right) => left.Equals(right);
        public static bool operator !=(PipelineOptions left, PipelineOptions right) => !left.Equals(right);

        public static explicit operator string(PipelineOptions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PipelineOptions other && Equals(other);
        public bool Equals(PipelineOptions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the source.
    /// </summary>
    [EnumType]
    public readonly struct PipelineRunSourceType : IEquatable<PipelineRunSourceType>
    {
        private readonly string _value;

        private PipelineRunSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PipelineRunSourceType AzureStorageBlob { get; } = new PipelineRunSourceType("AzureStorageBlob");

        public static bool operator ==(PipelineRunSourceType left, PipelineRunSourceType right) => left.Equals(right);
        public static bool operator !=(PipelineRunSourceType left, PipelineRunSourceType right) => !left.Equals(right);

        public static explicit operator string(PipelineRunSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PipelineRunSourceType other && Equals(other);
        public bool Equals(PipelineRunSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of the target.
    /// </summary>
    [EnumType]
    public readonly struct PipelineRunTargetType : IEquatable<PipelineRunTargetType>
    {
        private readonly string _value;

        private PipelineRunTargetType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PipelineRunTargetType AzureStorageBlob { get; } = new PipelineRunTargetType("AzureStorageBlob");

        public static bool operator ==(PipelineRunTargetType left, PipelineRunTargetType right) => left.Equals(right);
        public static bool operator !=(PipelineRunTargetType left, PipelineRunTargetType right) => !left.Equals(right);

        public static explicit operator string(PipelineRunTargetType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PipelineRunTargetType other && Equals(other);
        public bool Equals(PipelineRunTargetType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The type of source for the import pipeline.
    /// </summary>
    [EnumType]
    public readonly struct PipelineSourceType : IEquatable<PipelineSourceType>
    {
        private readonly string _value;

        private PipelineSourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PipelineSourceType AzureStorageBlobContainer { get; } = new PipelineSourceType("AzureStorageBlobContainer");

        public static bool operator ==(PipelineSourceType left, PipelineSourceType right) => left.Equals(right);
        public static bool operator !=(PipelineSourceType left, PipelineSourceType right) => !left.Equals(right);

        public static explicit operator string(PipelineSourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PipelineSourceType other && Equals(other);
        public bool Equals(PipelineSourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
    /// Whether or not public network access is allowed for the container registry.
    /// </summary>
    [EnumType]
    public readonly struct PublicNetworkAccess : IEquatable<PublicNetworkAccess>
    {
        private readonly string _value;

        private PublicNetworkAccess(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PublicNetworkAccess Enabled { get; } = new PublicNetworkAccess("Enabled");
        public static PublicNetworkAccess Disabled { get; } = new PublicNetworkAccess("Disabled");

        public static bool operator ==(PublicNetworkAccess left, PublicNetworkAccess right) => left.Equals(right);
        public static bool operator !=(PublicNetworkAccess left, PublicNetworkAccess right) => !left.Equals(right);

        public static explicit operator string(PublicNetworkAccess value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PublicNetworkAccess other && Equals(other);
        public bool Equals(PublicNetworkAccess other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
        public static ResourceIdentityType UserAssigned { get; } = new ResourceIdentityType("UserAssigned");
        public static ResourceIdentityType SystemAssigned_UserAssigned { get; } = new ResourceIdentityType("SystemAssigned, UserAssigned");
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
    /// The current status of the source trigger.
    /// </summary>
    [EnumType]
    public readonly struct TriggerStatus : IEquatable<TriggerStatus>
    {
        private readonly string _value;

        private TriggerStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TriggerStatus Enabled { get; } = new TriggerStatus("Enabled");
        public static TriggerStatus Disabled { get; } = new TriggerStatus("Disabled");

        public static bool operator ==(TriggerStatus left, TriggerStatus right) => left.Equals(right);
        public static bool operator !=(TriggerStatus left, TriggerStatus right) => !left.Equals(right);

        public static explicit operator string(TriggerStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TriggerStatus other && Equals(other);
        public bool Equals(TriggerStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
