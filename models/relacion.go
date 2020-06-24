package models

//Relacion ... define una relación entre un usuario y otro
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuariorelacionId"`
}
