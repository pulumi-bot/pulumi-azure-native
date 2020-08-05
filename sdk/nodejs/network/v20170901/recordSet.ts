// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes a DNS record set (a collection of DNS records with the same name and type).
 */
export class RecordSet extends pulumi.CustomResource {
    /**
     * Get an existing RecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RecordSet {
        return new RecordSet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20170901:RecordSet';

    /**
     * Returns true if the given object is an instance of RecordSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordSet.__pulumiType;
    }

    /**
     * The etag of the record set.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the record set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the record set.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20170901.RecordSetPropertiesResponse>;
    /**
     * The type of the record set.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RecordSetArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.recordType === undefined) {
                throw new Error("Missing required property 'recordType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.zoneName === undefined) {
                throw new Error("Missing required property 'zoneName'");
            }
            inputs["ARecords"] = args ? args.ARecords : undefined;
            inputs["TTL"] = args ? args.TTL : undefined;
            inputs["aaaaRecords"] = args ? args.aaaaRecords : undefined;
            inputs["caaRecords"] = args ? args.caaRecords : undefined;
            inputs["cnameRecord"] = args ? args.cnameRecord : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["mxRecords"] = args ? args.mxRecords : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nsRecords"] = args ? args.nsRecords : undefined;
            inputs["ptrRecords"] = args ? args.ptrRecords : undefined;
            inputs["recordType"] = args ? args.recordType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["soaRecord"] = args ? args.soaRecord : undefined;
            inputs["srvRecords"] = args ? args.srvRecords : undefined;
            inputs["txtRecords"] = args ? args.txtRecords : undefined;
            inputs["zoneName"] = args ? args.zoneName : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RecordSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RecordSet resource.
 */
export interface RecordSetArgs {
    /**
     * The list of A records in the record set.
     */
    readonly ARecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.ARecord>[]>;
    /**
     * The TTL (time-to-live) of the records in the record set.
     */
    readonly TTL?: pulumi.Input<number>;
    /**
     * The list of AAAA records in the record set.
     */
    readonly aaaaRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.AaaaRecord>[]>;
    /**
     * The list of CAA records in the record set.
     */
    readonly caaRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.CaaRecord>[]>;
    /**
     * The CNAME record in the  record set.
     */
    readonly cnameRecord?: pulumi.Input<inputs.network.v20170901.CnameRecord>;
    /**
     * The etag of the record set.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The metadata attached to the record set.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of MX records in the record set.
     */
    readonly mxRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.MxRecord>[]>;
    /**
     * The name of the record set, relative to the name of the zone.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The list of NS records in the record set.
     */
    readonly nsRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.NsRecord>[]>;
    /**
     * The list of PTR records in the record set.
     */
    readonly ptrRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.PtrRecord>[]>;
    /**
     * The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
     */
    readonly recordType: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SOA record in the record set.
     */
    readonly soaRecord?: pulumi.Input<inputs.network.v20170901.SoaRecord>;
    /**
     * The list of SRV records in the record set.
     */
    readonly srvRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.SrvRecord>[]>;
    /**
     * The list of TXT records in the record set.
     */
    readonly txtRecords?: pulumi.Input<pulumi.Input<inputs.network.v20170901.TxtRecord>[]>;
    /**
     * The name of the DNS zone (without a terminating dot).
     */
    readonly zoneName: pulumi.Input<string>;
}
