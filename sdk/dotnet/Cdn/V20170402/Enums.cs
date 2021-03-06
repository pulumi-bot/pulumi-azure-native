// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Cdn.V20170402
{
    /// <summary>
    /// Action of the geo filter, i.e. allow or block access.
    /// </summary>
    [EnumType]
    public readonly struct GeoFilterActions : IEquatable<GeoFilterActions>
    {
        private readonly string _value;

        private GeoFilterActions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static GeoFilterActions Block { get; } = new GeoFilterActions("Block");
        public static GeoFilterActions Allow { get; } = new GeoFilterActions("Allow");

        public static bool operator ==(GeoFilterActions left, GeoFilterActions right) => left.Equals(right);
        public static bool operator !=(GeoFilterActions left, GeoFilterActions right) => !left.Equals(right);

        public static explicit operator string(GeoFilterActions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GeoFilterActions other && Equals(other);
        public bool Equals(GeoFilterActions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
    /// </summary>
    [EnumType]
    public readonly struct OptimizationType : IEquatable<OptimizationType>
    {
        private readonly string _value;

        private OptimizationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OptimizationType GeneralWebDelivery { get; } = new OptimizationType("GeneralWebDelivery");
        public static OptimizationType GeneralMediaStreaming { get; } = new OptimizationType("GeneralMediaStreaming");
        public static OptimizationType VideoOnDemandMediaStreaming { get; } = new OptimizationType("VideoOnDemandMediaStreaming");
        public static OptimizationType LargeFileDownload { get; } = new OptimizationType("LargeFileDownload");
        public static OptimizationType DynamicSiteAcceleration { get; } = new OptimizationType("DynamicSiteAcceleration");

        public static bool operator ==(OptimizationType left, OptimizationType right) => left.Equals(right);
        public static bool operator !=(OptimizationType left, OptimizationType right) => !left.Equals(right);

        public static explicit operator string(OptimizationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OptimizationType other && Equals(other);
        public bool Equals(OptimizationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
    /// </summary>
    [EnumType]
    public readonly struct QueryStringCachingBehavior : IEquatable<QueryStringCachingBehavior>
    {
        private readonly string _value;

        private QueryStringCachingBehavior(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static QueryStringCachingBehavior IgnoreQueryString { get; } = new QueryStringCachingBehavior("IgnoreQueryString");
        public static QueryStringCachingBehavior BypassCaching { get; } = new QueryStringCachingBehavior("BypassCaching");
        public static QueryStringCachingBehavior UseQueryString { get; } = new QueryStringCachingBehavior("UseQueryString");
        public static QueryStringCachingBehavior NotSet { get; } = new QueryStringCachingBehavior("NotSet");

        public static bool operator ==(QueryStringCachingBehavior left, QueryStringCachingBehavior right) => left.Equals(right);
        public static bool operator !=(QueryStringCachingBehavior left, QueryStringCachingBehavior right) => !left.Equals(right);

        public static explicit operator string(QueryStringCachingBehavior value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is QueryStringCachingBehavior other && Equals(other);
        public bool Equals(QueryStringCachingBehavior other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Name of the pricing tier.
    /// </summary>
    [EnumType]
    public readonly struct SkuName : IEquatable<SkuName>
    {
        private readonly string _value;

        private SkuName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SkuName Standard_Verizon { get; } = new SkuName("Standard_Verizon");
        public static SkuName Premium_Verizon { get; } = new SkuName("Premium_Verizon");
        public static SkuName Custom_Verizon { get; } = new SkuName("Custom_Verizon");
        public static SkuName Standard_Akamai { get; } = new SkuName("Standard_Akamai");
        public static SkuName Standard_ChinaCdn { get; } = new SkuName("Standard_ChinaCdn");

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
}
