# Go webservice package
The goal of this package is to provide common convenience types, functionality, and guidance on structuring a new web service/API. 

Providing a consistently-structured message back to the client, helps with standardizing deserialization and error-checking on the client side.

Note: *design descisions highly opinionated*

# Responses
## Payload
The payload is a standardized container for responses from the server. It is a contract that guarantees deserializability on the client.
- Status: A status message for the consumer to understand the outcome of a request.
- Code: A status code provided for convenience. Generally, it should match the HTTP code. E.g. 200 for OK, 400 for a client error, etc.
- Messages: An array of messages the client can optionally check for. Can be used for displaying to the client GUI, detailed tracing, etc.   
- Data: The data requested by the client. This can be the requested resource, or any data the client expects in return for the request.
- Errors: see below.

## ApiError
The api error is a structure that provides the implementor of the API additional context when something goes wrong.
- ReferenceCode: This can be used to reference more specific documentation on the detailed error, or for the client to act upon a specific scenario.
- Message: Summary of *what* went wrong **and** *why*