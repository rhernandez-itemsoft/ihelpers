package datatable

//https://datatables.net/manual/server-side

//DataTable es la definición de la estructura de parametros que se envían y reciben en el server-side del Datatables.net
type DataTable struct {
	Draw    int         `json:"draw"`
	Start   int         `json:"start"`
	Length  int         `json:"length"`
	Search  DTSearch    `json:"search"`
	Order   []DTOrder   `json:"search"`
	Columns []DTColumns `json:"columns"`
}

//DTSearch la busqueda
type DTSearch struct {
	Value string `json:"value"`
	Regex bool   `json:"regex"`
}

//DTOrder El ordenamiento
type DTOrder struct {
	Column int    `json:"column"`
	Dir    string `json:"dir"`
}

//DTColumns Las columnas de la tabla
type DTColumns struct {
	Data       string   `json:"data"`
	Name       string   `json:"name"`
	Searchable bool     `json:"searchable"`
	Orderable  bool     `json:"orderable"`
	Search     DTSearch `json:"search"`
}

func OrderString(_dt DataTable) string {
	_order := ""
	for _, _ord := range _dt.Order {
		if _dt.Columns[_ord.Column].Orderable == true {
			if len(_order) > 0 {
				_order += ", "
			}
			_order += _dt.Columns[_ord.Column].Name + " " + _ord.Dir
		}
	}
	return _order
}
func SearchString(_dt DataTable) string {

	if _dt.Search.Regex == true {
		return ""
	}

	return _dt.Search.Value
}

func FilterString(_dt DataTable) (string, []string) {
	_search := ""
	var _values []string
	for _, _col := range _dt.Columns {
		if _col.Searchable == true {
			if len(_search) > 0 {
				_search += " and "
			}
			_search += _col.Name + " = ? "
			_values = append(_values, _col.Search.Value)
		}

	}

	return _search, _values
}
