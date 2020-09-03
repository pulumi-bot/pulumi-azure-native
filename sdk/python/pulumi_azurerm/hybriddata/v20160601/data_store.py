# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DataStore']


class DataStore(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 customer_secrets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CustomerSecretArgs']]]]] = None,
                 data_manager_name: Optional[pulumi.Input[str]] = None,
                 data_store_name: Optional[pulumi.Input[str]] = None,
                 data_store_type_id: Optional[pulumi.Input[str]] = None,
                 extended_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Data store.

        ## Example Usage
        ### DataStores_CreateOrUpdate_DataSinkPUT162

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        data_store = azurerm.hybriddata.v20160601.DataStore("dataStore",
            customer_secrets=[
                {
                    "algorithm": "RSA1_5",
                    "keyIdentifier": "StorageAccountAccessKey",
                    "keyValue": "Pses92T1jTpPBH2roHgNKArWVv57WomubD/9ukE2d0M89gIVnzc/0bfeoJVf0Med6Uvt4mzmOghFEVxqoBNXzKLgsCLVqvPkqdst4WzZbeel3k/NVkfDdf04eNPAm1FwM/hWZQlG3lGr/olCihW8AFHoiLEZWK8DC6UmHl8ittuuhzY/Ct8R5VXdMWavLdFg8G66TCSyH2aF/eeqzHcOBP4XFgbF2NxuvmEd/cY+y5lEon3TfdDwI0JcOumf5s4zHTWM5+StWa3SsvKxGPpJ27xik2FBo2kEqrtAAByi6HxXSinJB8DwXZicHIwjaOiTeiJUMADwZXwv1PLpwo5d1Q==:QBzHKcxhJkiIeLtFx84oLqDUd5p+oM2AwoKtZIgYVZhfKFw5VW2BfsigL2K7AGxyTNatwn6JZm9ylo8YhRZrn0eIVqLR4gCiRSwDHI7i6R/tqTfx8ZO/aJy6rTh/WW8d6vZOXXGeuRDAz6fYfjQKb9J/OhTq3cjfVouLt6bKdZsZve08NVZq0sNBYZftCabcOhVg5hamuDhQhemwqFMn6l1xrCWcq4e5YgJ90fbK5N66Wj5LNr2dU+scHH7YfM8a3IIhq51TObhXZ59oNnLhLGGA8j0K43MMKtQAnqpBc+hmwgwc8/DZLod1wnaPbJW5/fQ2HkF7vH9xakIip4bZ9Q==",
                },
                {
                    "algorithm": "RSA1_5",
                    "keyIdentifier": "StorageAccountAccessKeyForQueue",
                    "keyValue": "Pses92T1jTpPBH2roHgNKArWVv57WomubD/9ukE2d0M89gIVnzc/0bfeoJVf0Med6Uvt4mzmOghFEVxqoBNXzKLgsCLVqvPkqdst4WzZbeel3k/NVkfDdf04eNPAm1FwM/hWZQlG3lGr/olCihW8AFHoiLEZWK8DC6UmHl8ittuuhzY/Ct8R5VXdMWavLdFg8G66TCSyH2aF/eeqzHcOBP4XFgbF2NxuvmEd/cY+y5lEon3TfdDwI0JcOumf5s4zHTWM5+StWa3SsvKxGPpJ27xik2FBo2kEqrtAAByi6HxXSinJB8DwXZicHIwjaOiTeiJUMADwZXwv1PLpwo5d1Q==:QBzHKcxhJkiIeLtFx84oLqDUd5p+oM2AwoKtZIgYVZhfKFw5VW2BfsigL2K7AGxyTNatwn6JZm9ylo8YhRZrn0eIVqLR4gCiRSwDHI7i6R/tqTfx8ZO/aJy6rTh/WW8d6vZOXXGeuRDAz6fYfjQKb9J/OhTq3cjfVouLt6bKdZsZve08NVZq0sNBYZftCabcOhVg5hamuDhQhemwqFMn6l1xrCWcq4e5YgJ90fbK5N66Wj5LNr2dU+scHH7YfM8a3IIhq51TObhXZ59oNnLhLGGA8j0K43MMKtQAnqpBc+hmwgwc8/DZLod1wnaPbJW5/fQ2HkF7vH9xakIip4bZ9Q==",
                },
            ],
            data_manager_name="TestAzureSDKOperations",
            data_store_name="TestAzureStorage1",
            data_store_type_id="/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/AzureStorageAccount",
            extended_properties={
                "extendedSaKey": None,
                "extendedSaName": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
                "storageAccountNameForQueue": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
            },
            repository_id="/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.Storage/storageAccounts/dmsdatasink",
            resource_group_name="ResourceGroupForSDKTest",
            state="Enabled")

        ```
        ### DataStores_CreateOrUpdate_DataSourcePUT162

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        data_store = azurerm.hybriddata.v20160601.DataStore("dataStore",
            customer_secrets=[{
                "algorithm": "RSA1_5",
                "keyIdentifier": "ServiceEncryptionKey",
                "keyValue": "vZqtembBBg2RC/UyYwZiexGOqujLcMYCmaywqf0sURqIidjxlSp86FGz+T2eRnb1XlYCqFf1CzPzwLpwHEuTJ8LP5hTV1yUiM+YnyKHIGdlQajLcVcFy8ji9n+jSS4J9PjjHsr5AKzW1w+y76UgTEpX7K9kFDWFVyDGEujvuB2bYBlxlKolMCOu0WHZYkBBYLob6a3mQgCHbXYj1mqTmdhPW+J+8tyBCzG6cjlvRJ9hcp9Ss3HV9TV6hrbqlUU3lE1FX8O5Dr6/TXi7tIU7hGfmS5psE0Kz+2PsJTX1R1AbkBpKObPwPxPoC5jCXFxwfmZOrNQdjZ7nu5+JHaLZylw==:tS9oSCAvIwOrkYRyD/jLahSLZypl4eNexW5N/pGqf9vsVfzMhmxob+O/Io48uCPxvtdDksef09OVXpxgaC65K2Og49W9rtRt8cvGyyC41cx5D2DP9fxAu7d/lREP9cWHgrRJlZ4JJFcqy+m+yqYKl3WPrTA2yoZpISGbWAPkj0Hk3IwRr1lmqKfCWtp0jNHnrIJmQ5BQaDLGXpohKQSrrftqz7TdBCYuorSntQz8pqHgc8MTiYMgMtgZ+HRKQ1F5ctOlP+6LJMS6/OFl/tnYb5BD6rn/RufB4OHhVDe9ZD5GMtkwqkUvU9b1v2n31mb63JLApxIi/o8OsSpkA8ZTCg==",
            }],
            data_manager_name="TestAzureSDKOperations",
            data_store_name="TestStorSimpleSource1",
            data_store_type_id="/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStoreTypes/StorSimple8000Series",
            extended_properties={
                "extendedSaKey": None,
                "resource_id": "/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/dmsdatasource/providers/Microsoft.StorSimple/managers/dmsdatasource",
            },
            repository_id="/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/dmsdatasource/providers/Microsoft.StorSimple/managers/dmsdatasource",
            resource_group_name="ResourceGroupForSDKTest",
            state="Enabled")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CustomerSecretArgs']]]] customer_secrets: List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        :param pulumi.Input[str] data_manager_name: The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        :param pulumi.Input[str] data_store_name: The data store/repository name to be created or updated.
        :param pulumi.Input[str] data_store_type_id: The arm id of the data store type.
        :param pulumi.Input[Mapping[str, Any]] extended_properties: A generic json used differently by each data source type.
        :param pulumi.Input[str] repository_id: Arm Id for the manager resource to which the data source is associated. This is optional.
        :param pulumi.Input[str] resource_group_name: The Resource Group Name
        :param pulumi.Input[str] state: State of the data source.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['customer_secrets'] = customer_secrets
            if data_manager_name is None:
                raise TypeError("Missing required property 'data_manager_name'")
            __props__['data_manager_name'] = data_manager_name
            if data_store_name is None:
                raise TypeError("Missing required property 'data_store_name'")
            __props__['data_store_name'] = data_store_name
            if data_store_type_id is None:
                raise TypeError("Missing required property 'data_store_type_id'")
            __props__['data_store_type_id'] = data_store_type_id
            __props__['extended_properties'] = extended_properties
            __props__['repository_id'] = repository_id
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if state is None:
                raise TypeError("Missing required property 'state'")
            __props__['state'] = state
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:hybriddata/latest:DataStore"), pulumi.Alias(type_="azurerm:hybriddata/v20190601:DataStore")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataStore, __self__).__init__(
            'azurerm:hybriddata/v20160601:DataStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataStore':
        """
        Get an existing DataStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DataStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customerSecrets")
    def customer_secrets(self) -> pulumi.Output[Optional[List['outputs.CustomerSecretResponse']]]:
        """
        List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        """
        return pulumi.get(self, "customer_secrets")

    @property
    @pulumi.getter(name="dataStoreTypeId")
    def data_store_type_id(self) -> pulumi.Output[str]:
        """
        The arm id of the data store type.
        """
        return pulumi.get(self, "data_store_type_id")

    @property
    @pulumi.getter(name="extendedProperties")
    def extended_properties(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A generic json used differently by each data source type.
        """
        return pulumi.get(self, "extended_properties")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[Optional[str]]:
        """
        Arm Id for the manager resource to which the data source is associated. This is optional.
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the data source.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the object.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

