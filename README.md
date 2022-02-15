# Golang & ReactJS Chat

[![GolangCI](https://github.com/pmokeev/web-chat/actions/workflows/GolangCI.yml/badge.svg?branch=master)](https://github.com/pmokeev/web-chat/actions/workflows/GolangCI.yml)

This is a Golang & ReactJS chat application powered by WebSockets, which provides the main functions: registration on the platform, sending messages to the general chat.

![](https://i.ibb.co/RBkGfcF/1.jpg)
![](https://i.ibb.co/jvD55fS/2.jpg)
![](https://i.ibb.co/1z3KRb0/bz-Nw1-MWA5s4.jpg)

## Installation

Create an .env file at the root of the project with the content:

```environment
JWTSecretKey=-//-
POSTGRES_USER=-//-
POSTGRES_PASSWORD=-//-
POSTGRES_DB=-//-
```

The application is packaged in [docker](https://www.docker.com/) containers. You must also have docker-compose installed in order to run the application. Command to run the application:

```bash
sudo docker-compose up -d
```

## Features

- Design with bootstrap
- Registration with JWT token
- Communication between the server and the client in the chat occurs thanks to WebSockets
- Changing a user's password in a profile
- Other awesome features yet to be implemented

**To Do:**

- Emoji support
- Possibility to create closed rooms
- Private messages by command /msg [user]

## License
[MIT](https://choosealicense.com/licenses/mit/)