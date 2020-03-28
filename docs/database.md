# Database

## Collections

#### events
A monitoring event. Once inserted, an event should not change. The event with the latest timestamp and same dedup key should be used.
```
{
    "uid":"",
    "dedup_key":"",
    "timestamp": 1234123,
    "state": "trigger",
    "severity":"high",
    "message":"This is the event",
    "details": {}
}
```

#### event_metadata
Metadata information that can be changed/computed. Tied to event via the dedup_key.
```
{
    "dedup_key": "",
    "count":5,
    "event_group_id":"",
    "ticket_number":""
}
```

#### blackouts
Hides/silences incoming or existing events.
```
{
    "name":"",
    "description": "",
    "created_by":"",
    "type":"maintenance",
    "ticket_ref":"",
    "start_ts"123123,
    "stop_ts":123123
}
```


#### config
Config objects for the platform

```
{
    "type":"script_list",
    "scripts": [
        "enrich_with_server_data",
        "add_service_a_business_logic",
        "verify_maintenance"
    ]
}
```

```
{
    "type":"script",
    "metadata":{
        "name": "enrich_with_server_data",
        "description":"What the script does",
        "type":"parse",
        "initial-author":"your-email",
        "per_message_timeout_ms": 10000,
        "rate_limit_per_minute": -1
    },
    "script": "b64-encoded-js"
}
```
