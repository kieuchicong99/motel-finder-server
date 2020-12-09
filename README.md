# 1. Server side

> Golang environment development
>
> - Gin framework.
> - Mongodb : database
> - Swagger : documentaion api
>
> Run command:
> sure that work directory is **react-go-heroku/server**

```bash
 swag init // generator api doc
 go run main.go
```

# 2. Client side

> Nodejs enviroment development
>
> - Reactjs framework.
> - AntDesign: library use for layout, form, ...., etc
> - Redux
>
> Run command:

```bash
 npm install  or yarn install
 npm start or yarn start
```

# 3. Deploy

> Docker enviroment development
>
> - Docker image, docker hub
> - heroku container deploy
>
> Heroku deployment command usage

```bash
heroku login
heroku container:login
heroku create go-react-heroku
heroku stack:set container
heroku container:push web
heroku config:set PORT=8080
heroku container:release web

```
