// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNextGen.Resources.Latest
{
    /// <summary>
    /// The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode, resources are deployed without deleting existing resources that are not included in the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included in the template are deleted. Be careful when using Complete mode as you may unintentionally delete resources.
    /// </summary>
    [EnumType]
    public readonly struct DeploymentMode : IEquatable<DeploymentMode>
    {
        private readonly string _value;

        private DeploymentMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DeploymentMode Incremental { get; } = new DeploymentMode("Incremental");
        public static DeploymentMode Complete { get; } = new DeploymentMode("Complete");

        public static bool operator ==(DeploymentMode left, DeploymentMode right) => left.Equals(right);
        public static bool operator !=(DeploymentMode left, DeploymentMode right) => !left.Equals(right);

        public static explicit operator string(DeploymentMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DeploymentMode other && Equals(other);
        public bool Equals(DeploymentMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The scope to be used for evaluation of parameters, variables and functions in a nested template.
    /// </summary>
    [EnumType]
    public readonly struct ExpressionEvaluationOptionsScopeType : IEquatable<ExpressionEvaluationOptionsScopeType>
    {
        private readonly string _value;

        private ExpressionEvaluationOptionsScopeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ExpressionEvaluationOptionsScopeType NotSpecified { get; } = new ExpressionEvaluationOptionsScopeType("NotSpecified");
        public static ExpressionEvaluationOptionsScopeType Outer { get; } = new ExpressionEvaluationOptionsScopeType("Outer");
        public static ExpressionEvaluationOptionsScopeType Inner { get; } = new ExpressionEvaluationOptionsScopeType("Inner");

        public static bool operator ==(ExpressionEvaluationOptionsScopeType left, ExpressionEvaluationOptionsScopeType right) => left.Equals(right);
        public static bool operator !=(ExpressionEvaluationOptionsScopeType left, ExpressionEvaluationOptionsScopeType right) => !left.Equals(right);

        public static explicit operator string(ExpressionEvaluationOptionsScopeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExpressionEvaluationOptionsScopeType other && Equals(other);
        public bool Equals(ExpressionEvaluationOptionsScopeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the managed identity.
    /// </summary>
    [EnumType]
    public readonly struct ManagedServiceIdentityType : IEquatable<ManagedServiceIdentityType>
    {
        private readonly string _value;

        private ManagedServiceIdentityType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedServiceIdentityType UserAssigned { get; } = new ManagedServiceIdentityType("UserAssigned");

        public static bool operator ==(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => left.Equals(right);
        public static bool operator !=(ManagedServiceIdentityType left, ManagedServiceIdentityType right) => !left.Equals(right);

        public static explicit operator string(ManagedServiceIdentityType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedServiceIdentityType other && Equals(other);
        public bool Equals(ManagedServiceIdentityType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
    /// </summary>
    [EnumType]
    public readonly struct OnErrorDeploymentType : IEquatable<OnErrorDeploymentType>
    {
        private readonly string _value;

        private OnErrorDeploymentType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static OnErrorDeploymentType LastSuccessful { get; } = new OnErrorDeploymentType("LastSuccessful");
        public static OnErrorDeploymentType SpecificDeployment { get; } = new OnErrorDeploymentType("SpecificDeployment");

        public static bool operator ==(OnErrorDeploymentType left, OnErrorDeploymentType right) => left.Equals(right);
        public static bool operator !=(OnErrorDeploymentType left, OnErrorDeploymentType right) => !left.Equals(right);

        public static explicit operator string(OnErrorDeploymentType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is OnErrorDeploymentType other && Equals(other);
        public bool Equals(OnErrorDeploymentType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

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
    /// Type of the script.
    /// </summary>
    [EnumType]
    public readonly struct ScriptType : IEquatable<ScriptType>
    {
        private readonly string _value;

        private ScriptType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScriptType AzurePowerShell { get; } = new ScriptType("AzurePowerShell");
        public static ScriptType AzureCLI { get; } = new ScriptType("AzureCLI");

        public static bool operator ==(ScriptType left, ScriptType right) => left.Equals(right);
        public static bool operator !=(ScriptType left, ScriptType right) => !left.Equals(right);

        public static explicit operator string(ScriptType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScriptType other && Equals(other);
        public bool Equals(ScriptType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
