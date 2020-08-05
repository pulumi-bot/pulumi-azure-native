import * as azurerm from "../../sdk/nodejs";
import * as pulumi from "@pulumi/pulumi";

export type Location =
    "eastasia"           |
    "southeastasia"      |
    "centralus"          |
    "eastus"             |
    "eastus2"            |
    "westus"             |
    "westus2"            |
    "northcentralus"     |
    "southcentralus"     |
    "northeurope"        |
    "westeurope"         |
    "japanwest"          |
    "japaneast"          |
    "brazilsouth"        |
    "australiaeast"      |
    "australiasoutheast" |
    "australiacentral"   |
    "australiacentral2"  |
    "southindia"         |
    "centralindia"       |
    "westindia"          |
    "canadacentral"      |
    "canadaeast"         |
    "uksouth"            |
    "ukwest"             |
    "westcentralus"      |
    "koreacentral"       |
    "koreasouth"         |
    "francecentral"      |
    "francesouth"        |
    "southafricanorth"   |
    "southafricawest"    |
    "uaecentral"         |
    "uaenorth"           ;

type Api = "Sql" | "MongoDB" | "Cassandra" | "Gremlin" | "Table";

type DefaultConsistencyPolicy = {
    defaultConsistencyLevel: "Eventual" | "ConsistentPrefix" | "Session" | "Strong",
}

type BoundedStalenessConsistencyPolicy = {
    defaultConsistencyLevel: "BoundedStaleness",
    maxStalenessPrefix: number;
    maxIntervalInSeconds: number;
}

type ConsistencyPolicy = DefaultConsistencyPolicy | BoundedStalenessConsistencyPolicy;

type SingleLocation = {
    type: "single";
    region: Location;
}

type MultipleLocations = {
    type: "multiple";
    regions: Location[];
    enableMultipleWriteLocations: boolean;
    enableAutomaticFailover: boolean;
}

type Locationable = SingleLocation | MultipleLocations;

interface MyDatabaseAccountArgs {
    resourceGroup: azurerm.resources.v20200601.ResourceGroup;
    api: Api;
    consisencyPolicy: ConsistencyPolicy;
    location: Locationable;
};

export class DatabaseAccount {
    public id: pulumi.Output<string>;

    constructor(name: string, args: MyDatabaseAccountArgs) {

        let kind = "GlobalDocumentDB";
        let capabilities: azurerm.types.input.documentdb.v20200401.Capability[] = [];
        if (args.api === "MongoDB") {
            kind = "MongoDB";
        } else if (args.api === "Cassandra") {
            capabilities = [{ name: "EnableCassandra" }];
        } else if (args.api === "Gremlin") {
            capabilities = [{ name: "EnableGremlin" }];
        } else if (args.api === "Table") {
            capabilities = [{ name: "EnableTable" }];
        }

        const locations = args.location.type === "single" ? [args.location.region] : args.location.regions;
        const enableMultipleWriteLocations = args.location.type === "single" ? false : args.location.enableMultipleWriteLocations;
        const enableAutomaticFailover = args.location.type === "single" ? false : args.location.enableAutomaticFailover;

        const cosmosdbAccount = new azurerm.documentdb.v20200401.DatabaseAccount(name, {
            resourceGroupName: args.resourceGroup.name,
            name,
            location: locations[0],
            kind,
            consistencyPolicy: args.consisencyPolicy,        
            locations: locations.map((l, i) => ({
                locationName: l,
                failoverPriority: i,
                isZoneRedundant: false
            })),
            databaseAccountOfferType: "Standard",
            enableAutomaticFailover,
            capabilities,
            enableMultipleWriteLocations,
        });
        this.id = cosmosdbAccount.id;
    }
}
