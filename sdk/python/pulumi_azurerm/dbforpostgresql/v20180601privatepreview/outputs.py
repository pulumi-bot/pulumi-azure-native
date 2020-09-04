# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PrivateEndpointPropertyResponse',
    'PrivateLinkServiceConnectionStatePropertyResponse',
]

@pulumi.output_type
class PrivateEndpointPropertyResponse(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        :param str id: Resource id of the private endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource id of the private endpoint.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateLinkServiceConnectionStatePropertyResponse(dict):
    def __init__(__self__, *,
                 actions_required: str,
                 description: str,
                 status: str):
        """
        :param str actions_required: The actions required for private link service connection.
        :param str description: The private link service connection description.
        :param str status: The private link service connection status.
        """
        pulumi.set(__self__, "actions_required", actions_required)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> str:
        """
        The actions required for private link service connection.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The private link service connection description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The private link service connection status.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


