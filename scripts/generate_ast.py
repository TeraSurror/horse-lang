import sys


if len(sys.argv) < 2:
    print("Usage: python generate_ast.py <filename.go>")
    sys.exit(1)

file_name = sys.argv[1]

closed_braces = "{}"

types = [
    "Binary: left Expr, operator token.Token, right Expr",
    "Grouping: expression Expr",
    "Literal: value interface{}",
    "Unary: operator token.Token, right Expr",
]

with open(file_name, "w") as file:
    file.write("package ast\n")
    file.write("\n")
    file.write('import "horse-lang/src/token"\n')
    file.write("\n")
    file.write("type Expr interface {\n")
    file.write("\tassign(visitorExpr ExprVisitor) interface{}\n")
    file.write("}\n")
    file.write("\n")

    file.write("type ExprVisitor interface {\n")
    for type in types:
        name = type.split(":")[0].strip()
        file.write(f"\tVisit{name}Expr(expr *{name}) interface{closed_braces}\n")

    file.write("}\n")
    file.write("\n")

    for type in types:
        name = type.split(":")[0].strip()
        members = type.split(":")[1].split(", ")

        file.write(f"type {name} struct {'{'}\n")
        for member in members:
            member = member.strip()
            file.write(f"\t{member}\n")
        file.write("}\n")
        file.write("\n")

        file.write(
            f"func (expr *{name}) accept(visitor ExprVisitor) interface{closed_braces} {'{'}\n"
        )
        file.write(f"\treturn visitor.Visit{name}Expr(expr)\n")
        file.write("}\n")
        file.write("\n")
