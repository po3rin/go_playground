people : [...]

aaa: string

index:{
	mappings: {
		for p in (people) {
			"\(p.name)": {
				age: p.age
			}
		}
	}
}
