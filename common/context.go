package common

type Field struct {
	Name      string
	Type      string
	Nullable  string
	TableName string
	Comment   string
	Tag       string
}

type ParserInfo struct {
	ProjectPrefix string
	DSN 		  string
	// k: struct name, v: struct fields
	Structs map[string][]*Field
	
	ModelDependence *ModelDependence
	DaoDependence *DaoDependence
	ServiceDependence *ServiceDependence
}

var Parser = &ParserInfo{
	Structs: make(map[string][]*Field),
}