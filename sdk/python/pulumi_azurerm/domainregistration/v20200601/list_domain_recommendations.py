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
    'ListDomainRecommendationsResult',
    'AwaitableListDomainRecommendationsResult',
    'list_domain_recommendations',
]

@pulumi.output_type
class ListDomainRecommendationsResult:
    """
    Collection of domain name identifiers.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> str:
        """
        Link to next page of resources.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> List['outputs.NameIdentifierResponseResult']:
        """
        Collection of resources.
        """
        return pulumi.get(self, "value")


class AwaitableListDomainRecommendationsResult(ListDomainRecommendationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainRecommendationsResult(
            next_link=self.next_link,
            value=self.value)


def list_domain_recommendations(keywords: Optional[str] = None,
                                max_domain_recommendations: Optional[float] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainRecommendationsResult:
    """
    Use this data source to access information about an existing resource.

    :param str keywords: Keywords to be used for generating domain recommendations.
    :param float max_domain_recommendations: Maximum number of recommendations.
    """
    __args__ = dict()
    __args__['keywords'] = keywords
    __args__['maxDomainRecommendations'] = max_domain_recommendations
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:domainregistration/v20200601:listDomainRecommendations', __args__, opts=opts, typ=ListDomainRecommendationsResult).value

    return AwaitableListDomainRecommendationsResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
