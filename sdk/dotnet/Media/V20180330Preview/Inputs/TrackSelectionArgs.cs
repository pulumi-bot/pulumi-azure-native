// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Inputs
{

    /// <summary>
    /// Class to select a track
    /// </summary>
    public sealed class TrackSelectionArgs : Pulumi.ResourceArgs
    {
        [Input("trackSelections")]
        private InputList<Inputs.TrackPropertyConditionArgs>? _trackSelections;

        /// <summary>
        /// TrackSelections is a track property condition list which can specify track(s)
        /// </summary>
        public InputList<Inputs.TrackPropertyConditionArgs> TrackSelections
        {
            get => _trackSelections ?? (_trackSelections = new InputList<Inputs.TrackPropertyConditionArgs>());
            set => _trackSelections = value;
        }

        public TrackSelectionArgs()
        {
        }
    }
}
