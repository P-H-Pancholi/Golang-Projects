## Technologies needed
1. Go
2. Postgres

## Steps to install gator
1. Clone this repo.
2. Run go install . on root of this project
3. setup .gatorconfig.json at the home folder of your user
4. Content of .gatorconfig.json

```

{
    "db_url":"postgres://postgres:postgres@localhost:5433/gator?sslmode=disable",
    "current_user_name":"kahya"
}

```
Please add the database url of postgres in dburl section
you can now use gator command in your command line

## Using gator

1. To register user:
```
gator register <user_name>
```
2. To login as a user:
```
gator login <user_name>
```
3. To list all users:
```
gator users
```