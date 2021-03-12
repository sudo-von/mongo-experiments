#### Welcome to the Mongo Workbook

This little book accompanies the Mongo course taught over 2 days. Each exercise builds on the last, so it's best to work through these exercises in order.

We'll start off with the helper functions, find, count and distinct, then move into the aggregate pipeline and on to map reduce, finally integrating with Node to produce a service oriented architecture.

#### Welcome to MongoDB

Mongo DB is a highly scalable, NoSQL, schema free data store.

#### Here's why it's great

<ul>
    <li>It stores data as JSON so you can use the same data clientside and serverside. This means you write almost no wiring code, everything just works.
    </li>
    <li>
    Flexible Schema. If your requirements change, you can adapt.
    </li>
    <li>Unstructured data - you can store and retrieve unstructured data easily. It's just JSON. Not every document in a collection needs the same fields.</li>
    <li>Denormalised data - Group related content in a single document.</li>
    <li>Clean and simple API - Mongo is nice to talk to.</li>
</ul>

#### Here's why you might not like it

<ul>
    <li>Denormalised data means no joins. If your data is highly relational, Mongo is not your baby. Your data is organised into collections. If you need data from more than collection, you need to hit the database more than once.</li>
    <li>Flexible schema means no built in data validation. Your data is validated at the application tier. The database is dumb and will store whatever your application gives it, even junk and typos.</li>
    <li>Bugs - Mongo is new and there are still issues in the tracker. Not bad bugs, but occasionally things don't work as you might expect.</li>
    <li>No transactions - A SQL database allows you to bundle multiple writes into a transaction. If one write fails the whole transaction is rolled back. Mongo lacks this feature, writes are small and atomic. If you need a transaction you must build it yourself.</li>
    <li>Theoretical data loss - Mongo scales using a technique called sharding. It creates slaves that mirror data written to the master. If the master goes down before data is mirrored you may lose recent commits depending on your settings.</li>
</ul>

#### When you should use it

Mongo represents data as a tree. If your data is tree shaped, or can be made tree shaped, Mongo is great. If your data is a web or a network which can't be flattened out, you likely have relational data, and Mongo is perhaps not for you this time.

If you have unstructured data to store which can be represented as a series of nested lists Mongo will make your life more enjoyable.

Say you have a webpage full of widgets, and each of those widgets can contain arbitrary information. This is a semi-structured tree and Mongo might be a very good choice.

If you have big customer data to store, and each customer record contains lists of communications, subscriptions, etc, the data is tree shaped, and Mongo would again be a good chice.

If you have big data and you want to query it in interesting and complex ways, pulling useful aggregated data out the other side in suprisingly short timeframes, Mongo is perfect.

On the other hand, if your data looks like a web: comments, purchases, kittens, customers, sharks, exploding hats, etc, all linking to each other in a web, then you have relational data, and you may wish to stick with a relational database like Postgres, MySQL or MS SQL Server.

#### Why is it fast?

Mongo manages to be so fast because it does less. There's no magical difference in the architecture that makes it fast, it just has a simplified streamlined query language that is easier to optimise.

#### Connecting to the terminal

We connect to the Mongo terminal using the mongo command

```
```

By default Mongo will connect to localhost.

We can connect to a remote server by passing arguments, like so:

```
```

Once we connect to a Mongo instance we can type JavaScript directly into the console. We can create variables, do maths, write JSON.

<hr>

#### Exercise - connect to a console

Connect to the console at localhost. Try typing some JavaScript expressions.

```
Tell me how many seconds there are in a week
```

```
Tell me how many weeks there are in a human lifetime of 80 years.
```

<hr>

#### Creating a database

We can switch to a database in Mongo with the use command.

```
```

This will switch to writing to the petshop database. It doesn't matter if the database doesn't exist yet. It will be brought into existence when you first write a document to it.

You can find which database you are using simply by typing db. You can drop the current database and everything in it using db.dropDatabase.

```
```

<hr>

#### Exercise - Create a database

