# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class OpenShiftCluster(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The geo-location where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The cluster properties.
      * `apiserver_profile` (`dict`) - The cluster API server profile.
        * `ip` (`str`) - The IP of the cluster API server (immutable).
        * `url` (`str`) - The URL to access the cluster API server (immutable).
        * `visibility` (`str`) - API server visibility (immutable).

      * `cluster_profile` (`dict`) - The cluster profile.
        * `domain` (`str`) - The domain for the cluster (immutable).
        * `pull_secret` (`str`) - The pull secret for the cluster (immutable).
        * `resource_group_id` (`str`) - The ID of the cluster resource group (immutable).
        * `version` (`str`) - The version of the cluster (immutable).

      * `console_profile` (`dict`) - The console profile.
        * `url` (`str`) - The URL to access the cluster console (immutable).

      * `ingress_profiles` (`list`) - The cluster ingress profiles.
        * `ip` (`str`) - The IP of the ingress (immutable).
        * `name` (`str`) - The ingress profile name.  Must be "default" (immutable).
        * `visibility` (`str`) - Ingress visibility (immutable).

      * `master_profile` (`dict`) - The cluster master profile.
        * `subnet_id` (`str`) - The Azure resource ID of the master subnet (immutable).
        * `vm_size` (`str`) - The size of the master VMs (immutable).

      * `network_profile` (`dict`) - The cluster network profile.
        * `pod_cidr` (`str`) - The CIDR used for OpenShift/Kubernetes Pods (immutable).
        * `service_cidr` (`str`) - The CIDR used for OpenShift/Kubernetes Services (immutable).

      * `provisioning_state` (`str`) - The cluster provisioning state (immutable).
      * `service_principal_profile` (`dict`) - The cluster service principal profile.
        * `client_id` (`str`) - The client ID used for the cluster (immutable).
        * `client_secret` (`str`) - The client secret used for the cluster (immutable).

      * `worker_profiles` (`list`) - The cluster worker profiles.
        * `count` (`float`) - The number of worker VMs.  Must be between 3 and 20 (immutable).
        * `disk_size_gb` (`float`) - The disk size of the worker VMs.  Must be 128 or greater (immutable).
        * `name` (`str`) - The worker profile name.  Must be "worker" (immutable).
        * `subnet_id` (`str`) - The Azure resource ID of the worker subnet (immutable).
        * `vm_size` (`str`) - The size of the worker VMs (immutable).
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, apiserver_profile=None, cluster_profile=None, console_profile=None, ingress_profiles=None, location=None, master_profile=None, name=None, network_profile=None, provisioning_state=None, resource_group_name=None, service_principal_profile=None, tags=None, worker_profiles=None, __props__=None, __name__=None, __opts__=None):
        """
        OpenShiftCluster represents an Azure Red Hat OpenShift cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] apiserver_profile: The cluster API server profile.
        :param pulumi.Input[dict] cluster_profile: The cluster profile.
        :param pulumi.Input[dict] console_profile: The console profile.
        :param pulumi.Input[list] ingress_profiles: The cluster ingress profiles.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[dict] master_profile: The cluster master profile.
        :param pulumi.Input[str] name: The name of the OpenShift cluster resource.
        :param pulumi.Input[dict] network_profile: The cluster network profile.
        :param pulumi.Input[str] provisioning_state: The cluster provisioning state (immutable).
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[dict] service_principal_profile: The cluster service principal profile.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[list] worker_profiles: The cluster worker profiles.

        The **apiserver_profile** object supports the following:

          * `ip` (`pulumi.Input[str]`) - The IP of the cluster API server (immutable).
          * `url` (`pulumi.Input[str]`) - The URL to access the cluster API server (immutable).
          * `visibility` (`pulumi.Input[str]`) - API server visibility (immutable).

        The **cluster_profile** object supports the following:

          * `domain` (`pulumi.Input[str]`) - The domain for the cluster (immutable).
          * `pull_secret` (`pulumi.Input[str]`) - The pull secret for the cluster (immutable).
          * `resource_group_id` (`pulumi.Input[str]`) - The ID of the cluster resource group (immutable).
          * `version` (`pulumi.Input[str]`) - The version of the cluster (immutable).

        The **console_profile** object supports the following:

          * `url` (`pulumi.Input[str]`) - The URL to access the cluster console (immutable).

        The **ingress_profiles** object supports the following:

          * `ip` (`pulumi.Input[str]`) - The IP of the ingress (immutable).
          * `name` (`pulumi.Input[str]`) - The ingress profile name.  Must be "default" (immutable).
          * `visibility` (`pulumi.Input[str]`) - Ingress visibility (immutable).

        The **master_profile** object supports the following:

          * `subnet_id` (`pulumi.Input[str]`) - The Azure resource ID of the master subnet (immutable).
          * `vm_size` (`pulumi.Input[str]`) - The size of the master VMs (immutable).

        The **network_profile** object supports the following:

          * `pod_cidr` (`pulumi.Input[str]`) - The CIDR used for OpenShift/Kubernetes Pods (immutable).
          * `service_cidr` (`pulumi.Input[str]`) - The CIDR used for OpenShift/Kubernetes Services (immutable).

        The **service_principal_profile** object supports the following:

          * `client_id` (`pulumi.Input[str]`) - The client ID used for the cluster (immutable).
          * `client_secret` (`pulumi.Input[str]`) - The client secret used for the cluster (immutable).

        The **worker_profiles** object supports the following:

          * `count` (`pulumi.Input[float]`) - The number of worker VMs.  Must be between 3 and 20 (immutable).
          * `disk_size_gb` (`pulumi.Input[float]`) - The disk size of the worker VMs.  Must be 128 or greater (immutable).
          * `name` (`pulumi.Input[str]`) - The worker profile name.  Must be "worker" (immutable).
          * `subnet_id` (`pulumi.Input[str]`) - The Azure resource ID of the worker subnet (immutable).
          * `vm_size` (`pulumi.Input[str]`) - The size of the worker VMs (immutable).
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

            __props__['apiserver_profile'] = apiserver_profile
            __props__['cluster_profile'] = cluster_profile
            __props__['console_profile'] = console_profile
            __props__['ingress_profiles'] = ingress_profiles
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['master_profile'] = master_profile
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_profile'] = network_profile
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_principal_profile'] = service_principal_profile
            __props__['tags'] = tags
            __props__['worker_profiles'] = worker_profiles
            __props__['properties'] = None
            __props__['type'] = None
        super(OpenShiftCluster, __self__).__init__(
            'azurerm:redhatopenshift/v20200430:OpenShiftCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing OpenShiftCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return OpenShiftCluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop