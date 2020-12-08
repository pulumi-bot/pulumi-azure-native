// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.ContainerService.V20190930Preview
{
    /// <summary>
    /// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
    /// </summary>
    [EnumType]
    public readonly struct OSType : IEquatable<OSType>
    {
        private readonly string _value;

        private OSType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OSType Linux { get; } = new OSType("Linux");
        public static OSType Windows { get; } = new OSType("Windows");

        public static bool operator ==(OSType left, OSType right) => left.Equals(right);
        public static bool operator !=(OSType left, OSType right) => !left.Equals(right);

        public static explicit operator string(OSType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OSType other && Equals(other);
        public bool Equals(OSType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Define the role of the AgentPoolProfile.
    /// </summary>
    [EnumType]
    public readonly struct OpenShiftAgentPoolProfileRole : IEquatable<OpenShiftAgentPoolProfileRole>
    {
        private readonly string _value;

        private OpenShiftAgentPoolProfileRole(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OpenShiftAgentPoolProfileRole Compute { get; } = new OpenShiftAgentPoolProfileRole("compute");
        public static OpenShiftAgentPoolProfileRole Infra { get; } = new OpenShiftAgentPoolProfileRole("infra");

        public static bool operator ==(OpenShiftAgentPoolProfileRole left, OpenShiftAgentPoolProfileRole right) => left.Equals(right);
        public static bool operator !=(OpenShiftAgentPoolProfileRole left, OpenShiftAgentPoolProfileRole right) => !left.Equals(right);

        public static explicit operator string(OpenShiftAgentPoolProfileRole value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OpenShiftAgentPoolProfileRole other && Equals(other);
        public bool Equals(OpenShiftAgentPoolProfileRole other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Size of agent VMs.
    /// </summary>
    [EnumType]
    public readonly struct OpenShiftContainerServiceVMSize : IEquatable<OpenShiftContainerServiceVMSize>
    {
        private readonly string _value;

        private OpenShiftContainerServiceVMSize(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OpenShiftContainerServiceVMSize Standard_D2s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D2s_v3");
        public static OpenShiftContainerServiceVMSize Standard_D4s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D4s_v3");
        public static OpenShiftContainerServiceVMSize Standard_D8s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D8s_v3");
        public static OpenShiftContainerServiceVMSize Standard_D16s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D16s_v3");
        public static OpenShiftContainerServiceVMSize Standard_D32s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D32s_v3");
        public static OpenShiftContainerServiceVMSize Standard_D64s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_D64s_v3");
        public static OpenShiftContainerServiceVMSize Standard_DS4_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS4_v2");
        public static OpenShiftContainerServiceVMSize Standard_DS5_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS5_v2");
        public static OpenShiftContainerServiceVMSize Standard_F8s_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_F8s_v2");
        public static OpenShiftContainerServiceVMSize Standard_F16s_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_F16s_v2");
        public static OpenShiftContainerServiceVMSize Standard_F32s_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_F32s_v2");
        public static OpenShiftContainerServiceVMSize Standard_F64s_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_F64s_v2");
        public static OpenShiftContainerServiceVMSize Standard_F72s_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_F72s_v2");
        public static OpenShiftContainerServiceVMSize Standard_F8s { get; } = new OpenShiftContainerServiceVMSize("Standard_F8s");
        public static OpenShiftContainerServiceVMSize Standard_F16s { get; } = new OpenShiftContainerServiceVMSize("Standard_F16s");
        public static OpenShiftContainerServiceVMSize Standard_E4s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E4s_v3");
        public static OpenShiftContainerServiceVMSize Standard_E8s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E8s_v3");
        public static OpenShiftContainerServiceVMSize Standard_E16s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E16s_v3");
        public static OpenShiftContainerServiceVMSize Standard_E20s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E20s_v3");
        public static OpenShiftContainerServiceVMSize Standard_E32s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E32s_v3");
        public static OpenShiftContainerServiceVMSize Standard_E64s_v3 { get; } = new OpenShiftContainerServiceVMSize("Standard_E64s_v3");
        public static OpenShiftContainerServiceVMSize Standard_GS2 { get; } = new OpenShiftContainerServiceVMSize("Standard_GS2");
        public static OpenShiftContainerServiceVMSize Standard_GS3 { get; } = new OpenShiftContainerServiceVMSize("Standard_GS3");
        public static OpenShiftContainerServiceVMSize Standard_GS4 { get; } = new OpenShiftContainerServiceVMSize("Standard_GS4");
        public static OpenShiftContainerServiceVMSize Standard_GS5 { get; } = new OpenShiftContainerServiceVMSize("Standard_GS5");
        public static OpenShiftContainerServiceVMSize Standard_DS12_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS12_v2");
        public static OpenShiftContainerServiceVMSize Standard_DS13_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS13_v2");
        public static OpenShiftContainerServiceVMSize Standard_DS14_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS14_v2");
        public static OpenShiftContainerServiceVMSize Standard_DS15_v2 { get; } = new OpenShiftContainerServiceVMSize("Standard_DS15_v2");
        public static OpenShiftContainerServiceVMSize Standard_L4s { get; } = new OpenShiftContainerServiceVMSize("Standard_L4s");
        public static OpenShiftContainerServiceVMSize Standard_L8s { get; } = new OpenShiftContainerServiceVMSize("Standard_L8s");
        public static OpenShiftContainerServiceVMSize Standard_L16s { get; } = new OpenShiftContainerServiceVMSize("Standard_L16s");
        public static OpenShiftContainerServiceVMSize Standard_L32s { get; } = new OpenShiftContainerServiceVMSize("Standard_L32s");

        public static bool operator ==(OpenShiftContainerServiceVMSize left, OpenShiftContainerServiceVMSize right) => left.Equals(right);
        public static bool operator !=(OpenShiftContainerServiceVMSize left, OpenShiftContainerServiceVMSize right) => !left.Equals(right);

        public static explicit operator string(OpenShiftContainerServiceVMSize value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OpenShiftContainerServiceVMSize other && Equals(other);
        public bool Equals(OpenShiftContainerServiceVMSize other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
