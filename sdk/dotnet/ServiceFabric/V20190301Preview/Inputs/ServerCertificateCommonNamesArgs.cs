// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Inputs
{

    /// <summary>
    /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
    /// </summary>
    public sealed class ServerCertificateCommonNamesArgs : Pulumi.ResourceArgs
    {
        [Input("commonNames")]
        private InputList<Inputs.ServerCertificateCommonNameArgs>? _commonNames;

        /// <summary>
        /// The list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        public InputList<Inputs.ServerCertificateCommonNameArgs> CommonNames
        {
            get => _commonNames ?? (_commonNames = new InputList<Inputs.ServerCertificateCommonNameArgs>());
            set => _commonNames = value;
        }

        /// <summary>
        /// The local certificate store location.
        /// </summary>
        [Input("x509StoreName")]
        public Input<string>? X509StoreName { get; set; }

        public ServerCertificateCommonNamesArgs()
        {
        }
    }
}
