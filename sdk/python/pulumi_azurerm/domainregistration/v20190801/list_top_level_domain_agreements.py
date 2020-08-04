# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListTopLevelDomainAgreementsResult:
    """
    Collection of top-level domain legal agreements.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        __self__.next_link = next_link
        """
        Link to next page of resources.
        """
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        __self__.value = value
        """
        Collection of resources.
        """


class AwaitableListTopLevelDomainAgreementsResult(ListTopLevelDomainAgreementsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListTopLevelDomainAgreementsResult(
            next_link=self.next_link,
            value=self.value)


def list_top_level_domain_agreements(for_transfer=None, include_privacy=None, name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param bool for_transfer: If <code>true</code>, then the list of agreements will include agreements for domain transfer as well; otherwise, <code>false</code>.
    :param bool include_privacy: If <code>true</code>, then the list of agreements will include agreements for domain privacy as well; otherwise, <code>false</code>.
    :param str name: Name of the top-level domain.
    """
    __args__ = dict()
    __args__['forTransfer'] = for_transfer
    __args__['includePrivacy'] = include_privacy
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:domainregistration/v20190801:listTopLevelDomainAgreements', __args__, opts=opts).value

    return AwaitableListTopLevelDomainAgreementsResult(
        next_link=__ret__.get('nextLink'),
        value=__ret__.get('value'))
