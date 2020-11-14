// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Linked service resource type.
 */
export class LinkedService extends pulumi.CustomResource {
    /**
     * Get an existing LinkedService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LinkedService {
        return new LinkedService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:datafactory/v20170901preview:LinkedService';

    /**
     * Returns true if the given object is an instance of LinkedService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinkedService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinkedService.__pulumiType;
    }

    /**
     * Etag identifies change in the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of linked service.
     */
    public readonly properties!: pulumi.Output<outputs.datafactory.v20170901preview.AmazonMWSLinkedServiceResponse | outputs.datafactory.v20170901preview.AmazonRedshiftLinkedServiceResponse | outputs.datafactory.v20170901preview.AmazonS3LinkedServiceResponse | outputs.datafactory.v20170901preview.AzureBatchLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureDataLakeAnalyticsLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureDataLakeStoreLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureDatabricksLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureKeyVaultLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureMLLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureMySqlLinkedServiceResponse | outputs.datafactory.v20170901preview.AzurePostgreSqlLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureSearchLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureSqlDWLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureSqlDatabaseLinkedServiceResponse | outputs.datafactory.v20170901preview.AzureStorageLinkedServiceResponse | outputs.datafactory.v20170901preview.CassandraLinkedServiceResponse | outputs.datafactory.v20170901preview.ConcurLinkedServiceResponse | outputs.datafactory.v20170901preview.CosmosDbLinkedServiceResponse | outputs.datafactory.v20170901preview.CouchbaseLinkedServiceResponse | outputs.datafactory.v20170901preview.CustomDataSourceLinkedServiceResponse | outputs.datafactory.v20170901preview.Db2LinkedServiceResponse | outputs.datafactory.v20170901preview.DrillLinkedServiceResponse | outputs.datafactory.v20170901preview.DynamicsLinkedServiceResponse | outputs.datafactory.v20170901preview.EloquaLinkedServiceResponse | outputs.datafactory.v20170901preview.FileServerLinkedServiceResponse | outputs.datafactory.v20170901preview.FtpServerLinkedServiceResponse | outputs.datafactory.v20170901preview.GoogleBigQueryLinkedServiceResponse | outputs.datafactory.v20170901preview.GreenplumLinkedServiceResponse | outputs.datafactory.v20170901preview.HBaseLinkedServiceResponse | outputs.datafactory.v20170901preview.HDInsightLinkedServiceResponse | outputs.datafactory.v20170901preview.HDInsightOnDemandLinkedServiceResponse | outputs.datafactory.v20170901preview.HdfsLinkedServiceResponse | outputs.datafactory.v20170901preview.HiveLinkedServiceResponse | outputs.datafactory.v20170901preview.HttpLinkedServiceResponse | outputs.datafactory.v20170901preview.HubspotLinkedServiceResponse | outputs.datafactory.v20170901preview.ImpalaLinkedServiceResponse | outputs.datafactory.v20170901preview.JiraLinkedServiceResponse | outputs.datafactory.v20170901preview.MagentoLinkedServiceResponse | outputs.datafactory.v20170901preview.MariaDBLinkedServiceResponse | outputs.datafactory.v20170901preview.MarketoLinkedServiceResponse | outputs.datafactory.v20170901preview.MongoDbLinkedServiceResponse | outputs.datafactory.v20170901preview.MySqlLinkedServiceResponse | outputs.datafactory.v20170901preview.NetezzaLinkedServiceResponse | outputs.datafactory.v20170901preview.ODataLinkedServiceResponse | outputs.datafactory.v20170901preview.OdbcLinkedServiceResponse | outputs.datafactory.v20170901preview.OracleLinkedServiceResponse | outputs.datafactory.v20170901preview.PaypalLinkedServiceResponse | outputs.datafactory.v20170901preview.PhoenixLinkedServiceResponse | outputs.datafactory.v20170901preview.PostgreSqlLinkedServiceResponse | outputs.datafactory.v20170901preview.PrestoLinkedServiceResponse | outputs.datafactory.v20170901preview.QuickBooksLinkedServiceResponse | outputs.datafactory.v20170901preview.ResponsysLinkedServiceResponse | outputs.datafactory.v20170901preview.SalesforceLinkedServiceResponse | outputs.datafactory.v20170901preview.SalesforceMarketingCloudLinkedServiceResponse | outputs.datafactory.v20170901preview.SapBWLinkedServiceResponse | outputs.datafactory.v20170901preview.SapCloudForCustomerLinkedServiceResponse | outputs.datafactory.v20170901preview.SapEccLinkedServiceResponse | outputs.datafactory.v20170901preview.SapHanaLinkedServiceResponse | outputs.datafactory.v20170901preview.ServiceNowLinkedServiceResponse | outputs.datafactory.v20170901preview.SftpServerLinkedServiceResponse | outputs.datafactory.v20170901preview.ShopifyLinkedServiceResponse | outputs.datafactory.v20170901preview.SparkLinkedServiceResponse | outputs.datafactory.v20170901preview.SqlServerLinkedServiceResponse | outputs.datafactory.v20170901preview.SquareLinkedServiceResponse | outputs.datafactory.v20170901preview.SybaseLinkedServiceResponse | outputs.datafactory.v20170901preview.TeradataLinkedServiceResponse | outputs.datafactory.v20170901preview.VerticaLinkedServiceResponse | outputs.datafactory.v20170901preview.WebLinkedServiceResponse | outputs.datafactory.v20170901preview.XeroLinkedServiceResponse | outputs.datafactory.v20170901preview.ZohoLinkedServiceResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LinkedService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkedServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.factoryName === undefined) {
                throw new Error("Missing required property 'factoryName'");
            }
            if (!args || args.linkedServiceName === undefined) {
                throw new Error("Missing required property 'linkedServiceName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["factoryName"] = args ? args.factoryName : undefined;
            inputs["linkedServiceName"] = args ? args.linkedServiceName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:datafactory/latest:LinkedService" }, { type: "azure-nextgen:datafactory/v20180601:LinkedService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LinkedService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LinkedService resource.
 */
export interface LinkedServiceArgs {
    /**
     * The factory name.
     */
    readonly factoryName: pulumi.Input<string>;
    /**
     * The linked service name.
     */
    readonly linkedServiceName: pulumi.Input<string>;
    /**
     * Properties of linked service.
     */
    readonly properties: pulumi.Input<inputs.datafactory.v20170901preview.AmazonMWSLinkedService | inputs.datafactory.v20170901preview.AmazonRedshiftLinkedService | inputs.datafactory.v20170901preview.AmazonS3LinkedService | inputs.datafactory.v20170901preview.AzureBatchLinkedService | inputs.datafactory.v20170901preview.AzureDataLakeAnalyticsLinkedService | inputs.datafactory.v20170901preview.AzureDataLakeStoreLinkedService | inputs.datafactory.v20170901preview.AzureDatabricksLinkedService | inputs.datafactory.v20170901preview.AzureKeyVaultLinkedService | inputs.datafactory.v20170901preview.AzureMLLinkedService | inputs.datafactory.v20170901preview.AzureMySqlLinkedService | inputs.datafactory.v20170901preview.AzurePostgreSqlLinkedService | inputs.datafactory.v20170901preview.AzureSearchLinkedService | inputs.datafactory.v20170901preview.AzureSqlDWLinkedService | inputs.datafactory.v20170901preview.AzureSqlDatabaseLinkedService | inputs.datafactory.v20170901preview.AzureStorageLinkedService | inputs.datafactory.v20170901preview.CassandraLinkedService | inputs.datafactory.v20170901preview.ConcurLinkedService | inputs.datafactory.v20170901preview.CosmosDbLinkedService | inputs.datafactory.v20170901preview.CouchbaseLinkedService | inputs.datafactory.v20170901preview.CustomDataSourceLinkedService | inputs.datafactory.v20170901preview.Db2LinkedService | inputs.datafactory.v20170901preview.DrillLinkedService | inputs.datafactory.v20170901preview.DynamicsLinkedService | inputs.datafactory.v20170901preview.EloquaLinkedService | inputs.datafactory.v20170901preview.FileServerLinkedService | inputs.datafactory.v20170901preview.FtpServerLinkedService | inputs.datafactory.v20170901preview.GoogleBigQueryLinkedService | inputs.datafactory.v20170901preview.GreenplumLinkedService | inputs.datafactory.v20170901preview.HBaseLinkedService | inputs.datafactory.v20170901preview.HDInsightLinkedService | inputs.datafactory.v20170901preview.HDInsightOnDemandLinkedService | inputs.datafactory.v20170901preview.HdfsLinkedService | inputs.datafactory.v20170901preview.HiveLinkedService | inputs.datafactory.v20170901preview.HttpLinkedService | inputs.datafactory.v20170901preview.HubspotLinkedService | inputs.datafactory.v20170901preview.ImpalaLinkedService | inputs.datafactory.v20170901preview.JiraLinkedService | inputs.datafactory.v20170901preview.MagentoLinkedService | inputs.datafactory.v20170901preview.MariaDBLinkedService | inputs.datafactory.v20170901preview.MarketoLinkedService | inputs.datafactory.v20170901preview.MongoDbLinkedService | inputs.datafactory.v20170901preview.MySqlLinkedService | inputs.datafactory.v20170901preview.NetezzaLinkedService | inputs.datafactory.v20170901preview.ODataLinkedService | inputs.datafactory.v20170901preview.OdbcLinkedService | inputs.datafactory.v20170901preview.OracleLinkedService | inputs.datafactory.v20170901preview.PaypalLinkedService | inputs.datafactory.v20170901preview.PhoenixLinkedService | inputs.datafactory.v20170901preview.PostgreSqlLinkedService | inputs.datafactory.v20170901preview.PrestoLinkedService | inputs.datafactory.v20170901preview.QuickBooksLinkedService | inputs.datafactory.v20170901preview.ResponsysLinkedService | inputs.datafactory.v20170901preview.SalesforceLinkedService | inputs.datafactory.v20170901preview.SalesforceMarketingCloudLinkedService | inputs.datafactory.v20170901preview.SapBWLinkedService | inputs.datafactory.v20170901preview.SapCloudForCustomerLinkedService | inputs.datafactory.v20170901preview.SapEccLinkedService | inputs.datafactory.v20170901preview.SapHanaLinkedService | inputs.datafactory.v20170901preview.ServiceNowLinkedService | inputs.datafactory.v20170901preview.SftpServerLinkedService | inputs.datafactory.v20170901preview.ShopifyLinkedService | inputs.datafactory.v20170901preview.SparkLinkedService | inputs.datafactory.v20170901preview.SqlServerLinkedService | inputs.datafactory.v20170901preview.SquareLinkedService | inputs.datafactory.v20170901preview.SybaseLinkedService | inputs.datafactory.v20170901preview.TeradataLinkedService | inputs.datafactory.v20170901preview.VerticaLinkedService | inputs.datafactory.v20170901preview.WebLinkedService | inputs.datafactory.v20170901preview.XeroLinkedService | inputs.datafactory.v20170901preview.ZohoLinkedService>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
