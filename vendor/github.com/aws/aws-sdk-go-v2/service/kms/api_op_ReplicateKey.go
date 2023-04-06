// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replicates a multi-Region key into the specified Region. This operation creates
// a multi-Region replica key based on a multi-Region primary key in a different
// Region of the same Amazon Web Services partition. You can create multiple
// replicas of a primary key, but each must be in a different Region. To create a
// multi-Region primary key, use the CreateKey operation. This operation supports
// multi-Region keys, an KMS feature that lets you create multiple interoperable
// KMS keys in different Amazon Web Services Regions. Because these KMS keys have
// the same key ID, key material, and other metadata, you can use them
// interchangeably to encrypt data in one Amazon Web Services Region and decrypt it
// in a different Amazon Web Services Region without re-encrypting the data or
// making a cross-Region call. For more information about multi-Region keys, see
// Multi-Region keys in KMS
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
// in the Key Management Service Developer Guide. A replica key is a
// fully-functional KMS key that can be used independently of its primary and peer
// replica keys. A primary key and its replica keys share properties that make them
// interoperable. They have the same key ID
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id)
// and key material. They also have the same key spec
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-spec),
// key usage
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-usage),
// key material origin
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-origin),
// and automatic key rotation status
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html). KMS
// automatically synchronizes these shared properties among related multi-Region
// keys. All other properties of a replica key can differ, including its key policy
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html), tags
// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html),
// aliases (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html),
// and Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html). KMS
// pricing and quotas for KMS keys apply to each primary key and replica key. When
// this operation completes, the new replica key has a transient key state of
// Creating. This key state changes to Enabled (or PendingImport) after a few
// seconds when the process of creating the new replica key is complete. While the
// key state is Creating, you can manage key, but you cannot yet use it in
// cryptographic operations. If you are creating and using the replica key
// programmatically, retry on KMSInvalidStateException or call DescribeKey to check
// its KeyState value before using it. For details about the Creating key state,
// see Key states of KMS keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// Key Management Service Developer Guide. You cannot create more than one replica
// of a primary key in any Region. If the Region already includes a replica of the
// key you're trying to replicate, ReplicateKey returns an AlreadyExistsException
// error. If the key state of the existing replica is PendingDeletion, you can
// cancel the scheduled key deletion (CancelKeyDeletion) or wait for the key to be
// deleted. The new replica key you create will have the same shared properties
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-sync-properties)
// as the original replica key. The CloudTrail log of a ReplicateKey operation
// records a ReplicateKey operation in the primary key's Region and a CreateKey
// operation in the replica key's Region. If you replicate a multi-Region primary
// key with imported key material, the replica key is created with no key material.
// You must import the same key material that you imported into the primary key.
// For details, see Importing key material into multi-Region keys in the Key
// Management Service Developer Guide. To convert a replica key to a primary key,
// use the UpdatePrimaryRegion operation. ReplicateKey uses different default
// values for the KeyPolicy and Tags parameters than those used in the KMS console.
// For details, see the parameter descriptions. Cross-account use: No. You cannot
// use this operation to create a replica key in a different Amazon Web Services
// account. Required permissions:
//
// * kms:ReplicateKey on the primary key (in the
// primary key's Region). Include this permission in the primary key's key
// policy.
//
// * kms:CreateKey in an IAM policy in the replica Region.
//
// * To use the
// Tags parameter, kms:TagResource in an IAM policy in the replica Region.
//
// Related
// operations
//
// * CreateKey
//
// * UpdatePrimaryRegion
func (c *Client) ReplicateKey(ctx context.Context, params *ReplicateKeyInput, optFns ...func(*Options)) (*ReplicateKeyOutput, error) {
	if params == nil {
		params = &ReplicateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplicateKey", params, optFns, c.addOperationReplicateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplicateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplicateKeyInput struct {

	// Identifies the multi-Region primary key that is being replicated. To determine
	// whether a KMS key is a multi-Region primary key, use the DescribeKey operation
	// to check the value of the MultiRegionKeyType property. Specify the key ID or key
	// ARN of a multi-Region primary key. For example:
	//
	// * Key ID:
	// mrk-1234abcd12ab34cd56ef1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab
	//
	// To
	// get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// The Region ID of the Amazon Web Services Region for this replica key. Enter the
	// Region ID, such as us-east-1 or ap-southeast-2. For a list of Amazon Web
	// Services Regions in which KMS is supported, see KMS service endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/kms.html#kms_region) in the
	// Amazon Web Services General Reference. HMAC KMS keys are not supported in all
	// Amazon Web Services Regions. If you try to replicate an HMAC KMS key in an
	// Amazon Web Services Region in which HMAC keys are not supported, the
	// ReplicateKey operation returns an UnsupportedOperationException. For a list of
	// Regions in which HMAC KMS keys are supported, see HMAC keys in KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html) in the Key
	// Management Service Developer Guide. The replica must be in a different Amazon
	// Web Services Region than its primary key and other replicas of that primary key,
	// but in the same Amazon Web Services partition. KMS must be available in the
	// replica Region. If the Region is not enabled by default, the Amazon Web Services
	// account must be enabled in the Region. For information about Amazon Web Services
	// partitions, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference. For information about enabling and
	// disabling Regions, see Enabling a Region
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable)
	// and Disabling a Region
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-disable)
	// in the Amazon Web Services General Reference.
	//
	// This member is required.
	ReplicaRegion *string

	// Skips ("bypasses") the key policy lockout safety check. The default value is
	// false. Setting this value to true increases the risk that the KMS key becomes
	// unmanageable. Do not set this value to true indiscriminately. For more
	// information, see Default key policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key)
	// in the Key Management Service Developer Guide. Use this parameter only when you
	// intend to prevent the principal that is making the request from making a
	// subsequent PutKeyPolicy request on the KMS key.
	BypassPolicyLockoutSafetyCheck bool

	// A description of the KMS key. The default value is an empty string (no
	// description). The description is not a shared property of multi-Region keys. You
	// can specify the same description or a different description for each key in a
	// set of related multi-Region keys. KMS does not synchronize this property.
	Description *string

	// The key policy to attach to the KMS key. This parameter is optional. If you do
	// not provide a key policy, KMS attaches the default key policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)
	// to the KMS key. The key policy is not a shared property of multi-Region keys.
	// You can specify the same key policy or a different key policy for each key in a
	// set of related multi-Region keys. KMS does not synchronize this property. If you
	// provide a key policy, it must meet the following criteria:
	//
	// * The key policy
	// must allow the calling principal to make a subsequent PutKeyPolicy request on
	// the KMS key. This reduces the risk that the KMS key becomes unmanageable. For
	// more information, see Default key policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key)
	// in the Key Management Service Developer Guide. (To omit this condition, set
	// BypassPolicyLockoutSafetyCheck to true.)
	//
	// * Each statement in the key policy
	// must contain one or more principals. The principals in the key policy must exist
	// and be visible to KMS. When you create a new Amazon Web Services principal, you
	// might need to enforce a delay before including the new principal in a key policy
	// because the new principal might not be immediately visible to KMS. For more
	// information, see Changes that I make are not always immediately visible
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	// in the Amazon Web Services Identity and Access Management User Guide.
	//
	// A key
	// policy document can include only the following characters:
	//
	// * Printable ASCII
	// characters from the space character (\u0020) through the end of the ASCII
	// character range.
	//
	// * Printable characters in the Basic Latin and Latin-1
	// Supplement character set (through \u00FF).
	//
	// * The tab (\u0009), line feed
	// (\u000A), and carriage return (\u000D) special characters
	//
	// For information about
	// key policies, see Key policies in KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the
	// Key Management Service Developer Guide. For help writing and formatting a JSON
	// policy document, see the IAM JSON Policy Reference
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in
	// the Identity and Access Management User Guide .
	Policy *string

	// Assigns one or more tags to the replica key. Use this parameter to tag the KMS
	// key when it is created. To tag an existing KMS key, use the TagResource
	// operation. Tagging or untagging a KMS key can allow or deny permission to the
	// KMS key. For details, see ABAC for KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key
	// Management Service Developer Guide. To use this parameter, you must have
	// kms:TagResource
	// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// permission in an IAM policy. Tags are not a shared property of multi-Region
	// keys. You can specify the same tags or different tags for each key in a set of
	// related multi-Region keys. KMS does not synchronize this property. Each tag
	// consists of a tag key and a tag value. Both the tag key and the tag value are
	// required, but the tag value can be an empty (null) string. You cannot have more
	// than one tag on a KMS key with the same tag key. If you specify an existing tag
	// key with a different tag value, KMS replaces the current tag value with the
	// specified one. When you add tags to an Amazon Web Services resource, Amazon Web
	// Services generates a cost allocation report with usage and costs aggregated by
	// tags. Tags can also be used to control access to a KMS key. For details, see
	// Tagging Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ReplicateKeyOutput struct {

	// Displays details about the new replica key, including its Amazon Resource Name
	// (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// and Key states of KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html). It also
	// includes the ARN and Amazon Web Services Region of its primary key and other
	// replica keys.
	ReplicaKeyMetadata *types.KeyMetadata

	// The key policy of the new replica key. The value is a key policy document in
	// JSON format.
	ReplicaPolicy *string

	// The tags on the new replica key. The value is a list of tag key and tag value
	// pairs.
	ReplicaTags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplicateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReplicateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReplicateKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpReplicateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplicateKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opReplicateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ReplicateKey",
	}
}
