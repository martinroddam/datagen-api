Data Generator API
==========

This api returns a given number of randomly generated personal data records with some basic statistics.

Table of Contents
-----------------

  * [Requirements](#requirements)
  * [Usage](#usage)
  * [License](#license)


Requirements
------------

Data Generator API requires the following to run:

  * [Golang][golang] 1.10+


Usage
-----

Get the application:

```sh
go get github.com/martinroddam/datagen-api
```

Run the application:

```sh
datagen-api
```

Hit the endpoint:

```
GET http://localhost:8080/generate/{numRecords}
```

Example output:

```
GET http://localhost:8080/generate/1
```

```
    {
        "statistics": {
            "countryCounts": {
                "United Kingdom": 0,
                "United States of America": 0
            },
        "recordCount": 1
        },
        "records": [
            {
                "fullName": "MISS JARED O'REILLY",
                "address": "1022 Orn Drives",
                "city": "East Missourimouth",
                "state": "California",
                "postcode": "53281-3066",
                "country": "Sri Lanka",
                "email": "antwan@dickinson.name",
                "homePhone": "(630)148-3520 x66725",
                "mobilePhone": "1-563-322-0766 x767"
            }
        ]
    }
```

License
-------

Copyright &copy; 2018, Martin Roddam

[golang]: https://golang.org/