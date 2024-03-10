package data

const GetCatfactsList =`{
    "type": "array",
    "items": {
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
                        "type": "boolean"
                    },
                    "sentCount": {
                        "type": "integer"
                    }
                }
            },
            "_id": {
                "type": "string"
            },
            "user": {
                "type": "string"
            },
            "text": {
                "type": "string"
            },
            "__v": {
                "type": "integer"
            },
            "source": {
                "type": "string"
            },
            "updatedAt": {
                "type": "string"
            },
            "type": {
                "type": "string"
            },
            "createdAt": {
                "type": "string"
            },
            "deleted": {
                "type": "boolean"
            },
            "used": {
                "type": "boolean"
            }
        },
        "additionalProperties": false
    }
}`
