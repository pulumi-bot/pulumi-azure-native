// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20190701Preview
{
    /// <summary>
    /// The properties of the EventHubConsumerGroupInfo object.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:devices/v20190701preview:IotHubResourceEventHubConsumerGroup")]
    public partial class IotHubResourceEventHubConsumerGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The Event Hub-compatible consumer group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The tags.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableDictionary<string, string>> Properties { get; private set; } = null!;

        /// <summary>
        /// the resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IotHubResourceEventHubConsumerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotHubResourceEventHubConsumerGroup(string name, IotHubResourceEventHubConsumerGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:devices/v20190701preview:IotHubResourceEventHubConsumerGroup", name, args ?? new IotHubResourceEventHubConsumerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotHubResourceEventHubConsumerGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:devices/v20190701preview:IotHubResourceEventHubConsumerGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:devices/latest:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20160203:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20170119:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20170701:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20180122:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20180401:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20181201preview:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20190322:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20191104:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200301:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200401:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200615:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200801:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200831:IotHubResourceEventHubConsumerGroup"},
                    new Pulumi.Alias { Type = "azure-nextgen:devices/v20200831preview:IotHubResourceEventHubConsumerGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IotHubResourceEventHubConsumerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotHubResourceEventHubConsumerGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IotHubResourceEventHubConsumerGroup(name, id, options);
        }
    }

    public sealed class IotHubResourceEventHubConsumerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Event Hub-compatible endpoint in the IoT hub.
        /// </summary>
        [Input("eventHubEndpointName", required: true)]
        public Input<string> EventHubEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the consumer group to add.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public IotHubResourceEventHubConsumerGroupArgs()
        {
        }
    }
}
