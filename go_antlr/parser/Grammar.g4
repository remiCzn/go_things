grammar Grammar;

// règles syntaxiques

program: statement* return;

statement: term ';';
return: term;

term:
	LIT										# Lit
	| term op = OP1 term					# BOp
	| term op = OP2 term					# BOp
	| 'ifz' term 'then' term 'else' term	# IfZ
	| VAR									# Var
	| 'let' var = VAR '=' term				# Let;

// règles lexicales
OP1: '*' | '/';
OP2: '+' | '-';
LIT: '0' | [1-9][0-9]*;
WS: [ \t\n\r]+ -> channel(HIDDEN);
VAR: [a-zA-Z_]+ [a-zA-Z0-9_]*;