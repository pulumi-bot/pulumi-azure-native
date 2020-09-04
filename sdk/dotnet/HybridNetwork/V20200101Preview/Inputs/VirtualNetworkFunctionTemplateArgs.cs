// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridNetwork.V20200101Preview.Inputs
{

    /// <summary>
    /// The virtual network function template.
    /// </summary>
    public sealed class VirtualNetworkFunctionTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("virutalNetworkFunctionRoleConfigurations")]
        private InputList<Inputs.VirtualNetworkFunctionRoleConfigurationArgs>? _virutalNetworkFunctionRoleConfigurations;

        /// <summary>
        /// An array of virtual network function role definitions.
        /// </summary>
        public InputList<Inputs.VirtualNetworkFunctionRoleConfigurationArgs> VirutalNetworkFunctionRoleConfigurations
        {
            get => _virutalNetworkFunctionRoleConfigurations ?? (_virutalNetworkFunctionRoleConfigurations = new InputList<Inputs.VirtualNetworkFunctionRoleConfigurationArgs>());
            set => _virutalNetworkFunctionRoleConfigurations = value;
        }

        public VirtualNetworkFunctionTemplateArgs()
        {
        }
    }
}
