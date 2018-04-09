/*
 * Copyright (c) 2018 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dat

import (
	"context"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/topfreegames/extensions/jaeger"
)

// Trace wraps a Dat/PosgreSQL query and reports it to Jaeger
func Trace(ctx context.Context, statement string, next func() error) {
	var parent opentracing.SpanContext

	if ctx == nil {
		ctx = context.Background()
	}

	if span := opentracing.SpanFromContext(ctx); span != nil {
		parent = span.Context()
	}

	operation := strings.Fields(statement)[0]
	operationName := "SQL " + operation
	reference := opentracing.ChildOf(parent)
	tags := opentracing.Tags{
		"db.operation": operation,
		"db.type":      "postgres",
		"db.statement": statement,

		"span.kind": "client",
	}

	span := opentracing.StartSpan(operationName, reference, tags)
	defer span.Finish()
	defer jaeger.LogPanic(span)

	err := next()
	if err != nil {
		message := err.Error()
		jaeger.LogError(span, message)
	}
}
