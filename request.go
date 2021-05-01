package faceeasco

import "github.com/imroc/req"

func post() {
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic YWRtaW46YWRtaW4=",
	}

	param := req.Param{
		"name": "imroc",
		"cmd":  "add",
	}

	r := req.New()
	r.Post("http://foo.bar/api", header, param)
}
