// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getManagementGroup(args: GetManagementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:management/latest:getManagementGroup", {
        "expand": args.expand,
        "filter": args.filter,
        "groupId": args.groupId,
        "recurse": args.recurse,
    }, opts);
}

export interface GetManagementGroupArgs {
    /**
     * The $expand=children query string parameter allows clients to request inclusion of children in the response payload.  $expand=path includes the path from the root group to the current group.
     */
    readonly expand?: string;
    /**
     * A filter which allows the exclusion of subscriptions from results (i.e. '$filter=children.childType ne Subscription')
     */
    readonly filter?: string;
    /**
     * Management Group ID.
     */
    readonly groupId: string;
    /**
     * The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload. Note that  $expand=children must be passed up if $recurse is set to true.
     */
    readonly recurse?: boolean;
}

/**
 * The management group details.
 */
export interface GetManagementGroupResult {
    /**
     * The list of children.
     */
    readonly children?: outputs.management.latest.ManagementGroupChildInfoResponse[];
    /**
     * The details of a management group.
     */
    readonly details?: outputs.management.latest.ManagementGroupDetailsResponse;
    /**
     * The friendly name of the management group.
     */
    readonly displayName?: string;
    /**
     * The name of the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    readonly name: string;
    /**
     * The path from the root to the current group.
     */
    readonly path?: outputs.management.latest.ManagementGroupPathElementResponse[];
    /**
     * The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    readonly tenantId?: string;
    /**
     * The type of the resource.  For example, Microsoft.Management/managementGroups
     */
    readonly type: string;
}
