# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .database import *
from .database_vulnerability_assessment import *
from .database_vulnerability_assessment_rule_baseline import *
from .elastic_pool import *
from .failover_group import *
from .firewall_rule import *
from .get_database import *
from .get_database_vulnerability_assessment import *
from .get_database_vulnerability_assessment_rule_baseline import *
from .get_elastic_pool import *
from .get_failover_group import *
from .get_firewall_rule import *
from .get_instance_failover_group import *
from .get_instance_pool import *
from .get_job import *
from .get_job_agent import *
from .get_job_credential import *
from .get_job_step import *
from .get_job_target_group import *
from .get_managed_database import *
from .get_managed_database_sensitivity_label import *
from .get_managed_database_vulnerability_assessment import *
from .get_managed_database_vulnerability_assessment_rule_baseline import *
from .get_managed_instance import *
from .get_managed_instance_administrator import *
from .get_managed_instance_azure_ad_only_authentication import *
from .get_managed_instance_key import *
from .get_managed_instance_private_endpoint_connection import *
from .get_managed_instance_vulnerability_assessment import *
from .get_private_endpoint_connection import *
from .get_sensitivity_label import *
from .get_server import *
from .get_server_azure_ad_administrator import *
from .get_server_azure_ad_only_authentication import *
from .get_server_dns_alias import *
from .get_server_key import *
from .get_server_trust_group import *
from .get_server_vulnerability_assessment import *
from .get_sync_agent import *
from .get_sync_group import *
from .get_sync_member import *
from .get_transparent_data_encryption import *
from .get_virtual_network_rule import *
from .get_workload_classifier import *
from .get_workload_group import *
from .instance_failover_group import *
from .instance_pool import *
from .job import *
from .job_agent import *
from .job_credential import *
from .job_step import *
from .job_target_group import *
from .managed_database import *
from .managed_database_sensitivity_label import *
from .managed_database_vulnerability_assessment import *
from .managed_database_vulnerability_assessment_rule_baseline import *
from .managed_instance import *
from .managed_instance_administrator import *
from .managed_instance_azure_ad_only_authentication import *
from .managed_instance_key import *
from .managed_instance_private_endpoint_connection import *
from .managed_instance_vulnerability_assessment import *
from .private_endpoint_connection import *
from .sensitivity_label import *
from .server import *
from .server_azure_ad_administrator import *
from .server_azure_ad_only_authentication import *
from .server_dns_alias import *
from .server_key import *
from .server_trust_group import *
from .server_vulnerability_assessment import *
from .sync_agent import *
from .sync_group import *
from .sync_member import *
from .transparent_data_encryption import *
from .virtual_network_rule import *
from .workload_classifier import *
from .workload_group import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-nextgen:sql/v20200801preview:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:DatabaseVulnerabilityAssessment":
                return DatabaseVulnerabilityAssessment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:DatabaseVulnerabilityAssessmentRuleBaseline":
                return DatabaseVulnerabilityAssessmentRuleBaseline(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ElasticPool":
                return ElasticPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:FailoverGroup":
                return FailoverGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:FirewallRule":
                return FirewallRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:InstanceFailoverGroup":
                return InstanceFailoverGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:InstancePool":
                return InstancePool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:Job":
                return Job(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:JobAgent":
                return JobAgent(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:JobCredential":
                return JobCredential(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:JobStep":
                return JobStep(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:JobTargetGroup":
                return JobTargetGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedDatabase":
                return ManagedDatabase(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedDatabaseSensitivityLabel":
                return ManagedDatabaseSensitivityLabel(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedDatabaseVulnerabilityAssessment":
                return ManagedDatabaseVulnerabilityAssessment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
                return ManagedDatabaseVulnerabilityAssessmentRuleBaseline(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstance":
                return ManagedInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstanceAdministrator":
                return ManagedInstanceAdministrator(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstanceAzureADOnlyAuthentication":
                return ManagedInstanceAzureADOnlyAuthentication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstanceKey":
                return ManagedInstanceKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstancePrivateEndpointConnection":
                return ManagedInstancePrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ManagedInstanceVulnerabilityAssessment":
                return ManagedInstanceVulnerabilityAssessment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:PrivateEndpointConnection":
                return PrivateEndpointConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:SensitivityLabel":
                return SensitivityLabel(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:Server":
                return Server(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerAzureADAdministrator":
                return ServerAzureADAdministrator(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerAzureADOnlyAuthentication":
                return ServerAzureADOnlyAuthentication(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerDnsAlias":
                return ServerDnsAlias(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerKey":
                return ServerKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerTrustGroup":
                return ServerTrustGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:ServerVulnerabilityAssessment":
                return ServerVulnerabilityAssessment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:SyncAgent":
                return SyncAgent(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:SyncGroup":
                return SyncGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:SyncMember":
                return SyncMember(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:TransparentDataEncryption":
                return TransparentDataEncryption(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:VirtualNetworkRule":
                return VirtualNetworkRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:WorkloadClassifier":
                return WorkloadClassifier(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/v20200801preview:WorkloadGroup":
                return WorkloadGroup(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-nextgen", "sql/v20200801preview", _module_instance)

_register_module()