```
Use the use command to connect to a new database (If it doesn't exist, Mongo will create it when you write to it).

That was easy wasn't it. Don't worry, it gets a bit harder.
```

<hr>

#### Collections

Collections are sets of (usually) related documents. Your database can have as many collections as you like.

Because Mongo has no joins, a Mongo query can pull data from only one collection at a time. You will need to take this into account when deciding how to arrange your data.

You can create a collection using the createCollection command.

```
```

Collections will also be created automatically. If you write a document to a collection that doesn't exist that collection will be brought into being for you.

View your databases and collections using the show command, like this:

```
```

<hr>

#### Exercise - Create a collection

```
Use db.createCollection to create a collection. I'll leave the subject up to you.
```

```
Run show dbs and show collections to view your database and collections.
```

<hr>

#### Documents

Documents are JSON objects that live inside a collection. They can be any valid JSON format, with the caveat that they can't contain functions.

The size limit for a document is 16Mb which is more than ample for most use cases.

#### Creating a document

You can create a document by inserting it into a collection

```
```

<hr>

#### Exercise - Create some documents

```
Insert a couple of documents into your collection. I'll leave the subject matter up to you, perhaps cars or hats.
```

<hr>

#### Finding a document

You can find a document or documents matching a particular pattern using the find function.

If you want to find all the mammals in the mammals collection, you can do this easily.

```
```

<hr>

#### Exercise - documents

```
Use find() to list them out.
```

<hr>

#### Finding documents

Mongo comes with a set of convenience functions for performing common operations. Find is one such function. It allows us to find documents by providing a partial match, or an expression.

#### Uses

You can use find to:

<ul>
    <li>Find a document by id.</li>
    <li>Find a user by email.</li>
    <li>Find a list of all users with the same first name.</li>
    <li>Find all cats who are more than 12 years old.</li>
    <li>Find all gerbils called 'Herbie' who are bald, have three or more eyes, and who have exactly 3 legs.</li>
</ul>

#### Limitations

You <b>can't</b> use find to chain complex operators. You can do a few simple things like counting, but if you want to real power you need the aggregate pipeline, which is actually not at all scary and is quite easy to use.

The Aggregate pipeline allows us to chain operations together and pipe a set of documents from one operation to the next.

#### Using find

You can use find with no arguments to list documents in a collection.

```
```

This will list all of the codes, 20 at a time.

You can get the same result by passing an empty object, like so:

```
```

#### Finding by ID

Assuming you know the object ID of a document. You can pull that document by id like so:

```
```

The _id field of any collection is automatically indexed.

IDs are 12 byte BSON objects, not Strings which is why we need the ObectId function. If you want to read more on ObjectId, you can do so here.

#### Finding by partial match

Say you have a list of users and you want to find by name, you might do:

```
```

You can match on more than one field:

```
```

You can match on numbers:

```
```

You also match using a regex (although be aware this is slow on large data sets):

```
```

<hr>

#### Exercise

We need to start out by inserting some data which we can work with.

```
Paste the following into your terminal to create a petshop with some pets in it

use petshop;
db.pets.insert({name: "Mikey", species: "Gerbil"});
db.pets.insert({name: "Davey Bungooligan", species: "Piranha"});
db.pets.insert({name: "Suzy B", species: "Cat"});
db.pets.insert({name: "Mikey", species: "Hotdog"});
db.pets.insert({name: "Terrence", species: "Sausagedog"});
db.pets.insert({name: "Philomena Jones", species: "Cat"});
```


```
Add another piranha, and a naked mole rat called Henry.
```

```
Use find to list all the pets. 
```

```
Find the ID of Mikey the Gerbil.
```

```
Use find to find all the gerbils.
```

```
Find all the creatures named Mikey.
```

```
Find all the creatures named Mikey who are gerbils.
```

```
Find all the creatures with the string "dog" in their species.
```

<hr>

#### Finding with Expressions and comparison queries

We have seen how we can find elements by passing Mongo a partial match, like so:
db.people.find({name: 'Yolanda Sasquatch'})

