# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['Logger']


class Logger(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_buffered: Optional[pulumi.Input[bool]] = None,
                 logger_id: Optional[pulumi.Input[str]] = None,
                 logger_type: Optional[pulumi.Input[Union[str, 'LoggerType']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Logger details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] credentials: The name and SendRule connection string of the event hub for azureEventHub logger.
               Instrumentation key for applicationInsights logger.
        :param pulumi.Input[str] description: Logger description.
        :param pulumi.Input[bool] is_buffered: Whether records are buffered in the logger before publishing. Default is assumed to be true.
        :param pulumi.Input[str] logger_id: Logger identifier. Must be unique in the API Management service instance.
        :param pulumi.Input[Union[str, 'LoggerType']] logger_type: Logger type.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_id: Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
        :param pulumi.Input[str] service_name: The name of the API Management service.
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

            if credentials is None and not opts.urn:
                raise TypeError("Missing required property 'credentials'")
            __props__['credentials'] = credentials
            __props__['description'] = description
            __props__['is_buffered'] = is_buffered
            __props__['logger_id'] = logger_id
            if logger_type is None and not opts.urn:
                raise TypeError("Missing required property 'logger_type'")
            __props__['logger_type'] = logger_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_id'] = resource_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:apimanagement/v20180601preview:Logger"), pulumi.Alias(type_="azure-native:apimanagement:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement:Logger"), pulumi.Alias(type_="azure-native:apimanagement/latest:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/latest:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20160707:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20160707:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20161010:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20161010:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20170301:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20170301:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20180101:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180101:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20190101:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20190101:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20191201:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20191201preview:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201preview:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20200601preview:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20200601preview:Logger"), pulumi.Alias(type_="azure-native:apimanagement/v20201201:Logger"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20201201:Logger")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Logger, __self__).__init__(
            'azure-native:apimanagement/v20180601preview:Logger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Logger':
        """
        Get an existing Logger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["credentials"] = None
        __props__["description"] = None
        __props__["is_buffered"] = None
        __props__["logger_type"] = None
        __props__["name"] = None
        __props__["resource_id"] = None
        __props__["type"] = None
        return Logger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The name and SendRule connection string of the event hub for azureEventHub logger.
        Instrumentation key for applicationInsights logger.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Logger description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isBuffered")
    def is_buffered(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether records are buffered in the logger before publishing. Default is assumed to be true.
        """
        return pulumi.get(self, "is_buffered")

    @property
    @pulumi.getter(name="loggerType")
    def logger_type(self) -> pulumi.Output[str]:
        """
        Logger type.
        """
        return pulumi.get(self, "logger_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

