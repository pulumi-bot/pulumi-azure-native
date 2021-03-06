# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .get_management_lock import *
from .get_management_lock_at_resource_group_level import *
from .management_lock import *
from .management_lock_at_resource_group_level import *

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:authorization/v20150101:ManagementLock":
                return ManagementLock(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:authorization/v20150101:ManagementLockAtResourceGroupLevel":
                return ManagementLockAtResourceGroupLevel(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "authorization/v20150101", _module_instance)

_register_module()
