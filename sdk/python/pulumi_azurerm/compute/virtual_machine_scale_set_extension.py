# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VirtualMachineScaleSetExtension(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the extension.
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of a Virtual Machine Scale Set Extension.
      * `auto_upgrade_minor_version` (`bool`) - Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
      * `force_update_tag` (`str`) - If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
      * `protected_settings` (`dict`) - The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
      * `provision_after_extensions` (`list`) - Collection of extension names after which this extension needs to be provisioned.
      * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
      * `publisher` (`str`) - The name of the extension handler publisher.
      * `settings` (`dict`) - Json formatted public settings for the extension.
      * `type` (`str`) - Specifies the type of the extension; an example is "CustomScriptExtension".
      * `type_handler_version` (`str`) - Specifies the version of the script handler.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, vm_scale_set_name=None, vmss_extension_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Describes a Virtual Machine Scale Set Extension.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the extension.
        :param pulumi.Input[dict] properties: Describes the properties of a Virtual Machine Scale Set Extension.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] vm_scale_set_name: The name of the VM scale set where the extension should be create or updated.
        :param pulumi.Input[str] vmss_extension_name: The name of the VM scale set extension.

        The **properties** object supports the following:

          * `auto_upgrade_minor_version` (`pulumi.Input[bool]`) - Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
          * `force_update_tag` (`pulumi.Input[str]`) - If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
          * `protected_settings` (`pulumi.Input[dict]`) - The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
          * `provision_after_extensions` (`pulumi.Input[list]`) - Collection of extension names after which this extension needs to be provisioned.
          * `publisher` (`pulumi.Input[str]`) - The name of the extension handler publisher.
          * `settings` (`pulumi.Input[dict]`) - Json formatted public settings for the extension.
          * `type` (`pulumi.Input[str]`) - Specifies the type of the extension; an example is "CustomScriptExtension".
          * `type_handler_version` (`pulumi.Input[str]`) - Specifies the version of the script handler.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if vm_scale_set_name is None:
                raise TypeError("Missing required property 'vm_scale_set_name'")
            __props__['vm_scale_set_name'] = vm_scale_set_name
            if vmss_extension_name is None:
                raise TypeError("Missing required property 'vmss_extension_name'")
            __props__['vmss_extension_name'] = vmss_extension_name
            __props__['type'] = None
        super(VirtualMachineScaleSetExtension, __self__).__init__(
            'azurerm:compute:VirtualMachineScaleSetExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualMachineScaleSetExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachineScaleSetExtension(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
