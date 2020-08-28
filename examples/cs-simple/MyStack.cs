using Pulumi;
using Pulumi.AzureRM.Resources.V20200601;
using Pulumi.AzureRM.Storage.V20190601;
using Pulumi.AzureRM.Storage.V20190601.Inputs;
using Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var randomString = new RandomString("random", new RandomStringArgs
        {
            Length = 12,
            Special = false,
            Upper = false,
        });

        var resourceGroup = new ResourceGroup("resourceGroup", new ResourceGroupArgs
        {
            ResourceGroupName = randomString.Result,
            Location = "WestUS"
        });

        var storageAccount = new StorageAccount("sa", new StorageAccountArgs
        {
            ResourceGroupName = resourceGroup.Name,
            AccountName = randomString.Result,
            Location = "WestUS",
            Sku = new SkuArgs
            {
                Name = "Standard_LRS",
                Tier = "Standard"
            },
            Kind = "StorageV2"
        });
    }
}
