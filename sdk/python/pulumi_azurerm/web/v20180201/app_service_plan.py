# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AppServicePlan(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    location: pulumi.Output[str]
    """
    Resource Location.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    AppServicePlan resource specific properties
      * `free_offer_expiration_time` (`str`) - The time when the server farm free offer expires.
      * `geo_region` (`str`) - Geographical location for the App Service plan.
      * `hosting_environment_profile` (`dict`) - Specification for the App Service Environment to use for the App Service plan.
        * `id` (`str`) - Resource ID of the App Service Environment.
        * `name` (`str`) - Name of the App Service Environment.
        * `type` (`str`) - Resource type of the App Service Environment.

      * `hyper_v` (`bool`) - If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
      * `is_spot` (`bool`) - If <code>true</code>, this App Service Plan owns spot instances.
      * `is_xenon` (`bool`) - Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
      * `maximum_elastic_worker_count` (`float`) - Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
      * `maximum_number_of_workers` (`float`) - Maximum number of instances that can be assigned to this App Service plan.
      * `number_of_sites` (`float`) - Number of apps assigned to this App Service plan.
      * `per_site_scaling` (`bool`) - If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
        If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
      * `provisioning_state` (`str`) - Provisioning state of the App Service Environment.
      * `reserved` (`bool`) - If Linux app service plan <code>true</code>, <code>false</code> otherwise.
      * `resource_group` (`str`) - Resource group of the App Service plan.
      * `spot_expiration_time` (`str`) - The time when the server farm expires. Valid only if it is a spot server farm.
      * `status` (`str`) - App Service plan status.
      * `subscription` (`str`) - App Service plan subscription.
      * `target_worker_count` (`float`) - Scaling worker count.
      * `target_worker_size_id` (`float`) - Scaling worker size ID.
      * `worker_tier_name` (`str`) - Target worker tier assigned to the App Service plan.
    """
    sku: pulumi.Output[dict]
    """
    Description of a SKU for a scalable resource.
      * `capabilities` (`list`) - Capabilities of the SKU, e.g., is traffic manager enabled?
        * `name` (`str`) - Name of the SKU capability.
        * `reason` (`str`) - Reason of the SKU capability.
        * `value` (`str`) - Value of the SKU capability.

      * `capacity` (`float`) - Current number of instances assigned to the resource.
      * `family` (`str`) - Family code of the resource SKU.
      * `locations` (`list`) - Locations of the SKU.
      * `name` (`str`) - Name of the resource SKU.
      * `size` (`str`) - Size specifier of the resource SKU.
      * `sku_capacity` (`dict`) - Min, max, and default scale values of the SKU.
        * `default` (`float`) - Default number of workers for this App Service plan SKU.
        * `maximum` (`float`) - Maximum number of workers for this App Service plan SKU.
        * `minimum` (`float`) - Minimum number of workers for this App Service plan SKU.
        * `scale_type` (`str`) - Available scale configurations for an App Service plan.

      * `tier` (`str`) - Service tier of the resource SKU.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, free_offer_expiration_time=None, hosting_environment_profile=None, hyper_v=None, is_spot=None, is_xenon=None, kind=None, location=None, maximum_elastic_worker_count=None, name=None, per_site_scaling=None, reserved=None, resource_group_name=None, sku=None, spot_expiration_time=None, tags=None, target_worker_count=None, target_worker_size_id=None, worker_tier_name=None, __props__=None, __name__=None, __opts__=None):
        """
        App Service plan.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] free_offer_expiration_time: The time when the server farm free offer expires.
        :param pulumi.Input[dict] hosting_environment_profile: Specification for the App Service Environment to use for the App Service plan.
        :param pulumi.Input[bool] hyper_v: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
        :param pulumi.Input[bool] is_spot: If <code>true</code>, this App Service Plan owns spot instances.
        :param pulumi.Input[bool] is_xenon: Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[float] maximum_elastic_worker_count: Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
        :param pulumi.Input[str] name: Name of the App Service plan.
        :param pulumi.Input[bool] per_site_scaling: If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
               If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
        :param pulumi.Input[bool] reserved: If Linux app service plan <code>true</code>, <code>false</code> otherwise.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[dict] sku: Description of a SKU for a scalable resource.
        :param pulumi.Input[str] spot_expiration_time: The time when the server farm expires. Valid only if it is a spot server farm.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] target_worker_count: Scaling worker count.
        :param pulumi.Input[float] target_worker_size_id: Scaling worker size ID.
        :param pulumi.Input[str] worker_tier_name: Target worker tier assigned to the App Service plan.

        The **hosting_environment_profile** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID of the App Service Environment.

        The **sku** object supports the following:

          * `capabilities` (`pulumi.Input[list]`) - Capabilities of the SKU, e.g., is traffic manager enabled?
            * `name` (`pulumi.Input[str]`) - Name of the SKU capability.
            * `reason` (`pulumi.Input[str]`) - Reason of the SKU capability.
            * `value` (`pulumi.Input[str]`) - Value of the SKU capability.

          * `capacity` (`pulumi.Input[float]`) - Current number of instances assigned to the resource.
          * `family` (`pulumi.Input[str]`) - Family code of the resource SKU.
          * `locations` (`pulumi.Input[list]`) - Locations of the SKU.
          * `name` (`pulumi.Input[str]`) - Name of the resource SKU.
          * `size` (`pulumi.Input[str]`) - Size specifier of the resource SKU.
          * `sku_capacity` (`pulumi.Input[dict]`) - Min, max, and default scale values of the SKU.
            * `default` (`pulumi.Input[float]`) - Default number of workers for this App Service plan SKU.
            * `maximum` (`pulumi.Input[float]`) - Maximum number of workers for this App Service plan SKU.
            * `minimum` (`pulumi.Input[float]`) - Minimum number of workers for this App Service plan SKU.
            * `scale_type` (`pulumi.Input[str]`) - Available scale configurations for an App Service plan.

          * `tier` (`pulumi.Input[str]`) - Service tier of the resource SKU.
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

            __props__['free_offer_expiration_time'] = free_offer_expiration_time
            __props__['hosting_environment_profile'] = hosting_environment_profile
            __props__['hyper_v'] = hyper_v
            __props__['is_spot'] = is_spot
            __props__['is_xenon'] = is_xenon
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['maximum_elastic_worker_count'] = maximum_elastic_worker_count
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['per_site_scaling'] = per_site_scaling
            __props__['reserved'] = reserved
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['spot_expiration_time'] = spot_expiration_time
            __props__['tags'] = tags
            __props__['target_worker_count'] = target_worker_count
            __props__['target_worker_size_id'] = target_worker_size_id
            __props__['worker_tier_name'] = worker_tier_name
            __props__['properties'] = None
            __props__['type'] = None
        super(AppServicePlan, __self__).__init__(
            'azurerm:web/v20180201:AppServicePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AppServicePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AppServicePlan(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop