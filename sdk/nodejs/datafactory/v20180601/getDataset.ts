// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDataset(args: GetDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetDatasetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datafactory/v20180601:getDataset", {
        "datasetName": args.datasetName,
        "factoryName": args.factoryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDatasetArgs {
    /**
     * The dataset name.
     */
    readonly datasetName: string;
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Dataset resource type.
 */
export interface GetDatasetResult {
    /**
     * Etag identifies change in the resource.
     */
    readonly etag: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Dataset properties.
     */
    readonly properties: outputs.datafactory.v20180601.AmazonMWSObjectDatasetResponse | outputs.datafactory.v20180601.AmazonRedshiftTableDatasetResponse | outputs.datafactory.v20180601.AmazonS3DatasetResponse | outputs.datafactory.v20180601.AvroDatasetResponse | outputs.datafactory.v20180601.AzureBlobDatasetResponse | outputs.datafactory.v20180601.AzureBlobFSDatasetResponse | outputs.datafactory.v20180601.AzureDataExplorerTableDatasetResponse | outputs.datafactory.v20180601.AzureDataLakeStoreDatasetResponse | outputs.datafactory.v20180601.AzureDatabricksDeltaLakeDatasetResponse | outputs.datafactory.v20180601.AzureMariaDBTableDatasetResponse | outputs.datafactory.v20180601.AzureMySqlTableDatasetResponse | outputs.datafactory.v20180601.AzurePostgreSqlTableDatasetResponse | outputs.datafactory.v20180601.AzureSearchIndexDatasetResponse | outputs.datafactory.v20180601.AzureSqlDWTableDatasetResponse | outputs.datafactory.v20180601.AzureSqlMITableDatasetResponse | outputs.datafactory.v20180601.AzureSqlTableDatasetResponse | outputs.datafactory.v20180601.AzureTableDatasetResponse | outputs.datafactory.v20180601.BinaryDatasetResponse | outputs.datafactory.v20180601.CassandraTableDatasetResponse | outputs.datafactory.v20180601.CommonDataServiceForAppsEntityDatasetResponse | outputs.datafactory.v20180601.ConcurObjectDatasetResponse | outputs.datafactory.v20180601.CosmosDbMongoDbApiCollectionDatasetResponse | outputs.datafactory.v20180601.CosmosDbSqlApiCollectionDatasetResponse | outputs.datafactory.v20180601.CouchbaseTableDatasetResponse | outputs.datafactory.v20180601.CustomDatasetResponse | outputs.datafactory.v20180601.Db2TableDatasetResponse | outputs.datafactory.v20180601.DelimitedTextDatasetResponse | outputs.datafactory.v20180601.DocumentDbCollectionDatasetResponse | outputs.datafactory.v20180601.DrillTableDatasetResponse | outputs.datafactory.v20180601.DynamicsAXResourceDatasetResponse | outputs.datafactory.v20180601.DynamicsCrmEntityDatasetResponse | outputs.datafactory.v20180601.DynamicsEntityDatasetResponse | outputs.datafactory.v20180601.EloquaObjectDatasetResponse | outputs.datafactory.v20180601.ExcelDatasetResponse | outputs.datafactory.v20180601.FileShareDatasetResponse | outputs.datafactory.v20180601.GoogleAdWordsObjectDatasetResponse | outputs.datafactory.v20180601.GoogleBigQueryObjectDatasetResponse | outputs.datafactory.v20180601.GreenplumTableDatasetResponse | outputs.datafactory.v20180601.HBaseObjectDatasetResponse | outputs.datafactory.v20180601.HiveObjectDatasetResponse | outputs.datafactory.v20180601.HttpDatasetResponse | outputs.datafactory.v20180601.HubspotObjectDatasetResponse | outputs.datafactory.v20180601.ImpalaObjectDatasetResponse | outputs.datafactory.v20180601.InformixTableDatasetResponse | outputs.datafactory.v20180601.JiraObjectDatasetResponse | outputs.datafactory.v20180601.JsonDatasetResponse | outputs.datafactory.v20180601.MagentoObjectDatasetResponse | outputs.datafactory.v20180601.MariaDBTableDatasetResponse | outputs.datafactory.v20180601.MarketoObjectDatasetResponse | outputs.datafactory.v20180601.MicrosoftAccessTableDatasetResponse | outputs.datafactory.v20180601.MongoDbCollectionDatasetResponse | outputs.datafactory.v20180601.MongoDbV2CollectionDatasetResponse | outputs.datafactory.v20180601.MySqlTableDatasetResponse | outputs.datafactory.v20180601.NetezzaTableDatasetResponse | outputs.datafactory.v20180601.ODataResourceDatasetResponse | outputs.datafactory.v20180601.OdbcTableDatasetResponse | outputs.datafactory.v20180601.Office365DatasetResponse | outputs.datafactory.v20180601.OracleServiceCloudObjectDatasetResponse | outputs.datafactory.v20180601.OracleTableDatasetResponse | outputs.datafactory.v20180601.OrcDatasetResponse | outputs.datafactory.v20180601.ParquetDatasetResponse | outputs.datafactory.v20180601.PaypalObjectDatasetResponse | outputs.datafactory.v20180601.PhoenixObjectDatasetResponse | outputs.datafactory.v20180601.PostgreSqlTableDatasetResponse | outputs.datafactory.v20180601.PrestoObjectDatasetResponse | outputs.datafactory.v20180601.QuickBooksObjectDatasetResponse | outputs.datafactory.v20180601.RelationalTableDatasetResponse | outputs.datafactory.v20180601.ResponsysObjectDatasetResponse | outputs.datafactory.v20180601.RestResourceDatasetResponse | outputs.datafactory.v20180601.SalesforceMarketingCloudObjectDatasetResponse | outputs.datafactory.v20180601.SalesforceObjectDatasetResponse | outputs.datafactory.v20180601.SalesforceServiceCloudObjectDatasetResponse | outputs.datafactory.v20180601.SapBwCubeDatasetResponse | outputs.datafactory.v20180601.SapCloudForCustomerResourceDatasetResponse | outputs.datafactory.v20180601.SapEccResourceDatasetResponse | outputs.datafactory.v20180601.SapHanaTableDatasetResponse | outputs.datafactory.v20180601.SapOpenHubTableDatasetResponse | outputs.datafactory.v20180601.SapTableResourceDatasetResponse | outputs.datafactory.v20180601.ServiceNowObjectDatasetResponse | outputs.datafactory.v20180601.SharePointOnlineListResourceDatasetResponse | outputs.datafactory.v20180601.ShopifyObjectDatasetResponse | outputs.datafactory.v20180601.SnowflakeDatasetResponse | outputs.datafactory.v20180601.SparkObjectDatasetResponse | outputs.datafactory.v20180601.SqlServerTableDatasetResponse | outputs.datafactory.v20180601.SquareObjectDatasetResponse | outputs.datafactory.v20180601.SybaseTableDatasetResponse | outputs.datafactory.v20180601.TeradataTableDatasetResponse | outputs.datafactory.v20180601.VerticaTableDatasetResponse | outputs.datafactory.v20180601.WebTableDatasetResponse | outputs.datafactory.v20180601.XeroObjectDatasetResponse | outputs.datafactory.v20180601.XmlDatasetResponse | outputs.datafactory.v20180601.ZohoObjectDatasetResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
