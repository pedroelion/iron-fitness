package modelo

type Plano struct {
	Nome      string
	Preco     string
	Vantagens []string
	Destaque  bool // Novo campo para destacar o melhor plano
}

type Aula struct {
	Nome      string
	Descricao string
	Imagem    string
}

type Estrutura struct {
	Titulo string
	Imagem string
}

type PageData struct {
	Planos     []Plano
	Aulas      []Aula
	Estruturas []Estrutura
}

func GetPageData() PageData {
	return PageData{
		Planos: []Plano{
			{
				Nome:     "PLANO IRON",
				Preco:    "R$ 139,90",
				Destaque: false,
				Vantagens: []string{
					"Acesso a todas as unidades",
					"12 Meses de fidelidade",
					"Sem manutenção anual",
					"Sem taxa de adesão",
				},
			},
			{
				Nome:     "PLANO FLEX",
				Preco:    "R$ 149,90",
				Destaque: true, // Este plano terá um visual de destaque
				Vantagens: []string{
					"Acesso a todas as unidades",
					"Sem fidelidade",
					"Cancelamento a qualquer momento",
					"Acesso total às aulas coletivas",
				},
			},
		},
		Aulas: []Aula{
			{Nome: "MUAY THAI", Descricao: "Força, disciplina e condicionamento extremo.", Imagem: "https://images.unsplash.com/photo-1549719386-74dfcbf7dbed?q=80&w=800"},
			{Nome: "CYCLE", Descricao: "Queima calórica intensa com muita energia.", Imagem: "https://images.unsplash.com/photo-1534438327276-14e5300c3a48?q=80&w=800"},
			{Nome: "CROSS", Descricao: "Supere seus limites com treinos dinâmicos.", Imagem: "https://images.unsplash.com/photo-1534258936925-c58bed479fcb?q=80&w=800"},
		},
		Estruturas: []Estrutura{
			{Titulo: "MAQUINÁRIO IMPORTADO", Imagem: "https://images.unsplash.com/photo-1581009146145-b5ef050c2e1e?q=80&w=800"},
			{Titulo: "ÁREA DE PESO LIVRE GIGANTE", Imagem: "https://images.unsplash.com/photo-1540497077202-7c8a3999166f?q=80&w=800"},
			{Titulo: "ESTÚDIO COLETIVO", Imagem: "https://images.unsplash.com/photo-1518611012118-696072aa579a?q=80&w=800"},
		},
	}
}
