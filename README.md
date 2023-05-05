# JSON EVAL 
JSON EVAL is simple json evaluator that may be simplify business rule. 

# Operator 
Operator supported 
| No.      | Operator  | Precedence |
| ---      | ---       | ---
| 1        | <         | 1 (left to right)
| 2        | <=        | 1 (left to right)
| 3        | >         | 1 (left to right)
| 4        | >=        | 1 (left to right)
| 5        | !=        | 2 (left to right)
| 6.       | ==        | 2 (left to right)  
| 6        | &&        | 3 (left to right)
| 7        | \|\|      | 4 (left to right)

# Evaluate
json-eval can evaluate json value with logical expression.

We have below json:

```json 
{
    "glossary": {
        "title": "example glossary",
        "GlossDiv": {
            "total": 1000
            "title": "S",
            "GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
                    "SortAs": "SGML",
                    "GlossTerm": "Standard Generalized Markup Language",
                    "Acronym": "SGML",
                    "Abbrev": "ISO 8879:1986",
                    "GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
                        "GlossSeeAlso": [
                            "GML",
                            "XML"
                        ]
                    },
                    "GlossSee": "markup"
                }
            }
        }
    }
}
```
expression: 

~~~
glossary.title == "example glossary" && glossary.total > 100
~~~
# Array of Index
json eval also has capability to extract array value from json. for example, 
~~~
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[0]
~~~

Result:
```go 
"GML"
```

# Example
Code snippets
```go
jsnStr := `
    {
        "glossary": {
            "title": "example glossary",
            "GlossDiv": {
                "total": 1000
            }
        }
    }    
`
expression := `glossary.title == "example glossary" && glossary.total > 100`
mapEval := jsoneval.NewJsonEvaluator()
result, err := mapEval.EvaluateJson(expression, jsonStr)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(result)
}
```

Result:
```go 
true 
```


