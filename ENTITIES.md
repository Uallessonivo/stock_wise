1. **Entidades Principais**:

   a. **Usuário**:
      - Atributos:
        - ID
        - Nome
        - E-mail
        - Senha (ou método de autenticação)
        - Cargo/Função (por exemplo: administrador, operador, etc.)

   b. **Produto**:
      - Atributos:
        - ID
        - Nome
        - Categoria
        - Data de Validade
        - Lote
        - Fornecedor
        - Preço
        - Quantidade em Estoque
        - Quantidade Mínima Aceitável
        - Quantidade Máxima Recomendada
      - Relacionamentos:
        - Um Produto pode ter múltiplos Movimentos de Estoque
        - Um Produto pertence a um ou mais Fornecedores

   c. **Fornecedor**:
      - Atributos:
        - ID
        - Nome
        - Contato (e-mail, telefone, etc.)
      - Relacionamentos:
        - Um Fornecedor pode fornecer múltiplos Produtos

   d. **Movimento de Estoque**:
      - Atributos:
        - ID
        - Data e Hora
        - Tipo (Entrada ou Saída)
        - Quantidade
        - Motivo (Venda, Devolução, Transferência, etc.)
      - Relacionamentos:
        - Um Movimento de Estoque está associado a um Produto
        - Um Movimento de Estoque é realizado por um Usuário

Agora, com a adição da entidade de Usuário, é possível rastrear qual usuário realizou cada movimentação de estoque. Isso proporcionará um controle mais granular sobre as ações realizadas no sistema, permitindo uma maior segurança e rastreabilidade das operações.