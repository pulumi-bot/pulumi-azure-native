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

__all__ = ['Diagnostic']


class Diagnostic(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 always_log: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[pulumi.InputType['PipelineDiagnosticSettingsArgs']]] = None,
                 frontend: Optional[pulumi.Input[pulumi.InputType['PipelineDiagnosticSettingsArgs']]] = None,
                 http_correlation_protocol: Optional[pulumi.Input[str]] = None,
                 log_client_ip: Optional[pulumi.Input[bool]] = None,
                 logger_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sampling: Optional[pulumi.Input[pulumi.InputType['SamplingSettingsArgs']]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 verbosity: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Diagnostic details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] always_log: Specifies for what type of messages sampling settings should not apply.
        :param pulumi.Input[pulumi.InputType['PipelineDiagnosticSettingsArgs']] backend: Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        :param pulumi.Input[pulumi.InputType['PipelineDiagnosticSettingsArgs']] frontend: Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        :param pulumi.Input[str] http_correlation_protocol: Sets correlation protocol to use for Application Insights diagnostics.
        :param pulumi.Input[bool] log_client_ip: Log the ClientIP. Default is false.
        :param pulumi.Input[str] logger_id: Resource Id of a target logger.
        :param pulumi.Input[str] name: Diagnostic identifier. Must be unique in the current API Management service instance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['SamplingSettingsArgs']] sampling: Sampling settings for Diagnostic.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[str] verbosity: The verbosity level applied to traces emitted by trace policies.
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

            __props__['always_log'] = always_log
            __props__['backend'] = backend
            __props__['frontend'] = frontend
            __props__['http_correlation_protocol'] = http_correlation_protocol
            __props__['log_client_ip'] = log_client_ip
            if logger_id is None:
                raise TypeError("Missing required property 'logger_id'")
            __props__['logger_id'] = logger_id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sampling'] = sampling
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['verbosity'] = verbosity
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/v20170301:Diagnostic"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:Diagnostic"), pulumi.Alias(type_="azurerm:apimanagement/v20190101:Diagnostic")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Diagnostic, __self__).__init__(
            'azurerm:apimanagement/v20191201:Diagnostic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Diagnostic':
        """
        Get an existing Diagnostic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Diagnostic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alwaysLog")
    def always_log(self) -> Optional[str]:
        """
        Specifies for what type of messages sampling settings should not apply.
        """
        return pulumi.get(self, "always_log")

    @property
    @pulumi.getter
    def backend(self) -> Optional['outputs.PipelineDiagnosticSettingsResponse']:
        """
        Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def frontend(self) -> Optional['outputs.PipelineDiagnosticSettingsResponse']:
        """
        Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        """
        return pulumi.get(self, "frontend")

    @property
    @pulumi.getter(name="httpCorrelationProtocol")
    def http_correlation_protocol(self) -> Optional[str]:
        """
        Sets correlation protocol to use for Application Insights diagnostics.
        """
        return pulumi.get(self, "http_correlation_protocol")

    @property
    @pulumi.getter(name="logClientIp")
    def log_client_ip(self) -> Optional[bool]:
        """
        Log the ClientIP. Default is false.
        """
        return pulumi.get(self, "log_client_ip")

    @property
    @pulumi.getter(name="loggerId")
    def logger_id(self) -> str:
        """
        Resource Id of a target logger.
        """
        return pulumi.get(self, "logger_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sampling(self) -> Optional['outputs.SamplingSettingsResponse']:
        """
        Sampling settings for Diagnostic.
        """
        return pulumi.get(self, "sampling")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def verbosity(self) -> Optional[str]:
        """
        The verbosity level applied to traces emitted by trace policies.
        """
        return pulumi.get(self, "verbosity")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

