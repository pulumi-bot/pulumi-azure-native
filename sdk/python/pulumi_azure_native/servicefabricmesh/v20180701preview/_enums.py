# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DiagnosticsSinkKind',
    'HealthState',
    'IngressQoSLevel',
    'OperatingSystemTypes',
    'VolumeProvider',
]


class DiagnosticsSinkKind(str, Enum):
    """
    The kind of DiagnosticsSink.
    """
    INVALID = "Invalid"
    AZURE_INTERNAL_MONITORING_PIPELINE = "AzureInternalMonitoringPipeline"


class HealthState(str, Enum):
    """
    The health state of a resource such as Application, Service, or Network.
    """
    INVALID = "Invalid"
    OK = "Ok"
    WARNING = "Warning"
    ERROR = "Error"
    UNKNOWN = "Unknown"


class IngressQoSLevel(str, Enum):
    """
    The QoS tier for ingress.
    """
    BRONZE = "Bronze"


class OperatingSystemTypes(str, Enum):
    """
    The Operating system type required by the code in service.
    """
    LINUX = "Linux"
    WINDOWS = "Windows"


class VolumeProvider(str, Enum):
    """
    Provider of the volume.
    """
    SF_AZURE_FILE = "SFAzureFile"
