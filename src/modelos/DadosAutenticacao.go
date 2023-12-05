package modelos

// DadosAutenticacao contém os dados de autenticação token e id do usuário autenticado
type DadosAutenticacao struct{
	ID string `json:"id"`
	Token string `json:"token"`
}