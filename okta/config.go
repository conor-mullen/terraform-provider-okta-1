package okta

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/oktadeveloper/terraform-provider-okta/sdk"
)

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "Okta Terraform Provider")
	return adt.T.RoundTrip(req)
}

type (
	// AddHeaderTransport used to tack on default headers to outgoing requests
	AddHeaderTransport struct {
		T http.RoundTripper
	}

	// Config contains our provider schema values and Okta clients
	Config struct {
		orgName          string
		domain           string
		apiToken         string
		retryCount       int
		parallelism      int
		backoff          bool
		minWait          int
		maxWait          int
		logLevel         int
		requestTimeout   int
		oktaClient       *okta.Client
		supplementClient *sdk.ApiSupplement
		logger           hclog.Logger
	}
)

func (c *Config) loadAndValidate() error {
	c.logger = hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Level(c.logLevel),
		TimeFormat: "2006/01/02 03:04:05",
	})
	var httpClient *http.Client
	if c.backoff {
		retryableClient := retryablehttp.NewClient()
		retryableClient.RetryWaitMin = time.Second * time.Duration(c.minWait)
		retryableClient.RetryWaitMax = time.Second * time.Duration(c.maxWait)
		retryableClient.RetryMax = c.retryCount
		retryableClient.Logger = c.logger
		retryableClient.HTTPClient.Transport = logging.NewTransport("Okta", retryableClient.HTTPClient.Transport)
		retryableClient.ErrorHandler = errHandler
		httpClient = retryableClient.StandardClient()
	} else {
		httpClient = cleanhttp.DefaultClient()
		httpClient.Transport = logging.NewTransport("Okta", httpClient.Transport)
	}

	_, client, err := okta.NewClient(
		context.Background(),
		okta.WithOrgUrl(fmt.Sprintf("https://%v.%v", c.orgName, c.domain)),
		okta.WithToken(c.apiToken),
		okta.WithCache(false),
		okta.WithHttpClient(*httpClient),
		okta.WithRateLimitMaxBackOff(int64(c.maxWait)),
		okta.WithRequestTimeout(int64(c.requestTimeout)),
		okta.WithRateLimitMaxRetries(int32(c.retryCount)),
		okta.WithUserAgentExtra("okta-terraform/3.7.4"),
	)
	if err != nil {
		return err
	}
	c.oktaClient = client
	c.supplementClient = &sdk.ApiSupplement{
		BaseURL:         fmt.Sprintf("https://%s.%s", c.orgName, c.domain),
		Client:          httpClient,
		Token:           c.apiToken,
		RequestExecutor: client.GetRequestExecutor(),
	}
	return nil
}

func errHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	err = okta.CheckResponseForError(resp)
	if err != nil {
		oErr, ok := err.(*okta.Error)
		if ok {
			oErr.ErrorSummary = fmt.Sprintf("%s, giving up after %d attempt(s)", oErr.ErrorSummary, numTries)
			return resp, oErr
		}
		return resp, fmt.Errorf("%v: giving up after %d attempt(s)", err, numTries)
	}
	return resp, nil
}
