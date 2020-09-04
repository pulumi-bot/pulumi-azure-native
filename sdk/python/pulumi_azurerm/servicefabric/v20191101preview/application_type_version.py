# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['ApplicationTypeVersion']


class ApplicationTypeVersion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_package_url: Optional[pulumi.Input[str]] = None,
                 application_type_name: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An application type version resource for the specified application type name resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_package_url: The URL to the application package
        :param pulumi.Input[str] application_type_name: The name of the application type name resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster resource.
        :param pulumi.Input[str] location: It will be deprecated in New API, resource location depends on the parent resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Azure resource tags.
        :param pulumi.Input[str] version: The application type version.
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

            if app_package_url is None:
                raise TypeError("Missing required property 'app_package_url'")
            __props__['app_package_url'] = app_package_url
            if application_type_name is None:
                raise TypeError("Missing required property 'application_type_name'")
            __props__['application_type_name'] = application_type_name
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
            __props__['default_parameter_list'] = None
            __props__['etag'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:servicefabric/latest:ApplicationTypeVersion"), pulumi.Alias(type_="azurerm:servicefabric/v20170701preview:ApplicationTypeVersion"), pulumi.Alias(type_="azurerm:servicefabric/v20190301:ApplicationTypeVersion"), pulumi.Alias(type_="azurerm:servicefabric/v20190301preview:ApplicationTypeVersion"), pulumi.Alias(type_="azurerm:servicefabric/v20190601preview:ApplicationTypeVersion"), pulumi.Alias(type_="azurerm:servicefabric/v20200301:ApplicationTypeVersion")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApplicationTypeVersion, __self__).__init__(
            'azurerm:servicefabric/v20191101preview:ApplicationTypeVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApplicationTypeVersion':
        """
        Get an existing ApplicationTypeVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApplicationTypeVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appPackageUrl")
    def app_package_url(self) -> pulumi.Output[str]:
        """
        The URL to the application package
        """
        return pulumi.get(self, "app_package_url")

    @property
    @pulumi.getter(name="defaultParameterList")
    def default_parameter_list(self) -> pulumi.Output[Mapping[str, str]]:
        """
        List of application type parameters that can be overridden when creating or updating the application.
        """
        return pulumi.get(self, "default_parameter_list")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Azure resource etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        It will be deprecated in New API, resource location depends on the parent resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current deployment or provisioning state, which only appears in the response
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Azure resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

