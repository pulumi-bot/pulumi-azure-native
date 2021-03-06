// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Blockchain.V20180601Preview
{
    /// <summary>
    /// Gets or sets the blockchain protocol.
    /// </summary>
    [EnumType]
    public readonly struct BlockchainProtocol : IEquatable<BlockchainProtocol>
    {
        private readonly string _value;

        private BlockchainProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BlockchainProtocol NotSpecified { get; } = new BlockchainProtocol("NotSpecified");
        public static BlockchainProtocol Parity { get; } = new BlockchainProtocol("Parity");
        public static BlockchainProtocol Quorum { get; } = new BlockchainProtocol("Quorum");
        public static BlockchainProtocol Corda { get; } = new BlockchainProtocol("Corda");

        public static bool operator ==(BlockchainProtocol left, BlockchainProtocol right) => left.Equals(right);
        public static bool operator !=(BlockchainProtocol left, BlockchainProtocol right) => !left.Equals(right);

        public static explicit operator string(BlockchainProtocol value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BlockchainProtocol other && Equals(other);
        public bool Equals(BlockchainProtocol other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
