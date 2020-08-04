# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppDeployment(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    Deployment resource specific properties
      * `active` (`bool`) - True if deployment is currently active, false if completed and null if not started.
      * `author` (`str`) - Who authored the deployment.
      * `author_email` (`str`) - Author email.
      * `deployer` (`str`) - Who performed the deployment.
      * `details` (`str`) - Details on deployment.
      * `end_time` (`str`) - End time.
      * `message` (`str`) - Details about deployment status.
      * `start_time` (`str`) - Start time.
      * `status` (`float`) - Deployment status.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        User credentials used for publishing activity.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: ID of an existing deployment.
        :param pulumi.Input[dict] properties: Deployment resource specific properties
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.

        The **properties** object supports the following:

          * `active` (`pulumi.Input[bool]`) - True if deployment is currently active, false if completed and null if not started.
          * `author` (`pulumi.Input[str]`) - Who authored the deployment.
          * `author_email` (`pulumi.Input[str]`) - Author email.
          * `deployer` (`pulumi.Input[str]`) - Who performed the deployment.
          * `details` (`pulumi.Input[str]`) - Details on deployment.
          * `end_time` (`pulumi.Input[str]`) - End time.
          * `message` (`pulumi.Input[str]`) - Details about deployment status.
          * `start_time` (`pulumi.Input[str]`) - Start time.
          * `status` (`pulumi.Input[float]`) - Deployment status.
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

            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(WebAppDeployment, __self__).__init__(
            'azurerm:web/v20180201:WebAppDeployment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppDeployment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppDeployment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
