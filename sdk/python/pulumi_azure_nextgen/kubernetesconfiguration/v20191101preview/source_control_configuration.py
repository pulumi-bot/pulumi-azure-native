# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SourceControlConfiguration']


class SourceControlConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 cluster_resource_name: Optional[pulumi.Input[str]] = None,
                 cluster_rp: Optional[pulumi.Input[str]] = None,
                 enable_helm_operator: Optional[pulumi.Input[str]] = None,
                 helm_operator_properties: Optional[pulumi.Input[pulumi.InputType['HelmOperatorPropertiesArgs']]] = None,
                 operator_instance_name: Optional[pulumi.Input[str]] = None,
                 operator_namespace: Optional[pulumi.Input[str]] = None,
                 operator_params: Optional[pulumi.Input[str]] = None,
                 operator_scope: Optional[pulumi.Input[str]] = None,
                 operator_type: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source_control_configuration_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The SourceControl Configuration object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the kubernetes cluster.
        :param pulumi.Input[str] cluster_resource_name: The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
        :param pulumi.Input[str] cluster_rp: The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
        :param pulumi.Input[str] enable_helm_operator: Option to enable Helm Operator for this git configuration.
        :param pulumi.Input[pulumi.InputType['HelmOperatorPropertiesArgs']] helm_operator_properties: Properties for Helm operator.
        :param pulumi.Input[str] operator_instance_name: Instance name of the operator - identifying the specific configuration.
        :param pulumi.Input[str] operator_namespace: The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
        :param pulumi.Input[str] operator_params: Any Parameters for the Operator instance in string format.
        :param pulumi.Input[str] operator_scope: Scope at which the operator will be installed.
        :param pulumi.Input[str] operator_type: Type of the operator
        :param pulumi.Input[str] repository_url: Url of the SourceControl Repository.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] source_control_configuration_name: Name of the Source Control Configuration.
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

            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            if cluster_resource_name is None:
                raise TypeError("Missing required property 'cluster_resource_name'")
            __props__['cluster_resource_name'] = cluster_resource_name
            if cluster_rp is None:
                raise TypeError("Missing required property 'cluster_rp'")
            __props__['cluster_rp'] = cluster_rp
            __props__['enable_helm_operator'] = enable_helm_operator
            __props__['helm_operator_properties'] = helm_operator_properties
            __props__['operator_instance_name'] = operator_instance_name
            __props__['operator_namespace'] = operator_namespace
            __props__['operator_params'] = operator_params
            __props__['operator_scope'] = operator_scope
            __props__['operator_type'] = operator_type
            __props__['repository_url'] = repository_url
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_control_configuration_name is None:
                raise TypeError("Missing required property 'source_control_configuration_name'")
            __props__['source_control_configuration_name'] = source_control_configuration_name
            __props__['compliance_status'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['repository_public_key'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:kubernetesconfiguration/v20201001preview:SourceControlConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SourceControlConfiguration, __self__).__init__(
            'azure-nextgen:kubernetesconfiguration/v20191101preview:SourceControlConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SourceControlConfiguration':
        """
        Get an existing SourceControlConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SourceControlConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="complianceStatus")
    def compliance_status(self) -> pulumi.Output['outputs.ComplianceStatusResponse']:
        """
        Compliance Status of the Configuration
        """
        return pulumi.get(self, "compliance_status")

    @property
    @pulumi.getter(name="enableHelmOperator")
    def enable_helm_operator(self) -> pulumi.Output[Optional[str]]:
        """
        Option to enable Helm Operator for this git configuration.
        """
        return pulumi.get(self, "enable_helm_operator")

    @property
    @pulumi.getter(name="helmOperatorProperties")
    def helm_operator_properties(self) -> pulumi.Output[Optional['outputs.HelmOperatorPropertiesResponse']]:
        """
        Properties for Helm operator.
        """
        return pulumi.get(self, "helm_operator_properties")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operatorInstanceName")
    def operator_instance_name(self) -> pulumi.Output[Optional[str]]:
        """
        Instance name of the operator - identifying the specific configuration.
        """
        return pulumi.get(self, "operator_instance_name")

    @property
    @pulumi.getter(name="operatorNamespace")
    def operator_namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
        """
        return pulumi.get(self, "operator_namespace")

    @property
    @pulumi.getter(name="operatorParams")
    def operator_params(self) -> pulumi.Output[Optional[str]]:
        """
        Any Parameters for the Operator instance in string format.
        """
        return pulumi.get(self, "operator_params")

    @property
    @pulumi.getter(name="operatorScope")
    def operator_scope(self) -> pulumi.Output[Optional[str]]:
        """
        Scope at which the operator will be installed.
        """
        return pulumi.get(self, "operator_scope")

    @property
    @pulumi.getter(name="operatorType")
    def operator_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of the operator
        """
        return pulumi.get(self, "operator_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the resource provider.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="repositoryPublicKey")
    def repository_public_key(self) -> pulumi.Output[str]:
        """
        Public Key associated with this SourceControl configuration (either generated within the cluster or provided by the user).
        """
        return pulumi.get(self, "repository_public_key")

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Output[Optional[str]]:
        """
        Url of the SourceControl Repository.
        """
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

