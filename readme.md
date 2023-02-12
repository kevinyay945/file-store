# File Store

- partially compatible with imgur
- Support Google Drive / PCloud for asset storage

## Api

### [POST] /v1/image
#### Description
Upload Asset and get preview link
#### Header
- Authorization: Client-ID {your-client-id}
#### Body
- Type: form-data
- Key/Type
  - image/File
#### Response
##### Success
- Status: 200
- Data:

```json
{
  "data": {
    "link": "{preview link}"
  }
}
```
##### Fail
- Status: 400
- Data:
```
{error message}
```

### [GET] /v1/temp-link/obsidian/:fileName
#### Description
Redirect link to storage space public link
It doesn't redirect if server can't find the public link
#### Response
##### Success
- Status: 302
