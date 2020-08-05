# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class SiteInstanceDeploymentSlot(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource
    """
    location: pulumi.Output[str]
    """
    Resource Location
    """
    name: pulumi.Output[str]
    """
    Resource Name
    """
    properties: pulumi.Output[dict]
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, active=None, author=None, author_email=None, deployer=None, details=None, end_time=None, instance_id=None, kind=None, location=None, message=None, name=None, resource_group_name=None, slot=None, start_time=None, status=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents user credentials used for publishing activity

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Active
        :param pulumi.Input[str] author: Author
        :param pulumi.Input[str] author_email: AuthorEmail
        :param pulumi.Input[str] deployer: Deployer
        :param pulumi.Input[str] details: Detail
        :param pulumi.Input[str] end_time: EndTime
        :param pulumi.Input[str] instance_id: Id of web app instance
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] message: Message
        :param pulumi.Input[str] name: Resource Id
        :param pulumi.Input[str] resource_group_name: Name of resource group
        :param pulumi.Input[str] slot: Name of web app slot. If not specified then will default to production slot.
        :param pulumi.Input[str] start_time: StartTime
        :param pulumi.Input[float] status: Status
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[str] type: Resource type
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

            __props__['active'] = active
            __props__['author'] = author
            __props__['author_email'] = author_email
            __props__['deployer'] = deployer
            __props__['details'] = details
            __props__['end_time'] = end_time
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['message'] = message
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if slot is None:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['start_time'] = start_time
            __props__['status'] = status
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['properties'] = None
        super(SiteInstanceDeploymentSlot, __self__).__init__(
            'azurerm:web/v20150801:SiteInstanceDeploymentSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing SiteInstanceDeploymentSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SiteInstanceDeploymentSlot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
