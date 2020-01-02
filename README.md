# Dev Challenge for Go developers - GRPC

Dear most valuable applicant! This challenge is designed to see your way of thinking, your level of knowledge in the subject and your definition of "complete".

## Outline

The task is to create a GPRC microservice in Go, that behaves as a simple in-memory key-value storage.

Payload will be simple structure – a message and a numeric value –, indexed by numbers (IDs).

Example (not actual output):

```
1 => ( "message": "Hello World", "value": 0 )
2 => ( "message": "GPMD", "value": 123.456 )
```

No initial values are required, the example above is just for the understanding of the possible data types.

Go is an errors first language – please use this approach to write a safe program.

## The small details we are happy to see

* The right balance of briefness and clarity
* Testing – we prefer TDD for non-spike code
* Avoiding data races
* Go module
* Line of sight (Align the happy path to the left edge)
* No fluff – use what's necessary
* Go vet saying OK
