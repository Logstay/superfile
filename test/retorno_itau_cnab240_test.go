package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRetornoItauCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB240Cobranca)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/cobranca_itau_cnab240.ret")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, 52, len(retorno.Header))
	assert.Equal(t, 16, len(retorno.Trailer))
	assert.Equal(t, 1, len(retorno.Lotes))
	assert.Equal(t, 48, len(retorno.Lotes[0].Header))
	assert.Equal(t, 4, len(retorno.Lotes[0].Segmentos()))
	assert.Equal(t, 24, len(retorno.Lotes[0].Trailer))
	assert.Equal(t, int64(1), retorno.TotalLotes())
	assert.Equal(t, 0, len(retorno.Falhas()))
}
