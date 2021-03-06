# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = ['GalleryImage']


class GalleryImage(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gallery_image_name: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 is_override: Optional[pulumi.Input[bool]] = None,
                 is_plan_authorized: Optional[pulumi.Input[bool]] = None,
                 lab_account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 unique_identifier: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an image from the Azure Marketplace
        API Version: 2018-10-15.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gallery_image_name: The name of the gallery Image.
        :param pulumi.Input[bool] is_enabled: Indicates whether this gallery image is enabled.
        :param pulumi.Input[bool] is_override: Indicates whether this gallery has been overridden for this lab account
        :param pulumi.Input[bool] is_plan_authorized: Indicates if the plan has been authorized for programmatic deployment.
        :param pulumi.Input[str] lab_account_name: The name of the lab Account.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] provisioning_state: The provisioning status of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] unique_identifier: The unique immutable identifier of a resource (Guid).
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

            __props__['gallery_image_name'] = gallery_image_name
            __props__['is_enabled'] = is_enabled
            __props__['is_override'] = is_override
            __props__['is_plan_authorized'] = is_plan_authorized
            if lab_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'lab_account_name'")
            __props__['lab_account_name'] = lab_account_name
            __props__['location'] = location
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['unique_identifier'] = unique_identifier
            __props__['author'] = None
            __props__['created_date'] = None
            __props__['description'] = None
            __props__['icon'] = None
            __props__['image_reference'] = None
            __props__['latest_operation_result'] = None
            __props__['name'] = None
            __props__['plan_id'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:labservices:GalleryImage"), pulumi.Alias(type_="azure-native:labservices/latest:GalleryImage"), pulumi.Alias(type_="azure-nextgen:labservices/latest:GalleryImage"), pulumi.Alias(type_="azure-native:labservices/v20181015:GalleryImage"), pulumi.Alias(type_="azure-nextgen:labservices/v20181015:GalleryImage")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GalleryImage, __self__).__init__(
            'azure-native:labservices:GalleryImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GalleryImage':
        """
        Get an existing GalleryImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["author"] = None
        __props__["created_date"] = None
        __props__["description"] = None
        __props__["icon"] = None
        __props__["image_reference"] = None
        __props__["is_enabled"] = None
        __props__["is_override"] = None
        __props__["is_plan_authorized"] = None
        __props__["latest_operation_result"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["plan_id"] = None
        __props__["provisioning_state"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["unique_identifier"] = None
        return GalleryImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[str]:
        """
        The author of the gallery image.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[str]:
        """
        The creation date of the gallery image.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the gallery image.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def icon(self) -> pulumi.Output[str]:
        """
        The icon of the gallery image.
        """
        return pulumi.get(self, "icon")

    @property
    @pulumi.getter(name="imageReference")
    def image_reference(self) -> pulumi.Output['outputs.GalleryImageReferenceResponse']:
        """
        The image reference of the gallery image.
        """
        return pulumi.get(self, "image_reference")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this gallery image is enabled.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="isOverride")
    def is_override(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this gallery has been overridden for this lab account
        """
        return pulumi.get(self, "is_override")

    @property
    @pulumi.getter(name="isPlanAuthorized")
    def is_plan_authorized(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the plan has been authorized for programmatic deployment.
        """
        return pulumi.get(self, "is_plan_authorized")

    @property
    @pulumi.getter(name="latestOperationResult")
    def latest_operation_result(self) -> pulumi.Output['outputs.LatestOperationResultResponse']:
        """
        The details of the latest operation. ex: status, error
        """
        return pulumi.get(self, "latest_operation_result")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        """
        The third party plan that applies to this image
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

