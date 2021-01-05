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
    'GetStreamingEndpointResult',
    'AwaitableGetStreamingEndpointResult',
    'get_streaming_endpoint',
]

@pulumi.output_type
class GetStreamingEndpointResult:
    """
    The StreamingEndpoint.
    """
    def __init__(__self__, access_control=None, availability_set_name=None, cdn_enabled=None, cdn_profile=None, cdn_provider=None, created=None, cross_site_access_policies=None, custom_host_names=None, description=None, free_trial_end_time=None, host_name=None, id=None, last_modified=None, location=None, max_cache_age=None, name=None, provisioning_state=None, resource_state=None, scale_units=None, tags=None, type=None):
        if access_control and not isinstance(access_control, dict):
            raise TypeError("Expected argument 'access_control' to be a dict")
        pulumi.set(__self__, "access_control", access_control)
        if availability_set_name and not isinstance(availability_set_name, str):
            raise TypeError("Expected argument 'availability_set_name' to be a str")
        pulumi.set(__self__, "availability_set_name", availability_set_name)
        if cdn_enabled and not isinstance(cdn_enabled, bool):
            raise TypeError("Expected argument 'cdn_enabled' to be a bool")
        pulumi.set(__self__, "cdn_enabled", cdn_enabled)
        if cdn_profile and not isinstance(cdn_profile, str):
            raise TypeError("Expected argument 'cdn_profile' to be a str")
        pulumi.set(__self__, "cdn_profile", cdn_profile)
        if cdn_provider and not isinstance(cdn_provider, str):
            raise TypeError("Expected argument 'cdn_provider' to be a str")
        pulumi.set(__self__, "cdn_provider", cdn_provider)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if cross_site_access_policies and not isinstance(cross_site_access_policies, dict):
            raise TypeError("Expected argument 'cross_site_access_policies' to be a dict")
        pulumi.set(__self__, "cross_site_access_policies", cross_site_access_policies)
        if custom_host_names and not isinstance(custom_host_names, list):
            raise TypeError("Expected argument 'custom_host_names' to be a list")
        pulumi.set(__self__, "custom_host_names", custom_host_names)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if free_trial_end_time and not isinstance(free_trial_end_time, str):
            raise TypeError("Expected argument 'free_trial_end_time' to be a str")
        pulumi.set(__self__, "free_trial_end_time", free_trial_end_time)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_modified and not isinstance(last_modified, str):
            raise TypeError("Expected argument 'last_modified' to be a str")
        pulumi.set(__self__, "last_modified", last_modified)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_cache_age and not isinstance(max_cache_age, float):
            raise TypeError("Expected argument 'max_cache_age' to be a float")
        pulumi.set(__self__, "max_cache_age", max_cache_age)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_state and not isinstance(resource_state, str):
            raise TypeError("Expected argument 'resource_state' to be a str")
        pulumi.set(__self__, "resource_state", resource_state)
        if scale_units and not isinstance(scale_units, int):
            raise TypeError("Expected argument 'scale_units' to be a int")
        pulumi.set(__self__, "scale_units", scale_units)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accessControl")
    def access_control(self) -> Optional['outputs.StreamingEndpointAccessControlResponse']:
        """
        The access control definition of the StreamingEndpoint.
        """
        return pulumi.get(self, "access_control")

    @property
    @pulumi.getter(name="availabilitySetName")
    def availability_set_name(self) -> Optional[str]:
        """
        The name of the AvailabilitySet used with this StreamingEndpoint for high availability streaming.  This value can only be set at creation time.
        """
        return pulumi.get(self, "availability_set_name")

    @property
    @pulumi.getter(name="cdnEnabled")
    def cdn_enabled(self) -> Optional[bool]:
        """
        The CDN enabled flag.
        """
        return pulumi.get(self, "cdn_enabled")

    @property
    @pulumi.getter(name="cdnProfile")
    def cdn_profile(self) -> Optional[str]:
        """
        The CDN profile name.
        """
        return pulumi.get(self, "cdn_profile")

    @property
    @pulumi.getter(name="cdnProvider")
    def cdn_provider(self) -> Optional[str]:
        """
        The CDN provider name.
        """
        return pulumi.get(self, "cdn_provider")

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        The exact time the StreamingEndpoint was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="crossSiteAccessPolicies")
    def cross_site_access_policies(self) -> Optional['outputs.CrossSiteAccessPoliciesResponse']:
        """
        The StreamingEndpoint access policies.
        """
        return pulumi.get(self, "cross_site_access_policies")

    @property
    @pulumi.getter(name="customHostNames")
    def custom_host_names(self) -> Optional[Sequence[str]]:
        """
        The custom host names of the StreamingEndpoint
        """
        return pulumi.get(self, "custom_host_names")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The StreamingEndpoint description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="freeTrialEndTime")
    def free_trial_end_time(self) -> str:
        """
        The free trial expiration time.
        """
        return pulumi.get(self, "free_trial_end_time")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        The StreamingEndpoint host name.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> str:
        """
        The exact time the StreamingEndpoint was last modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The Azure Region of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxCacheAge")
    def max_cache_age(self) -> Optional[float]:
        """
        Max cache age
        """
        return pulumi.get(self, "max_cache_age")

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
        The provisioning state of the StreamingEndpoint.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> str:
        """
        The resource state of the StreamingEndpoint.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="scaleUnits")
    def scale_units(self) -> int:
        """
        The number of scale units.  Use the Scale operation to adjust this value.
        """
        return pulumi.get(self, "scale_units")

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
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetStreamingEndpointResult(GetStreamingEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamingEndpointResult(
            access_control=self.access_control,
            availability_set_name=self.availability_set_name,
            cdn_enabled=self.cdn_enabled,
            cdn_profile=self.cdn_profile,
            cdn_provider=self.cdn_provider,
            created=self.created,
            cross_site_access_policies=self.cross_site_access_policies,
            custom_host_names=self.custom_host_names,
            description=self.description,
            free_trial_end_time=self.free_trial_end_time,
            host_name=self.host_name,
            id=self.id,
            last_modified=self.last_modified,
            location=self.location,
            max_cache_age=self.max_cache_age,
            name=self.name,
            provisioning_state=self.provisioning_state,
            resource_state=self.resource_state,
            scale_units=self.scale_units,
            tags=self.tags,
            type=self.type)


def get_streaming_endpoint(account_name: Optional[str] = None,
                           resource_group_name: Optional[str] = None,
                           streaming_endpoint_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamingEndpointResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    :param str streaming_endpoint_name: The name of the StreamingEndpoint.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['streamingEndpointName'] = streaming_endpoint_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:media/v20190501preview:getStreamingEndpoint', __args__, opts=opts, typ=GetStreamingEndpointResult).value

    return AwaitableGetStreamingEndpointResult(
        access_control=__ret__.access_control,
        availability_set_name=__ret__.availability_set_name,
        cdn_enabled=__ret__.cdn_enabled,
        cdn_profile=__ret__.cdn_profile,
        cdn_provider=__ret__.cdn_provider,
        created=__ret__.created,
        cross_site_access_policies=__ret__.cross_site_access_policies,
        custom_host_names=__ret__.custom_host_names,
        description=__ret__.description,
        free_trial_end_time=__ret__.free_trial_end_time,
        host_name=__ret__.host_name,
        id=__ret__.id,
        last_modified=__ret__.last_modified,
        location=__ret__.location,
        max_cache_age=__ret__.max_cache_age,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        resource_state=__ret__.resource_state,
        scale_units=__ret__.scale_units,
        tags=__ret__.tags,
        type=__ret__.type)
