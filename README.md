# RelaxaSearch

RelaxaSearch aims to be a full text search engine and lightweight replacemenet for ElasticSearch, built on top of Bleve using Go.

RelaxaSearch derives from portuguese word "Relaxa", which basically means be calm. Only those who already tried to setup ElasticSearch, or OpenSearch knows how stressfull setting them up it can be sometimes.

On other hands, Relax Search tries to be simplest and easiest solution for a full text search engine.

## Installation



## Usage
You can just run the binary on any directory you want. Please note that relaxasearch will look up for a directory called `relaxasearchData`. This diretory **should be on the same directory that you are running relaxasearch**. It will be responsable for storing the index data. If you hadn't set it up, relaxasearch will throw an error on terminal. 

## How to build
```shell
go build main.go

./main
```

## Routes

**Index routes**
- **GET** `/relaxasearch/v1` - Returns a json containing all the indexed created 
- **GET** `/relaxasearch/v1/:indexName/count` - Returns a json containing the number of docs of an index
- **GET** `/relaxasearch/v1/:indexName` - Returns informations about a specific index
- **POST** - `/relaxasearch/v1/:indexName` - Creates a index.
- **DELETE** - `/relaxasearch/v1/:indexName` - Deletes a index.

**Document routes**
- **GET** `/relaxasearch/v1/:indexName/:docId` - Returns a json containing all the indexed created 

- **POST** - `/relaxasearch/v1/:indexName/:docId` - Creates or updates a document with given id. The request should contains a json body with the document fields.
- **POST** - `/relaxasearch/v1/:indexName/search` - Search data in a index. Please see the next session to see how to build search queries. 
- **DELETE** - `/relaxasearch/v1/:indexName/:docId` - Deletes a a document from an index.

## Searching
To perform a search, you have to do a POST request to `/relaxasearch/v1/:indexName/search` with the following request body:

```json
{
	"size": 10, // the size of results
	"from": 0, // for pagination
	"explain": false, // explains the query result
	"highlight": {}, // fields to be highlighted
	"query": {
		"boost": 1,
		"query": "your query"
	},
	"fields": [
		"*" // you can determine which fields you want to search against.
	]
}
```
## Adding a document
To add a document, just do a **POST** request to `/relaxasearch/v1/:indexName/:docId` with a json body contaning your document fields:
```json
{
	"title": "My document",
	"content": "My document content here",
    "keywords":"document, new, keyword"
}
```

Please note that you can create or update a document using the same route, so be careful to do not remove an existing document accidentally when you are creating a new one.

## Roadmap
There is a lot of things to be finished. Some of them are:

- [ ] UI for basic search test 
- [ ] Protect routes with a token based auth
- [ ] Improve document indexing with filters and normalizers
- [ ] Improve memory usage on indexing operations
- [ ] ...

  
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/)