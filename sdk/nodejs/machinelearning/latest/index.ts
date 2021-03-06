// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getWebService";
export * from "./getWorkspace";
export * from "./listWorkspaceKeys";
export * from "./webService";
export * from "./workspace";

// Export enums:
export * from "../../types/enums/machinelearning/latest";

// Import resources to register:
import { WebService } from "./webService";
import { Workspace } from "./workspace";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:machinelearning/latest:WebService":
                return new WebService(name, <any>undefined, { urn })
            case "azure-native:machinelearning/latest:Workspace":
                return new Workspace(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "machinelearning/latest", _module)
