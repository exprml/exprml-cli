package main

import (
	"fmt"
	"github.com/exprml/exprml-go"
	exprmlpb "github.com/exprml/exprml-go/pb/exprml/v1"
	"io"
	"log"
	"os"
)

//go:generate go run github.com/Jumpaku/cyamli@v1.1.7 generate golang -schema-path=cli.yaml -out-path=cli.gen.go -package=main
//go:generate go run github.com/Jumpaku/cyamli@v1.1.7 generate docs -schema-path=cli.yaml -out-path=docs.md -all -format=markdown
func main() {
	cli := NewCLI()
	cli.FUNC = func(subcmd []string, in CLI_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)

		fmt.Println(GetDoc(subcmd))
		return nil
	}
	cli.Version.FUNC = func(subcmd []string, in CLI_Version_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		fmt.Println("v0.0.2")
		return nil
	}
	cli.Validate.FUNC = func(subcmd []string, in CLI_Validate_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)
		inFile := readerOrStdin(in.Opt_InputPath)
		defer inFile.Close()

		s, err := io.ReadAll(inFile)
		panicOnError(err, "fail to read file")

		decodeResult := exprml.NewDecoder().Decode(&exprmlpb.DecodeInput{Yaml: string(s)})
		if decodeResult.IsError {
			return fmt.Errorf("fail to decode yaml: %s", decodeResult.ErrorMessage)
		}
		parseResult := exprml.NewParser().Parse(&exprmlpb.ParseInput{Value: decodeResult.Value})
		if parseResult.IsError {
			return fmt.Errorf("fail to parse AST: %v", parseResult.ErrorMessage)
		}

		return nil
	}
	cli.Evaluate.FUNC = func(subcmd []string, in CLI_Evaluate_Input, inErr error) error {
		panicOnInputError(inErr, subcmd)
		exitAfterHelp(in.Opt_Help, subcmd)

		var format exprmlpb.EncodeInput_Format
		switch in.Opt_Format {
		case "json":
			format = exprmlpb.EncodeInput_JSON
		case "yaml":
			format = exprmlpb.EncodeInput_YAML
		default:
			log.Panicf("format must be 'json' or 'yaml': %q\n", in.Opt_Format)
		}

		inFile := readerOrStdin(in.Opt_InputPath)
		defer inFile.Close()
		outFile := writerOrStdout(in.Opt_OutputPath)
		defer outFile.Close()

		s, err := io.ReadAll(inFile)
		panicOnError(err, "fail to read file")

		decodeResult := exprml.NewDecoder().Decode(&exprmlpb.DecodeInput{Yaml: string(s)})
		panicIf(decodeResult.IsError, "fail to decode yaml: %s", decodeResult.ErrorMessage)

		parseResult := exprml.NewParser().Parse(&exprmlpb.ParseInput{Value: decodeResult.Value})
		panicIf(parseResult.IsError, "fail to parse AST: %v", parseResult.ErrorMessage)

		evaluateResult := exprml.NewEvaluator(nil).EvaluateExpr(&exprmlpb.EvaluateInput{Expr: parseResult.Expr})
		if evaluateResult.Status != exprmlpb.EvaluateOutput_OK {
			log.Panicln(fmt.Errorf("fail to evaluate expression: status = %q: %q: %v",
				evaluateResult.Status.String(),
				exprml.Format(evaluateResult.ErrorPath),
				evaluateResult.ErrorMessage))
		}

		v := exprml.NewEncoder().Encode(&exprmlpb.EncodeInput{
			Format: format,
			Value:  evaluateResult.Value,
		})
		panicIf(v.IsError, "fail to encode value: %s", v.ErrorMessage)

		_, err = io.WriteString(outFile, v.Result)
		panicOnError(err, "fail to write file")

		return nil
	}
	if err := Run(cli, os.Args); err != nil {
		panic(err)
	}
}

func panicOnInputError(inErr error, subcmd []string) {
	if inErr != nil {
		fmt.Fprintln(os.Stderr, GetDoc(subcmd))
		log.Panicln(inErr)
	}
}
func panicOnError(err error, format string, args ...any) {
	if err != nil {
		log.Panicln(fmt.Errorf(format+": %w", append(append([]any{}, args...), err)))
	}
}
func panicIf(panicCond bool, format string, args ...any) {
	if panicCond {
		log.Panicln(fmt.Errorf(format, args...))
	}
}
func readerOrStdin(inFile string) io.ReadCloser {
	if inFile == "" {
		return io.NopCloser(os.Stdin)
	}
	f, err := os.Open(inFile)
	panicOnError(err, "fail to open file", inFile)
	return f
}
func writerOrStdout(outFile string) io.WriteCloser {
	if outFile == "" {
		return os.Stdout
	}
	f, err := os.Create(outFile)
	panicOnError(err, "fail to create file", outFile)
	return f

}
func exitAfterHelp(flag bool, subcmd []string) {
	if flag {
		fmt.Println(GetDoc(subcmd))
		os.Exit(0)
	}
}
