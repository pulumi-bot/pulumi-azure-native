# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ComplianceState',
]


class ComplianceState(str, Enum):
    """
    The compliance state that should be set on the resource.
    """
    COMPLIANT = "Compliant"
    NON_COMPLIANT = "NonCompliant"
    UNKNOWN = "Unknown"
