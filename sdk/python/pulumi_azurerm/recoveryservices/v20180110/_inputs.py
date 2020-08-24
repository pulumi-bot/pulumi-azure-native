# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AddRecoveryServicesProviderInputPropertiesArgs',
    'AddVCenterRequestPropertiesArgs',
    'CreateNetworkMappingInputPropertiesArgs',
    'CreatePolicyInputPropertiesArgs',
    'CreateProtectionContainerMappingInputPropertiesArgs',
    'CreateRecoveryPlanInputPropertiesArgs',
    'EnableMigrationInputPropertiesArgs',
    'EnableMigrationProviderSpecificInputArgs',
    'EnableProtectionInputPropertiesArgs',
    'EnableProtectionProviderSpecificInputArgs',
    'FabricCreationInputPropertiesArgs',
    'FabricSpecificCreateNetworkMappingInputArgs',
    'FabricSpecificCreationInputArgs',
    'IdentityProviderInputArgs',
    'PolicyProviderSpecificInputArgs',
    'RecoveryPlanActionArgs',
    'RecoveryPlanGroupArgs',
    'RecoveryPlanProtectedItemArgs',
    'ReplicationProviderSpecificContainerMappingInputArgs',
    'StorageMappingInputPropertiesArgs',
]

@pulumi.input_type
class AddRecoveryServicesProviderInputPropertiesArgs:
    def __init__(__self__, *,
                 authentication_identity_input: pulumi.Input['IdentityProviderInputArgs'],
                 machine_name: pulumi.Input[str],
                 resource_access_identity_input: pulumi.Input['IdentityProviderInputArgs']):
        """
        The properties of an add provider request.
        :param pulumi.Input['IdentityProviderInputArgs'] authentication_identity_input: The identity provider input for DRA authentication.
        :param pulumi.Input[str] machine_name: The name of the machine where the provider is getting added.
        :param pulumi.Input['IdentityProviderInputArgs'] resource_access_identity_input: The identity provider input for resource access.
        """
        pulumi.set(__self__, "authentication_identity_input", authentication_identity_input)
        pulumi.set(__self__, "machine_name", machine_name)
        pulumi.set(__self__, "resource_access_identity_input", resource_access_identity_input)

    @property
    @pulumi.getter(name="authenticationIdentityInput")
    def authentication_identity_input(self) -> pulumi.Input['IdentityProviderInputArgs']:
        """
        The identity provider input for DRA authentication.
        """
        return pulumi.get(self, "authentication_identity_input")

    @authentication_identity_input.setter
    def authentication_identity_input(self, value: pulumi.Input['IdentityProviderInputArgs']):
        pulumi.set(self, "authentication_identity_input", value)

    @property
    @pulumi.getter(name="machineName")
    def machine_name(self) -> pulumi.Input[str]:
        """
        The name of the machine where the provider is getting added.
        """
        return pulumi.get(self, "machine_name")

    @machine_name.setter
    def machine_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "machine_name", value)

    @property
    @pulumi.getter(name="resourceAccessIdentityInput")
    def resource_access_identity_input(self) -> pulumi.Input['IdentityProviderInputArgs']:
        """
        The identity provider input for resource access.
        """
        return pulumi.get(self, "resource_access_identity_input")

    @resource_access_identity_input.setter
    def resource_access_identity_input(self, value: pulumi.Input['IdentityProviderInputArgs']):
        pulumi.set(self, "resource_access_identity_input", value)


