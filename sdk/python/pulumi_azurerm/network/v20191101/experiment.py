# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Experiment(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The properties of an Experiment
      * `description` (`str`) - The description of the details or intents of the Experiment
      * `enabled_state` (`str`) - The state of the Experiment
      * `endpoint_a` (`dict`) - The endpoint A of an experiment
        * `endpoint` (`str`) - The endpoint URL
        * `name` (`str`) - The name of the endpoint

      * `endpoint_b` (`dict`) - The endpoint B of an experiment
      * `resource_state` (`str`) - Resource status.
      * `script_file_uri` (`str`) - The uri to the Script used in the Experiment
      * `status` (`str`) - The description of Experiment status from the server side
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, description=None, enabled_state=None, endpoint_a=None, endpoint_b=None, location=None, name=None, profile_name=None, resource_group_name=None, resource_state=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Defines the properties of an Experiment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the details or intents of the Experiment
        :param pulumi.Input[str] enabled_state: The state of the Experiment
        :param pulumi.Input[dict] endpoint_a: The endpoint A of an experiment
        :param pulumi.Input[dict] endpoint_b: The endpoint B of an experiment
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The Experiment identifier associated with the Experiment
        :param pulumi.Input[str] profile_name: The Profile identifier associated with the Tenant and Partner
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[str] resource_state: Resource status.
        :param pulumi.Input[dict] tags: Resource tags.

        The **endpoint_a** object supports the following:

          * `endpoint` (`pulumi.Input[str]`) - The endpoint URL
          * `name` (`pulumi.Input[str]`) - The name of the endpoint
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
            __props__['enabled_state'] = enabled_state
            __props__['endpoint_a'] = endpoint_a
            __props__['endpoint_b'] = endpoint_b
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_state'] = resource_state
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Experiment, __self__).__init__(
            'azurerm:network/v20191101:Experiment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Experiment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Experiment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
