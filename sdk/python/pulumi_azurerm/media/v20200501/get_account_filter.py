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
    'GetAccountFilterResult',
    'AwaitableGetAccountFilterResult',
    'get_account_filter',
]

@pulumi.output_type
class GetAccountFilterResult:
    """
    An Account Filter.
    """
    def __init__(__self__, first_quality=None, name=None, presentation_time_range=None, tracks=None, type=None):
        if first_quality and not isinstance(first_quality, dict):
            raise TypeError("Expected argument 'first_quality' to be a dict")
        pulumi.set(__self__, "first_quality", first_quality)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if presentation_time_range and not isinstance(presentation_time_range, dict):
            raise TypeError("Expected argument 'presentation_time_range' to be a dict")
        pulumi.set(__self__, "presentation_time_range", presentation_time_range)
        if tracks and not isinstance(tracks, list):
            raise TypeError("Expected argument 'tracks' to be a list")
        pulumi.set(__self__, "tracks", tracks)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="firstQuality")
    def first_quality(self) -> Optional['outputs.FirstQualityResponse']:
        """
        The first quality.
        """
        return pulumi.get(self, "first_quality")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="presentationTimeRange")
    def presentation_time_range(self) -> Optional['outputs.PresentationTimeRangeResponse']:
        """
        The presentation time range.
        """
        return pulumi.get(self, "presentation_time_range")

    @property
    @pulumi.getter
    def tracks(self) -> Optional[List['outputs.FilterTrackSelectionResponse']]:
        """
        The tracks selection conditions.
        """
        return pulumi.get(self, "tracks")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetAccountFilterResult(GetAccountFilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountFilterResult(
            first_quality=self.first_quality,
            name=self.name,
            presentation_time_range=self.presentation_time_range,
            tracks=self.tracks,
            type=self.type)


def get_account_filter(account_name: Optional[str] = None,
                       name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountFilterResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The Media Services account name.
    :param str name: The Account Filter name
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:media/v20200501:getAccountFilter', __args__, opts=opts, typ=GetAccountFilterResult).value

    return AwaitableGetAccountFilterResult(
        first_quality=__ret__.first_quality,
        name=__ret__.name,
        presentation_time_range=__ret__.presentation_time_range,
        tracks=__ret__.tracks,
        type=__ret__.type)
