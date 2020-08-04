# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetBudgetResult:
    """
    A budget resource.
    """
    def __init__(__self__, e_tag=None, name=None, properties=None, type=None):
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        __self__.e_tag = e_tag
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the budget.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableGetBudgetResult(GetBudgetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBudgetResult(
            e_tag=self.e_tag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_budget(name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: Budget Name.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:consumption/v20180831:getBudget', __args__, opts=opts).value

    return AwaitableGetBudgetResult(
        e_tag=__ret__.get('eTag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
