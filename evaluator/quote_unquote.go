package evaluator

import (
	"fmt"
	"monkey/ast"
	"monkey/object"
	"monkey/token"
)

/**
 * quote() 函数是一个入口，该函数会将其函数体内的所有语句中的
 * unquote() 函数的参数进行替换并求值；
 * 最后，返回一个 *object.Quote{Node: node} 结构体
 */
func quote(node ast.Node, env *object.Environment) object.Object {
	// fmt.Printf(">>>>> before evalUnquoteCalls: %#v \n", node)
	node = evalUnquoteCalls(node, env)
	// fmt.Printf(">>>>> after  evalUnquoteCalls: %#v \n", node)
	return &object.Quote{Node: node}
}


/**
 * 	使用 Modify 函数递归修改所有 unquote 标记的节点

 * 	将所有 unquote(x) 调用，替换为参数 x 对应的 ast 节点，然后对该 ast 节点求值。
 *
 *  说白了，unquote() 没有函数体，不是真正的函数调用，而是一个替换标记，
 *  凡是被 unquote 标记的 x，需要被替换为 env 中 "x" 这个名称对应的 ast 节点。


// 例如：
// 在求值过程中遇到函数参数（例如：标识符x，结构是 ast.Identifier)，
// Eval 函数 会在 evalEnv 中查找到 x 对应的 Quote 结构体
// （例如：整数字面量 object.Quote{Node: *ast.IntegerLiteral}  
// 或者 各种表达式 object.Quote{Node: *ast.Expression}），
// 最后，将求值结果即该 object.Quote 转化为 ast.Node 并返回。

 */
func evalUnquoteCalls(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquoteCall(node) {
			return node
		}

		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}
        fmt.Printf("call.Arguments[0]: %#v \n", call.Arguments[0])

		if len(call.Arguments) != 1 {
			return node
		}

		// 从 evalEnv 这个 map 中找到对应的参数
		// 例如：
		// let mm = macro(a) { quote(unquote(a)); };
		// mm(5)
		// 如果此时在替换 unquote(a)，则 call.Arguments[0] 就是：
		// &ast.Identifier{Token:token.Token{Type:"IDENT", Literal:"a"}, Value:"a"}
		// 而在 evalEnv 中 标识符a 对应的是 整数字面量 5。
		// 
		unquoted := Eval(call.Arguments[0], env)
		
        fmt.Printf("unquoted: %#v \n", unquoted)
		return convertObjectToASTNode(unquoted)
	})
}

func isUnquoteCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return callExpression.Function.TokenLiteral() == "unquote"
}

func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type: token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}	
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}

	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}

	case *object.Quote:
		return obj.Node
	
	default:
		return nil	
	}
}