@pulumi.input_type
class AddVCenterRequestPropertiesArgs:
    def __init__(__self__, *,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 process_server_id: Optional[pulumi.Input[str]] = None,
                 run_as_account_id: Optional[pulumi.Input[str]] = None):
        """
        The properties of an add vCenter request.
        :param pulumi.Input[str] friendly_name: The friendly name of the vCenter.
        :param pulumi.Input[str] ip_address: The IP address of the vCenter to be discovered.
        :param pulumi.Input[str] port: The port number for discovery.
        :param pulumi.Input[str] process_server_id: The process server Id from where the discovery is orchestrated.
        :param pulumi.Input[str] run_as_account_id: The account Id which has privileges to discover the vCenter.
        """
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if process_server_id is not None:
            pulumi.set(__self__, "process_server_id", process_server_id)
        if run_as_account_id is not None:
            pulumi.set(__self__, "run_as_account_id", run_as_account_id)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name of the vCenter.
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the vCenter to be discovered.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The port number for discovery.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="processServerId")
    def process_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The process server Id from where the discovery is orchestrated.
        """
        return pulumi.get(self, "process_server_id")

    @process_server_id.setter
    def process_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "process_server_id", value)

    @property
    @pulumi.getter(name="runAsAccountId")
    def run_as_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account Id which has privileges to discover the vCenter.
        """
        return pulumi.get(self, "run_as_account_id")

    @run_as_account_id.setter
    def run_as_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "run_as_account_id", value)


@pulumi.input_type
class CreateNetworkMappingInputPropertiesArgs:
    def __init__(__self__, *,
                 fabric_specific_details: Optional[pulumi.Input['FabricSpecificCreateNetworkMappingInputArgs']] = None,
                 recovery_fabric_name: Optional[pulumi.Input[str]] = None,
                 recovery_network_id: Optional[pulumi.Input[str]] = None):
        """
        Common input details for network mapping operation.
        :param pulumi.Input['FabricSpecificCreateNetworkMappingInputArgs'] fabric_specific_details: Fabric specific input properties.
        :param pulumi.Input[str] recovery_fabric_name: Recovery fabric Name.
        :param pulumi.Input[str] recovery_network_id: Recovery network Id.
        """
        if fabric_specific_details is not None:
            pulumi.set(__self__, "fabric_specific_details", fabric_specific_details)
        if recovery_fabric_name is not None:
            pulumi.set(__self__, "recovery_fabric_name", recovery_fabric_name)
        if recovery_network_id is not None:
            pulumi.set(__self__, "recovery_network_id", recovery_network_id)

    @property
    @pulumi.getter(name="fabricSpecificDetails")
    def fabric_specific_details(self) -> Optional[pulumi.Input['FabricSpecificCreateNetworkMappingInputArgs']]:
        """
        Fabric specific input properties.
        """
        return pulumi.get(self, "fabric_specific_details")

    @fabric_specific_details.setter
    def fabric_specific_details(self, value: Optional[pulumi.Input['FabricSpecificCreateNetworkMappingInputArgs']]):
        pulumi.set(self, "fabric_specific_details", value)

    @property
    @pulumi.getter(name="recoveryFabricName")
    def recovery_fabric_name(self) -> Optional[pulumi.Input[str]]:
        """
        Recovery fabric Name.
        """
        return pulumi.get(self, "recovery_fabric_name")

    @recovery_fabric_name.setter
    def recovery_fabric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recovery_fabric_name", value)

    @property
    @pulumi.getter(name="recoveryNetworkId")
    def recovery_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        Recovery network Id.
        """
        return pulumi.get(self, "recovery_network_id")

    @recovery_network_id.setter
    def recovery_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recovery_network_id", value)


@pulumi.input_type
class CreatePolicyInputPropertiesArgs:
    def __init__(__self__, *,
                 provider_specific_input: Optional[pulumi.Input['PolicyProviderSpecificInputArgs']] = None):
        """
        Policy creation properties.
        :param pulumi.Input['PolicyProviderSpecificInputArgs'] provider_specific_input: The ReplicationProviderSettings.
        """
        if provider_specific_input is not None:
            pulumi.set(__self__, "provider_specific_input", provider_specific_input)

    @property
    @pulumi.getter(name="providerSpecificInput")
    def provider_specific_input(self) -> Optional[pulumi.Input['PolicyProviderSpecificInputArgs']]:
        """
        The ReplicationProviderSettings.
        """
        return pulumi.get(self, "provider_specific_input")

    @provider_specific_input.setter
    def provider_specific_input(self, value: Optional[pulumi.Input['PolicyProviderSpecificInputArgs']]):
        pulumi.set(self, "provider_specific_input", value)


@pulumi.input_type
class CreateProtectionContainerMappingInputPropertiesArgs:
    def __init__(__self__, *,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 provider_specific_input: Optional[pulumi.Input['ReplicationProviderSpecificContainerMappingInputArgs']] = None,
                 target_protection_container_id: Optional[pulumi.Input[str]] = None):
        """
        Configure pairing input properties.
        :param pulumi.Input[str] policy_id: Applicable policy.
        :param pulumi.Input['ReplicationProviderSpecificContainerMappingInputArgs'] provider_specific_input: Provider specific input for pairing.
        :param pulumi.Input[str] target_protection_container_id: The target unique protection container name.
        """
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if provider_specific_input is not None:
            pulumi.set(__self__, "provider_specific_input", provider_specific_input)
        if target_protection_container_id is not None:
            pulumi.set(__self__, "target_protection_container_id", target_protection_container_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Applicable policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="providerSpecificInput")
    def provider_specific_input(self) -> Optional[pulumi.Input['ReplicationProviderSpecificContainerMappingInputArgs']]:
        """
        Provider specific input for pairing.
        """
        return pulumi.get(self, "provider_specific_input")

    @provider_specific_input.setter
    def provider_specific_input(self, value: Optional[pulumi.Input['ReplicationProviderSpecificContainerMappingInputArgs']]):
        pulumi.set(self, "provider_specific_input", value)

    @property
    @pulumi.getter(name="targetProtectionContainerId")
    def target_protection_container_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target unique protection container name.
        """
        return pulumi.get(self, "target_protection_container_id")

    @target_protection_container_id.setter
    def target_protection_container_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_protection_container_id", value)


