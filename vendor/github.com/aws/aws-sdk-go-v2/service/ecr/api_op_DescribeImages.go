// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeImagesInput struct {
	_ struct{} `type:"structure"`

	// The filter key and value with which to filter your DescribeImages results.
	Filter *DescribeImagesFilter `locationName:"filter" type:"structure"`

	// The list of image IDs for the requested repository.
	ImageIds []ImageIdentifier `locationName:"imageIds" min:"1" type:"list"`

	// The maximum number of repository results returned by DescribeImages in paginated
	// output. When this parameter is used, DescribeImages only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeImages
	// request with the returned nextToken value. This value can be between 1 and
	// 1000. If this parameter is not used, then DescribeImages returns up to 100
	// results and a nextToken value, if applicable. This option cannot be used
	// when you specify images with imageIds.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated DescribeImages request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This option cannot be used when you specify images with imageIds.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to describe images. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository that contains the images to describe.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeImagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeImagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeImagesInput"}
	if s.ImageIds != nil && len(s.ImageIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageIds", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageIds != nil {
		for i, v := range s.ImageIds {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ImageIds", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeImagesOutput struct {
	_ struct{} `type:"structure"`

	// A list of ImageDetail objects that contain data about the image.
	ImageDetails []ImageDetail `locationName:"imageDetails" type:"list"`

	// The nextToken value to include in a future DescribeImages request. When the
	// results of a DescribeImages request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeImagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeImages = "DescribeImages"

// DescribeImagesRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Returns metadata about the images in a repository, including image size,
// image tags, and creation date.
//
// Beginning with Docker version 1.9, the Docker client compresses image layers
// before pushing them to a V2 Docker registry. The output of the docker images
// command shows the uncompressed image size, so it may return a larger image
// size than the image sizes returned by DescribeImages.
//
//    // Example sending a request using DescribeImagesRequest.
//    req := client.DescribeImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/DescribeImages
func (c *Client) DescribeImagesRequest(input *DescribeImagesInput) DescribeImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeImages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeImagesInput{}
	}

	req := c.newRequest(op, input, &DescribeImagesOutput{})
	return DescribeImagesRequest{Request: req, Input: input, Copy: c.DescribeImagesRequest}
}

// DescribeImagesRequest is the request type for the
// DescribeImages API operation.
type DescribeImagesRequest struct {
	*aws.Request
	Input *DescribeImagesInput
	Copy  func(*DescribeImagesInput) DescribeImagesRequest
}

// Send marshals and sends the DescribeImages API request.
func (r DescribeImagesRequest) Send(ctx context.Context) (*DescribeImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImagesResponse{
		DescribeImagesOutput: r.Request.Data.(*DescribeImagesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeImagesRequestPaginator returns a paginator for DescribeImages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeImagesRequest(input)
//   p := ecr.NewDescribeImagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeImagesPaginator(req DescribeImagesRequest) DescribeImagesPaginator {
	return DescribeImagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeImagesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeImagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeImagesPaginator struct {
	aws.Pager
}

func (p *DescribeImagesPaginator) CurrentPage() *DescribeImagesOutput {
	return p.Pager.CurrentPage().(*DescribeImagesOutput)
}

// DescribeImagesResponse is the response type for the
// DescribeImages API operation.
type DescribeImagesResponse struct {
	*DescribeImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImages request.
func (r *DescribeImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}