# JSON EVAL 
JSON EVAL is simple json evaluator that may be simplify business rule. 

##Operator 
Operator supported 
println!("{}",
    table!(
        "{^:10:}" => "No", "{^:10:}" => "Operator";
        "1", "<", "2", "<=", "3", ">", "4", "||","5", "&&"
    ).format()
);