@pulumi.input_type
class CreateRecoveryPlanInputPropertiesArgs:
    def __init__(__self__, *,
                 groups: pulumi.Input[List[pulumi.Input['RecoveryPlanGroupArgs']]],
                 primary_fabric_id: pulumi.Input[str],
                 recovery_fabric_id: pulumi.Input[str],
                 failover_deployment_model: Optional[pulumi.Input[str]] = None):
        """
        Recovery plan creation properties.
        :param pulumi.Input[List[pulumi.Input['RecoveryPlanGroupArgs']]] groups: The recovery plan groups.
        :param pulumi.Input[str] primary_fabric_id: The primary fabric Id.
        :param pulumi.Input[str] recovery_fabric_id: The recovery fabric Id.
        :param pulumi.Input[str] failover_deployment_model: The failover deployment model.
        """
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "primary_fabric_id", primary_fabric_id)
        pulumi.set(__self__, "recovery_fabric_id", recovery_fabric_id)
        if failover_deployment_model is not None:
            pulumi.set(__self__, "failover_deployment_model", failover_deployment_model)

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Input[List[pulumi.Input['RecoveryPlanGroupArgs']]]:
        """
        The recovery plan groups.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: pulumi.Input[List[pulumi.Input['RecoveryPlanGroupArgs']]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="primaryFabricId")
    def primary_fabric_id(self) -> pulumi.Input[str]:
        """
        The primary fabric Id.
        """
        return pulumi.get(self, "primary_fabric_id")

    @primary_fabric_id.setter
    def primary_fabric_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_fabric_id", value)

    @property
    @pulumi.getter(name="recoveryFabricId")
    def recovery_fabric_id(self) -> pulumi.Input[str]:
        """
        The recovery fabric Id.
        """
        return pulumi.get(self, "recovery_fabric_id")

    @recovery_fabric_id.setter
    def recovery_fabric_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "recovery_fabric_id", value)

    @property
    @pulumi.getter(name="failoverDeploymentModel")
    def failover_deployment_model(self) -> Optional[pulumi.Input[str]]:
        """
        The failover deployment model.
        """
        return pulumi.get(self, "failover_deployment_model")

    @failover_deployment_model.setter
    def failover_deployment_model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "failover_deployment_model", value)


@pulumi.input_type
class EnableMigrationInputPropertiesArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 provider_specific_details: pulumi.Input['EnableMigrationProviderSpecificInputArgs']):
        """
        Enable migration input properties.
        :param pulumi.Input[str] policy_id: The policy Id.
        :param pulumi.Input['EnableMigrationProviderSpecificInputArgs'] provider_specific_details: The provider specific details.
        """
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "provider_specific_details", provider_specific_details)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The policy Id.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="providerSpecificDetails")
    def provider_specific_details(self) -> pulumi.Input['EnableMigrationProviderSpecificInputArgs']:
        """
        The provider specific details.
        """
        return pulumi.get(self, "provider_specific_details")

    @provider_specific_details.setter
    def provider_specific_details(self, value: pulumi.Input['EnableMigrationProviderSpecificInputArgs']):
        pulumi.set(self, "provider_specific_details", value)


@pulumi.input_type
class EnableMigrationProviderSpecificInputArgs:
    def __init__(__self__, *,
                 instance_type: pulumi.Input[str]):
        """
        Enable migration provider specific input.
        :param pulumi.Input[str] instance_type: The class type.
        """
        pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        The class type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class EnableProtectionInputPropertiesArgs:
    def __init__(__self__, *,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protectable_item_id: Optional[pulumi.Input[str]] = None,
                 provider_specific_details: Optional[pulumi.Input['EnableProtectionProviderSpecificInputArgs']] = None):
        """
        Enable protection input properties.
        :param pulumi.Input[str] policy_id: The Policy Id.
        :param pulumi.Input[str] protectable_item_id: The protectable item Id.
        :param pulumi.Input['EnableProtectionProviderSpecificInputArgs'] provider_specific_details: The ReplicationProviderInput. For HyperVReplicaAzure provider, it will be AzureEnableProtectionInput object. For San provider, it will be SanEnableProtectionInput object. For HyperVReplicaAzure provider, it can be null.
        """
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protectable_item_id is not None:
            pulumi.set(__self__, "protectable_item_id", protectable_item_id)
        if provider_specific_details is not None:
            pulumi.set(__self__, "provider_specific_details", provider_specific_details)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Policy Id.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="protectableItemId")
    def protectable_item_id(self) -> Optional[pulumi.Input[str]]:
        """
        The protectable item Id.
        """
        return pulumi.get(self, "protectable_item_id")

    @protectable_item_id.setter
    def protectable_item_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protectable_item_id", value)

    @property
    @pulumi.getter(name="providerSpecificDetails")
    def provider_specific_details(self) -> Optional[pulumi.Input['EnableProtectionProviderSpecificInputArgs']]:
        """
        The ReplicationProviderInput. For HyperVReplicaAzure provider, it will be AzureEnableProtectionInput object. For San provider, it will be SanEnableProtectionInput object. For HyperVReplicaAzure provider, it can be null.
        """
        return pulumi.get(self, "provider_specific_details")

    @provider_specific_details.setter
    def provider_specific_details(self, value: Optional[pulumi.Input['EnableProtectionProviderSpecificInputArgs']]):
        pulumi.set(self, "provider_specific_details", value)


@pulumi.input_type
class EnableProtectionProviderSpecificInputArgs:
    def __init__(__self__, *,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Enable protection provider specific input.
        :param pulumi.Input[str] instance_type: The class type.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The class type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class FabricCreationInputPropertiesArgs:
    def __init__(__self__, *,
                 custom_details: Optional[pulumi.Input['FabricSpecificCreationInputArgs']] = None):
        """
        Properties of site details provided during the time of site creation
        :param pulumi.Input['FabricSpecificCreationInputArgs'] custom_details: Fabric provider specific creation input.
        """
        if custom_details is not None:
            pulumi.set(__self__, "custom_details", custom_details)

    @property
    @pulumi.getter(name="customDetails")
    def custom_details(self) -> Optional[pulumi.Input['FabricSpecificCreationInputArgs']]:
        """
        Fabric provider specific creation input.
        """
        return pulumi.get(self, "custom_details")

    @custom_details.setter
    def custom_details(self, value: Optional[pulumi.Input['FabricSpecificCreationInputArgs']]):
        pulumi.set(self, "custom_details", value)


@pulumi.input_type
class FabricSpecificCreateNetworkMappingInputArgs:
    def __init__(__self__, *,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Input details specific to fabrics during Network Mapping.
        :param pulumi.Input[str] instance_type: The instance type.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The instance type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class FabricSpecificCreationInputArgs:
    def __init__(__self__, *,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Fabric provider specific settings.
        :param pulumi.Input[str] instance_type: Gets the class type.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        Gets the class type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class IdentityProviderInputArgs:
    def __init__(__self__, *,
                 aad_authority: pulumi.Input[str],
                 application_id: pulumi.Input[str],
                 audience: pulumi.Input[str],
                 object_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str]):
        """
        Identity provider input.
        :param pulumi.Input[str] aad_authority: The base authority for Azure Active Directory authentication.
        :param pulumi.Input[str] application_id: The application/client Id for the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        :param pulumi.Input[str] audience: The intended Audience of the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        :param pulumi.Input[str] object_id: The object Id of the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        :param pulumi.Input[str] tenant_id: The tenant Id for the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        """
        pulumi.set(__self__, "aad_authority", aad_authority)
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "audience", audience)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="aadAuthority")
    def aad_authority(self) -> pulumi.Input[str]:
        """
        The base authority for Azure Active Directory authentication.
        """
        return pulumi.get(self, "aad_authority")

    @aad_authority.setter
    def aad_authority(self, value: pulumi.Input[str]):
        pulumi.set(self, "aad_authority", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The application/client Id for the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def audience(self) -> pulumi.Input[str]:
        """
        The intended Audience of the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: pulumi.Input[str]):
        pulumi.set(self, "audience", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[str]:
        """
        The object Id of the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The tenant Id for the service principal with which the on-premise management/data plane components would communicate with our Azure services.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class PolicyProviderSpecificInputArgs:
    def __init__(__self__, *,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Base class for provider specific input
        :param pulumi.Input[str] instance_type: The class type.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The class type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class RecoveryPlanActionArgs:
    def __init__(__self__, *,
                 action_name: pulumi.Input[str],
                 failover_directions: pulumi.Input[List[pulumi.Input[str]]],
                 failover_types: pulumi.Input[List[pulumi.Input[str]]]):
        """
        Recovery plan action details.
        :param pulumi.Input[str] action_name: The action name.
        :param pulumi.Input[List[pulumi.Input[str]]] failover_directions: The list of failover directions.
        :param pulumi.Input[List[pulumi.Input[str]]] failover_types: The list of failover types.
        """
        pulumi.set(__self__, "action_name", action_name)
        pulumi.set(__self__, "failover_directions", failover_directions)
        pulumi.set(__self__, "failover_types", failover_types)

    @property
    @pulumi.getter(name="actionName")
    def action_name(self) -> pulumi.Input[str]:
        """
        The action name.
        """
        return pulumi.get(self, "action_name")

    @action_name.setter
    def action_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "action_name", value)

    @property
    @pulumi.getter(name="failoverDirections")
    def failover_directions(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The list of failover directions.
        """
        return pulumi.get(self, "failover_directions")

    @failover_directions.setter
    def failover_directions(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "failover_directions", value)

    @property
    @pulumi.getter(name="failoverTypes")
    def failover_types(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The list of failover types.
        """
        return pulumi.get(self, "failover_types")

    @failover_types.setter
    def failover_types(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "failover_types", value)


@pulumi.input_type
class RecoveryPlanGroupArgs:
    def __init__(__self__, *,
                 group_type: pulumi.Input[str],
                 end_group_actions: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]] = None,
                 replication_protected_items: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanProtectedItemArgs']]]] = None,
                 start_group_actions: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]] = None):
        """
        Recovery plan group details.
        :param pulumi.Input[str] group_type: The group type.
        :param pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]] end_group_actions: The end group actions.
        :param pulumi.Input[List[pulumi.Input['RecoveryPlanProtectedItemArgs']]] replication_protected_items: The list of protected items.
        :param pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]] start_group_actions: The start group actions.
        """
        pulumi.set(__self__, "group_type", group_type)
        if end_group_actions is not None:
            pulumi.set(__self__, "end_group_actions", end_group_actions)
        if replication_protected_items is not None:
            pulumi.set(__self__, "replication_protected_items", replication_protected_items)
        if start_group_actions is not None:
            pulumi.set(__self__, "start_group_actions", start_group_actions)

    @property
    @pulumi.getter(name="groupType")
    def group_type(self) -> pulumi.Input[str]:
        """
        The group type.
        """
        return pulumi.get(self, "group_type")

    @group_type.setter
    def group_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_type", value)

    @property
    @pulumi.getter(name="endGroupActions")
    def end_group_actions(self) -> Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]]:
        """
        The end group actions.
        """
        return pulumi.get(self, "end_group_actions")

    @end_group_actions.setter
    def end_group_actions(self, value: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]]):
        pulumi.set(self, "end_group_actions", value)

    @property
    @pulumi.getter(name="replicationProtectedItems")
    def replication_protected_items(self) -> Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanProtectedItemArgs']]]]:
        """
        The list of protected items.
        """
        return pulumi.get(self, "replication_protected_items")

    @replication_protected_items.setter
    def replication_protected_items(self, value: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanProtectedItemArgs']]]]):
        pulumi.set(self, "replication_protected_items", value)

    @property
    @pulumi.getter(name="startGroupActions")
    def start_group_actions(self) -> Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]]:
        """
        The start group actions.
        """
        return pulumi.get(self, "start_group_actions")

    @start_group_actions.setter
    def start_group_actions(self, value: Optional[pulumi.Input[List[pulumi.Input['RecoveryPlanActionArgs']]]]):
        pulumi.set(self, "start_group_actions", value)


@pulumi.input_type
class RecoveryPlanProtectedItemArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None):
        """
        Recovery plan protected item.
        :param pulumi.Input[str] id: The ARM Id of the recovery plan protected item.
        :param pulumi.Input[str] virtual_machine_id: The virtual machine Id.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ARM Id of the recovery plan protected item.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[str]]:
        """
        The virtual machine Id.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_machine_id", value)


@pulumi.input_type
class ReplicationProviderSpecificContainerMappingInputArgs:
    def __init__(__self__, *,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Provider specific input for pairing operations.
        :param pulumi.Input[str] instance_type: The class type.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The class type.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class StorageMappingInputPropertiesArgs:
    def __init__(__self__, *,
                 target_storage_classification_id: Optional[pulumi.Input[str]] = None):
        """
        Storage mapping input properties.
        :param pulumi.Input[str] target_storage_classification_id: The ID of the storage object.
        """
        if target_storage_classification_id is not None:
            pulumi.set(__self__, "target_storage_classification_id", target_storage_classification_id)

    @property
    @pulumi.getter(name="targetStorageClassificationId")
    def target_storage_classification_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the storage object.
        """
        return pulumi.get(self, "target_storage_classification_id")

    @target_storage_classification_id.setter
    def target_storage_classification_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_storage_classification_id", value)


