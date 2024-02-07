# setu-engine

# Usage 
## Run the [whatsapp-ki-maya](https://github.com/veryshyjelly/whatsapp-ki-maya) server
or
## Run the [telegram-ki-maya](https://github.com/veryshyjelly/telegram-ki-maya) server

### Get their IP address

## Using Docker
- Build the docker image
```bash
docker build -t setu-engine .
```
- Run the docker container
```bash
docker run -d -p 8070:8070 --name setu-engine setu-engine
```

### Bridging the chats

### Obtain chat id from whatsapp 
- Send .id message in the chat with the bot present 
### Obtain chat id from telegram
- Send /id message in the chat with the bot present

### Post the bridge request using postman
- Create a new request
- Set the request type to `POST`
- Set the request URL to `http://localhost:8070/bridge`
- Set the request body to 
```json
{
    "first_chat_id": "first_chat_id",
    "first_url" : "ip_address_of_first_host",
    "second_chat_id": "second_chat_id",
    "second_url" : "ip_address_of_second_host"
}
```

### Notes
- You can bridge any two chats from any two hosts