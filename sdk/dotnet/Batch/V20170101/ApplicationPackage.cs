// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20170101
{
    /// <summary>
    /// An application package which represents a particular version of an application.
    /// </summary>
    public partial class ApplicationPackage : Pulumi.CustomResource
    {
        /// <summary>
        /// The format of the application package, if the package is active.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// The time at which the package was last activated, if the package is active.
        /// </summary>
        [Output("lastActivationTime")]
        public Output<string> LastActivationTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the application package.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The storage URL at which the application package is stored.
        /// </summary>
        [Output("storageUrl")]
        public Output<string> StorageUrl { get; private set; } = null!;

        /// <summary>
        /// The UTC time at which the storage URL will expire.
        /// </summary>
        [Output("storageUrlExpiry")]
        public Output<string> StorageUrlExpiry { get; private set; } = null!;

        /// <summary>
        /// The version of the application package. 
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationPackage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationPackage(string name, ApplicationPackageArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:batch/v20170101:ApplicationPackage", name, args ?? new ApplicationPackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationPackage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:batch/v20170101:ApplicationPackage", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:batch/latest:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20151201:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20170501:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20170901:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20181201:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20190401:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20190801:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20200301:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20200501:ApplicationPackage"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20200901:ApplicationPackage"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationPackage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationPackage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApplicationPackage(name, id, options);
        }
    }

    public sealed class ApplicationPackageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The ID of the application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The version of the application.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ApplicationPackageArgs()
        {
        }
    }
}
