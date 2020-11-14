// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DesktopVirtualization.V20201019Preview.Inputs
{

    /// <summary>
    /// Schema for MSIX Package Dependencies properties.
    /// </summary>
    public sealed class MsixPackageDependenciesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of package dependency.
        /// </summary>
        [Input("dependencyName")]
        public Input<string>? DependencyName { get; set; }

        /// <summary>
        /// Dependency version required.
        /// </summary>
        [Input("minVersion")]
        public Input<string>? MinVersion { get; set; }

        /// <summary>
        /// Name of dependency publisher.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        public MsixPackageDependenciesArgs()
        {
        }
    }
}
