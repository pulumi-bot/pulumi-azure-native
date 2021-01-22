# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetProductsResult',
    'AwaitableGetProductsResult',
    'get_products',
]

@pulumi.output_type
class GetProductsResult:
    """
    Pageable list of products.
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
    def next_link(self) -> Optional[str]:
        """
        URI to the next page.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[Sequence['outputs.ProductResponseResult']]:
        """
        List of products.
        """
        return pulumi.get(self, "value")


class AwaitableGetProductsResult(GetProductsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductsResult(
            next_link=self.next_link,
            value=self.value)


def get_products(product_name: Optional[str] = None,
                 registration_name: Optional[str] = None,
                 resource_group: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductsResult:
    """
    Use this data source to access information about an existing resource.

    :param str product_name: Name of the product.
    :param str registration_name: Name of the Azure Stack registration.
    :param str resource_group: Name of the resource group.
    """
    __args__ = dict()
    __args__['productName'] = product_name
    __args__['registrationName'] = registration_name
    __args__['resourceGroup'] = resource_group
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:azurestack/latest:getProducts', __args__, opts=opts, typ=GetProductsResult).value

    return AwaitableGetProductsResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
