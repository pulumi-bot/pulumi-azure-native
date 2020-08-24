# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GalleryImageVersion']


class GalleryImageVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gallery_image_name: Optional[pulumi.Input[str]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publishing_profile: Optional[pulumi.Input[pulumi.InputType['GalleryImageVersionPublishingProfileArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Specifies information about the gallery Image Version that you want to create or update.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gallery_image_name: The name of the gallery Image Definition in which the Image Version is to be created.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Image Definition resides.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
        :param pulumi.Input[pulumi.InputType['GalleryImageVersionPublishingProfileArgs']] publishing_profile: The publishing profile of a gallery Image Version.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if gallery_image_name is None:
                raise TypeError("Missing required property 'gallery_image_name'")
            __props__['gallery_image_name'] = gallery_image_name
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if publishing_profile is None:
                raise TypeError("Missing required property 'publishing_profile'")
            __props__['publishing_profile'] = publishing_profile
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['provisioning_state'] = None
            __props__['replication_status'] = None
            __props__['storage_profile'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:compute/v20180601:GalleryImageVersion"), pulumi.Alias(type_="azurerm:compute/v20190701:GalleryImageVersion"), pulumi.Alias(type_="azurerm:compute/v20191201:GalleryImageVersion")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GalleryImageVersion, __self__).__init__(
            'azurerm:compute/v20190301:GalleryImageVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GalleryImageVersion':
        """
        Get an existing GalleryImageVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GalleryImageVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publishingProfile")
    def publishing_profile(self) -> 'outputs.GalleryImageVersionPublishingProfileResponse':
        """
        The publishing profile of a gallery Image Version.
        """
        return pulumi.get(self, "publishing_profile")

    @property
    @pulumi.getter(name="replicationStatus")
    def replication_status(self) -> 'outputs.ReplicationStatusResponse':
        """
        This is the replication status of the gallery Image Version.
        """
        return pulumi.get(self, "replication_status")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> 'outputs.GalleryImageVersionStorageProfileResponse':
        """
        This is the storage profile of a Gallery Image Version.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

