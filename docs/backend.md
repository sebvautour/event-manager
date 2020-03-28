# Backend Service
Main backend service that process the monitoring events. It is responsible for the following:
- Applies business logic to events
- Performs external requests for enrichment data
- Correlates events together

The backend is the only service that **writes any modifications** to the backend. It has a cache to store temporary data such as for enrichment.

The overall functionality the backend should provide is the following:
 - consume event and action messages from a message bus
 - ability to run a combination of scripts, in a certain order, against events or actions.
 - Provide functions that can be used by the scripts, such as:
   - HTTP request
   - Simple cache (GET, GETALL, SET, DELETE)
   - API, interactions with the Event Manager API endpoints
   - SaveEvent function to save an event to the DB
   - Log
 - Audit and monitor script executions
 - Each event/action should also have a context attached to it. Allowing for extra metadata information to be shared across scripts.

 ## Files
  * scripts mapping config (scripts.json) - provides the script filename and the order they should be run

    ```
    [
        "enrich_with_server_data",
        "add_service_a_business_logic",
        "verify_maintenance"
    ]
    ```


  * Script file - Javascript script containing the business logic. Needs to have a script_metadata function that returns the following object:
  
    ```
    {
        "name": "enrich_with_server_data",
        "description":"What the script does",
        "type":"parse",
        "initial-author":"your-email",
        "per_message_timeout_ms": 10000,
        "rate_limit_per_minute": -1
    }
    ```
