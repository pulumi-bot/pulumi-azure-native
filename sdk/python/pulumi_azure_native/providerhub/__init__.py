# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .default_rollout import *
from .get_default_rollout import *
from .get_notification_registration import *
from .get_provider_registration import *
from .get_resource_type_registration import *
from .get_skus import *
from .get_skus_nested_resource_type_first import *
from .get_skus_nested_resource_type_second import *
from .get_skus_nested_resource_type_third import *
from .notification_registration import *
from .operation_by_provider_registration import *
from .provider_registration import *
from .resource_type_registration import *
from .skus import *
from .skus_nested_resource_type_first import *
from .skus_nested_resource_type_second import *
from .skus_nested_resource_type_third import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20201120,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:providerhub:DefaultRollout":
                return DefaultRollout(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:NotificationRegistration":
                return NotificationRegistration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:OperationByProviderRegistration":
                return OperationByProviderRegistration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:ProviderRegistration":
                return ProviderRegistration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:ResourceTypeRegistration":
                return ResourceTypeRegistration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:Skus":
                return Skus(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:SkusNestedResourceTypeFirst":
                return SkusNestedResourceTypeFirst(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:SkusNestedResourceTypeSecond":
                return SkusNestedResourceTypeSecond(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:providerhub:SkusNestedResourceTypeThird":
                return SkusNestedResourceTypeThird(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "providerhub", _module_instance)

_register_module()
