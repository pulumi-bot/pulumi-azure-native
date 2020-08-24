# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetWebhookResult',
    'AwaitableGetWebhookResult',
    'get_webhook',
]

@pulumi.output_type
class GetWebhookResult:
    """
    An object that represents a webhook for a container registry.
    """
    def __init__(__self__, actions=None, location=None, name=None, provisioning_state=None, scope=None, status=None, tags=None, type=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> List[str]:
        """
        The list of actions that trigger the webhook to post notifications.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the webhook at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The scope of repositories where the event can be triggered. For example, 'foo:*' means events for all tags under repository 'foo'. 'foo:bar' means events for 'foo:bar' only. 'foo' is equivalent to 'foo:latest'. Empty means all events.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the webhook at the time the operation was called.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetWebhookResult(GetWebhookResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebhookResult(
            actions=self.actions,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            scope=self.scope,
            status=self.status,
            tags=self.tags,
            type=self.type)


def get_webhook(name: Optional[str] = None,
                registry_name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebhookResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the webhook.
    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerregistry/v20190501:getWebhook', __args__, opts=opts, typ=GetWebhookResult).value

    return AwaitableGetWebhookResult(
        actions=__ret__.actions,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        scope=__ret__.scope,
        status=__ret__.status,
        tags=__ret__.tags,
        type=__ret__.type)
