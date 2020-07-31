// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Batch.V20181201.Inputs
{

    public sealed class CloudServiceConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Possible values are: 2 - OS Family 2, equivalent to Windows Server 2008 R2 SP1. 3 - OS Family 3, equivalent to Windows Server 2012. 4 - OS Family 4, equivalent to Windows Server 2012 R2. 5 - OS Family 5, equivalent to Windows Server 2016. For more information, see Azure Guest OS Releases (https://azure.microsoft.com/documentation/articles/cloud-services-guestos-update-matrix/#releases).
        /// </summary>
        [Input("osFamily", required: true)]
        public Input<string> OsFamily { get; set; } = null!;

        /// <summary>
        /// The default value is * which specifies the latest operating system version for the specified OS family.
        /// </summary>
        [Input("osVersion")]
        public Input<string>? OsVersion { get; set; }

        public CloudServiceConfigurationArgs()
        {
        }
    }
}