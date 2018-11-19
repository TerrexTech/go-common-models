This repository provides the models commonly used across TerrexTech projects.  
It also includes [`bootstrap`][0] package for conveniently creating `Event` and `EventMeta` tables and associated Keyspace in Cassandra.

#### Included models:

* **Command**
* **Document**
* **Error**
* **Event**
* **EventMeta**
* **EventStoreQuery**
* **LogEntry**

The models provided are intended to be imported by any libraries dealing with the respective models, to ensure consistency across structures/schema.

**Go Docs:**
 * **[bootstrap][0]**
 * **[definition][1]**
 * **[model][2]**

### Usage:

As mentioned above, this library can be used to directly create the tables `events`, `events_meta`, and the associated keyspace.

To create the tables, simply call the [`bootstrap#Event`][3] and the [`bootstrap#EventMeta`][4] methods.

The required information (such as Cassandra Hosts, Keyspace/Table names, etc.) is read from the Environemnt.

The default configuration can be found in the `.env` file at the root.

  [0]: https://godoc.org/github.com/TerrexTech/go-common-models/bootstrap
  [1]: https://godoc.org/github.com/TerrexTech/go-common-models/definition
  [2]: https://godoc.org/github.com/TerrexTech/go-common-models/model
  [3]: https://godoc.org/github.com/TerrexTech/go-common-models/bootstrap#Event
  [4]: https://godoc.org/github.com/TerrexTech/go-common-models/bootstrap#EventMeta

 