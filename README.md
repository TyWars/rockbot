# rockbot
rockbot


# Run MongoDB container

https://hub.docker.com/r/_/mongo/
docker run -p 27017:27017 --name bot-mongo -d mongo

 docker exec -it some-mongo mongo admin
connecting to: admin
> db.createUser({ user: 'rockbot', pwd: 'rockbot', roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] });
Successfully added user: {
    "user" : "jsmith",
    "roles" : [
        {
            "role" : "userAdminAnyDatabase",
            "db" : "admin"
        }
    ]
}