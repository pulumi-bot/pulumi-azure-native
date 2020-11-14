// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getMigrateProject(args: GetMigrateProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetMigrateProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:migrate/v20180901preview:getMigrateProject", {
        "migrateProjectName": args.migrateProjectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMigrateProjectArgs {
    /**
     * Name of the Azure Migrate project.
     */
    readonly migrateProjectName: string;
    /**
     * Name of the Azure Resource Group that migrate project is part of.
     */
    readonly resourceGroupName: string;
}

/**
 * Migrate Project REST Resource.
 */
export interface GetMigrateProjectResult {
    /**
     * Gets or sets the eTag for concurrency control.
     */
    readonly eTag?: string;
    /**
     * Gets or sets the Azure location in which migrate project is created.
     */
    readonly location?: string;
    /**
     * Gets the name of the migrate project.
     */
    readonly name: string;
    /**
     * Gets or sets the nested properties.
     */
    readonly properties: outputs.migrate.v20180901preview.MigrateProjectPropertiesResponse;
    /**
     * Gets or sets the tags.
     */
    readonly tags?: outputs.migrate.v20180901preview.MigrateProjectResponseTags;
    /**
     * Handled by resource provider. Type = Microsoft.Migrate/MigrateProject.
     */
    readonly type: string;
}
