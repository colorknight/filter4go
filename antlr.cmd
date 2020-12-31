alias antlr='java -jar /Users/tools/jar_repo/org/antlr/antlr-4.7.1-complete.jar'

antlr -Dlanguage=Go -o ./ -package filter -no-listener ./filter/filter.g4

antlr -Dlanguage=Go -o ./ -package oper -no-listener ./oper/oper.g4