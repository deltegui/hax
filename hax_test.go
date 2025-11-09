package hax_test

import (
	"hax"
	"net/http"
	"strings"
	"testing"
)

func strip(s string) string {
	r := strings.NewReplacer(
		//" ", "",
		"\t", "",
		"\n", "",
		"\r", "",
	)
	return r.Replace(s)
}

func render(tree hax.INode) string {
	renderer := hax.StringRenderer{}
	renderer.Render(tree)
	return renderer.String()
}

func checkExpected(t *testing.T, current, expected string) {
	if current != expected {
		t.Error("Expected result: ", current, " to be equal to ", expected)
	}
}

func matchTreeString(t *testing.T, tree hax.INode, str string) {
	result := render(tree)
	checkExpected(t, result, str)
}

func TestOneNode(t *testing.T) {
	t.Run("Should match one node", func(t *testing.T) {
		matchTreeString(
			t,
			hax.Div(),
			"<div></div>",
		)
	})

	t.Run("Should match one node with text", func(t *testing.T) {
		matchTreeString(
			t,
			hax.H1().Text("Hola"),
			"<h1>Hola</h1>",
		)
	})

	t.Run("Should match one node with attributes", func(t *testing.T) {
		matchTreeString(
			t,
			hax.A().Attribute("href", "deltegui.com").Text("Go to deltegui.com"),
			`<a href="deltegui.com">Go to deltegui.com</a>`,
		)
	})

	t.Run("Should match one node with classes", func(t *testing.T) {
		matchTreeString(
			t,
			hax.Button().Class("btn", "btn-primary").Text("I'm a button"),
			`<button class="btn btn-primary">I&#39;m a button</button>`,
		)
	})

	t.Run("Should match one node with style", func(t *testing.T) {
		matchTreeString(
			t,
			hax.Button().Style("background-color", "green").Style("border-radious", "50%").Text("Say hello!"),
			`<button style="background-color: green;border-radious: 50%;">Say hello!</button>`,
		)
	})

	t.Run("Should match one node with ID", func(t *testing.T) {
		matchTreeString(
			t,
			hax.Div().Id("container").Text("Hola"),
			`<div id="container">Hola</div>`,
		)
	})
}

func TestBody(t *testing.T) {
	t.Run("Should match one node with ID", func(t *testing.T) {
		matchTreeString(
			t,
			hax.Div().
				Id("container").
				Body(
					hax.H1().Class("title").Text("Demo"),
					hax.Form().
						Action("/endpoint/to/backend").
						Method(http.MethodGet).
						Body(
							hax.A().Href("www.google.es").Text("Go to google"),
							hax.Div().Class("form-control").Body(
								hax.Label().Text("Name").For("name-input"),
								hax.Input().Value("hola mundo").Placeholder("Nombre").Id("name-input"),
								hax.Output().For("name-input"),
							),
						),
				),
			strip(`
				<div id="container">
					<h1 class="title">Demo</h1>

					<form action="/endpoint/to/backend" method="GET">
						<a href="www.google.es">Go to google</a>
						<div class="form-control">
							<label for="name-input">Name</label>
							<input value="hola mundo" placeholder="Nombre" id="name-input"/>
							<output for="name-input"></output>
						</div>
					</form>
				</div>`),
		)
	})
}
