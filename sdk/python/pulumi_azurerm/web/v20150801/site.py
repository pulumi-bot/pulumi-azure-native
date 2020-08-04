# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Site(pulumi.CustomResource):
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
    def __init__(__self__, resource_name, opts=None, force_dns_registration=None, id=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, skip_custom_domain_verification=None, skip_dns_registration=None, tags=None, ttl_in_seconds=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a web app

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] force_dns_registration: If true, web app hostname is force registered with DNS
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] name: Resource Name
        :param pulumi.Input[str] resource_group_name: Name of the resource group
        :param pulumi.Input[str] skip_custom_domain_verification: If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
        :param pulumi.Input[str] skip_dns_registration: If true web app hostname is not registered with DNS on creation. This parameter is
                           only used for app creation
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[str] ttl_in_seconds: Time to live in seconds for web app's default domain name
        :param pulumi.Input[str] type: Resource type

        The **properties** object supports the following:

          * `client_affinity_enabled` (`pulumi.Input[bool]`) - Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
          * `client_cert_enabled` (`pulumi.Input[bool]`) - Specifies if the client certificate is enabled for the web app
          * `cloning_info` (`pulumi.Input[dict]`) - This is only valid for web app creation. If specified, web app is cloned from 
                        a source web app
            * `app_settings_overrides` (`pulumi.Input[dict]`) - Application settings overrides for cloned web app. If specified these settings will override the settings cloned 
                          from source web app. If not specified, application settings from source web app are retained.
            * `clone_custom_host_names` (`pulumi.Input[bool]`) - If true, clone custom hostnames from source web app
            * `clone_source_control` (`pulumi.Input[bool]`) - Clone source control from source web app
            * `configure_load_balancing` (`pulumi.Input[bool]`) - If specified configure load balancing for source and clone site
            * `correlation_id` (`pulumi.Input[str]`) - Correlation Id of cloning operation. This id ties multiple cloning operations
                          together to use the same snapshot
            * `hosting_environment` (`pulumi.Input[str]`) - Hosting environment
            * `overwrite` (`pulumi.Input[bool]`) - Overwrite destination web app
            * `source_web_app_id` (`pulumi.Input[str]`) - ARM resource id of the source web app. Web app resource id is of the form 
                          /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName} for production slots and 
                          /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slotName} for other slots
            * `traffic_manager_profile_id` (`pulumi.Input[str]`) - ARM resource id of the traffic manager profile to use if it exists. Traffic manager resource id is of the form 
                          /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{profileName}
            * `traffic_manager_profile_name` (`pulumi.Input[str]`) - Name of traffic manager profile to create. This is only needed if traffic manager profile does not already exist

          * `container_size` (`pulumi.Input[float]`) - Size of a function container
          * `enabled` (`pulumi.Input[bool]`) - True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
          * `gateway_site_name` (`pulumi.Input[str]`) - Name of gateway app associated with web app
          * `host_name_ssl_states` (`pulumi.Input[list]`) - Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
            * `name` (`pulumi.Input[str]`) - Host name
            * `ssl_state` (`pulumi.Input[str]`) - SSL type
            * `thumbprint` (`pulumi.Input[str]`) - SSL cert thumbprint
            * `to_update` (`pulumi.Input[bool]`) - Set this flag to update existing host name
            * `virtual_ip` (`pulumi.Input[str]`) - Virtual IP address assigned to the host name if IP based SSL is enabled

          * `host_names_disabled` (`pulumi.Input[bool]`) - Specifies if the public hostnames are disabled the web app.
                        If set to true the app is only accessible via API Management process
          * `hosting_environment_profile` (`pulumi.Input[dict]`) - Specification for the hosting environment (App Service Environment) to use for the web app
            * `id` (`pulumi.Input[str]`) - Resource id of the hostingEnvironment (App Service Environment)
            * `name` (`pulumi.Input[str]`) - Name of the hostingEnvironment (App Service Environment) (read only)
            * `type` (`pulumi.Input[str]`) - Resource type of the hostingEnvironment (App Service Environment) (read only)

          * `max_number_of_workers` (`pulumi.Input[float]`) - Maximum number of workers
                        This only applies to function container
          * `micro_service` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of web app
          * `scm_site_also_stopped` (`pulumi.Input[bool]`) - If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
          * `server_farm_id` (`pulumi.Input[str]`)
          * `site_config` (`pulumi.Input[dict]`) - Configuration of web app
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `kind` (`pulumi.Input[str]`) - Kind of resource
            * `location` (`pulumi.Input[str]`) - Resource Location
            * `name` (`pulumi.Input[str]`) - Resource Name
            * `properties` (`pulumi.Input[dict]`)
              * `always_on` (`pulumi.Input[bool]`) - Always On
              * `api_definition` (`pulumi.Input[dict]`) - Information about the formal API definition for the web app.
                * `url` (`pulumi.Input[str]`) - The URL of the API definition.

              * `app_command_line` (`pulumi.Input[str]`) - App Command Line to launch
              * `app_settings` (`pulumi.Input[list]`) - Application Settings
                * `name` (`pulumi.Input[str]`) - Pair name
                * `value` (`pulumi.Input[str]`) - Pair value

              * `auto_heal_enabled` (`pulumi.Input[bool]`) - Auto heal enabled
              * `auto_heal_rules` (`pulumi.Input[dict]`) - Auto heal rules
                * `actions` (`pulumi.Input[dict]`) - Actions - Actions to be executed when a rule is triggered
                  * `action_type` (`pulumi.Input[str]`) - ActionType - predefined action to be taken
                  * `custom_action` (`pulumi.Input[dict]`) - CustomAction - custom action to be taken
                    * `exe` (`pulumi.Input[str]`) - Executable to be run
                    * `parameters` (`pulumi.Input[str]`) - Parameters for the executable

                  * `min_process_execution_time` (`pulumi.Input[str]`) - MinProcessExecutionTime - minimum time the process must execute
                                before taking the action

                * `triggers` (`pulumi.Input[dict]`) - Triggers - Conditions that describe when to execute the auto-heal actions
                  * `private_bytes_in_kb` (`pulumi.Input[float]`) - PrivateBytesInKB - Defines a rule based on private bytes
                  * `requests` (`pulumi.Input[dict]`) - Requests - Defines a rule based on total requests
                    * `count` (`pulumi.Input[float]`) - Count
                    * `time_interval` (`pulumi.Input[str]`) - TimeInterval

                  * `slow_requests` (`pulumi.Input[dict]`) - SlowRequests - Defines a rule based on request execution time
                    * `count` (`pulumi.Input[float]`) - Count
                    * `time_interval` (`pulumi.Input[str]`) - TimeInterval
                    * `time_taken` (`pulumi.Input[str]`) - TimeTaken

                  * `status_codes` (`pulumi.Input[list]`) - StatusCodes - Defines a rule based on status codes
                    * `count` (`pulumi.Input[float]`) - Count
                    * `status` (`pulumi.Input[float]`) - HTTP status code
                    * `sub_status` (`pulumi.Input[float]`) - SubStatus
                    * `time_interval` (`pulumi.Input[str]`) - TimeInterval
                    * `win32_status` (`pulumi.Input[float]`) - Win32 error code

              * `auto_swap_slot_name` (`pulumi.Input[str]`) - Auto swap slot name
              * `connection_strings` (`pulumi.Input[list]`) - Connection strings
                * `connection_string` (`pulumi.Input[str]`) - Connection string value
                * `name` (`pulumi.Input[str]`) - Name of connection string
                * `type` (`pulumi.Input[str]`) - Type of database

              * `cors` (`pulumi.Input[dict]`) - Cross-Origin Resource Sharing (CORS) settings.
                * `allowed_origins` (`pulumi.Input[list]`) - Gets or sets the list of origins that should be allowed to make cross-origin
                              calls (for example: http://example.com:12345). Use "*" to allow all.

              * `default_documents` (`pulumi.Input[list]`) - Default documents
              * `detailed_error_logging_enabled` (`pulumi.Input[bool]`) - Detailed error logging enabled
              * `document_root` (`pulumi.Input[str]`) - Document root
              * `experiments` (`pulumi.Input[dict]`) - This is work around for polymorphic types
                * `ramp_up_rules` (`pulumi.Input[list]`) - List of {Microsoft.Web.Hosting.Administration.RampUpRule} objects.
                  * `action_host_name` (`pulumi.Input[str]`) - Hostname of a slot to which the traffic will be redirected if decided to. E.g. mysite-stage.azurewebsites.net
                  * `change_decision_callback_url` (`pulumi.Input[str]`) - Custom decision algorithm can be provided in TiPCallback site extension which Url can be specified. See TiPCallback site extension for the scaffold and contracts.
                                https://www.siteextensions.net/packages/TiPCallback/
                  * `change_interval_in_minutes` (`pulumi.Input[float]`) - [Optional] Specifies interval in minutes to reevaluate ReroutePercentage
                  * `change_step` (`pulumi.Input[float]`) - [Optional] In auto ramp up scenario this is the step to add/remove from {Microsoft.Web.Hosting.Administration.RampUpRule.ReroutePercentage} until it reaches 
                                {Microsoft.Web.Hosting.Administration.RampUpRule.MinReroutePercentage} or {Microsoft.Web.Hosting.Administration.RampUpRule.MaxReroutePercentage}. Site metrics are checked every N minutes specified in {Microsoft.Web.Hosting.Administration.RampUpRule.ChangeIntervalInMinutes}.
                                Custom decision algorithm can be provided in TiPCallback site extension which Url can be specified in {Microsoft.Web.Hosting.Administration.RampUpRule.ChangeDecisionCallbackUrl}
                  * `max_reroute_percentage` (`pulumi.Input[float]`) - [Optional] Specifies upper boundary below which ReroutePercentage will stay.
                  * `min_reroute_percentage` (`pulumi.Input[float]`) - [Optional] Specifies lower boundary above which ReroutePercentage will stay.
                  * `name` (`pulumi.Input[str]`) - Name of the routing rule. The recommended name would be to point to the slot which will receive the traffic in the experiment.
                  * `reroute_percentage` (`pulumi.Input[float]`) - Percentage of the traffic which will be redirected to {Microsoft.Web.Hosting.Administration.RampUpRule.ActionHostName}

              * `handler_mappings` (`pulumi.Input[list]`) - Handler mappings
                * `arguments` (`pulumi.Input[str]`) - Command-line arguments to be passed to the script processor.
                * `extension` (`pulumi.Input[str]`) - Requests with this extension will be handled using the specified FastCGI application.
                * `script_processor` (`pulumi.Input[str]`) - The absolute path to the FastCGI application.

              * `http_logging_enabled` (`pulumi.Input[bool]`) - HTTP logging Enabled
              * `ip_security_restrictions` (`pulumi.Input[list]`) - Ip Security restrictions
                * `ip_address` (`pulumi.Input[str]`) - IP address the security restriction is valid for
                * `subnet_mask` (`pulumi.Input[str]`) - Subnet mask for the range of IP addresses the restriction is valid for

              * `java_container` (`pulumi.Input[str]`) - Java container
              * `java_container_version` (`pulumi.Input[str]`) - Java container version
              * `java_version` (`pulumi.Input[str]`) - Java version
              * `limits` (`pulumi.Input[dict]`) - Site limits
                * `max_disk_size_in_mb` (`pulumi.Input[float]`) - Maximum allowed disk size usage in MB
                * `max_memory_in_mb` (`pulumi.Input[float]`) - Maximum allowed memory usage in MB
                * `max_percentage_cpu` (`pulumi.Input[float]`) - Maximum allowed CPU usage percentage

              * `load_balancing` (`pulumi.Input[str]`) - Site load balancing
              * `local_my_sql_enabled` (`pulumi.Input[bool]`) - Local mysql enabled
              * `logs_directory_size_limit` (`pulumi.Input[float]`) - HTTP Logs Directory size limit
              * `managed_pipeline_mode` (`pulumi.Input[str]`) - Managed pipeline mode
              * `metadata` (`pulumi.Input[list]`) - Site Metadata
              * `net_framework_version` (`pulumi.Input[str]`) - Net Framework Version
              * `node_version` (`pulumi.Input[str]`) - Version of Node
              * `number_of_workers` (`pulumi.Input[float]`) - Number of workers
              * `php_version` (`pulumi.Input[str]`) - Version of PHP
              * `publishing_password` (`pulumi.Input[str]`) - Publishing password
              * `publishing_username` (`pulumi.Input[str]`) - Publishing user name
              * `python_version` (`pulumi.Input[str]`) - Version of Python
              * `remote_debugging_enabled` (`pulumi.Input[bool]`) - Remote Debugging Enabled
              * `remote_debugging_version` (`pulumi.Input[str]`) - Remote Debugging Version
              * `request_tracing_enabled` (`pulumi.Input[bool]`) - Enable request tracing
              * `request_tracing_expiration_time` (`pulumi.Input[str]`) - Request tracing expiration time
              * `scm_type` (`pulumi.Input[str]`) - SCM type
              * `tracing_options` (`pulumi.Input[str]`) - Tracing options
              * `use32_bit_worker_process` (`pulumi.Input[bool]`) - Use 32 bit worker process
              * `virtual_applications` (`pulumi.Input[list]`) - Virtual applications
                * `physical_path` (`pulumi.Input[str]`)
                * `preload_enabled` (`pulumi.Input[bool]`)
                * `virtual_directories` (`pulumi.Input[list]`)
                  * `physical_path` (`pulumi.Input[str]`)
                  * `virtual_path` (`pulumi.Input[str]`)

                * `virtual_path` (`pulumi.Input[str]`)

              * `vnet_name` (`pulumi.Input[str]`) - Vnet name
              * `web_sockets_enabled` (`pulumi.Input[bool]`) - Web socket enabled.

            * `tags` (`pulumi.Input[dict]`) - Resource tags
            * `type` (`pulumi.Input[str]`) - Resource type
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

            __props__['force_dns_registration'] = force_dns_registration
            __props__['id'] = id
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['skip_custom_domain_verification'] = skip_custom_domain_verification
            __props__['skip_dns_registration'] = skip_dns_registration
            __props__['tags'] = tags
            __props__['ttl_in_seconds'] = ttl_in_seconds
            __props__['type'] = type
        super(Site, __self__).__init__(
            'azurerm:web/v20150801:Site',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Site resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Site(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
