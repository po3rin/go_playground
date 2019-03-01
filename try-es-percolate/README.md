# try Elasticsearch percolate

```go
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
                }
            }
        }
    }
}

GET testindex?pretty
GET testindex/_mapping?pretty

PUT testindex/queries/1?refresh
{
    "query" : {
        "match" : {
            "title" : "bonsai tree"
        }
    }
}

PUT testindex/queries/1?refresh
{
    "query" : {
        "match" : {
            "title" : "golang"
        }
    }
}

PUT testindex/queries/2?refresh
{
    "query" : {
        "match" : {
            "title" : "elasticsearch"
        }
    }
}

GET /testindex/_search
{
    "query" : {
        "percolate" : {
            "field" : "query",
            "document_type" : "articles",
            "document" : {
                "title" : "A new bonsai tree in the office"
            }
        }
    }
}

PUT /testindex/title/1
{
    "title" : "develop app using golang and elasticsearch"
}

GET /testindex/_search

GET /testindex/_search
{
    "query" : {
        "percolate" : {
            "field": "query",
            "document_type" : "articles",
            "index" : "testindex",
            "type" : "title",
            "id" : "1"
        }
    }
}

PUT /testindex/queries/1?refresh
{
    "query" : {
        "match" : {
            "title" : "brown fox"
        }
    }
}

PUT /testindex/queries/2?refresh
{
    "query" : {
        "match" : {
            "title" : "lazy dog"
        }
    }
}

GET testindex/_search
{
    "query" : {
        "percolate" : {
            "field": "query",
            "document_type" : "articles",
            "document" : {
                "title" : "The quick brown fox jumps over the lazy dog"
            }
        }
    },
    "highlight": {
      "fields": {
        "title": {}
      }
    }
}

PUT /testindex/title/1
{
  "title" : "The quick brown fox jumps over the lazy dog"
}

GET /testindex/_search
{
    "query" : {
        "percolate" : {
            "field": "query",
            "document_type" : "articles",
            "index" : "testindex",
            "type" : "title",
            "id" : "1"
        }
    }
}

DELETE /testindex
```