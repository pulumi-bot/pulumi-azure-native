# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .commitment_plan import *
from .get_commitment_plan import *
from .get_web_service import *
from .get_workspace import *
from .list_workspace_keys import *
from .web_service import *
from .workspace import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    latest,
    v20160401,
    v20160501preview,
    v20170101,
    v20191001,
)

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure-native:machinelearning:CommitmentPlan":
                return CommitmentPlan(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:machinelearning:WebService":
                return WebService(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:machinelearning:Workspace":
                return Workspace(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "machinelearning", _module_instance)

_register_module()
