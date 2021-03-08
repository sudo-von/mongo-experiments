# NoSQL databases

NoSQL databases have 4 large families:
<ul>
<li> Key value stores </li>
<li> Based on graphs </li>
<li> Wide-column stores </li>
<li> Document based </li>
</ul>

<b>Key value stores</b>: They store the information using keys and values. We use them to store cache, user session information or very simple things. They are very fast to consult but we can't use them in more complex cases where we need more special structures. The best example of these databases is Redis.

<b>Graph Databases</b>: They allow us to establish connections between our entities to make queries in a more efficient way than in relational databases (as well as Twitter or Medium where each publication has different relationships between its users, likes, etc...). For example: Neo4j or JanusGraph.

<b>Wide-column Stores</b>: They have a row key and a column key to make very fast queries and store large amounts of information, but modeling the data can become a bit complicated. We use them in Big Data, IoT, recommendation systems, among others. For example: Cassandra or HBase.

<b>Document Databases</b>: They allow us to store documents within collections, they have very good performance and flexibility that allows us to model real life cases in a simple and effective way. For example: MongoDB or CouchBase.

<hr>

# MongoDB

MongoDB is a free and open source Non-relational document based database that allows us to store a large number of documents in a distributed way. Mongo is also the name of the company that develops the code for this database.

One of its main characteristics is that it allows us to store our structures or documents in JSON format (not exactly JSON, but its something very similar, we will see it later) to have great flexibility when modeling real life situations.

Because its a distributed database, we can speak not of one but of several servers, which we know as the MongoDB Cluster. Thanks to this we obtain a great horizontal scalability (scalability in number of servers).

MongoDB is “Schema Less” which allows our documents to have different structures without affecting their operation, something that we cannot do with relational database tables.

<hr>

# MongoDB Atlas

We have several providers that allow us to use or rent MongoDB as a service and in this case we are going to use MongoDB Atlas because its developed by the same people who develop MongoDB.

MongoDB Atlas has the following features:

<ul>
<li> Automatic provisioning of clusters with MongoDB </li>
<li> High availability </li>
<li> Highly scalable </li>
<li> Safe </li>
<li> Available on AWS, GCP and Microsoft Azure </li>
<li> Easy monitoring and optimization </li>
</ul>

<hr>

#### Basic MongoDB commands

Run MongoDB.
<code>mongo</code>
Create and use a DB.
<code>use DB</code>
List all DBS.
<code>show dbs</code>
Create collection and insert a document.
<code>db.COLLECTION.insertOne({"name":"VoN"})</code>
List all documents from a collection.
<code>db.COLLECTION.find()</code>
List all collections.
<code>show collections</code>
Find one document.
<code>db.COLLECTION.findOne()</code>

<hr>

#### What are MongoDB drivers?

They are the libraries that we use to communicate our application with our database.

Without our drivers we cannot work with our database clusters.

#### How to add the drivers within our project?

We use a dependency manager. We add it in our dependency manager.

```
Python: python -m pip install pyongo
Node.js: npm install mongodb --save
GO: dep ensure -add go.mongodb.org/mongo-driver/mongo
```

#### Quick start across most languages

<ol>
<li> Create MongoDB connection. </li>
<li> Get MongoDB connection. </li>
<li> Access to MongoDB connection. </li>
<li> CRUD. </li>
</ol>

#### Databases, Collections and Documents in MongoDB

```
Database:

    Physical container for collections.
    Each database has its own file on the file system.
    A cluster can have multiple databases.

Collections:

    Grouping of documents.
    Equivalent to a table in relational databases.
    It does not impose a scheme.

Documents:

    A record within a collection.
    It is analogous to a JSON object (BSON).
    The basic unit within MongoDB.
    The driver takes care of the BSON transformations.
```

<hr>

#### Data types

<ul>
    <li> Strings: They are used to save texts. </li>
    <li> Boolean: True or false. </li>
    <li> ObjectId: They use the exact time in which we generate the query to always generate unique IDs. They exist in BSON but not in JSON. </li>
    <li> Date: They serve us to save dates and do range operations between them. </li>
    <li> Numbers: Doubles, Integers, Integers 64 bits and Decimals. </li>
    <li> Embedded documents: Documents within other documents ({}). </li>
    <li> Arrays: Arrays or lists of any other type of data, even other lists. </li>
</ul>

<hr>


#### Schemas and relationships

Schemas are the way we organize our documents in our collections. MongoDB does not impose any schema, but we can follow good practices and structure our documents in a similar way (not the same) to take advantage of the flexibility and scalability of the database without increasing the complexity of our applications.

Relationships are the way our entities or documents are linked to each other. For example: A career has multiple courses and each course has multiple classes.

#### Relationships between documents

Embedded documents help us to save information in a single document and save us the time it takes to consult different documents from references. However, references are still very important when we need to update information in different places continuously.

#### Operators

<ul>
    <li><b>$ eq: </b> = </li>
    <li><b>$ gt: </b> > </li>
    <li><b>$ gte: </b> >= </li>
    <li><b>$ lt: </b> < </li>
    <li><b>$ lte: </b> <= </li>
    <li><b>$ ne: </b> != </li>
    <li><b>$ in: </b> values ​​within a range. </li>
    <li><b>$ nin: </b> values ​​that are not within a range. </li>
    <li><b>$ and: </b> Match queries with a logical AND. </li>
    <li><b>$ not: </b> Inverts the effect of a query. </li>
    <li><b>$ nor: </b> Match queries with a logical NOR. </li>
    <li><b>$ or: </b> Match queries with a logical OR. </li>
    <li><b>$ exist: </b> Documents that have a specific field. </li>
    <li><b>$ type: </b> Documents that have a field of a specific type. </li>
    <li><b>$ all: </b> Arrays containing all elements of the query. </li>
    <li><b>$ elemMatch: </b> Documents that meet the $ elemMatch condition in one of its elements. </li>
    <li><b>$ size: </b> Documents that contain an array type field of a specific size. </li>
</ul>