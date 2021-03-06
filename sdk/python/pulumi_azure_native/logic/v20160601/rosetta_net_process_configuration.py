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

__all__ = ['RosettaNetProcessConfiguration']


class RosettaNetProcessConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activity_settings: Optional[pulumi.Input[pulumi.InputType['RosettaNetPipActivitySettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initiator_role_settings: Optional[pulumi.Input[pulumi.InputType['RosettaNetPipRoleSettingsArgs']]] = None,
                 integration_account_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 process_code: Optional[pulumi.Input[str]] = None,
                 process_name: Optional[pulumi.Input[str]] = None,
                 process_version: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 responder_role_settings: Optional[pulumi.Input[pulumi.InputType['RosettaNetPipRoleSettingsArgs']]] = None,
                 rosetta_net_process_configuration_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The integration account RosettaNet process configuration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RosettaNetPipActivitySettingsArgs']] activity_settings: The RosettaNet process configuration activity settings.
        :param pulumi.Input[str] description: The integration account RosettaNet ProcessConfiguration properties.
        :param pulumi.Input[pulumi.InputType['RosettaNetPipRoleSettingsArgs']] initiator_role_settings: The RosettaNet initiator role settings.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: The metadata.
        :param pulumi.Input[str] process_code: The integration account RosettaNet process code.
        :param pulumi.Input[str] process_name: The integration account RosettaNet process name.
        :param pulumi.Input[str] process_version: The integration account RosettaNet process version.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[pulumi.InputType['RosettaNetPipRoleSettingsArgs']] responder_role_settings: The RosettaNet responder role settings.
        :param pulumi.Input[str] rosetta_net_process_configuration_name: The integration account RosettaNet ProcessConfiguration name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
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

            if activity_settings is None and not opts.urn:
                raise TypeError("Missing required property 'activity_settings'")
            __props__['activity_settings'] = activity_settings
            __props__['description'] = description
            if initiator_role_settings is None and not opts.urn:
                raise TypeError("Missing required property 'initiator_role_settings'")
            __props__['initiator_role_settings'] = initiator_role_settings
            if integration_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
            __props__['location'] = location
            __props__['metadata'] = metadata
            if process_code is None and not opts.urn:
                raise TypeError("Missing required property 'process_code'")
            __props__['process_code'] = process_code
            if process_name is None and not opts.urn:
                raise TypeError("Missing required property 'process_name'")
            __props__['process_name'] = process_name
            if process_version is None and not opts.urn:
                raise TypeError("Missing required property 'process_version'")
            __props__['process_version'] = process_version
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if responder_role_settings is None and not opts.urn:
                raise TypeError("Missing required property 'responder_role_settings'")
            __props__['responder_role_settings'] = responder_role_settings
            __props__['rosetta_net_process_configuration_name'] = rosetta_net_process_configuration_name
            __props__['tags'] = tags
            __props__['changed_time'] = None
            __props__['created_time'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:logic/v20160601:RosettaNetProcessConfiguration"), pulumi.Alias(type_="azure-native:logic:RosettaNetProcessConfiguration"), pulumi.Alias(type_="azure-nextgen:logic:RosettaNetProcessConfiguration"), pulumi.Alias(type_="azure-native:logic/latest:RosettaNetProcessConfiguration"), pulumi.Alias(type_="azure-nextgen:logic/latest:RosettaNetProcessConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RosettaNetProcessConfiguration, __self__).__init__(
            'azure-native:logic/v20160601:RosettaNetProcessConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RosettaNetProcessConfiguration':
        """
        Get an existing RosettaNetProcessConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activity_settings"] = None
        __props__["changed_time"] = None
        __props__["created_time"] = None
        __props__["description"] = None
        __props__["initiator_role_settings"] = None
        __props__["location"] = None
        __props__["metadata"] = None
        __props__["name"] = None
        __props__["process_code"] = None
        __props__["process_name"] = None
        __props__["process_version"] = None
        __props__["responder_role_settings"] = None
        __props__["tags"] = None
        __props__["type"] = None
        return RosettaNetProcessConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activitySettings")
    def activity_settings(self) -> pulumi.Output['outputs.RosettaNetPipActivitySettingsResponse']:
        """
        The RosettaNet process configuration activity settings.
        """
        return pulumi.get(self, "activity_settings")

    @property
    @pulumi.getter(name="changedTime")
    def changed_time(self) -> pulumi.Output[str]:
        """
        The changed time.
        """
        return pulumi.get(self, "changed_time")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        The created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The integration account RosettaNet ProcessConfiguration properties.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="initiatorRoleSettings")
    def initiator_role_settings(self) -> pulumi.Output['outputs.RosettaNetPipRoleSettingsResponse']:
        """
        The RosettaNet initiator role settings.
        """
        return pulumi.get(self, "initiator_role_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Gets the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="processCode")
    def process_code(self) -> pulumi.Output[str]:
        """
        The integration account RosettaNet process code.
        """
        return pulumi.get(self, "process_code")

    @property
    @pulumi.getter(name="processName")
    def process_name(self) -> pulumi.Output[str]:
        """
        The integration account RosettaNet process name.
        """
        return pulumi.get(self, "process_name")

    @property
    @pulumi.getter(name="processVersion")
    def process_version(self) -> pulumi.Output[str]:
        """
        The integration account RosettaNet process version.
        """
        return pulumi.get(self, "process_version")

    @property
    @pulumi.getter(name="responderRoleSettings")
    def responder_role_settings(self) -> pulumi.Output['outputs.RosettaNetPipRoleSettingsResponse']:
        """
        The RosettaNet responder role settings.
        """
        return pulumi.get(self, "responder_role_settings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Gets the resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

