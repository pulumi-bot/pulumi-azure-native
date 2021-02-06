# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LogProfile']


class LogProfile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 categories: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_profile_name: Optional[pulumi.Input[str]] = None,
                 retention_policy: Optional[pulumi.Input[pulumi.InputType['RetentionPolicyArgs']]] = None,
                 service_bus_rule_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The log profile resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] categories: the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
        :param pulumi.Input[str] log_profile_name: The name of the log profile.
        :param pulumi.Input[pulumi.InputType['RetentionPolicyArgs']] retention_policy: the retention policy for the events in the log.
        :param pulumi.Input[str] service_bus_rule_id: The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
        :param pulumi.Input[str] storage_account_id: the resource id of the storage account to which you would like to send the Activity Log.
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

            if categories is None and not opts.urn:
                raise TypeError("Missing required property 'categories'")
            __props__['categories'] = categories
            __props__['location'] = location
            if locations is None and not opts.urn:
                raise TypeError("Missing required property 'locations'")
            __props__['locations'] = locations
            if log_profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'log_profile_name'")
            __props__['log_profile_name'] = log_profile_name
            if retention_policy is None and not opts.urn:
                raise TypeError("Missing required property 'retention_policy'")
            __props__['retention_policy'] = retention_policy
            __props__['service_bus_rule_id'] = service_bus_rule_id
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:insights/latest:LogProfile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LogProfile, __self__).__init__(
            'azure-nextgen:insights/v20160301:LogProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LogProfile':
        """
        Get an existing LogProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return LogProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def categories(self) -> pulumi.Output[Sequence[str]]:
        """
        the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPolicy")
    def retention_policy(self) -> pulumi.Output['outputs.RetentionPolicyResponse']:
        """
        the retention policy for the events in the log.
        """
        return pulumi.get(self, "retention_policy")

    @property
    @pulumi.getter(name="serviceBusRuleId")
    def service_bus_rule_id(self) -> pulumi.Output[Optional[str]]:
        """
        The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
        """
        return pulumi.get(self, "service_bus_rule_id")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        the resource id of the storage account to which you would like to send the Activity Log.
        """
        return pulumi.get(self, "storage_account_id")

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
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

