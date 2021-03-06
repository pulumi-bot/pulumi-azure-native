// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./account";
export * from "./dataSet";
export * from "./dataSetMapping";
export * from "./getAccount";
export * from "./getDataSet";
export * from "./getDataSetMapping";
export * from "./getInvitation";
export * from "./getShare";
export * from "./getShareSubscription";
export * from "./getSynchronizationSetting";
export * from "./getTrigger";
export * from "./invitation";
export * from "./listShareSubscriptionSourceShareSynchronizationSettings";
export * from "./listShareSubscriptionSynchronizationDetails";
export * from "./listShareSubscriptionSynchronizations";
export * from "./listShareSynchronizationDetails";
export * from "./listShareSynchronizations";
export * from "./share";
export * from "./shareSubscription";
export * from "./synchronizationSetting";
export * from "./trigger";

// Export enums:
export * from "../../types/enums/datashare/latest";

// Import resources to register:
import { Account } from "./account";
import { DataSet } from "./dataSet";
import { DataSetMapping } from "./dataSetMapping";
import { Invitation } from "./invitation";
import { Share } from "./share";
import { ShareSubscription } from "./shareSubscription";
import { SynchronizationSetting } from "./synchronizationSetting";
import { Trigger } from "./trigger";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:datashare/latest:Account":
                return new Account(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:DataSet":
                return new DataSet(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:DataSetMapping":
                return new DataSetMapping(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:Invitation":
                return new Invitation(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:Share":
                return new Share(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:ShareSubscription":
                return new ShareSubscription(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:SynchronizationSetting":
                return new SynchronizationSetting(name, <any>undefined, { urn })
            case "azure-native:datashare/latest:Trigger":
                return new Trigger(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "datashare/latest", _module)
