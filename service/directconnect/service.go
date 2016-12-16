// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package directconnect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Direct Connect links your internal network to an AWS Direct Connect location
// over a standard 1 gigabit or 10 gigabit Ethernet fiber-optic cable. One end
// of the cable is connected to your router, the other to an AWS Direct Connect
// router. With this connection in place, you can create virtual interfaces
// directly to the AWS cloud (for example, to Amazon Elastic Compute Cloud (Amazon
// EC2) and Amazon Simple Storage Service (Amazon S3)) and to Amazon Virtual
// Private Cloud (Amazon VPC), bypassing Internet service providers in your
// network path. An AWS Direct Connect location provides access to AWS in the
// region it is associated with, as well as access to other US regions. For
// example, you can provision a single connection to any AWS Direct Connect
// location in the US and use it to access public AWS services in all US Regions
// and AWS GovCloud (US).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type DirectConnect struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "directconnect" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName     // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the DirectConnect client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a DirectConnect client from just a session.
//     svc := directconnect.New(mySession)
//
//     // Create a DirectConnect client with additional configuration
//     svc := directconnect.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DirectConnect {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DirectConnect {
	svc := &DirectConnect{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-10-25",
				JSONVersion:   "1.1",
				TargetPrefix:  "OvertureService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DirectConnect operation and runs any
// custom request initialization.
func (c *DirectConnect) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
