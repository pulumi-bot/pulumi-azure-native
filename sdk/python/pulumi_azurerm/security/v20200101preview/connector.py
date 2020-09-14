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

__all__ = ['Connector']


class Connector(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication_details: Optional[pulumi.Input[Union[pulumi.InputType['AwAssumeRoleAuthenticationDetailsPropertiesArgs'], pulumi.InputType['AwsCredsAuthenticationDetailsPropertiesArgs'], pulumi.InputType['GcpCredentialsDetailsPropertiesArgs']]]] = None,
                 connector_name: Optional[pulumi.Input[str]] = None,
                 hybrid_compute_settings: Optional[pulumi.Input[pulumi.InputType['HybridComputeSettingsPropertiesArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The connector setting

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[pulumi.InputType['AwAssumeRoleAuthenticationDetailsPropertiesArgs'], pulumi.InputType['AwsCredsAuthenticationDetailsPropertiesArgs'], pulumi.InputType['GcpCredentialsDetailsPropertiesArgs']]] authentication_details: Settings for authentication management, these settings are relevant only for the cloud connector.
        :param pulumi.Input[str] connector_name: Name of the cloud account connector
        :param pulumi.Input[pulumi.InputType['HybridComputeSettingsPropertiesArgs']] hybrid_compute_settings: Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
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

            __props__['authentication_details'] = authentication_details
            if connector_name is None:
                raise TypeError("Missing required property 'connector_name'")
            __props__['connector_name'] = connector_name
            __props__['hybrid_compute_settings'] = hybrid_compute_settings
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:security/latest:Connector")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Connector, __self__).__init__(
            'azurerm:security/v20200101preview:Connector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connector':
        """
        Get an existing Connector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Connector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authenticationDetails")
    def authentication_details(self) -> pulumi.Output[Optional[Any]]:
        """
        Settings for authentication management, these settings are relevant only for the cloud connector.
        """
        return pulumi.get(self, "authentication_details")

    @property
    @pulumi.getter(name="hybridComputeSettings")
    def hybrid_compute_settings(self) -> pulumi.Output[Optional['outputs.HybridComputeSettingsPropertiesResponse']]:
        """
        Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
        """
        return pulumi.get(self, "hybrid_compute_settings")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

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

