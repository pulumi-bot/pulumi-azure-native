# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['GalleryImage']


class GalleryImage(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disallowed: Optional[pulumi.Input[pulumi.InputType['DisallowedArgs']]] = None,
                 end_of_life_date: Optional[pulumi.Input[str]] = None,
                 eula: Optional[pulumi.Input[str]] = None,
                 gallery_image_name: Optional[pulumi.Input[str]] = None,
                 gallery_name: Optional[pulumi.Input[str]] = None,
                 hyper_v_generation: Optional[pulumi.Input[Union[str, 'HyperVGeneration']]] = None,
                 identifier: Optional[pulumi.Input[pulumi.InputType['GalleryImageIdentifierArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 os_state: Optional[pulumi.Input['OperatingSystemStateTypes']] = None,
                 os_type: Optional[pulumi.Input['OperatingSystemTypes']] = None,
                 privacy_statement_uri: Optional[pulumi.Input[str]] = None,
                 purchase_plan: Optional[pulumi.Input[pulumi.InputType['ImagePurchasePlanArgs']]] = None,
                 recommended: Optional[pulumi.Input[pulumi.InputType['RecommendedMachineConfigurationArgs']]] = None,
                 release_note_uri: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Specifies information about the gallery Image Definition that you want to create or update.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of this gallery Image Definition resource. This property is updatable.
        :param pulumi.Input[pulumi.InputType['DisallowedArgs']] disallowed: Describes the disallowed disk types.
        :param pulumi.Input[str] end_of_life_date: The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
        :param pulumi.Input[str] eula: The Eula agreement for the gallery Image Definition.
        :param pulumi.Input[str] gallery_image_name: The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
        :param pulumi.Input[str] gallery_name: The name of the Shared Image Gallery in which the Image Definition is to be created.
        :param pulumi.Input[Union[str, 'HyperVGeneration']] hyper_v_generation: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        :param pulumi.Input[pulumi.InputType['GalleryImageIdentifierArgs']] identifier: This is the gallery Image Definition identifier.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input['OperatingSystemStateTypes'] os_state: This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
        :param pulumi.Input['OperatingSystemTypes'] os_type: This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
        :param pulumi.Input[str] privacy_statement_uri: The privacy statement uri.
        :param pulumi.Input[pulumi.InputType['ImagePurchasePlanArgs']] purchase_plan: Describes the gallery Image Definition purchase plan. This is used by marketplace images.
        :param pulumi.Input[pulumi.InputType['RecommendedMachineConfigurationArgs']] recommended: The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
        :param pulumi.Input[str] release_note_uri: The release note uri.
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

            __props__['description'] = description
            __props__['disallowed'] = disallowed
            __props__['end_of_life_date'] = end_of_life_date
            __props__['eula'] = eula
            if gallery_image_name is None and not opts.urn:
                raise TypeError("Missing required property 'gallery_image_name'")
            __props__['gallery_image_name'] = gallery_image_name
            if gallery_name is None and not opts.urn:
                raise TypeError("Missing required property 'gallery_name'")
            __props__['gallery_name'] = gallery_name
            __props__['hyper_v_generation'] = hyper_v_generation
            if identifier is None and not opts.urn:
                raise TypeError("Missing required property 'identifier'")
            __props__['identifier'] = identifier
            __props__['location'] = location
            if os_state is None and not opts.urn:
                raise TypeError("Missing required property 'os_state'")
            __props__['os_state'] = os_state
            if os_type is None and not opts.urn:
                raise TypeError("Missing required property 'os_type'")
            __props__['os_type'] = os_type
            __props__['privacy_statement_uri'] = privacy_statement_uri
            __props__['purchase_plan'] = purchase_plan
            __props__['recommended'] = recommended
            __props__['release_note_uri'] = release_note_uri
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:compute/latest:GalleryImage"), pulumi.Alias(type_="azure-nextgen:compute/v20180601:GalleryImage"), pulumi.Alias(type_="azure-nextgen:compute/v20190301:GalleryImage"), pulumi.Alias(type_="azure-nextgen:compute/v20191201:GalleryImage"), pulumi.Alias(type_="azure-nextgen:compute/v20200930:GalleryImage")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(GalleryImage, __self__).__init__(
            'azure-nextgen:compute/v20190701:GalleryImage',
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

        return GalleryImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this gallery Image Definition resource. This property is updatable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disallowed(self) -> pulumi.Output[Optional['outputs.DisallowedResponse']]:
        """
        Describes the disallowed disk types.
        """
        return pulumi.get(self, "disallowed")

    @property
    @pulumi.getter(name="endOfLifeDate")
    def end_of_life_date(self) -> pulumi.Output[Optional[str]]:
        """
        The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
        """
        return pulumi.get(self, "end_of_life_date")

    @property
    @pulumi.getter
    def eula(self) -> pulumi.Output[Optional[str]]:
        """
        The Eula agreement for the gallery Image Definition.
        """
        return pulumi.get(self, "eula")

    @property
    @pulumi.getter(name="hyperVGeneration")
    def hyper_v_generation(self) -> pulumi.Output[Optional[str]]:
        """
        The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        """
        return pulumi.get(self, "hyper_v_generation")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output['outputs.GalleryImageIdentifierResponse']:
        """
        This is the gallery Image Definition identifier.
        """
        return pulumi.get(self, "identifier")

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
    @pulumi.getter(name="osState")
    def os_state(self) -> pulumi.Output[str]:
        """
        This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
        """
        return pulumi.get(self, "os_state")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[str]:
        """
        This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="privacyStatementUri")
    def privacy_statement_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The privacy statement uri.
        """
        return pulumi.get(self, "privacy_statement_uri")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="purchasePlan")
    def purchase_plan(self) -> pulumi.Output[Optional['outputs.ImagePurchasePlanResponse']]:
        """
        Describes the gallery Image Definition purchase plan. This is used by marketplace images.
        """
        return pulumi.get(self, "purchase_plan")

    @property
    @pulumi.getter
    def recommended(self) -> pulumi.Output[Optional['outputs.RecommendedMachineConfigurationResponse']]:
        """
        The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
        """
        return pulumi.get(self, "recommended")

    @property
    @pulumi.getter(name="releaseNoteUri")
    def release_note_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The release note uri.
        """
        return pulumi.get(self, "release_note_uri")

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

