FORMAT: 1A
HOST: https://api.hetzner.cloud/v1

# Hetzner Cloud API

This is the official API documentation for the Public Hetzner Cloud.

# Group Overview


# Overview [/]

## Introduction [GET /overview/intro]

The Hetzner Cloud API operates over HTTPS and uses JSON as its data format. The API is a RESTful
API and utilizes HTTP methods and HTTP status codes to specify requests and responses.


As an alternative to working directly with our API you may also consider to use:

* Our CLI program [hcloud](https://github.com/hetznercloud/cli)

* Our [library for Go](https://github.com/hetznercloud/hcloud-go)

Also you can find a [list of libraries, tools, and integrations on GitHub.](https://github.com/hetznercloud/awesome-hcloud)

If you are developing integrations based on our API and your product is Open Source you may be eligible for a free one time €50 (excl. VAT) credit on your account. Please contact us via the the support page on your Cloud Console and let us know the following:

* The type of integration you would like to develop
* Link to the GitHub repo you will use for the project
* Link to some other Open Source work you have already done (if you have done so)

+ Response 204 (application/dummy)
## Getting Started [GET /overview/gettingstarted]

>#### Example: Getting Started Request
```bash
curl -H "Authorization: Bearer jEheVytlAoFl7F8MqUQ7jAo2hOXASztX" \
    https://api.hetzner.cloud/v1/servers
```

>#### Example: Getting Started Response
```json
{
    "servers": [],
    "meta": {
        "pagination": {
            "page": 1,
            "per_page": 25,
            "previous_page": null,
            "next_page": null,
            "last_page": 1,
            "total_entries": 0
        }
    }
}
```

To get started using the API you first need an API token. Sign in into the
[Hetzner Cloud Console](https://console.hetzner.cloud/) choose a project, go to `Access` → `Tokens`, and create a new token. Make
sure to copy the token because it won’t be shown to you again.
A token is bound to a project, to interact with the API of another project you have to create a new token inside the project.
Let’s say your new token is `jEheVytlAoFl7F8MqUQ7jAo2hOXASztX`.

You’re now ready to do your first request against the API. To get a list
of all servers in your project, issue the example request on the right side using [curl](https://curl.haxx.se/).

Make sure to replace the token in the example command with the token you have
just created. Since your project probably does not contain any servers yet,
the example response will look like the response on the right side.
We will almost always provide a resource root like `servers` inside the example response.
A response can also contain a `meta` object with information like [Pagination](#overview-pagination).

+ Response 204 (application/dummy)
## Authentication [GET /overview/auth]

>#### Example: Authorization header
```bash
Authorization: Bearer LRK9DAWQ1ZAEFSrCNEEzLCUwhYX1U3g7wMg4dTlkkDC96fyDuyJ39nVbVjCKSDfj
```


All requests to the Hetzner Cloud API must be authenticated via a API token. Include your
secret API token in every request you send to the API with the `Authorization` HTTP header.

```
Authorization: Bearer <token>
```

To create a new API token for your project, switch into the [Hetzner Cloud Console](https://console.hetzner.cloud/) choose a project, go to `Access` → `Tokens`, and create a new token.

+ Response 204 (application/dummy)
## Errors [GET /overview/errors]

>#### Example: Error response
```json
{
  "error": {
    "code": "invalid_input",
    "message": "invalid input in field 'broken_field': is too long",
    "details": {
      "fields": [
        {
          "name": "broken_field",
          "messages": ["is too long"]
        }
      ]
    }
  }
}
```

Errors are indicated by HTTP status codes. Further, the response of the
request which generated the error contains an error code, an error message, and,
optionally, error details. The schema of the error details object depends on
the error code.

The error response contains the following keys:

| Key       | Meaning                                                               |
|-----------|-----------------------------------------------------------------------|
| `code`    | Short string indicating the type of error (machine-parsable)          |
| `message` | Textual description on what has gone wrong                            |
| `details` | An object providing for details on the error (schema depends on code) |

### Error Codes

| Code                      | Description                                                                                                    |
|---------------------------|----------------------------------------------------------------------------------------------------------------|
| `forbidden`               | Insufficient permissions for this request                                                                      |
| `invalid_input`           | Error while parsing or processing the input                                                                    |
| `json_error`              | Invalid JSON input in your request                                                                             |
| `locked`                  | The item you are trying to access is locked (there is already an action running)                               |
| `not_found`               | Entity not found                                                                                               |
| `rate_limit_exceeded`     | Error when sending too many requests                                                                           |
| `resource_limit_exceeded` | Error when exceeding the maximum quantity of a resource for an account                                         |
| `resource_unavailable`    | The requested resource is currently unavailable                                                                |
| `service_error`           | Error within a service                                                                                         |
| `uniqueness_error`        | One or more of the objects fields must be unique                                                               |
| `protected`               | The action you are trying to start is protected for this resource                                              |
| `maintenance`             | Cannot perform operation due to maintenance                                                                    |

### Error Details

**invalid_input**

```json
{
  "error": {
    "code": "invalid_input",
    "message": "invalid input in field 'broken_field': is too long",
    "details": {
      "fields": [
        {
          "name": "broken_field",
          "messages": ["is too long"]
        }
      ]
    }
  }
}
```

**uniqueness_error**

```json
{
  "error": {
    "code": "uniqueness_error",
    "message": "SSH key with the same fingerprint already exists",
    "details": {
      "fields": [
        {
          "name": "public_key",
        }
      ]
    }
  }
}
```

+ Response 204 (application/dummy)
## Pagination [GET /overview]

Responses which return multiple items support pagination. If they do support pagination,
it can be controlled with following query string parameters:

* A `page` parameter specifies the page to fetch. The number of the first page is 1.
* A `per_page` parameter specifies the number of items returned per page. The default
  value is 25, the maximum value is 50 except otherwise specified in the documentation.


>#### Example: Pagination Link header
```bash
Link: <https://api.hetzner.cloud/actions?page=2&per_page=5>; rel="prev",
      <https://api.hetzner.cloud/actions?page=4&per_page=5>; rel="next",
      <https://api.hetzner.cloud/actions?page=6&per_page=5>; rel="last"
```

>Line breaks have been added for display purposes only and responses may only contain
>some of the above `rel` values.


Responses contain a `Link` header with pagination information.

Additionally, if the response body is JSON and the root object is an object, that object
has a `pagination` object inside the `meta` object with pagination information:

```json
{
    "servers": [...],
    "meta": {
        "pagination": {
            "page": 2,
            "per_page": 25,
            "previous_page": 1,
            "next_page": 3,
            "last_page": 4,
            "total_entries": 100
        }
    }
}
```

The keys `previous_page`, `next_page`, `last_page`, and `total_entries` may be `null` when
on the first page, last page, or when the total number of entries is unknown.

+ Response 204 (application/dummy)
## Sorting [GET /overview/sorting]

>#### Example: Sorting
```bash
https://api.hetzner.cloud/actions?sort=status
https://api.hetzner.cloud/actions?sort=status:asc
https://api.hetzner.cloud/actions?sort=status:desc
https://api.hetzner.cloud/actions?sort=status:asc&sort=command:desc
```

Some responses which return multiple items support sorting. If they do support
sorting the documentation states which fields can be used for sorting. You
specify sorting with the `sort` query string parameter. You can sort by
multiple fields. You can set the sort direction by appending `:asc` or `:desc`
to the field name. By default, ascending sorting is used.

+ Response 204 (application/dummy)
## Rate Limiting [GET /overview/ratelimit]

All requests, whether they are authenticated or not, are subject to rate
limiting. If you have reached your limit, your requests will be handled with
a `429 Too Many Requests` error. Burst requests are allowed. Responses contain
serveral headers which provide information about your current rate limit status.

*  The `RateLimit-Limit` header contains the total number of requests you
   can perform per hour.

*  The `RateLimit-Remaining` header contains the number of requests remaining
   in the current rate limit time frame.

*  The `RateLimit-Reset` header contains a UNIX timestamp of the point in time
   when your rate limit will have recovered and you will have the full number
   of requests available again.

The default limit is 3600 requests per hour and per project. The number of remaining
requests increases gradually. For example, when your limit is 3600 requests per hour,
the number of remaining requests will increase by 1 every second.

+ Response 204 (application/dummy)
## Changelog [GET /overview/changelog]

**2018-04-05:** Add deprecated flag to images

Property `deprecated` has been added to images. It's `null` or an ISO8601 timestamp which when set indicates the time after which it will no longer be possible to use the image to create a new server. This applies only to system images. Images that have deprecated timestamp in the past are also not returned under [Get all images endpoint](#images-get-all-images), they however can still be queried individually via their ID and the [Get single image endpoint](#images-get-an-image).

<br>

**2018-03-28:** Add resource protection

Servers, snapshots and Floating IPs can now be protected from accidental deletion and rebuild via the protection configuration. Additionally, this introduced the new [Image Actions](#image-actions) resource.

<br>

**2018-03-20:** Filter SSH keys by fingerprint

Similar to filter by name, SSH keys can now be filtered by fingerprint.

<br>

**2018-03-16:** Use SSH key name in server creation

In addition to SSH key IDs, SSH key names can now be used for injecting an SSH key on server create.

<br>

**2018-02-26:** Add deprecated flag to ISOs

Property `deprecated` has been added to ISOs. It's `null` or an ISO8601 timestamp which when set indicates the time after which it will no longer be possible to attach that ISO to servers. ISOs that have deprecated timestamp in the past are also not returned under [get all ISOs endpoint](#isos-get-all-isos), they however can still be queried individually via their ID and the [get single ISO endpoint](#isos-get-an-iso).

<br>

**2018-02-21:** Add server action to request console


Request credentials for remote access via VNC over websocket was added to the API. For more information see the [documentation for the request console endpoint](#server-actions-request-console-for-a-server)

+ Response 204 (application/dummy)

# Group Resources

# Actions [/actions]

Actions show the results and progress of asynchronous requests to the API. 

## List all Actions [GET /actions{?status,sort}]

Returns all action objects. You can select specific actions only and sort the results by using URI parameters.

+ Parameters
    + status (enum[string], optional) - Can be used multiple times. Response will have only actions with specified statuses.
        + Members
            + `running`
            + `success`
            + `error`
    + sort (enum[string], optional) - Can be used multiple times.
        + Members
            + `id`
            + `id:asc`
            + `id:desc`
            + `command`
            + `command:asc`
            + `command:desc`
            + `status`
            + `status:asc`
            + `status:desc`
            + `progress`
            + `progress:asc`
            + `progress:desc`
            + `started`
            + `started:asc`
            + `started:desc`
            + `finished`
            + `finished:asc`
            + `finished:desc`
+ Response 200 (application/json)
    The `actions` key in the reply contains an array of action objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/actions
     ```



    + Attributes
        + `actions` (array[Action], required, fixed-type)
    + Body

## Get one Action [GET /actions/{id}]

Returns a specific action object.

+ Parameters
    + id: `1337` (string) - ID of the action

+ Response 200 (application/json)
    The `action` key in the reply has this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/actions/1337
     ```



    + Attributes
        + `action` (Action, required)
    + Body
# Servers [/servers]

Servers are virtual machines that can be provisioned.

## Get all Servers [GET /servers{?name}]

Returns all existing server objects.

+ Parameters
    + name (string, optional) - Can be used to filter servers by their name. The response will only contain the server matching the specified name.
+ Response 200 (application/json)
    The `servers` key in the reply contains an array of server objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of server |
    | name | string | Name of the server (must be unique per project and a valid hostname as per RFC 1123) |
    | status | string | Status of the server <br>Choices: `running`, `initializing`, `starting`, `stopping`, `off`, `deleting`, `migrating`, `rebuilding`, `unknown` |
    | created | string | Point in time when the server was created (in ISO-8601 format) |
    | public_net | object | Public network information. The servers ipv4 address can be found in `public_net->ipv4->ip` |
    | server_type | object | Type of server - determines how much ram, disk and cpu a server has |
    | datacenter | object | Datacenter this server is located at |
    | image | object,&nbsp;null | Image this server was created from. |
    | iso | object,&nbsp;null | ISO image that is attached to this server. Null if no ISO is attached. |
    | rescue_enabled | boolean | True if rescue mode is enabled: Server will then boot into rescue system on next reboot. |
    | locked | boolean | True if server has been locked and is not available to user. |
    | backup_window | string,&nbsp;null | Time window (UTC) in which the backup will run, or null if the backups are not enabled |
    | outgoing_traffic | number,&nbsp;null | Outbound Traffic for the current billing period in bytes |
    | ingoing_traffic | number,&nbsp;null | Inbound Traffic for the current billing period in bytes |
    | included_traffic | number | Free Traffic for the current billing period in bytes |
    | protection | object | Protection configuration for the server |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers
     ```



    + Attributes
        + `servers` (array[Server], required, fixed-type)
    + Body

## Get a Server [GET /servers/{id}]

Returns a specific server object. The server must exist inside the project.


+ Parameters
    + id: `42` (string) - ID of the server

+ Response 200 (application/json)
    The `server` key in the reply contains a server object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of server |
    | name | string | Name of the server (must be unique per project and a valid hostname as per RFC 1123) |
    | status | string | Status of the server <br>Choices: `running`, `initializing`, `starting`, `stopping`, `off`, `deleting`, `migrating`, `rebuilding`, `unknown` |
    | created | string | Point in time when the server was created (in ISO-8601 format) |
    | public_net | object | Public network information. The servers ipv4 address can be found in `public_net->ipv4->ip` |
    | server_type | object | Type of server - determines how much ram, disk and cpu a server has |
    | datacenter | object | Datacenter this server is located at |
    | image | object,&nbsp;null | Image this server was created from. |
    | iso | object,&nbsp;null | ISO image that is attached to this server. Null if no ISO is attached. |
    | rescue_enabled | boolean | True if rescue mode is enabled: Server will then boot into rescue system on next reboot. |
    | locked | boolean | True if server has been locked and is not available to user. |
    | backup_window | string,&nbsp;null | Time window (UTC) in which the backup will run, or null if the backups are not enabled |
    | outgoing_traffic | number,&nbsp;null | Outbound Traffic for the current billing period in bytes |
    | ingoing_traffic | number,&nbsp;null | Inbound Traffic for the current billing period in bytes |
    | included_traffic | number | Free Traffic for the current billing period in bytes |
    | protection | object | Protection configuration for the server |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42
     ```



    + Attributes
        + `server` (Server, required)
    + Body

## Create a Server [POST /servers]

Creates a new server. Returns preliminary information about the server as well as an action that covers progress of creation.

+ Request (application/json)
    Please note that server names must be unique per project and valid hostnames as per RFC 1123
    (i.e. may only contain letters, digits, periods, and dashes).
    
    For `server_type` you can either use the ID as listed in [`/server_types`](#server-types-get-all-server-types) or its name.
    
    For `image` you can either use the ID as listed in [`/images`](#images-get-all-images) or its name.
    
    If you want to create the server in a location, you must set `location` to the ID or name as listed in [`/locations`](#locations-get-all-locations). This is the recommended way. You can be even more specific by setting `datacenter` to the ID or name as listed in [`/datacenters`](#datacenters-get-all-datacenters). However directly specifying the datacenter is discouraged since supply availability in datacenters varies greatly and datacenters may be out of stock for extended periods of time or not serve certain server types at all.
    
    For accessing your server we strongly recommend to use SSH keys by passing the respective key ids in `ssh_keys`. If you do not specify any `ssh_keys` we will generate a root password for you and
    return it in the response.
    
    
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | name | string&nbsp;(required) | Name of the server to create (must be unique per project and a valid hostname as per RFC 1123) |
    | server_type | string&nbsp;(required) | ID or name of the server type this server should be created with |
    | start_after_create | boolean&nbsp;(optional) | Start Server right after creation. Defaults to true. |
    | image | string&nbsp;(required) | ID or name of the image the server is created from |
    | ssh_keys | array&nbsp;(optional) | SSH key IDs or names which should be injected into the server at creation time |
    | user_data | string&nbsp;(optional) | Cloud-Init user data to use during server creation. This field is limited to 32KiB. |
    | location | string&nbsp;(optional) | ID or name of location to create server in. |
    | datacenter | string&nbsp;(optional) | ID or name of datacenter to create server in. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"name": "my-server", "server_type": "cx11", "location": "nbg1", "start_after_create": true, "image": "ubuntu-16.04", "ssh_keys": ["my-ssh-key"], "user_data": "#cloud-config\nruncmd:\n- [touch, /root/cloud-init-worked]\n"}' \
  https://api.hetzner.cloud/v1/servers
     ```



    + Attributes(Server create)

+ Response 201 (application/json)
    The `server` key in the reply contains a server object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of server |
    | name | string | Name of the server (must be unique per project and a valid hostname as per RFC 1123) |
    | status | string | Status of the server <br>Choices: `initializing`, `starting`, `running`, `stopping`, `off`, `deleting`, `migrating`, `rebuilding` |
    | created | string | Point in time when the server was created (in ISO-8601 format) |
    | public_net | object | Public network information. The servers ipv4 address can be found in `public_net->ipv4->ip` |
    | server_type | object | Type of server - determines how much ram, disk and cpu a server has |
    | datacenter | object | Datacenter this server is located at |
    | image | object,&nbsp;null | Image this server was created from. |
    | iso | object,&nbsp;null | ISO image that is attached to this server. Null if no ISO is attached. |
    | rescue_enabled | boolean | True if rescue mode is enabled: Server will then boot into rescue system on next reboot. |
    | locked | boolean | True if server has been locked and is not available to user. |
    | backup_window | string,&nbsp;null | Time window (UTC) in which the backup will run, or null if the backups are not enabled |
    | outgoing_traffic | number,&nbsp;null | Outbound Traffic for the current billing period in bytes |
    | ingoing_traffic | number,&nbsp;null | Inbound Traffic for the current billing period in bytes |
    | included_traffic | number | Free Traffic for the current billing period in bytes |
    | protection | object | Protection configuration for the server |




    
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    #### Call specific error codes
    
    | Code | Description |
    |------|-------------|
    | `placement_error` | An error during placement of the server occured |
    
    + Attributes
        + `server` (Server, required)
            + `status`: `initializing` (enum[string], required) - Status of the server
                + Members
                    + `initializing`
                    + `starting`
                    + `running`
                    + `stopping`
                    + `off`
                    + `deleting`
                    + `migrating`
                    + `rebuilding`
        + `action` (ActionRunning, required)
            + `command`: `create_server` (string)
        + `root_password`: `YItygq1v3GYjjMomLaKc` (string, required, nullable) - Root password when no SSH keys have been specified

    + Body

## Change Name of a Server [PUT /servers/{id}]

Changes the name of a server.

Please note that server names must be unique per project and valid hostnames as per RFC 1123
(i.e. may only contain letters, digits, periods, and dashes).

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | name | string&nbsp;(optional) | New name to set |

    > #### Example curl
     ```bash
     curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"name": "new-name"}' \
  https://api.hetzner.cloud/v1/servers/42
     ```




    + Attributes
        + name: `new-name` (string) - New name to set

+ Response 200 (application/json)
    The `server` key in the reply contains the modified server object with the new name.
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of server |
    | name | string | Server name |
    | status | string | Status of the server <br>Choices: `running`, `initializing`, `starting`, `stopping`, `off`, `deleting`, `migrating`, `rebuilding`, `unknown` |
    | created | string | Point in time when the server was created (in ISO-8601 format) |
    | public_net | object | Public network information. The servers ipv4 address can be found in `public_net->ipv4->ip` |
    | server_type | object | Type of server - determines how much ram, disk and cpu a server has |
    | datacenter | object | Datacenter this server is located at |
    | image | object,&nbsp;null | Image this server was created from. |
    | iso | object,&nbsp;null | ISO image that is attached to this server. Null if no ISO is attached. |
    | rescue_enabled | boolean | True if rescue mode is enabled: Server will then boot into rescue system on next reboot. |
    | locked | boolean | True if server has been locked and is not available to user. |
    | backup_window | string,&nbsp;null | Time window (UTC) in which the backup will run, or null if the backups are not enabled |
    | outgoing_traffic | number,&nbsp;null | Outbound Traffic for the current billing period in bytes |
    | ingoing_traffic | number,&nbsp;null | Inbound Traffic for the current billing period in bytes |
    | included_traffic | number | Free Traffic for the current billing period in bytes |
    | protection | object | Protection configuration for the server |




    + Attributes
        + `server` (Server, required)
            + name: `new-name` (string) - Server name
    + Body

## Delete a Server [DELETE /servers/{id}]

Deletes a server. This immediately removes the server from your account, and it is no longer accessible.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 200 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42
     ```



    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `delete_server` (string)

## Get Metrics for a Server [GET /servers/{id}/metrics{?type,start,end,step}]

Get Metrics for specified server.

You must specify the type of metric to get: `cpu`, `disk` or `network`. You can also specify more than one type by comma separation, e.g. `cpu,disk`.

Depending on the type you will get different time series data:

|Type | Timeseries | Unit | Description |
|---- |------------|------|-------------|
|cpu|cpu | percent | Percent CPU usage |
|disk|disk.0.iops.read | iop/s | Number of read IO operations per second|
||disk.0.iops.write | iop/s | Number of write IO operations per second|
||disk.0.bandwidth.read |bytes/s |Bytes read per second|
||disk.0.bandwidth.write | bytes/s | Bytes written per second |
|network|network.0.pps.<span>in</span>| packets/s | Public Network interface packets per second received|
||network.0.pps.out | packets/s | Public Network interface packets per second sent|
||network.0.bandwidth.<span>in</span> | bytes/s | Public Network interface bytes/s received|
||network.0.bandwidth.out | bytes/s | Public Network interface bytes/s sent|


Metrics are available for the last 30 days only.

If you do not provide the step argument we will automatically adjust it so that a maximum of 100 samples are returned.

We limit the number of samples returned to a maximum of 500 and will adjust the step parameter accordingly.

 + Parameters
     + id: `42` (string) - ID of the server
     + type: `cpu` (enum[string], required) - Type of metrics to get
       + `cpu`
       + `disk`
       + `network`
     + start: `2017-01-01T00:00:00Z` (string, required) - Start of period to get Metrics for (in ISO-8601 format)
     + end: `2017-01-01T23:00:00Z` (string, required) - End of period to get Metrics for (in ISO-8601 format)
     + step: `60` (number, optional) - Resolution of results in seconds

+ Response 200 (application/json)
    The `metrics` key in the reply contains a metrics object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | start | string | Start of period of metrics reported (in ISO-8601 format) |
    | end | string | End of period of metrics reported (in ISO-8601 format) |
    | step | number | Resolution of results in seconds. |
    | time_series | object | Hash with timeseries information, containing the name of timeseries as key |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/metrics?type=cpu&start=2017-01-01T00:00:00Z&end=2017-01-01T23:00:00Z
     ```



    
    The `time_series` key in the returned metrics object is a hash with one key per series returned with this structure:
    
    ```
    "time_series": {
        "name_of_timeseries": {
            "values": [
              [
                1435781470.622,
                "42"
              ],
              [
                1435781471.622,
                "43"
              ]
            ]
        }
    }
    ```
    
    Contrary to the example on the right the timestamp is a number and not a string.

    + Attributes
        + `metrics` (Metrics Generic, required)


# Server Actions [/servers/{id}/actions]

## Get all Actions for a Server [GET /servers/{id}/actions{?status,sort}]

Returns all action objects for a server.

+ Parameters
    + id: `42` (string) - ID of the server
    + status (enum[string], optional) - Can be used multiple times. Response will have only actions with specified statuses.
        + Members
            + `running`
            + `success`
            + `error`
    + sort (enum[string], optional) - Can be used multiple times.
        + Members
            + `id`
            + `id:asc`
            + `id:desc`
            + `command`
            + `command:asc`
            + `command:desc`
            + `status`
            + `status:asc`
            + `status:desc`
            + `progress`
            + `progress:asc`
            + `progress:desc`
            + `started`
            + `started:asc`
            + `started:desc`
            + `finished`
            + `finished:asc`
            + `finished:desc`

+ Response 200 (application/json)
    The `actions` key in the reply contains an array of action objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions
     ```



    + Attributes
        + `actions` (array[Action], required, fixed-type)
    + Body

## Get a specific Action for a Server [GET /servers/{id}/actions/{action_id}]

Returns a specific action object for a Server.

+ Parameters
    + id: `42` (string) - ID of the server
    + action_id: `1337` (string) - ID of the action

+ Response 200 (application/json)
    The `action` key in the reply has this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/1337
     ```



    + Attributes
        + `action` (Action, required)
    + Body

## Power on a Server [POST /servers/{id}/actions/poweron]

Starts a server by turning its power on.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/poweron
     ```



        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `start_server` (string)
    + Body

## Soft-reboot a Server [POST /servers/{id}/actions/reboot]

Reboots a server gracefully by sending an ACPI request. The server operating system must
support ACPI and react to the request, otherwise the server will not reboot.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/reboot
     ```



        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `reboot_server` (string)
    + Body

## Reset a Server [POST /servers/{id}/actions/reset]

Cuts power to a server and starts it again. This forcefully stops it without giving the
server operating system time to gracefully stop. This may lead to data loss, it's equivalent to
pulling the power cord and plugging it in again. Reset should only be used when reboot
does not work.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/reset
     ```



        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `reset_server` (string)
    + Body

## Shutdown a Server [POST /servers/{id}/actions/shutdown]

Shuts down a server gracefully by sending an ACPI shutdown request. The server operating system must
support ACPI and react to the request, otherwise the server will not shut down.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/shutdown
     ```



        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `shutdown_server` (string)
    + Body

## Power off a Server [POST /servers/{id}/actions/poweroff]

Cuts power to the server. This forcefully stops it without giving the server operating
system time to gracefully stop. May lead to data loss, equivalent to pulling the power
cord. Power off should only be used when shutdown does not work.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/poweroff
     ```



        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `stop_server` (string)
    + Body

## Reset root Password of a Server [POST /servers/{id}/actions/reset_password]

Resets the root password. Only works for Linux systems that are running the qemu guest agent. Server must be powered on (state `on`) in order for this operation to succeed.

This will generate a new password for this server and return it.

If this does not succeed you can use the rescue system to netboot the server and manually change your server password by hand.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `root_password` key in the reply contains the new root password that will be active if the action succeeds.

    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/reset_password
     ```



        
    + Attributes
        + `root_password`: `zCWbFhnu950dUTko5f40` (string) - Password that will be set for this server once the action succeeds.
        + `action` (ActionRunning, required)
            + `command`: `reset_password` (string)
    + Body

## Enable Rescue Mode for a Server [POST /servers/{id}/actions/enable_rescue]

Enable the Hetzner Rescue System for this server. The next time a Server with enabled rescue mode boots it will start a special minimal Linux distribution designed for repair and reinstall.

In case a server cannot boot on its own you can use this to access a server’s disks.

Rescue Mode is automatically disabled when you first boot into it or if you do not use it for 60 minutes.

Enabling rescue mode will not [reboot](#server-actions-soft-reboot-a-server) your server — you will have to do this yourself.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    Specify `type` to select the kind if rescue environment you want to boot.
    
    To access the booted rescue system by SSH key pass the respective SSH key ids in `ssh_keys`.
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | type | string&nbsp;(optional) | Type of rescue system to boot (default: `linux64`) <br>Choices: `linux64`, `linux32`, `freebsd64` |
    | ssh_keys | array&nbsp;(optional) | Array of SSH key IDs which should be injected into the rescue system. Only available for types: `linux64` and `linux32`. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"type": "linux64", "ssh_keys": [2323]}' \
  https://api.hetzner.cloud/v1/servers/42/actions/enable_rescue
     ```



    + Attributes
        + `type`: `linux64` (enum[string], optional) - Type of rescue system to boot (default: `linux64`)
            + Members
                + `linux64`
                + `linux32`
                + `freebsd64`
        + `ssh_keys`: `2323` (array[number], optional, fixed-type) - Array of SSH key IDs which should be injected into the rescue system. Only available for types: `linux64` and `linux32`.

+ Response 201 (application/json)
    The `root_password` key in the reply contains the root password that can be used to access the booted rescue system.
    
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
            
    #### Call specific error codes
    
    | Code | Description |
    |------|-------------|
    | `rescue_already_enabled` | Returned when the rescue system is already enabled |
    + Attributes
        + `root_password`: `zCWbFhnu950dUTko5f40` (string) - Root Password for Server once it gets booted in rescue mode
        + `action` (Action, required)
            + `command`: `enable_rescue` (string)
    + Body



## Disable Rescue Mode for a Server [POST /servers/{id}/actions/disable_rescue]

Disables the Hetzner Rescue System for a server. This makes a server start from its disks on next reboot.

Rescue Mode is automatically disabled when you first boot into it or if you do not use it for 60 minutes.

Disabling rescue mode will not reboot your server — you will have to do this yourself.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/disable_rescue
     ```




    
    
    #### Call specific error codes
    
    | Code | Description |
    |------|-------------|
    | `rescue_already_disabled` | Returned when the rescue system is already disabled |
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `disable_rescue` (string)
    + Body


## Create Image from a Server [POST /servers/{id}/actions/create_image]

Creates an image (snapshot) from a server by copying the contents of its disks. This creates a
snapshot of the current state of the disk and copies it into an image. If the server is currently running
you must make sure that its disk content is consistent. Otherwise, the created image may not be readable.

To make sure disk content is consistent, we recommend to shut down the server prior to creating an image.

You can either create a `backup` image that is bound to the server and therefore will be deleted when the
server is deleted, or you can create an `snapshot` image which is completely independent of the server it was
created from and will survive server deletion. Backup images are only available when the backup option is
enabled for the Server. Snapshot images are billed on a per GB basis.

+ Parameters
    + id: `42` (string) - ID of the server


+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | description | string&nbsp;(optional) | Description of the image. If you do not set this we auto-generate one for you. |
    | type | string&nbsp;(optional) | Type of image to create (default: `snapshot`) <br>Choices: `snapshot`, `backup` |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"description": "my image", "type": "snapshot"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/create_image
     ```



    + Attributes
        + description: `my image` (string, optional) - Description of the image. If you do not set this we auto-generate one for you.
        + `type`: `snapshot` (enum[string], optional) - Type of image to create (default: `snapshot`)
            + Members
                + `snapshot`
                + `backup`

+ Response 201 (application/json)
    The `image` key in the reply contains an the created image, which is an object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the image |
    | type | string | Type of the image <br>Choices: `snapshot`, `system`, `backup` |
    | status | string | Whether the image can be used or if it's still being created <br>Choices: `creating`, `available` |
    | name | string,&nbsp;null | Unique identifier of the image. This value is only set for system images. |
    | description | string | Description of the image |
    | image_size | number,&nbsp;null | Size of the image file in our storage in GB. For snapshot images this is the value relevant for calculating costs for the image. |
    | disk_size | number | Size of the disk contained in the image in GB. |
    | created | string | Point in time when the image was created (in ISO-8601 format) |
    | created_from | object,&nbsp;null | Information about the server the image was created from |
    | bound_to | number,&nbsp;null | ID of server the image is bound to. Only set for images of type `backup`. |
    | os_flavor | string | Flavor of operating system contained in the image <br>Choices: `ubuntu`, `centos`, `debian`, `fedora`, `unknown` |
    | os_version | string,&nbsp;null | Operating system version |
    | rapid_deploy | boolean | Indicates that rapid deploy of the image is available |
    | protection | object | Protection configuration for the image |
    | deprecated | string,&nbsp;null | Point in time when the image is considered to be deprecated (in ISO-8601 format) |




    
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
        
    + Attributes
        + `image` (Snapshot Image, required) - The newly created image
            + `description`: `my image` (string, required) - Description of the image
            + `status`: `creating` (enum[string], required) - Whether the image can be used or if it's still being created
                + Members
                    + `creating`
                    + `available`
        + `action` (Action, required)
            + `command`: `create_image` (string)


## Rebuild a Server from an Image [POST /servers/{id}/actions/rebuild]

Rebuilds a server overwriting its disk with the content of an image, thereby **destroying all data** on the target server

The image can either be one you have created earlier (`backup` or `snapshot` image) or it can be a
completely fresh `system` image provided by us. You can get a list of all available images with `GET /images`.

Your server will automatically be powered off before the rebuild command executes.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    To select which image to rebuild from you can either pass an id or a name as the `image` argument.
    Passing a name only works for `system` images since the other image types do not have a name set.

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | image | string&nbsp;(required) | ID or name of image to rebuilt from. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"image": "ubuntu-16.04"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/rebuild
     ```



    + Attributes
        + image: `ubuntu-16.04` (string, required) - ID or name of image to rebuilt from.

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `rebuild_server` (string)
        + `root_password` (string, required, nullable) - New root password when not using SSH keys
    + Body


## Change the Type of a Server [POST /servers/{id}/actions/change_type]

Changes the type (Cores, RAM and disk sizes) of a server.

Server must be powered off for this command to succeed.

This copies the content of its disk, and starts it again.

You can only migrate to server types with the same `storage_type` and equal or bigger disks. Shrinking disks is
not possible as it might destroy data.

If the disk gets upgraded, the server type can not be downgraded any more. If you plan to downgrade the server type, set `upgrade_disk` to `false`.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | upgrade_disk | boolean&nbsp;(optional) | If false, do not upgrade the disk. This allows downgrading the server type later. |
    | server_type | string&nbsp;(required) | ID or name of server type the server should migrate to |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"upgrade_disk": true, "server_type": "cx11"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/change_type
     ```



    + Attributes
        + upgrade_disk: `true` (boolean) - If false, do not upgrade the disk. This allows downgrading the server type later.
        + `server_type`: `cx11` (string, required) - ID or name of server type the server should migrate to
        
+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    #### Call specific error codes
    
    | Code | Description |
    |------|-------------|
    | `server_not_stopped` | Returned when the server was not powered off |
    | `invalid_server_type` | Returned when a type change to the new server_type is not possible |
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `change_server_type` (string)
    + Body


## Enable and Configure Backups for a Server [POST /servers/{id}/actions/enable_backup]

Enables and configures the automatic daily backup option for the server. Enabling automatic backups will increase the price of
 the server by 20%. In return, you will get seven slots where images of type backup can be stored.

Backups are automatically created daily.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    Passing the `backup_window` will select the time window for the daily backup to run. All times are in UTC. `22-02` means
    that the backup will be started between 10 PM and 2 AM. This will be done on a best-effort base, so we cannot
    guarantee the backup window will be met.
    
    If you do not provide a time window one will be randomly selected.
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | backup_window | string&nbsp;(optional) | Time window (UTC) in which the backup will run <br>Choices: `22-02`, `02-06`, `06-10`, `10-14`, `14-18`, `18-22` |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"backup_window": "22-02"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/enable_backup
     ```



    + Attributes
        + `backup_window`: `22-02` (enum[string], optional) - Time window (UTC) in which the backup will run
            + Members
                + `22-02`
                + `02-06`
                + `06-10`
                + `10-14`
                + `14-18`
                + `18-22`

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + `action` (Action, required)
            + `command`: `enable_backup` (string)
    + Body

## Disable Backups for a Server [POST /servers/{id}/actions/disable_backup]

Disables the automatic backup option and deletes all existing Backups for a Server.
No more additional charges for backups will be made.

Caution: This immediately removes all existing backups for the server!

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/disable_backup
     ```


   
    + Attributes
        + `action` (Action, required)
            + `command`: `disable_backup` (string)
    + Body


## Attach an ISO to a Server [POST /servers/{id}/actions/attach_iso]

Attaches an ISO to a server. The Server will immediately see it as a new disk.
An already attached ISO will automatically be detached before the new ISO is attached.

Servers with attached ISOs have a modified boot order: They will try to boot from the ISO first before falling back to
hard disk.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | iso | string&nbsp;(required) | ID or name of ISO to attach to the server as listed in GET `/isos` |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"iso": "FreeBSD-11.0-RELEASE-amd64-dvd1"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/attach_iso
     ```



    + Attributes
        + iso: `FreeBSD-11.0-RELEASE-amd64-dvd1` (string, required) - ID or name of ISO to attach to the server as listed in GET `/isos`

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
           
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `attach_iso` (string)
    + Body

## Detach an ISO from a Server [POST /servers/{id}/actions/detach_iso]

Detaches an ISO from a server. In case no ISO image is attached to the server,
the status of the returned action is immediately set to `success`.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/detach_iso
     ```


   
    
        
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `detach_iso` (string)
    + Body

## Change reverse DNS entry for this server [POST /servers/{id}/actions/change_dns_ptr]

Changes the hostname that will appear when getting the hostname belonging to the primary IPs (ipv4 and ipv6) of this server.

Floating IPs assigned to the server are not affected by this.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    Select the IP address for which to change the dns entry by passing `ip`. It can be either ipv4 or ipv6.
    The target hostname is set by passing `dns_ptr`.
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | ip | string&nbsp;(required) | Primary IP address for which the reverse DNS entry should be set. |
    | dns_ptr | string,&nbsp;null&nbsp;(required) | Hostname to set as a reverse DNS PTR entry. Will reset to original value if `null` |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"ip": "1.2.3.4", "dns_ptr": "server01.example.com"}' \
  https://api.hetzner.cloud/v1/servers/42/actions/change_dns_ptr
     ```



    + Attributes
        + `ip`: `1.2.3.4` (string, required) - Primary IP address for which the reverse DNS entry should be set.
        + `dns_ptr`: `server01.example.com` (string, required, nullable) - Hostname to set as a reverse DNS PTR entry. Will reset to original value if `null`

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |



   
    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `change_dns_ptr` (string)
    + Body

## Change protection for a Server [POST /servers/{id}/actions/change_protection]

Changes the protection configuration of the server.

+ Parameters
    + id: `42` (string) - ID of the server

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | delete | boolean&nbsp;(optional) | If true, prevents the server from being deleted (currently delete and rebuild attribute needs to have the same value) |
    | rebuild | boolean&nbsp;(optional) | If true, prevents the server from being rebuilt` (currently delete and rebuild attribute needs to have the same value) |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"delete": true, "rebuild": true}' \
  https://api.hetzner.cloud/v1/servers/42/actions/change_protection
     ```



    + Attributes
        + `delete`: `true` (boolean, optional) - If true, prevents the server from being deleted (currently delete and rebuild attribute needs to have the same value)
        + `rebuild`: `true` (boolean, optional) - If true, prevents the server from being rebuilt` (currently delete and rebuild attribute needs to have the same value)

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
        
    + Attributes
        + `action` (Action, required)
            + `command`: `change_protection` (string)
    + Body






## Request Console for a Server [POST /servers/{id}/actions/request_console]

Requests credentials for remote access via vnc over websocket to keyboard, monitor, and mouse for a server.
The provided url is valid for 1 minute, after this period a new url needs to be created to connect to the server.
How long the connection is open after the initial connect is not subject to this timeout.

+ Parameters
    + id: `42` (string) - ID of the server

+ Response 201 (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | wss_url | string | URL of websocket proxy to use. This includes a token which is valid for a limited time only. |
    | password | string | VNC password to use for this connection. This password only works in combination with a wss_url with valid token. |
    | action | object | Action created by this call |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/servers/42/actions/request_console
     ```




    The `action` key in the reply contains an action object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    
        
    + Attributes
        + `wss_url`: `wss://console.hetzner.cloud/?server_id=1&token=3db32d15-af2f-459c-8bf8-dee1fd05f49c` (string, required) - URL of websocket proxy to use. This includes a token which is valid for a limited time only.
        + `password`: `9MQaTg2VAGI0FIpc10k3UpRXcHj2wQ6x` (string, required) - VNC password to use for this connection. This password only works in combination with a wss_url with valid token.
        + `action` (Action, required) - Action created by this call
            + `command`: `request_console` (string)
    + Body


# Floating IPs [/floating_ips]

Floating IPs help you to create highly available setups. You can assign a Floating IP
to any server. The server can then use this IP. You can reassign it
to a different server at any time, or you can choose to unassign the IP from servers all together.

Floating IPs can be used globally. This means you can assign a Floating IP to a server in one
location and later reassign it to a server in a different location. For optimal routing and latency Floating IPs
should be used in the location they were created in.

For Floating IPs to work with your server, you must configure them inside your operation system.

Floating IPs of type `ipv4` use a single ipv4 address as their `ip` property. Floating IPs of type `ipv6` use a
/64 network such as `fc00::/64` as their `ip` property. Any IP address within that network can be used on your host.

Floating IPs are billed on a monthly basis.

## Get all Floating IPs [GET /floating_ips]

Returns all floating ip objects.

+ Response 200 (application/json)
    The `floating_ips` key in the reply contains an array of floating ip objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the Floating IP |
    | description | string,&nbsp;null | Description of the Floating IP |
    | ip | string | IP address of the Floating IP |
    | type | string | Type of the Floating IP <br>Choices: `ipv4`, `ipv6` |
    | server | number,&nbsp;null | Id of the Server the Floating IP is assigned to, null if it is not assigned at all. |
    | dns_ptr | array | Array of reverse DNS entries |
    | home_location | object | Location the Floating IP was created in. Routing is optimized for this location. |
    | blocked | boolean | Whether the IP is blocked |
    | protection | object | Protection configuration for the Floating IP |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips
     ```



    + Attributes
        + floating_ips (array[FloatingIP], required, fixed-type)

## Get a specific Floating IP [GET /floating_ips/{id}]

Returns a specific floating ip object.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Response 200 (application/json)
    The `floating_ip` key in the reply contains a floating ip object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the Floating IP |
    | description | string,&nbsp;null | Description of the Floating IP |
    | ip | string | IP address of the Floating IP |
    | type | string | Type of the Floating IP <br>Choices: `ipv4`, `ipv6` |
    | server | number,&nbsp;null | Id of the Server the Floating IP is assigned to, null if it is not assigned at all. |
    | dns_ptr | array | Array of reverse DNS entries |
    | home_location | object | Location the Floating IP was created in. Routing is optimized for this location. |
    | blocked | boolean | Whether the IP is blocked |
    | protection | object | Protection configuration for the Floating IP |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips/4711
     ```



    + Attributes
        + floating_ip (FloatingIP, required)

## Create a Floating IP [POST /floating_ips]

Creates a new Floating IP assigned to a server. If you want to create a Floating IP that is not bound
to a server, you need to provide the `home_location` key instead of `server`. This can be either the ID or
the name of the location this IP shall be created in. Note that a Floating IP can be assigned to a server
in any location later on. For optimal routing it is advised to use the Floating IP in the same Location
it was created in.

+ Request (application/json)
    The `type` argument is required while `home_location` and `server` are mutually exclusive.

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | type | string&nbsp;(required) | Floating IP type <br>Choices: `ipv4`, `ipv6` |
    | server | number&nbsp;(optional) | Server to assign the Floating IP to |
    | home_location | string&nbsp;(optional) | Home location (routing is optimized for that location). Only optional if server argument is passed. |
    | description | string&nbsp;(optional) |  |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"type": "ipv4", "server": 42, "home_location": "nbg1", "description": "Web Frontend"}' \
  https://api.hetzner.cloud/v1/floating_ips
     ```



    + Attributes
        + `type`: `ipv4` (enum[string], required) - Floating IP type
            + Members
                + `ipv4`
                + `ipv6`
        + `server`: `42` (number, optional) - Server to assign the Floating IP to
        + `home_location`: `nbg1` (string, optional) - Home location (routing is optimized for that location). Only optional if server argument is passed.
        + `description`: `Web Frontend` (string, optional) – Description

+ Response 201 (application/json)
    The `floating_ip` key in the reply contains the object that was just created:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the Floating IP |
    | description | string,&nbsp;null | Description of the Floating IP |
    | ip | string | IP address of the Floating IP |
    | type | string | Type of the Floating IP <br>Choices: `ipv4`, `ipv6` |
    | server | number,&nbsp;null | Id of the Server the Floating IP is assigned to, null if it is not assigned at all. |
    | dns_ptr | array | Array of reverse DNS entries |
    | home_location | object | Location the Floating IP was created in. Routing is optimized for this location. |
    | blocked | boolean | Whether the IP is blocked |
    | protection | object | Protection configuration for the Floating IP |




    
    If you chose to bind the Floating IP to a server on creation the `action` key in the reply contains the action
    that tracks assignment of the IP to the specified server:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + floating_ip (FloatingIP, required)
        + `action` (ActionRunning, optional) - Assign action (only present if a server was provided)
            + `command`: `assign_floating_ip` (string)
      



## Change description of a Floating IP [PUT /floating_ips/{id}]

Changes the description of a Floating IP.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | description | string&nbsp;(optional) | New Description to set |

    > #### Example curl
     ```bash
     curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"description": "New description"}' \
  https://api.hetzner.cloud/v1/floating_ips/4711
     ```



    + Attributes
        + description: `New description` (string, optional) - New Description to set

+ Response 200 (application/json)
    The `floating_ip` key in the reply contains the modified Floating IP object with the new description:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the Floating IP |
    | description | string,&nbsp;null | Description of the Floating IP |
    | ip | string | IP address of the Floating IP |
    | type | string | Type of the Floating IP <br>Choices: `ipv4`, `ipv6` |
    | server | number,&nbsp;null | Id of the Server the Floating IP is assigned to, null if it is not assigned at all. |
    | dns_ptr | array | Array of reverse DNS entries |
    | home_location | object | Location the Floating IP was created in. Routing is optimized for this location. |
    | blocked | boolean | Whether the IP is blocked |
    | protection | object | Protection configuration for the Floating IP |




    + Attributes
        + `floating_ip` (FloatingIP, required)
            + `description`: `New description` (string, nullable, required) - Description of the Floating IP        
        

## Delete a Floating IP [DELETE /floating_ips/{id}]

Deletes a Floating IP. If it is currently assigned to a server it will automatically get unassigned.

> #### Example curl
 ```bash
 curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips/4711
 ```




+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Response 204



# Floating IP Actions [/floating_ips/{id}/actions/]

## Get all Actions for a Floating IP [GET /floating_ips/{id}/actions{?sort}]

Returns all action objects for a Floating IP. You can sort the results by using the `sort` URI parameter.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP
    + sort (enum[string], optional) - Can be used multiple times.
        + Members
            + `id`
            + `id:asc`
            + `id:desc`
            + `command`
            + `command:asc`
            + `command:desc`
            + `status`
            + `status:asc`
            + `status:desc`
            + `progress`
            + `progress:asc`
            + `progress:desc`
            + `started`
            + `started:asc`
            + `started:desc`
            + `finished`
            + `finished:asc`
            + `finished:desc`

+ Response 200 (application/json)
    The `actions` key in the reply contains an array of action objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips/4711/actions
     ```



    + Attributes
        + actions (array, required, fixed-type)
            + (Action)
                + `command`: `assign_floating_ip` (string)

## Get an Action for a Floating IP [GET /floating_ips/{id}/actions/{action_id}]

Returns a specific action object for a Floating IP.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP
    + action_id: `1337` (string) - ID of the action

+ Response 200 (application/json)
    The `action` key in the reply has this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips/4711/actions/1337
     ```



    + Attributes
        + action (Action, required)
            + `command`: `assign_floating_ip` (string)

## Assign a Floating IP to a Server [POST /floating_ips/{id}/actions/assign]

Assigns a Floating IP to a server. 

+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | server | number&nbsp;(required) | ID of the server the Floating IP shall be assigned to |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"server": 43}' \
  https://api.hetzner.cloud/v1/floating_ips/4711/actions/assign
     ```



    + Attributes
        + server: `43` (number, required) - ID of the server the Floating IP shall be assigned to

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + action (ActionRunning, required)
            + `command`: `assign_floating_ip` (string)
            + `resources` (array, required, fixed-type) - Resources the action relates to
                + (ResourceRef)
                    + `id`: `43` (number, required) - ID of resource referenced
        

## Unassign a Floating IP [POST /floating_ips/{id}/actions/unassign]

Unassigns a Floating IP, resulting in it being unreachable. You may assign it to a server again at a later time.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/floating_ips/4711/actions/unassign
     ```



    + Attributes
        + action (ActionRunning, required)
            + `command`: `unassign_floating_ip` (string)

## Change reverse DNS entry for a Floating IP [POST /floating_ips/{id}/actions/change_dns_ptr]

Changes the hostname that will appear when getting the hostname belonging to this Floating IP.

+ Parameters
    + id: `4711` (string) - ID of the Floating IP

+ Request (application/json)
    Select the IP address for which to change the dns entry by passing `ip`. For a Floating IP of type `ipv4` this must exactly
    match the IP address of the Floating IP. For a Floating IP of type `ipv6` this must be a single IP within the ipv6 /64 range
    that belongs to this Floating IP.
    
    The target hostname is set by passing `dns_ptr`.
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | ip | string&nbsp;(required) | IP address for which to set the reverse DNS entry |
    | dns_ptr | string,&nbsp;null&nbsp;(required) | Hostname to set as a reverse DNS PTR entry, will reset to original default value if `null` |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"ip": "1.2.3.4", "dns_ptr": "server02.example.com"}' \
  https://api.hetzner.cloud/v1/floating_ips/4711/actions/change_dns_ptr
     ```



    + Attributes
        + ip: `1.2.3.4` (string, required) - IP address for which to set the reverse DNS entry
        + dns_ptr: `server02.example.com` (string, required, nullable) - Hostname to set as a reverse DNS PTR entry, will reset to original default value if `null`

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `running`, `success`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + `action` (ActionRunning, required)
            + `command`: `change_dns_ptr` (string)
            + `resources` (array, required, fixed-type) - Resources the action relates to
                + (ResourceRef)
                    + `id`: `4711` (number, required) - ID of resource referenced
                    + `type`: `floating_ip` (string, required) - Type of resource referenced

## Change protection [POST /floating_ips/{id}/actions/change_protection]

Changes the protection configuration of the Floating IP.


+ Parameters
    + id: `42` (string) - ID of the Floating IP

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | delete | boolean&nbsp;(optional) | If true, prevents the Floating IP from being deleted |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"delete": true}' \
  https://api.hetzner.cloud/v1/floating_ips/42/actions/change_protection
     ```



    + Attributes
        + `delete`: `true` (boolean, optional) - If true, prevents the Floating IP from being deleted

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + `action` (Action, required)
            + `command`: `change_protection` (string)
    + Body
# SSH Keys [/ssh_keys]

SSH keys are public keys you provide to the cloud system. They can be injected into servers at creation time. We highly
recommend that you use keys instead of passwords to manage your servers.

## Get all SSH keys [GET /ssh_keys{?name,fingerprint}]

Returns all SSH key objects.

+ Parameters
    + name (string, optional) - Can be used to filter SSH keys by their name. The response will only contain the SSH key matching the specified name.
    + fingerprint (string, optional) - Can be used to filter SSH keys by their fingerprint. The response will only contain the SSH key matching the specified fingerprint.
+ Response 200 (application/json)
    The `ssh_keys` key in the reply contains an array of SSH key objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the SSH key |
    | name | string | Name of the SSH key (must be unique per project) |
    | fingerprint | string | Fingerprint of public key |
    | public_key | string | Public key |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/ssh_keys
     ```



    + Attributes
        + `ssh_keys` (array[SSHKey], required, fixed-type)
    + Body

## Get an SSH key [GET /ssh_keys/{id}]

Returns a specific SSH key object.

+ Parameters
    + id: `2323` (string) - ID of the SSH key

+ Response 200 (application/json)
    The `ssh_key` key in the reply contains an SSH key object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the SSH key |
    | name | string | Name of the SSH key (must be unique per project) |
    | fingerprint | string | Fingerprint of public key |
    | public_key | string | Public key |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/ssh_keys/2323
     ```



    + Attributes
        + `ssh_key` (SSHKey, required)
    + Body

## Create an SSH key [POST /ssh_keys]

Creates a new SSH key with the given `name` and `public_key`. Once an SSH key is created, it can be used in other calls
such as creating servers.

+ Request (application/json)
    + Attributes
        + `name`: `My ssh key` (string, required) - Name of the SSH key
        + `public_key`: `ssh-rsa AAAjjk76kgf...Xt` (string, required) - Public key

+ Response 201 (application/json)
    The `ssh_key` key in the reply contains the object that was just created:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the SSH key |
    | name | string | Name of the SSH key (must be unique per project) |
    | fingerprint | string | Fingerprint of public key |
    | public_key | string | Public key |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"name": "My ssh key", "public_key": "ssh-rsa AAAjjk76kgf...Xt"}' \
  https://api.hetzner.cloud/v1/ssh_keys
     ```



    + Attributes
        + `ssh_key` (SSHKey, required)
    + Body

## Change the name of an SSH key [PUT /ssh_keys/{id}]

Changes the name of an SSH key.

+ Parameters
    + id: `2323` (string) - ID of the SSH key

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | name | string&nbsp;(optional) | New name Name to set |

    > #### Example curl
     ```bash
     curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"name": "New name"}' \
  https://api.hetzner.cloud/v1/ssh_keys/2323
     ```



    + Attributes
        + `name`: `New name` (string, optional) - New name Name to set

+ Response 200 (application/json)
    The `ssh_key` key in the reply contains the modified SSH key object with the new description:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the SSH key |
    | name | string | Name of the SSH key |
    | fingerprint | string | Fingerprint of public key |
    | public_key | string | Public key |




    + Attributes
        + `ssh_key` (SSHKey, required)
            + `name`: `New name` (string, required) - Name of the SSH key
    + Body

## Delete an SSH key [DELETE /ssh_keys/{id}]

Deletes an SSH key. It cannot be used anymore.

> #### Example curl
 ```bash
 curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/ssh_keys/2323
 ```




+ Parameters
    + id: `2323` (string) - ID of the SSH key

+ Response 204
# Server Types [/server_types]

Server types define kinds of servers offered. Each type has a hourly and a monthly cost.
You will pay whichever cost is lower for your usage of this specific server. Costs may differ between locations.

Currency for all amounts is €. All prices exclude VAT.

## Get all Server Types [GET /server_types{?name}]

Gets all server type objects.

+ Parameters
    + name (string, optional) - Can be used to filter server types by their name. The response will only contain the server type matching the specified name.
+ Response 200 (application/json)
    The `server_types` key in the reply contains an array of server type objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the server type |
    | name | string | Unique identifier of the server type |
    | description | string | Description of the server type |
    | cores | number | Number of cpu cores a server of this type will have |
    | memory | number | Memory a server of this type will have in GB |
    | disk | number | Disk size a server of this type will have in GB |
    | prices | array | Prices in different Locations |
    | storage_type | string | Type of server boot drive. Local has higher speed. Network has better availability. <br>Choices: `local`, `network` |
    | cpu_type | string | Type of cpu. <br>Choices: `shared` |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/server_types
     ```



    + Attributes
        + `server_types` (array[ServerType], required, fixed-type)
    + Body

## Get a Server Type [GET /server_types/{id}]

Gets a specific server type object.


+ Parameters
    + id: `1` (string) - ID of server type

+ Response 200 (application/json)
    The `server_type` key in the reply contains a server type object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the server type |
    | name | string | Unique identifier of the server type |
    | description | string | Description of the server type |
    | cores | number | Number of cpu cores a server of this type will have |
    | memory | number | Memory a server of this type will have in GB |
    | disk | number | Disk size a server of this type will have in GB |
    | prices | array | Prices in different Locations |
    | storage_type | string | Type of server boot drive. Local has higher speed. Network has better availability. <br>Choices: `local`, `network` |
    | cpu_type | string | Type of cpu. <br>Choices: `shared` |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/server_types/1
     ```



    + Attributes
        + `server_type` (ServerType, required)
    + Body
# Locations [/locations]

Datacenters are organized by locations. Datacenters in the same location are connected with very low latency links.

## Get all Locations [GET /locations{?name}]

Returns all location objects.

+ Parameters
    + name (string, optional) - Can be used to filter locations by their name. The response will only contain the location matching the specified name.
+ Response 200 (application/json)
    The `locations` key in the reply contains an array of location objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the location |
    | name | string | Unique identifier of the location |
    | description | string | Description of the location |
    | country | string | ISO 3166-1 alpha-2 code of the country the location resides in |
    | city | string | City the location is closest to |
    | latitude | number | Latitude of the city closest to the location |
    | longitude | number | Longitude of the city closest to the location |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/locations
     ```



    + Attributes
        + `locations` (array[Location], required, fixed-type)


## Get a Location [GET /locations/{id}]

Returns a specific location object.

+ Parameters
    + id: `1` (string) - ID of location

+ Response 200 (application/json)
    The `location` key in the reply contains a location object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the location |
    | name | string | Unique identifier of the location |
    | description | string | Description of the location |
    | country | string | ISO 3166-1 alpha-2 code of the country the location resides in |
    | city | string | City the location is closest to |
    | latitude | number | Latitude of the city closest to the location |
    | longitude | number | Longitude of the city closest to the location |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/locations/1
     ```



    + Attributes
        + `location` (Location, required)
    + Body
# Datacenters [/datacenters]

Each datacenter represents a physical entity where virtual machines are hosted. Note that datacenters are **not** guaranteed to be entirely independent failure domains.

Datacenters in the same location are connected by very low latency links.

Datacenter names are made up of their location and datacenter number, for example `fsn1-dc8` means datacenter `8` in location `fsn1`.


## Get all Datacenters [GET /datacenters{?name}]

Returns all datacenter objects.

+ Parameters
    + name (string, optional) - Can be used to filter datacenters by their name. The response will only contain the datacenter matching the specified name. When the name does not match the datacenter name format, an `invalid_input` error is returned.
+ Response 200 (application/json)
    The `datacenters` key in the reply contains an array of datacenter objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the datacenter |
    | name | string | Unique identifier of the datacenter |
    | description | string | Description of the datacenter |
    | location | object | Location where the datacenter resides in |
    | server_types | object | The server types the datacenter can handle |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/datacenters
     ```



    + Attributes
        + `datacenters` (array[Datacenter], required, fixed-type)
        + `recommendation` `1` (number) - The datacenter which is recommended to be used to create new servers.

## Get a Datacenter [GET /datacenters/{id}]

Returns a specific datacenter object.

+ Parameters
    + id: `1` (string) - ID of datacenter

+ Response 200 (application/json)
    The `datacenter` key in the reply contains a datacenter object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the datacenter |
    | name | string | Unique identifier of the datacenter |
    | description | string | Description of the datacenter |
    | location | object | Location where the datacenter resides in |
    | server_types | object | The server types the datacenter can handle |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/datacenters/1
     ```



    + Attributes
        + `datacenter` (Datacenter, required)
    + Body
# Images [/images]

Images are blueprints for your VM disks. They can be of different types:

### System images

Distribution images maintained by us, e.g. "Ubuntu 16.04"

### Snapshot Images

Maintained by you, for example "Ubuntu 16.04 with my own settings". These are billed per GB per month.

### Backup images

Daily Backups of your server. Will automatically be created for servers which have backups enabled (`POST /servers/{id}/actions/enable_backup`)

Bound to exactly one server. If you delete the server, you also delete all backups bound to it. You may convert backup images to snapshot images to keep them.

These are billed at 20% of your instance price for 7 backup slots.

## Get all Images  [GET /images{?sort,type,bound_to,name}]

Returns all image objects. You can select specific image types only and sort the results by using URI parameters.

+ Parameters
    + sort (enum[string], optional) - Can be used multiple times.
        + Members
            + `id`
            + `id:asc`
            + `id:desc`
            + `name`
            + `name:asc`
            + `name:desc`
            + `created`
            + `created:asc`
            + `created:desc`
    + type (enum[string], optional) - Can be used multiple times.
        + Members
            + `system`
            + `snapshot`
            + `backup`
    + bound_to (string, optional) - Can be used multiple times. Server Id linked to the image. Only available for images of type `backup`
    + name (string, optional) - Can be used to filter images by their name. The response will only contain the image matching the specified name.

+ Response 200 (application/json)
    The `images` key in the reply contains an array of image objects with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the image |
    | type | string | Type of the image <br>Choices: `system`, `snapshot`, `backup` |
    | status | string | Whether the image can be used or if it's still being created <br>Choices: `available`, `creating` |
    | name | string,&nbsp;null | Unique identifier of the image. This value is only set for system images. |
    | description | string | Description of the image |
    | image_size | number,&nbsp;null | Size of the image file in our storage in GB. For snapshot images this is the value relevant for calculating costs for the image. |
    | disk_size | number | Size of the disk contained in the image in GB. |
    | created | string | Point in time when the image was created (in ISO-8601 format) |
    | created_from | object,&nbsp;null | Information about the server the image was created from |
    | bound_to | number,&nbsp;null | ID of server the image is bound to. Only set for images of type `backup`. |
    | os_flavor | string | Flavor of operating system contained in the image <br>Choices: `ubuntu`, `centos`, `debian`, `fedora`, `unknown` |
    | os_version | string,&nbsp;null | Operating system version |
    | rapid_deploy | boolean | Indicates that rapid deploy of the image is available |
    | protection | object | Protection configuration for the image |
    | deprecated | string,&nbsp;null | Point in time when the image is considered to be deprecated (in ISO-8601 format) |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/images
     ```



    + Attributes
        + `images` (array[Image], required, fixed-type)
    + Body

## Get an Image  [GET /images/{id}]

Returns a specific image object.

+ Parameters
    + id: `4711` (string) - ID of the image

+ Response 200 (application/json)
    The `image` key in the reply contains an image object with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the image |
    | type | string | Type of the image <br>Choices: `system`, `snapshot`, `backup` |
    | status | string | Whether the image can be used or if it's still being created <br>Choices: `available`, `creating` |
    | name | string,&nbsp;null | Unique identifier of the image. This value is only set for system images. |
    | description | string | Description of the image |
    | image_size | number,&nbsp;null | Size of the image file in our storage in GB. For snapshot images this is the value relevant for calculating costs for the image. |
    | disk_size | number | Size of the disk contained in the image in GB. |
    | created | string | Point in time when the image was created (in ISO-8601 format) |
    | created_from | object,&nbsp;null | Information about the server the image was created from |
    | bound_to | number,&nbsp;null | ID of server the image is bound to. Only set for images of type `backup`. |
    | os_flavor | string | Flavor of operating system contained in the image <br>Choices: `ubuntu`, `centos`, `debian`, `fedora`, `unknown` |
    | os_version | string,&nbsp;null | Operating system version |
    | rapid_deploy | boolean | Indicates that rapid deploy of the image is available |
    | protection | object | Protection configuration for the image |
    | deprecated | string,&nbsp;null | Point in time when the image is considered to be deprecated (in ISO-8601 format) |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/images/4711
     ```



    + Attributes
        + `image` (Image, required)
    + Body

## Update an Image [PUT /images/{id}]

Updates the Image. You may change the description or convert a Backup image to a Snapshot Image. Only images of type `snapshot` and
`backup` can be updated.

+ Parameters
    + id: `4711` (string) - ID of the image to be updated

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | description | string&nbsp;(optional) | New description of Image |
    | type | string&nbsp;(optional) | Destination image type to convert to <br>Choices: `snapshot` |

    > #### Example curl
     ```bash
     curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"description": "My new Image description", "type": "the new image type"}' \
  https://api.hetzner.cloud/v1/images/4711
     ```



    + Attributes
        + `description`: `My new Image description` (string, optional) - New description of Image
        + `type`: `the new image type` (enum[string], optional) - Destination image type to convert to
            + Members
                + `snapshot`

+ Response 200 (application/json)
    The `image` key in the reply contains the modified image object:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the image |
    | type | string | Type of the image <br>Choices: `snapshot`, `system`, `backup` |
    | status | string | Whether the image can be used or if it's still being created <br>Choices: `available`, `creating` |
    | name | string,&nbsp;null | Unique identifier of the image. This value is only set for system images. |
    | description | string | Description of the image |
    | image_size | number,&nbsp;null | Size of the image file in our storage in GB. For snapshot images this is the value relevant for calculating costs for the image. |
    | disk_size | number | Size of the disk contained in the image in GB. |
    | created | string | Point in time when the image was created (in ISO-8601 format) |
    | created_from | object,&nbsp;null | Information about the server the image was created from |
    | bound_to | number,&nbsp;null | ID of server the image is bound to. Only set for images of type `backup`. |
    | os_flavor | string | Flavor of operating system contained in the image <br>Choices: `ubuntu`, `centos`, `debian`, `fedora`, `unknown` |
    | os_version | string,&nbsp;null | Operating system version |
    | rapid_deploy | boolean | Indicates that rapid deploy of the image is available |
    | protection | object | Protection configuration for the image |
    | deprecated | string,&nbsp;null | Point in time when the image is considered to be deprecated (in ISO-8601 format) |




    + Attributes
        + `image` (Snapshot Image)
            + `description`: `My new Image description` (string, required) - Description of the image

## Delete an Image [DELETE /images/{id}]

Deletes an Image. Only images of type `snapshot` and `backup` can be deleted.

> #### Example curl
 ```bash
 curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/images/4711
 ```




+ Parameters
    + id: `4711` (string) - ID of the Image

+ Response 204


# Image Actions [/images/{id}/actions/]

## Get all Actions for an Image [GET /images/{id}/actions{?sort}]

Returns all action objects for an image. You can sort the results by using the `sort` URI parameter.

+ Parameters
    + id: `4711` (string) - ID of the Image
    + sort (enum[string], optional) - Can be used multiple times.
        + Members
            + `id`
            + `id:asc`
            + `id:desc`
            + `command`
            + `command:asc`
            + `command:desc`
            + `status`
            + `status:asc`
            + `status:desc`
            + `progress`
            + `progress:asc`
            + `progress:desc`
            + `started`
            + `started:asc`
            + `started:desc`
            + `finished`
            + `finished:asc`
            + `finished:desc`

+ Response 200 (application/json)
    The `actions` key in the reply contains an array of action objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/images/4711/actions
     ```



    + Attributes
        + actions (array, required, fixed-type)
            + (Action)
                + `command`: `change_protection` (string)

## Get an Action for an Image [GET /images/{id}/actions/{action_id}]

Returns a specific action object for an image.

+ Parameters
    + id: `4711` (string) - ID of the image
    + action_id: `1337` (string) - ID of the action

+ Response 200 (application/json)
    The `action` key in the reply contains an array of action objects with this structure:

    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/images/4711/actions/1337
     ```



    + Attributes
        + action (Action, required)
            + `command`: `change_protection` (string)

## Change protection for an Image [POST /images/{id}/actions/change_protection]

Changes the protection configuration of the image. Can only be used on snapshots.

+ Parameters
    + id: `42` (string) - ID of the image

+ Request (application/json)
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | delete | boolean&nbsp;(optional) | If true, prevents the snapshot from being deleted |

    > #### Example curl
     ```bash
     curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $API_TOKEN" \
  -d '{"delete": true}' \
  https://api.hetzner.cloud/v1/images/42/actions/change_protection
     ```



    + Attributes
        + `delete`: `true` (boolean, optional) - If true, prevents the snapshot from being deleted

+ Response 201 (application/json)
    The `action` key in the reply contains an action object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the action |
    | command | string | Command executed in the action |
    | status | string | Status of the action <br>Choices: `success`, `running`, `error` |
    | progress | number | Progress of action in percent |
    | started | string | Point in time when the action was started (in ISO-8601 format) |
    | finished | string,&nbsp;null | Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null. |
    | resources | array | Resources the action relates to |
    | error | object,&nbsp;null | Error message for the action if error occured, otherwise null. |




    + Attributes
        + `action` (Action, required)
            + `command`: `change_protection` (string)
    + Body


# ISOs [/isos]

ISOs are Read-Only images of DVDs. While we recommend using our image functionality to install your servers we also provide some
stock ISOs so you can install more exotic operating systems by yourself.

On request our support uploads a private ISO just for you. These are marked with type `private` and only visible in your
project.

To attach an ISO to your server use `POST /servers/{id}/actions/attach_iso`.

## Get all ISOs  [GET /isos{?name}]

Returns all available iso objects.

+ Parameters
    + name (string, optional) - Can be used to filter isos by their name. The response will only contain the iso matching the specified name.
+ Response 200 (application/json)
    The `isos` key in the reply contains an array of iso objects with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the ISO |
    | name | string,&nbsp;null | Unique identifier of the ISO. Only set for public ISOs |
    | description | string | Description of the ISO |
    | type | string | Type of the ISO <br>Choices: `public`, `private` |
    | deprecated | string,&nbsp;null | ISO 8601 timestamp of deprecation, null if ISO is still available. After the deprecation time it will no longer be possible to attach the ISO to servers. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/isos
     ```



    + Attributes
        + `isos` (array[ISO], required, fixed-type)
    + Body

## Get an ISO  [GET /isos/{id}]

Returns a specific iso object.

+ Parameters
    + id: `4711` (string) - ID of the ISO

+ Response 200 (application/json)
    The `iso` key in the reply contains an array of iso objects with this structure:
        
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | id | number | ID of the ISO |
    | name | string,&nbsp;null | Unique identifier of the ISO. Only set for public ISOs |
    | description | string | Description of the ISO |
    | type | string | Type of the ISO <br>Choices: `public`, `private` |
    | deprecated | string,&nbsp;null | ISO 8601 timestamp of deprecation, null if ISO is still available. After the deprecation time it will no longer be possible to attach the ISO to servers. |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/isos/4711
     ```



    + Attributes
        + `iso` (ISO, required)
    + Body







# Pricing [/pricing]

Returns prices for resources.

## Get all prices  [GET /pricing]

Returns prices for all resources available on the platform.
VAT and currency of the project owner are used for calculations.



Both net and gross prices are included in the response.

+ Response 200 (application/json)
    The `pricing` key in the reply contains an pricing object with this structure:
    
    | Parameter | Type | Description |
    |-------------|-------------|-------------|
    | currency | string | Currency the returned prices are expressed in, coded according to ISO 4217. |
    | vat_rate | string | The VAT rate used for calculating prices with VAT |
    | image | object | The cost of one 1GB Image for the full month. |
    | floating_ip | object | The cost of one floating IP per month. |
    | traffic | object | The cost of additional traffic per GB |
    | server_backup | object | Will increase base server costs by specific percentage. |
    | server_types | array | Costs of server types per location and type |

    > #### Example curl
     ```bash
     curl -H "Authorization: Bearer $API_TOKEN" \
  https://api.hetzner.cloud/v1/pricing
     ```


    
    + Attributes
        + `pricing` (Price, required, fixed-type)
    + Body


# Data Structures

## SSHKey
+ `id`: `2323` (number, required) - ID of the SSH key
+ `name`: `My ssh key` (string, required) - Name of the SSH key (must be unique per project)
+ `fingerprint`: `b7:2f:30:a0:2f:6c:58:6c:21:04:58:61:ba:06:3b:2f` (string, required) - Fingerprint of public key
+ `public_key`: `ssh-rsa AAAjjk76kgf...Xt` (string, required) - Public key

## Server create (object)
+ `name`: `my-server` (string, required) - Name of the server to create (must be unique per project and a valid hostname as per RFC 1123)
+ `server_type`: `cx11` (string, required) - ID or name of the server type this server should be created with
+ One Of
    + `location`: `nbg1` (string, optional) - ID or name of location to create server in.
    + `datacenter`: `nbg1-dc3` (string, optional) - ID or name of datacenter to create server in.
+ `start_after_create`: `true` (boolean, optional) - Start Server right after creation. Defaults to true.
+ `image`: `ubuntu-16.04` (string, required) - ID or name of the image the server is created from
+ `ssh_keys`: `my-ssh-key` (array, optional, fixed-type) - SSH key IDs or names which should be injected into the server at creation time
+ `user_data`: `#cloud-config\nruncmd:\n- [touch, /root/cloud-init-worked]\n` (string, optional) - Cloud-Init user data to use during server creation. This field is limited to 32KiB.


## Server (object)
+ `id`: `42` (number, required) - ID of server
+ `name`: `my-server` (string, required) - Name of the server (must be unique per project and a valid hostname as per RFC 1123)
+ `status`: `running` (enum[string], required) - Status of the server
    + Members
        + `running`
        + `initializing`
        + `starting`
        + `stopping`
        + `off`
        + `deleting`
        + `migrating`
        + `rebuilding`
        + `unknown`
+ `created`: `2016-01-30T23:50+00:00` (string, required) - Point in time when the server was created (in ISO-8601 format)
+ `public_net` (Public Network, required) - Public network information. The servers ipv4 address can be found in `public_net->ipv4->ip`
+ `server_type` (ServerType, required) - Type of server - determines how much ram, disk and cpu a server has
+ `datacenter` (Datacenter, required) - Datacenter this server is located at
+ `image` (Image, required, nullable) - Image this server was created from.
+ `iso` (ISO, required, nullable) - ISO image that is attached to this server. Null if no ISO is attached.
+ `rescue_enabled`: `false` (boolean, required) - True if rescue mode is enabled: Server will then boot into rescue system on next reboot.
+ `locked`: `false` (boolean, required) - True if server has been locked and is not available to user.
+ `backup_window`: `22-02` (string, required, nullable) - Time window (UTC) in which the backup will run, or null if the backups are not enabled
+ `outgoing_traffic`: `123456` (number, required, nullable) - Outbound Traffic for the current billing period in bytes
+ `ingoing_traffic`: `123456` (number, required, nullable) - Inbound Traffic for the current billing period in bytes
+ `included_traffic`: `654321` (number, required) - Free Traffic for the current billing period in bytes
+ `protection` (ServerProtection, required) - Protection configuration for the server


## ServerType
+ `id`: `1` (number, required) - ID of the server type
+ `name`: `cx11` (string, required) - Unique identifier of the server type
+ `description`: `CX11` (string, required) - Description of the server type
+ `cores`: `1` (number, required) - Number of cpu cores a server of this type will have
+ `memory`: `1` (number, required) - Memory a server of this type will have in GB
+ `disk`: `25` (number, required) - Disk size a server of this type will have in GB
+ `prices` (array[ServerTypeLocationPrice], required, fixed-type) - Prices in different Locations
+ `storage_type`: `local` (enum[string], required) - Type of server boot drive. Local has higher speed. Network has better availability.
    + Members
        + `local`
        + `network`
+ `cpu_type`: `shared` (enum[string], required) - Type of cpu.
    + Members
        + `shared`
        

## Location
+ `id`: `1` (number, required) - ID of the location
+ `name`: `fsn1` (string, required) - Unique identifier of the location
+ `description`: `Falkenstein DC Park 1` (string, required) - Description of the location
+ `country`: `DE` (string, required) - ISO 3166-1 alpha-2 code of the country the location resides in
+ `city`: `Falkenstein` (string, required) - City the location is closest to
+ `latitude`: `50.476120` (number, required) - Latitude of the city closest to the location
+ `longitude`: `12.370071` (number, required) - Longitude of the city closest to the location

## Datacenter
+ `id`: `1` (number, required) - ID of the datacenter
+ `name`: `fsn1-dc8` (string, required) - Unique identifier of the datacenter
+ `description`: `Falkenstein 1 DC 8` (string, required) - Description of the datacenter
+ `location` (Location, required) - Location where the datacenter resides in
+ `server_types` (DatacenterServerTypes, required) - The server types the datacenter can handle

## DatacenterServerTypes
+ `supported`: 1,2,3 (array[number], required, fixed-type) - IDs of server types that are supported in the datacenter
+ `available`: 1,2,3 (array[number], required, fixed-type) - IDs of server types that are supported and for which the datacenter has enough resources left

## Image
+ `id`: `4711` (number, required) - ID of the image
+ `type`: `snapshot` (enum[string], required) - Type of the image
    + Members
        + `system`
        + `snapshot`
        + `backup`

+ `status`: `available` (enum[string], required) - Whether the image can be used or if it's still being created
    + Members
        + `available`
        + `creating`
+ `name`: `ubuntu-16.04` (string, nullable, required) - Unique identifier of the image. This value is only set for system images.
+ `description`: `Ubuntu 16.04 Standard 64 bit` (string, required) - Description of the image
+ `image_size`: `2.3` (number, required, nullable) - Size of the image file in our storage in GB. For snapshot images this is the value relevant for calculating costs for the image.
+ `disk_size`: `10` (number, required) - Size of the disk contained in the image in GB.
+ `created`: `2016-01-30T23:50+00:00` (string, required) - Point in time when the image was created (in ISO-8601 format)
+ `created_from` (ImageCreator, required, nullable) - Information about the server the image was created from
+ `bound_to` (number, required, nullable) - ID of server the image is bound to. Only set for images of type `backup`.
+ `os_flavor`: `ubuntu` (enum[string], required) - Flavor of operating system contained in the image
    + Members
        + `ubuntu`
        + `centos`
        + `debian`
        + `fedora`
        + `unknown`
+ `os_version`: `16.04` (string, required, nullable) - Operating system version
+ `rapid_deploy`: `false` (boolean) - Indicates that rapid deploy of the image is available
+ `protection` (ImageProtection, required) - Protection configuration for the image
+ `deprecated`: `2018-02-28T00:00:00+00:00` (string, required, nullable) - Point in time when the image is considered to be deprecated (in ISO-8601 format)

## Snapshot Image (Image)
+ `name` (string, nullable, required) - Unique identifier of the image. This value is only set for system images.
+ `type`: `snapshot` (enum[string], required) - Type of the image
    + Members
        + `snapshot`
        + `system`
        + `backup`

## ImageCreator (object)
+ `id`: `1` (number, required) - ID of the server the image was created from
+ `name`: `Server` (string, required) - Server name at the time the image was created

## Action
+ `id`: `1337` (number, required) - ID of the action
+ `command`: `create_server` (string, required) - Command executed in the action
+ `status`: `success` (enum[string], required) - Status of the action
    + Members
        + `success`
        + `running`
        + `error`
+ `progress`: `100` (number, required) - Progress of action in percent
+ `started`: `2016-01-30T23:55:00+00:00` (string, required) - Point in time when the action was started (in ISO-8601 format)
+ `finished`: `2016-01-30T23:56:00+00:00` (string, required, nullable) - Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null.
+ `resources` (array[ResourceRef], required, fixed-type) - Resources the action relates to
+ `error` (ActionErrorRef, required, nullable) - Error message for the action if error occured, otherwise null.

## ActionRunning (Action)
+ `status`: `running` (enum[string], required) - Status of the action
    + Members
        + `running`
        + `success`
        + `error`
+ `progress`: `0` (number, required) - Progress of action in percent
+ `started`: `2016-01-30T23:50+00:00` (string, required) - Point in time when the action was started (in ISO-8601 format)
+ `finished` (string, required, nullable) - Point in time when the action was finished (in ISO-8601 format). Only set if the action is finished otherwise null.
+ `resources` (array[ResourceRef], required, fixed-type) - Resources the action relates to
+ `error` (ActionErrorRef, required, nullable) - Error message for the action if error occured, otherwise null.

## ActionErrorRef
+ `code`: `action_failed` (string, required) - Fixed machine readable code
+ `message`: `Action failed` (string, required) - Humanized error message.


## ResourceRef
+ `id`: `42` (number, required) - ID of resource referenced
+ `type`: `server` (string, required) - Type of resource referenced

## ErrorRef
+ `code`: `server_create_failed` (string, required) - Fixed machine readable code
+ `message`: `Creating the server failed` (string, required) - Humanized error message.

## FloatingIP
+ `id`: `4711` (number, required) - ID of the Floating IP
+ `description`: `Web Frontend` (string, nullable, required) - Description of the Floating IP
+ `ip`: `131.232.99.1` (string, required) - IP address of the Floating IP
+ `type`: `ipv4` (enum[string], required) - Type of the Floating IP
    + Members
        + `ipv4`
        + `ipv6`
+ `server`: `42` (number, nullable, required) -Id of the Server the Floating IP is assigned to, null if it is not assigned at all.
+ `dns_ptr` (array[FloatingIPrDNS], required, fixed-type) - Array of reverse DNS entries
+ `home_location` (Location, required) - Location the Floating IP was created in. Routing is optimized for this location.
+ `blocked`: `false` (boolean, required) - Whether the IP is blocked
+ `protection` (FloatingIPProtection, required) - Protection configuration for the Floating IP

## FloatingIPrDNS (object)
+ `ip`: `2001:db8::1` (string, required) - Single IPv4 or IPv6 address
+ `dns_ptr`: `server.example.com` (string, required) - DNS pointer for the specific IP address

## ISO (object)
+ `id`: 4711 (number, required) - ID of the ISO
+ `name`: `FreeBSD-11.0-RELEASE-amd64-dvd1` (string, required, nullable) - Unique identifier of the ISO. Only set for public ISOs
+ `description`: `FreeBSD 11.0 x64` (string, required) - Description of the ISO
+ `type`: `public` (enum[string], required) - Type of the ISO
    + Members
        + `public`
        + `private`
+ `deprecated`: `2018-02-28T00:00:00+00:00` (string, required, nullable) - ISO 8601 timestamp of deprecation, null if ISO is still available. After the deprecation time it will no longer be possible to attach the ISO to servers.


## Metrics Generic (object)
  + `start`: `2017-01-01T00:00:00+00:00` (string, required) - Start of period of metrics reported (in ISO-8601 format)
  + `end`: `2017-01-01T23:00:00+00:00` (string, required) - End of period of metrics reported (in ISO-8601 format)
  + `step`: `60` (number, required) - Resolution of results in seconds.
  + `time_series` (object, required) - Hash with timeseries information, containing the name of timeseries as key
    + *name_of_timeseries (string, required)* (Metrics Values, required)

## Metrics Values (object)
+ `values` (array, required, fixed-type) - Metrics Timestamps with values
  + 1435781470.622, 42 (array)


## IPv4 (object)
+ `ip`: `1.2.3.4` (string, required) - IP address (v4) of this server.
+ `blocked`: `false` (boolean, required) - If the IP is blocked by our anti abuse dept
+ `dns_ptr`: `server01.example.com` (string, required) - Reverse DNS PTR entry for the IPv4 addresses of this server.

## IPv6 (object)
+ `ip`: `2001:db8::/64` (string, required) - IP address (v4) of this server.
+ `blocked`: `false` (boolean, required) - If the IP is blocked by our anti abuse dept
+ `dns_ptr` (array[IPv6rDNS], required, fixed-type, nullable) - Reverse DNS PTR entries for the IPv6 addresses of this server, `null` by default.

## IPv6rDNS (object)
+ `ip`: `2001:db8::1` (string, required) - Single IPv6 address of this server for which the reverse DNS entry has been set up.
+ `dns_ptr`: `server.example.com` (string, required) - DNS pointer for the specific IP address.

# Public Network (object)
+ `ipv4` (IPv4, required) - IP address (v4) and its reverse dns entry of this server.
+ `ipv6` (IPv6, required) - IPv6 network assigned to this server and its reverse dns entry.
+ `floating_ips`: `478` (array[number], required, fixed-type) - IDs of floating IPs assigned to this server.


## NetGrossPrice (object)
+ `net`: `1` (string, required) - Price without VAT
+ `gross`: `1.19` (string, required) - Price with VAT added

## ImagePrice (object)
+ `price_per_gb_month` (NetGrossPrice, required)

## FloatingIPPrice (object)
+ `price_monthly` (NetGrossPrice, required)

## TrafficPrice (object)
+ `price_per_tb` (NetGrossPrice, required)

## BackupPrice (object)
+ `percentage`: `20` (string, required) - Percentage by how much the base price will increase

## ServerTypeLocationPrice (object)
+ `location`: `fsn1` (string, required) - Name of the location the price is for
+ `price_hourly` (NetGrossPrice, required) - Hourly costs for a server type in this location
+ `price_monthly` (NetGrossPrice, required) - Monthly costs for a server type in this location

## ServerTypesPricing (object)
+ `id`: 4 (number, required) - ID of the server type the price is for
+ `name`: `CX11` (string, required) - Name of the server type the price is for
+ `prices` (array[ServerTypeLocationPrice], required, fixed-type) - Server type costs per location

## Price (object)
+ `currency`: `EUR` (string, required) - Currency the returned prices are expressed in, coded according to ISO 4217.
+ `vat_rate`: `19.00` (string, required) - The VAT rate used for calculating prices with VAT
+ `image` (ImagePrice, required) - The cost of one 1GB Image for the full month.
+ `floating_ip` (FloatingIPPrice, required) - The cost of one floating IP per month.
+ `traffic` (TrafficPrice, required) - The cost of additional traffic per GB
+ `server_backup` (BackupPrice, required) - Will increase base server costs by specific percentage.
+ `server_types` (array[ServerTypesPricing], required, fixed-type) - Costs of server types per location and type

## ServerProtection
+ `delete`: `false` (boolean, required) - If true, prevents the server from being deleted
+ `rebuild`: `false` (boolean, required) - If true, prevents the server from being rebuilt

## FloatingIPProtection
+ `delete`: `false` (boolean, required) - If true, prevents the Floating IP from being deleted

## ImageProtection
+ `delete`: `false` (boolean, required) - If true, prevents the snapshot from being deleted


