# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .database import *
from .database_threat_detection_policy import *
from .disaster_recovery_configuration import *
from .elastic_pool import *
from .firewall_rule import *
from .geo_backup_policy import *
from .get_database import *
from .get_database_threat_detection_policy import *
from .get_disaster_recovery_configuration import *
from .get_elastic_pool import *
from .get_firewall_rule import *
from .get_geo_backup_policy import *
from .get_server import *
from .get_server_azure_ad_administrator import *
from .get_server_communication_link import *
from .get_transparent_data_encryption import *
from .server import *
from .server_azure_ad_administrator import *
from .server_communication_link import *
from .transparent_data_encryption import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-nextgen:sql/latest:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:DatabaseThreatDetectionPolicy":
                return DatabaseThreatDetectionPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:DisasterRecoveryConfiguration":
                return DisasterRecoveryConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:ElasticPool":
                return ElasticPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:FirewallRule":
                return FirewallRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:GeoBackupPolicy":
                return GeoBackupPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:Server":
                return Server(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:ServerAzureADAdministrator":
                return ServerAzureADAdministrator(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:ServerCommunicationLink":
                return ServerCommunicationLink(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-nextgen:sql/latest:TransparentDataEncryption":
                return TransparentDataEncryption(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-nextgen", "sql/latest", _module_instance)

_register_module()
