// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20190501Preview
{
    [EnumType]
    public readonly struct TokenCertificateName : IEquatable<TokenCertificateName>
    {
        private readonly string _value;

        private TokenCertificateName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TokenCertificateName Certificate1 { get; } = new TokenCertificateName("certificate1");
        public static TokenCertificateName Certificate2 { get; } = new TokenCertificateName("certificate2");

        public static bool operator ==(TokenCertificateName left, TokenCertificateName right) => left.Equals(right);
        public static bool operator !=(TokenCertificateName left, TokenCertificateName right) => !left.Equals(right);

        public static explicit operator string(TokenCertificateName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TokenCertificateName other && Equals(other);
        public bool Equals(TokenCertificateName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The password name "password1" or "password2"
    /// </summary>
    [EnumType]
    public readonly struct TokenPasswordName : IEquatable<TokenPasswordName>
    {
        private readonly string _value;

        private TokenPasswordName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TokenPasswordName Password1 { get; } = new TokenPasswordName("password1");
        public static TokenPasswordName Password2 { get; } = new TokenPasswordName("password2");

        public static bool operator ==(TokenPasswordName left, TokenPasswordName right) => left.Equals(right);
        public static bool operator !=(TokenPasswordName left, TokenPasswordName right) => !left.Equals(right);

        public static explicit operator string(TokenPasswordName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TokenPasswordName other && Equals(other);
        public bool Equals(TokenPasswordName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the token example enabled or disabled.
    /// </summary>
    [EnumType]
    public readonly struct TokenStatus : IEquatable<TokenStatus>
    {
        private readonly string _value;

        private TokenStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TokenStatus Enabled { get; } = new TokenStatus("enabled");
        public static TokenStatus Disabled { get; } = new TokenStatus("disabled");

        public static bool operator ==(TokenStatus left, TokenStatus right) => left.Equals(right);
        public static bool operator !=(TokenStatus left, TokenStatus right) => !left.Equals(right);

        public static explicit operator string(TokenStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TokenStatus other && Equals(other);
        public bool Equals(TokenStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
