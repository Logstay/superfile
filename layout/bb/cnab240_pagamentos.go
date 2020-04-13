package bb

const CNAB240Pagamentos = `
# FORMATO: BANCO DO BRASIL - FEBRABAN CNAB240
# OBJETIVO DO ARQUIVO: COBRANÇA BANCÁRIA
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 240 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita.
#
servico: "pagamentos"
versao: "083"
layout: "cnab240"

remessa:
  header_arquivo:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: "0000"
    tipo_servico_01:
      pos: [8, 8]
      picture: "9(1)"
      default: "0"
    exclusivo_febraban_01:
      pos: [9, 17]
      picture: "X(9)"
      default: ""
    tipo_inscricao_empresa:
      pos: [18, 18]
      picture: "9(1)"
    numero_inscricao_empresa:
      pos: [19, 32]
      picture: "9(14)"
    numero_convenio:
      pos: [33, 41]
      picture: "9(9)"
    codigo_convenio:
      pos: [42, 45]
      picture: "9(4)"
    reservado_banco_01:
      pos: [46, 50]
      picture: "X(5)"
      default: ""
    arquivo_teste:
      pos: [51, 52]
      picture: "X(2)"
      default: ""
    agencia_mantenedora_conta:
      pos: [53, 57]
      picture: "9(5)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    numero_conta_corrente:
      pos: [59, 70]
      picture: "9(12)"
    digito_verificador_conta:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    nome_banco:
      pos: [103, 132]
      picture: "X(30)"
      default: "BANCO DO BRASIL S.A."
    exclusivo_febraban_02:
      pos: [133, 142]
      picture: "X(10)"
      default: ""
    codigo_remessa_retorno:
      pos: [143, 143]
      picture: "9(1)"
    data_geracao_arquivo:
      pos: [144, 151]
      picture: "9(8)"
    hora_geracao_arquivo:
      pos: [152, 157]
      picture: "9(6)"
    numero_sequencial_arquivo:
      pos: [158, 163]
      picture: "9(6)"
    versao_layout_arquivo:
      pos: [164, 166]
      picture: "9(3)"
      default: "084"
    densidade_gravacao_arquivo:
      pos: [167, 171]
      picture: "9(5)"
      default: 0
    reservado_banco_02:
      pos: [172, 191]
      picture: "X(20)"
      default: ""
    reservado_empresa_01:
      pos: [192, 211]
      picture: "X(20)"
      default: ""
    exclusivo_febraban_03:
      pos: [212, 222]
      picture: "X(11)"
      default: ""
    identificacao_cobranca_sem_papel:
      pos: [223, 225]
      picture: "X(3)"
      default: ""
    exclusivo_vans:
      pos: [226, 228]
      picture: "9(3)"
      default: 000
    tipo_servico_02:
      pos: [229, 230]
      picture: "9(2)"
      default: 00
    codigos_ocorrencias:
      pos: [231, 240]
      picture: "9(10)"
      default: "0000000000"

  trailer_arquivo:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 1
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: "9999"
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 9
    exclusivo_febraban_01:
      pos: [9, 17]
      picture: "X(9)"
      default: ""
    quantidade_lotes_arquivo:
      pos: [18, 23]
      picture: "9(6)"
    quantidade_registros_arquivo:
      pos: [24, 29]
      picture: "9(6)"
    quantidade_contas_conciliacao_lotes:
      pos: [30, 35]
      picture: "9(6)"
      default: 0
    exclusivo_febraban_02:
      pos: [36, 240]
      picture: "X(205)"
      default: ""

  header_lote:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 1
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 1
    tipo_operacao:
      pos: [9, 9]
      picture: "X(1)"
      default: "C"
    tipo_servico:
      pos: [10, 11]
      picture: "9(2)"
    forma_lancamento:
      pos: [12, 13]
      picture: "9(2)"
    versao_layout_lote:
      pos: [14, 16]
      picture: "9(3)"
      default: "083"
    exclusivo_febraban_01:
      pos: [17, 17]
      picture: "X(1)"
      default: ""
    tipo_inscricao_empresa:
      pos: [18, 18]
      picture: "9(1)"
    numero_inscricao_empresa:
      pos: [19, 32]
      picture: "9(14)"
    numero_convenio:
      pos: [33, 41]
      picture: "9(9)"
    codigo_convenio:
      pos: [42, 45]
      picture: "9(4)"
    reservado_banco_01:
      pos: [46, 50]
      picture: "X(5)"
      default: ""
    arquivo_teste:
      pos: [51, 52]
      picture: "X(2)"
      default: ""
    agencia_mantenedora_conta:
      pos: [53, 57]
      picture: "9(5)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    numero_conta_corrente:
      pos: [59, 70]
      picture: "9(12)"
    digito_verificador_conta:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    mensagem:
      pos: [103, 142]
      picture: "X(40)"
      default: "0"
    logradouro:
      pos: [143, 172]
      picture: "X(30)"
    numero:
      pos: [173, 177]
      picture: "9(5)"
    complemento:
      pos: [178, 192]
      picture: "X(15)"
    cidade:
      pos: [193, 212]
      picture: "X(20)"
    cep:
      pos: [213, 217]
      picture: "9(5)"
    complemento_cep:
      pos: [218, 220]
      picture: "X(3)"
    estado:
      pos: [221, 222]
      picture: "X(2)"
    exclusivo_febraban_02:
      pos: [223, 230]
      picture: "X(8)"
      default: ""
    codigos_ocorrencias_retorno:
      pos: [231, 240]
      picture: "X(10)"
      default: "0000000000"

  trailer_lote:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: "001"
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: "5"
    exclusivo_febraban_01:
      pos: [9, 17]
      picture: "X(9)"
      default: ""
    quantidade_registros_lote:
      pos: [18, 23]
      picture: "9(6)"
    somatoria_valores:
      pos: [24, 41]
      picture: "9(16)V9(2)"
    somatoria_quantidade_moedas:
      pos: [42, 59]
      picture: "9(13)V9(5)"
      default: "000000000000000000"
    numero_aviso_debito:
      pos: [60, 65]
      picture: "9(6)"
      default: "000000"
    exclusivo_febraban_02:
      pos: [66, 230]
      picture: "X(165)"
      default: ""
    codigos_ocorrencias_retorno:
      pos: [231, 240]
      picture: "X(10)"

  detalhes:
    segmento_a:
      codigo_banco:
        pos: [1, 3]
        picture: "9(3)"
      lote_servico:
        pos: [4, 7]
        picture: "9(4)"
      tipo_registro:
        pos: [8, 8]
        picture: "9(1)"
        default: "3"
      numero_sequencial_registro_lote:
        pos: [9, 13]
        picture: "9(5)"
      codigo_segmento_registro_detalhe:
        pos: [14, 14]
        picture: "X(1)"
        default: "A"
      tipo_movimento:
        pos: [15, 15]
        picture: "9(1)"
      codigo_instrucao_movimento:
        pos: [16, 17]
        picture: "9(2)"
      codigo_camara_centralizadora:
        pos: [18, 20]
        picture: "9(3)"
      codigo_banco_favorecido:
        pos: [21, 23]
        picture: "9(3)"
      agencia_mantenedora_conta_favorecido:
        pos: [24, 28]
        picture: "9(5)"
      digito_verificador_agencia:
        pos: [29, 29]
        picture: "X(1)"
      numero_conta_corrente:
        pos: [30, 41]
        picture: "9(12)"
      digito_verificador_conta:
        pos: [42, 42]
        picture: "X(1)"
      digito_verificador_agencia_conta:
        pos: [43, 43]
        picture: "X(1)"
      nome_favorecido:
        pos: [44, 73]
        picture: "X(30)"
      numero_documento_atribuido_empresa:
        pos: [74, 93]
        picture: "X(20)"
      data_pagamento:
        pos: [94, 101]
        picture: "9(8)"
      tipo_moeda:
        pos: [102, 104]
        picture: "X(3)"
      quantidade_moeda:
        pos: [105, 119]
        picture: "9(10)V9(5)"
      valor_pagamento:
        pos: [120, 134]
        picture: "9(13)V9(2)"
      numero_documento_atribuido_banco:
        pos: [135, 154]
        picture: "X(20)"
      data_real_efetivacao_pagamento:
        pos: [155, 162]
        picture: "9(8)"
      valor_real_efetivacao_pagamento:
        pos: [163, 177]
        picture: "9(13)V9(2)"
      outras_informacoes:
        pos: [178, 217]
        picture: "X(40)"
      complemento_tipo_servico:
        pos: [218, 219]
        picture: "X(2)"
      codigo_finalidade_ted:
        pos: [220, 224]
        picture: "X(5)"
      complemento_finalidade_pagamento:
        pos: [225, 226]
        picture: "X(2)"
      exclusivo_febraban_01:
        pos: [227, 229]
        picture: "X(3)"
        default: ""
      aviso_favorecido:
        pos: [230, 230]
        picture: "9(1)"
      codigos_ocorrencias_retorno:
        pos: [231, 240]
        picture: "X(10)"

    segmento_b:
      codigo_banco:
        pos: [1, 3]
        picture: "9(3)"
        default: "001"
      lote_servico:
        pos: [4, 7]
        picture: "9(4)"
      tipo_registro:
        pos: [8, 8]
        picture: "9(1)"
        default: "3"
      numero_sequencial_registro_lote:
        pos: [9, 13]
        picture: "9(5)"
      codigo_segmento_registro_detalhe:
        pos: [14, 14]
        picture: "X(1)"
        default: "B"
      exclusivo_febraban_01:
        pos: [15, 17]
        picture: "X(3)"
        default: ""
      tipo_inscricao_favorecido:
        pos: [18, 18]
        picture: "9(1)"
      numero_inscricao_favorecido:
        pos: [19, 32]
        picture: "9(14)"
      logradouro:
        pos: [33, 62]
        picture: "X(30)"
        default: ""
      numero:
        pos: [63, 67]
        picture: "9(5)"
        default: "0"
      complemento:
        pos: [68, 82]
        picture: "X(15)"
        default: ""
      bairro:
        pos: [83, 97]
        picture: "X(15)"
        default: ""
      cidade:
        pos: [98, 117]
        picture: "X(20)"
        default: ""
      cep:
        pos: [118, 122]
        picture: "9(5)"
        default: "0"
      complemento_cep:
        pos: [123, 125]
        picture: "X(3)"
        default: ""
      estado:
        pos: [126, 127]
        picture: "X(2)"
        default: ""
      data_vencimento:
        pos: [128, 135]
        picture: "9(8)"
        default: "0"
      valor_documento:
        pos: [136, 150]
        picture: "9(13)V9(2)"
        default: "0"
      valor_abatimento:
        pos: [151, 165]
        picture: "9(13)V9(2)"
        default: "0"
      valor_desconto:
        pos: [166, 180]
        picture: "9(13)V9(2)"
        default: "0"
      valor_mora:
        pos: [181, 195]
        picture: "9(13)V9(2)"
        default: "0"
      valor_multa:
        pos: [196, 210]
        picture: "9(13)V9(2)"
        default: "0"
      codigo_doc_favorecido:
        pos: [211, 225]
        picture: "X(15)"
        default: ""
      aviso_favorecido:
        pos: [226, 226]
        picture: "9(1)"
        default: "0"
      uso_exclusivo_siape:
        pos: [227, 232]
        picture: "9(6)"
        default: "0"
      codigo_ispb:
        pos: [233, 240]
        picture: "9(8)"
        default: "0"

    segmento_c:
      codigo_banco:
        pos: [1, 3]
        picture: "9(3)"
        default: "001"
      lote_servico:
        pos: [4, 7]
        picture: "9(4)"
      tipo_registro:
        pos: [8, 8]
        picture: "9(1)"
        default: "3"
      numero_sequencial_registro_lote:
        pos: [9, 13]
        picture: "9(5)"
      codigo_segmento_registro_detalhe:
        pos: [14, 14]
        picture: "X(1)"
        default: "C"
      exclusivo_febraban_01:
        pos: [15, 17]
        picture: "X(3)"
        default: ""
      valor_ir:
        pos: [18, 32]
        picture: "9(15)"
        default: "0"
      valor_iss:
        pos: [33, 47]
        picture: "9(13)V9(2)"
        default: "0"
      valor_iof:
        pos: [48, 62]
        picture: "9(13)V9(2)"
        default: "0"
      valor_outras_deducoes:
        pos: [63, 77]
        picture: "9(13)V9(2)"
        default: "0"
      valor_outros_acrescimos:
        pos: [78, 92]
        picture: "9(13)V9(2)"
        default: "0"
      agencia_favorecido:
        pos: [93, 97]
        picture: "9(5)"
        default: "0"
      digito_verificador_agencia:
        pos: [98, 98]
        picture: "X(1)"
        default: "0"
      numero_conta_corrente:
        pos: [99, 110]
        picture: "9(12)"
        default: "0"
      digito_verificador_conta:
        pos: [111, 111]
        picture: "X(1)"
        default: "0"
      digito_verificador_agencia_conta:
        pos: [112, 112]
        picture: "X(1)"
        default: "0"
      valor_inss:
        pos: [113, 127]
        picture: "9(13)V9(2)"
        default: "0"
      numero_conta_pagamento_creditada:
        pos: [128, 147]
        picture: "9(20)"
        default: "0"
      exclusivo_febraban_02:
        pos: [148, 240]
        picture: "X(93)"
        default: ""
`
