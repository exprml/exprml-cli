# exprml-cli

A CLI interpreter for the ExprML language ( https://github.com/exprml/exprml-language ).

## Installation

```shell
go install github.com/exprml/exprml-cli@latest
```

## Usage

```shell
exprml-cli evaluate < input.yaml
```

```yaml
# input.yaml
cat: [Hello, ", ", ExprML, "!"]
# => Hello, ExprML!
```

The CLI document is available at https://github.com/exprml/exprml-cli/blob/main/docs.md .

