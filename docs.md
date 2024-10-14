# exprml-cli (v0.0.2)

## exprml-cli

### Description

exprml command line interface

### Syntax

```shell
exprml-cli  [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Shows help message.  

### Subcommands

* evaluate:  
  Evaluates a YAML expression.  

* validate:  
  Validates a YAML file.  

* version:  
  Shows the version of the exprml-cli command.  


## exprml-cli evaluate

### Description

Evaluates a YAML expression.

### Syntax

```shell
exprml-cli evaluate [<option>]...
```

### Options

* `-format=<string>`, `-f=<string>`  (default=`"yaml"`):  
  Output format. One of `yaml` or `json`.  

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Show help message.  

* `-input-path=<string>`, `-i=<string>`  (default=`""`):  
  Input YAML file path. stdin is used if not provided.  

* `-output-path=<string>`, `-o=<string>`  (default=`""`):  
  Output file path. stdout is used if not provided.  


## exprml-cli validate

### Description

Validates a YAML file.

### Syntax

```shell
exprml-cli validate [<option>]...
```

### Options

* `-help[=<boolean>]`, `-h[=<boolean>]`  (default=`false`):  
  Show help message.  

* `-input-path=<string>`, `-i=<string>`  (default=`""`):  
  Input YAML file path. stdin is used if not provided.  


## exprml-cli version

### Description

Shows the version of the exprml-cli command.

### Syntax

```shell
exprml-cli version
```


