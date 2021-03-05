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
db.inventory.insertOne(
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