package HelloWorld

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"io/ioutil"
	"net/http"
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

	req, _ := http.NewRequest("GET", "https://slack.com/api/users.list", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("token", "xoxp-539270626848-539900207523-540778240739-edfae5fc0b4dda7d30e144e104be3c5d")
	q.Add("pretty", "1")
	req.URL.RawQuery = q.Encode()

	log.Infof("La request es [%s]", req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Infof("Error on response.\n[ERRO] -", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Infof("Response from slack [%s]", string([]byte(body)))

	return true, nil
}
