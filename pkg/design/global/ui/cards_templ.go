// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/onsonr/sonr/pkg/design/global/styles"

func Card(id string, size styles.Size) templ.Component {
	return renderCard(id, size.CardAttributes())
}

func renderCard(id string, attrs templ.Attributes) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `global/ui/cards.templ`, Line: 10, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><div class=\"w-full h-full\"><div class=\"row\"><div class=\"col-md-12 space-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfileCard() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative max-w-sm overflow-hidden bg-white border rounded-lg shadow-sm border-neutral-200/60\"><img src=\"https://cdn.devdojo.com/images/august2023/wallpaper.jpeg\" class=\"relative z-20 object-cover w-full h-32\"><div class=\"absolute top-0 z-50 flex items-center w-full mt-2 translate-y-24 px-7 -translate-x-0\"><div class=\"w-20 h-20 p-1 bg-white rounded-full\"><img src=\"https://cdn.devdojo.com/images/august2023/adam.jpeg\" class=\"w-full h-full rounded-full\"></div><a href=\"https://twitter.com/adamwathan\" target=\"_blank\" class=\"block mt-6 ml-2\"><h5 class=\"text-lg font-bold leading-none tracking-tight text-neutral-900\">Adam Wathan</h5><small class=\"block mt-1 text-sm font-medium leading-none text-neutral-500\">adamwathan</small></a> <button class=\"absolute right-0 inline-flex items-center justify-center w-auto px-5 mt-6 text-sm font-medium transition-colors duration-100 rounded-full h-9 mr-7 hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:opacity-50 bg-neutral-900 disabled:pointer-events-none hover:bg-neutral-800 text-neutral-100\"><span>Follow</span></button></div><div class=\"relative pb-6 p-7\"><p class=\"mt-12 mb-6 text-neutral-500 text-\">Creator of @tailwindcss. Listener of Slayer. Austin 3:16. BTW, Pines UI is super cool!</p><div class=\"flex items-center justify-between pr-2 text-neutral-500\"><div class=\"relative flex w-16\"><img src=\"https://cdn.devdojo.com/images/august2023/caleb.jpeg\" class=\"relative z-30 w-8 h-8 border-2 border-white rounded-full\"> <img src=\"https://cdn.devdojo.com/images/august2023/taylor.jpeg\" class=\"z-20 w-8 h-8 -translate-x-4 border-2 border-white rounded-full\"> <img src=\"https://cdn.devdojo.com/images/august2023/adam.jpeg\" class=\"z-10 w-8 h-8 border-2 border-white rounded-full -translate-x-7\"></div><a href=\"https://twitter.com/adamwathan/following\" target=\"_blank\" class=\"text-sm hover:underline\"><strong class=\"text-neutral-800\">673</strong> Following</a> <a href=\"https://twitter.com/adamwathan/followers\" target=\"_blank\" class=\"text-sm hover:underline\"><strong class=\"text-neutral-800\">168.6K</strong> Followers</a></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
