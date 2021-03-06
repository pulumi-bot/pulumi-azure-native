# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'QueryStringCachingBehavior',
    'SkuName',
]


class QueryStringCachingBehavior(str, Enum):
    """
    Defines the query string caching behavior.
    """
    IGNORE_QUERY_STRING = "IgnoreQueryString"
    BYPASS_CACHING = "BypassCaching"
    USE_QUERY_STRING = "UseQueryString"
    NOT_SET = "NotSet"


class SkuName(str, Enum):
    """
    Name of the pricing tier
    """
    STANDARD_VERIZON = "Standard_Verizon"
    PREMIUM_VERIZON = "Premium_Verizon"
    CUSTOM_VERIZON = "Custom_Verizon"
    STANDARD_AKAMAI = "Standard_Akamai"
