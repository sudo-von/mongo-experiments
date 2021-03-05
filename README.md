# Bases de datos NoSQL

Las bases de datos NoSQL tienen 4 grandes familias: 
<ul>
<li>Key Value Stores</li>
<li>Basadas en grafos</li>
<li>Columnares</li>
<li>Basadas en documentos</li>
</ul>

<b>Key Value Stores</b>: Guardan la información en formato de llaves y valores. Las usamos para guardar cache, información de sesión de los usuarios o cosas muy sencillas. Son muy rápidas de consultar pero no podemos usarlas en casos más complejos donde necesitamos estructuras más especiales. El mejor ejemplo de estas bases de datos es Redis.

<b>Graph Databases</b>: Bases de datos basadas en Grafos. Nos permiten establecer conexiones entre nuestras entidades para realizar consultas de una forma más eficiente que en bases de datos relacionales (así como Twitter o Medium donde cada publicación tiene diferentes relaciones entre sus usuarios, likes, etc). Por ejemplo: Neo4j o JanusGraph.

<b>Wide-column Stores</b>: Bases de datos columnares. Tienen una llave de fila y otra de columnas para hacer consultas muy rápidas y guardar grandes cantidades de información pero modelar los datos se puede volver un poco complicado. Las usamos en Big Data, IoT, sistemas de recomendaciones, entre otras. Por ejemplo: Cassandra o HBase.

<b>Document Databases</b>: Bases de datos basadas en documentos. Nos permiten guardar documentos dentro de colecciones, tiene muy buena performance y flexibilidad que nos permite modelar casos de la vida real de forma sencilla y efectiva. Por ejemplo: MongoDB o CouchBase.

<hr>

# MongoDB

MongoDB es una base de datos gratis y de código abierto No Relacional basada en documentos que nos permite guardar una gran cantidad de documentos de forma distribuida. Mongo también es el nombre de la compañía que desarrolla el código de esta base de datos.

Una de sus principales características es que nos permite guardar nuestras estructuras o documentos en formato JSON (no exactamente JSON, pero si algo muy parecido, lo veremos más adelante) para tener una gran flexibilidad a la hora de modelar situaciones de la vida real.

Por ser una base de datos distribuida podemos hablar no de uno sino de varios servidores, lo que conocemos como el Cluster de MongoDB. Gracias a esto obtenemos una gran escalabilidad de forma horizontal (escalabilidad en cantidad de servidores).

MongoDB es “Schema Less” lo que permite que nuestros documentos tengan estructuras diferentes sin afectar su funcionamiento, algo que no podemos hacer con las tablas de las bases de datos relacionales. Su lenguaje para realizar queries, índices y agregaciones es muy expresivo.

<hr>

# MongoDB Atlas

Tenemos varios proveedores que nos permiten utilizar o alquilar MongoDB como servicio y en este caso vamos a usar MongoDB Atlas por ser desarrollado por las mismas personas que desarrollan MongoDB.

MongoDB Atlas tiene las siguientes características:

<ul>
<li>Aprovisionamiento automático de clusters con MongoDB</li>
<li>Alta disponibilidad</li>
<li>Altamente escalable</li>
<li>Seguro</li>
<li>Disponible en AWS, GCP y Microsoft Azure</li>
<li>Fácil monitoreo y optimización</li>
</ul>

<hr>

#### Comandos básicos de MongoDB

Correr MongoDB.
<code>mongo</code>
Crear y usar una DB.
<code>use DB</code>
Listar DBS.
<code>show dbs</code>
Crear colección e insertar documento.
<code>db.COLLECTION.insertOne({"name":"VoN"})</code>
Listar documentos de una colección.
<code>db.COLLECTION.find()</code>
Listar las colecciones.
<code>show collections</code>
Mostrar un documento cualquiera.
<code>db.COLLECTION.findOne()</code>
Saber que DB estamos usando.
<code>db</code>

<hr>

<b>Ejemplo de un insert:</b>
```
db.fake_users.insertOne(
    {
        "_id": "1",
        "username": "fake",
        "name": "data",
        "active": false,
        "skills": ["Golang"],
        "cash" : 0
    }
)
```
<hr>

#### ¿Qué son los drivers en MongoDB?

Son las librerías que utilizamos para comunicar nuestra aplicación con nuestra base de datos.

Sin nuestros drivers no podemos trabajar con nuestros cluster de base de datos.

#### ¿Cómo agregar los drivers dentro de nuestro proyecto?

Usamos un gestor de dependencias. Lo agregamos en nuestro gestor de dependencias.

```
Python: python -m pip install pyongo
Node.js: npm install mongodb --save
GO: dep ensure -add go.mongodb.org/mongo-driver/mongo
```

#### Inicio rápido transversal a la mayoría de lenguajes

<ol>
<li>Crear conexión MongoCliente.</li>
<li>Obtener la base de datos MongoDatabase.</li>
<li>Acceder a una colección MongoCollection.</li>
<li>CRUD.</li>
</ol>

