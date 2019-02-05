package HelloWorld

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"gopkg.in/resty.v1"
	"os"
)

var log = logger.GetLogger("activity-helloworld")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	name := context.GetInput("name").(string)
	salutation := context.GetInput("salutation").(string)
	log.Infof("The Flogo engine says [%s] to [%s]", salutation, name)
	context.SetOutput("result", "The Flogo engine says "+salutation+" to "+name)


	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"token": os.Getenv("SLACK_XOXP_TOKEN"),
			"pretty": "1",
		}).
		SetHeader("Accept", "application/x-www-form-urlencoded").
		Get("https://slack.com/api/users.list")
	if err != nil {
		log.Infof("Error on response.\n[ERRO] -", err)
	}

	log.Infof("LA RESPONSE SFVC", resp)

	return true, nil
}
