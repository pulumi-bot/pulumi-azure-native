// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Attestation.V20180901Preview.Inputs
{

    public sealed class JSONWebKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The "alg" (algorithm) parameter identifies the algorithm intended for
        /// use with the key.  The values used should either be registered in the
        /// IANA "JSON Web Signature and Encryption Algorithms" registry
        /// established by [JWA] or be a value that contains a Collision-
        /// Resistant Name.
        /// </summary>
        [Input("alg", required: true)]
        public Input<string> Alg { get; set; } = null!;

        /// <summary>
        /// The "crv" (curve) parameter identifies the curve type
        /// </summary>
        [Input("crv")]
        public Input<string>? Crv { get; set; }

        /// <summary>
        /// RSA private exponent or ECC private key
        /// </summary>
        [Input("d")]
        public Input<string>? D { get; set; }

        /// <summary>
        /// RSA Private Key Parameter
        /// </summary>
        [Input("dp")]
        public Input<string>? Dp { get; set; }

        /// <summary>
        /// RSA Private Key Parameter
        /// </summary>
        [Input("dq")]
        public Input<string>? Dq { get; set; }

        /// <summary>
        /// RSA public exponent, in Base64
        /// </summary>
        [Input("e")]
        public Input<string>? E { get; set; }

        /// <summary>
        /// Symmetric key
        /// </summary>
        [Input("k")]
        public Input<string>? K { get; set; }

        /// <summary>
        /// The "kid" (key ID) parameter is used to match a specific key.  This
        /// is used, for instance, to choose among a set of keys within a JWK Set
        /// during key rollover.  The structure of the "kid" value is
        /// unspecified.  When "kid" values are used within a JWK Set, different
        /// keys within the JWK Set SHOULD use distinct "kid" values.  (One
        /// example in which different keys might use the same "kid" value is if
        /// they have different "kty" (key type) values but are considered to be
        /// equivalent alternatives by the application using them.)  The "kid"
        /// value is a case-sensitive string.
        /// </summary>
        [Input("kid", required: true)]
        public Input<string> Kid { get; set; } = null!;

        /// <summary>
        /// The "kty" (key type) parameter identifies the cryptographic algorithm
        /// family used with the key, such as "RSA" or "EC". "kty" values should
        /// either be registered in the IANA "JSON Web Key Types" registry
        /// established by [JWA] or be a value that contains a Collision-
        /// Resistant Name.  The "kty" value is a case-sensitive string.
        /// </summary>
        [Input("kty", required: true)]
        public Input<string> Kty { get; set; } = null!;

        /// <summary>
        /// RSA modulus, in Base64
        /// </summary>
        [Input("n")]
        public Input<string>? N { get; set; }

        /// <summary>
        /// RSA secret prime
        /// </summary>
        [Input("p")]
        public Input<string>? P { get; set; }

        /// <summary>
        /// RSA secret prime, with p &lt; q
        /// </summary>
        [Input("q")]
        public Input<string>? Q { get; set; }

        /// <summary>
        /// RSA Private Key Parameter
        /// </summary>
        [Input("qi")]
        public Input<string>? Qi { get; set; }

        /// <summary>
        /// Use ("public key use") identifies the intended use of
        /// the public key. The "use" parameter is employed to indicate whether
        /// a public key is used for encrypting data or verifying the signature
        /// on data. Values are commonly "sig" (signature) or "enc" (encryption).
        /// </summary>
        [Input("use", required: true)]
        public Input<string> Use { get; set; } = null!;

        /// <summary>
        /// X coordinate for the Elliptic Curve point
        /// </summary>
        [Input("x")]
        public Input<string>? X { get; set; }

        [Input("x5c")]
        private InputList<string>? _x5c;

        /// <summary>
        /// The "x5c" (X.509 certificate chain) parameter contains a chain of one
        /// or more PKIX certificates [RFC5280].  The certificate chain is
        /// represented as a JSON array of certificate value strings.  Each
        /// string in the array is a base64-encoded (Section 4 of [RFC4648] --
        /// not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value.
        /// The PKIX certificate containing the key value MUST be the first
        /// certificate.
        /// </summary>
        public InputList<string> X5c
        {
            get => _x5c ?? (_x5c = new InputList<string>());
            set => _x5c = value;
        }

        /// <summary>
        /// Y coordinate for the Elliptic Curve point
        /// </summary>
        [Input("y")]
        public Input<string>? Y { get; set; }

        public JSONWebKeyArgs()
        {
        }
    }
}
