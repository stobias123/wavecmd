# Wavecmd

I want to create simple events from the cli, to overlay over my dashboards..

## Usage

**Config**
Set your token and address in `~/.wavecmd.yaml`
```yaml
address: my-wavefront-tenant.wavefront.com
token: foo-bar-long-api-token
```

**Create an event**
```bash
$ wavecmd event create -n "testing" -t "deployment" -m "testing some cli" | jq .

{
  "annotations": {
    "details": "testing more cli",
    "severity": "",
    "type": "deployment"
  },
  "name": "demo-for-readme",
  "id": "123123123123:demo-for-readme:0",
  "startTime": 1592505935000,
  "tags": [
    "cli"
  ],
  "Severity": "",
  "Type": "deployment",
  "Details": "testing more cli",
  "isEphemeral": false
}
```

**Close an event**
```bash
$ wavecmd event close --id "123123123123:demo-for-readme:0" | jq .
{
  "annotations": {
    "details": "testing more cli",
    "severity": "",
    "type": "deployment"
  },
  "name": "demo-for-readme",
  "id": "123123123:demo-for-readme:1",
  "startTime": 1592505935000,
  "endTime": 1592505984609,
  "tags": [
    "cli"
  ],
  "Severity": "",
  "Type": "deployment",
  "Details": "testing more cli",
  "isEphemeral": false
}
```