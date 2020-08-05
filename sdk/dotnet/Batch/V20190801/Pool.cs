// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20190801
{
    /// <summary>
    /// Contains information about a pool.
    /// </summary>
    public partial class Pool : Pulumi.CustomResource
    {
        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties associated with the pool.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PoolPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Pool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pool(string name, PoolArgs args, CustomResourceOptions? options = null)
            : base("azurerm:batch/v20190801:Pool", name, args ?? new PoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:batch/v20190801:Pool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Pool(name, id, options);
        }
    }

    public sealed class PoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("applicationLicenses")]
        private InputList<string>? _applicationLicenses;

        /// <summary>
        /// The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
        /// </summary>
        public InputList<string> ApplicationLicenses
        {
            get => _applicationLicenses ?? (_applicationLicenses = new InputList<string>());
            set => _applicationLicenses = value;
        }

        [Input("applicationPackages")]
        private InputList<Inputs.ApplicationPackageReferenceArgs>? _applicationPackages;

        /// <summary>
        /// Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
        /// </summary>
        public InputList<Inputs.ApplicationPackageReferenceArgs> ApplicationPackages
        {
            get => _applicationPackages ?? (_applicationPackages = new InputList<Inputs.ApplicationPackageReferenceArgs>());
            set => _applicationPackages = value;
        }

        [Input("certificates")]
        private InputList<Inputs.CertificateReferenceArgs>? _certificates;

        /// <summary>
        /// For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
        /// </summary>
        public InputList<Inputs.CertificateReferenceArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.CertificateReferenceArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
        /// </summary>
        [Input("deploymentConfiguration")]
        public Input<Inputs.DeploymentConfigurationArgs>? DeploymentConfiguration { get; set; }

        /// <summary>
        /// The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
        /// </summary>
        [Input("interNodeCommunication")]
        public Input<string>? InterNodeCommunication { get; set; }

        /// <summary>
        /// The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
        /// </summary>
        [Input("maxTasksPerNode")]
        public Input<int>? MaxTasksPerNode { get; set; }

        [Input("metadata")]
        private InputList<Inputs.MetadataItemArgs>? _metadata;

        /// <summary>
        /// The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
        /// </summary>
        public InputList<Inputs.MetadataItemArgs> Metadata
        {
            get => _metadata ?? (_metadata = new InputList<Inputs.MetadataItemArgs>());
            set => _metadata = value;
        }

        [Input("mountConfiguration")]
        private InputList<Inputs.MountConfigurationArgs>? _mountConfiguration;

        /// <summary>
        /// This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
        /// </summary>
        public InputList<Inputs.MountConfigurationArgs> MountConfiguration
        {
            get => _mountConfiguration ?? (_mountConfiguration = new InputList<Inputs.MountConfigurationArgs>());
            set => _mountConfiguration = value;
        }

        /// <summary>
        /// The pool name. This must be unique within the account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The network configuration for a pool.
        /// </summary>
        [Input("networkConfiguration")]
        public Input<Inputs.NetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
        /// </summary>
        [Input("scaleSettings")]
        public Input<Inputs.ScaleSettingsArgs>? ScaleSettings { get; set; }

        /// <summary>
        /// In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
        /// </summary>
        [Input("startTask")]
        public Input<Inputs.StartTaskArgs>? StartTask { get; set; }

        /// <summary>
        /// If not specified, the default is spread.
        /// </summary>
        [Input("taskSchedulingPolicy")]
        public Input<Inputs.TaskSchedulingPolicyArgs>? TaskSchedulingPolicy { get; set; }

        [Input("userAccounts")]
        private InputList<Inputs.UserAccountArgs>? _userAccounts;
        public InputList<Inputs.UserAccountArgs> UserAccounts
        {
            get => _userAccounts ?? (_userAccounts = new InputList<Inputs.UserAccountArgs>());
            set => _userAccounts = value;
        }

        /// <summary>
        /// For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        public PoolArgs()
        {
        }
    }
}
