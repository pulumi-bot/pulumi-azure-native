# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetPrefixResult',
    'AwaitableGetPrefixResult',
    'get_prefix',
]

@pulumi.output_type
class GetPrefixResult:
    """
    The peering service prefix class.
    """
    def __init__(__self__, error_message=None, events=None, learned_type=None, name=None, prefix=None, prefix_validation_state=None, provisioning_state=None, type=None):
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if events and not isinstance(events, list):
            raise TypeError("Expected argument 'events' to be a list")
        pulumi.set(__self__, "events", events)
        if learned_type and not isinstance(learned_type, str):
            raise TypeError("Expected argument 'learned_type' to be a str")
        pulumi.set(__self__, "learned_type", learned_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if prefix and not isinstance(prefix, str):
            raise TypeError("Expected argument 'prefix' to be a str")
        pulumi.set(__self__, "prefix", prefix)
        if prefix_validation_state and not isinstance(prefix_validation_state, str):
            raise TypeError("Expected argument 'prefix_validation_state' to be a str")
        pulumi.set(__self__, "prefix_validation_state", prefix_validation_state)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> str:
        """
        The error message for validation state
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter
    def events(self) -> List['outputs.PeeringServicePrefixEventResponse']:
        """
        The list of events for peering service prefix
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter(name="learnedType")
    def learned_type(self) -> str:
        """
        The prefix learned type
        """
        return pulumi.get(self, "learned_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        The prefix from which your traffic originates.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="prefixValidationState")
    def prefix_validation_state(self) -> str:
        """
        The prefix validation state
        """
        return pulumi.get(self, "prefix_validation_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetPrefixResult(GetPrefixResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrefixResult(
            error_message=self.error_message,
            events=self.events,
            learned_type=self.learned_type,
            name=self.name,
            prefix=self.prefix,
            prefix_validation_state=self.prefix_validation_state,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_prefix(expand: Optional[str] = None,
               peering_service_name: Optional[str] = None,
               prefix_name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrefixResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: The properties to be expanded.
    :param str peering_service_name: The name of the peering service.
    :param str prefix_name: The name of the prefix.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['peeringServiceName'] = peering_service_name
    __args__['prefixName'] = prefix_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:peering/v20190901preview:getPrefix', __args__, opts=opts, typ=GetPrefixResult).value

    return AwaitableGetPrefixResult(
        error_message=__ret__.error_message,
        events=__ret__.events,
        learned_type=__ret__.learned_type,
        name=__ret__.name,
        prefix=__ret__.prefix,
        prefix_validation_state=__ret__.prefix_validation_state,
        provisioning_state=__ret__.provisioning_state,
        type=__ret__.type)
