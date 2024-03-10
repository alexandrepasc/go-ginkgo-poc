package data

const GetCatfactsId = `{
    "type": "object",
    "required": [
        "status",
        "_id",
        "user",
        "text",
        "__v",
        "source",
        "updatedAt",
        "type",
        "createdAt",
        "deleted",
        "used"
    ],
    "properties": {
        "status": {
            "type": "object",
            "required": [
                "verified",
                "sentCount"
            ],
            "properties": {
                "verified": {
                    "type": "boolean",
                    "enum": [
                        true
                    ]
                },
                "sentCount": {
                    "type": "integer",
                    "enum": [
                        1
                    ]
                }
            }
        },
        "_id": {
            "type": "string",
            "enum": [
                "58e008780aac31001185ed05"
            ]
        },
        "user": {
            "type": "object",
            "required": [
                "name",
                "_id",
                "photo"
            ],
            "properties": {
                "name": {
                    "type": "object",
                    "required": [
                        "first",
                        "last"
                    ],
                    "properties": {
                        "first": {
                            "type": "string",
                            "enum": [
                                "Kasimir"
                            ]
                        },
                        "last": {
                            "type": "string",
                            "enum": [
                                "Schulz"
                            ]
                        }
                    },
                    "additionalProperties": false
                },
                "_id": {
                    "type": "string",
                    "enum": [
                        "58e007480aac31001185ecef"
                    ]
                },
                "photo": {
                    "type": "string",
                    "enum": [
                        "https://lh6.googleusercontent.com/-BS_rskGd3kA/AAAAAAAAAAI/AAAAAAAAADg/yAxrX9QabMg/photo.jpg?sz=200"
                    ]
                }
            },
            "additionalProperties": false
        },
        "text": {
            "type": "string",
            "enum": [
                "Owning a cat can reduce the risk of stroke and heart attack by a third."
            ]
        },
        "__v": {
            "type": "integer",
            "enum": [
                0
            ]
        },
        "source": {
            "type": "string",
            "enum": [
                "user"
            ]
        },
        "updatedAt": {
            "type": "string",
            "enum": [
                "2020-08-23T20:20:01.611Z"
            ]
        },
        "type": {
            "type": "string",
            "enum": [
                "cat"
            ]
        },
        "createdAt": {
            "type": "string",
            "enum": [
                "2018-03-29T20:20:03.844Z"
            ]
        },
        "deleted": {
            "type": "boolean",
            "enum": [
                false
            ]
        },
        "used": {
            "type": "boolean",
            "enum": [
                false
            ]
        }
    },
    "additionalProperties": false
}`
