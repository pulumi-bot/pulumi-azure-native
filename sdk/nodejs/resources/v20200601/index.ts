// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./deployment";
export * from "./deploymentAtManagementGroupScope";
export * from "./deploymentAtScope";
export * from "./deploymentAtSubscriptionScope";
export * from "./deploymentAtTenantScope";
export * from "./getDeployment";
export * from "./getDeploymentAtManagementGroupScope";
export * from "./getDeploymentAtScope";
export * from "./getDeploymentAtSubscriptionScope";
export * from "./getDeploymentAtTenantScope";
export * from "./getResource";
export * from "./getResourceGroup";
export * from "./getTagAtScope";
export * from "./resource";
export * from "./resourceGroup";
export * from "./tagAtScope";

// Export enums:
export * from "../../types/enums/resources/v20200601";

// Import resources to register:
import { Deployment } from "./deployment";
import { DeploymentAtManagementGroupScope } from "./deploymentAtManagementGroupScope";
import { DeploymentAtScope } from "./deploymentAtScope";
import { DeploymentAtSubscriptionScope } from "./deploymentAtSubscriptionScope";
import { DeploymentAtTenantScope } from "./deploymentAtTenantScope";
import { Resource } from "./resource";
import { ResourceGroup } from "./resourceGroup";
import { TagAtScope } from "./tagAtScope";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-nextgen:resources/v20200601:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:DeploymentAtManagementGroupScope":
                return new DeploymentAtManagementGroupScope(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:DeploymentAtScope":
                return new DeploymentAtScope(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:DeploymentAtSubscriptionScope":
                return new DeploymentAtSubscriptionScope(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:DeploymentAtTenantScope":
                return new DeploymentAtTenantScope(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:Resource":
                return new Resource(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:ResourceGroup":
                return new ResourceGroup(name, <any>undefined, { urn })
            case "azure-nextgen:resources/v20200601:TagAtScope":
                return new TagAtScope(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-nextgen", "resources/v20200601", _module)
