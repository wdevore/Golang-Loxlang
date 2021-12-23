# Grammar

```
program       -> declaration* EOF ;
declaration   -> varDecl | statement;
varDecl       -> "var" IDENTIFIER ( "=" expression )? ";" ;

statement     -> exprStmt
                 | breakStmt
                 | continueStmt
                 | forStmt
                 | ifStmt
                 | printStmt
                 | whileStmt
                 | block ;
exprStmt      -> expression ";" ;
forStmt       -> "for" "(" ( varDecl | exprStmt | ";" )
                 expression? ";"
                 expression? ")" statement ;
ifStmt        -> "if" "(" expression ")" statement
                 ( "else" statement )? ;
printStmt     -> "print" expression ";" ;
whileStmt     -> "while" "(" expression ")" statement ;
breakStmt     -> "break" ";" ;
continueStmt  -> "continue" ";" ;

block         -> "{" declaration* "}" ;

expression    -> assignment ;
assignment    -> IDENTIFIER "=" assignment
                 | logic_or ;
logic_or      -> logic_and ( "or" logic_and )* ;
logic_and     -> equality ( "and" equality )* ;
equality      -> comparison ( ( "!=" | "==" ) comparison )* ;
comparison    -> term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term          -> factor ( ( "-" | "+" ) factor )* ;
factor        -> unary ( ( "/" | "*" ) unary )* ;
unary         -> ( "!" | "-" ) unary | primary ;
primary       -> "true" | "false" | "nil"
                 | NUMBER | STRING
                 | "(" expression ")"
                 | IDENTIFIER ;

```