# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PolicyType',
]


class PolicyType(str, Enum):
    """
    The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
    """
    NOT_SPECIFIED = "NotSpecified"
    BUILT_IN = "BuiltIn"
    CUSTOM = "Custom"
