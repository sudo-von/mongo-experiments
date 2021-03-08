### Exercises

<hr>
Connect to the console at localhost. Try typing some JavaScript expressions.

```
Tell me how many seconds there are in a week

solution: 
60 * 60 * 24 * 7;

Tell me how many weeks there are in a human lifetime of 80 years.

solution: 
52 * 80;
```

<hr>

```
Use the use command to connect to a new database (If it doesn't exist, Mongo will create it when you write to it).

solution: 
use test;
```

<hr>

```
Use db.createCollection to create a collection. I'll leave the subject up to you.

solution:
db.users;

Run show dbs and show collections to view your database and collections.

solution:
show dbs;
show collections;
```

<hr>

```
Insert a couple of documents into your collection. I'll leave the subject matter up to you, perhaps cars or hats.

solution:
db.users.insert(
    {
        "username": "von",
        "name": "sudo_von",
        "active": true,
        "animes_watched" : 99
    }
)
db.users.insert(
    {
        "username": "nov",
        "name": "sudo_nov",
        "active": false,
        "animes_watched" : 100
    }
)
```

<hr>

```
Use find() to list them out.

solution:
db.users.find()
```

<hr>