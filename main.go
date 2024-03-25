package main

import (
	"fmt"

	"example.com/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type kotlinListener struct {
	*parser.BaseKotlinParserListener
}

func (l *kotlinListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	functionName := ctx.SimpleIdentifier().GetText()
	fmt.Println("Found function declaration:", functionName)
}

func (l *kotlinListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	className := ctx.SimpleIdentifier().GetText()
	fmt.Println("Found class declaration:", className)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("fun main() {}")

	// Create the Lexer
	lexer := parser.NewKotlinLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewKotlinParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&kotlinListener{}, p.KotlinFile())
}
