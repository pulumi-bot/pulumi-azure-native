# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListAssetContainerSasResult:
    """
    The Asset Storage container SAS URLs.
    """
    def __init__(__self__, asset_container_sas_urls=None):
        if asset_container_sas_urls and not isinstance(asset_container_sas_urls, list):
            raise TypeError("Expected argument 'asset_container_sas_urls' to be a list")
        __self__.asset_container_sas_urls = asset_container_sas_urls
        """
        The list of Asset container SAS URLs.
        """


class AwaitableListAssetContainerSasResult(ListAssetContainerSasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAssetContainerSasResult(
            asset_container_sas_urls=self.asset_container_sas_urls)


def list_asset_container_sas(account_name=None, asset_name=None, expiry_time=None, permissions=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The Media Services account name.
    :param str asset_name: The Asset name.
    :param str expiry_time: The SAS URL expiration time.  This must be less than 24 hours from the current time.
    :param str permissions: The permissions to set on the SAS URL.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['assetName'] = asset_name
    __args__['expiryTime'] = expiry_time
    __args__['permissions'] = permissions
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:media/v20180701:listAssetContainerSas', __args__, opts=opts).value

    return AwaitableListAssetContainerSasResult(
        asset_container_sas_urls=__ret__.get('assetContainerSasUrls'))
