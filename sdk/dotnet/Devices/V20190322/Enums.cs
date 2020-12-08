// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Devices.V20190322
{
    /// <summary>
    /// The permissions assigned to the shared access policy.
    /// </summary>
    [EnumType]
    public readonly struct AccessRights : IEquatable<AccessRights>
    {
        private readonly string _value;

        private AccessRights(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessRights RegistryRead { get; } = new AccessRights("RegistryRead");
        public static AccessRights RegistryWrite { get; } = new AccessRights("RegistryWrite");
        public static AccessRights ServiceConnect { get; } = new AccessRights("ServiceConnect");
        public static AccessRights DeviceConnect { get; } = new AccessRights("DeviceConnect");
        public static AccessRights RegistryRead_RegistryWrite { get; } = new AccessRights("RegistryRead, RegistryWrite");
        public static AccessRights RegistryRead_ServiceConnect { get; } = new AccessRights("RegistryRead, ServiceConnect");
        public static AccessRights RegistryRead_DeviceConnect { get; } = new AccessRights("RegistryRead, DeviceConnect");
        public static AccessRights RegistryWrite_ServiceConnect { get; } = new AccessRights("RegistryWrite, ServiceConnect");
        public static AccessRights RegistryWrite_DeviceConnect { get; } = new AccessRights("RegistryWrite, DeviceConnect");
        public static AccessRights ServiceConnect_DeviceConnect { get; } = new AccessRights("ServiceConnect, DeviceConnect");
        public static AccessRights RegistryRead_RegistryWrite_ServiceConnect { get; } = new AccessRights("RegistryRead, RegistryWrite, ServiceConnect");
        public static AccessRights RegistryRead_RegistryWrite_DeviceConnect { get; } = new AccessRights("RegistryRead, RegistryWrite, DeviceConnect");
        public static AccessRights RegistryRead_ServiceConnect_DeviceConnect { get; } = new AccessRights("RegistryRead, ServiceConnect, DeviceConnect");
        public static AccessRights RegistryWrite_ServiceConnect_DeviceConnect { get; } = new AccessRights("RegistryWrite, ServiceConnect, DeviceConnect");
        public static AccessRights RegistryRead_RegistryWrite_ServiceConnect_DeviceConnect { get; } = new AccessRights("RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect");

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
    /// The capabilities and features enabled for the IoT hub.
    /// </summary>
    [EnumType]
    public readonly struct Capabilities : IEquatable<Capabilities>
    {
        private readonly string _value;

        private Capabilities(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Capabilities None { get; } = new Capabilities("None");
        public static Capabilities DeviceManagement { get; } = new Capabilities("DeviceManagement");

        public static bool operator ==(Capabilities left, Capabilities right) => left.Equals(right);
        public static bool operator !=(Capabilities left, Capabilities right) => !left.Equals(right);

        public static explicit operator string(Capabilities value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Capabilities other && Equals(other);
        public bool Equals(Capabilities other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The name of the SKU.
    /// </summary>
    [EnumType]
    public readonly struct IotHubSku : IEquatable<IotHubSku>
    {
        private readonly string _value;

        private IotHubSku(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IotHubSku F1 { get; } = new IotHubSku("F1");
        public static IotHubSku S1 { get; } = new IotHubSku("S1");
        public static IotHubSku S2 { get; } = new IotHubSku("S2");
        public static IotHubSku S3 { get; } = new IotHubSku("S3");
        public static IotHubSku B1 { get; } = new IotHubSku("B1");
        public static IotHubSku B2 { get; } = new IotHubSku("B2");
        public static IotHubSku B3 { get; } = new IotHubSku("B3");

        public static bool operator ==(IotHubSku left, IotHubSku right) => left.Equals(right);
        public static bool operator !=(IotHubSku left, IotHubSku right) => !left.Equals(right);

        public static explicit operator string(IotHubSku value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IotHubSku other && Equals(other);
        public bool Equals(IotHubSku other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The desired action for requests captured by this rule.
    /// </summary>
    [EnumType]
    public readonly struct IpFilterActionType : IEquatable<IpFilterActionType>
    {
        private readonly string _value;

        private IpFilterActionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static IpFilterActionType Accept { get; } = new IpFilterActionType("Accept");
        public static IpFilterActionType Reject { get; } = new IpFilterActionType("Reject");

        public static bool operator ==(IpFilterActionType left, IpFilterActionType right) => left.Equals(right);
        public static bool operator !=(IpFilterActionType left, IpFilterActionType right) => !left.Equals(right);

        public static explicit operator string(IpFilterActionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IpFilterActionType other && Equals(other);
        public bool Equals(IpFilterActionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The source that the routing rule is to be applied to, such as DeviceMessages.
    /// </summary>
    [EnumType]
    public readonly struct RoutingSource : IEquatable<RoutingSource>
    {
        private readonly string _value;

        private RoutingSource(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RoutingSource Invalid { get; } = new RoutingSource("Invalid");
        public static RoutingSource DeviceMessages { get; } = new RoutingSource("DeviceMessages");
        public static RoutingSource TwinChangeEvents { get; } = new RoutingSource("TwinChangeEvents");
        public static RoutingSource DeviceLifecycleEvents { get; } = new RoutingSource("DeviceLifecycleEvents");
        public static RoutingSource DeviceJobLifecycleEvents { get; } = new RoutingSource("DeviceJobLifecycleEvents");

        public static bool operator ==(RoutingSource left, RoutingSource right) => left.Equals(right);
        public static bool operator !=(RoutingSource left, RoutingSource right) => !left.Equals(right);

        public static explicit operator string(RoutingSource value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RoutingSource other && Equals(other);
        public bool Equals(RoutingSource other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
