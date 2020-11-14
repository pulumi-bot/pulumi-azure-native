// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getLiveEvent(args: GetLiveEventArgs, opts?: pulumi.InvokeOptions): Promise<GetLiveEventResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:media/v20180701:getLiveEvent", {
        "accountName": args.accountName,
        "liveEventName": args.liveEventName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLiveEventArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The name of the Live Event.
     */
    readonly liveEventName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * The Live Event.
 */
export interface GetLiveEventResult {
    /**
     * The exact time the Live Event was created.
     */
    readonly created: string;
    /**
     * The Live Event access policies.
     */
    readonly crossSiteAccessPolicies?: outputs.media.v20180701.CrossSiteAccessPoliciesResponse;
    /**
     * The Live Event description.
     */
    readonly description?: string;
    /**
     * The Live Event encoding.
     */
    readonly encoding?: outputs.media.v20180701.LiveEventEncodingResponse;
    /**
     * The Live Event input.
     */
    readonly input: outputs.media.v20180701.LiveEventInputResponse;
    /**
     * The exact time the Live Event was last modified.
     */
    readonly lastModified: string;
    /**
     * The Azure Region of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The Live Event preview.
     */
    readonly preview?: outputs.media.v20180701.LiveEventPreviewResponse;
    /**
     * The provisioning state of the Live Event.
     */
    readonly provisioningState: string;
    /**
     * The resource state of the Live Event.
     */
    readonly resourceState: string;
    /**
     * The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.
     */
    readonly streamOptions?: string[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
     */
    readonly vanityUrl?: boolean;
}
