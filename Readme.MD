# FeatureToggle Golang Project

A feature toggle is a technique in software development that attempts to provide an alternative to maintaining multiple branches in source code, such that a software feature can be tested even before it is completed and ready for release. A feature toggle is used to hide, enable or disable the feature during runtime.

This is a light, simple and fast microservice application builded with golang. It uses a mongodb database to store the features.

### How to use

To run the application, you just need to use the docker-compose to build the infrastructure and then start the golang aplication using "go run main.go".

### How it works

Feature have a very simple struct:
{
    "name": "featureName"
    "clients": [
        "yourClientId1",
        "yourClientId2"
    ]
}

Endpoints:

- GetAll features - GET .../feature
- Insert feature - POST .../feature
- Validate if a feature is enabled for a client - GET ../feature/{featureName}/{client}

### To do's

- Update feature clients.
- Request Authentication.
