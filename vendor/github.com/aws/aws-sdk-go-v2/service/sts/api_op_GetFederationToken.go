// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of temporary security credentials (consisting of an access key
// ID, a secret access key, and a security token) for a user. A typical use is in a
// proxy application that gets temporary security credentials on behalf of
// distributed applications inside a corporate network.
//
// You must call the GetFederationToken operation using the long-term security
// credentials of an IAM user. As a result, this call is appropriate in contexts
// where those credentials can be safeguarded, usually in a server-based
// application. For a comparison of GetFederationToken with the other API
// operations that produce temporary credentials, see [Requesting Temporary Security Credentials]and [Comparing the Amazon Web Services STS API operations] in the IAM User Guide.
//
// Although it is possible to call GetFederationToken using the security
// credentials of an Amazon Web Services account root user rather than an IAM user
// that you create for the purpose of a proxy application, we do not recommend it.
// For more information, see [Safeguard your root user credentials and don't use them for everyday tasks]in the IAM User Guide.
//
// You can create a mobile-based or browser-based app that can authenticate users
// using a web identity provider like Login with Amazon, Facebook, Google, or an
// OpenID Connect-compatible identity provider. In this case, we recommend that you
// use [Amazon Cognito]or AssumeRoleWithWebIdentity . For more information, see [Federation Through a Web-based Identity Provider] in the IAM User
// Guide.
//
// # Session duration
//
// The temporary credentials are valid for the specified duration, from 900
// seconds (15 minutes) up to a maximum of 129,600 seconds (36 hours). The default
// session duration is 43,200 seconds (12 hours). Temporary credentials obtained by
// using the root user credentials have a maximum duration of 3,600 seconds (1
// hour).
//
// # Permissions
//
// You can use the temporary credentials created by GetFederationToken in any
// Amazon Web Services service with the following exceptions:
//
//   - You cannot call any IAM operations using the CLI or the Amazon Web Services
//     API. This limitation does not apply to console sessions.
//
//   - You cannot call any STS operations except GetCallerIdentity .
//
// You can use temporary credentials for single sign-on (SSO) to the console.
//
// You must pass an inline or managed [session policy] to this operation. You can pass a single
// JSON policy document to use as an inline session policy. You can also specify up
// to 10 managed policy Amazon Resource Names (ARNs) to use as managed session
// policies. The plaintext that you use for both inline and managed session
// policies can't exceed 2,048 characters.
//
// Though the session policy parameters are optional, if you do not pass a policy,
// then the resulting federated user session has no permissions. When you pass
// session policies, the session permissions are the intersection of the IAM user
// policies and the session policies that you pass. This gives you a way to further
// restrict the permissions for a federated user. You cannot use session policies
// to grant more permissions than those that are defined in the permissions policy
// of the IAM user. For more information, see [Session Policies]in the IAM User Guide. For
// information about using GetFederationToken to create temporary security
// credentials, see [GetFederationToken—Federation Through a Custom Identity Broker].
//
// You can use the credentials to access a resource that has a resource-based
// policy. If that policy specifically references the federated user session in the
// Principal element of the policy, the session has the permissions allowed by the
// policy. These permissions are granted in addition to the permissions granted by
// the session policies.
//
// # Tags
//
// (Optional) You can pass tag key-value pairs to your session. These are called
// session tags. For more information about session tags, see [Passing Session Tags in STS]in the IAM User
// Guide.
//
// You can create a mobile-based or browser-based app that can authenticate users
// using a web identity provider like Login with Amazon, Facebook, Google, or an
// OpenID Connect-compatible identity provider. In this case, we recommend that you
// use [Amazon Cognito]or AssumeRoleWithWebIdentity . For more information, see [Federation Through a Web-based Identity Provider] in the IAM User
// Guide.
//
// An administrator must grant you the permissions necessary to pass session tags.
// The administrator can also create granular permissions to allow you to pass only
// specific session tags. For more information, see [Tutorial: Using Tags for Attribute-Based Access Control]in the IAM User Guide.
//
// Tag key–value pairs are not case sensitive, but case is preserved. This means
// that you cannot have separate Department and department tag keys. Assume that
// the user that you are federating has the Department = Marketing tag and you
// pass the department = engineering session tag. Department and department are
// not saved as separate tags, and the session tag passed in the request takes
// precedence over the user tag.
//
// [Federation Through a Web-based Identity Provider]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity
// [session policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Amazon Cognito]: http://aws.amazon.com/cognito/
// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
// [GetFederationToken—Federation Through a Custom Identity Broker]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getfederationtoken
// [Comparing the Amazon Web Services STS API operations]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison
// [Safeguard your root user credentials and don't use them for everyday tasks]: https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#lock-away-credentials
// [Requesting Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html
// [Tutorial: Using Tags for Attribute-Based Access Control]: https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_attribute-based-access-control.html
func (c *Client) GetFederationToken(ctx context.Context, params *GetFederationTokenInput, optFns ...func(*Options)) (*GetFederationTokenOutput, error) {
	if params == nil {
		params = &GetFederationTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFederationToken", params, optFns, c.addOperationGetFederationTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFederationTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFederationTokenInput struct {

	// The name of the federated user. The name is used as an identifier for the
	// temporary security credentials (such as Bob ). For example, you can reference
	// the federated user name in a resource-based policy, such as in an Amazon S3
	// bucket policy.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can also
	// include underscores or any of the following characters: =,.@-
	//
	// This member is required.
	Name *string

	// The duration, in seconds, that the session should last. Acceptable durations
	// for federation sessions range from 900 seconds (15 minutes) to 129,600 seconds
	// (36 hours), with 43,200 seconds (12 hours) as the default. Sessions obtained
	// using root user credentials are restricted to a maximum of 3,600 seconds (one
	// hour). If the specified duration is longer than one hour, the session obtained
	// by using root user credentials defaults to one hour.
	DurationSeconds *int32

	// An IAM policy in JSON format that you want to use as an inline session policy.
	//
	// You must pass an inline or managed [session policy] to this operation. You can pass a single
	// JSON policy document to use as an inline session policy. You can also specify up
	// to 10 managed policy Amazon Resource Names (ARNs) to use as managed session
	// policies.
	//
	// This parameter is optional. However, if you do not pass any session policies,
	// then the resulting federated user session has no permissions.
	//
	// When you pass session policies, the session permissions are the intersection of
	// the IAM user policies and the session policies that you pass. This gives you a
	// way to further restrict the permissions for a federated user. You cannot use
	// session policies to grant more permissions than those that are defined in the
	// permissions policy of the IAM user. For more information, see [Session Policies]in the IAM User
	// Guide.
	//
	// The resulting credentials can be used to access a resource that has a
	// resource-based policy. If that policy specifically references the federated user
	// session in the Principal element of the policy, the session has the permissions
	// allowed by the policy. These permissions are granted in addition to the
	// permissions that are granted by the session policies.
	//
	// The plaintext that you use for both inline and managed session policies can't
	// exceed 2,048 characters. The JSON policy characters can be any ASCII character
	// from the space character to the end of the valid character list (\u0020 through
	// \u00FF). It can also include the tab (\u0009), linefeed (\u000A), and carriage
	// return (\u000D) characters.
	//
	// An Amazon Web Services conversion compresses the passed inline session policy,
	// managed policy ARNs, and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates
	// by percentage how close the policies and tags for your request are to the upper
	// size limit.
	//
	// [session policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	Policy *string

	// The Amazon Resource Names (ARNs) of the IAM managed policies that you want to
	// use as a managed session policy. The policies must exist in the same account as
	// the IAM user that is requesting federated access.
	//
	// You must pass an inline or managed [session policy] to this operation. You can pass a single
	// JSON policy document to use as an inline session policy. You can also specify up
	// to 10 managed policy Amazon Resource Names (ARNs) to use as managed session
	// policies. The plaintext that you use for both inline and managed session
	// policies can't exceed 2,048 characters. You can provide up to 10 managed policy
	// ARNs. For more information about ARNs, see [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]in the Amazon Web Services General
	// Reference.
	//
	// This parameter is optional. However, if you do not pass any session policies,
	// then the resulting federated user session has no permissions.
	//
	// When you pass session policies, the session permissions are the intersection of
	// the IAM user policies and the session policies that you pass. This gives you a
	// way to further restrict the permissions for a federated user. You cannot use
	// session policies to grant more permissions than those that are defined in the
	// permissions policy of the IAM user. For more information, see [Session Policies]in the IAM User
	// Guide.
	//
	// The resulting credentials can be used to access a resource that has a
	// resource-based policy. If that policy specifically references the federated user
	// session in the Principal element of the policy, the session has the permissions
	// allowed by the policy. These permissions are granted in addition to the
	// permissions that are granted by the session policies.
	//
	// An Amazon Web Services conversion compresses the passed inline session policy,
	// managed policy ARNs, and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates
	// by percentage how close the policies and tags for your request are to the upper
	// size limit.
	//
	// [session policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	// [Session Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session
	// [Amazon Resource Names (ARNs) and Amazon Web Services Service Namespaces]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	PolicyArns []types.PolicyDescriptorType

	// A list of session tags. Each session tag consists of a key name and an
	// associated value. For more information about session tags, see [Passing Session Tags in STS]in the IAM User
	// Guide.
	//
	// This parameter is optional. You can pass up to 50 session tags. The plaintext
	// session tag keys can’t exceed 128 characters and the values can’t exceed 256
	// characters. For these and additional limits, see [IAM and STS Character Limits]in the IAM User Guide.
	//
	// An Amazon Web Services conversion compresses the passed inline session policy,
	// managed policy ARNs, and session tags into a packed binary format that has a
	// separate limit. Your request can fail for this limit even if your plaintext
	// meets the other requirements. The PackedPolicySize response element indicates
	// by percentage how close the policies and tags for your request are to the upper
	// size limit.
	//
	// You can pass a session tag with the same key as a tag that is already attached
	// to the user you are federating. When you do, session tags override a user tag
	// with the same key.
	//
	// Tag key–value pairs are not case sensitive, but case is preserved. This means
	// that you cannot have separate Department and department tag keys. Assume that
	// the role has the Department = Marketing tag and you pass the department =
	// engineering session tag. Department and department are not saved as separate
	// tags, and the session tag passed in the request takes precedence over the role
	// tag.
	//
	// [Passing Session Tags in STS]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
	// [IAM and STS Character Limits]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful GetFederationToken request, including temporary Amazon Web
// Services credentials that can be used to make Amazon Web Services requests.
type GetFederationTokenOutput struct {

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// The size of the security token that STS API operations return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size.
	Credentials *types.Credentials

	// Identifiers for the federated user associated with the credentials (such as
	// arn:aws:sts::123456789012:federated-user/Bob or 123456789012:Bob ). You can use
	// the federated user's ARN in your resource-based policies, such as an Amazon S3
	// bucket policy.
	FederatedUser *types.FederatedUser

	// A percentage value that indicates the packed size of the session policies and
	// session tags combined passed in the request. The request fails if the packed
	// size is greater than 100 percent, which means the policies and tags exceeded the
	// allowed space.
	PackedPolicySize *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFederationTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetFederationToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetFederationToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFederationToken"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpGetFederationTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFederationToken(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetFederationToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFederationToken",
	}
}
