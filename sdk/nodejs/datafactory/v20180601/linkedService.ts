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
    public static readonly __pulumiType = 'azure-nextgen:datafactory/v20180601:LinkedService';

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
    public readonly properties!: pulumi.Output<outputs.datafactory.v20180601.AmazonMWSLinkedServiceResponse | outputs.datafactory.v20180601.AmazonRedshiftLinkedServiceResponse | outputs.datafactory.v20180601.AmazonS3LinkedServiceResponse | outputs.datafactory.v20180601.AzureBatchLinkedServiceResponse | outputs.datafactory.v20180601.AzureBlobFSLinkedServiceResponse | outputs.datafactory.v20180601.AzureBlobStorageLinkedServiceResponse | outputs.datafactory.v20180601.AzureDataExplorerLinkedServiceResponse | outputs.datafactory.v20180601.AzureDataLakeAnalyticsLinkedServiceResponse | outputs.datafactory.v20180601.AzureDataLakeStoreLinkedServiceResponse | outputs.datafactory.v20180601.AzureDatabricksDeltaLakeLinkedServiceResponse | outputs.datafactory.v20180601.AzureDatabricksLinkedServiceResponse | outputs.datafactory.v20180601.AzureFileStorageLinkedServiceResponse | outputs.datafactory.v20180601.AzureFunctionLinkedServiceResponse | outputs.datafactory.v20180601.AzureKeyVaultLinkedServiceResponse | outputs.datafactory.v20180601.AzureMLLinkedServiceResponse | outputs.datafactory.v20180601.AzureMLServiceLinkedServiceResponse | outputs.datafactory.v20180601.AzureMariaDBLinkedServiceResponse | outputs.datafactory.v20180601.AzureMySqlLinkedServiceResponse | outputs.datafactory.v20180601.AzurePostgreSqlLinkedServiceResponse | outputs.datafactory.v20180601.AzureSearchLinkedServiceResponse | outputs.datafactory.v20180601.AzureSqlDWLinkedServiceResponse | outputs.datafactory.v20180601.AzureSqlDatabaseLinkedServiceResponse | outputs.datafactory.v20180601.AzureSqlMILinkedServiceResponse | outputs.datafactory.v20180601.AzureStorageLinkedServiceResponse | outputs.datafactory.v20180601.AzureTableStorageLinkedServiceResponse | outputs.datafactory.v20180601.CassandraLinkedServiceResponse | outputs.datafactory.v20180601.CommonDataServiceForAppsLinkedServiceResponse | outputs.datafactory.v20180601.ConcurLinkedServiceResponse | outputs.datafactory.v20180601.CosmosDbLinkedServiceResponse | outputs.datafactory.v20180601.CosmosDbMongoDbApiLinkedServiceResponse | outputs.datafactory.v20180601.CouchbaseLinkedServiceResponse | outputs.datafactory.v20180601.CustomDataSourceLinkedServiceResponse | outputs.datafactory.v20180601.Db2LinkedServiceResponse | outputs.datafactory.v20180601.DrillLinkedServiceResponse | outputs.datafactory.v20180601.DynamicsAXLinkedServiceResponse | outputs.datafactory.v20180601.DynamicsCrmLinkedServiceResponse | outputs.datafactory.v20180601.DynamicsLinkedServiceResponse | outputs.datafactory.v20180601.EloquaLinkedServiceResponse | outputs.datafactory.v20180601.FileServerLinkedServiceResponse | outputs.datafactory.v20180601.FtpServerLinkedServiceResponse | outputs.datafactory.v20180601.GoogleAdWordsLinkedServiceResponse | outputs.datafactory.v20180601.GoogleBigQueryLinkedServiceResponse | outputs.datafactory.v20180601.GoogleCloudStorageLinkedServiceResponse | outputs.datafactory.v20180601.GreenplumLinkedServiceResponse | outputs.datafactory.v20180601.HBaseLinkedServiceResponse | outputs.datafactory.v20180601.HDInsightLinkedServiceResponse | outputs.datafactory.v20180601.HDInsightOnDemandLinkedServiceResponse | outputs.datafactory.v20180601.HdfsLinkedServiceResponse | outputs.datafactory.v20180601.HiveLinkedServiceResponse | outputs.datafactory.v20180601.HttpLinkedServiceResponse | outputs.datafactory.v20180601.HubspotLinkedServiceResponse | outputs.datafactory.v20180601.ImpalaLinkedServiceResponse | outputs.datafactory.v20180601.InformixLinkedServiceResponse | outputs.datafactory.v20180601.JiraLinkedServiceResponse | outputs.datafactory.v20180601.MagentoLinkedServiceResponse | outputs.datafactory.v20180601.MariaDBLinkedServiceResponse | outputs.datafactory.v20180601.MarketoLinkedServiceResponse | outputs.datafactory.v20180601.MicrosoftAccessLinkedServiceResponse | outputs.datafactory.v20180601.MongoDbAtlasLinkedServiceResponse | outputs.datafactory.v20180601.MongoDbLinkedServiceResponse | outputs.datafactory.v20180601.MongoDbV2LinkedServiceResponse | outputs.datafactory.v20180601.MySqlLinkedServiceResponse | outputs.datafactory.v20180601.NetezzaLinkedServiceResponse | outputs.datafactory.v20180601.ODataLinkedServiceResponse | outputs.datafactory.v20180601.OdbcLinkedServiceResponse | outputs.datafactory.v20180601.Office365LinkedServiceResponse | outputs.datafactory.v20180601.OracleLinkedServiceResponse | outputs.datafactory.v20180601.OracleServiceCloudLinkedServiceResponse | outputs.datafactory.v20180601.PaypalLinkedServiceResponse | outputs.datafactory.v20180601.PhoenixLinkedServiceResponse | outputs.datafactory.v20180601.PostgreSqlLinkedServiceResponse | outputs.datafactory.v20180601.PrestoLinkedServiceResponse | outputs.datafactory.v20180601.QuickBooksLinkedServiceResponse | outputs.datafactory.v20180601.ResponsysLinkedServiceResponse | outputs.datafactory.v20180601.RestServiceLinkedServiceResponse | outputs.datafactory.v20180601.SalesforceLinkedServiceResponse | outputs.datafactory.v20180601.SalesforceMarketingCloudLinkedServiceResponse | outputs.datafactory.v20180601.SalesforceServiceCloudLinkedServiceResponse | outputs.datafactory.v20180601.SapBWLinkedServiceResponse | outputs.datafactory.v20180601.SapCloudForCustomerLinkedServiceResponse | outputs.datafactory.v20180601.SapEccLinkedServiceResponse | outputs.datafactory.v20180601.SapHanaLinkedServiceResponse | outputs.datafactory.v20180601.SapOpenHubLinkedServiceResponse | outputs.datafactory.v20180601.SapTableLinkedServiceResponse | outputs.datafactory.v20180601.ServiceNowLinkedServiceResponse | outputs.datafactory.v20180601.SftpServerLinkedServiceResponse | outputs.datafactory.v20180601.SharePointOnlineListLinkedServiceResponse | outputs.datafactory.v20180601.ShopifyLinkedServiceResponse | outputs.datafactory.v20180601.SnowflakeLinkedServiceResponse | outputs.datafactory.v20180601.SparkLinkedServiceResponse | outputs.datafactory.v20180601.SqlServerLinkedServiceResponse | outputs.datafactory.v20180601.SquareLinkedServiceResponse | outputs.datafactory.v20180601.SybaseLinkedServiceResponse | outputs.datafactory.v20180601.TeradataLinkedServiceResponse | outputs.datafactory.v20180601.VerticaLinkedServiceResponse | outputs.datafactory.v20180601.WebLinkedServiceResponse | outputs.datafactory.v20180601.XeroLinkedServiceResponse | outputs.datafactory.v20180601.ZohoLinkedServiceResponse>;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:datafactory/latest:LinkedService" }, { type: "azure-nextgen:datafactory/v20170901preview:LinkedService" }] };
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
    readonly properties: pulumi.Input<inputs.datafactory.v20180601.AmazonMWSLinkedService | inputs.datafactory.v20180601.AmazonRedshiftLinkedService | inputs.datafactory.v20180601.AmazonS3LinkedService | inputs.datafactory.v20180601.AzureBatchLinkedService | inputs.datafactory.v20180601.AzureBlobFSLinkedService | inputs.datafactory.v20180601.AzureBlobStorageLinkedService | inputs.datafactory.v20180601.AzureDataExplorerLinkedService | inputs.datafactory.v20180601.AzureDataLakeAnalyticsLinkedService | inputs.datafactory.v20180601.AzureDataLakeStoreLinkedService | inputs.datafactory.v20180601.AzureDatabricksDeltaLakeLinkedService | inputs.datafactory.v20180601.AzureDatabricksLinkedService | inputs.datafactory.v20180601.AzureFileStorageLinkedService | inputs.datafactory.v20180601.AzureFunctionLinkedService | inputs.datafactory.v20180601.AzureKeyVaultLinkedService | inputs.datafactory.v20180601.AzureMLLinkedService | inputs.datafactory.v20180601.AzureMLServiceLinkedService | inputs.datafactory.v20180601.AzureMariaDBLinkedService | inputs.datafactory.v20180601.AzureMySqlLinkedService | inputs.datafactory.v20180601.AzurePostgreSqlLinkedService | inputs.datafactory.v20180601.AzureSearchLinkedService | inputs.datafactory.v20180601.AzureSqlDWLinkedService | inputs.datafactory.v20180601.AzureSqlDatabaseLinkedService | inputs.datafactory.v20180601.AzureSqlMILinkedService | inputs.datafactory.v20180601.AzureStorageLinkedService | inputs.datafactory.v20180601.AzureTableStorageLinkedService | inputs.datafactory.v20180601.CassandraLinkedService | inputs.datafactory.v20180601.CommonDataServiceForAppsLinkedService | inputs.datafactory.v20180601.ConcurLinkedService | inputs.datafactory.v20180601.CosmosDbLinkedService | inputs.datafactory.v20180601.CosmosDbMongoDbApiLinkedService | inputs.datafactory.v20180601.CouchbaseLinkedService | inputs.datafactory.v20180601.CustomDataSourceLinkedService | inputs.datafactory.v20180601.Db2LinkedService | inputs.datafactory.v20180601.DrillLinkedService | inputs.datafactory.v20180601.DynamicsAXLinkedService | inputs.datafactory.v20180601.DynamicsCrmLinkedService | inputs.datafactory.v20180601.DynamicsLinkedService | inputs.datafactory.v20180601.EloquaLinkedService | inputs.datafactory.v20180601.FileServerLinkedService | inputs.datafactory.v20180601.FtpServerLinkedService | inputs.datafactory.v20180601.GoogleAdWordsLinkedService | inputs.datafactory.v20180601.GoogleBigQueryLinkedService | inputs.datafactory.v20180601.GoogleCloudStorageLinkedService | inputs.datafactory.v20180601.GreenplumLinkedService | inputs.datafactory.v20180601.HBaseLinkedService | inputs.datafactory.v20180601.HDInsightLinkedService | inputs.datafactory.v20180601.HDInsightOnDemandLinkedService | inputs.datafactory.v20180601.HdfsLinkedService | inputs.datafactory.v20180601.HiveLinkedService | inputs.datafactory.v20180601.HttpLinkedService | inputs.datafactory.v20180601.HubspotLinkedService | inputs.datafactory.v20180601.ImpalaLinkedService | inputs.datafactory.v20180601.InformixLinkedService | inputs.datafactory.v20180601.JiraLinkedService | inputs.datafactory.v20180601.MagentoLinkedService | inputs.datafactory.v20180601.MariaDBLinkedService | inputs.datafactory.v20180601.MarketoLinkedService | inputs.datafactory.v20180601.MicrosoftAccessLinkedService | inputs.datafactory.v20180601.MongoDbAtlasLinkedService | inputs.datafactory.v20180601.MongoDbLinkedService | inputs.datafactory.v20180601.MongoDbV2LinkedService | inputs.datafactory.v20180601.MySqlLinkedService | inputs.datafactory.v20180601.NetezzaLinkedService | inputs.datafactory.v20180601.ODataLinkedService | inputs.datafactory.v20180601.OdbcLinkedService | inputs.datafactory.v20180601.Office365LinkedService | inputs.datafactory.v20180601.OracleLinkedService | inputs.datafactory.v20180601.OracleServiceCloudLinkedService | inputs.datafactory.v20180601.PaypalLinkedService | inputs.datafactory.v20180601.PhoenixLinkedService | inputs.datafactory.v20180601.PostgreSqlLinkedService | inputs.datafactory.v20180601.PrestoLinkedService | inputs.datafactory.v20180601.QuickBooksLinkedService | inputs.datafactory.v20180601.ResponsysLinkedService | inputs.datafactory.v20180601.RestServiceLinkedService | inputs.datafactory.v20180601.SalesforceLinkedService | inputs.datafactory.v20180601.SalesforceMarketingCloudLinkedService | inputs.datafactory.v20180601.SalesforceServiceCloudLinkedService | inputs.datafactory.v20180601.SapBWLinkedService | inputs.datafactory.v20180601.SapCloudForCustomerLinkedService | inputs.datafactory.v20180601.SapEccLinkedService | inputs.datafactory.v20180601.SapHanaLinkedService | inputs.datafactory.v20180601.SapOpenHubLinkedService | inputs.datafactory.v20180601.SapTableLinkedService | inputs.datafactory.v20180601.ServiceNowLinkedService | inputs.datafactory.v20180601.SftpServerLinkedService | inputs.datafactory.v20180601.SharePointOnlineListLinkedService | inputs.datafactory.v20180601.ShopifyLinkedService | inputs.datafactory.v20180601.SnowflakeLinkedService | inputs.datafactory.v20180601.SparkLinkedService | inputs.datafactory.v20180601.SqlServerLinkedService | inputs.datafactory.v20180601.SquareLinkedService | inputs.datafactory.v20180601.SybaseLinkedService | inputs.datafactory.v20180601.TeradataLinkedService | inputs.datafactory.v20180601.VerticaLinkedService | inputs.datafactory.v20180601.WebLinkedService | inputs.datafactory.v20180601.XeroLinkedService | inputs.datafactory.v20180601.ZohoLinkedService>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
