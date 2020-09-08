package utils

import (
	"github.com/ontio/ontology/common"
)

type Token struct {
	TokenName       string
	TokenType       uint8
	ContractAddress common.Address
}

func NewToken(TokenName string,
	TokenType uint8,
	ContractAddress common.Address) Token {
	return Token{
		TokenName:       TokenName,
		TokenType:       TokenType,
		ContractAddress: ContractAddress,
	}
}

func (token *Token) Serialize(sink *common.ZeroCopySink) {
	sink.WriteString(token.TokenName)
	sink.WriteByte(token.TokenType)
	sink.WriteAddress(token.ContractAddress)
}
