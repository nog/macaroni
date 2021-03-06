package macaroni

import (
	"encoding/json"

	"github.com/Songmu/horenso"
)

var testReport horenso.Report
var testReportJSON = []byte(`{
  "command": "perl -E 'say 1;warn \"$$\\n\";'",
  "commandArgs": [
    "perl",
    "-E",
    "say 1;warn \"$$\\n\";"
  ],
  "output": "1\n95030\n",
  "stdout": "1\n",
  "stderr": "95030\n",
  "exitCode": 0,
  "result": "command exited with code: 0",
  "pid": 95030,
  "startAt": "2015-12-28T00:37:10.494282399+09:00",
  "endAt": "2015-12-28T00:37:10.546466379+09:00",
  "hostname": "webserver.example.com",
  "systemTime": 0.034632,
  "userTime": 0.026523
}`)

func init() {
	json.Unmarshal(testReportJSON, &testReport)
}
