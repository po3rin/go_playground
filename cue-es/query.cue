queries: [...]
useDataSort: bool
size: int

{
	"query": {
		"bool": {
			"should": [
				for q in (queries) {
					{
						"fields": [
							"title",
						]
						"query": (q)
					}
				}]
		}
	}
	if useDataSort {
		"sort": [
			{
				"post_date": {
					"format": "strict_date_optional_time_nanos"
				}
			},
		]
	}
	"size": (size)
}