We can also find using expressions. We define these using JSON, like so:

```
```

We can use operators like this:

```
```

See the full list here:

http://docs.mongodb.org/manual/reference/operator/query/

<hr>

#### Exercise

Copy the following code into a Mongo terminal. It will create a collection of people, some of whom will have cats.

Optionally modify the code so that some people have piranhas, and some have dachshunds.

```
use people;
(function() {
  var names = [
    'Yolanda',
    'Iska',
    'Malone',
    'Frank',
    'Foxton',
    'Pirate',
    'Poppelhoffen',
    'Elbow',
    'Fluffy',
    'Paphat'
  ]
  var randName = function() {
    var n = names.length;
    return [
      names[Math.floor(Math.random() * n)],
      names[Math.floor(Math.random() * n)]
    ].join(' ');
  }
  var randAge = function(n) {
    return Math.floor(Math.random() * n);
  }
  for (var i = 0; i < 1000; ++i) {
    var person = {
      name: randName(),
      age: randAge(100)
    }
    if (Math.random() 0.8) {
      person.cat = {
        name: randName(),
        age: randAge(18)
      }
    }
    db.people.insert(person);
  };
})();
```

```
Use find to get all the people who are exactly 99 years old.
```

```
Find all the people who are eligible for a bus pass (people older than 65).
```

```
Find all the teenagers, greater than 12 and less than 20.
```

<hr>

#### $exists

We can use exists to filter on the existence of non-existence of a field. We might find all the breakfasts with eggs:

```
```

<hr>

#### Exercise - $exists

```
Find all the people with cats.
```

```
Find all the pensioners with cats.
```

```
Find all the teenagers with teenage cats.
```

<hr>

#### \$gt and \$lt

We can use \$gt and \$lt to find documents that have fields which are greater than or less than a value:

```
```

#### \$where

We can even filter using an arbitrary JavaScript expression using $where. This will allow us to compare two fields in a single document.

```
```

Here we find all the sandwiches with jam and peanut butter where the jam quotient outweighs the peanut butter.

Warning: It's easy to overuse $where since it appears to do everything with plain old JavaScript. \$where is eval-ing a JavaScript expression and as such is slow. Mongo can make no optimisations here, and must execute the JavaScript on every single document in the collection. Prefer the native operators where possible.

#### 
Exercise - $where

```
Use $where to find all the people whohave a cat.
```

```
Find all the people who are youngerthan their cats. Remember, not everyonehas a cat, so you will need to use aboolean && to filter out the non-catowners.
```

```
Does anyone have the same name as theircat? Re-run the insertion script tocreate more records until someone does.
```

<hr>

#### 
Projection

Find takes a second parameter which allows you to whitelist fields to pass into the output document. We call this projection.

You can choose fields to pass though, like so:

```
{
  ham: 4,
  eggs: 2
}
{
  cheese: 6,
  lime: 0.5
}
db.breakfast.find({}, {
  eggs: true,
  lime: true
});
```

This will yield:

```
{
  eggs: 2
},
{
  lime: 0.5
}
```

<hr>

#### Exercise - Tidy up your output

```
Use projection to format your array of people. 
We want only the names.
```

```
Output just the names of the people whoare 99 years old.
```

```
Output only the cats, like this:
{ "cat" : { "name" : "Fluffy Frank", "age" : 13 } }
```

When you output the cats, you will need to find only people who have cats, where cats $exists, or you will have gaps in your data.

<hr>

#### Aggregation

We can do much more complex projection, even creating new fields based on expressions using the aggregate pipeline. More on this in a bit.

#### Excluding the id field

You will notice that the ID field is always passed through project by default. This is often desirable, but you may wish to hide it, perhaps to conceal your implementation, or to keep your communication over the wire tight.

You can do this easily by passing _id: false:

```
db.breakfast.find({}, {
  eggs: true,
  lime: true,
  _id: false
});
```

<hr>

#### Exercise - remove the ids

```
List the cats. 
Remove the ids from the output.
```

<hr>
