# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetTopicAuthorizationRuleResult',
    'AwaitableGetTopicAuthorizationRuleResult',
    'get_topic_authorization_rule',
]

@pulumi.output_type
class GetTopicAuthorizationRuleResult:
    """
    Description of a namespace authorization rule.
    """
    def __init__(__self__, name=None, rights=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rights and not isinstance(rights, list):
            raise TypeError("Expected argument 'rights' to be a list")
        pulumi.set(__self__, "rights", rights)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rights(self) -> List[str]:
        """
        The rights associated with the rule.
        """
        return pulumi.get(self, "rights")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetTopicAuthorizationRuleResult(GetTopicAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicAuthorizationRuleResult(
            name=self.name,
            rights=self.rights,
            type=self.type)


def get_topic_authorization_rule(name: Optional[str] = None,
                                 namespace_name: Optional[str] = None,
                                 resource_group_name: Optional[str] = None,
                                 topic_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicAuthorizationRuleResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The authorization rule name.
    :param str namespace_name: The namespace name
    :param str resource_group_name: Name of the Resource group within the Azure subscription.
    :param str topic_name: The topic name.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:servicebus/v20170401:getTopicAuthorizationRule', __args__, opts=opts, typ=GetTopicAuthorizationRuleResult).value

    return AwaitableGetTopicAuthorizationRuleResult(
        name=__ret__.name,
        rights=__ret__.rights,
        type=__ret__.type)
