# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListGlobalDomainRegistrationDomainRecommendationsResult:
    """
    Collection of domain name identifiers
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        __self__.next_link = next_link
        """
        Link to next page of resources
        """
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        __self__.value = value
        """
        Collection of resources
        """


class AwaitableListGlobalDomainRegistrationDomainRecommendationsResult(ListGlobalDomainRegistrationDomainRecommendationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListGlobalDomainRegistrationDomainRecommendationsResult(
            next_link=self.next_link,
            value=self.value)


def list_global_domain_registration_domain_recommendations(keywords=None, max_domain_recommendations=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str keywords: Keywords to be used for generating domain recommendations
    :param float max_domain_recommendations: Maximum number of recommendations
    """
    __args__ = dict()
    __args__['keywords'] = keywords
    __args__['maxDomainRecommendations'] = max_domain_recommendations
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:domainregistration/v20150801:listGlobalDomainRegistrationDomainRecommendations', __args__, opts=opts).value

    return AwaitableListGlobalDomainRegistrationDomainRecommendationsResult(
        next_link=__ret__.get('nextLink'),
        value=__ret__.get('value'))