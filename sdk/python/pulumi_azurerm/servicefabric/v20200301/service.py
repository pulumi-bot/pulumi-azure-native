# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Service(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Azure resource etag.
    """
    location: pulumi.Output[str]
    """
    It will be deprecated in New API, resource location depends on the parent resource.
    """
    name: pulumi.Output[str]
    """
    Azure resource name.
    """
    properties: pulumi.Output[dict]
    """
    The service resource properties.
      * `correlation_scheme` (`list`) - A list that describes the correlation of the service with other services.
        * `scheme` (`str`) - The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
        * `service_name` (`str`) - The name of the service that the correlation relationship is established with.

      * `default_move_cost` (`str`) - Specifies the move cost for the service.
      * `partition_description` (`dict`) - Describes how the service is partitioned.
        * `partition_scheme` (`str`) - Specifies how the service is partitioned.

      * `placement_constraints` (`str`) - The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
      * `provisioning_state` (`str`) - The current deployment or provisioning state, which only appears in the response
      * `service_dns_name` (`str`) - Dns name used for the service. If this is specified, then the service can be accessed via its DNS name instead of service name.
      * `service_kind` (`str`) - The kind of service (Stateless or Stateful).
      * `service_load_metrics` (`list`) - The service load metrics is given as an array of ServiceLoadMetricDescription objects.
        * `default_load` (`float`) - Used only for Stateless services. The default amount of load, as a number, that this service creates for this metric.
        * `name` (`str`) - The name of the metric. If the service chooses to report load during runtime, the load metric name should match the name that is specified in Name exactly. Note that metric names are case sensitive.
        * `primary_default_load` (`float`) - Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Primary replica.
        * `secondary_default_load` (`float`) - Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Secondary replica.
        * `weight` (`str`) - The service load metric relative weight, compared to other metrics configured for this service, as a number.

      * `service_package_activation_mode` (`str`) - The activation Mode of the service package
      * `service_placement_policies` (`list`) - A list that describes the correlation of the service with other services.
        * `type` (`str`) - The type of placement policy for a service fabric service. Following are the possible values.

      * `service_type_name` (`str`) - The name of the service type
    """
    tags: pulumi.Output[dict]
    """
    Azure resource tags.
    """
    type: pulumi.Output[str]
    """
    Azure resource type.
    """
    def __init__(__self__, resource_name, opts=None, application_name=None, cluster_name=None, correlation_scheme=None, default_move_cost=None, location=None, name=None, partition_description=None, placement_constraints=None, resource_group_name=None, service_dns_name=None, service_kind=None, service_load_metrics=None, service_package_activation_mode=None, service_placement_policies=None, service_type_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The service resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_name: The name of the application resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster resource.
        :param pulumi.Input[list] correlation_scheme: A list that describes the correlation of the service with other services.
        :param pulumi.Input[str] default_move_cost: Specifies the move cost for the service.
        :param pulumi.Input[str] location: It will be deprecated in New API, resource location depends on the parent resource.
        :param pulumi.Input[str] name: The name of the service resource in the format of {applicationName}~{serviceName}.
        :param pulumi.Input[dict] partition_description: Describes how the service is partitioned.
        :param pulumi.Input[str] placement_constraints: The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_dns_name: Dns name used for the service. If this is specified, then the service can be accessed via its DNS name instead of service name.
        :param pulumi.Input[str] service_kind: The kind of service (Stateless or Stateful).
        :param pulumi.Input[list] service_load_metrics: The service load metrics is given as an array of ServiceLoadMetricDescription objects.
        :param pulumi.Input[str] service_package_activation_mode: The activation Mode of the service package
        :param pulumi.Input[list] service_placement_policies: A list that describes the correlation of the service with other services.
        :param pulumi.Input[str] service_type_name: The name of the service type
        :param pulumi.Input[dict] tags: Azure resource tags.

        The **correlation_scheme** object supports the following:

          * `scheme` (`pulumi.Input[str]`) - The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
          * `service_name` (`pulumi.Input[str]`) - The name of the service that the correlation relationship is established with.

        The **partition_description** object supports the following:

          * `partition_scheme` (`pulumi.Input[str]`) - Specifies how the service is partitioned.

        The **service_load_metrics** object supports the following:

          * `default_load` (`pulumi.Input[float]`) - Used only for Stateless services. The default amount of load, as a number, that this service creates for this metric.
          * `name` (`pulumi.Input[str]`) - The name of the metric. If the service chooses to report load during runtime, the load metric name should match the name that is specified in Name exactly. Note that metric names are case sensitive.
          * `primary_default_load` (`pulumi.Input[float]`) - Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Primary replica.
          * `secondary_default_load` (`pulumi.Input[float]`) - Used only for Stateful services. The default amount of load, as a number, that this service creates for this metric when it is a Secondary replica.
          * `weight` (`pulumi.Input[str]`) - The service load metric relative weight, compared to other metrics configured for this service, as a number.

        The **service_placement_policies** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of placement policy for a service fabric service. Following are the possible values.
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

            if application_name is None:
                raise TypeError("Missing required property 'application_name'")
            __props__['application_name'] = application_name
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['correlation_scheme'] = correlation_scheme
            __props__['default_move_cost'] = default_move_cost
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['partition_description'] = partition_description
            __props__['placement_constraints'] = placement_constraints
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_dns_name'] = service_dns_name
            if service_kind is None:
                raise TypeError("Missing required property 'service_kind'")
            __props__['service_kind'] = service_kind
            __props__['service_load_metrics'] = service_load_metrics
            __props__['service_package_activation_mode'] = service_package_activation_mode
            __props__['service_placement_policies'] = service_placement_policies
            __props__['service_type_name'] = service_type_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(Service, __self__).__init__(
            'azurerm:servicefabric/v20200301:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Service(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop