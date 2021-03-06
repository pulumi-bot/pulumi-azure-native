# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'EnforcementMode',
    'ParameterType',
    'PolicyType',
    'ResourceIdentityType',
]


class EnforcementMode(str, Enum):
    """
    The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
    """
    DEFAULT = "Default"
    DO_NOT_ENFORCE = "DoNotEnforce"


class ParameterType(str, Enum):
    """
    The data type of the parameter.
    """
    STRING = "String"
    ARRAY = "Array"
    OBJECT = "Object"
    BOOLEAN = "Boolean"
    INTEGER = "Integer"
    FLOAT = "Float"
    DATE_TIME = "DateTime"


class PolicyType(str, Enum):
    """
    The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
    """
    NOT_SPECIFIED = "NotSpecified"
    BUILT_IN = "BuiltIn"
    CUSTOM = "Custom"
    STATIC = "Static"


class ResourceIdentityType(str, Enum):
    """
    The identity type. This is the only required field when adding a system assigned identity to a resource.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    NONE = "None"
