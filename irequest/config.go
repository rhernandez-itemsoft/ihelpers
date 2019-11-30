package irequest

//# parametros del request:    https://datatables.net/manual/server-side
//#documentacion del tatable:  https://datatables.net/examples/styling/bootstrap4
//#datatable version angular:  https://l-lin.github.io/angular-datatables/#/getting-started

//RequestDT request del tata table
type RequestDT struct {
	Draw    uint32
	Start   uint32
	Length  uint32
	Search  searchDT
	Order   []orderDT
	Columns []columnsDT
}

type searchDT struct {
	Value string
	Regex bool
}
type orderDT struct {
	Dir    string
	Column string
}

type columnsDT struct {
	Data       string
	Name       string
	Searchable bool
	Orderable  bool
	Search     searchDT
}
