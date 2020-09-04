// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getLiveOutput(args: GetLiveOutputArgs, opts?: pulumi.InvokeOptions): Promise<GetLiveOutputResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:media/v20190501preview:getLiveOutput", {
        "accountName": args.accountName,
        "liveEventName": args.liveEventName,
        "liveOutputName": args.liveOutputName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLiveOutputArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The name of the Live Event.
     */
    readonly liveEventName: string;
    /**
     * The name of the Live Output.
     */
    readonly liveOutputName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * The Live Output.
 */
export interface GetLiveOutputResult {
    /**
     * ISO 8601 timespan duration of the archive window length. This is duration that customer want to retain the recorded content.
     */
    readonly archiveWindowLength: string;
    /**
     * The asset name.
     */
    readonly assetName: string;
    /**
     * The exact time the Live Output was created.
     */
    readonly created: string;
    /**
     * The description of the Live Output.
     */
    readonly description?: string;
    /**
     * The HLS configuration.
     */
    readonly hls?: outputs.media.v20190501preview.HlsResponse;
    /**
     * The exact time the Live Output was last modified.
     */
    readonly lastModified: string;
    /**
     * The manifest file name.  If not provided, the service will generate one automatically.
     */
    readonly manifestName?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The output snapshot time.
     */
    readonly outputSnapTime?: number;
    /**
     * The provisioning state of the Live Output.
     */
    readonly provisioningState: string;
    /**
     * The resource state of the Live Output.
     */
    readonly resourceState: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
