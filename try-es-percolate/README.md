# try Elasticsearch percolate

```
GET _cat/indices?v
GET testindex?pretty
GET testindex/_mapping?pretty
GET /testindex/_search

DELETE /testindex

PUT /testindex
{
    "mappings": {
        "articles": {
            "properties": {
                "title": {
                    "type": "text"
                }
            }
        },
        "queries": {
            "properties": {
                "query": {
                    "type": "percolator"
                },
                "user_id": {
                    "type": "keyword"
                }
            }
        }
    }
}

PUT testindex/queries/1?refresh
{
    "query" : {
        "match" : {
            "title" : "golang"
        }
    },
    "user_id" : "po3rin"
}

PUT testindex/queries/2?refresh
{
    "query" : {
        "match" : {
            "title" : "elasticsearch"
        }
    },
    "user_id" : "po3rin"
}

PUT testindex/queries/3?refresh
{
    "query" : {
        "match" : {
            "title" : "python"
        }
    },
    "user_id" : "po4rin"
}

PUT /testindex/articles/1
{
    "title" : "develop app using golang and elasticsearch"
}

GET /testindex/_search
{
    "query" : {
        "percolate" : {
            "field": "query",
            "document_type" : "articles",
            "index" : "testindex",
            "type" : "articles",
            "id" : "1"
        }
    }
}

GET testindex/_search
{
  "query": {
    "bool": {
      "filter": [
        {
          "percolate" : {
            "field" : "query",
            "document_type" : "articles",
            "index" : "testindex",
            "type" : "articles",
            "id" : "1"
          }
        },
        { "term": { "user_id": "po3rin" } }
      ]
    }
  }
}
```
