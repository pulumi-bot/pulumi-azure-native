// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Job Resource.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:databox/v20200401:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Reason for cancellation.
     */
    public /*out*/ readonly cancellationReason!: pulumi.Output<string>;
    /**
     * Delivery Info of Job.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.databox.v20200401.JobDeliveryInfoResponse | undefined>;
    /**
     * Delivery type of Job.
     */
    public readonly deliveryType!: pulumi.Output<string | undefined>;
    /**
     * Details of a job run. This field will only be sent for expand details filter.
     */
    public readonly details!: pulumi.Output<outputs.databox.v20200401.DataBoxDiskJobDetailsResponse | outputs.databox.v20200401.DataBoxHeavyJobDetailsResponse | outputs.databox.v20200401.DataBoxJobDetailsResponse | undefined>;
    /**
     * Top level error for the job.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.databox.v20200401.CloudErrorResponse>;
    /**
     * Msi identity of the resource
     */
    public readonly identity!: pulumi.Output<outputs.databox.v20200401.ResourceIdentityResponse | undefined>;
    /**
     * Describes whether the job is cancellable or not.
     */
    public /*out*/ readonly isCancellable!: pulumi.Output<boolean>;
    /**
     * Flag to indicate cancellation of scheduled job.
     */
    public /*out*/ readonly isCancellableWithoutFee!: pulumi.Output<boolean>;
    /**
     * Describes whether the job is deletable or not.
     */
    public /*out*/ readonly isDeletable!: pulumi.Output<boolean>;
    /**
     * Is Prepare To Ship Enabled on this job
     */
    public /*out*/ readonly isPrepareToShipEnabled!: pulumi.Output<boolean>;
    /**
     * Describes whether the shipping address is editable or not.
     */
    public /*out*/ readonly isShippingAddressEditable!: pulumi.Output<boolean>;
    /**
     * The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The sku type.
     */
    public readonly sku!: pulumi.Output<outputs.databox.v20200401.SkuResponse>;
    /**
     * Time at which the job was started in UTC ISO 8601 format.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * Name of the stage which is in progress.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of the data transfer.
     */
    public readonly transferType!: pulumi.Output<string>;
    /**
     * Type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            if (!args || args.transferType === undefined) {
                throw new Error("Missing required property 'transferType'");
            }
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["deliveryType"] = args ? args.deliveryType : undefined;
            inputs["details"] = args ? args.details : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["jobName"] = args ? args.jobName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transferType"] = args ? args.transferType : undefined;
            inputs["cancellationReason"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["isCancellable"] = undefined /*out*/;
            inputs["isCancellableWithoutFee"] = undefined /*out*/;
            inputs["isDeletable"] = undefined /*out*/;
            inputs["isPrepareToShipEnabled"] = undefined /*out*/;
            inputs["isShippingAddressEditable"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["cancellationReason"] = undefined /*out*/;
            inputs["deliveryInfo"] = undefined /*out*/;
            inputs["deliveryType"] = undefined /*out*/;
            inputs["details"] = undefined /*out*/;
            inputs["error"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["isCancellable"] = undefined /*out*/;
            inputs["isCancellableWithoutFee"] = undefined /*out*/;
            inputs["isDeletable"] = undefined /*out*/;
            inputs["isPrepareToShipEnabled"] = undefined /*out*/;
            inputs["isShippingAddressEditable"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["transferType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:databox/latest:Job" }, { type: "azure-nextgen:databox/v20180101:Job" }, { type: "azure-nextgen:databox/v20190901:Job" }, { type: "azure-nextgen:databox/v20201101:Job" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Delivery Info of Job.
     */
    readonly deliveryInfo?: pulumi.Input<inputs.databox.v20200401.JobDeliveryInfo>;
    /**
     * Delivery type of Job.
     */
    readonly deliveryType?: pulumi.Input<string>;
    /**
     * Details of a job run. This field will only be sent for expand details filter.
     */
    readonly details?: pulumi.Input<inputs.databox.v20200401.DataBoxDiskJobDetails | inputs.databox.v20200401.DataBoxHeavyJobDetails | inputs.databox.v20200401.DataBoxJobDetails>;
    /**
     * Msi identity of the resource
     */
    readonly identity?: pulumi.Input<inputs.databox.v20200401.ResourceIdentity>;
    /**
     * The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Resource Group Name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku type.
     */
    readonly sku: pulumi.Input<inputs.databox.v20200401.Sku>;
    /**
     * The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of the data transfer.
     */
    readonly transferType: pulumi.Input<string>;
}
