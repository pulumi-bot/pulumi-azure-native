# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ComponentFavorite(pulumi.CustomResource):
    category: pulumi.Output[str]
    """
    Favorite category, as defined by the user at creation time.
    """
    config: pulumi.Output[str]
    """
    Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
    """
    favorite_id: pulumi.Output[str]
    """
    Internally assigned unique id of the favorite definition.
    """
    favorite_type: pulumi.Output[str]
    """
    Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
    """
    is_generated_from_template: pulumi.Output[bool]
    """
    Flag denoting wether or not this favorite was generated from a template.
    """
    name: pulumi.Output[str]
    """
    The user-defined name of the favorite.
    """
    source_type: pulumi.Output[str]
    """
    The source of the favorite definition.
    """
    tags: pulumi.Output[list]
    """
    A list of 0 or more tags that are associated with this favorite definition
    """
    time_modified: pulumi.Output[str]
    """
    Date and time in UTC of the last modification that was made to this favorite definition.
    """
    user_id: pulumi.Output[str]
    """
    Unique user id of the specific user that owns this favorite.
    """
    version: pulumi.Output[str]
    """
    This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
    """
    def __init__(__self__, resource_name, opts=None, category=None, config=None, favorite_type=None, is_generated_from_template=None, name=None, source_type=None, tags=None, version=None, favorite_id=None, resource_group_name=None, resource_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Properties that define a favorite that is associated to an Application Insights component.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: Favorite category, as defined by the user at creation time.
        :param pulumi.Input[str] config: Configuration of this particular favorite, which are driven by the Azure portal UX. Configuration data is a string containing valid JSON
        :param pulumi.Input[str] favorite_type: Enum indicating if this favorite definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        :param pulumi.Input[bool] is_generated_from_template: Flag denoting wether or not this favorite was generated from a template.
        :param pulumi.Input[str] name: The user-defined name of the favorite.
        :param pulumi.Input[str] source_type: The source of the favorite definition.
        :param pulumi.Input[list] tags: A list of 0 or more tags that are associated with this favorite definition
        :param pulumi.Input[str] version: This instance's version of the data model. This can change as new features are added that can be marked favorite. Current examples include MetricsExplorer (ME) and Search.
        :param pulumi.Input[str] favorite_id: The Id of a specific favorite defined in the Application Insights component
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_name: The name of the Application Insights component resource.
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

            __props__['category'] = category
            __props__['config'] = config
            __props__['favorite_type'] = favorite_type
            __props__['is_generated_from_template'] = is_generated_from_template
            __props__['name'] = name
            __props__['source_type'] = source_type
            __props__['tags'] = tags
            __props__['version'] = version
            if favorite_id is None:
                raise TypeError("Missing required property 'favorite_id'")
            __props__['favorite_id'] = favorite_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name is None:
                raise TypeError("Missing required property 'resource_name'")
            __props__['resource_name'] = resource_name
            __props__['favorite_id'] = None
            __props__['time_modified'] = None
            __props__['user_id'] = None
        super(ComponentFavorite, __self__).__init__(
            'azurerm:insights:ComponentFavorite',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ComponentFavorite resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ComponentFavorite(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
