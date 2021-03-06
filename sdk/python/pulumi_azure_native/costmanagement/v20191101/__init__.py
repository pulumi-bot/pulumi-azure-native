# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .export import *
from .get_export import *
from .get_view import *
from .get_view_by_scope import *
from .view import *
from .view_by_scope import *
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
            if typ == "azure-native:costmanagement/v20191101:Export":
                return Export(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:costmanagement/v20191101:View":
                return View(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure-native:costmanagement/v20191101:ViewByScope":
                return ViewByScope(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure-native", "costmanagement/v20191101", _module_instance)

_register_module()
