// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Confluent.Latest
{
    /// <summary>
    /// SaaS Offer Status
    /// </summary>
    [EnumType]
    public readonly struct SaaSOfferStatus : IEquatable<SaaSOfferStatus>
    {
        private readonly string _value;

        private SaaSOfferStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SaaSOfferStatus Started { get; } = new SaaSOfferStatus("Started");
        public static SaaSOfferStatus PendingFulfillmentStart { get; } = new SaaSOfferStatus("PendingFulfillmentStart");
        public static SaaSOfferStatus InProgress { get; } = new SaaSOfferStatus("InProgress");
        public static SaaSOfferStatus Subscribed { get; } = new SaaSOfferStatus("Subscribed");
        public static SaaSOfferStatus Suspended { get; } = new SaaSOfferStatus("Suspended");
        public static SaaSOfferStatus Reinstated { get; } = new SaaSOfferStatus("Reinstated");
        public static SaaSOfferStatus Succeeded { get; } = new SaaSOfferStatus("Succeeded");
        public static SaaSOfferStatus Failed { get; } = new SaaSOfferStatus("Failed");
        public static SaaSOfferStatus Unsubscribed { get; } = new SaaSOfferStatus("Unsubscribed");
        public static SaaSOfferStatus Updating { get; } = new SaaSOfferStatus("Updating");

        public static bool operator ==(SaaSOfferStatus left, SaaSOfferStatus right) => left.Equals(right);
        public static bool operator !=(SaaSOfferStatus left, SaaSOfferStatus right) => !left.Equals(right);

        public static explicit operator string(SaaSOfferStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SaaSOfferStatus other && Equals(other);
        public bool Equals(SaaSOfferStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
