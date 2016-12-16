// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/query"
)

// Amazon Simple Notification Service (Amazon SNS) is a web service that enables
// you to build distributed web-enabled applications. Applications can use Amazon
// SNS to easily push real-time notification messages to interested subscribers
// over multiple delivery protocols. For more information about this product
// see http://aws.amazon.com/sns (http://aws.amazon.com/sns/). For detailed
// information about Amazon SNS features and their associated API calls, see
// the Amazon SNS Developer Guide (http://docs.aws.amazon.com/sns/latest/dg/).
//
// We also provide SDKs that enable you to access Amazon SNS from your preferred
// programming language. The SDKs contain functionality that automatically takes
// care of tasks such as: cryptographically signing your service requests, retrying
// requests, and handling error responses. For a list of available SDKs, go
// to Tools for Amazon Web Services (http://aws.amazon.com/tools/).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type SNS struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "sns"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the SNS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a SNS client from just a session.
//     svc := sns.New(mySession)
//
//     // Create a SNS client with additional configuration
//     svc := sns.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *SNS {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *SNS {
	svc := &SNS{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-03-31",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a SNS operation and runs any
// custom request initialization.
func (c *SNS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
