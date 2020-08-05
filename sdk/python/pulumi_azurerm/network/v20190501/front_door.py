# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class FrontDoor(pulumi.CustomResource):
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
    Properties of the Front Door Load Balancer
      * `backend_pools` (`list`) - Backend pools available to routing rules.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the Front Door Backend Pool
          * `backends` (`list`) - The set of backends for this pool
            * `address` (`str`) - Location of the backend (IP address or FQDN)
            * `backend_host_header` (`str`) - The value to use as the host header sent to the backend. If blank or unspecified, this defaults to the incoming host.
            * `enabled_state` (`str`) - Whether to enable use of this backend. Permitted values are 'Enabled' or 'Disabled'
            * `http_port` (`float`) - The HTTP TCP port number. Must be between 1 and 65535.
            * `https_port` (`float`) - The HTTPS TCP port number. Must be between 1 and 65535.
            * `priority` (`float`) - Priority to use for load balancing. Higher priorities will not be used for load balancing if any lower priority backend is healthy.
            * `weight` (`float`) - Weight of this endpoint for load balancing purposes.

          * `health_probe_settings` (`dict`) - L7 health probe settings for a backend pool
            * `id` (`str`) - Resource ID.

          * `load_balancing_settings` (`dict`) - Load balancing settings for a backend pool
          * `resource_state` (`str`) - Resource status.

        * `type` (`str`) - Resource type.

      * `backend_pools_settings` (`dict`) - Settings for all backendPools
        * `enforce_certificate_name_check` (`str`) - Whether to enforce certificate name check on HTTPS requests to all backend pools. No effect on non-HTTPS requests.
        * `send_recv_timeout_seconds` (`float`) - Send and receive timeout on forwarding request to the backend. When timeout is reached, the request fails and returns.

      * `cname` (`str`) - The host that each frontendEndpoint must CNAME to.
      * `enabled_state` (`str`) - Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
      * `friendly_name` (`str`) - A friendly name for the frontDoor
      * `frontend_endpoints` (`list`) - Frontend endpoints available to routing rules.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the Frontend endpoint
          * `custom_https_configuration` (`dict`) - The configuration specifying how to enable HTTPS
            * `certificate_source` (`str`) - Defines the source of the SSL certificate
            * `front_door_certificate_source_parameters` (`dict`) - Parameters required for enabling SSL with Front Door-managed certificates (if certificateSource=FrontDoor)
              * `certificate_type` (`str`) - Defines the type of the certificate used for secure connections to a frontendEndpoint

            * `key_vault_certificate_source_parameters` (`dict`) - KeyVault certificate source parameters (if certificateSource=AzureKeyVault)
              * `secret_name` (`str`) - The name of the Key Vault secret representing the full certificate PFX
              * `secret_version` (`str`) - The version of the Key Vault secret representing the full certificate PFX
              * `vault` (`dict`) - The Key Vault containing the SSL certificate
                * `id` (`str`) - Resource ID.

            * `minimum_tls_version` (`str`) - The minimum TLS version required from the clients to establish an SSL handshake with Front Door.
            * `protocol_type` (`str`) - Defines the TLS extension protocol that is used for secure delivery

          * `custom_https_provisioning_state` (`str`) - Provisioning status of Custom Https of the frontendEndpoint.
          * `custom_https_provisioning_substate` (`str`) - Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
          * `host_name` (`str`) - The host name of the frontendEndpoint. Must be a domain name.
          * `resource_state` (`str`) - Resource status.
          * `session_affinity_enabled_state` (`str`) - Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
          * `session_affinity_ttl_seconds` (`float`) - UNUSED. This field will be ignored. The TTL to use in seconds for session affinity, if applicable.
          * `web_application_firewall_policy_link` (`dict`) - Defines the Web Application Firewall policy for each host (if applicable)
            * `id` (`str`) - Resource ID.

        * `type` (`str`) - Resource type.

      * `health_probe_settings` (`list`) - Health probe settings associated with this Front Door instance.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the health probe settings
          * `enabled_state` (`str`) - Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
          * `health_probe_method` (`str`) - Configures which HTTP method to use to probe the backends defined under backendPools.
          * `interval_in_seconds` (`float`) - The number of seconds between health probes.
          * `path` (`str`) - The path to use for the health probe. Default is /
          * `protocol` (`str`) - Protocol scheme to use for this probe
          * `resource_state` (`str`) - Resource status.

        * `type` (`str`) - Resource type.

      * `load_balancing_settings` (`list`) - Load balancing settings associated with this Front Door instance.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the load balancing settings
          * `additional_latency_milliseconds` (`float`) - The additional latency in milliseconds for probes to fall into the lowest latency bucket
          * `resource_state` (`str`) - Resource status.
          * `sample_size` (`float`) - The number of samples to consider for load balancing decisions
          * `successful_samples_required` (`float`) - The number of samples within the sample period that must succeed

        * `type` (`str`) - Resource type.

      * `provisioning_state` (`str`) - Provisioning state of the Front Door.
      * `resource_state` (`str`) - Resource status of the Front Door.
      * `routing_rules` (`list`) - Routing rules associated with this Front Door.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the Front Door Routing Rule
          * `accepted_protocols` (`list`) - Protocol schemes to match for this rule
          * `enabled_state` (`str`) - Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
          * `frontend_endpoints` (`list`) - Frontend endpoints associated with this rule
          * `patterns_to_match` (`list`) - The route patterns of the rule.
          * `resource_state` (`str`) - Resource status.
          * `route_configuration` (`dict`) - A reference to the routing configuration.

        * `type` (`str`) - Resource type.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, backend_pools=None, backend_pools_settings=None, enabled_state=None, friendly_name=None, frontend_endpoints=None, health_probe_settings=None, load_balancing_settings=None, location=None, name=None, resource_group_name=None, resource_state=None, routing_rules=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] backend_pools: Backend pools available to routing rules.
        :param pulumi.Input[dict] backend_pools_settings: Settings for all backendPools
        :param pulumi.Input[str] enabled_state: Operational status of the Front Door load balancer. Permitted values are 'Enabled' or 'Disabled'
        :param pulumi.Input[str] friendly_name: A friendly name for the frontDoor
        :param pulumi.Input[list] frontend_endpoints: Frontend endpoints available to routing rules.
        :param pulumi.Input[list] health_probe_settings: Health probe settings associated with this Front Door instance.
        :param pulumi.Input[list] load_balancing_settings: Load balancing settings associated with this Front Door instance.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Name of the Front Door which is globally unique.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[str] resource_state: Resource status of the Front Door.
        :param pulumi.Input[list] routing_rules: Routing rules associated with this Front Door.
        :param pulumi.Input[dict] tags: Resource tags.

        The **backend_pools** object supports the following:

          * `backends` (`pulumi.Input[list]`) - The set of backends for this pool
            * `address` (`pulumi.Input[str]`) - Location of the backend (IP address or FQDN)
            * `backend_host_header` (`pulumi.Input[str]`) - The value to use as the host header sent to the backend. If blank or unspecified, this defaults to the incoming host.
            * `enabled_state` (`pulumi.Input[str]`) - Whether to enable use of this backend. Permitted values are 'Enabled' or 'Disabled'
            * `http_port` (`pulumi.Input[float]`) - The HTTP TCP port number. Must be between 1 and 65535.
            * `https_port` (`pulumi.Input[float]`) - The HTTPS TCP port number. Must be between 1 and 65535.
            * `priority` (`pulumi.Input[float]`) - Priority to use for load balancing. Higher priorities will not be used for load balancing if any lower priority backend is healthy.
            * `weight` (`pulumi.Input[float]`) - Weight of this endpoint for load balancing purposes.

          * `health_probe_settings` (`pulumi.Input[dict]`) - L7 health probe settings for a backend pool
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `load_balancing_settings` (`pulumi.Input[dict]`) - Load balancing settings for a backend pool
          * `name` (`pulumi.Input[str]`) - Resource name.
          * `resource_state` (`pulumi.Input[str]`) - Resource status.

        The **backend_pools_settings** object supports the following:

          * `enforce_certificate_name_check` (`pulumi.Input[str]`) - Whether to enforce certificate name check on HTTPS requests to all backend pools. No effect on non-HTTPS requests.
          * `send_recv_timeout_seconds` (`pulumi.Input[float]`) - Send and receive timeout on forwarding request to the backend. When timeout is reached, the request fails and returns.

        The **frontend_endpoints** object supports the following:

          * `host_name` (`pulumi.Input[str]`) - The host name of the frontendEndpoint. Must be a domain name.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Resource name.
          * `resource_state` (`pulumi.Input[str]`) - Resource status.
          * `session_affinity_enabled_state` (`pulumi.Input[str]`) - Whether to allow session affinity on this host. Valid options are 'Enabled' or 'Disabled'
          * `session_affinity_ttl_seconds` (`pulumi.Input[float]`) - UNUSED. This field will be ignored. The TTL to use in seconds for session affinity, if applicable.
          * `web_application_firewall_policy_link` (`pulumi.Input[dict]`) - Defines the Web Application Firewall policy for each host (if applicable)
            * `id` (`pulumi.Input[str]`) - Resource ID.

        The **health_probe_settings** object supports the following:

          * `enabled_state` (`pulumi.Input[str]`) - Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
          * `health_probe_method` (`pulumi.Input[str]`) - Configures which HTTP method to use to probe the backends defined under backendPools.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `interval_in_seconds` (`pulumi.Input[float]`) - The number of seconds between health probes.
          * `name` (`pulumi.Input[str]`) - Resource name.
          * `path` (`pulumi.Input[str]`) - The path to use for the health probe. Default is /
          * `protocol` (`pulumi.Input[str]`) - Protocol scheme to use for this probe
          * `resource_state` (`pulumi.Input[str]`) - Resource status.

        The **load_balancing_settings** object supports the following:

          * `additional_latency_milliseconds` (`pulumi.Input[float]`) - The additional latency in milliseconds for probes to fall into the lowest latency bucket
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Resource name.
          * `resource_state` (`pulumi.Input[str]`) - Resource status.
          * `sample_size` (`pulumi.Input[float]`) - The number of samples to consider for load balancing decisions
          * `successful_samples_required` (`pulumi.Input[float]`) - The number of samples within the sample period that must succeed

        The **routing_rules** object supports the following:

          * `accepted_protocols` (`pulumi.Input[list]`) - Protocol schemes to match for this rule
          * `enabled_state` (`pulumi.Input[str]`) - Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
          * `frontend_endpoints` (`pulumi.Input[list]`) - Frontend endpoints associated with this rule
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Resource name.
          * `patterns_to_match` (`pulumi.Input[list]`) - The route patterns of the rule.
          * `resource_state` (`pulumi.Input[str]`) - Resource status.
          * `route_configuration` (`pulumi.Input[dict]`) - A reference to the routing configuration.
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

            __props__['backend_pools'] = backend_pools
            __props__['backend_pools_settings'] = backend_pools_settings
            __props__['enabled_state'] = enabled_state
            __props__['friendly_name'] = friendly_name
            __props__['frontend_endpoints'] = frontend_endpoints
            __props__['health_probe_settings'] = health_probe_settings
            __props__['load_balancing_settings'] = load_balancing_settings
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_state'] = resource_state
            __props__['routing_rules'] = routing_rules
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(FrontDoor, __self__).__init__(
            'azurerm:network/v20190501:FrontDoor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing FrontDoor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FrontDoor(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
