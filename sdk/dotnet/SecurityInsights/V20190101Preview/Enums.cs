// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.SecurityInsights.V20190101Preview
{
    /// <summary>
    /// The kind of the alert rule
    /// </summary>
    [EnumType]
    public readonly struct AlertRuleKind : IEquatable<AlertRuleKind>
    {
        private readonly string _value;

        private AlertRuleKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AlertRuleKind Scheduled { get; } = new AlertRuleKind("Scheduled");
        public static AlertRuleKind MicrosoftSecurityIncidentCreation { get; } = new AlertRuleKind("MicrosoftSecurityIncidentCreation");
        public static AlertRuleKind Fusion { get; } = new AlertRuleKind("Fusion");
        public static AlertRuleKind MLBehaviorAnalytics { get; } = new AlertRuleKind("MLBehaviorAnalytics");
        public static AlertRuleKind ThreatIntelligence { get; } = new AlertRuleKind("ThreatIntelligence");
        public static AlertRuleKind Anomaly { get; } = new AlertRuleKind("Anomaly");

        public static bool operator ==(AlertRuleKind left, AlertRuleKind right) => left.Equals(right);
        public static bool operator !=(AlertRuleKind left, AlertRuleKind right) => !left.Equals(right);

        public static explicit operator string(AlertRuleKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AlertRuleKind other && Equals(other);
        public bool Equals(AlertRuleKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The severity of the incident
    /// </summary>
    [EnumType]
    public readonly struct CaseSeverity : IEquatable<CaseSeverity>
    {
        private readonly string _value;

        private CaseSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Critical severity
        /// </summary>
        public static CaseSeverity Critical { get; } = new CaseSeverity("Critical");
        /// <summary>
        /// High severity
        /// </summary>
        public static CaseSeverity High { get; } = new CaseSeverity("High");
        /// <summary>
        /// Medium severity
        /// </summary>
        public static CaseSeverity Medium { get; } = new CaseSeverity("Medium");
        /// <summary>
        /// Low severity
        /// </summary>
        public static CaseSeverity Low { get; } = new CaseSeverity("Low");
        /// <summary>
        /// Informational severity
        /// </summary>
        public static CaseSeverity Informational { get; } = new CaseSeverity("Informational");

        public static bool operator ==(CaseSeverity left, CaseSeverity right) => left.Equals(right);
        public static bool operator !=(CaseSeverity left, CaseSeverity right) => !left.Equals(right);

        public static explicit operator string(CaseSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CaseSeverity other && Equals(other);
        public bool Equals(CaseSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the data connector
    /// </summary>
    [EnumType]
    public readonly struct DataConnectorKind : IEquatable<DataConnectorKind>
    {
        private readonly string _value;

        private DataConnectorKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DataConnectorKind AzureActiveDirectory { get; } = new DataConnectorKind("AzureActiveDirectory");
        public static DataConnectorKind AzureSecurityCenter { get; } = new DataConnectorKind("AzureSecurityCenter");
        public static DataConnectorKind MicrosoftCloudAppSecurity { get; } = new DataConnectorKind("MicrosoftCloudAppSecurity");
        public static DataConnectorKind ThreatIntelligence { get; } = new DataConnectorKind("ThreatIntelligence");
        public static DataConnectorKind ThreatIntelligenceTaxii { get; } = new DataConnectorKind("ThreatIntelligenceTaxii");
        public static DataConnectorKind Office365 { get; } = new DataConnectorKind("Office365");
        public static DataConnectorKind OfficeATP { get; } = new DataConnectorKind("OfficeATP");
        public static DataConnectorKind AmazonWebServicesCloudTrail { get; } = new DataConnectorKind("AmazonWebServicesCloudTrail");
        public static DataConnectorKind AzureAdvancedThreatProtection { get; } = new DataConnectorKind("AzureAdvancedThreatProtection");
        public static DataConnectorKind MicrosoftDefenderAdvancedThreatProtection { get; } = new DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection");
        public static DataConnectorKind Dynamics365 { get; } = new DataConnectorKind("Dynamics365");
        public static DataConnectorKind MicrosoftThreatProtection { get; } = new DataConnectorKind("MicrosoftThreatProtection");
        public static DataConnectorKind MicrosoftThreatIntelligence { get; } = new DataConnectorKind("MicrosoftThreatIntelligence");

        public static bool operator ==(DataConnectorKind left, DataConnectorKind right) => left.Equals(right);
        public static bool operator !=(DataConnectorKind left, DataConnectorKind right) => !left.Equals(right);

        public static explicit operator string(DataConnectorKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DataConnectorKind other && Equals(other);
        public bool Equals(DataConnectorKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The entity query kind
    /// </summary>
    [EnumType]
    public readonly struct EntityTimelineKind : IEquatable<EntityTimelineKind>
    {
        private readonly string _value;

        private EntityTimelineKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// activity
        /// </summary>
        public static EntityTimelineKind Activity { get; } = new EntityTimelineKind("Activity");
        /// <summary>
        /// bookmarks
        /// </summary>
        public static EntityTimelineKind Bookmark { get; } = new EntityTimelineKind("Bookmark");
        /// <summary>
        /// security alerts
        /// </summary>
        public static EntityTimelineKind SecurityAlert { get; } = new EntityTimelineKind("SecurityAlert");

        public static bool operator ==(EntityTimelineKind left, EntityTimelineKind right) => left.Equals(right);
        public static bool operator !=(EntityTimelineKind left, EntityTimelineKind right) => !left.Equals(right);

        public static explicit operator string(EntityTimelineKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is EntityTimelineKind other && Equals(other);
        public bool Equals(EntityTimelineKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The reason the incident was closed
    /// </summary>
    [EnumType]
    public readonly struct IncidentClassification : IEquatable<IncidentClassification>
    {
        private readonly string _value;

        private IncidentClassification(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Incident classification was undetermined
        /// </summary>
        public static IncidentClassification Undetermined { get; } = new IncidentClassification("Undetermined");
        /// <summary>
        /// Incident was true positive
        /// </summary>
        public static IncidentClassification TruePositive { get; } = new IncidentClassification("TruePositive");
        /// <summary>
        /// Incident was benign positive
        /// </summary>
        public static IncidentClassification BenignPositive { get; } = new IncidentClassification("BenignPositive");
        /// <summary>
        /// Incident was false positive
        /// </summary>
        public static IncidentClassification FalsePositive { get; } = new IncidentClassification("FalsePositive");

        public static bool operator ==(IncidentClassification left, IncidentClassification right) => left.Equals(right);
        public static bool operator !=(IncidentClassification left, IncidentClassification right) => !left.Equals(right);

        public static explicit operator string(IncidentClassification value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentClassification other && Equals(other);
        public bool Equals(IncidentClassification other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The classification reason the incident was closed with
    /// </summary>
    [EnumType]
    public readonly struct IncidentClassificationReason : IEquatable<IncidentClassificationReason>
    {
        private readonly string _value;

        private IncidentClassificationReason(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Classification reason was suspicious activity
        /// </summary>
        public static IncidentClassificationReason SuspiciousActivity { get; } = new IncidentClassificationReason("SuspiciousActivity");
        /// <summary>
        /// Classification reason was suspicious but expected
        /// </summary>
        public static IncidentClassificationReason SuspiciousButExpected { get; } = new IncidentClassificationReason("SuspiciousButExpected");
        /// <summary>
        /// Classification reason was incorrect alert logic
        /// </summary>
        public static IncidentClassificationReason IncorrectAlertLogic { get; } = new IncidentClassificationReason("IncorrectAlertLogic");
        /// <summary>
        /// Classification reason was inaccurate data
        /// </summary>
        public static IncidentClassificationReason InaccurateData { get; } = new IncidentClassificationReason("InaccurateData");

        public static bool operator ==(IncidentClassificationReason left, IncidentClassificationReason right) => left.Equals(right);
        public static bool operator !=(IncidentClassificationReason left, IncidentClassificationReason right) => !left.Equals(right);

        public static explicit operator string(IncidentClassificationReason value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentClassificationReason other && Equals(other);
        public bool Equals(IncidentClassificationReason other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The severity of the incident
    /// </summary>
    [EnumType]
    public readonly struct IncidentSeverity : IEquatable<IncidentSeverity>
    {
        private readonly string _value;

        private IncidentSeverity(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// High severity
        /// </summary>
        public static IncidentSeverity High { get; } = new IncidentSeverity("High");
        /// <summary>
        /// Medium severity
        /// </summary>
        public static IncidentSeverity Medium { get; } = new IncidentSeverity("Medium");
        /// <summary>
        /// Low severity
        /// </summary>
        public static IncidentSeverity Low { get; } = new IncidentSeverity("Low");
        /// <summary>
        /// Informational severity
        /// </summary>
        public static IncidentSeverity Informational { get; } = new IncidentSeverity("Informational");

        public static bool operator ==(IncidentSeverity left, IncidentSeverity right) => left.Equals(right);
        public static bool operator !=(IncidentSeverity left, IncidentSeverity right) => !left.Equals(right);

        public static explicit operator string(IncidentSeverity value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentSeverity other && Equals(other);
        public bool Equals(IncidentSeverity other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The status of the incident
    /// </summary>
    [EnumType]
    public readonly struct IncidentStatus : IEquatable<IncidentStatus>
    {
        private readonly string _value;

        private IncidentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An active incident which isn't being handled currently
        /// </summary>
        public static IncidentStatus New { get; } = new IncidentStatus("New");
        /// <summary>
        /// An active incident which is being handled
        /// </summary>
        public static IncidentStatus Active { get; } = new IncidentStatus("Active");
        /// <summary>
        /// A non-active incident
        /// </summary>
        public static IncidentStatus Closed { get; } = new IncidentStatus("Closed");

        public static bool operator ==(IncidentStatus left, IncidentStatus right) => left.Equals(right);
        public static bool operator !=(IncidentStatus left, IncidentStatus right) => !left.Equals(right);

        public static explicit operator string(IncidentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is IncidentStatus other && Equals(other);
        public bool Equals(IncidentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the setting
    /// </summary>
    [EnumType]
    public readonly struct SettingKind : IEquatable<SettingKind>
    {
        private readonly string _value;

        private SettingKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SettingKind EyesOn { get; } = new SettingKind("EyesOn");
        public static SettingKind EntityAnalytics { get; } = new SettingKind("EntityAnalytics");
        public static SettingKind Ueba { get; } = new SettingKind("Ueba");

        public static bool operator ==(SettingKind left, SettingKind right) => left.Equals(right);
        public static bool operator !=(SettingKind left, SettingKind right) => !left.Equals(right);

        public static explicit operator string(SettingKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SettingKind other && Equals(other);
        public bool Equals(SettingKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The source of the watchlist
    /// </summary>
    [EnumType]
    public readonly struct Source : IEquatable<Source>
    {
        private readonly string _value;

        private Source(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Source Local_file { get; } = new Source("Local file");
        public static Source Remote_storage { get; } = new Source("Remote storage");

        public static bool operator ==(Source left, Source right) => left.Equals(right);
        public static bool operator !=(Source left, Source right) => !left.Equals(right);

        public static explicit operator string(Source value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Source other && Equals(other);
        public bool Equals(Source other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The kind of the entity.
    /// </summary>
    [EnumType]
    public readonly struct ThreatIntelligenceResourceKind : IEquatable<ThreatIntelligenceResourceKind>
    {
        private readonly string _value;

        private ThreatIntelligenceResourceKind(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Entity represents threat intelligence indicator in the system.
        /// </summary>
        public static ThreatIntelligenceResourceKind Indicator { get; } = new ThreatIntelligenceResourceKind("indicator");

        public static bool operator ==(ThreatIntelligenceResourceKind left, ThreatIntelligenceResourceKind right) => left.Equals(right);
        public static bool operator !=(ThreatIntelligenceResourceKind left, ThreatIntelligenceResourceKind right) => !left.Equals(right);

        public static explicit operator string(ThreatIntelligenceResourceKind value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ThreatIntelligenceResourceKind other && Equals(other);
        public bool Equals(ThreatIntelligenceResourceKind other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
