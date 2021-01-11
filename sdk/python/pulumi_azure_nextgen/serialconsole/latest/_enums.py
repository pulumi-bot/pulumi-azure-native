# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'SerialPortState',
]


class SerialPortState(str, Enum):
    """
    Specifies whether the port is enabled for a serial console connection.
    """
    ENABLED = "enabled"
    DISABLED = "disabled"
