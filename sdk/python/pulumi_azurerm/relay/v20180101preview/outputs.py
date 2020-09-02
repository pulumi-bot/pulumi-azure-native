# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'SkuResponse',
]

@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    Private endpoint object properties.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        Private endpoint object properties.
        :param str id: Full identifier of the private endpoint resource.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Full identifier of the private endpoint resource.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    An object that represents the approval state of the private link connection.
    """
    def __init__(__self__, *,
                 action_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        An object that represents the approval state of the private link connection.
        :param str action_required: A message indicating if changes on the service provider require any updates on the consumer.
        :param str description: The reason for approval or rejection.
        :param str status: Indicates whether the connection has been approved, rejected or removed by the Relay Namespace owner.
        """
        if action_required is not None:
            pulumi.set(__self__, "action_required", action_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionRequired")
    def action_required(self) -> Optional[str]:
        """
        A message indicating if changes on the service provider require any updates on the consumer.
        """
        return pulumi.get(self, "action_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The reason for approval or rejection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates whether the connection has been approved, rejected or removed by the Relay Namespace owner.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    SKU of the namespace.
    """
    def __init__(__self__, *,
                 name: str,
                 tier: Optional[str] = None):
        """
        SKU of the namespace.
        :param str name: Name of this SKU.
        :param str tier: The tier of this SKU.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        """
        The tier of this SKU.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


