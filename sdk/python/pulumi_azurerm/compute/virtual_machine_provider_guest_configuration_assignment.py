# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VirtualMachineProviderGuestConfigurationAssignment(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Region where the VM is located.
    """
    name: pulumi.Output[str]
    """
    Name of the guest configuration assignment.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the Guest configuration assignment.
      * `assignment_hash` (`str`) - Combined hash of the configuration package and parameters.
      * `compliance_status` (`str`) - A value indicating compliance status of the machine for the assigned guest configuration.
      * `context` (`str`) - The source which initiated the guest configuration assignment. Ex: Azure Policy
      * `guest_configuration` (`dict`) - The guest configuration to assign.
        * `configuration_parameter` (`list`) - The configuration parameters for the guest configuration.
          * `name` (`str`) - Name of the configuration parameter.
          * `value` (`str`) - Value of the configuration parameter.

        * `configuration_setting` (`dict`) - The configuration setting for the guest configuration.
          * `action_after_reboot` (`str`) - Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration
          * `allow_module_overwrite` (`str`) - If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
          * `configuration_mode` (`str`) - Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
          * `configuration_mode_frequency_mins` (`float`) - How often, in minutes, the current configuration is checked and applied. This property is ignored if the ConfigurationMode property is set to ApplyOnly. The default value is 15.
          * `reboot_if_needed` (`str`) - Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
          * `refresh_frequency_mins` (`float`) - The time interval, in minutes, at which the LCM checks a pull service to get updated configurations. This value is ignored if the LCM is not configured in pull mode. The default value is 30.

        * `content_hash` (`str`) - Combined hash of the guest configuration package and configuration parameters.
        * `content_uri` (`str`) - Uri of the storage where guest configuration package is uploaded.
        * `kind` (`str`) - Kind of the guest configuration. For example:DSC
        * `name` (`str`) - Name of the guest configuration.
        * `version` (`str`) - Version of the guest configuration.

      * `last_compliance_status_checked` (`str`) - Date and time when last compliance status was checked.
      * `latest_assignment_report` (`dict`) - Last reported guest configuration assignment report.
        * `assignment` (`dict`) - Configuration details of the guest configuration assignment.
          * `configuration` (`dict`) - Information about the configuration.
            * `name` (`str`) - Name of the configuration.
            * `version` (`str`) - Version of the configuration.

          * `name` (`str`) - Name of the guest configuration assignment.

        * `compliance_status` (`str`) - A value indicating compliance status of the machine for the assigned guest configuration.
        * `end_time` (`str`) - End date and time of the guest configuration assignment compliance status check.
        * `id` (`str`) - ARM resource id of the report for the guest configuration assignment.
        * `operation_type` (`str`) - Type of report, Consistency or Initial
        * `report_id` (`str`) - GUID that identifies the guest configuration assignment report under a subscription, resource group.
        * `resources` (`list`) - The list of resources for which guest configuration assignment compliance is checked.
          * `compliance_status` (`str`) - A value indicating compliance status of the machine for the assigned guest configuration.
          * `properties` (`dict`) - Properties of a guest configuration assignment resource.
          * `reasons` (`list`) - Compliance reason and reason code for a resource.
            * `code` (`str`) - Code for the compliance of the guest configuration assignment resource.
            * `phrase` (`str`) - Reason for the compliance of the guest configuration assignment resource.

        * `start_time` (`str`) - Start date and time of the guest configuration assignment compliance status check.
        * `vm` (`dict`) - Information about the VM.
          * `id` (`str`) - Azure resource Id of the VM.
          * `uuid` (`str`) - UUID(Universally Unique Identifier) of the VM.

      * `latest_report_id` (`str`) - Id of the latest report for the guest configuration assignment. 
      * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
      * `target_resource_id` (`str`) - VM resource Id.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, guest_configuration_assignment_name=None, location=None, name=None, properties=None, resource_group_name=None, vm_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Guest configuration assignment is an association between a machine and guest configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] guest_configuration_assignment_name: Name of the guest configuration assignment.
        :param pulumi.Input[str] location: Region where the VM is located.
        :param pulumi.Input[str] name: Name of the guest configuration assignment.
        :param pulumi.Input[dict] properties: Properties of the Guest configuration assignment.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[str] vm_name: The name of the virtual machine.

        The **properties** object supports the following:

          * `context` (`pulumi.Input[str]`) - The source which initiated the guest configuration assignment. Ex: Azure Policy
          * `guest_configuration` (`pulumi.Input[dict]`) - The guest configuration to assign.
            * `configuration_parameter` (`pulumi.Input[list]`) - The configuration parameters for the guest configuration.
              * `name` (`pulumi.Input[str]`) - Name of the configuration parameter.
              * `value` (`pulumi.Input[str]`) - Value of the configuration parameter.

            * `configuration_setting` (`pulumi.Input[dict]`) - The configuration setting for the guest configuration.
              * `action_after_reboot` (`pulumi.Input[str]`) - Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration
              * `allow_module_overwrite` (`pulumi.Input[str]`) - If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
              * `configuration_mode` (`pulumi.Input[str]`) - Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
              * `configuration_mode_frequency_mins` (`pulumi.Input[float]`) - How often, in minutes, the current configuration is checked and applied. This property is ignored if the ConfigurationMode property is set to ApplyOnly. The default value is 15.
              * `reboot_if_needed` (`pulumi.Input[str]`) - Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
              * `refresh_frequency_mins` (`pulumi.Input[float]`) - The time interval, in minutes, at which the LCM checks a pull service to get updated configurations. This value is ignored if the LCM is not configured in pull mode. The default value is 30.

            * `kind` (`pulumi.Input[str]`) - Kind of the guest configuration. For example:DSC
            * `name` (`pulumi.Input[str]`) - Name of the guest configuration.
            * `version` (`pulumi.Input[str]`) - Version of the guest configuration.

          * `latest_assignment_report` (`pulumi.Input[dict]`) - Last reported guest configuration assignment report.
            * `assignment` (`pulumi.Input[dict]`) - Configuration details of the guest configuration assignment.
              * `configuration` (`pulumi.Input[dict]`) - Information about the configuration.

            * `resources` (`pulumi.Input[list]`) - The list of resources for which guest configuration assignment compliance is checked.
              * `reasons` (`pulumi.Input[list]`) - Compliance reason and reason code for a resource.

            * `vm` (`pulumi.Input[dict]`) - Information about the VM.
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

            if guest_configuration_assignment_name is None:
                raise TypeError("Missing required property 'guest_configuration_assignment_name'")
            __props__['guest_configuration_assignment_name'] = guest_configuration_assignment_name
            __props__['location'] = location
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if vm_name is None:
                raise TypeError("Missing required property 'vm_name'")
            __props__['vm_name'] = vm_name
            __props__['type'] = None
        super(VirtualMachineProviderGuestConfigurationAssignment, __self__).__init__(
            'azurerm:compute:VirtualMachineProviderGuestConfigurationAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualMachineProviderGuestConfigurationAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachineProviderGuestConfigurationAssignment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
