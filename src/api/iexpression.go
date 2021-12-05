package api

/*
abstract class Expr {
	// defineVisitor
	interface Visitor<R> {
		R visitBinaryExpr(Binary expr);
		R visitGroupingExpr(Grouping expr);
		// ...
	}

	// loop defineType
	static class Binary extends Expr {
		Binary(left, op, right) {
			this.left = left;
			this.op = op;
			this.right = right;
		}
		@override
		<R> R accept(Visitor<R> visitor) {
			return visitor.visitBinaryExpr(this);
		}

		// Field
		final left
		final op
		final right
	}

	static class Grouping extends Expr {
		Grouping(expr) {
			this.expr = expr
		}
		@override
		<R> R accept(Visitor<R> visitor) {
			return visitor.visitGroupingExpr(this);
		}

		// Field
		final expr
	}
	//...

	abstract <R> R accept(Visitor<R> visitor);
}
*/
