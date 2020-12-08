// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Compute.V20201001Preview
{
    /// <summary>
    /// Update mode for the cloud service. Role instances are allocated to update domains when the service is deployed. Updates can be initiated manually in each update domain or initiated automatically in all update domains.
    /// Possible Values are &lt;br /&gt;&lt;br /&gt;**Auto**&lt;br /&gt;&lt;br /&gt;**Manual** &lt;br /&gt;&lt;br /&gt;**Simultaneous**&lt;br /&gt;&lt;br /&gt;
    /// If not specified, the default value is Auto. If set to Manual, PUT UpdateDomain must be called to apply the update. If set to Auto, the update is automatically applied to each update domain in sequence.
    /// </summary>
    [EnumType]
    public readonly struct CloudServiceUpgradeMode : IEquatable<CloudServiceUpgradeMode>
    {
        private readonly string _value;

        private CloudServiceUpgradeMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudServiceUpgradeMode Auto { get; } = new CloudServiceUpgradeMode("Auto");
        public static CloudServiceUpgradeMode Manual { get; } = new CloudServiceUpgradeMode("Manual");
        public static CloudServiceUpgradeMode Simultaneous { get; } = new CloudServiceUpgradeMode("Simultaneous");

        public static bool operator ==(CloudServiceUpgradeMode left, CloudServiceUpgradeMode right) => left.Equals(right);
        public static bool operator !=(CloudServiceUpgradeMode left, CloudServiceUpgradeMode right) => !left.Equals(right);

        public static explicit operator string(CloudServiceUpgradeMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudServiceUpgradeMode other && Equals(other);
        public bool Equals(CloudServiceUpgradeMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
