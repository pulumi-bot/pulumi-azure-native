// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301.Inputs
{

    /// <summary>
    /// Defines a map that contains specific application health policies for different applications.
    /// Each entry specifies as key the application name and as value an ApplicationHealthPolicy used to evaluate the application health.
    /// The application name should include the 'fabric:' URI scheme.
    /// The map is empty by default.
    /// </summary>
    public sealed class ApplicationHealthPolicyMapArgs : Pulumi.ResourceArgs
    {
        public ApplicationHealthPolicyMapArgs()
        {
        }
    }
}