#### Bases de datos, Colecciones y Documentos en MongoDB

```
Base de datos:

    Contenedor físico de colecciones.
    Cada base de datos tiene su archivo propio en el sistema de archivos.
    Un cluster puede tener múltiples bases de datos.

Colecciones:

    Agrupación de documentos.
    Equivalente a una tabla en las bases de datos relacionales.
    No impone un esquema.

Documentos:

    Un registro dentro de una colección.
    Es análogo a un objeto JSON (BSON).
    La unidad básica dentro de MongoDB.
    El driver se encarga de las transformaciones de los BSON.
```

#### Operaciones CRUD desde la consola de MongoDB

<b>InsertOne()</b> 
Guarda un documento dentro de la colección en la BD. Si no existe la colección MONGODB la crea. Al hacer un insert dentro de una colección sino especifica el campo _id (es un campo obligatorio que la BD pide) Mongo lo crea automáticamente como un objeto, en el caso de especificarlo puede ser tomado como un string si así lo deseas.

```
db.fake_users.insertOne(
    {
        "_id": "1",
        "username": "fake",
        "name": "data",
        "active": false,
        "skills": ["Golang"],
        "cash" : 0
    }
)
```

<b>InsertMany()</b>
Recibe un arreglo en formato JSON y los ingresa a MongoDB automáticamente MongoDB. Tiene atomicidad dentro de los documentos, esto quiere decir que en el las operaciones de escritura garantiza que al escribir un documento la operación es atómica (se escribe y si no se escribe se hace rollback).

```
db.fake_users.insertMany(
    [
    {
        "username": "fake",
        "name": "data",
        "active": false,
        "skills": ["Golang"],
        "cash" : 0
    },
        {
        "username": "fake2",
        "name": "data2",
        "active": true,
        "skills": ["Python"],
        "cash" : 1
    }
    ]
)
```

<b>Find</b> ({condicion}).
Se hace una consulta con un filtro y nos permite buscar uno o varios documentos que cumplan la condición.

```
db.fake_users.find(
    {
        "username": "fake"
    }
)
```

<b>FindOne()</b>
Devuelve los documentos en orden natural y retorna la primera coincidencia (orden de índices).

```
db.fake_users.findOne(
    {
        "username" : "fake"
    }
)
```

```
db.fake_users.findOne(
    {
        "username" : "fake",
        "cash" : { $lte : 5} 
    }
)
```

<b>UpdateOne</b> (Filtro, nuevo valor). 

```
db.fake_users.updateOne(
    {
        "username": "fake"
    },
    {
        "$set": { "username": "not fake" }
    }
)
```

<b>UpdateMany</b> (Filtro, nuevo valor). 


```
db.fake_users.updateMany(
    {
        "username": "not fake"
    },
    {
        "$set": { "username": "fake" }
    }
)
```

<b>DeleteOne()</b>
Hace el paso del filtro, borrará según su orden natural que encuentre con el filtro.

```
db.fake_users.deleteOne(
    {
        "username": "fake"
    }
)
```

<b>DeleteMany()</b>
Compara el filtro con la cantidad de doc. Que cuplen con el filtro y los borra

```
db.fake_users.deleteMany(
    {
        "username": "fake"
    }
)
```

<hr>

#### Tipos de datos

<ul>
    <li>Strings: Nos sirven para guardar textos.</li>
    <li>Boolean: Información cierta o falsa (true y false).</li>
    <li>ObjectId: Utilizan el tiempo exacto en el que generamos la consulta para siempre generan IDs únicos. Existen en BSON pero no en JSON.</li>
    <li>Date: Nos sirven para guardar fechas y hacer operaciones de rangos entre ellas.</li>
    <li>Números: Doubles, Integers, Integers 64 bits y Decimals.</li>
    <li>Documentos Embebidos: Documentos dentro de otros documentos ({}).</li>
    <li>Arrays: Arreglos o listas de cualquier otro tipo de datos, incluso, de otras listas.</li>
</ul>

<hr>

#### Esquemas y relaciones

Los esquemas son la forma en que organizamos nuestros documentos en nuestras colecciones. MongoDB no impone ningún esquema pero podemos seguir buenas prácticas y estructurar nuestros documentos de forma parecida (no igual) para aprovechar la flexibilidad y escalabilidad de la base de datos sin aumentar la complejidad de nuestras aplicaciones.

Las relaciones son la forma en que nuestras entidades o documentos sen encuentran enlazados unos con otros. Por ejemplo: Una carrera tiene multiples cursos y cada curso tiene multiples clases.

#### Relaciones entre documentos

Las documentos embebidos nos ayudan a guardar la información en un solo documento y nos ahorra el tiempo que tardamos en consultar diferentes documentos a partir de referencias. Sin embargo, las referencias siguen siendo muy importantes cuando debemos actualizar información en diferentes lugares de forma continua.