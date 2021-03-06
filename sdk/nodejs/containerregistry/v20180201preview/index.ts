// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./buildStep";
export * from "./buildTask";
export * from "./getBuildLogLink";
export * from "./getBuildStep";
export * from "./getBuildTask";
export * from "./getRegistryBuildSourceUploadUrl";
export * from "./listBuildStepBuildArguments";
export * from "./listBuildTaskSourceRepositoryProperties";

// Export enums:
export * from "../../types/enums/containerregistry/v20180201preview";

// Import resources to register:
import { BuildStep } from "./buildStep";
import { BuildTask } from "./buildTask";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:containerregistry/v20180201preview:BuildStep":
                return new BuildStep(name, <any>undefined, { urn })
            case "azure-native:containerregistry/v20180201preview:BuildTask":
                return new BuildTask(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "containerregistry/v20180201preview", _module)
