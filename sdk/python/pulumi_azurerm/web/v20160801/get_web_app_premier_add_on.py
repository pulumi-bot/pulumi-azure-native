# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetWebAppPremierAddOnResult',
    'AwaitableGetWebAppPremierAddOnResult',
    'get_web_app_premier_add_on',
]

@pulumi.output_type
class GetWebAppPremierAddOnResult:
    """
    Premier add-on.
    """
    def __init__(__self__, kind=None, location=None, marketplace_offer=None, marketplace_publisher=None, name=None, premier_add_on_name=None, product=None, sku=None, tags=None, type=None, vendor=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if marketplace_offer and not isinstance(marketplace_offer, str):
            raise TypeError("Expected argument 'marketplace_offer' to be a str")
        pulumi.set(__self__, "marketplace_offer", marketplace_offer)
        if marketplace_publisher and not isinstance(marketplace_publisher, str):
            raise TypeError("Expected argument 'marketplace_publisher' to be a str")
        pulumi.set(__self__, "marketplace_publisher", marketplace_publisher)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if premier_add_on_name and not isinstance(premier_add_on_name, str):
            raise TypeError("Expected argument 'premier_add_on_name' to be a str")
        pulumi.set(__self__, "premier_add_on_name", premier_add_on_name)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vendor and not isinstance(vendor, str):
            raise TypeError("Expected argument 'vendor' to be a str")
        pulumi.set(__self__, "vendor", vendor)

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="marketplaceOffer")
    def marketplace_offer(self) -> Optional[str]:
        """
        Premier add on Marketplace offer.
        """
        return pulumi.get(self, "marketplace_offer")

    @property
    @pulumi.getter(name="marketplacePublisher")
    def marketplace_publisher(self) -> Optional[str]:
        """
        Premier add on Marketplace publisher.
        """
        return pulumi.get(self, "marketplace_publisher")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="premierAddOnName")
    def premier_add_on_name(self) -> Optional[str]:
        """
        Premier add on Name.
        """
        return pulumi.get(self, "premier_add_on_name")

    @property
    @pulumi.getter
    def product(self) -> Optional[str]:
        """
        Premier add on Product.
        """
        return pulumi.get(self, "product")

    @property
    @pulumi.getter
    def sku(self) -> Optional[str]:
        """
        Premier add on SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vendor(self) -> Optional[str]:
        """
        Premier add on Vendor.
        """
        return pulumi.get(self, "vendor")


class AwaitableGetWebAppPremierAddOnResult(GetWebAppPremierAddOnResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppPremierAddOnResult(
            kind=self.kind,
            location=self.location,
            marketplace_offer=self.marketplace_offer,
            marketplace_publisher=self.marketplace_publisher,
            name=self.name,
            premier_add_on_name=self.premier_add_on_name,
            product=self.product,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            vendor=self.vendor)


def get_web_app_premier_add_on(name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppPremierAddOnResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Add-on name.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:web/v20160801:getWebAppPremierAddOn', __args__, opts=opts, typ=GetWebAppPremierAddOnResult).value

    return AwaitableGetWebAppPremierAddOnResult(
        kind=__ret__.kind,
        location=__ret__.location,
        marketplace_offer=__ret__.marketplace_offer,
        marketplace_publisher=__ret__.marketplace_publisher,
        name=__ret__.name,
        premier_add_on_name=__ret__.premier_add_on_name,
        product=__ret__.product,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type,
        vendor=__ret__.vendor)
