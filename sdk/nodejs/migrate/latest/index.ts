// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./assessment";
export * from "./getAssessment";
export * from "./getGroup";
export * from "./getHyperVCollector";
export * from "./getMoveCollection";
export * from "./getMoveResource";
export * from "./getProject";
export * from "./getVMwareCollector";
export * from "./group";
export * from "./hyperVCollector";
export * from "./moveCollection";
export * from "./moveResource";
export * from "./project";
export * from "./vmwareCollector";

// Export enums:
export * from "../../types/enums/migrate/latest";

// Import resources to register:
import { Assessment } from "./assessment";
import { Group } from "./group";
import { HyperVCollector } from "./hyperVCollector";
import { MoveCollection } from "./moveCollection";
import { MoveResource } from "./moveResource";
import { Project } from "./project";
import { VMwareCollector } from "./vmwareCollector";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:migrate/latest:Assessment":
                return new Assessment(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:Group":
                return new Group(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:HyperVCollector":
                return new HyperVCollector(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:MoveCollection":
                return new MoveCollection(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:MoveResource":
                return new MoveResource(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:Project":
                return new Project(name, <any>undefined, { urn })
            case "azure-native:migrate/latest:VMwareCollector":
                return new VMwareCollector(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "migrate/latest", _module)
