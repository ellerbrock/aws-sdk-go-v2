// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// SFN provides the API operation methods for making requests to
// AWS Step Functions. See this package's package overview docs
// for details on the service.
//
// SFN methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type SFN struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*SFN)

// Used for custom request initialization logic
var initRequest func(*SFN, *aws.Request)

// Service information constants
const (
	ServiceName = "states"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the SFN client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a SFN client from just a config.
//     svc := sfn.New(myConfig)
//
//     // Create a SFN client with additional configuration
//     svc := sfn.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *SFN {
	var signingName string
	signingRegion := config.Region

	svc := &SFN{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-11-23",
				JSONVersion:   "1.0",
				TargetPrefix:  "AWSStepFunctions",
			},
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
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a SFN operation and runs any
// custom request initialization.
func (c *SFN) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
