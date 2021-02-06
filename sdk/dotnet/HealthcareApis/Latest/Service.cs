// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HealthcareApis.Latest
{
    /// <summary>
    /// The description of the service.
    /// Latest API Version: 2021-01-11.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:healthcareapis/latest:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// An etag associated with the resource, used for optimistic concurrency when editing it.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Setting indicating whether the service has a managed identity associated with it.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServicesResourceResponseIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The kind of the service.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The common properties of a service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ServicesPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:healthcareapis/latest:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:healthcareapis/latest:Service", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:healthcareapis/v20180820preview:Service"},
                    new Pulumi.Alias { Type = "azure-nextgen:healthcareapis/v20190916:Service"},
                    new Pulumi.Alias { Type = "azure-nextgen:healthcareapis/v20200315:Service"},
                    new Pulumi.Alias { Type = "azure-nextgen:healthcareapis/v20200330:Service"},
                    new Pulumi.Alias { Type = "azure-nextgen:healthcareapis/v20210111:Service"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An etag associated with the resource, used for optimistic concurrency when editing it.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Setting indicating whether the service has a managed identity associated with it.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServicesResourceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The kind of the service.
        /// </summary>
        [Input("kind", required: true)]
        public Input<Pulumi.AzureNextGen.HealthcareApis.Latest.Kind> Kind { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The common properties of a service.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ServicesPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the service instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the service instance.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServiceArgs()
        {
        }
    }
}
