# sgmock

SendGrid API v3 mock for using in tests

## API

 * /v3/mail/send -- mock for SendGrid API v3 [send email endpoint]
                    (https://sendgrid.com/docs/API_Reference/api_v3.html)
 * /ctl/list -- list all sent messages, message format is the same like for
                send operation
 * /ctl/clear -- clear list of sent messages

Examples

Send message:

```http
POST http://dockerhost:9001/v3/mail/send
Content-Type: application/json
Accept: application/json
Authorization: Bearer key

{
    "personalizations": [
        {
            "subject": "test",
            "to": [
                {
                    "email": "user@mail.com"
                }
            ]
        }
    ],
    "from": {
        "email": "sender@mail.com"
    },
    "content": {
        "type": "text/plain",
        "value": "Hello!"
    }
}
```

List messages:

```http
GET http://dockerhost:9001/ctl/list
```

Clear messages:

```http
POST http://dockerhost:9001/ctl/clear
```

## Docker

Image: vbogretsov/0.1.0

Port: 9001

Environment variables: ${SGMOCK_KEY} -- test value for SendGrid API key

## License

See the [LICENSE](https://github.com/vbogretsov/sgmock/blob/master/LICENSE) file.
