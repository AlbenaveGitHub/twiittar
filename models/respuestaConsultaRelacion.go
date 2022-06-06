package models

/*RespuestaConsultaRelacion tiene el true o el false que se obtiene de consultar la realacion entre 2 ususario*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
