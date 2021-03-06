# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AdditionalWorkspaceDataType',
    'AdditionalWorkspaceType',
    'AssessmentStatusCode',
    'AssessmentType',
    'Category',
    'DataSource',
    'ExportData',
    'ImplementationEffort',
    'Protocol',
    'RecommendationConfigStatus',
    'RecommendationType',
    'SecuritySolutionStatus',
    'Severity',
    'Source',
    'Status',
    'StatusReason',
    'Threats',
    'UnmaskedIpLoggingStatus',
    'UserImpact',
]


class AdditionalWorkspaceDataType(str, Enum):
    """
    Data types sent to workspace.
    """
    ALERTS = "Alerts"
    RAW_EVENTS = "RawEvents"


class AdditionalWorkspaceType(str, Enum):
    """
    Workspace type.
    """
    SENTINEL = "Sentinel"


class AssessmentStatusCode(str, Enum):
    """
    Programmatic code for the status of the assessment
    """
    HEALTHY = "Healthy"
    UNHEALTHY = "Unhealthy"
    NOT_APPLICABLE = "NotApplicable"


class AssessmentType(str, Enum):
    """
    BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
    """
    BUILT_IN = "BuiltIn"
    CUSTOM_POLICY = "CustomPolicy"
    CUSTOMER_MANAGED = "CustomerManaged"
    VERIFIED_PARTNER = "VerifiedPartner"


class Category(str, Enum):
    """
    The category of resource that is at risk when the assessment is unhealthy
    """
    COMPUTE = "Compute"
    NETWORKING = "Networking"
    DATA = "Data"
    IDENTITY_AND_ACCESS = "IdentityAndAccess"
    IO_T = "IoT"


class DataSource(str, Enum):
    TWIN_DATA = "TwinData"


class ExportData(str, Enum):
    RAW_EVENTS = "RawEvents"


class ImplementationEffort(str, Enum):
    """
    The implementation effort required to remediate this assessment
    """
    LOW = "Low"
    MODERATE = "Moderate"
    HIGH = "High"


class Protocol(str, Enum):
    TCP = "TCP"
    UDP = "UDP"
    ALL = "*"


class RecommendationConfigStatus(str, Enum):
    """
    Recommendation status. When the recommendation status is disabled recommendations are not generated.
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class RecommendationType(str, Enum):
    """
    The type of IoT Security recommendation.
    """
    IO_T_ACR_AUTHENTICATION = "IoT_ACRAuthentication"
    IO_T_AGENT_SENDS_UNUTILIZED_MESSAGES = "IoT_AgentSendsUnutilizedMessages"
    IO_T_BASELINE = "IoT_Baseline"
    IO_T_EDGE_HUB_MEM_OPTIMIZE = "IoT_EdgeHubMemOptimize"
    IO_T_EDGE_LOGGING_OPTIONS = "IoT_EdgeLoggingOptions"
    IO_T_INCONSISTENT_MODULE_SETTINGS = "IoT_InconsistentModuleSettings"
    IO_T_INSTALL_AGENT = "IoT_InstallAgent"
    IO_T_IP_FILTER_DENY_ALL = "IoT_IPFilter_DenyAll"
    IO_T_IP_FILTER_PERMISSIVE_RULE = "IoT_IPFilter_PermissiveRule"
    IO_T_OPEN_PORTS = "IoT_OpenPorts"
    IO_T_PERMISSIVE_FIREWALL_POLICY = "IoT_PermissiveFirewallPolicy"
    IO_T_PERMISSIVE_INPUT_FIREWALL_RULES = "IoT_PermissiveInputFirewallRules"
    IO_T_PERMISSIVE_OUTPUT_FIREWALL_RULES = "IoT_PermissiveOutputFirewallRules"
    IO_T_PRIVILEGED_DOCKER_OPTIONS = "IoT_PrivilegedDockerOptions"
    IO_T_SHARED_CREDENTIALS = "IoT_SharedCredentials"
    IO_T_VULNERABLE_TLS_CIPHER_SUITE = "IoT_VulnerableTLSCipherSuite"


class SecuritySolutionStatus(str, Enum):
    """
    Status of the IoT Security solution.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class Severity(str, Enum):
    """
    The severity level of the assessment
    """
    LOW = "Low"
    MEDIUM = "Medium"
    HIGH = "High"


class Source(str, Enum):
    """
    The platform where the assessed resource resides
    """
    AZURE = "Azure"
    ON_PREMISE = "OnPremise"
    ON_PREMISE_SQL = "OnPremiseSql"


class Status(str, Enum):
    """
    The status of the port
    """
    REVOKED = "Revoked"
    INITIATED = "Initiated"


class StatusReason(str, Enum):
    """
    A description of why the `status` has its value
    """
    EXPIRED = "Expired"
    USER_REQUESTED = "UserRequested"
    NEWER_REQUEST_INITIATED = "NewerRequestInitiated"


class Threats(str, Enum):
    """
    Threats impact of the assessment
    """
    ACCOUNT_BREACH = "accountBreach"
    DATA_EXFILTRATION = "dataExfiltration"
    DATA_SPILLAGE = "dataSpillage"
    MALICIOUS_INSIDER = "maliciousInsider"
    ELEVATION_OF_PRIVILEGE = "elevationOfPrivilege"
    THREAT_RESISTANCE = "threatResistance"
    MISSING_COVERAGE = "missingCoverage"
    DENIAL_OF_SERVICE = "denialOfService"


class UnmaskedIpLoggingStatus(str, Enum):
    """
    Unmasked IP address logging status
    """
    DISABLED = "Disabled"
    ENABLED = "Enabled"


class UserImpact(str, Enum):
    """
    The user impact of the assessment
    """
    LOW = "Low"
    MODERATE = "Moderate"
    HIGH = "High"
