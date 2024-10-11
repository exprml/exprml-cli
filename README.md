# exprml-cli

Interpreter for the ExprML language.

## Installation

```shell
go install github.com/exprml/exprml-cli@latest
```

## Usage


```shell
exprml-cli evaluate < input.yaml
```

where the `input.yaml` is as follows:
```yaml
# input.yaml
cat: [Hello, ", ", ExprML, "!"]
# => Hello, ExprML!
```


The CLI document is available at https://github.com/exprml/exprml-cli/blob/main/docs.md .

