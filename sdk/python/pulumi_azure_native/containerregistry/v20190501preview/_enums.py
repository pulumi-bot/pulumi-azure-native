# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'TokenCertificateName',
    'TokenPasswordName',
    'TokenStatus',
]


class TokenCertificateName(str, Enum):
    CERTIFICATE1 = "certificate1"
    CERTIFICATE2 = "certificate2"


class TokenPasswordName(str, Enum):
    """
    The password name "password1" or "password2"
    """
    PASSWORD1 = "password1"
    PASSWORD2 = "password2"


class TokenStatus(str, Enum):
    """
    The status of the token example enabled or disabled.
    """
    ENABLED = "enabled"
    DISABLED = "disabled"
