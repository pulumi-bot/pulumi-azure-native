# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['GalleryApplication']


class GalleryApplication(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_of_life_date: Optional[pulumi.Input[str]] = None,
                 eula: Optional[pulumi.Input[str]] = None,
                 gallery_application_name: Optional[pulumi.Input[str]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 privacy_statement_uri: Optional[pulumi.Input[str]] = None,
                 release_note_uri: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 supported_os_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Specifies information about the gallery Application Definition that you want to create or update.

        ## Example Usage
        ### Create or update a simple gallery Application.

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        gallery_application = azurerm.compute.v20190701.GalleryApplication("galleryApplication",
            description="This is the gallery application description.",
            eula="This is the gallery application EULA.",
            gallery_application_name="myGalleryApplicationName",
            gallery_name="myGalleryName",
            location="West US",
            privacy_statement_uri="myPrivacyStatementUri}",
            release_note_uri="myReleaseNoteUri",
            resource_group_name="myResourceGroup",
            supported_os_type="Windows")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this gallery Application Definition resource. This property is updatable.
        :param pulumi.Input[str] end_of_life_date: The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
        :param pulumi.Input[str] eula: The Eula agreement for the gallery Application Definition.
        :param pulumi.Input[str] gallery_application_name: The name of the gallery Application Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
        :param pulumi.Input[str] gallery_name: The name of the Shared Application Gallery in which the Application Definition is to be created.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] privacy_statement_uri: The privacy statement uri.
        :param pulumi.Input[str] release_note_uri: The release note uri.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] supported_os_type: This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
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

            __props__['description'] = description
            __props__['end_of_life_date'] = end_of_life_date
            __props__['eula'] = eula
            if gallery_application_name is None:
                raise TypeError("Missing required property 'gallery_application_name'")
            __props__['gallery_application_name'] = gallery_application_name
            if gallery_name is None:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['privacy_statement_uri'] = privacy_statement_uri
            __props__['release_note_uri'] = release_note_uri
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if supported_os_type is None:
                raise TypeError("Missing required property 'supported_os_type'")
            __props__['supported_os_type'] = supported_os_type
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:compute/latest:GalleryApplication"), pulumi.Alias(type_="azurerm:compute/v20190301:GalleryApplication"), pulumi.Alias(type_="azurerm:compute/v20191201:GalleryApplication")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GalleryApplication, __self__).__init__(
            'azurerm:compute/v20190701:GalleryApplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GalleryApplication':
        """
        Get an existing GalleryApplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GalleryApplication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this gallery Application Definition resource. This property is updatable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endOfLifeDate")
    def end_of_life_date(self) -> pulumi.Output[Optional[str]]:
        """
        The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
        """
        return pulumi.get(self, "end_of_life_date")

    @property
    @pulumi.getter
    def eula(self) -> pulumi.Output[Optional[str]]:
        """
        The Eula agreement for the gallery Application Definition.
        """
        return pulumi.get(self, "eula")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privacyStatementUri")
    def privacy_statement_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The privacy statement uri.
        """
        return pulumi.get(self, "privacy_statement_uri")

    @property
    @pulumi.getter(name="releaseNoteUri")
    def release_note_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The release note uri.
        """
        return pulumi.get(self, "release_note_uri")

    @property
    @pulumi.getter(name="supportedOSType")
    def supported_os_type(self) -> pulumi.Output[str]:
        """
        This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
        """
        return pulumi.get(self, "supported_os_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

