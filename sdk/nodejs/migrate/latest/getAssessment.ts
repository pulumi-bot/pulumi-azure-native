// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * An assessment created for a group in the Migration project.
 * Latest API Version: 2019-10-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:migrate:getAssessment'. */
export function getAssessment(args: GetAssessmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentResult> {
    pulumi.log.warn("getAssessment is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:migrate:getAssessment'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:migrate/latest:getAssessment", {
        "assessmentName": args.assessmentName,
        "groupName": args.groupName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAssessmentArgs {
    /**
     * Unique name of an assessment within a project.
     */
    readonly assessmentName: string;
    /**
     * Unique name of a group within a project.
     */
    readonly groupName: string;
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: string;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: string;
}

/**
 * An assessment created for a group in the Migration project.
 */
export interface GetAssessmentResult {
    /**
     * For optimistic concurrency control.
     */
    readonly eTag?: string;
    /**
     * Path reference to this assessment. /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessment/{assessmentName}
     */
    readonly id: string;
    /**
     * Unique name of an assessment.
     */
    readonly name: string;
    /**
     * Properties of the assessment.
     */
    readonly properties: outputs.migrate.latest.AssessmentPropertiesResponse;
    /**
     * Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
     */
    readonly type: string;
}
