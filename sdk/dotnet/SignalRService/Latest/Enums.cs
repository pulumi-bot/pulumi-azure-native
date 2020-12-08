// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.SignalRService.Latest
{
    /// <summary>
    /// Default action when no other rule matches
    /// </summary>
    [EnumType]
    public readonly struct ACLAction : IEquatable<ACLAction>
    {
        private readonly string _value;

        private ACLAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ACLAction Allow { get; } = new ACLAction("Allow");
        public static ACLAction Deny { get; } = new ACLAction("Deny");

        public static bool operator ==(ACLAction left, ACLAction right) => left.Equals(right);
        public static bool operator !=(ACLAction left, ACLAction right) => !left.Equals(right);

        public static explicit operator string(ACLAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ACLAction other && Equals(other);
        public bool Equals(ACLAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// FeatureFlags is the supported features of Azure SignalR service.
    /// - ServiceMode: Flag for backend server for SignalR service. Values allowed: "Default": have your own backend server; "Serverless": your application doesn't have a backend server; "Classic": for backward compatibility. Support both Default and Serverless mode but not recommended; "PredefinedOnly": for future use.
    /// - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log category respectively.
    /// </summary>
    [EnumType]
    public readonly struct FeatureFlags : IEquatable<FeatureFlags>
    {
        private readonly string _value;

        private FeatureFlags(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FeatureFlags ServiceMode { get; } = new FeatureFlags("ServiceMode");
        public static FeatureFlags EnableConnectivityLogs { get; } = new FeatureFlags("EnableConnectivityLogs");
        public static FeatureFlags EnableMessagingLogs { get; } = new FeatureFlags("EnableMessagingLogs");

        public static bool operator ==(FeatureFlags left, FeatureFlags right) => left.Equals(right);
        public static bool operator !=(FeatureFlags left, FeatureFlags right) => !left.Equals(right);

        public static explicit operator string(FeatureFlags value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FeatureFlags other && Equals(other);
        public bool Equals(FeatureFlags other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    /// </summary>
    [EnumType]
    public readonly struct PrivateLinkServiceConnectionStatus : IEquatable<PrivateLinkServiceConnectionStatus>
    {
        private readonly string _value;

        private PrivateLinkServiceConnectionStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static PrivateLinkServiceConnectionStatus Pending { get; } = new PrivateLinkServiceConnectionStatus("Pending");
        public static PrivateLinkServiceConnectionStatus Approved { get; } = new PrivateLinkServiceConnectionStatus("Approved");
        public static PrivateLinkServiceConnectionStatus Rejected { get; } = new PrivateLinkServiceConnectionStatus("Rejected");
        public static PrivateLinkServiceConnectionStatus Disconnected { get; } = new PrivateLinkServiceConnectionStatus("Disconnected");

        public static bool operator ==(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => left.Equals(right);
        public static bool operator !=(PrivateLinkServiceConnectionStatus left, PrivateLinkServiceConnectionStatus right) => !left.Equals(right);

        public static explicit operator string(PrivateLinkServiceConnectionStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is PrivateLinkServiceConnectionStatus other && Equals(other);
        public bool Equals(PrivateLinkServiceConnectionStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
    /// </summary>
    [EnumType]
    public readonly struct ServiceKind : IEquatable<ServiceKind>
    {
        private readonly string _value;

        private ServiceKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceKind SignalR { get; } = new ServiceKind("SignalR");
        public static ServiceKind RawWebSockets { get; } = new ServiceKind("RawWebSockets");

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
    /// Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
    /// </summary>
    [EnumType]
    public readonly struct SignalRRequestType : IEquatable<SignalRRequestType>
    {
        private readonly string _value;

        private SignalRRequestType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SignalRRequestType ClientConnection { get; } = new SignalRRequestType("ClientConnection");
        public static SignalRRequestType ServerConnection { get; } = new SignalRRequestType("ServerConnection");
        public static SignalRRequestType RESTAPI { get; } = new SignalRRequestType("RESTAPI");

        public static bool operator ==(SignalRRequestType left, SignalRRequestType right) => left.Equals(right);
        public static bool operator !=(SignalRRequestType left, SignalRRequestType right) => !left.Equals(right);

        public static explicit operator string(SignalRRequestType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignalRRequestType other && Equals(other);
        public bool Equals(SignalRRequestType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Optional tier of this particular SKU. 'Standard' or 'Free'. 
    /// 
    /// `Basic` is deprecated, use `Standard` instead.
    /// </summary>
    [EnumType]
    public readonly struct SignalRSkuTier : IEquatable<SignalRSkuTier>
    {
        private readonly string _value;

        private SignalRSkuTier(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SignalRSkuTier Free { get; } = new SignalRSkuTier("Free");
        public static SignalRSkuTier Basic { get; } = new SignalRSkuTier("Basic");
        public static SignalRSkuTier Standard { get; } = new SignalRSkuTier("Standard");
        public static SignalRSkuTier Premium { get; } = new SignalRSkuTier("Premium");

        public static bool operator ==(SignalRSkuTier left, SignalRSkuTier right) => left.Equals(right);
        public static bool operator !=(SignalRSkuTier left, SignalRSkuTier right) => !left.Equals(right);

        public static explicit operator string(SignalRSkuTier value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SignalRSkuTier other && Equals(other);
        public bool Equals(SignalRSkuTier other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
