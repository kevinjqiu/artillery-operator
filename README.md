```
    _         _   _ _ _                 
   / \   _ __| |_(_) | | ___ _ __ _   _ 
  / _ \ | '__| __| | | |/ _ \ '__| | | |
 / ___ \| |  | |_| | | |  __/ |  | |_| |
/_/   \_\_|   \__|_|_|_|\___|_|   \__, |
                                  |___/ 
  ___                       _             
 / _ \ _ __   ___ _ __ __ _| |_ ___  _ __ 
| | | | '_ \ / _ \ '__/ _` | __/ _ \| '__|
| |_| | |_) |  __/ | | (_| | || (_) | |   
 \___/| .__/ \___|_|  \__,_|\__\___/|_|   
      |_|                                 

```

A kubernetes operator for managing artillery (artillery.io) jobs

Sample manifest:

```yaml
apiVersion: artillery.idempotent.ca/v1alpha1
kind: Artillery
spec:
  testScript:
    config:
    target: "https://staging1.local"
    phases:
        - duration: 60
        arrivalRate: 5
        - duration: 120
        arrivalRate: 5
        rampTo: 50
        - duration: 600
        arrivalRate: 50
    payload:
        path: "keywords.csv"
        fields:
        - "keywords"
    scenarios:
    - name: "Search and buy"
        flow:
        - post:
            url: "/search"
            body: "kw={{ keywords }}"
            capture:
                json: "$.results[0].id"
                as: "id"
        - get:
            url: "/details/{{ id }}"
        - think: 3
        - post:
            url: "/cart"
            json:
                productId: "{{ id }}"
```