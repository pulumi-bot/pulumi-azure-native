# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ContainerGroup(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the container group, if configured.
      * `principal_id` (`str`) - The principal id of the container group identity. This property will only be provided for a system assigned identity.
      * `tenant_id` (`str`) - The tenant id associated with the container group. This property will only be provided for a system assigned identity.
      * `type` (`str`) - The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the container group.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with the container group. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    The resource name.
    """
    properties: pulumi.Output[dict]
    """
    The container group properties
      * `containers` (`list`) - The containers within the container group.
        * `name` (`str`) - The user-provided name of the container instance.
        * `properties` (`dict`) - The properties of the container instance.
          * `command` (`list`) - The commands to execute within the container instance in exec form.
          * `environment_variables` (`list`) - The environment variables to set in the container instance.
            * `name` (`str`) - The name of the environment variable.
            * `secure_value` (`str`) - The value of the secure environment variable.
            * `value` (`str`) - The value of the environment variable.

          * `image` (`str`) - The name of the image used to create the container instance.
          * `instance_view` (`dict`) - The instance view of the container instance. Only valid in response.
            * `current_state` (`dict`) - Current container instance state.
              * `detail_status` (`str`) - The human-readable status of the container instance state.
              * `exit_code` (`float`) - The container instance exit codes correspond to those from the `docker run` command.
              * `finish_time` (`str`) - The date-time when the container instance state finished.
              * `start_time` (`str`) - The date-time when the container instance state started.
              * `state` (`str`) - The state of the container instance.

            * `events` (`list`) - The events of the container instance.
              * `count` (`float`) - The count of the event.
              * `first_timestamp` (`str`) - The date-time of the earliest logged event.
              * `last_timestamp` (`str`) - The date-time of the latest logged event.
              * `message` (`str`) - The event message.
              * `name` (`str`) - The event name.
              * `type` (`str`) - The event type.

            * `previous_state` (`dict`) - Previous container instance state.
            * `restart_count` (`float`) - The number of times that the container instance has been restarted.

          * `liveness_probe` (`dict`) - The liveness probe.
            * `exec` (`dict`) - The execution command to probe
              * `command` (`list`) - The commands to execute within the container.

            * `failure_threshold` (`float`) - The failure threshold.
            * `http_get` (`dict`) - The Http Get settings to probe
              * `path` (`str`) - The path to probe.
              * `port` (`float`) - The port number to probe.
              * `scheme` (`str`) - The scheme.

            * `initial_delay_seconds` (`float`) - The initial delay seconds.
            * `period_seconds` (`float`) - The period seconds.
            * `success_threshold` (`float`) - The success threshold.
            * `timeout_seconds` (`float`) - The timeout seconds.

          * `ports` (`list`) - The exposed ports on the container instance.
            * `port` (`float`) - The port number exposed within the container group.
            * `protocol` (`str`) - The protocol associated with the port.

          * `readiness_probe` (`dict`) - The readiness probe.
          * `resources` (`dict`) - The resource requirements of the container instance.
            * `limits` (`dict`) - The resource limits of this container instance.
              * `cpu` (`float`) - The CPU limit of this container instance.
              * `gpu` (`dict`) - The GPU limit of this container instance.
                * `count` (`float`) - The count of the GPU resource.
                * `sku` (`str`) - The SKU of the GPU resource.

              * `memory_in_gb` (`float`) - The memory limit in GB of this container instance.

            * `requests` (`dict`) - The resource requests of this container instance.
              * `cpu` (`float`) - The CPU request of this container instance.
              * `gpu` (`dict`) - The GPU request of this container instance.
              * `memory_in_gb` (`float`) - The memory request in GB of this container instance.

          * `volume_mounts` (`list`) - The volume mounts available to the container instance.
            * `mount_path` (`str`) - The path within the container where the volume should be mounted. Must not contain colon (:).
            * `name` (`str`) - The name of the volume mount.
            * `read_only` (`bool`) - The flag indicating whether the volume mount is read-only.

      * `diagnostics` (`dict`) - The diagnostic information for a container group.
        * `log_analytics` (`dict`) - Container group log analytics information.
          * `log_type` (`str`) - The log type to be used.
          * `metadata` (`dict`) - Metadata for log analytics.
          * `workspace_id` (`str`) - The workspace id for log analytics
          * `workspace_key` (`str`) - The workspace key for log analytics

      * `dns_config` (`dict`) - The DNS config information for a container group.
        * `name_servers` (`list`) - The DNS servers for the container group.
        * `options` (`str`) - The DNS options for the container group.
        * `search_domains` (`str`) - The DNS search domains for hostname lookup in the container group.

      * `encryption_properties` (`dict`) - The encryption properties for a container group.
        * `key_name` (`str`) - The encryption key name.
        * `key_version` (`str`) - The encryption key version.
        * `vault_base_url` (`str`) - The keyvault base url.

      * `image_registry_credentials` (`list`) - The image registry credentials by which the container group is created from.
        * `password` (`str`) - The password for the private registry.
        * `server` (`str`) - The Docker image registry server without a protocol such as "http" and "https".
        * `username` (`str`) - The username for the private registry.

      * `init_containers` (`list`) - The init containers for a container group.
        * `name` (`str`) - The name for the init container.
        * `properties` (`dict`) - The properties for the init container.
          * `command` (`list`) - The command to execute within the init container in exec form.
          * `environment_variables` (`list`) - The environment variables to set in the init container.
          * `image` (`str`) - The image of the init container.
          * `instance_view` (`dict`) - The instance view of the init container. Only valid in response.
            * `current_state` (`dict`) - The current state of the init container.
            * `events` (`list`) - The events of the init container.
            * `previous_state` (`dict`) - The previous state of the init container.
            * `restart_count` (`float`) - The number of times that the init container has been restarted.

          * `volume_mounts` (`list`) - The volume mounts available to the init container.

      * `instance_view` (`dict`) - The instance view of the container group. Only valid in response.
      * `ip_address` (`dict`) - The IP address type of the container group.
        * `dns_name_label` (`str`) - The Dns name label for the IP.
        * `fqdn` (`str`) - The FQDN for the IP.
        * `ip` (`str`) - The IP exposed to the public internet.
        * `ports` (`list`) - The list of ports exposed on the container group.
          * `port` (`float`) - The port number.
          * `protocol` (`str`) - The protocol associated with the port.

        * `type` (`str`) - Specifies if the IP is exposed to the public internet or private VNET.

      * `network_profile` (`dict`) - The network profile information for a container group.
        * `id` (`str`) - The identifier for a network profile.

      * `os_type` (`str`) - The operating system type required by the containers in the container group.
      * `provisioning_state` (`str`) - The provisioning state of the container group. This only appears in the response.
      * `restart_policy` (`str`) - Restart policy for all containers within the container group. 
        - `Always` Always restart
        - `OnFailure` Restart on failure
        - `Never` Never restart
      * `sku` (`str`) - The SKU for a container group.
      * `volumes` (`list`) - The list of volumes that can be mounted by containers in this container group.
        * `azure_file` (`dict`) - The Azure File volume.
          * `read_only` (`bool`) - The flag indicating whether the Azure File shared mounted as a volume is read-only.
          * `share_name` (`str`) - The name of the Azure File share to be mounted as a volume.
          * `storage_account_key` (`str`) - The storage account access key used to access the Azure File share.
          * `storage_account_name` (`str`) - The name of the storage account that contains the Azure File share.

        * `empty_dir` (`dict`) - The empty directory volume.
        * `git_repo` (`dict`) - The git repo volume.
          * `directory` (`str`) - Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
          * `repository` (`str`) - Repository URL
          * `revision` (`str`) - Commit hash for the specified revision.

        * `name` (`str`) - The name of the volume.
        * `secret` (`dict`) - The secret volume.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, containers=None, diagnostics=None, dns_config=None, encryption_properties=None, identity=None, image_registry_credentials=None, init_containers=None, ip_address=None, location=None, name=None, network_profile=None, os_type=None, resource_group_name=None, restart_policy=None, sku=None, tags=None, volumes=None, __props__=None, __name__=None, __opts__=None):
        """
        A container group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] containers: The containers within the container group.
        :param pulumi.Input[dict] diagnostics: The diagnostic information for a container group.
        :param pulumi.Input[dict] dns_config: The DNS config information for a container group.
        :param pulumi.Input[dict] encryption_properties: The encryption properties for a container group.
        :param pulumi.Input[dict] identity: The identity of the container group, if configured.
        :param pulumi.Input[list] image_registry_credentials: The image registry credentials by which the container group is created from.
        :param pulumi.Input[list] init_containers: The init containers for a container group.
        :param pulumi.Input[dict] ip_address: The IP address type of the container group.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The name of the container group.
        :param pulumi.Input[dict] network_profile: The network profile information for a container group.
        :param pulumi.Input[str] os_type: The operating system type required by the containers in the container group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] restart_policy: Restart policy for all containers within the container group. 
               - `Always` Always restart
               - `OnFailure` Restart on failure
               - `Never` Never restart
        :param pulumi.Input[str] sku: The SKU for a container group.
        :param pulumi.Input[dict] tags: The resource tags.
        :param pulumi.Input[list] volumes: The list of volumes that can be mounted by containers in this container group.

        The **containers** object supports the following:

          * `command` (`pulumi.Input[list]`) - The commands to execute within the container instance in exec form.
          * `environment_variables` (`pulumi.Input[list]`) - The environment variables to set in the container instance.
            * `name` (`pulumi.Input[str]`) - The name of the environment variable.
            * `secure_value` (`pulumi.Input[str]`) - The value of the secure environment variable.
            * `value` (`pulumi.Input[str]`) - The value of the environment variable.

          * `image` (`pulumi.Input[str]`) - The name of the image used to create the container instance.
          * `liveness_probe` (`pulumi.Input[dict]`) - The liveness probe.
            * `exec` (`pulumi.Input[dict]`) - The execution command to probe
              * `command` (`pulumi.Input[list]`) - The commands to execute within the container.

            * `failure_threshold` (`pulumi.Input[float]`) - The failure threshold.
            * `http_get` (`pulumi.Input[dict]`) - The Http Get settings to probe
              * `path` (`pulumi.Input[str]`) - The path to probe.
              * `port` (`pulumi.Input[float]`) - The port number to probe.
              * `scheme` (`pulumi.Input[str]`) - The scheme.

            * `initial_delay_seconds` (`pulumi.Input[float]`) - The initial delay seconds.
            * `period_seconds` (`pulumi.Input[float]`) - The period seconds.
            * `success_threshold` (`pulumi.Input[float]`) - The success threshold.
            * `timeout_seconds` (`pulumi.Input[float]`) - The timeout seconds.

          * `name` (`pulumi.Input[str]`) - The user-provided name of the container instance.
          * `ports` (`pulumi.Input[list]`) - The exposed ports on the container instance.
            * `port` (`pulumi.Input[float]`) - The port number exposed within the container group.
            * `protocol` (`pulumi.Input[str]`) - The protocol associated with the port.

          * `readiness_probe` (`pulumi.Input[dict]`) - The readiness probe.
          * `resources` (`pulumi.Input[dict]`) - The resource requirements of the container instance.
            * `limits` (`pulumi.Input[dict]`) - The resource limits of this container instance.
              * `cpu` (`pulumi.Input[float]`) - The CPU limit of this container instance.
              * `gpu` (`pulumi.Input[dict]`) - The GPU limit of this container instance.
                * `count` (`pulumi.Input[float]`) - The count of the GPU resource.
                * `sku` (`pulumi.Input[str]`) - The SKU of the GPU resource.

              * `memory_in_gb` (`pulumi.Input[float]`) - The memory limit in GB of this container instance.

            * `requests` (`pulumi.Input[dict]`) - The resource requests of this container instance.
              * `cpu` (`pulumi.Input[float]`) - The CPU request of this container instance.
              * `gpu` (`pulumi.Input[dict]`) - The GPU request of this container instance.
              * `memory_in_gb` (`pulumi.Input[float]`) - The memory request in GB of this container instance.

          * `volume_mounts` (`pulumi.Input[list]`) - The volume mounts available to the container instance.
            * `mount_path` (`pulumi.Input[str]`) - The path within the container where the volume should be mounted. Must not contain colon (:).
            * `name` (`pulumi.Input[str]`) - The name of the volume mount.
            * `read_only` (`pulumi.Input[bool]`) - The flag indicating whether the volume mount is read-only.

        The **diagnostics** object supports the following:

          * `log_analytics` (`pulumi.Input[dict]`) - Container group log analytics information.
            * `log_type` (`pulumi.Input[str]`) - The log type to be used.
            * `metadata` (`pulumi.Input[dict]`) - Metadata for log analytics.
            * `workspace_id` (`pulumi.Input[str]`) - The workspace id for log analytics
            * `workspace_key` (`pulumi.Input[str]`) - The workspace key for log analytics

        The **dns_config** object supports the following:

          * `name_servers` (`pulumi.Input[list]`) - The DNS servers for the container group.
          * `options` (`pulumi.Input[str]`) - The DNS options for the container group.
          * `search_domains` (`pulumi.Input[str]`) - The DNS search domains for hostname lookup in the container group.

        The **encryption_properties** object supports the following:

          * `key_name` (`pulumi.Input[str]`) - The encryption key name.
          * `key_version` (`pulumi.Input[str]`) - The encryption key version.
          * `vault_base_url` (`pulumi.Input[str]`) - The keyvault base url.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the container group.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with the container group. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **image_registry_credentials** object supports the following:

          * `password` (`pulumi.Input[str]`) - The password for the private registry.
          * `server` (`pulumi.Input[str]`) - The Docker image registry server without a protocol such as "http" and "https".
          * `username` (`pulumi.Input[str]`) - The username for the private registry.

        The **init_containers** object supports the following:

          * `command` (`pulumi.Input[list]`) - The command to execute within the init container in exec form.
          * `environment_variables` (`pulumi.Input[list]`) - The environment variables to set in the init container.
          * `image` (`pulumi.Input[str]`) - The image of the init container.
          * `name` (`pulumi.Input[str]`) - The name for the init container.
          * `volume_mounts` (`pulumi.Input[list]`) - The volume mounts available to the init container.

        The **ip_address** object supports the following:

          * `dns_name_label` (`pulumi.Input[str]`) - The Dns name label for the IP.
          * `ip` (`pulumi.Input[str]`) - The IP exposed to the public internet.
          * `ports` (`pulumi.Input[list]`) - The list of ports exposed on the container group.
            * `port` (`pulumi.Input[float]`) - The port number.
            * `protocol` (`pulumi.Input[str]`) - The protocol associated with the port.

          * `type` (`pulumi.Input[str]`) - Specifies if the IP is exposed to the public internet or private VNET.

        The **network_profile** object supports the following:

          * `id` (`pulumi.Input[str]`) - The identifier for a network profile.

        The **volumes** object supports the following:

          * `azure_file` (`pulumi.Input[dict]`) - The Azure File volume.
            * `read_only` (`pulumi.Input[bool]`) - The flag indicating whether the Azure File shared mounted as a volume is read-only.
            * `share_name` (`pulumi.Input[str]`) - The name of the Azure File share to be mounted as a volume.
            * `storage_account_key` (`pulumi.Input[str]`) - The storage account access key used to access the Azure File share.
            * `storage_account_name` (`pulumi.Input[str]`) - The name of the storage account that contains the Azure File share.

          * `empty_dir` (`pulumi.Input[dict]`) - The empty directory volume.
          * `git_repo` (`pulumi.Input[dict]`) - The git repo volume.
            * `directory` (`pulumi.Input[str]`) - Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
            * `repository` (`pulumi.Input[str]`) - Repository URL
            * `revision` (`pulumi.Input[str]`) - Commit hash for the specified revision.

          * `name` (`pulumi.Input[str]`) - The name of the volume.
          * `secret` (`pulumi.Input[dict]`) - The secret volume.
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

            if containers is None:
                raise TypeError("Missing required property 'containers'")
            __props__['containers'] = containers
            __props__['diagnostics'] = diagnostics
            __props__['dns_config'] = dns_config
            __props__['encryption_properties'] = encryption_properties
            __props__['identity'] = identity
            __props__['image_registry_credentials'] = image_registry_credentials
            __props__['init_containers'] = init_containers
            __props__['ip_address'] = ip_address
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_profile'] = network_profile
            if os_type is None:
                raise TypeError("Missing required property 'os_type'")
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restart_policy'] = restart_policy
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['volumes'] = volumes
            __props__['properties'] = None
            __props__['type'] = None
        super(ContainerGroup, __self__).__init__(
            'azurerm:containerinstance/v20191201:ContainerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ContainerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ContainerGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
