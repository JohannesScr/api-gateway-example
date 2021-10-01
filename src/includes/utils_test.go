package includes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	type data struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	tt := []struct {
		name string
		body interface{}
		data data
	}{
		{
			name: "no body",
			body: nil,
			data: data{
				FirstName: "",
				LastName:  "",
			},
		},
		{
			name: "additional body fields",
			body: struct {
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Age       int    `json:"age"`
			}{
				FirstName: "james",
				Age:       5,
			},
			data: data{
				FirstName: "james",
				LastName:  "",
			},
		},
		{
			name: "missing body fields",
			body: data{
				FirstName: "james",
			},
			data: data{
				FirstName: "james",
				LastName:  "",
			},
		},
		{
			name: "correct",
			body: data{
				FirstName: "james",
				LastName:  "bond",
			},
			data: data{
				FirstName: "james",
				LastName:  "bond",
			},
		},
	}

	for _, tc := range tt {
		// run the actual sub-test
		t.Run(tc.name, func(t *testing.T) {
			bs, _ := json.Marshal(tc.body)
			br := bytes.NewReader(bs)
			req := httptest.NewRequest("POST", "/", br)
			rec := httptest.NewRecorder()

			// test struct
			ts := data{}
			err := Decode(rec, req, &ts)
			if err != nil {
				t.Errorf("unable to decode %v", tc.body)
			}
			if ts.FirstName != tc.data.FirstName {
				t.Errorf("expected '%v' got '%v'", tc.data.FirstName, ts.FirstName)
			}
			if ts.LastName != tc.data.LastName {
				t.Errorf("expected '%v' got '%v'", tc.data.LastName, ts.LastName)
			}
		})
	}
}

func TestQueryParams(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	q := url.Values{}
	q.Add("userId", "12345678")
	q.Add("perm", "1")
	q.Add("perm", "2")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()

	qp, _ := QueryParams(rec, req)
	if qp["userId"] == nil {
		t.Errorf("expected %v got %v", []string{"12345678"}, nil)
	}
	if qp["userId"][0] != "12345678" {
		t.Errorf("expected %v got %v", "12345678", qp["userId"][0])
	}
	if len(qp["userId"]) != 1 {
		t.Errorf("expected %v got %v", 1, len(qp["userId"]))
	}
	if qp["perm"] == nil {
		t.Errorf("expected %v got %v", []string{"1", "2"}, nil)
	}
	if len(qp["perm"]) != 2 {
		t.Errorf("expected %v got %v", 2, len(qp["perm"]))
	}
	if qp["perm"][0] != "1" {
		t.Errorf("expected '%v' got '%v'", "1", qp["perm"][0])
	}
}

func TestResp_Respond(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	type Payload = struct {
		String string            `json:"string"`
		Slice  []int             `json:"slice"`
		Map    map[string]string `json:"map"`
		Struct interface{}       `json:"struct"`
	}

	payload := Payload{
		String: "MyString",
		Slice:  []int{1, 2, 3},
		Map: map[string]string{
			"uno":  "one",
			"dos":  "two",
			"tres": "three",
		},
		Struct: struct {
			String string `json:"string"`
			Bool   bool   `json:"bool"`
		}{String: "StructString", Bool: true},
	}


	resp := Resp{
		Status:  123,
		Message: "My Test Message",
		Data: payload,
	}
	resp.Respond(rec, req)
	res := rec.Result()

	resBody := Resp{}
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unable to read body: %v", err)
	}
	err = json.Unmarshal(bs, &resBody)
	if err != nil {
		t.Errorf("unable to decode body: %v", err)
	}

	if res.StatusCode != resp.Status {
		t.Errorf("expected %v got %v", resp.Status, res.StatusCode)
	}
	if resBody.Status != 0 {
		t.Errorf("expected %v got %v", 0, resBody.Status)
	}
}

func TestEnv_Load(t *testing.T) {
	e := Env{}
	if e.Vars != nil {
		t.Errorf("expected nil got %v", e.Vars)
	}

	e.Load("./../../testdata/.env.testfile")

	myEnvVar := os.Getenv("MY_ENV_VAR")
	myInt := os.Getenv("MY_INT")
	if myEnvVar != "this is my env var" {
		t.Errorf("expected '%v' got '%v'", "this is my env var", myEnvVar)
	}
	if myInt != "1234" {
		t.Errorf("expected '%v' got '%v'", "1234", myInt)
	}

	if e.Vars["MY_ENV_VAR"] != "this is my env var" {
		t.Errorf("expected '%v' got '%v'", "this is my env var", e.Vars["MY_ENV_VAR"])
	}
	if e.Vars["MY_INT"] != "1234" {
		t.Errorf("expected '%v' got '%v'", "1234", e.Vars["MY_INT"])
	}
}