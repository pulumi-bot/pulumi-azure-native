# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .database import *
from .elastic_pool import *
from .get_database import *
from .get_elastic_pool import *
from .get_instance_failover_group import *
from .get_managed_database_vulnerability_assessment import *
from .get_managed_database_vulnerability_assessment_rule_baseline import *
from .get_managed_instance_key import *
from .instance_failover_group import *
from .managed_database_vulnerability_assessment import *
from .managed_database_vulnerability_assessment_rule_baseline import *
from .managed_instance_key import *
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
            if typ == "azure-native:sql/v20171001preview:Database":
                return Database(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:sql/v20171001preview:ElasticPool":
                return ElasticPool(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:sql/v20171001preview:InstanceFailoverGroup":
                return InstanceFailoverGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:sql/v20171001preview:ManagedDatabaseVulnerabilityAssessment":
                return ManagedDatabaseVulnerabilityAssessment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:sql/v20171001preview:ManagedDatabaseVulnerabilityAssessmentRuleBaseline":
                return ManagedDatabaseVulnerabilityAssessmentRuleBaseline(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:sql/v20171001preview:ManagedInstanceKey":
                return ManagedInstanceKey(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "sql/v20171001preview", _module_instance)

_register_module()
