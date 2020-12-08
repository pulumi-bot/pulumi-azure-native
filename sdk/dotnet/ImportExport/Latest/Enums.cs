// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ImportExport.Latest
{
    /// <summary>
    /// The drive's current state. 
    /// </summary>
    [EnumType]
    public readonly struct DriveState : IEquatable<DriveState>
    {
        private readonly string _value;

        private DriveState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DriveState Specified { get; } = new DriveState("Specified");
        public static DriveState Received { get; } = new DriveState("Received");
        public static DriveState NeverReceived { get; } = new DriveState("NeverReceived");
        public static DriveState Transferring { get; } = new DriveState("Transferring");
        public static DriveState Completed { get; } = new DriveState("Completed");
        public static DriveState CompletedMoreInfo { get; } = new DriveState("CompletedMoreInfo");
        public static DriveState ShippedBack { get; } = new DriveState("ShippedBack");

        public static bool operator ==(DriveState left, DriveState right) => left.Equals(right);
        public static bool operator !=(DriveState left, DriveState right) => !left.Equals(right);

        public static explicit operator string(DriveState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DriveState other && Equals(other);
        public bool Equals(DriveState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
