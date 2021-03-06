# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .get_registration_assignment import *
from .get_registration_definition import *
from .registration_assignment import *
from .registration_definition import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20180601preview,
    v20190401preview,
    v20190601,
    v20190901,
    v20200201preview,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:managedservices:RegistrationAssignment":
                return RegistrationAssignment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:managedservices:RegistrationDefinition":
                return RegistrationDefinition(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "managedservices", _module_instance)

_register_module()
