// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./artifactSourceResource";
export * from "./customImageResource";
export * from "./formulaResource";
export * from "./getArtifactSourceResource";
export * from "./getCustomImageResource";
export * from "./getFormulaResource";
export * from "./getLabResource";
export * from "./getPolicyResource";
export * from "./getScheduleResource";
export * from "./getVirtualMachineResource";
export * from "./getVirtualNetworkResource";
export * from "./labResource";
export * from "./listLabVhds";
export * from "./policyResource";
export * from "./scheduleResource";
export * from "./virtualMachineResource";
export * from "./virtualNetworkResource";

// Export enums:
export * from "../../types/enums/devtestlab/v20150521preview";

// Import resources to register:
import { ArtifactSourceResource } from "./artifactSourceResource";
import { CustomImageResource } from "./customImageResource";
import { FormulaResource } from "./formulaResource";
import { LabResource } from "./labResource";
import { PolicyResource } from "./policyResource";
import { ScheduleResource } from "./scheduleResource";
import { VirtualMachineResource } from "./virtualMachineResource";
import { VirtualNetworkResource } from "./virtualNetworkResource";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:devtestlab/v20150521preview:ArtifactSourceResource":
                return new ArtifactSourceResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:CustomImageResource":
                return new CustomImageResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:FormulaResource":
                return new FormulaResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:LabResource":
                return new LabResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:PolicyResource":
                return new PolicyResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:ScheduleResource":
                return new ScheduleResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:VirtualMachineResource":
                return new VirtualMachineResource(name, <any>undefined, { urn })
            case "azure-native:devtestlab/v20150521preview:VirtualNetworkResource":
                return new VirtualNetworkResource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "devtestlab/v20150521preview", _module)
