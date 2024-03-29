# Go API for Projects
REST API written in Go that displays my personal projects as a JSON object to curious callers. 

# Usage

## List all projects
### Definition
`GET /projects`

### Response
- `200 OK` on success

### Get All Projects Request Example (Python)
```python
  import requests
  import json

  r = requests.get("http://api.davidamante.com/projects") 
  data = json.dumps(r.json(), sort_keys=True, indent=4) # format json to be more readable
  print(data)
```

### Response
```json
{
  "Projects": [
    {
      "id": 31,
      "title": "Password Generator",
      "link": "https://github.com/aManNamedTed/arduino/tree/master/22e60",
      "date": "2019-08-17T00:00:00.000Z",
      "description": "Generates a single 32 character password using various sensors and a base-93 alphabet. Over 200 bits of security!",
      "technologies": [
        "Arduino", 
        "C++"
      ],
      "categories": [
        "Embedded Systems", 
        "Security", 
        "Application", 
        "Software"
      ]
    },
    {
      "id": 30,
      "title": "Personal Website",
      "link": "https://davidamante.com",
      "date": "2019-07-19T00:00:00.000Z",
      "description": "Version 4.0 of my personal website that serves as my professional portfolio. (the website you are currently viewing!)",
      "technologies": [
        "React", 
        "Javascript", 
        "Firebase", 
        "HTML", 
        "CSS"
      ],
      "categories": [
        "Web", 
        "Mobile", 
        "Software"
      ]
    },
    {
      "id": 29,
      "title": "Automated Email Sender",
      "link": "https://github.com/aManNamedTed/personalprojects/blob/master/auto_email_sender",
      "date": "2019-02-01T00:00:00.000Z",
      "description": "Google Apps Script that sends graded assignment scores to a list of student emails via Google Sheets.",
      "technologies": [
        "Javascript"
      ],
      "categories": [
        "Developer Tools", 
        "Application", 
        "Back-End", 
        "Software"
      ]
    }, 
    {
      etc...
    }
  ]
}
```

## List a single project
### Definition
`GET /projects/<id>`

### Response
- `200 OK` on success
- `204 No Content` if project does not exist

### Get Single Project Request Example (Python)
```python
  import requests
  import json
  
  r = requests.get("http://api.davidamante.com/projects/31")
  data = json.dumps(r.json(), sort_keys=True, indent=4) # format json to be more readable
  print(data)
```  

### Response
```json
{
  "id": 31,
  "title": "Password Generator",
  "link": "https://github.com/aManNamedTed/arduino/tree/master/22e60",
  "date": "2019-08-17T00:00:00.000Z",
  "description": "Generates a single 32 character password using various sensors and a base-93 alphabet. Over 200 bits of security!",
  "technologies": [
    "Arduino", 
    "C++"
  ],
  "categories": [
    "Embedded Systems", 
    "Security", 
    "Application", 
    "Software"
  ]
}
```
